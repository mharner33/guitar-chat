// Command api serves the Guitar Theory Chat HTTP API and the embedded web UI.
//
// It loads config, connects to Postgres, applies migrations, verifies the
// embedding-dimension guard, then serves the chi router until SIGINT/SIGTERM.
package main

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

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := run(); err != nil {
		slog.Error("api exited with error", "err", err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()

	// Local dev convenience: pick up ./.env if present. Real environment
	// variables take precedence; in containers the file is absent.
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
	slog.Info("configuration loaded",
		"addr", cfg.Addr,
		"chat_model", cfg.LLMChatModel,
		"embed_model", cfg.LLMEmbedModel,
		"embedding_dimensions", cfg.EmbeddingDimensions,
		"observability_enabled", cfg.ObservabilityEnabled(),
	)

	// Startup DB work runs under a deadline. The pool's lifetime is not tied to
	// this context: pgxpool.New uses ctx only for the initial connect.
	startCtx, cancelStart := context.WithTimeout(ctx, 30*time.Second)
	defer cancelStart()

	st, err := store.Open(startCtx, cfg.DatabaseURL)
	if err != nil {
		return err
	}
	defer st.Close()

	if err := st.Migrate(startCtx, migrations.FS); err != nil {
		return err
	}
	slog.Info("migrations applied")

	if err := st.CheckEmbeddingDimensions(startCtx, cfg.EmbeddingDimensions); err != nil {
		return err
	}
	cancelStart()
	slog.Info("embedding dimensions verified", "dimensions", cfg.EmbeddingDimensions)

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
}
