# Phase 1 — Server + UI + Chat Stub Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Stand up the chi HTTP server, the embedded two-pane web UI, and a wired-but-stubbed `POST /api/chat` that validates input, persists the conversation + verbatim user message, and returns the real response shape with a placeholder answer.

**Architecture:** A single `http.Handler` built in `internal/server` (chi router + middleware) wires `GET /healthz`, `POST /api/chat`, and the `go:embed`-served UI. The chat handler depends on a small consumer-defined `ChatStore` interface (faked in tests); `*store.Store` implements it with pgx SQL. `cmd/api` runs it with timeouts and graceful shutdown. No RAG, LLM, embeddings, or observability — those are later phases.

**Tech Stack:** Go 1.26, `github.com/go-chi/chi/v5`, `pgx/v5`, `log/slog`, stdlib `net/http`/`embed`/`testing`/`httptest`. Vanilla HTML/CSS/JS frontend (no framework, no build step).

## Global Constraints

- Go 1.26.2 (module floor); `agent.md` floor is "Go 1.23+". Standard library first; every new dependency must be justified (chi is the sanctioned router per `agent.md`).
- Routing: `chi`. DB: `pgx/v5`. Logging: `log/slog` JSON. No ORMs, no codegen.
- Errors wrap with `%w` and are returned; no panics outside `main`; handle/log at the boundary.
- Thread `ctx` through every call; the answer path has a 30s deadline.
- All config is read once in `internal/config`; add any new variable to `.env.example` (this plan adds none).
- SQL lives in `internal/store`; migrations are forward-only and never edited.
- Frontend: no frameworks, no npm, no build; `web/app.js` uses `fetch` only.
- Tests: table-driven, fakes for store; **no network/DB in unit tests** (the one store integration test is `DATABASE_URL`-gated and skips when unset).
- Never bypass RAG (the stub must not call an LLM or assert theory) and never log secrets.
- Demo scope: no auth, no streaming, no k8s, no multi-agent.
- Definition of Done: `go vet ./...`, `staticcheck ./...`, `go test ./...`, `gofmt -l .` all clean.
- Code stays MIT (root `LICENSE`).

---

## File Structure

- Create `internal/store/chat.go` — `EnsureConversation`, `AddMessage` (pgx SQL).
- Create `internal/store/chat_test.go` — `DATABASE_URL`-gated integration test.
- Create `internal/server/respond.go` — JSON + error-envelope write helpers.
- Create `internal/server/server.go` — `New(...)`, chi router, middleware, `/healthz`, `ChatStore` interface, UI mount.
- Create `internal/server/chat.go` — `POST /api/chat` handler, request/response types, validation, stub answer.
- Create `internal/server/server_test.go` — healthz + static-serving + error-helper tests, shared fake store.
- Create `internal/server/chat_test.go` — `/api/chat` handler tests.
- Create `web/web.go` — `//go:embed` the three asset files.
- Create `web/index.html`, `web/app.js`, `web/style.css` — two-pane UI.
- Modify `cmd/api/main.go` — replace the Phase-1 stub comment with server construction + `http.Server` + graceful shutdown.
- Modify `go.mod` / `go.sum` — add chi (via `go get`).

---

## Task 1: Store persistence (`internal/store/chat.go`)

**Files:**
- Create: `internal/store/chat.go`
- Test: `internal/store/chat_test.go`

**Interfaces:**
- Consumes: `*store.Store` + `store.Store.Pool()` (`*pgxpool.Pool`) from `internal/store/store.go`.
- Produces:
  - `func (s *Store) EnsureConversation(ctx context.Context, id string) (string, error)` — empty `id` mints a new conversation (`gen_random_uuid()`), returns the id; non-empty `id` is inserted `ON CONFLICT DO NOTHING` and returned unchanged.
  - `func (s *Store) AddMessage(ctx context.Context, conversationID, role, content string) error` — inserts one row into `messages` (`trace_id` left NULL).

- [ ] **Step 1: Write the failing test**

Create `internal/store/chat_test.go`:

