package server

import (
	"context"
	"io/fs"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/mharner33/guitar-chat/internal/config"
)

// ChatStore is the persistence the chat handler needs. Declared here (the
// consumer) so it can be faked in tests; *store.Store implements it.
type ChatStore interface {
	EnsureConversation(ctx context.Context, id string) (string, error)
	AddMessage(ctx context.Context, conversationID, role, content string) error
}

type api struct {
	chat ChatStore
}

// New builds the HTTP handler: middleware, health, chat, and the embedded UI.
// cfg is carried for the seam that later phases (retrieval/history knobs) need.
func New(cfg config.Config, chat ChatStore, assets fs.FS) http.Handler {
	_ = cfg // reserved for Phase 2+
	a := &api{chat: chat}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(accessLog)
	r.Use(middleware.Timeout(30 * time.Second))

	r.Get("/healthz", handleHealth)
	r.Post("/api/chat", a.handleChat)
	r.Handle("/*", http.FileServerFS(assets))

	return r
}

func handleHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// accessLog emits one slog-JSON line per request.
func accessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)
		slog.Info("http request",
			"method", r.Method,
			"path", r.URL.Path,
			"status", ww.Status(),
			"bytes", ww.BytesWritten(),
			"duration_ms", time.Since(start).Milliseconds(),
			"request_id", middleware.GetReqID(r.Context()),
		)
	})
}
