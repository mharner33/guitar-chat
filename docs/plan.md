# Guitar Theory Chat — Implementation Plan

Step-by-step build plan, grounded in `docs/design.md` (§4.1 layout, §12 milestones) and the rules in `agent.md`. Ordered so each step builds on a working previous one, and each ends in something runnable and observable.

## Phase 0 — Scaffolding (prereq for everything)

1. **Project skeleton & config.** Create the `cmd/`, `internal/`, `web/`, `migrations/`, `corpus/`, `evals/` trees. Build `internal/config/config.go` first — the env struct read *once* (DATABASE_URL, all `LLM_*`, `EMBEDDING_DIMENSIONS`, retrieval/rewrite/history knobs, `DD_*`, `JUDGE_MODEL`, `ONLINE_EVAL_SAMPLE_RATE`). Write `.env.example` alongside it with every variable.
   - *Done when:* `go run ./cmd/api` loads config and fails loudly on missing required vars.

2. **Dev infra.** `docker-compose.yml` (pgvector/pgvector:pg16 + app + one-shot ingest), `Dockerfile` (distroless, `go:embed` web/), and `migrations/001_init.sql` (vector extension, all tables, hnsw cosine index). Wire migrations to run at app start; add the `EMBEDDING_DIMENSIONS` vs `vector(N)` guard in `config`/startup.
   - *Done when:* `docker compose up db` comes up healthy with the schema applied.

## Phase 1 — Milestone 1: server shell + UI + health

3. **HTTP server.** `internal/server/server.go` (chi router, slog JSON middleware, timeouts), `GET /healthz` → `{"status":"ok"}`. `cmd/api/main.go` wires config → store → server.

4. **Frontend shell.** `web/index.html` + `app.js` + `style.css` — chat window, input, send button, `go:embed`-served at `/`. No real backend call yet.
   - *Done (M1):* page loads, health green, Postgres + migrations up.

## Phase 2 — Milestone 2: RAG end-to-end (no observability yet)

5. **Interfaces + fakes.** Define the single-method `Retriever`, `Embedder`, `LLMClient` at their consumers (`internal/rag/…`, `internal/llm/client.go`). Write fakes now so later tests never hit the network.

6. **LLM + embedder clients.** OpenAI-compatible `internal/llm/client.go` (chat) and `internal/rag/embed.go` (HTTP embeddings, respecting `LLM_EMBED_BASE_URL`).

7. **Ingestion.** `internal/rag/chunk.go` (heading-aware `##`, ~800-token max, ~100 overlap) + `cmd/ingest/main.go` + `internal/store` upserts. Implement the **idempotent reconcile**: content-hash chunk IDs, delete-missing-chunks-per-file, delete-removed-documents-after-walk. Seed a few `corpus/*.md` pages + `corpus/LICENSE` + `corpus/ATTRIBUTION.md`.
   - *Done when:* `go run ./cmd/ingest` twice produces zero duplicate rows.

8. **Retriever.** `internal/rag/retriever.go` — pgvector top-K cosine search.

9. **Pipeline + chat handler.** `internal/rag/pipeline.go` (`Answer(ctx, question)`), `internal/rag/prompt.go` (system prompt + assembly), `internal/server/chat.go`. Implement the full answer path: validate → resolve conversation → load history → rewrite-query (follow-ups only, with raw-question fallback) → embed → retrieve → **score floor → canned `"I don't have enough context"` when empty** → generate → persist verbatim user message + answer → return `{answer, sources[], conversation_id, trace_id}`.
   - *Done (M2):* grounded answers with source chips in the UI; off-corpus questions return the canned answer.

## Phase 3 — Milestone 3: Datadog Agent Observability

10. **observe wrappers.** `internal/observe/llmobs.go` — thin typed-span helpers (`StartWorkflowSpan`/`Embedding`/`Retrieval`/`LLM`, per-kind annotate, token metrics, `Finish(WithError)`). **Verify the actual SDK API against `llmobs/options.go`** before coding; no SDK imports outside this package. Agentless US1 setup; degrade to logs if misconfigured.

11. **Instrument the pipeline.** Wrap the Phase 2 pipeline so each request emits exactly one trace: `chat.answer` → (`rewrite.query` on follow-ups) → `embed.question` → `retrieve.chunks` (record *all* scores) → `generate.answer` (only when a chunk passed; with token metrics). No `generate.answer` on empty retrieval.
    - *Done (M3):* full span tree visible in US1 Agent Observability per request.

## Phase 4 — Milestone 4: offline eval harness + CI gate

12. **Eval harness.** `internal/evaljudge/` (golden loader, LLM-as-judge, strict-JSON scoring) + `cmd/eval/main.go`. Run the *full* RAG pipeline per golden row, score correctness/groundedness/guitar-friendliness 0–2, insert one `eval_runs` row, enforce the gate (mean correctness ≥ 0.8, no groundedness-0 row) via exit code. Write `evals/golden.jsonl` (30–50 rows) + `evals/judge_prompt.md`.

13. **CI.** GitHub Actions running `go vet`, `staticcheck`, `go test ./...`, `gofmt -l .`, and the eval gate.
    - *Done (M4):* `go run ./cmd/eval` prints scores; CI red on gate failure.

## Phase 5 — Milestone 5: online judge + feedback

14. **Online judge.** Async, sampled (`ONLINE_EVAL_SAMPLE_RATE`), *after* the response returns — capture root `SpanID()`/`TraceID()`, run in a goroutine with its own `judge.answer` span `AddLink`ed back, submit via `SubmitEvaluationFromSpan`. Never block the response; skip on canned answers.

15. **Feedback.** `POST /api/feedback` → local `feedback` table + Datadog eval-metric HTTP API (`event_kind=feedback`) in `internal/observe/datadog.go`; 👍/👎 buttons in `web/app.js` using the stored `trace_id`.
    - *Done (M5):* judge scores and 👍/👎 appear on the right traces.

## Phase 6 — Milestone 6: dashboards + demo

16. **Dashboards/monitors** (request rate, p95 latency, token spend, judge-score distribution, thumbs ratio, empty-retrieval rate) and the monitors from §9.4.

17. **Demo storyline rehearsal** (§9.5), including the degraded-re-ingest → groundedness-drop step that proves the evals catch regressions.
    - *Done (M6):* full demo flow per the Definition of Done in `agent.md`.

## Sequencing notes

- Observability (Phase 3) comes *after* a working RAG path rather than interleaved — it's easier to instrument a pipeline you can already see working end-to-end. If the experimental llmobs SDK is the riskiest unknown, consider a thin spike of step 10 earlier to de-risk the API.
- Steps 5–9 are the critical path and the bulk of the code.
- Every phase respects the `agent.md` checks (`go vet`, `staticcheck`, `go test ./...`, `gofmt -l .`) and the demo-scope guardrails (no auth, k8s, multi-agent, or streaming in v1).
