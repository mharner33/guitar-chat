package store

import (
	"context"
	"fmt"
)

// EnsureConversation returns the canonical conversation id. An empty id mints a
// new conversation (gen_random_uuid()); a non-empty id is created if absent and
// returned unchanged, so a client-held id (e.g. from localStorage) works.
func (s *Store) EnsureConversation(ctx context.Context, id string) (string, error) {
	if id == "" {
		var newID string
		if err := s.pool.QueryRow(ctx,
			`INSERT INTO conversations (id) VALUES (gen_random_uuid()) RETURNING id`,
		).Scan(&newID); err != nil {
			return "", fmt.Errorf("insert conversation: %w", err)
		}
		return newID, nil
	}
	if _, err := s.pool.Exec(ctx,
		`INSERT INTO conversations (id) VALUES ($1) ON CONFLICT (id) DO NOTHING`, id,
	); err != nil {
		return "", fmt.Errorf("ensure conversation %s: %w", id, err)
	}
	return id, nil
}

// AddMessage stores one chat message. role must be 'user' or 'assistant'
// (enforced by a CHECK constraint). trace_id is left NULL until Phase 3.
func (s *Store) AddMessage(ctx context.Context, conversationID, role, content string) error {
	if _, err := s.pool.Exec(ctx,
		`INSERT INTO messages (conversation_id, role, content) VALUES ($1, $2, $3)`,
		conversationID, role, content,
	); err != nil {
		return fmt.Errorf("add message: %w", err)
	}
	return nil
}
