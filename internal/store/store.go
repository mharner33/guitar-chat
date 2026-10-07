// Package store owns the Postgres connection pool and all SQL. It uses pgx/v5
// directly — no ORM, no codegen (see agent.md conventions).
package store

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"sort"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Store wraps a pgx connection pool.
type Store struct {
	pool *pgxpool.Pool
}

// Open connects to Postgres using the given DSN and verifies connectivity.
func Open(ctx context.Context, databaseURL string) (*Store, error) {
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, fmt.Errorf("connect postgres: %w", err)
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping postgres: %w", err)
	}
	return &Store{pool: pool}, nil
}

// Close releases the pool.
func (s *Store) Close() { s.pool.Close() }

// Pool exposes the underlying pool for packages that run their own queries.
func (s *Store) Pool() *pgxpool.Pool { return s.pool }

// Migrate applies every *.sql file in fsys, in lexical order, that has not yet
// been applied. Applied versions are tracked in schema_migrations. Each file is
// run in its own transaction, so a failure leaves the database at the last
// fully-applied version. Migrations are forward-only.
func (s *Store) Migrate(ctx context.Context, fsys fs.FS) error {
	if _, err := s.pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version    TEXT PRIMARY KEY,
			applied_at TIMESTAMPTZ NOT NULL DEFAULT now()
		)`); err != nil {
		return fmt.Errorf("create schema_migrations: %w", err)
	}

	entries, err := fs.Glob(fsys, "*.sql")
	if err != nil {
		return fmt.Errorf("list migrations: %w", err)
	}
	sort.Strings(entries)

	for _, name := range entries {
		var exists bool
		if err := s.pool.QueryRow(ctx,
			`SELECT EXISTS (SELECT 1 FROM schema_migrations WHERE version = $1)`, name,
		).Scan(&exists); err != nil {
			return fmt.Errorf("check migration %s: %w", name, err)
		}
		if exists {
			continue
		}

		sqlBytes, err := fs.ReadFile(fsys, name)
		if err != nil {
			return fmt.Errorf("read migration %s: %w", name, err)
		}

		if err := s.applyMigration(ctx, name, string(sqlBytes)); err != nil {
			return fmt.Errorf("apply migration %s: %w", name, err)
		}
	}
	return nil
}

func (s *Store) applyMigration(ctx context.Context, name, sql string) error {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx) //nolint:errcheck // no-op after a successful Commit

	if _, err := tx.Exec(ctx, sql); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx,
		`INSERT INTO schema_migrations (version) VALUES ($1)`, name,
	); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

// CheckEmbeddingDimensions verifies that the chunks.embedding column's vector
// width equals want. api and ingest call this at startup and refuse to run on a
// mismatch, because a disagreement between the configured embedding model and
// the applied migration silently corrupts retrieval.
func (s *Store) CheckEmbeddingDimensions(ctx context.Context, want int) error {
	// pgvector stores the declared dimension directly in atttypmod (-1 when
	// the column is an unconstrained vector).
	var typmod int
	err := s.pool.QueryRow(ctx, `
		SELECT a.atttypmod
		FROM pg_attribute a
		JOIN pg_class c ON c.oid = a.attrelid
		WHERE c.relname = 'chunks' AND a.attname = 'embedding' AND a.attnum > 0`,
	).Scan(&typmod)
	if errors.Is(err, pgx.ErrNoRows) {
		return errors.New("chunks.embedding column not found; run migrations first")
	}
	if err != nil {
		return fmt.Errorf("inspect embedding column: %w", err)
	}
	if typmod != want {
		return fmt.Errorf("EMBEDDING_DIMENSIONS=%d does not match chunks.embedding vector(%d); "+
			"a width change requires a new migration", want, typmod)
	}
	return nil
}
