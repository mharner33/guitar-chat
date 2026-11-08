package corpus

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

const bundleDir = "../../musicwiki"

func TestReadMusicWikiRealBundle(t *testing.T) {
	docs, err := ReadMusicWiki(bundleDir)
	if err != nil {
		t.Fatalf("ReadMusicWiki: %v", err)
	}
	if len(docs) != 409 {
		t.Errorf("got %d documents, want 409", len(docs))
	}

	kinds := map[string]int{}
	byID := map[string]Document{}
	for _, d := range docs {
		byID[d.ID] = d
		for _, c := range d.Chunks {
			kinds[strings.Fields(c.Section)[0]]++
			if !strings.HasPrefix(c.Text, d.Title) {
				t.Fatalf("%s %q: text does not start with title %q", d.ID, c.Section, d.Title)
			}
		}
	}
	want := map[string]int{"overview": 409, "pattern": 429, "voicing": 684, "progression": 24}
	for k, n := range want {
		if kinds[k] != n {
			t.Errorf("%s chunks = %d, want %d", k, kinds[k], n)
		}
	}

	cmaj7 := byID["wiki:/chords/c-major-7th-chord"]
	if !hasChunk(cmaj7, "voicing 1 (x-3-2-0-0-0)") {
		t.Errorf("C maj7 missing voicing 1 x-3-2-0-0-0; sections: %v", sections(cmaj7))
	}
	gmaj := byID["wiki:/keys/g-major"]
	if c := chunk(gmaj, "progression"); !strings.Contains(c.Text, "vii°: F# diminished") {
		t.Errorf("G major progression = %q", c.Text)
	}
}

// Positions 1–3 of A minor pentatonic must match the hand-checked
// lesson-creator fixtures (lesson-creator/fixtures.py POSITIONS).
func TestMinorPentatonicMatchesLessonFixtures(t *testing.T) {
	fixtures := []map[int][]int{
		{6: {5, 8}, 5: {5, 7}, 4: {5, 7}, 3: {5, 7}, 2: {5, 8}, 1: {5, 8}},
		{6: {8, 10}, 5: {7, 10}, 4: {7, 10}, 3: {7, 9}, 2: {8, 10}, 1: {8, 10}},
		{6: {10, 12}, 5: {10, 12}, 4: {10, 12}, 3: {9, 12}, 2: {10, 13}, 1: {10, 12}},
	}
	docs, err := ReadMusicWiki(bundleDir)
	if err != nil {
		t.Fatalf("ReadMusicWiki: %v", err)
	}
	var doc Document
	for _, d := range docs {
		if d.ID == "wiki:/scales/a-minor-pentatonic-scale" {
			doc = d
		}
	}
	lineRE := regexp.MustCompile(`^String (\d) \([^)]*\): (.*)\.$`)
	fretRE := regexp.MustCompile(`fret (\d+)`)
	var patterns []Chunk
	for _, c := range doc.Chunks {
		if strings.HasPrefix(c.Section, "pattern ") {
			patterns = append(patterns, c)
		}
	}
	if len(patterns) != len(fixtures) {
		t.Fatalf("got %d patterns, want %d", len(patterns), len(fixtures))
	}
	for i, c := range patterns {
		got := map[int][]int{}
		for _, line := range strings.Split(c.Text, "\n")[1:] {
			m := lineRE.FindStringSubmatch(line)
			if m == nil {
				t.Fatalf("pattern %d: unparseable line %q", i+1, line)
			}
			s, _ := strconv.Atoi(m[1])
			for _, fm := range fretRE.FindAllStringSubmatch(m[2], -1) {
				f, _ := strconv.Atoi(fm[1])
				got[s] = append(got[s], f)
			}
		}
		for s, frets := range fixtures[i] {
			if len(got[s]) != len(frets) || got[s][0] != frets[0] || got[s][1] != frets[1] {
				t.Errorf("pattern %d string %d: frets %v, fixture %v", i+1, s, got[s], frets)
			}
		}
	}
}

func TestReadMusicWikiErrors(t *testing.T) {
	tests := []struct {
		name    string
		files   map[string]string
		wantErr string
	}{
		{
			name: "duplicate slug",
			files: map[string]string{
				"intervals/a.md": "---\ntype: interval\nslug: /intervals/x\ntitle: A\ntheory: {}\n---\n",
				"intervals/b.md": "---\ntype: interval\nslug: /intervals/x\ntitle: B\ntheory: {}\n---\n",
			},
			wantErr: `duplicate slug "/intervals/x"`,
		},
		{
			name: "unresolved relationship",
			files: map[string]string{
				"scales/s.md": "---\ntype: scale\nslug: /scales/s\ntitle: S\ntheory: {}\nused_by: [/scales/missing]\n---\n",
			},
			wantErr: `unknown slug "/scales/missing"`,
		},
		{
			name: "bad node names file",
			files: map[string]string{
				"chords/bad.md": "---\ntype: Concept\ntitle: Bad\ntheory: {}\n---\n",
			},
			wantErr: "bad.md",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := writeTree(t, tt.files)
			_, err := ReadMusicWiki(dir)
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("err = %v, want containing %q", err, tt.wantErr)
			}
		})
	}
}

func TestReadMusicWikiSkipsIndexAndRootFiles(t *testing.T) {
	dir := writeTree(t, map[string]string{
		"index.md":           "---\ntype: Index\n---\n",
		"log.md":             "---\ntype: Log\n---\n",
		"README.md":          "# readme",
		"intervals/index.md": "---\ntype: Index\n---\n",
		"intervals/u.md":     "---\ntype: interval\nslug: /intervals/u\ntitle: U\ntheory: {}\n---\n",
	})
	docs, err := ReadMusicWiki(dir)
	if err != nil {
		t.Fatalf("ReadMusicWiki: %v", err)
	}
	if len(docs) != 1 || docs[0].ID != "wiki:/intervals/u" {
		t.Fatalf("docs = %+v, want only wiki:/intervals/u", docs)
	}
}

func writeTree(t *testing.T, files map[string]string) string {
	t.Helper()
	dir := t.TempDir()
	for name, content := range files {
		p := filepath.Join(dir, name)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return dir
}

func chunk(d Document, section string) Chunk {
	for _, c := range d.Chunks {
		if c.Section == section {
			return c
		}
	}
	return Chunk{}
}

func hasChunk(d Document, section string) bool { return chunk(d, section).Section != "" }

func sections(d Document) []string {
	out := make([]string, len(d.Chunks))
	for i, c := range d.Chunks {
		out[i] = c.Section
	}
	return out
}