```go
package store

import (
	"context"
	"os"
	"testing"
	"time"
)

// Integration test: requires a live DATABASE_URL with migrations applied.
// Skips when unset so `go test ./...` stays network-free by default.
func TestEnsureConversationAndAddMessage(t *testing.T) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Skip("DATABASE_URL not set; skipping store integration test")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	st, err := Open(ctx, dsn)
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	defer st.Close()

	// Mint a new conversation.
	id, err := st.EnsureConversation(ctx, "")
	if err != nil {
		t.Fatalf("EnsureConversation(new): %v", err)
	}
	if id == "" {
		t.Fatal("EnsureConversation(new) returned empty id")
	}

	// Clean up whatever we insert, regardless of outcome.
	t.Cleanup(func() {
		_, _ = st.Pool().Exec(context.Background(), `DELETE FROM messages WHERE conversation_id = $1`, id)
		_, _ = st.Pool().Exec(context.Background(), `DELETE FROM conversations WHERE id = $1`, id)
	})

	// Idempotent for an existing id.
	again, err := st.EnsureConversation(ctx, id)
	if err != nil {
		t.Fatalf("EnsureConversation(existing): %v", err)
	}
	if again != id {
		t.Fatalf("EnsureConversation(existing) = %q, want %q", again, id)
	}

	if err := st.AddMessage(ctx, id, "user", "test question"); err != nil {
		t.Fatalf("AddMessage(user): %v", err)
	}
	if err := st.AddMessage(ctx, id, "assistant", "test answer"); err != nil {
		t.Fatalf("AddMessage(assistant): %v", err)
	}

	var count int
	if err := st.Pool().QueryRow(ctx,
		`SELECT count(*) FROM messages WHERE conversation_id = $1`, id).Scan(&count); err != nil {
		t.Fatalf("count messages: %v", err)
	}
	if count != 2 {
		t.Fatalf("message count = %d, want 2", count)
	}
}
```

- [ ] **Step 2: Run test to verify it fails**

Run: `go test ./internal/store/ -run TestEnsureConversationAndAddMessage`
Expected: compile failure — `st.EnsureConversation`/`st.AddMessage` undefined.

- [ ] **Step 3: Write minimal implementation**

Create `internal/store/chat.go`:

```go
package store

import (
	"context"
	"fmt"
)

// EnsureConversation returns the canonical conversation id. An empty id mints a
// new conversation (gen_random_uuid()); a non-empty id is created if absent and
// returned unchanged, so a client-held id (e.g. from localStorage) works.
func (s *Store) EnsureConversation(ctx context.Context, id string) (string, error) {
	if id == "" {
		var newID string
		if err := s.pool.QueryRow(ctx,
			`INSERT INTO conversations (id) VALUES (gen_random_uuid()) RETURNING id`,
		).Scan(&newID); err != nil {
			return "", fmt.Errorf("insert conversation: %w", err)
		}
		return newID, nil
	}
	if _, err := s.pool.Exec(ctx,
		`INSERT INTO conversations (id) VALUES ($1) ON CONFLICT (id) DO NOTHING`, id,
	); err != nil {
		return "", fmt.Errorf("ensure conversation %s: %w", id, err)
	}
	return id, nil
}

// AddMessage stores one chat message. role must be 'user' or 'assistant'
// (enforced by a CHECK constraint). trace_id is left NULL until Phase 3.
func (s *Store) AddMessage(ctx context.Context, conversationID, role, content string) error {
	if _, err := s.pool.Exec(ctx,
		`INSERT INTO messages (conversation_id, role, content) VALUES ($1, $2, $3)`,
		conversationID, role, content,
	); err != nil {
		return fmt.Errorf("add message: %w", err)
	}
	return nil
}
```

- [ ] **Step 4: Run tests to verify they pass**

Run (unit mode, no DB): `go test ./internal/store/` → Expected: PASS with the integration test SKIPped.
Run (integration): `set -a; source .env; set +a; go test ./internal/store/ -run TestEnsureConversationAndAddMessage -v` → Expected: PASS.
Also: `gofmt -l internal/store/ ` prints nothing; `go vet ./internal/store/` clean.

- [ ] **Step 5: Commit**

```bash
git add internal/store/chat.go internal/store/chat_test.go
git commit -m "feat(store): conversation + message persistence (EnsureConversation, AddMessage)

Co-Authored-By: Claude Opus 4.8 (1M context) <noreply@anthropic.com>"
```

---

## Task 2: Server foundation — router, middleware, /healthz, ChatStore (`internal/server`)

