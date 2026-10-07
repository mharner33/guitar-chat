agent.md — Instructions for AI Coding Agents
Project
Guitar Theory Chat: a demo chat application that answers guitar music theory questions using RAG over a curated markdown corpus, with answer quality measured by offline + online LLM-as-judge evaluations, instrumented end-to-end with Datadog Agent Observability.
Backend: Go (single module, cmd/api serves, cmd/ingest builds the index, cmd/eval runs the eval harness)
Frontend: vanilla HTML/CSS/JS in web/ (no framework, no build step), served by the Go binary via go:embed
Database: Postgres 16 + pgvector (documents, chunks, conversations, feedback, eval runs)
LLM: OpenAI-compatible REST API. The demo provider is OpenRouter (a small chat model). Chat and embeddings may use different base URLs. Vector width is EMBEDDING_DIMENSIONS and must match the embedding model.
Observability: Datadog Agent Observability on US1 (DD_SITE=datadoghq.com), agentless, via the experimental github.com/DataDog/dd-trace-go/v2/llmobs Go SDK. That SDK is the only instrumentation path; do not add an OpenTelemetry exporter.
Design doc: docs/design.md — read it before making architectural decisions. It is the source of truth for scope.
Commands
# run everything (db + app + one-shot ingest)
docker compose up --build

# local dev
go run ./cmd/api            # needs DATABASE_URL, LLM_* env (see .env.example)
go run ./cmd/ingest          # chunk + embed corpus into pgvector (idempotent)
go run ./cmd/eval            # offline eval harness over evals/golden.jsonl

# checks — run before every commit; CI runs the same
go vet ./...
staticcheck ./...
go test ./...
gofmt -l .                    # must output nothing

Postgres for local dev: docker compose up db. Migrations in migrations/ run automatically at app start.
Architecture Summary
One trace per chat request:
POST /api/chat
└─ workflow span "chat.answer"          (root; session = conversation id)
   ├─ llm span "rewrite.query"          (follow-ups only; standalone search query)
   ├─ embedding span "embed.question"   (search query → vector)
   ├─ retrieval span "retrieve.chunks"  (pgvector top-K cosine search, then score floor)
   └─ llm span "generate.answer"        (only when a chunk passes the floor;
                                          RAG prompt + history → answer, token metrics)

Empty retrieval (no chunk at or above RETRIEVAL_MIN_SCORE) returns "I don't have enough context" with sources []. That trace has embedding and retrieval spans, records the rejected scores, and has no generate.answer span. rewrite.query is still present when this was a follow-up. The online judge does not run.

after a generated answer returns (async, sampled; not after the canned empty answer):
└─ workflow span "judge.answer"         (own trace, AddLink back to the answer trace;
                                          evals joined via captured SpanID/TraceID)

