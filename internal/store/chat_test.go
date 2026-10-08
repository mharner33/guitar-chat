package store

import (
	"context"
	"os"
	"testing"
	"time"
)

// Integration test: requires a live DATABASE_URL with migrations applied.
// Skips when unset so `go test ./...` stays network-free by default.
func TestEnsureConversationAndAddMessage(t *testing.T) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Skip("DATABASE_URL not set; skipping store integration test")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	st, err := Open(ctx, dsn)
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	defer st.Close()

	// Mint a new conversation.
	id, err := st.EnsureConversation(ctx, "")
	if err != nil {
		t.Fatalf("EnsureConversation(new): %v", err)
	}
	if id == "" {
		t.Fatal("EnsureConversation(new) returned empty id")
	}

	// Clean up whatever we insert, regardless of outcome.
	t.Cleanup(func() {
		_, _ = st.Pool().Exec(context.Background(), `DELETE FROM messages WHERE conversation_id = $1`, id)
		_, _ = st.Pool().Exec(context.Background(), `DELETE FROM conversations WHERE id = $1`, id)
	})

	// Idempotent for an existing id.
	again, err := st.EnsureConversation(ctx, id)
	if err != nil {
		t.Fatalf("EnsureConversation(existing): %v", err)
	}
	if again != id {
		t.Fatalf("EnsureConversation(existing) = %q, want %q", again, id)
	}

	if err := st.AddMessage(ctx, id, "user", "test question"); err != nil {
		t.Fatalf("AddMessage(user): %v", err)
	}
	if err := st.AddMessage(ctx, id, "assistant", "test answer"); err != nil {
		t.Fatalf("AddMessage(assistant): %v", err)
	}

	var count int
	if err := st.Pool().QueryRow(ctx,
		`SELECT count(*) FROM messages WHERE conversation_id = $1`, id).Scan(&count); err != nil {
		t.Fatalf("count messages: %v", err)
	}
	if count != 2 {
		t.Fatalf("message count = %d, want 2", count)
	}
}