**Files:**
- Create: `internal/server/respond.go`
- Create: `internal/server/server.go`
- Create: `internal/server/server_test.go`
- Modify: `go.mod`, `go.sum` (add chi)

**Interfaces:**
- Consumes: `config.Config` from `internal/config`.
- Produces:
  - `type ChatStore interface { EnsureConversation(ctx context.Context, id string) (string, error); AddMessage(ctx context.Context, conversationID, role, content string) error }`
  - `func New(cfg config.Config, chat ChatStore, assets fs.FS) http.Handler`
  - `func writeJSON(w http.ResponseWriter, status int, v any)` and `func writeError(w http.ResponseWriter, status int, code, message string)`
  - A shared `fakeChatStore` test helper (in `server_test.go`) used by Task 3's tests too.

- [ ] **Step 1: Add the chi dependency**

Run: `go get github.com/go-chi/chi/v5@latest`
Expected: `go.mod` now requires `github.com/go-chi/chi/v5`; `go.sum` updated.

- [ ] **Step 2: Write the failing tests**

Create `internal/server/server_test.go`:

```go
package server

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
	if got := rr.Body.String(); !contains(got, "guitar-chat") {
		t.Fatalf("index body = %q, want it to contain marker", got)
	}
}

func contains(s, sub string) bool {
	return len(s) >= len(sub) && (s == sub || indexOf(s, sub) >= 0)
}
func indexOf(s, sub string) int {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}
```

- [ ] **Step 3: Run tests to verify they fail**

Run: `go test ./internal/server/`
Expected: compile failure — `New`, `ChatStore` undefined.

- [ ] **Step 4: Write `respond.go`**

Create `internal/server/respond.go`:

```go
package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type errorEnvelope struct {
	Error errorBody `json:"error"`
}

type errorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// writeJSON writes v as a JSON response with the given status.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		slog.Error("write json response", "err", err)
	}
}

// writeError writes the standard {"error":{"code","message"}} envelope.
func writeError(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, errorEnvelope{Error: errorBody{Code: code, Message: message}})
}
```

- [ ] **Step 5: Write `server.go`**

Create `internal/server/server.go`:

```go
package server

import (
	"context"
	"io/fs"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/mharner33/guitar-chat/internal/config"
)

// ChatStore is the persistence the chat handler needs. Declared here (the
// consumer) so it can be faked in tests; *store.Store implements it.
type ChatStore interface {
	EnsureConversation(ctx context.Context, id string) (string, error)
	AddMessage(ctx context.Context, conversationID, role, content string) error
}

type api struct {
	chat ChatStore
}

// New builds the HTTP handler: middleware, health, chat, and the embedded UI.
// cfg is carried for the seam that later phases (retrieval/history knobs) need.
func New(cfg config.Config, chat ChatStore, assets fs.FS) http.Handler {
	_ = cfg // reserved for Phase 2+
	a := &api{chat: chat}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(accessLog)
	r.Use(middleware.Timeout(30 * time.Second))

	r.Get("/healthz", handleHealth)
	r.Post("/api/chat", a.handleChat)
	r.Handle("/*", http.FileServerFS(assets))

	return r
}

func handleHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// accessLog emits one slog-JSON line per request.
func accessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)
		slog.Info("http request",
			"method", r.Method,
			"path", r.URL.Path,
			"status", ww.Status(),
			"bytes", ww.BytesWritten(),
			"duration_ms", time.Since(start).Milliseconds(),
			"request_id", middleware.GetReqID(r.Context()),
		)
	})
}
```

> Note: `a.handleChat` is defined in Task 3 (`chat.go`, same package). The package will not compile until Task 3 adds it, so Steps 6–7 below run after `chat.go` exists. If implementing strictly task-by-task, add a temporary `func (a *api) handleChat(w http.ResponseWriter, r *http.Request) { writeError(w, http.StatusNotImplemented, "not_implemented", "chat arrives in Task 3") }` in `server.go`, then delete it when Task 3 adds the real one. (Subagent-driven execution: implement Task 2 and Task 3 as a pair.)

- [ ] **Step 6: Run tests to verify they pass**

Run: `go test ./internal/server/ -run 'TestHealthz|TestStaticIndexServed'`
Expected: PASS.

- [ ] **Step 7: Commit**

