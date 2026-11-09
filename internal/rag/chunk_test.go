package rag

import (
	"strings"
	"testing"
)

func TestChunkMarkdown(t *testing.T) {
	tests := []struct {
		name string
		body string
		want []Chunk
	}{
		{
			name: "intro and sections, H1 dropped",
			body: "# Title\n\nIntro text.\n\n## First\n\nAlpha.\n\n## Second\nBeta.\n",
			want: []Chunk{
				{Section: "Introduction", Text: "Intro text."},
				{Section: "First", Text: "Alpha."},
				{Section: "Second", Text: "Beta."},
			},
		},
		{
			name: "empty sections skipped",
			body: "## Empty\n\n\n## Full\nGamma.",
			want: []Chunk{{Section: "Full", Text: "Gamma."}},
		},
		{
			name: "CRLF line endings",
			body: "## One\r\nDelta.\r\n\r\n## Two\r\nEpsilon.\r\n",
			want: []Chunk{{Section: "One", Text: "Delta."}, {Section: "Two", Text: "Epsilon."}},
		},
		{
			name: "### subheadings stay inside their section",
			body: "## Top\nA.\n### Sub\nB.",
			want: []Chunk{{Section: "Top", Text: "A.\n### Sub\nB."}},
		},
		{
			name: "no content",
			body: "# Only a title\n\n",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ChunkMarkdown(tt.body)
			if len(got) != len(tt.want) {
				t.Fatalf("got %d chunks %q, want %d %q", len(got), got, len(tt.want), tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("chunk %d = %q, want %q", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestChunkMarkdownSplitsLongSectionWithOverlap(t *testing.T) {
	a := strings.Repeat("a", 1500)
	b := strings.Repeat("b", 1500)
	c := strings.Repeat("c", 150) // small enough to be the overlap
	d := strings.Repeat("d", 1500)
	body := "## Long\n" + strings.Join([]string{a, b, c, d}, "\n\n")

	got := ChunkMarkdown(body)

	if len(got) != 2 {
		t.Fatalf("got %d chunks, want 2", len(got))
	}
	if got[0].Section != "Long" || got[0].Text != a+"\n\n"+b+"\n\n"+c {
		t.Errorf("part 1 = section %q, %d chars; want Long with a,b,c", got[0].Section, len(got[0].Text))
	}
	if got[1].Section != "Long (part 2)" || got[1].Text != c+"\n\n"+d {
		t.Errorf("part 2 = section %q; want \"Long (part 2)\" starting with overlap c then d", got[1].Section)
	}
}

func TestChunkMarkdownKeepsOversizedParagraphWhole(t *testing.T) {
	huge := strings.Repeat("z", 5000)
	got := ChunkMarkdown("## Big\n" + huge)
	if len(got) != 1 || got[0].Text != huge {
		t.Fatalf("got %d chunks; want the 5000-char paragraph kept as one chunk", len(got))
	}
}

func TestChunkMarkdownNoOverlapOnlyChunk(t *testing.T) {
	a := strings.Repeat("a", 1500)
	b := strings.Repeat("b", 1500)
	c := strings.Repeat("c", 150)    // becomes the overlap seed for part 2
	big := strings.Repeat("g", 3100) // too big to share a chunk with the 152-char overlap seed
	body := "## Long\n" + strings.Join([]string{a, b, c, big}, "\n\n")

	got := ChunkMarkdown(body)

	if len(got) != 2 {
		t.Fatalf("got %d chunks, want 2 (no overlap-only chunk)", len(got))
	}
	if got[0].Section != "Long" || got[0].Text != a+"\n\n"+b+"\n\n"+c {
		t.Errorf("part 1 = section %q, %d chars; want Long with a,b,c", got[0].Section, len(got[0].Text))
	}
	if got[1].Section != "Long (part 2)" || got[1].Text != c+"\n\n"+big {
		t.Errorf("part 2 = section %q, %d chars; want \"Long (part 2)\" with overlap c then big", got[1].Section, len(got[1].Text))
	}
}
