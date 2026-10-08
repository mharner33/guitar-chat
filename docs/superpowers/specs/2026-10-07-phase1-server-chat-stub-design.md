# Phase 1 — Server + UI + Chat Stub (Design Spec)

Date: 2026-10-07
Status: Approved
Milestone: Phase 1 / Milestone 1 (`docs/plan.md` steps 3–4), extended with a wired-but-stubbed `POST /api/chat`.

## 1. Goal & scope

Deliver the HTTP server shell, the embedded web UI, and a **wired `POST /api/chat`** that
exercises the full request → validate → persist → respond seam while returning a
**placeholder answer**. The real RAG pipeline is Phase 2; this build stands up everything
around it so Phase 2 only has to fill in the answer generation.

### In scope
- `chi` server with middleware (request id, panic recovery, slog-JSON access log, 30s timeout).
- `GET /healthz` → `200 {"status":"ok"}` (checks nothing — not Postgres, not the LLM).
- `POST /api/chat`: validate input, resolve/create the conversation, persist the verbatim
  user message, return the real response shape with a placeholder answer and `sources: []`.
- Embedded web UI (`go:embed`): two-pane layout — a left "recent queries" sidebar
  (client-side `localStorage`) and a right chat window.
- `cmd/api/main.go` wires config → store → server and runs an `http.Server` with graceful
  shutdown.

### Out of scope (later phases, do not build here)
- RAG pipeline and the `Embedder` / `Retriever` / `LLMClient` interfaces + clients (Phase 2).
- Ingestion and the `corpus/` content (Phase 2).
- Datadog Agent Observability and a real `trace_id` (Phase 3).
- `POST /api/feedback` and the 👍/👎 UI (Phase 5).
- Offline/online evals (Phases 4–5).

## 2. API contracts (authoritative: `docs/design.md §6`)

### `POST /api/chat`
Request:
```json
{ "conversation_id": "uuid-or-null", "question": "..." }
```
Response `200`:
```json
{
  "conversation_id": "…",
  "answer": "…placeholder for Phase 1…",
  "sources": [],
  "trace_id": ""
}
```
- `conversation_id` null → mint a new conversation (new UUID). A non-null id missing from the
  table is created with that id (so a `localStorage` id works).
- The user question is stored **verbatim**.
- `trace_id` is `""` in Phase 1 (populated in Phase 3); stored `messages.trace_id` stays NULL.
- `sources` is always `[]` in Phase 1. The element shape for later phases is
  `{ "n": int, "title": string, "section": string, "score": float }`.

### `GET /healthz`
`200 {"status":"ok"}`. No dependency checks.

### Errors
Envelope `{"error": {"code": "...", "message": "..."}}` with the proper status code.
All Phase 1 validation failures are `400`:

| Condition | `code` |
|---|---|
| Body is not valid JSON | `invalid_json` |
| `question` empty/whitespace after trim | `invalid_request` |
| `question` exceeds the length cap | `question_too_long` |
| `conversation_id` present but not a valid UUID | `invalid_conversation_id` |

Length cap: a package const in `internal/server` (`maxQuestionRunes = 4000`). Counted in runes.

## 3. Components / files

New:
- `internal/server/server.go` — `New(cfg config.Config, store ChatStore) http.Handler`
  builds the `chi` router, mounts middleware, and registers routes (`GET /healthz`,
  `POST /api/chat`, UI at `/*`).
- `internal/server/chat.go` — the `/api/chat` handler, request/response structs, validation,
  conversation resolution, persistence, and the stub answer. Also declares the `ChatStore`
  interface (consumer-defined) that the handler needs.
- `internal/server/respond.go` — `writeJSON` and `writeError` helpers + the error envelope type.
- `internal/store/chat.go` — SQL for conversation + message persistence (SQL lives in `store`).
- `web/web.go` — `//go:embed index.html app.js style.css` exposing `embed.FS` (`package web`).
- `web/index.html`, `web/app.js`, `web/style.css` — the two-pane UI.
- `internal/server/server_test.go`, `internal/server/chat_test.go` — table-driven tests with a
  fake `ChatStore` and `httptest` (no DB, no network).

Modified:
- `cmd/api/main.go` — replace the `// Phase 1 wires the chi server and routes here` stub:
  construct the server and run `http.Server{Addr: cfg.Addr, Handler: …}` with
  `ReadHeaderTimeout`/`ReadTimeout`/`WriteTimeout`/`IdleTimeout`, and graceful shutdown via
  `signal.NotifyContext(..., os.Interrupt, syscall.SIGTERM)` + `srv.Shutdown(ctx)`.
- `.env.example` — only if a new config variable is introduced; Phase 1 introduces none (the
  length cap is a const), so no change expected.

## 4. Interfaces & boundaries

The handler depends on a small consumer-defined interface, matching the design's
fake-able-stages convention:
```go
// internal/server/chat.go
type ChatStore interface {
    // EnsureConversation returns the canonical conversation id. An empty id mints a new
    // conversation; a non-empty id is created if absent and returned unchanged.
    EnsureConversation(ctx context.Context, id string) (string, error)
    // AddMessage stores one message (role is "user" or "assistant"). trace_id is left NULL.
    AddMessage(ctx context.Context, conversationID, role, content string) error
}
```
`*store.Store` implements `ChatStore`. Unit tests use a fake that records calls and can
return injected errors.