```bash
git add go.mod go.sum internal/server/respond.go internal/server/server.go internal/server/server_test.go
git commit -m "feat(server): chi router, middleware, /healthz, ChatStore seam

Co-Authored-By: Claude Opus 4.8 (1M context) <noreply@anthropic.com>"
```

---

## Task 3: `POST /api/chat` handler (`internal/server/chat.go`)

**Files:**
- Create: `internal/server/chat.go`
- Test: `internal/server/chat_test.go`

**Interfaces:**
- Consumes: `api.chat` (`ChatStore`), `writeJSON`/`writeError` (Task 2), `fakeChatStore` (Task 2 test helper).
- Produces: `func (a *api) handleChat(w http.ResponseWriter, r *http.Request)` and the const `stubAnswer`.

- [ ] **Step 1: Write the failing tests**

Create `internal/server/chat_test.go`:

```go
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
```

- [ ] **Step 2: Run tests to verify they fail**

Run: `go test ./internal/server/ -run TestChat`
Expected: compile failure — `chatResponse`, `stubAnswer`, `handleChat` undefined. (If a temporary `handleChat` stub exists from Task 2 Step 5's note, delete it now.)

- [ ] **Step 3: Write `chat.go`**

Create `internal/server/chat.go`:

```go
package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"regexp"
	"strings"
)

const (
	maxQuestionRunes = 4000
	// stubAnswer is a fixed placeholder: it calls no LLM and asserts no theory,
	// honoring the "never bypass RAG" guardrail. Real answers arrive in Phase 2.
	stubAnswer = "The answer pipeline isn't wired up yet — grounded answers arrive in Phase 2."
)

var uuidRE = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)

type chatRequest struct {
	ConversationID *string `json:"conversation_id"`
	Question       string  `json:"question"`
}

type source struct {
	N       int     `json:"n"`
	Title   string  `json:"title"`
	Section string  `json:"section"`
	Score   float64 `json:"score"`
}

type chatResponse struct {
	ConversationID string   `json:"conversation_id"`
	Answer         string   `json:"answer"`
	Sources        []source `json:"sources"`
	TraceID        string   `json:"trace_id"`
}

func (a *api) handleChat(w http.ResponseWriter, r *http.Request) {
	var req chatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid_json", "request body is not valid JSON")
		return
	}

	question := strings.TrimSpace(req.Question)
	if question == "" {
		writeError(w, http.StatusBadRequest, "invalid_request", "question is required")
		return
	}
	if len([]rune(question)) > maxQuestionRunes {
		writeError(w, http.StatusBadRequest, "question_too_long", "question exceeds the maximum length")
		return
	}

	id := ""
	if req.ConversationID != nil && *req.ConversationID != "" {
		if !uuidRE.MatchString(*req.ConversationID) {
			writeError(w, http.StatusBadRequest, "invalid_conversation_id", "conversation_id must be a UUID")
			return
		}
		id = *req.ConversationID
	}

	ctx := r.Context()
	convID, err := a.chat.EnsureConversation(ctx, id)
	if err != nil {
		slog.Error("ensure conversation", "err", err)
		writeError(w, http.StatusInternalServerError, "internal", "could not process the request")
		return
	}
	if err := a.chat.AddMessage(ctx, convID, "user", question); err != nil {
		slog.Error("add user message", "err", err)
		writeError(w, http.StatusInternalServerError, "internal", "could not process the request")
		return
	}
	if err := a.chat.AddMessage(ctx, convID, "assistant", stubAnswer); err != nil {
		slog.Error("add assistant message", "err", err)
		writeError(w, http.StatusInternalServerError, "internal", "could not process the request")
		return
	}

	writeJSON(w, http.StatusOK, chatResponse{
		ConversationID: convID,
		Answer:         stubAnswer,
		Sources:        []source{},
		TraceID:        "",
	})
}
```

- [ ] **Step 4: Run tests to verify they pass**

Run: `go test ./internal/server/`
Expected: PASS (all healthz, static, and chat tests).
Also: `gofmt -l internal/server/` prints nothing; `go vet ./internal/server/` clean.

- [ ] **Step 5: Commit**

```bash
git add internal/server/chat.go internal/server/chat_test.go
git commit -m "feat(server): POST /api/chat stub — validate, persist, placeholder answer

Co-Authored-By: Claude Opus 4.8 (1M context) <noreply@anthropic.com>"
```

---

## Task 4: Embedded web UI (`web/`)

**Files:**
- Create: `web/web.go`, `web/index.html`, `web/app.js`, `web/style.css`
- Test: `web/web_test.go`

**Interfaces:**
- Produces: `web.Assets` (`embed.FS`) containing `index.html`, `app.js`, `style.css` — consumed by `cmd/api` in Task 5 and already mountable by `server.New` (Task 2).

- [ ] **Step 1: Write the failing test**

Create `web/web_test.go`:

```go
package web

import (
	"io/fs"
	"testing"
)

func TestAssetsEmbedded(t *testing.T) {
	for _, name := range []string{"index.html", "app.js", "style.css"} {
		b, err := fs.ReadFile(Assets, name)
		if err != nil {
			t.Fatalf("read %s: %v", name, err)
		}
		if len(b) == 0 {
			t.Fatalf("%s is empty", name)
		}
	}
}
```

- [ ] **Step 2: Run test to verify it fails**

Run: `go test ./web/`
Expected: compile failure — `Assets` undefined (and no Go file in `web/`).

- [ ] **Step 3: Create the asset files, then `web.go`**

Create `web/index.html`:

```html
<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <title>Guitar Theory Chat</title>
  <link rel="stylesheet" href="/style.css" />
</head>
<body>
  <div id="app">
    <aside id="sidebar">
      <div class="sidebar-head">
        <h2>Recent</h2>
        <button id="clear-recent" type="button" title="Clear recent queries">Clear</button>
      </div>
      <ul id="recent-list"></ul>
    </aside>
    <main id="chat">
      <header><h1>Guitar Theory Chat</h1></header>
      <div id="thread" aria-live="polite"></div>
      <form id="composer">
        <textarea id="question" rows="2" placeholder="Ask a guitar theory question…" required></textarea>
        <button id="send" type="submit">Send</button>
      </form>
    </main>
  </div>
  <script src="/app.js"></script>
</body>
</html>
```

Create `web/style.css`:

```css
* { box-sizing: border-box; }
body { margin: 0; font: 15px/1.5 system-ui, sans-serif; color: #1a1a1a; background: #f5f5f7; }
#app { display: flex; height: 100vh; }
#sidebar { width: 240px; flex: 0 0 240px; border-right: 1px solid #ddd; background: #fff; display: flex; flex-direction: column; }
.sidebar-head { display: flex; align-items: center; justify-content: space-between; padding: 12px 14px; border-bottom: 1px solid #eee; }
.sidebar-head h2 { font-size: 13px; text-transform: uppercase; letter-spacing: .04em; color: #666; margin: 0; }
#clear-recent { font-size: 12px; background: none; border: none; color: #888; cursor: pointer; }
#clear-recent:hover { color: #333; }
#recent-list { list-style: none; margin: 0; padding: 6px; overflow-y: auto; }
#recent-list li { margin: 0; }
#recent-list button { width: 100%; text-align: left; border: none; background: none; padding: 8px 10px; border-radius: 6px; cursor: pointer; font-size: 13px; color: #333; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }
#recent-list button:hover { background: #f0f0f3; }
#chat { flex: 1; display: flex; flex-direction: column; min-width: 0; }
#chat header { padding: 14px 20px; border-bottom: 1px solid #ddd; background: #fff; }
#chat header h1 { font-size: 18px; margin: 0; }
#thread { flex: 1; overflow-y: auto; padding: 20px; display: flex; flex-direction: column; gap: 12px; }
.msg { max-width: 72%; padding: 10px 14px; border-radius: 12px; white-space: pre-wrap; }
.msg.user { align-self: flex-end; background: #2563eb; color: #fff; }
.msg.assistant { align-self: flex-start; background: #fff; border: 1px solid #e3e3e6; }
.msg.error { align-self: center; background: #fde8e8; color: #9b1c1c; border: 1px solid #f5c2c2; font-size: 13px; }
.sources { margin-top: 8px; display: flex; flex-wrap: wrap; gap: 6px; }
.chip { font-size: 12px; background: #eef2ff; color: #3730a3; border-radius: 999px; padding: 2px 8px; }
#composer { display: flex; gap: 8px; padding: 14px 20px; border-top: 1px solid #ddd; background: #fff; }
#question { flex: 1; resize: none; padding: 8px 10px; border: 1px solid #ccc; border-radius: 8px; font: inherit; }
#send { padding: 0 18px; border: none; border-radius: 8px; background: #2563eb; color: #fff; font-weight: 600; cursor: pointer; }
#send:disabled, #question:disabled { opacity: .6; cursor: not-allowed; }
@media (max-width: 640px) {
  #app { flex-direction: column; height: auto; min-height: 100vh; }
  #sidebar { width: 100%; flex: none; max-height: 30vh; border-right: none; border-bottom: 1px solid #ddd; }
}
```

Create `web/app.js`:

```js
const RECENT_KEY = "gc_recent";
const CONV_KEY = "gc_conversation";
const RECENT_MAX = 20;

const thread = document.getElementById("thread");
const form = document.getElementById("composer");
const input = document.getElementById("question");
const sendBtn = document.getElementById("send");
const recentList = document.getElementById("recent-list");
const clearBtn = document.getElementById("clear-recent");

function getRecent() {
  try { return JSON.parse(localStorage.getItem(RECENT_KEY)) || []; }
  catch { return []; }
}
function setRecent(list) { localStorage.setItem(RECENT_KEY, JSON.stringify(list)); }

function pushRecent(q) {
  let list = getRecent().filter((x) => x !== q);
  list.unshift(q);
  list = list.slice(0, RECENT_MAX);
  setRecent(list);
  renderRecent();
}

function renderRecent() {
  const list = getRecent();
  recentList.innerHTML = "";
  for (const q of list) {
    const li = document.createElement("li");
    const btn = document.createElement("button");
    btn.type = "button";
    btn.textContent = q;
    btn.title = q;
    btn.addEventListener("click", () => { input.value = q; input.focus(); });
    li.appendChild(btn);
    recentList.appendChild(li);
  }
}

function addMessage(role, text, sources) {
  const el = document.createElement("div");
  el.className = "msg " + role;
  el.textContent = text;
  if (sources && sources.length) {
    const wrap = document.createElement("div");
    wrap.className = "sources";
    for (const s of sources) {
      const chip = document.createElement("span");
      chip.className = "chip";
      chip.textContent = `[${s.n}] ${s.title} · ${s.section}`;
      wrap.appendChild(chip);
    }
    el.appendChild(wrap);
  }
  thread.appendChild(el);
  thread.scrollTop = thread.scrollHeight;
}

function setLoading(on) { input.disabled = on; sendBtn.disabled = on; }

form.addEventListener("submit", async (e) => {
  e.preventDefault();
  const question = input.value.trim();
  if (!question) return;

  addMessage("user", question);
  pushRecent(question);
  input.value = "";
  setLoading(true);

  try {
    const res = await fetch("/api/chat", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        conversation_id: localStorage.getItem(CONV_KEY) || null,
        question,
      }),
    });
    const data = await res.json();
    if (!res.ok) {
      addMessage("error", (data.error && data.error.message) || "Request failed");
      return;
    }
    if (data.conversation_id) localStorage.setItem(CONV_KEY, data.conversation_id);
    addMessage("assistant", data.answer, data.sources);
  } catch (err) {
    addMessage("error", "Network error: " + err.message);
  } finally {
    setLoading(false);
    input.focus();
  }
});

clearBtn.addEventListener("click", () => { setRecent([]); renderRecent(); });

renderRecent();
```

Create `web/web.go`:

```go
// Package web holds the embedded static assets for the chat UI, served by the
// api binary via go:embed (no framework, no build step).
package web

import "embed"

//go:embed index.html app.js style.css
var Assets embed.FS
```

- [ ] **Step 4: Run test to verify it passes**

Run: `go test ./web/`
Expected: PASS.
Also: `gofmt -l web/` prints nothing; `go vet ./web/` clean.

- [ ] **Step 5: Commit**

```bash
git add web/
git commit -m "feat(web): embedded two-pane chat UI (chat window + recent-queries sidebar)

Co-Authored-By: Claude Opus 4.8 (1M context) <noreply@anthropic.com>"
```

---

## Task 5: Wire and run the server (`cmd/api/main.go`)

**Files:**
- Modify: `cmd/api/main.go` (replace the `// Phase 1 wires the chi server and routes here` block and its `return nil`)

**Interfaces:**
- Consumes: `server.New` (Task 2), `web.Assets` (Task 4), `*store.Store` (implements `server.ChatStore` via Task 1).

- [ ] **Step 1: Replace the end of `run()`**

In `cmd/api/main.go`, replace these two lines:

```go
	// Phase 1 wires the chi server and routes here.
	return nil
```

with:

```go
	srv := &http.Server{
		Addr:              cfg.Addr,
		Handler:           server.New(cfg, st, web.Assets),
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      35 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	sigCtx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	errCh := make(chan error, 1)
	go func() {
		slog.Info("http server listening", "addr", cfg.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		return fmt.Errorf("http server: %w", err)
	case <-sigCtx.Done():
		slog.Info("shutdown signal received")
		stop()
		shutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(shutCtx); err != nil {
			return fmt.Errorf("graceful shutdown: %w", err)
		}
		return nil
	}
```

- [ ] **Step 2: Update the import block**

Set `cmd/api/main.go`'s imports to:

```go
import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mharner33/guitar-chat/internal/config"
	"github.com/mharner33/guitar-chat/internal/server"
	"github.com/mharner33/guitar-chat/internal/store"
	"github.com/mharner33/guitar-chat/migrations"
	"github.com/mharner33/guitar-chat/web"
)
```

- [ ] **Step 3: Build and vet**

Run: `go build ./... && go vet ./... && gofmt -l .`
Expected: build succeeds, vet clean, `gofmt -l .` prints nothing.

- [ ] **Step 4: Full test suite**

Run: `go test ./...`
Expected: PASS (store integration test SKIPs without `DATABASE_URL`).

- [ ] **Step 5: Manual end-to-end verification (against Supabase)**

```bash
set -a; source .env; set +a
go run ./cmd/api &   # note the PID
sleep 1
curl -s localhost:8080/healthz                       # -> {"status":"ok"}
curl -s localhost:8080/ | grep -o '<title>[^<]*'     # -> <title>Guitar Theory Chat
curl -s -X POST localhost:8080/api/chat \
  -H 'Content-Type: application/json' \
  -d '{"conversation_id":null,"question":"what is a mode?"}'
# -> {"conversation_id":"<uuid>","answer":"The answer pipeline isn't wired up yet …","sources":[],"trace_id":""}
curl -s -X POST localhost:8080/api/chat -H 'Content-Type: application/json' -d '{"question":""}'
# -> 400 {"error":{"code":"invalid_request", ...}}
kill %1
```

Then confirm persistence in the DB:

```bash
psql "$DATABASE_URL" -tAc "select role, left(content,30) from messages order by id desc limit 2;"
# -> assistant | The answer pipeline isn't wir
#    user      | what is a mode?
```

Also open `http://localhost:8080/` in a browser: send a question (stub answer appears), confirm it shows in the Recent sidebar, click it to prefill, and reload to confirm Recent persists.

- [ ] **Step 6: Commit**

```bash
git add cmd/api/main.go
git commit -m "feat(api): serve chi handler + embedded UI with graceful shutdown

Co-Authored-By: Claude Opus 4.8 (1M context) <noreply@anthropic.com>"
```

---

## Self-Review notes (addressed)

- **Spec coverage:** server+middleware+timeouts (Task 2), `/healthz` (Task 2), `/api/chat` validate/resolve/persist/stub (Tasks 1+3), error envelope + codes (Tasks 2+3), two-pane UI + localStorage recent (Task 4), `cmd/api` run + graceful shutdown (Task 5), DoD checks (Tasks' run steps). All spec §2–§11 items map to a task.
- **Type consistency:** `EnsureConversation(ctx, id string) (string, error)` and `AddMessage(ctx, conversationID, role, content string) error` are identical in `store` (Task 1), the `ChatStore` interface (Task 2), the fake (Task 2), and the handler (Task 3). `chatResponse`/`source`/`errorEnvelope` field names match across handler and tests. `server.New(cfg, chat, assets)` signature matches its call in Task 5.
- **Cross-task compile note:** Task 2's `server.go` references `a.handleChat` (Task 3). Implement Tasks 2 and 3 together, or use the temporary not-implemented stub noted in Task 2 Step 5. Flagged there explicitly.
- **No placeholders:** every code step contains complete, runnable code.
