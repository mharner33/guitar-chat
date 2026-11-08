package store

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"os"
	"testing"
	"time"
)

func TestChunkID(t *testing.T) {
	a := ChunkID("wiki:/x", "overview", "text")
	if len(a) != 64 {
		t.Fatalf("len = %d, want 64 hex chars", len(a))
	}
	if a != ChunkID("wiki:/x", "overview", "text") {
		t.Error("ChunkID not deterministic")
	}
	// Field boundaries matter: moving characters between fields changes the id.
	if ChunkID("ab", "c", "d") == ChunkID("a", "bc", "d") {
		t.Error("ChunkID ignores field boundaries")
	}
}

func TestVectorLiteral(t *testing.T) {
	if got := vectorLiteral([]float32{0.5, -1, 2.25}); got != "[0.5,-1,2.25]" {
		t.Errorf("vectorLiteral = %q", got)
	}
}

// Integration: requires DATABASE_URL with migrations applied (vector(1536)).
func TestSyncAndReconcile(t *testing.T) {
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

	// Unique prefixes isolate this run from real data and other runs.
	pa, pb := testPrefix(t), testPrefix(t)
	t.Cleanup(func() {
		for _, p := range []string{pa, pb} {
			_, _ = st.Pool().Exec(context.Background(), `DELETE FROM documents WHERE starts_with(id, $1)`, p)
		}
	})
	vec := make([]float32, 1536)
	mk := func(doc, text string) NewChunk {
		return NewChunk{ID: ChunkID(doc, "s", text), Section: "s", Text: text, Embedding: vec}
	}

	docA, docB := pa+"one", pb+"one"
	c1, c2 := mk(docA, "first"), mk(docA, "second")
	if err := st.SyncDocument(ctx, docA, "One", []string{c1.ID, c2.ID}, []NewChunk{c1, c2}); err != nil {
		t.Fatalf("SyncDocument A: %v", err)
	}
	b1 := mk(docB, "b")
	if err := st.SyncDocument(ctx, docB, "B", []string{b1.ID}, []NewChunk{b1}); err != nil {
		t.Fatalf("SyncDocument B: %v", err)
	}

	// Re-sync keeping only c1: c2 is deleted, nothing re-inserted.
	if err := st.SyncDocument(ctx, docA, "One", []string{c1.ID}, nil); err != nil {
		t.Fatalf("resync: %v", err)
	}
	ids, err := st.ChunkIDs(ctx, docA)
	if err != nil || len(ids) != 1 || !ids[c1.ID] {
		t.Fatalf("ChunkIDs after resync = %v, %v; want only c1", ids, err)
	}

	// Empty keep list removes every chunk of that document (not none).
	if err := st.SyncDocument(ctx, docA, "One", nil, nil); err != nil {
		t.Fatalf("empty resync: %v", err)
	}
	if ids, _ := st.ChunkIDs(ctx, docA); len(ids) != 0 {
		t.Fatalf("ChunkIDs after empty keep = %v, want none", ids)
	}

	// Reconcile prefix A with nothing kept: deletes A's doc, not B's.
	n, err := st.DeleteMissingDocuments(ctx, pa, nil)
	if err != nil || n != 1 {
		t.Fatalf("DeleteMissingDocuments(A) = %d, %v; want 1", n, err)
	}
	if ids, _ := st.ChunkIDs(ctx, docB); len(ids) != 1 {
		t.Fatalf("source B chunks = %v, want untouched", ids)
	}
	if _, err := st.DeleteMissingDocuments(ctx, "", nil); err == nil {
		t.Fatal("DeleteMissingDocuments with empty prefix succeeded, want error")
	}
}

func testPrefix(t *testing.T) string {
	t.Helper()
	b := make([]byte, 6)
	if _, err := rand.Read(b); err != nil {
		t.Fatal(err)
	}
	return "test-" + hex.EncodeToString(b) + ":"
}