## 5. Store SQL (`internal/store/chat.go`)

- `EnsureConversation(ctx, id string) (string, error)`:
  - `id == ""`: `INSERT INTO conversations (id) VALUES (gen_random_uuid()) RETURNING id`.
  - `id != ""`: `INSERT INTO conversations (id) VALUES ($1) ON CONFLICT (id) DO NOTHING`,
    then return `id`. (`gen_random_uuid()` is built-in on the Supabase Postgres 17 target.)
- `AddMessage(ctx, conversationID, role, content string) error`:
  `INSERT INTO messages (conversation_id, role, content) VALUES ($1, $2, $3)`
  (`trace_id` omitted → NULL; `role` already CHECK-constrained to `user`/`assistant`).

UUID validation of a caller-supplied `conversation_id` happens in the handler before calling
the store (format check; invalid → `400 invalid_conversation_id`).

## 6. `POST /api/chat` data flow (stub)

1. Decode JSON body → `invalid_json` on failure.
2. Trim `question`; empty → `invalid_request`; `len([]rune) > maxQuestionRunes` →
   `question_too_long`.
3. If `conversation_id` present, validate UUID format → `invalid_conversation_id` on failure.
4. `convID, err := store.EnsureConversation(ctx, idOrEmpty)` (empty string when request id was
   null/absent).
5. `store.AddMessage(ctx, convID, "user", question)` — verbatim.
6. Build the stub answer (see §7) and `store.AddMessage(ctx, convID, "assistant", answer)`.
7. Respond `200 {conversation_id: convID, answer, sources: [], trace_id: ""}`.

Store errors → `500` with code `internal` and the error logged (never the SQL/secret).

## 7. Stub answer

A fixed, explicit placeholder that calls no LLM and asserts no theory, honoring the
"never bypass RAG / no ungrounded answers" guardrail:

> "The answer pipeline isn't wired up yet — grounded answers arrive in Phase 2."

Defined as a package const. The placeholder assistant message **is persisted** so the full
persistence path is exercised and visible in the DB.

## 8. Middleware & timeouts

- `middleware.RequestID`, `middleware.Recoverer` (defensive — conventions forbid panics
  outside `main`, but a recovered panic must become a `500`, not crash the server).
- A small slog-JSON access-log middleware: method, path, status, duration, request id.
- `middleware.Timeout(30 * time.Second)` — the answer-path deadline seam.
- `http.Server`: `ReadHeaderTimeout: 10s`, `ReadTimeout: 15s`, `WriteTimeout: 35s`
  (> the 30s handler timeout so `TimeoutHandler`/middleware fires first), `IdleTimeout: 60s`.

## 9. Frontend (two-pane, vanilla)

- **Left pane — Recent queries.** `localStorage` key `gc_recent`: array of strings,
  newest-first, deduped (re-sending moves an entry to the top), capped at 20. Each entry is a
  button that **prefills and focuses the textarea** (no auto-send). A "Clear" control empties
  the list. Re-renders on each send.
- **Right pane — Chat window.** Message thread (user + assistant bubbles), textarea + send
  button. `conversation_id` persisted under `localStorage` key `gc_conversation`. `fetch` POSTs
  to `/api/chat`; while in flight the input + button are disabled (loading state); API errors
  render inline in the thread. Source-chip rendering code is present but renders nothing for the
  empty `sources` of the stub.
- **Layout.** Flexbox: sidebar ~240px + chat flex-fill; a `max-width: 640px` media query stacks
  the sidebar above the chat so mobile isn't broken. No framework, no npm, no build step.
- **Serving.** `go:embed` via `web/web.go`; the server mounts it at `/*` (e.g.
  `http.FileServerFS(web.Assets)`), with `/` serving `index.html`.

The recent-queries feature is entirely browser-side: no server, store, or test changes.

## 10. Testing

Unit (`internal/server`, table-driven, fake `ChatStore`, `httptest` — no DB/network):
- `GET /healthz` → `200`, body `{"status":"ok"}`.
- `POST /api/chat` happy path, null `conversation_id` → `200`, response shape correct,
  `conversation_id` non-empty, `sources == []`, `trace_id == ""`; fake recorded one
  `EnsureConversation("")` and two `AddMessage` calls (`user` verbatim, then `assistant`).
- `POST /api/chat` with a valid provided `conversation_id` → passed through to
  `EnsureConversation` and echoed back.
- Validation matrix → `400` with the right `code`: malformed JSON, empty/whitespace question,
  over-length question, non-UUID `conversation_id`.
- Store error from the fake → `500` code `internal`.
- `GET /` → `200` and body contains an index.html marker string.

## 11. Definition of Done

- `go vet ./...`, `staticcheck ./...`, `go test ./...`, `gofmt -l .` all clean.
- `go run ./cmd/api` serves: `/healthz` green, UI loads, a chat message round-trips the stub and
  persists two `messages` rows (+ a `conversations` row) in the configured database.
- No new `.env` variables (none introduced); if that changes, update `.env.example` and
  `docs/design.md`.
