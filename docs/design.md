Guitar Theory Chat — Design Document
Status: Draft v1.1 (October 2026) Owner: Mike Harner Purpose: A demo application that answers guitar music-theory questions in a chat UI, instrumented end-to-end with Datadog Agent Observability, and continuously quality-checked with RAG-grounded evaluations.

1. Goals and Non-Goals
Goals
Chat experience: a single-page chat window that answers guitar music theory questions ("Why does a major scale sound different from a minor scale?", "What chords work over D Dorian?").
RAG pipeline: every answer is grounded in a curated knowledge corpus via retrieval-augmented generation, with sources surfaced in the answer.
Evaluations: answer quality is measured two ways — an offline golden-dataset eval harness (LLM-as-judge) and online evals + user feedback submitted to Datadog.
Datadog Agent Observability: every chat request produces one trace on US1 (workflow → optional query rewrite → embedding → retrieval → LLM when retrieval hits) visible in Datadog, with token counts, latency, model, and eval scores. Instrumentation uses the experimental Go llmobs SDK only.
Go backend, plain JavaScript frontend: no frontend framework, no database beyond Postgres.
Non-Goals (demo scope)
User authentication, multi-tenancy, rate limiting
Kubernetes, horizontal scale, HA
Multi-agent orchestration, tool calling beyond retrieval
Streaming token-by-token responses (v1 returns complete answers; streaming is future work)
Editing the corpus from the UI

2. Architecture Overview
┌─────────────────────────────────────────────────────────────────┐
│ Browser                                                          │
│  index.html + app.js (vanilla JS, fetch-based chat, 👍/👎)        │
└──────────────┬──────────────────────────────────────────────────┘
               │ POST /api/chat   POST /api/feedback   GET /healthz
┌──────────────▼──────────────────────────────────────────────────┐
│ Go backend (single binary: cmd/api)                              │
│                                                                  │
│  HTTP layer (chi)                                                │
│    └─ ChatHandler                                                │
│         ├─ RAG pipeline ("workflow" root span)                   │
│         │    ├─ rewrite.query (follow-ups only) → llm span       │
│         │    ├─ Embedder.Embed()        → embedding span         │
│         │    ├─ Retriever.Search()      → retrieval span         │
│         │    ├─ score floor; empty → canned answer, no llm span  │
│         │    ├─ Prompt assembly         → attributes on root     │
│         │    └─ LLMClient.Chat()        → llm span (model+tokens)│
│         ├─ OnlineJudge.Evaluate()      → judge workflow (linked; │
│         │                                  skipped when no LLM)  │
│         └─ Evals via SDK + feedback via HTTP API                │
│                                                                  │
│  Offline: cmd/ingest (chunk + embed corpus)                       │
│           cmd/eval  (golden dataset + LLM-as-judge harness)       │
│                                                                  │
│  dd-trace-go llmobs SDK (experimental Go support)                │
└──────────────┬───────────────────────────┬──────────────────────┘
               │ SQL + pgvector             │ HTTPS
┌──────────────▼──────────────┐  ┌─────────▼─────────────────────┐
│ Postgres 16 + pgvector      │  │ OpenRouter chat (small model) │
│  - documents, chunks        │  │ embeddings endpoint may differ│
│  - conversations, feedback  │  └───────────────────────────────┘
│  - eval runs                │
└─────────────────────────────┘  ┌────────────────────────────────┐
                                 │ Datadog US1 (datadoghq.com)     │
                                 │  Agent Observability traces,    │
                                 │  eval metrics, feedback, logs   │
                                 └────────────────────────────────┘

