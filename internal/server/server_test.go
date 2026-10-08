package server

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/mharner33/guitar-chat/internal/config"
)

// fakeChatStore records calls and returns injected values/errors.
// Shared across server_test.go and chat_test.go.
type fakeChatStore struct {
	ensureID  string // returned when the requested id is empty; defaults below
	ensureErr error
	addErr    error

	ensureCalls []string
	messages    []fakeMsg
}

type fakeMsg struct{ conv, role, content string }

func (f *fakeChatStore) EnsureConversation(_ context.Context, id string) (string, error) {
	f.ensureCalls = append(f.ensureCalls, id)
	if f.ensureErr != nil {
		return "", f.ensureErr
	}
	if id != "" {
		return id, nil
	}
	if f.ensureID != "" {
		return f.ensureID, nil
	}
	return "11111111-1111-1111-1111-111111111111", nil
}

func (f *fakeChatStore) AddMessage(_ context.Context, conv, role, content string) error {
	if f.addErr != nil {
		return f.addErr
	}
	f.messages = append(f.messages, fakeMsg{conv, role, content})
	return nil
}

func newTestHandler(chat ChatStore) http.Handler {
	assets := fstest.MapFS{
		"index.html": &fstest.MapFile{Data: []byte("<!doctype html><title>guitar-chat</title>")},
	}
	return New(config.Config{}, chat, assets)
}

func TestHealthz(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	newTestHandler(&fakeChatStore{}).ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rr.Code)
	}
	var body map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if body["status"] != "ok" {
		t.Fatalf("body = %v, want status=ok", body)
	}
}

func TestStaticIndexServed(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	newTestHandler(&fakeChatStore{}).ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rr.Code)
	}
	if got := rr.Body.String(); !strings.Contains(got, "guitar-chat") {
		t.Fatalf("index body = %q, want it to contain marker", got)
	}
}
