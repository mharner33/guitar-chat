// Package corpus reads the ingest sources — the generated musicwiki OKF bundle
// and the hand-written guides — into documents ready to embed. It knows the
// source formats; cmd/ingest only embeds and stores what it returns.
package corpus

import "github.com/mharner33/guitar-chat/internal/rag"

// Document id prefixes. Each source owns one, and reconcile is scoped to it.
const (
	WikiPrefix  = "wiki:"
	GuidePrefix = "guide:"
)

// Chunk is one embeddable section of a document.
type Chunk = rag.Chunk

// Document is one source file, ready to embed.
type Document struct {
	ID     string // "wiki:/scales/…" or "guide:…"
	Title  string
	Chunks []Chunk
}
