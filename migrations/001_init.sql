-- 001_init.sql — initial schema for Guitar Theory Chat.
-- Forward-only. Never edit an applied migration; a vector-width change is a new migration.
--
-- The vector(N) width below is the single source of truth for the embedding
-- dimension. EMBEDDING_DIMENSIONS in the environment must equal this N; api and
-- ingest refuse to start on a mismatch. 1536 matches text-embedding-3-small.

CREATE EXTENSION IF NOT EXISTS vector;

CREATE TABLE IF NOT EXISTS documents (
    id          TEXT PRIMARY KEY,            -- slug of the corpus file
    title       TEXT NOT NULL,
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS chunks (
    id          TEXT PRIMARY KEY,            -- sha256(doc_id + section + text)
    document_id TEXT NOT NULL REFERENCES documents(id) ON DELETE CASCADE,
    section     TEXT NOT NULL,
    text        TEXT NOT NULL,
    embedding   vector(1536) NOT NULL
);

CREATE INDEX IF NOT EXISTS chunks_embedding_idx ON chunks
    USING hnsw (embedding vector_cosine_ops);

CREATE TABLE IF NOT EXISTS conversations (
    id         UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS messages (
    id              BIGSERIAL PRIMARY KEY,
    conversation_id UUID NOT NULL REFERENCES conversations(id),
    role            TEXT NOT NULL CHECK (role IN ('user','assistant')),
    content         TEXT NOT NULL,
    trace_id        TEXT,                     -- Datadog trace id, for feedback joins
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS messages_conversation_idx
    ON messages (conversation_id, created_at);

CREATE TABLE IF NOT EXISTS feedback (
    id              BIGSERIAL PRIMARY KEY,
    conversation_id UUID NOT NULL,
    trace_id        TEXT NOT NULL,
    rating          SMALLINT NOT NULL CHECK (rating IN (-1, 1)),
    comment         TEXT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS eval_runs (
    id          BIGSERIAL PRIMARY KEY,
    kind        TEXT NOT NULL,                -- 'offline'; online judgments are not stored here
    dataset     TEXT,                         -- 'golden' when offline
    score       REAL,
    passed      INTEGER,
    total       INTEGER,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);
