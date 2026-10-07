# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Read these first

This repo has two authoritative documents. Read them before making any change — they are not optional background:

- **`agent.md`** — the owner's canonical instruction file for AI coding agents (same role as this file; the owner's convention names it `agent.md`). It defines coding conventions, the observability/RAG/eval rules, and the Definition of Done. **It takes precedence over this file where they overlap.** If a community-standard `AGENTS.md` is expected, symlink it: `ln -s agent.md AGENTS.md` (keep `agent.md` as the source of truth).
- **`docs/design.md`** — the source of truth for scope and architecture. Consult §1 (Non-Goals) before expanding scope, and §4–9 before architectural decisions.

## Project state

This is a **greenfield project**: the design is complete but no Go source has been written yet. `go.mod` and `go.sum` exist with dependencies resolved (notably the experimental `github.com/DataDog/dd-trace-go/v2` llmobs SDK, pinned at `v2.10.1`); the `cmd/`, `internal/`, `web/`, `migrations/`, `corpus/`, and `evals/` trees described below do not exist yet and are to be built per `docs/design.md §4.1`. When implementing, follow the layout and milestones in the design doc rather than inventing new structure.

Note: `go.mod` declares `go 1.26.2`. (`agent.md` text says "Go 1.23+" as a floor.)

## What this is

Guitar Theory Chat: a demo chat app that answers guitar music-theory questions using RAG over a curated markdown corpus, quality-checked by offline + online LLM-as-judge evals, and instrumented end-to-end with **Datadog Agent Observability** via the experimental Go llmobs SDK. The instrumentation and the evals *are the point of the demo* — they are never skipped to "ship faster."

## Commands

```bash
# run the whole demo (db + app + one-shot ingest)
docker compose up --build
docker compose up db            # just Postgres for local dev; migrations run at app start

# local dev (need DATABASE_URL and LLM_* env — see .env.example)
go run ./cmd/api                # HTTP server + embedded web/ UI
go run ./cmd/ingest             # chunk + embed corpus into pgvector (idempotent)
go run ./cmd/eval               # offline eval harness over evals/golden.jsonl; exit code IS the gate

# checks — all must pass before every commit; CI runs the same set
go vet ./...
staticcheck ./...
go test ./...
gofmt -l .                      # must print nothing

# run a single test
go test ./internal/rag -run TestRetriever
```

## Architecture big picture

Single Go module, three binaries under `cmd/` (`api`, `ingest`, `eval`), with all logic in `internal/` packages: `config`, `server` (chi routes), `rag` (pipeline, chunker, retriever, embedder, prompts), `llm` (OpenAI-compatible client), `store` (pgx/v5 over Postgres 16 + pgvector), `observe` (llmobs wrappers + Datadog HTTP), `evaljudge`. Frontend is vanilla HTML/CSS/JS in `web/`, served by the Go binary via `go:embed` — no framework, no npm, no build step.

**The retriever, embedder, and LLM client are small single-method interfaces** (`Retriever`, `Embedder`, `LLMClient`) defined at their consumers so every stage is fake-able in unit tests (no network in unit tests). Keep them that way.

### The one trace per chat request (the core invariant)

`POST /api/chat` produces **exactly one** Agent Observability trace, rooted at a `chat.answer` workflow span (session = conversation id):

```
workflow "chat.answer"
├─ llm "rewrite.query"        (follow-ups only — when history is non-empty)
├─ embedding "embed.question" (search query → vector)
├─ retrieval "retrieve.chunks"(pgvector top-K cosine, records ALL scores incl. rejected)
└─ llm "generate.answer"      (ONLY when a chunk passes RETRIEVAL_MIN_SCORE)
```

