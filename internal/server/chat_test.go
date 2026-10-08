package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func postChat(t *testing.T, chat ChatStore, body string) *httptest.ResponseRecorder {
	t.Helper()
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/chat", strings.NewReader(body))
	newTestHandler(chat).ServeHTTP(rr, req)
	return rr
}

func TestChatHappyPathNewConversation(t *testing.T) {
	f := &fakeChatStore{ensureID: "22222222-2222-2222-2222-222222222222"}
	rr := postChat(t, f, `{"conversation_id": null, "question": "what is a mode?"}`)

	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200 (body %s)", rr.Code, rr.Body.String())
	}
	var resp chatResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if resp.ConversationID != f.ensureID {
		t.Fatalf("conversation_id = %q, want %q", resp.ConversationID, f.ensureID)
	}
	if resp.Answer != stubAnswer {
		t.Fatalf("answer = %q, want stub", resp.Answer)
	}
	if resp.Sources == nil || len(resp.Sources) != 0 {
		t.Fatalf("sources = %v, want non-nil empty slice", resp.Sources)
	}
	if resp.TraceID != "" {
		t.Fatalf("trace_id = %q, want empty", resp.TraceID)
	}
	// Persistence: one EnsureConversation("") and two messages, user verbatim first.
	if len(f.ensureCalls) != 1 || f.ensureCalls[0] != "" {
		t.Fatalf("ensureCalls = %v, want one empty", f.ensureCalls)
	}
	if len(f.messages) != 2 {
		t.Fatalf("messages = %v, want 2", f.messages)
	}
	if f.messages[0] != (fakeMsg{f.ensureID, "user", "what is a mode?"}) {
		t.Fatalf("messages[0] = %v, want verbatim user", f.messages[0])
	}
	if f.messages[1].role != "assistant" || f.messages[1].content != stubAnswer {
		t.Fatalf("messages[1] = %v, want assistant stub", f.messages[1])
	}
}

func TestChatPassesThroughProvidedConversationID(t *testing.T) {
	f := &fakeChatStore{}
	id := "33333333-3333-3333-3333-333333333333"
	rr := postChat(t, f, `{"conversation_id": "`+id+`", "question": "hi"}`)
	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rr.Code)
	}
	var resp chatResponse
	_ = json.Unmarshal(rr.Body.Bytes(), &resp)
	if resp.ConversationID != id {
		t.Fatalf("conversation_id = %q, want %q", resp.ConversationID, id)
	}
	if len(f.ensureCalls) != 1 || f.ensureCalls[0] != id {
		t.Fatalf("ensureCalls = %v, want [%q]", f.ensureCalls, id)
	}
}

func TestChatValidation(t *testing.T) {
	cases := []struct {
		name, body, wantCode string
	}{
		{"malformed json", `{not json`, "invalid_json"},
		{"empty question", `{"question": ""}`, "invalid_request"},
		{"whitespace question", `{"question": "   "}`, "invalid_request"},
		{"too long", `{"question": "` + strings.Repeat("a", 4001) + `"}`, "question_too_long"},
		{"bad conversation id", `{"conversation_id": "not-a-uuid", "question": "hi"}`, "invalid_conversation_id"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			rr := postChat(t, &fakeChatStore{}, tc.body)
			if rr.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want 400", rr.Code)
			}
			var env errorEnvelope
			if err := json.Unmarshal(rr.Body.Bytes(), &env); err != nil {
				t.Fatalf("decode: %v", err)
			}
			if env.Error.Code != tc.wantCode {
				t.Fatalf("code = %q, want %q", env.Error.Code, tc.wantCode)
			}
		})
	}
}

func TestChatStoreErrorIs500(t *testing.T) {
	f := &fakeChatStore{ensureErr: errTest}
	rr := postChat(t, f, `{"question": "hi"}`)
	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, want 500", rr.Code)
	}
	var env errorEnvelope
	_ = json.Unmarshal(rr.Body.Bytes(), &env)
	if env.Error.Code != "internal" {
		t.Fatalf("code = %q, want internal", env.Error.Code)
	}
}

var errTest = errTestType("boom")

type errTestType string

func (e errTestType) Error() string { return string(e) }
