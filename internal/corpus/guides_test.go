package corpus

import (
	"strings"
	"testing"
)

func TestReadGuides(t *testing.T) {
	dir := writeTree(t, map[string]string{
		"modes.md":  "---\ntitle: Modes\nrelated: [/scales/d-dorian-scale]\n---\n\n# Modes\n\nIntro.\n\n## Dorian\nNatural 6.\n",
		"README.md": "# not a guide",
	})
	docs, err := ReadGuides(dir, map[string]bool{"/scales/d-dorian-scale": true})
	if err != nil {
		t.Fatalf("ReadGuides: %v", err)
	}
	if len(docs) != 1 || docs[0].ID != "guide:modes" || docs[0].Title != "Modes" {
		t.Fatalf("docs = %+v", docs)
	}
	want := []Chunk{
		{Section: "Introduction", Text: "Modes — Introduction\nIntro."},
		{Section: "Dorian", Text: "Modes — Dorian\nNatural 6."},
	}
	for i, c := range docs[0].Chunks {
		if c != want[i] {
			t.Errorf("chunk %d = %q, want %q", i, c, want[i])
		}
	}
}

func TestReadGuidesErrors(t *testing.T) {
	tests := []struct {
		name, content, wantErr string
	}{
		{"missing title", "---\ndescription: d\n---\n\n## A\nB.\n", `missing required field "title"`},
		{"unknown related slug", "---\ntitle: T\nrelated: [/scales/typo]\n---\n\n## A\nB.\n", `related slug "/scales/typo" is not a musicwiki node`},
		{"no content", "---\ntitle: T\n---\n\n# T\n", "has no content"},
		{"no frontmatter", "# T\n\n## A\nB.\n", "missing frontmatter"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := writeTree(t, map[string]string{"g.md": tt.content})
			_, err := ReadGuides(dir, map[string]bool{})
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) || !strings.Contains(err.Error(), "g.md") {
				t.Fatalf("err = %v, want containing %q and the file name", err, tt.wantErr)
			}
		})
	}
}

// The committed guides must load against the committed bundle.
func TestRealGuidesLoad(t *testing.T) {
	sources, err := Load([]string{"../../musicwiki", "../../guides"})
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if got := len(sources[1].Docs); got != 8 {
		t.Errorf("got %d guides, want 8", got)
	}
}
