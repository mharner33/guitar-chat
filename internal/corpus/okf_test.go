package corpus

import (
	"strings"
	"testing"
)

const validScale = `---
type: scale
slug: /scales/x-scale
title: X Scale
theory:
  intervals: ['1', b3]
  semitones: [0, 3]
  spelling: [A, C]
musical_forms:
- form: pattern
  system: 2-note-per-string
  start_fret: 5
  end_fret: 8
  notes:
  - {string: 6, fret: 5, interval: '1', note: A, is_root: true}
---

# X Scale

Body.
`

func TestParseNodeValid(t *testing.T) {
	n, err := parseNode("x.md", []byte(validScale))
	if err != nil {
		t.Fatalf("parseNode: %v", err)
	}
	if n.Slug != "/scales/x-scale" || n.Title != "X Scale" || n.Type != "scale" {
		t.Errorf("got slug=%q title=%q type=%q", n.Slug, n.Title, n.Type)
	}
	if got := n.MusicalForms[0].Notes[0]; !got.IsRoot || got.String != 6 || got.Fret != 5 {
		t.Errorf("pattern note = %+v", got)
	}
	if !strings.Contains(n.Body, "Body.") {
		t.Errorf("body = %q", n.Body)
	}
}

func TestParseNodeCRLF(t *testing.T) {
	raw := strings.ReplaceAll(validScale, "\n", "\r\n")
	if _, err := parseNode("x.md", []byte(raw)); err != nil {
		t.Fatalf("parseNode with CRLF: %v", err)
	}
}

func TestParseNodeErrors(t *testing.T) {
	tests := []struct {
		name    string
		raw     string
		wantErr string
	}{
		{"no frontmatter", "# Title\n", "missing frontmatter"},
		{"unterminated", "---\ntype: scale\n", "unterminated frontmatter"},
		{"bad yaml", "---\ntype: [scale\n---\n", "parse frontmatter"},
		{"missing slug", "---\ntype: scale\ntitle: T\ntheory: {}\n---\n", `missing required field "slug"`},
		{"missing title", "---\ntype: scale\nslug: /s\ntheory: {}\n---\n", `missing required field "title"`},
		{"missing type", "---\nslug: /s\ntitle: T\n---\n", `missing required field "type"`},
		{"unknown type", "---\ntype: mode\nslug: /s\ntitle: T\ntheory: {}\n---\n", `unknown node type "mode"`},
		{"missing theory", "---\ntype: scale\nslug: /s\ntitle: T\n---\n", `missing required field "theory"`},
		{"string out of range", "---\ntype: scale\nslug: /s\ntitle: T\ntheory: {}\nmusical_forms:\n- form: pattern\n  notes:\n  - {string: 7, fret: 1}\n---\n", "string 7 out of range"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := parseNode("bad.md", []byte(tt.raw))
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("err = %v, want containing %q", err, tt.wantErr)
			}
			if !strings.Contains(err.Error(), "bad.md") {
				t.Errorf("err %q does not name the file", err)
			}
		})
	}
}

func TestParseNodeKeyNeedsNoTheory(t *testing.T) {
	raw := "---\ntype: key\nslug: /keys/c-major\ntitle: C Major\n---\n\n# C Major\n"
	if _, err := parseNode("k.md", []byte(raw)); err != nil {
		t.Fatalf("key without theory: %v", err)
	}
}
