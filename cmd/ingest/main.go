// Command ingest chunks and embeds the corpus sources listed in CORPUS_DIRS
// (the musicwiki bundle and the hand-written guides) into pgvector. It is
// idempotent and reconciles each source's documents to what is on disk.
package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/mharner33/guitar-chat/internal/config"
	"github.com/mharner33/guitar-chat/internal/corpus"
	"github.com/mharner33/guitar-chat/internal/rag"
	"github.com/mharner33/guitar-chat/internal/store"
	"github.com/mharner33/guitar-chat/migrations"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	if err := run(); err != nil {
		slog.Error("ingest failed", "err", err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()

	// Local dev convenience, same as cmd/api: pick up ./.env if present. Real
	// environment variables take precedence; in containers the file is absent.
	loaded, err := config.LoadDotEnv(".env")
	if err != nil {
		return err
	}
	if loaded {
		slog.Info("loaded .env")
	}

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	// Read and validate every source before touching the database, so a bad
	// file or a mistyped directory cannot delete anything.
	sources, err := corpus.Load(cfg.CorpusDirs)
	if err != nil {
		return err
	}

	st, err := store.Open(ctx, cfg.DatabaseURL)
	if err != nil {
		return err
	}
	defer st.Close()

	if err := st.Migrate(ctx, migrations.FS); err != nil {
		return err
	}
	if err := st.CheckEmbeddingDimensions(ctx, cfg.EmbeddingDimensions); err != nil {
		return err
	}

	emb := &rag.HTTPEmbedder{
		BaseURL:    cfg.EmbedBaseURL(),
		APIKey:     cfg.LLMAPIKey,
		Model:      cfg.LLMEmbedModel,
		Dimensions: cfg.EmbeddingDimensions,
		Client:     &http.Client{Timeout: embedTimeout},
	}
	s, err := ingest(ctx, st, emb, sources)
	if err != nil {
		return err
	}
	slog.Info("ingest complete",
		"documents", s.Documents, "chunks", s.Chunks, "embedded", s.Embedded, "deleted_documents", s.Deleted)
	return nil
}
