package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/mharner33/guitar-chat/internal/corpus"
	"github.com/mharner33/guitar-chat/internal/store"
)

// fakeStore keeps chunks per document in memory with the real store's semantics.
type fakeStore struct {
	docs  map[string]map[string]bool // doc id → chunk ids
	syncs int
}

func newFakeStore() *fakeStore { return &fakeStore{docs: map[string]map[string]bool{}} }

func (f *fakeStore) ChunkIDs(_ context.Context, docID string) (map[string]bool, error) {
	out := map[string]bool{}
	for id := range f.docs[docID] {
		out[id] = true
	}
	return out, nil
}

func (f *fakeStore) SyncDocument(_ context.Context, docID, _ string, keep []string, added []store.NewChunk) error {
	f.syncs++
	cur := map[string]bool{}
	for _, c := range added {
		if len(c.Embedding) == 0 {
			return fmt.Errorf("chunk %s inserted without embedding", c.ID)
		}
	}
	for _, id := range keep {
		cur[id] = true
	}
	f.docs[docID] = cur
	return nil
}

func (f *fakeStore) DeleteMissingDocuments(_ context.Context, prefix string, keep []string) (int64, error) {
	kept := map[string]bool{}
	for _, id := range keep {
		kept[id] = true
	}
	var n int64
	for id := range f.docs {
		if strings.HasPrefix(id, prefix) && !kept[id] {
			delete(f.docs, id)
			n++
		}
	}
	return n, nil
}

type fakeEmbedder struct {
	batches []int
	err     error
}

func (e *fakeEmbedder) Embed(_ context.Context, texts []string) ([][]float32, error) {
	if e.err != nil {
		return nil, e.err
	}
	e.batches = append(e.batches, len(texts))
	out := make([][]float32, len(texts))
	for i := range out {
		out[i] = []float32{1}
	}
	return out, nil
}

func doc(id string, texts ...string) corpus.Document {
	d := corpus.Document{ID: id, Title: id}
	for _, t := range texts {
		d.Chunks = append(d.Chunks, corpus.Chunk{Section: "s", Text: t})
	}
	return d
}

func sources(wiki, guides []corpus.Document) []corpus.Source {
	return []corpus.Source{
		{Dir: "musicwiki", Prefix: corpus.WikiPrefix, Docs: wiki},
		{Dir: "guides", Prefix: corpus.GuidePrefix, Docs: guides},
	}
}

func TestIngestIsIdempotent(t *testing.T) {
	st, emb := newFakeStore(), &fakeEmbedder{}
	src := sources([]corpus.Document{doc("wiki:/a", "x", "y")}, []corpus.Document{doc("guide:g", "z")})

	first, err := ingest(context.Background(), st, emb, src)
	if err != nil || first.Embedded != 3 || first.Chunks != 3 {
		t.Fatalf("first run = %+v, %v; want 3 embedded", first, err)
	}
	second, err := ingest(context.Background(), st, emb, src)
	if err != nil || second.Embedded != 0 || second.Chunks != 3 || second.Deleted != 0 {
		t.Fatalf("second run = %+v, %v; want 0 embedded, 3 chunks, 0 deleted", second, err)
	}
}

func TestIngestReembedsOnlyChangedChunks(t *testing.T) {
	st, emb := newFakeStore(), &fakeEmbedder{}
	_, _ = ingest(context.Background(), st, emb, sources([]corpus.Document{doc("wiki:/a", "x", "y")}, []corpus.Document{doc("guide:g", "z")}))

	got, err := ingest(context.Background(), st, emb, sources([]corpus.Document{doc("wiki:/a", "x", "y2")}, []corpus.Document{doc("guide:g", "z")}))
	if err != nil || got.Embedded != 1 {
		t.Fatalf("= %+v, %v; want 1 embedded", got, err)
	}
	if ids := st.docs["wiki:/a"]; len(ids) != 2 || ids[store.ChunkID("wiki:/a", "s", "y")] {
		t.Fatalf("wiki:/a chunks = %v; want x and y2 only", ids)
	}
}

func TestIngestReconcileIsScopedPerSource(t *testing.T) {
	st, emb := newFakeStore(), &fakeEmbedder{}
	_, _ = ingest(context.Background(), st, emb, sources(
		[]corpus.Document{doc("wiki:/a", "x")},
		[]corpus.Document{doc("guide:g1", "p"), doc("guide:g2", "q")},
	))

	got, err := ingest(context.Background(), st, emb, sources(
		[]corpus.Document{doc("wiki:/a", "x")},
		[]corpus.Document{doc("guide:g1", "p")},
	))
	if err != nil || got.Deleted != 1 {
		t.Fatalf("= %+v, %v; want 1 deleted", got, err)
	}
	if _, ok := st.docs["guide:g2"]; ok {
		t.Error("guide:g2 still present")
	}
	if _, ok := st.docs["wiki:/a"]; !ok {
		t.Error("wiki:/a deleted by guides reconcile")
	}
}

func TestIngestBatchesAcrossDocuments(t *testing.T) {
	var docs []corpus.Document
	for i := range 130 {
		docs = append(docs, doc(fmt.Sprintf("wiki:/%d", i), fmt.Sprintf("t%d", i)))
	}
	emb := &fakeEmbedder{}
	if _, err := ingest(context.Background(), newFakeStore(), emb, []corpus.Source{{Prefix: corpus.WikiPrefix, Docs: docs}}); err != nil {
		t.Fatal(err)
	}
	if fmt.Sprint(emb.batches) != "[64 64 2]" {
		t.Errorf("batches = %v, want [64 64 2]", emb.batches)
	}
}

func TestIngestEmbedFailureWritesNothing(t *testing.T) {
	st := newFakeStore()
	st.docs["guide:old"] = map[string]bool{"id": true}
	emb := &fakeEmbedder{err: errors.New("429 rate limited")}

	_, err := ingest(context.Background(), st, emb, sources(nil, []corpus.Document{doc("guide:new", "z")}))
	if err == nil || !strings.Contains(err.Error(), "429") {
		t.Fatalf("err = %v, want embed error", err)
	}
	if st.syncs != 0 {
		t.Errorf("SyncDocument called %d times after embed failure", st.syncs)
	}
	if _, ok := st.docs["guide:old"]; !ok {
		t.Error("guide:old deleted after embed failure")
	}
}

func TestIngestRealCorpus(t *testing.T) {
	sources, err := corpus.Load([]string{"../../musicwiki", "../../guides"})
	if err != nil {
		t.Fatal(err)
	}
	st, emb := newFakeStore(), &fakeEmbedder{}

	first, err := ingest(context.Background(), st, emb, sources)
	if err != nil {
		t.Fatal(err)
	}
	if first.Documents != 417 || first.Embedded != first.Chunks {
		t.Fatalf("first run = %+v; want 417 documents, embedded == chunks", first)
	}
	t.Logf("first run: %+v", first)

	second, err := ingest(context.Background(), st, emb, sources)
	if err != nil {
		t.Fatal(err)
	}
	if second.Embedded != 0 || second.Deleted != 0 || second.Chunks != first.Chunks {
		t.Fatalf("second run = %+v; want 0 embedded, 0 deleted, %d chunks", second, first.Chunks)
	}
}
