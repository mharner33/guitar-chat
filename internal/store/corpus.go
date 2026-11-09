package store

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
)

// ChunkID is a chunk's content hash: sha256 over doc id, section and text,
// each NUL-terminated so field boundaries count. Identical content keeps the
// same id across runs, which is what makes ingest idempotent.
func ChunkID(docID, section, text string) string {
	h := sha256.New()
	for _, part := range []string{docID, section, text} {
		h.Write([]byte(part))
		h.Write([]byte{0})
	}
	return hex.EncodeToString(h.Sum(nil))
}

// NewChunk is a chunk to insert, with its embedding.
type NewChunk struct {
	ID        string
	Section   string
	Text      string
	Embedding []float32
}

// ChunkIDs returns the ids of a document's stored chunks.
func (s *Store) ChunkIDs(ctx context.Context, docID string) (map[string]bool, error) {
	rows, err := s.pool.Query(ctx, `SELECT id FROM chunks WHERE document_id = $1`, docID)
	if err != nil {
		return nil, fmt.Errorf("list chunks of %s: %w", docID, err)
	}
	ids, err := pgx.CollectRows(rows, pgx.RowTo[string])
	if err != nil {
		return nil, fmt.Errorf("list chunks of %s: %w", docID, err)
	}
	out := make(map[string]bool, len(ids))
	for _, id := range ids {
		out[id] = true
	}
	return out, nil
}

// SyncDocument makes a document's stored chunks exactly keepIDs, in one
// transaction: upsert the document, insert added chunks, delete the rest.
func (s *Store) SyncDocument(ctx context.Context, docID, title string, keepIDs []string, added []NewChunk) error {
	if keepIDs == nil {
		keepIDs = []string{} // nil would encode as NULL, and NOT (id = ANY(NULL)) deletes nothing
	}
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("sync %s: begin: %w", docID, err)
	}
	defer tx.Rollback(ctx) //nolint:errcheck // no-op after a successful Commit

	if _, err := tx.Exec(ctx, `
		INSERT INTO documents (id, title) VALUES ($1, $2)
		ON CONFLICT (id) DO UPDATE SET title = EXCLUDED.title, updated_at = now()`,
		docID, title); err != nil {
		return fmt.Errorf("sync %s: upsert document: %w", docID, err)
	}
	for _, c := range added {
		if _, err := tx.Exec(ctx, `
			INSERT INTO chunks (id, document_id, section, text, embedding)
			VALUES ($1, $2, $3, $4, $5::text::vector)
			ON CONFLICT (id) DO NOTHING`,
			c.ID, docID, c.Section, c.Text, vectorLiteral(c.Embedding)); err != nil {
			return fmt.Errorf("sync %s: insert chunk %q: %w", docID, c.Section, err)
		}
	}
	if _, err := tx.Exec(ctx,
		`DELETE FROM chunks WHERE document_id = $1 AND NOT (id = ANY($2))`,
		docID, keepIDs); err != nil {
		return fmt.Errorf("sync %s: delete stale chunks: %w", docID, err)
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("sync %s: commit: %w", docID, err)
	}
	return nil
}

// DeleteMissingDocuments deletes documents whose id starts with prefix and is
// not in keep; their chunks cascade. Scoping by prefix keeps one source's
// reconcile from touching another's documents.
func (s *Store) DeleteMissingDocuments(ctx context.Context, prefix string, keep []string) (int64, error) {
	if prefix == "" {
		return 0, errors.New("delete missing documents: empty prefix would match every document")
	}
	if keep == nil {
		keep = []string{}
	}
	tag, err := s.pool.Exec(ctx,
		`DELETE FROM documents WHERE starts_with(id, $1) AND NOT (id = ANY($2))`,
		prefix, keep)
	if err != nil {
		return 0, fmt.Errorf("delete missing %s documents: %w", prefix, err)
	}
	return tag.RowsAffected(), nil
}

// vectorLiteral formats v as a pgvector text literal, e.g. "[0.5,-1]".
func vectorLiteral(v []float32) string {
	var b strings.Builder
	b.WriteByte('[')
	for i, f := range v {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteString(strconv.FormatFloat(float64(f), 'g', -1, 32))
	}
	b.WriteByte(']')
	return b.String()
}
