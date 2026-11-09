package corpus

import (
	"errors"
	"fmt"
	"strings"

	"go.yaml.in/yaml/v3"
)

// node is the subset of an OKF concept file that ingest uses. Unknown fields
// (playback, instrument, tuning, …) are ignored.
type node struct {
	Type            string   `yaml:"type"`
	Slug            string   `yaml:"slug"`
	Title           string   `yaml:"title"`
	Description     string   `yaml:"description"`
	RelativeMajor   string   `yaml:"relative_major"`
	RelativeMinor   string   `yaml:"relative_minor"`
	Prerequisites   []string `yaml:"prerequisites"`
	UsedBy          []string `yaml:"used_by"`
	DegreeFunctions []string `yaml:"degree_functions"`
	Theory          *theory  `yaml:"theory"`
	Render          *render  `yaml:"render"`
	MusicalForms    []form   `yaml:"musical_forms"`

	Body string `yaml:"-"`
	Path string `yaml:"-"`
}

type theory struct {
	Intervals []string `yaml:"intervals"`
	Semitones []int    `yaml:"semitones"`
	Spelling  []string `yaml:"spelling"`
}

type render struct {
	PreferredPositions []position `yaml:"preferred_positions"`
}

// position is one ranked chord voicing.
type position struct {
	StartFret int           `yaml:"start_fret"`
	EndFret   int           `yaml:"end_fret"`
	Notes     []frettedNote `yaml:"notes"`
	Tags      []string      `yaml:"tags"`
}

// frettedNote is one note on the neck. String 1 = high E … 6 = low E.
type frettedNote struct {
	String   int    `yaml:"string"`
	Fret     int    `yaml:"fret"`
	Note     string `yaml:"note"`
	Interval string `yaml:"interval"`
	Role     string `yaml:"role"`
	Finger   int    `yaml:"finger"`
	IsRoot   bool   `yaml:"is_root"`
}

// form is one musical_forms entry; which fields are set depends on Form.
type form struct {
	Form      string        `yaml:"form"`
	System    string        `yaml:"system"`
	StartFret int           `yaml:"start_fret"`
	EndFret   int           `yaml:"end_fret"`
	Notes     []frettedNote `yaml:"notes"`
	Events    []degreeEvent `yaml:"events"`
}

// degreeEvent is one diatonic chord in a key's progression form.
type degreeEvent struct {
	Roman    string `yaml:"roman"`
	Root     string `yaml:"root"`
	Quality  string `yaml:"quality"`
	Function string `yaml:"function"`
	Triad    string `yaml:"triad"`
	Seventh  string `yaml:"seventh"`
}

var knownTypes = map[string]bool{"interval": true, "scale": true, "Concept": true, "key": true}

// splitFrontmatter separates a "---\n…\n---\n" YAML block from the body.
func splitFrontmatter(raw []byte) ([]byte, string, error) {
	s := strings.ReplaceAll(string(raw), "\r\n", "\n")
	if !strings.HasPrefix(s, "---\n") {
		return nil, "", errors.New("missing frontmatter")
	}
	rest := s[len("---\n"):]
	end := strings.Index(rest, "\n---\n")
	if end < 0 {
		return nil, "", errors.New("unterminated frontmatter")
	}
	return []byte(rest[:end]), rest[end+len("\n---\n"):], nil
}

// parseNode parses and validates one OKF concept file. Errors name the file.
func parseNode(path string, raw []byte) (node, error) {
	front, body, err := splitFrontmatter(raw)
	if err != nil {
		return node{}, fmt.Errorf("%s: %w", path, err)
	}
	var n node
	if err := yaml.Unmarshal(front, &n); err != nil {
		return node{}, fmt.Errorf("%s: parse frontmatter: %w", path, err)
	}
	n.Body, n.Path = body, path
	if err := validateNode(n); err != nil {
		return node{}, fmt.Errorf("%s: %w", path, err)
	}
	return n, nil
}

func validateNode(n node) error {
	for _, f := range []struct{ name, val string }{
		{"type", n.Type}, {"slug", n.Slug}, {"title", n.Title},
	} {
		if strings.TrimSpace(f.val) == "" {
			return fmt.Errorf("missing required field %q", f.name)
		}
	}
	if !knownTypes[n.Type] {
		return fmt.Errorf("unknown node type %q", n.Type)
	}
	if n.Type != "key" && n.Theory == nil {
		return fmt.Errorf("missing required field %q", "theory")
	}
	var notes []frettedNote
	for _, f := range n.MusicalForms {
		notes = append(notes, f.Notes...)
	}
	if n.Render != nil {
		for _, p := range n.Render.PreferredPositions {
			notes = append(notes, p.Notes...)
		}
	}
	for _, fn := range notes {
		if fn.String < 1 || fn.String > 6 {
			return fmt.Errorf("string %d out of range 1-6", fn.String)
		}
	}
	return nil
}