Key decision — how a Go app talks to Datadog Agent Observability. The officially documented SDKs cover Python (ddtrace), Node.js (dd-trace), and Java (dd-trace-java) (Datadog SDK reference). However, dd-trace-go now ships an experimental Go SDK for Agent Observability in dd-trace-go/llmobs (import path github.com/DataDog/dd-trace-go/v2/llmobs, API reference). It exposes typed spans for every Agent Observability span kind, per-kind input/output annotation, token metrics, and evaluation submission — exactly what this demo needs to showcase.
The design uses the experimental Go llmobs package as the only instrumentation path. v1 does not implement the OpenTelemetry or raw span-intake fallbacks below. The demo is also a test of this SDK, so every chat request is instrumented with it:
StartWorkflowSpan, StartAgentSpan, StartToolSpan, StartTaskSpan, StartEmbeddingSpan, StartRetrievalSpan, StartLLMSpan create typed spans; each returns (span, ctx) so child spans nest by passing the returned context down the call chain.
Per-kind annotation: AnnotateLLMIO (messages), AnnotateEmbeddingIO (documents in, embedding text out), AnnotateRetrievalIO (query in, scored documents out), AnnotateTextIO (workflow/agent/tool/task), plus WithAnnotatedMetadata, WithAnnotatedMetrics (token counts via MetricKeyInputTokens/MetricKeyOutputTokens/MetricKeyTotalTokens), and WithAnnotatedTags.
Span options: WithMLApp, WithModelName, WithModelProvider, WithSessionID.
Evaluation submission: SubmitEvaluationFromSpan (joins by span/trace ID) and SubmitEvaluationFromTag (joins by tag key-value pair), with boolean, categorical (string), or numeric score values.
Finish(WithError(err)) marks failures with stack traces.
Caveats to manage: the package is explicitly experimental — it may change or be removed without notice and is not covered by dd-trace-go's compatibility promise (API reference). Mitigations: pin the dd-trace-go version in go.mod; wrap every llmobs call behind internal/observe so a future API change touches one package; and treat the package README in the repo as the source of truth since official Datadog documentation for the Go path does not exist yet.
Documented fallbacks, not built in v1, if the experiment bites: (1) emit OpenTelemetry spans following the OTel v1.37+ GenAI semantic conventions and export OTLP with the dd-otlp-source=llmobs header — Datadog maps gen_ai.* attributes into the Agent Observability schema (OTel instrumentation guide); (2) submit spans and evaluations directly via the HTTP API — POST /api/intake/llm-obs/v1/trace/spans and POST /api/intake/llm-obs/v2/eval-metric (Agent Observability API). Feedback already uses the eval-metric HTTP API, because the Go SDK has no feedback API. Do not add an OTel exporter unless this design is revised.