Key packages (all under internal/): config, server (chi routes), rag (pipeline, chunker, retriever, embedder, prompt templates), llm (provider client), store (pgx), observe (llmobs wrappers + Datadog HTTP), evaljudge (golden dataset + judge).
Interfaces are small and single-method (Retriever, Embedder, LLMClient) so each stage is fake-able in tests. Keep them that way.
Coding Conventions
Go 1.23+. Standard library first; every dependency must be justified.
chi for routing, pgx/v5 for Postgres, slog JSON logging. No ORMs, no codegen.
Errors: wrap with %w, return them; no panics outside main. Handle and log at the boundary.
Context: thread ctx through everything; every LLM/DB/HTTP call gets a deadline (30s answer path, 60s judge path).
All config via environment variables read once in internal/config into a struct — no config lookups scattered in code. Update .env.example when adding a variable.
Frontend: no frameworks, no npm, no build. web/app.js uses fetch only.
Migrations: forward-only SQL files in migrations/, numbered (001_…), never edit an applied migration.
Tests: table-driven, fakes for LLMClient/Embedder/Retriever. No network in unit tests.
Observability Rules (important)
The instrumentation IS the point of this demo. Do not skip or defer it.
The llmobs package is experimental (github.com/DataDog/dd-trace-go/v2/llmobs): it may change or be removed without notice and is outside dd-trace-go's compatibility promise. Therefore:
Pin the exact dd-trace-go version in go.mod; never use loose version ranges.
Every llmobs call lives in internal/observe behind thin helpers. No direct SDK imports elsewhere. If the SDK API breaks, only internal/observe changes.
The package is not documented on Datadog's docs site; when in doubt about setup/start options, check the package README at https://github.com/DataDog/dd-trace-go/tree/main/llmobs, the pkg.go.dev reference, or the module source itself (llmobs/options.go has the authoritative StartOption list). Expected env vars mirror the other SDKs (DD_LLMOBS_ML_APP, DD_LLMOBS_AGENTLESS_ENABLED, DD_API_KEY, DD_SITE) but verify against the source.
Site is US1: DD_SITE=datadoghq.com, intake host https://api.datadoghq.com. Export is agentless (DD_LLMOBS_AGENTLESS_ENABLED=true). DD_LLMOBS_ML_APP=guitar-chat. v1 does not implement the OpenTelemetry fallback in docs/design.md §2; feedback still uses the eval-metric HTTP API.
Every chat request must produce exactly one trace. It always contains embed.question and retrieve.chunks. Follow-ups also contain rewrite.query. generate.answer is present only when the answer model was called. Use the typed constructors (StartWorkflowSpan, StartEmbeddingSpan, StartRetrievalSpan, StartLLMSpan) and annotate with the per-kind IO methods (AnnotateTextIO, AnnotateEmbeddingIO, AnnotateRetrievalIO, AnnotateLLMIO) plus WithAnnotatedMetrics for token counts (MetricKeyInputTokens / MetricKeyOutputTokens / MetricKeyTotalTokens). Embedding and retrieval annotations record the search query that was actually embedded (the rewrite output on follow-ups, the raw question otherwise). Do not emit an llm span for a path that did not call the provider: LLM spans are the billable kind.
On failure, finish spans with Finish(WithError(err)) — errors must be visible in the trace, not just logs.
Judge scores are submitted with llmobs.SubmitEvaluationFromSpan(...) joined to the answer trace. The online judge runs asynchronously after a generated answer is returned: capture the root span's SpanID()/TraceID() in the handler, run the judge in a goroutine with its own small workflow span linked via AddLink, and submit the evaluation using a struct carrying the captured IDs (anything implementing SpanID() + TraceID() works). Never hold the root span open waiting for the judge. Skip the judge when generate.answer did not run. Online scores are not written to eval_runs. Thumbs up/down goes to the local feedback table and to Datadog via the eval-metric HTTP API with event_kind=feedback (the Go SDK has no feedback API yet) — see internal/observe/datadog.go.
Never log secrets or send API keys to providers; span annotations contain question/answer text only (this is a demo, keep payloads clean).
If Agent Observability is misconfigured (missing DD_API_KEY etc.), the app must still work — degrade to logs, never crash on telemetry failures.
RAG Rules
The system prompt constrains answers to the provided context with numbered [1], [2] citations; answers must include source chips. If the supplied chunks still do not support an answer, the model says "I don't have enough context" and adds no theory that is not in the context. Do not weaken this constraint without running the eval harness and updating evals/golden.jsonl expectations.
The answer model sees the original question, the last HISTORY_MESSAGE_LIMIT messages (default 10), and the chunks that passed the score floor. The stored user message is the verbatim question.
Follow-up retrieval: when history is empty, embed the question as written and do not call the rewriter. When history is present, one completion named rewrite.query (temperature 0, REWRITE_MAX_TOKENS default 64) rewrites the last REWRITE_HISTORY_MESSAGES messages (default 4, two turns) into a standalone search query. Embed that query. On rewrite error, timeout, or blank output, embed the raw question and continue; the rewrite failure must not fail the chat request. Do not embed prior assistant answers.
Ingestion is idempotent (chunk IDs are content hashes) and reconciles to the current corpus. For each corpus file, upsert the document and its current chunks, then delete chunk rows for that document whose hashes were not in the file just read. After the walk, delete documents whose slugs are no longer in corpus/, and their chunks. Re-running cmd/ingest must never duplicate rows, and a degraded or removed file must not leave old chunks searchable.
Chunking is heading-aware (## splits, ~800 token max, ~100 token overlap). Chunk metadata (doc_title, section) flows to retrieval spans and to UI source chips.
Retrieval uses cosine similarity. Drop chunks below RETRIEVAL_MIN_SCORE (default 0.5). If none remain, or retrieval errors, return the exact answer "I don't have enough context" with sources [] and do not call the answer model. Never let the LLM free-run on guitar theory without grounding.
Evaluation Workflow
evals/golden.jsonl: one JSON object per line — {"id", "question", "expected_points": [...], "rubric"}. Cover modes, chord construction, intervals on the fretboard, chord-scale relationships, enharmonics.
cmd/eval runs the full RAG pipeline per row (not just the LLM), then an LLM-as-judge pass scoring 0–2 on correctness, groundedness, and guitar-friendliness. Judge output is strict JSON; a parse failure counts as a failed row (correctness 0, and the row does not count as passed). Golden rows are single-turn, so rewrite.query does not run.
At the end of the run, insert one eval_runs row: kind offline, dataset golden, score = mean correctness (parse failures count as 0), passed = rows with a parsed judge result and groundedness > 0, total = golden rows attempted. Online judgments do not insert eval_runs rows.
Gate: mean correctness ≥ 0.8 and no row with groundedness 0, or the run fails. The process exit code is the gate.
Run go run ./cmd/eval after any change to prompts, retrieval parameters, models, or corpus content. Add a golden row for every bug you fix in answer quality.
Online judge (judge.answer span) runs sampled (ONLINE_EVAL_SAMPLE_RATE) only after generate.answer, and must never block or fail the user response. Its groundedness and helpfulness scores go to Datadog on the answer trace only.
Definition of Done
A feature is done when all of these hold:
go vet ./..., staticcheck ./..., go test ./..., gofmt -l . all pass.
The full demo flow works: docker compose up → ingest → ask questions → answers with sources → 👍/👎 works.
Every request produces one US1 Agent Observability trace: workflow, then rewrite.query when the turn is a follow-up, then embedding, then retrieval. generate.answer is present, with model, token, and latency data, when the answer model ran. Empty retrieval omits generate.answer and submits no evaluation.
Judge scores appear as evaluations on traces that called the answer model; 👍/👎 appears as feedback on the corresponding trace.
Prompt/model/retrieval changes pass the offline eval gate.
New config variables are added to .env.example and documented in docs/design.md.
Guardrails
Demo scope: no auth, no Kubernetes, no multi-agent orchestration, no streaming (v1). If a request pulls toward these, stop and check docs/design.md §1 Non-Goals; propose a design change instead of silently expanding scope.
Never bypass RAG (no direct LLM answers without retrieved context) and never skip instrumentation to "ship faster" — both are the core purpose of the demo.
Never commit secrets or .env. Never edit applied migrations.
Code and web/ stay MIT (root LICENSE). Files under corpus/ are CC BY-SA 4.0, including hand-written pages. Keep corpus/LICENSE and corpus/ATTRIBUTION.md accurate when adding or adapting Open Music Theory material.
When the llmobs SDK surprises you (it's experimental — APIs may not match docs), fix internal/observe wrappers, add a regression test, and note the workaround in this file's Observability Rules section.
File Naming Note
This file is agent.md per the project owner's convention. If your coding agent looks for the community-standard AGENTS.md, symlink it: ln -s agent.md AGENTS.md (keep agent.md as the source of truth).
