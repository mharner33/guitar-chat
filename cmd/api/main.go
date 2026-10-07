// Command api serves the Guitar Theory Chat HTTP API and the embedded web UI.
//
// Phase 0: loads config, connects to Postgres, applies migrations, and enforces
// the embedding-dimension guard. The chi server and routes are added in Phase 1
// (see docs/plan.md).
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
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := run(); err != nil {
		slog.Error("startup failed", "err", err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		return err
	}
	slog.Info("configuration loaded",
		"addr", cfg.Addr,
		"chat_model", cfg.LLMChatModel,
		"embed_model", cfg.LLMEmbedModel,
		"embedding_dimensions", cfg.EmbeddingDimensions,
		"observability_enabled", cfg.ObservabilityEnabled(),
	)

	st, err := store.Open(ctx, cfg.DatabaseURL)
	if err != nil {
		return err
	}
	defer st.Close()

	if err := st.Migrate(ctx, migrations.FS); err != nil {
		return err
	}
	slog.Info("migrations applied")

	if err := st.CheckEmbeddingDimensions(ctx, cfg.EmbeddingDimensions); err != nil {
		return err
	}
	slog.Info("embedding dimensions verified", "dimensions", cfg.EmbeddingDimensions)

	// Phase 1 wires the chi server and routes here.
	return nil
}
