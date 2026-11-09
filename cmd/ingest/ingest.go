package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mharner33/guitar-chat/internal/corpus"
	"github.com/mharner33/guitar-chat/internal/rag"
	"github.com/mharner33/guitar-chat/internal/store"
)

const (
	embedBatchSize = 64
	embedTimeout   = 60 * time.Second
)

// docStore is the slice of *store.Store that ingest needs; tests fake it.
type docStore interface {
	ChunkIDs(ctx context.Context, docID string) (map[string]bool, error)
	SyncDocument(ctx context.Context, docID, title string, keepIDs []string, added []store.NewChunk) error
	DeleteMissingDocuments(ctx context.Context, prefix string, keep []string) (int64, error)
}

type stats struct {
	Documents int   // documents synced
	Chunks    int   // chunks present after the run
	Embedded  int   // chunks embedded this run (new or changed)
	Deleted   int64 // documents removed by reconcile
}

type docPlan struct {
	doc   corpus.Document
	keep  []string
	added []store.NewChunk
}

// ingest syncs each source in turn. Per source it plans every document, embeds
// only chunks whose ids are new (in batches across documents), and only then
// writes — so an embedding failure leaves that source untouched. Finally it
// deletes the source's documents that were not seen.
//
// Sources are processed in order and each commits independently: if a later
// source fails, earlier sources stay synced and reconciled. A SyncDocument
// failure mid-source leaves that source's earlier documents synced and its
// reconcile skipped (no rows lost; a re-run converges).
func ingest(ctx context.Context, st docStore, emb rag.Embedder, sources []corpus.Source) (stats, error) {
	var s stats
	for _, src := range sources {
		plans := make([]docPlan, 0, len(src.Docs))
		for _, d := range src.Docs {
			p, err := planDoc(ctx, st, d)
			if err != nil {
				return s, err
			}
			plans = append(plans, p)
		}

		var pending []*store.NewChunk
		for i := range plans {
			for j := range plans[i].added {
				pending = append(pending, &plans[i].added[j])
			}
		}
		if err := embedAll(ctx, emb, pending); err != nil {
			return s, fmt.Errorf("%s: %w", src.Dir, err)
		}

		keepDocs := make([]string, 0, len(plans))
		for _, p := range plans {
			if err := st.SyncDocument(ctx, p.doc.ID, p.doc.Title, p.keep, p.added); err != nil {
				return s, err
			}
			keepDocs = append(keepDocs, p.doc.ID)
			s.Documents++
			s.Chunks += len(p.keep)
			s.Embedded += len(p.added)
		}
		n, err := st.DeleteMissingDocuments(ctx, src.Prefix, keepDocs)
		if err != nil {
			return s, fmt.Errorf("reconcile %s: %w", src.Dir, err)
		}
		s.Deleted += n
	}
	return s, nil
}

func planDoc(ctx context.Context, st docStore, d corpus.Document) (docPlan, error) {
	existing, err := st.ChunkIDs(ctx, d.ID)
	if err != nil {
		return docPlan{}, err
	}
	p := docPlan{doc: d}
	seen := map[string]bool{}
	for _, c := range d.Chunks {
		id := store.ChunkID(d.ID, c.Section, c.Text)
		if seen[id] {
			continue
		}
		seen[id] = true
		p.keep = append(p.keep, id)
		if !existing[id] {
			p.added = append(p.added, store.NewChunk{ID: id, Section: c.Section, Text: c.Text})
		}
	}
	return p, nil
}

func embedAll(ctx context.Context, emb rag.Embedder, pending []*store.NewChunk) error {
	for start := 0; start < len(pending); start += embedBatchSize {
		end := min(start+embedBatchSize, len(pending))
		texts := make([]string, 0, end-start)
		for _, c := range pending[start:end] {
			texts = append(texts, c.Text)
		}
		bctx, cancel := context.WithTimeout(ctx, embedTimeout)
		vecs, err := emb.Embed(bctx, texts)
		cancel()
		if err != nil {
			return fmt.Errorf("embed: %w", err)
		}
		if len(vecs) != len(texts) {
			return fmt.Errorf("embed: got %d vectors for %d texts", len(vecs), len(texts))
		}
		for i, v := range vecs {
			pending[start+i].Embedding = v
		}
	}
	return nil
}