- **Empty retrieval** (no chunk ≥ `RETRIEVAL_MIN_SCORE`, default 0.5) returns the exact string `"I don't have enough context"` with `sources: []` and **no `generate.answer` span** — the answer model is never called, and the online judge does not run. Never let the LLM free-run on guitar theory without grounded context.
- **LLM spans are the only billable span kind.** Never emit an `llm` span on a path that did not call the provider.
- **Follow-up rewriting**: empty history → embed the raw question, no rewriter. Non-empty history → one `rewrite.query` completion (temp 0) rewrites recent messages into a standalone search query; on error/timeout/blank output, fall back to embedding the raw question — a rewrite failure must never fail the chat request. Never embed prior assistant answers.
- **Online judge** (`judge.answer`) runs async, sampled, *after* the response returns, in its own workflow span linked back via `AddLink` using the captured root `SpanID()`/`TraceID()`. Never hold the root span open for it.

### Observability is isolated on purpose

The llmobs SDK is **experimental** (may change/break without notice, outside dd-trace-go's compatibility promise). Therefore: pin the exact `dd-trace-go` version in `go.mod` (never loose ranges), and keep **every** llmobs call behind thin helpers in `internal/observe` — no direct SDK imports anywhere else, so an API break touches one package. The Go path is undocumented on Datadog's site; verify setup/start options against the package source (`llmobs/options.go`), the README at `github.com/DataDog/dd-trace-go/tree/main/llmobs`, or pkg.go.dev. Export is **agentless to US1** (`DD_SITE=datadoghq.com`, `DD_LLMOBS_AGENTLESS_ENABLED=true`, `DD_LLMOBS_ML_APP=guitar-chat`). Do **not** add an OpenTelemetry exporter — v1 uses the llmobs SDK only. If telemetry is misconfigured, the app must degrade to logs, never crash.

### Ingestion is idempotent and reconciling

`cmd/ingest` chunk IDs are content hashes (`sha256(doc_id + section + text)`), chunking is heading-aware (`##` splits, ~800 token max, ~100 token overlap). For each corpus file: upsert the doc + current chunks, then delete that doc's chunk rows whose hashes weren't in the file just read. After the walk, delete documents whose slugs are gone from `corpus/`. Re-running must never duplicate rows, and a degraded/removed file must not leave stale chunks searchable (this is what makes the "groundedness drops after a degraded re-ingest" demo step true). Both `api` and `ingest` refuse to start if `EMBEDDING_DIMENSIONS` ≠ the `vector(N)` column width.

## Conventions that aren't obvious

- All config is read **once** in `internal/config` into a struct — no scattered env lookups. Update `.env.example` whenever you add a variable.
- Thread `ctx` everywhere; every LLM/DB/HTTP call gets a deadline (30s answer path, 60s judge path). Errors wrap with `%w` and return; no panics outside `main`.
- Migrations are forward-only numbered SQL files (`001_…`) in `migrations/`; **never edit an applied migration**. A vector-width change is a new migration.
- Tests are table-driven with fakes for `LLMClient`/`Embedder`/`Retriever`.
- Run `go run ./cmd/eval` after **any** change to prompts, retrieval params, models, or corpus content, and add a golden row for every answer-quality bug you fix. The gate (mean correctness ≥ 0.8, no row with groundedness 0) is enforced by the process exit code in CI.

## Guardrails

- **Demo scope** (per `docs/design.md §1`): no auth, no Kubernetes, no multi-agent orchestration, no streaming in v1. If a request pulls toward these, stop and propose a design change rather than silently expanding scope.
- Never bypass RAG (no ungrounded LLM answers) and never skip instrumentation — both are the demo's core purpose.
- Licensing: root `LICENSE` is MIT and covers the Go code and `web/`. Everything under `corpus/` is CC BY-SA 4.0 (including hand-written pages). Keep `corpus/LICENSE` and `corpus/ATTRIBUTION.md` accurate when adding or adapting Open Music Theory material.
- When the llmobs SDK surprises you, fix the `internal/observe` wrappers, add a regression test, and note the workaround in `agent.md`'s Observability Rules section.