3. Technology Stack
Component
Choice
Rationale
Backend language
Go 1.23+
Required. Single static binary, strong concurrency for fan-out evals.
HTTP router
go-chi/chi/v5
Idiomatic, minimal middleware for logging/tracing.
Postgres driver
pgx/v5 + pgxpool
Standard, supports pgvector via pgvector/pgvector-go.
Vector store
Postgres 16 + pgvector
One database for everything; user already runs Postgres. Retriever is an interface, so swapping in a dedicated vector DB later is a config change (pgvector).
LLM + embeddings
OpenAI-compatible REST (configurable base URL)
OpenAI-compatible so the provider stays pluggable. The demo chat model is a small model on OpenRouter (https://openrouter.ai/api/v1). Embeddings may use the same base URL or LLM_EMBED_BASE_URL. Vector width is EMBEDDING_DIMENSIONS.
LLM Observability
github.com/DataDog/dd-trace-go/v2/llmobs (experimental)
Native typed spans for every Agent Observability span kind; wrapped behind internal/observe because the package is experimental and not under compatibility promise (repo).
Frontend
Vanilla HTML/CSS/JS
Required. One page, one JS file, fetch API.
Packaging
Docker Compose (app + postgres)
docker compose up runs the demo.
CI checks
go vet, staticcheck, go test ./...
Fast, standard.


Corpus source: seed content from openly licensed music-theory references (for example Open Music Theory) plus hand-written guitar-specific pages (CAGED system, chord construction, modes on the fretboard). Ingestion is a separate command so the corpus can evolve without redeploying.

Licenses: root LICENSE is MIT and covers the Go code and web/. Everything under corpus/ is CC BY-SA 4.0, including hand-written pages, because those files share a directory with adapted Open Music Theory text. corpus/LICENSE is the BY-SA 4.0 text. corpus/ATTRIBUTION.md credits Open Music Theory with title, authors, source URL, the BY-SA 4.0 link, and a statement that the pages were excerpted and adapted. Keep that file current when corpus text changes.

4. Backend Design
4.1 Repository layout
guitar-chat/
├── agent.md                    # instructions for AI coding agents
├── LICENSE                     # MIT: code and web/
├── docker-compose.yml
├── Dockerfile
├── .env.example
├── docs/design.md              # this document
├── corpus/                     # markdown knowledge base (one topic per file)
│   ├── LICENSE                 # CC BY-SA 4.0
│   ├── ATTRIBUTION.md          # Open Music Theory credit
│   ├── modes.md
│   ├── chord-construction.md
│   └── ...
├── cmd/
│   ├── api/main.go             # HTTP server
│   ├── ingest/main.go           # corpus → chunks → embeddings → pgvector
│   └── eval/main.go            # offline eval harness
├── internal/
│   ├── config/config.go        # env-driven config struct
│   ├── server/server.go        # chi routes, middleware
│   ├── server/chat.go          # POST /api/chat
│   ├── server/feedback.go      # POST /api/feedback
│   ├── rag/                    # pipeline orchestration
│   │   ├── pipeline.go         # Answer(ctx, question) → RAG orchestration
│   │   ├── chunk.go            # heading-aware markdown chunking
│   │   ├── retriever.go        # Retriever interface + pgvector impl
│   │   ├── embed.go            # Embedder interface + HTTP impl
│   │   └── prompt.go           # prompt templates + assembly
│   ├── llm/client.go           # LLMClient interface + OpenAI-compatible impl
│   ├── store/store.go          # pgx queries (docs, chunks, convos, feedback)
│   ├── observe/llmobs.go      # dd-trace-go llmobs setup + span helpers
│   ├── observe/datadog.go     # feedback submission (eval-metric API)
│   └── evaljudge/              # golden dataset loader, LLM-as-judge, scoring
├── evals/
│   ├── golden.jsonl            # {question, expected_points[], rubric}
│   └── judge_prompt.md
├── web/
│   ├── index.html
│   ├── app.js
│   └── style.css
└── migrations/
    └── 001_init.sql            # tables + pgvector extension + index

4.2 Core interfaces
Small, single-method interfaces defined at their consumers — standard Go style. This keeps the RAG pipeline testable with fakes and makes each pluggable component replaceable.
// internal/rag/retriever.go
type Chunk struct {
    ID       string
    DocTitle string
    Section  string
    Text     string
    Score    float64
}

type Retriever interface {
    Search(ctx context.Context, queryVec []float64, k int) ([]Chunk, error)
}

// internal/rag/embed.go
type Embedder interface {
    Embed(ctx context.Context, texts []string) ([][]float64, error)
}

// internal/llm/client.go
type ChatRequest struct {
    Model       string
    System      string
    Messages    []Message
    Temperature float64
    MaxTokens   int
}

type LLMClient interface {
    Chat(ctx context.Context, req ChatRequest) (ChatResponse, error)
}

4.3 Configuration
All configuration via environment variables (see .env.example). Key variables:
Variable
Purpose
Example
DATABASE_URL
Postgres DSN
postgres://app:app@localhost:5432/theory
LLM_BASE_URL
OpenAI-compatible chat base URL. Demo provider is OpenRouter.
https://openrouter.ai/api/v1
LLM_EMBED_BASE_URL
Embeddings base URL. Empty means use LLM_BASE_URL.
—
LLM_API_KEY
Provider key
—
LLM_CHAT_MODEL
Answer and rewrite model. OpenRouter slug for a small instruct model.
<org>/<small-instruct-model>
LLM_EMBED_MODEL
Embedding model. Its output width must equal EMBEDDING_DIMENSIONS.
text-embedding-3-small
EMBEDDING_DIMENSIONS
Vector width. Must match vector(N) in the applied migration. api and ingest refuse to start when they differ.
1536
RETRIEVAL_TOP_K
Chunks retrieved per query, before the score floor
5
RETRIEVAL_MIN_SCORE
Minimum cosine similarity. Chunks below this are dropped. None left means the canned answer.
0.5
HISTORY_MESSAGE_LIMIT
Messages from the conversation passed to the answer model
10
REWRITE_HISTORY_MESSAGES
Messages passed to rewrite.query. Used only when history is non-empty.
4
REWRITE_MAX_TOKENS
Max tokens for rewrite.query. Temperature is 0.
64
DD_LLMOBS_ML_APP
App name in Agent Observability
guitar-chat
DD_LLMOBS_AGENTLESS_ENABLED
Agentless span submission. This is the demo export path.
true
DD_API_KEY / DD_SITE
Datadog credentials and site. US1.
datadoghq.com
JUDGE_MODEL
Model used for LLM-as-judge. May be the same small OpenRouter model.
<org>/<small-instruct-model>
ONLINE_EVAL_SAMPLE_RATE
Fraction of generated answers judged. The canned empty answer is never judged.
1.0


The llmobs SDK's global setup (start/stop entry point, agentless vs Agent-based export) is not yet documented on Datadog's docs site — the package predates official documentation. The environment variables above are the expected mirror of the Python/Node/Java SDKs (SDK reference), but the exact llmobs.Start(...) options and defaults must be verified against the package source/README during implementation — e.g., go get github.com/DataDog/dd-trace-go/v2@latest and inspect llmobs/options.go for the authoritative StartOption list. Pin the resolved version exactly in go.mod. Export for this demo is agentless to US1: DD_SITE=datadoghq.com, API host https://api.datadoghq.com.

5. RAG Pipeline
5.1 Ingestion (cmd/ingest)
Read every corpus/*.md file. Skip corpus/LICENSE and corpus/ATTRIBUTION.md.
Chunk with heading-aware splitting: split on ## headings, then sub-split any block exceeding ~800 tokens at paragraph boundaries, keeping ~100 tokens of overlap. Each chunk carries {doc_title, section, text}.
Embed chunk texts in batches of 64 via Embedder.
Upsert the document and its chunks. Chunk ids are sha256(doc_id + section + text), so re-running does not duplicate identical rows.
For each file, delete chunk rows of that document whose ids are not in the set just written.
After the corpus walk, delete documents whose slugs are no longer in corpus/, and their chunks. Re-ingesting a degraded file or removing a file leaves only the current corpus searchable. This is what makes demo step 6 (groundedness drops after a degraded re-ingest) true.
Run manually: go run ./cmd/ingest. Also runnable inside the compose stack via a one-shot service. Refuse to run when EMBEDDING_DIMENSIONS does not match the chunks.embedding column width.
5.2 Answer path (POST /api/chat)
Validate question (non-empty, ≤ 2000 chars).
Resolve the conversation. A null conversation_id mints a new UUID and inserts the row. A non-null id that is not in conversations is inserted as given, so a localStorage id works.
Load the last HISTORY_MESSAGE_LIMIT messages (default 10).
Build the search query. When that history is empty, the query is the raw question and rewrite.query is not called. When it is not empty, one completion rewrites the last REWRITE_HISTORY_MESSAGES messages (default 4) into a standalone search query: temperature 0, max tokens REWRITE_MAX_TOKENS (default 64). The span is an llm span named rewrite.query. On error, timeout, or blank output, embed the raw question and continue. Do not embed prior assistant answers. The stored user message stays verbatim.
Embed the search query (embedding span). The span records the text that was embedded.
Retrieve top-K chunks by cosine similarity in pgvector (retrieval span). Annotate every returned score, including ones the floor later rejects.
Drop chunks with similarity < RETRIEVAL_MIN_SCORE (default 0.5). If none remain, or retrieval fails, return answer "I don't have enough context", sources [], and the conversation and trace ids. Do not call the answer model. Do not start generate.answer. Do not run the online judge.
Otherwise assemble the system prompt:
You are a guitar music theory tutor. Answer using ONLY the provided
context. If the context does not support an answer, say "I don't have enough context"
and do not add theory that is not in the context. Cite sources like [1], [2]
matching the numbered context chunks. Keep answers under 300 words
unless asked for more. Use fretboard-friendly language (string/fret,
shapes, intervals in semitones) when helpful.
Call the LLM with the original question, the surviving chunks, and the last HISTORY_MESSAGE_LIMIT messages (llm span generate.answer with model, provider, and token metrics).
Persist the user message verbatim and the assistant answer.
Return {answer, sources[], conversation_id, trace_id} — trace_id comes from the root workflow span's TraceID() method, letting the frontend attach feedback to the exact Agent Observability trace.
5.3 Data model
CREATE EXTENSION IF NOT EXISTS vector;

CREATE TABLE documents (
    id          TEXT PRIMARY KEY,      -- slug of corpus file
    title       TEXT NOT NULL,
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE chunks (
    id          TEXT PRIMARY KEY,      -- sha256(doc_id + section + text)
    document_id TEXT NOT NULL REFERENCES documents(id) ON DELETE CASCADE,
    section     TEXT NOT NULL,
    text        TEXT NOT NULL,
    embedding   vector(1536) NOT NULL  -- dimension = embedding model output
);

CREATE INDEX chunks_embedding_idx ON chunks
    USING hnsw (embedding vector_cosine_ops);

CREATE TABLE conversations (
    id         UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE messages (
    id              BIGSERIAL PRIMARY KEY,
    conversation_id UUID NOT NULL REFERENCES conversations(id),
    role            TEXT NOT NULL CHECK (role IN ('user','assistant')),
    content         TEXT NOT NULL,
    trace_id        TEXT,             -- Datadog trace id, for feedback joins
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE feedback (
    id              BIGSERIAL PRIMARY KEY,
    conversation_id UUID NOT NULL,
    trace_id        TEXT NOT NULL,
    rating          SMALLINT NOT NULL CHECK (rating IN (-1, 1)),
    comment         TEXT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE eval_runs (
    id          BIGSERIAL PRIMARY KEY,
    kind        TEXT NOT NULL,        -- 'offline'; one row per cmd/eval run. Online judgments are not stored here.
    dataset     TEXT,                 -- 'golden' when offline
    score       REAL,
    passed      INTEGER,
    total       INTEGER,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

vector(N) is chosen once, in 001_init.sql, before that migration is applied, and EMBEDDING_DIMENSIONS must be the same N. 1536 matches text-embedding-3-small. A different embedding model needs its width written into the migration before the first apply. After the migration is applied, a width change is a new migration. api and ingest compare the env value to the column type and exit on mismatch. The chat model on OpenRouter does not determine this width.

6. API Contracts
POST /api/chat
// request
{ "conversation_id": "uuid-or-null", "question": "What's the difference between Mixolydian and Ionian?" }

conversation_id null mints a new conversation. A non-null id missing from the table is created with that id. The user question is stored verbatim even when retrieval used a rewritten query.

// response 200
{
  "conversation_id": "…",
  "answer": "…tutor answer with [1] citations…",
  "sources": [
    { "n": 1, "title": "Modes", "section": "Mixolydian", "score": 0.83 }
  ],
  "trace_id": "1234567890123456789"
}

POST /api/feedback
// request
{ "conversation_id": "…", "trace_id": "…", "rating": 1, "comment": "optional" }
// response 204

GET /healthz → 200 {"status":"ok"} when the process is up. It does not check Postgres or the LLM. Compose waits on the db healthcheck, not on this route, before starting ingest.
Errors are {"error": {"code": "...", "message": "..."}} with proper status codes.

7. Frontend (Vanilla JS)
web/index.html + web/app.js + web/style.css — served by Go's http.FileServer at /.
Chat window: message list, input box, send button; Enter submits.
Each assistant message renders markdown-lite (bold, code, numbered citations) and shows source chips (doc title + section) beneath the answer.
Thumbs up / thumbs down buttons on every assistant message call POST /api/feedback with the stored trace_id.
Conversation id kept in localStorage; "New chat" button resets it.
Loading state disables input while a request is in flight; errors render inline in the thread.
No build step, no npm, no framework.

8. Evaluations
Two layers, both ultimately landing in Datadog:
8.1 Offline eval harness (cmd/eval)
evals/golden.jsonl: 30–50 curated questions with expected_points (facts a correct answer must contain) and a rubric (tone, guitar-friendliness, citation usage). Sample rows: mode identification, chord-scale relationships, interval fretboard math, enharmonic equivalents, chord construction.
For each row, run the full RAG pipeline (not just the LLM), then an LLM-as-judge pass scoring 0–2 on: correctness (contains expected points, no contradictions), groundedness (claims supported by retrieved context), guitar-friendliness. Judge output is strict JSON; parse failures count as failures (correctness 0, and the row is not passed). Golden rows have no history, so rewrite.query does not run.
Per-row scores are submitted to Datadog as Agent Observability evaluations via llmobs.SubmitEvaluationFromSpan (joined to that row's trace) or SubmitEvaluationFromTag (whole-run joins with a tag like eval_run:<id>), so quality trends appear next to traces.
At the end of the run, insert one eval_runs row and no more: kind offline, dataset golden, score = mean correctness (parse failures count as 0), passed = rows with a parsed judge result and groundedness > 0, total = golden rows attempted.
Gate: the process fails if mean correctness < 0.8 or any single row scores 0 on groundedness. CI fails the build on that gate. This makes prompt/model changes measurable before merging.
8.2 Online evals (in-request)
After each generated answer (sampled at ONLINE_EVAL_SAMPLE_RATE, default 100% for the demo), an asynchronous judge call scores the answer for groundedness and helpfulness using the retrieved context as reference. It does not run when the answer was the canned empty-retrieval string. It runs after the response has been returned, so it is deliberately not a child span of the answer trace — the root workflow span is finished when the handler returns, and holding it open would inflate request latency.
Instead, the handler captures the root span's SpanID()/TraceID() (both exposed by the typed spans) and the judge goroutine starts its own small judge.answer workflow span, linking back to the answer trace with AddLink(SpanLink{...}). llmobs.SubmitEvaluationFromSpan accepts any value implementing the EvaluatedSpan interface (SpanID() + TraceID()), so a tiny struct capturing those IDs joins the evaluation to the original answer trace even after it is finished.
Online scores are not written to eval_runs. They exist on the Datadog trace.
Judge failures never block the user response; they are logged with the trace id.
8.3 End-user feedback
Thumbs up/down from the UI is submitted both to the local feedback table and to Datadog. The Go SDK does not document a feedback API yet, so this uses the eval-metric HTTP endpoint with event_kind=feedback, joined to the trace by trace_id (or a feedback_join_key tag set on the root span), which is how Agent Observability links end-user feedback to traces (Evaluations API).
Datadog's own evaluation features complement this: production traces can be promoted into versioned datasets, experiments can compare prompt/model changes against the same golden data, and annotation queues support human review — all within Agent Observability (Agent Observability product). The harness built here also serves as the demo of external evaluations via API.

9. Datadog Agent Observability Integration
9.1 Span model
One trace per chat request, built with the llmobs package's typed span constructors. Span kinds map one-to-one to Datadog's taxonomy (workflow, agent, tool, task, embedding, retrieval, llm), where an LLM span is one call to an LLM provider and is the only billable span type — workflow, tool, agent, embedding, and retrieval spans are free (Datadog product page). Do not emit an llm span on a path that did not call the provider.
Pipeline step
llmobs call
Annotation
Chat request root
StartWorkflowSpan(ctx, "chat.answer", WithSessionID(conversationID))
AnnotateTextIO(question, answer)
Query rewrite (follow-ups only)
StartLLMSpan(ctx, "rewrite.query", WithModelName(chatModel), WithModelProvider(...))
AnnotateLLMIO with the last REWRITE_HISTORY_MESSAGES and the standalone query. Skipped on the first turn. On failure the span finishes with the error and retrieval embeds the raw question.
Question embedding
StartEmbeddingSpan(ctx, "embed.question", WithModelName(embedModel))
AnnotateEmbeddingIO with an EmbeddedDocument for the search query that was embedded (rewrite output, or the raw question), plus token metrics
Vector search
StartRetrievalSpan(ctx, "retrieve.chunks")
AnnotateRetrievalIO with that same query and scored RetrievedDocument values (text, score, doc/section metadata), including chunks below RETRIEVAL_MIN_SCORE
LLM answer call
StartLLMSpan(ctx, "generate.answer", WithModelName(...), WithModelProvider(...))
Present only when at least one chunk passed the floor. AnnotateLLMIO with prompt/response messages, token metrics (MetricKeyInputTokens, MetricKeyOutputTokens, MetricKeyTotalTokens), metadata (temperature, top_k)
Empty retrieval
no generate.answer span
Workflow output is the exact string "I don't have enough context". sources is empty. Online judge does not start.
Online judge call
its own StartWorkflowSpan(ctx, "judge.answer") after a generated answer returns, linked to the answer trace via AddLink (see §8.2)
same annotations + judge metadata. Not started for the canned empty answer.
Error paths
span.Finish(WithError(err)) on any failure
errors captured with stack traces


(Field names above are descriptive, not literal — the exact EmbeddedDocument/RetrievedDocument struct fields and StartSpanOptions must be taken from the SDK API reference at implementation time, since the package is experimental and may change.)
All llmobs calls go through internal/observe helpers that wrap the SDK, so the experimental dependency is isolated to one package and call sites stay stable if the SDK API shifts. Each typed span exposes SpanID()/TraceID() (plus APMTraceID() for APM correlation), which are the join keys for evaluation submission (Go API reference).
9.2 Export
The SDK supports agentless submission (spans sent directly to Datadog, no local Agent) and Agent-based export. This demo uses agentless export to US1: DD_LLMOBS_AGENTLESS_ENABLED=true, DD_SITE=datadoghq.com, intake at https://api.datadoghq.com. Because the Go package's setup API is not yet documented, confirm the exact start options, defaults, and environment variables against the package source (see §4.3) — the expected variables mirror the other SDKs (SDK reference). The optional compose datadog-agent service is not the demo path.
9.3 Evaluations and feedback submission
Judge scores (online + offline): llmobs.SubmitEvaluationFromSpan(label, value, span, opts...) — numeric scores for groundedness/helpfulness/correctness, joined to the trace by span/trace ID; SubmitEvaluationFromTag(label, value, JoinTag{Key, Value}) is available for batch joins such as tagging all spans of an eval run (Go API reference).
End-user feedback: the Go SDK does not yet document a feedback API, so thumbs up/down goes through the eval-metric HTTP endpoint with event_kind=feedback, joined by trace_id or a feedback_join_key tag set on the root span (Evaluations API). This is a clean example of mixing SDK and API ingestion in one app.
9.4 Dashboards, monitors, logs
Dashboard: request rate, p95 answer latency, LLM token spend, judge score distribution, thumbs up/down ratio, retrieval hit count.
Monitors: groundedness score < 0.7 over 30 min; thumbs-down ratio > 20%; LLM error rate; empty-retrieval rate.
Logs: Go logs via slog JSON, correlated to traces via the span's APMTraceID(); eval failures and low-confidence answers logged with the LLMObs trace id.
9.5 Demo storyline
The point of the app is the demo: (1) ask questions in the UI, (2) open the trace in US1 Agent Observability and walk the span tree — question, query rewrite on a follow-up, embedding, retrieval with scores, LLM answer with token counts, (3) show judge scores attached to the trace, (4) thumbs-down a bad answer and see the feedback land on that trace, (5) run the offline harness and show the single eval_runs row plus score trends, (6) re-ingest a degraded corpus (old chunks for that file are deleted) and watch groundedness scores drop — proving the evals catch regressions. An off-corpus question shows the canned "I don't have enough context" answer, a trace with rejected retrieval scores, and no generate.answer span.

10. Deployment
docker compose up:
db: pgvector/pgvector:pg16 with migrations applied on start.
app: Go binary in a distroless image; web/ embedded via go:embed; env from .env.
ingest (one-shot): runs after db health, then exits.
Optional datadog-agent service with OTLP ingest enabled. Not used by the demo; export is agentless to US1.
No other infrastructure. Local dev: go run ./cmd/api + local Postgres.

11. Security and Safety Notes
No user auth (demo) — do not log or send more than necessary to providers; Sensitive Data Scanner rules can be applied in Datadog, and Agent Observability integrates with it for redaction (Evaluations).
Secrets only via environment; .env.example documents every variable; .env git-ignored.
Prompt-injection surface is low (corpus is curated), but the system prompt still constrains answers to provided context and the judge checks groundedness.
Input length caps and request timeouts (30s answer path, 60s judge path) everywhere.

12. Milestones
#
Milestone
Demo check
1
Go server, /healthz, chat UI shell, Postgres + migrations
page loads, health green
2
Ingestion + pgvector retrieval, /api/chat end-to-end
grounded answers with sources
3
llmobs spans → Datadog US1, full span tree per request
trace visible in Agent Observability
4
Offline eval harness + golden dataset + CI gate
go run ./cmd/eval scores printed
5
Online judge + eval-metric submission + feedback buttons
scores and 👍/👎 on traces
6
Dashboard, monitors, demo storyline walkthrough
full demo flow


Definition of done (per milestone 3+)
Every chat request yields exactly one US1 trace. It always contains embed.question and retrieve.chunks. Follow-ups also contain rewrite.query. generate.answer, with model, token, and latency attributes, is present when the answer model ran, and absent when retrieval fell below RETRIEVAL_MIN_SCORE.
Every judged answer has eval scores visible on its trace in Datadog.
Every 👍/👎 is visible on the corresponding trace.
Offline eval gate passes in CI.

13. Future Work
SSE token streaming
Hybrid retrieval (pg full-text + vector, RRF fusion) and cross-encoder reranking
Conversation memory summarization for long chats
Datadog experiments: A/B prompt versions against the golden dataset
Auth + per-user session evaluation in Agent Observability

14. References
dd-trace-go llmobs package (experimental Go SDK, source)
llmobs Go API reference (pkg.go.dev)
Datadog Agent Observability product page
Datadog Agent Observability docs
SDK reference (Python, Node.js, Java; span kinds, ml_app)
OpenTelemetry instrumentation for Agent Observability
Agent Observability HTTP API (spans + evaluations)
External evaluations (source:otel, decimal span IDs)
Evaluations in Agent Observability
Datadog blog: OTel GenAI semantic conventions support
OpenTelemetry GenAI semantic conventions
pgvector
Open Music Theory (corpus candidate, CC BY-SA)
