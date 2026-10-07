// Command ingest chunks and embeds the corpus into pgvector. It is idempotent
// and reconciles the database to the current corpus/ directory.
//
// Phase 0: loads config and enforces the embedding-dimension guard. The
// chunking, embedding, and reconcile logic are added in Phase 2 (see docs/plan.md).
package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/mharner33/guitar-chat/internal/config"
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

	cfg, err := config.Load()
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

	slog.Info("ingest: config and schema verified; corpus ingestion added in Phase 2")
	return nil
}
