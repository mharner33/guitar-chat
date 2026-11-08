package corpus

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Standard tuning names, keyed by OKF string number (1 = high E … 6 = low E).
var stringNames = map[int]string{6: "low E", 5: "A", 4: "D", 3: "G", 2: "B", 1: "high E"}

var (
	linkRE       = regexp.MustCompile(`\[([^\]]+)\]\([^)]*\)`)
	renderNoteRE = regexp.MustCompile(`\s*\d+ computed fingerings are in the render block[^.]*\.`)
	blankRunRE   = regexp.MustCompile(`\n{3,}`)
)

// flatten turns one node into prose chunks: an overview, plus one chunk per
// scale pattern, chord voicing, and key progression. titles resolves every
// slug in the bundle; an unresolved slug is an error.
func flatten(n node, titles map[string]string) ([]Chunk, error) {
	overview, err := overviewChunk(n, titles)
	if err != nil {
		return nil, err
	}
	chunks := []Chunk{overview}
	patterns, progressions := 0, 0
	for _, f := range n.MusicalForms {
		switch f.Form {
		case "pattern":
			patterns++
			chunks = append(chunks, patternChunk(n.Title, patterns, f))
		case "progression":
			progressions++
			c, err := progressionChunk(n.Title, progressions, f, titles)
			if err != nil {
				return nil, err
			}
			chunks = append(chunks, c)
		}
	}
	if n.Render != nil {
		for i, p := range n.Render.PreferredPositions {
			chunks = append(chunks, voicingChunk(n.Title, i+1, p))
		}
	}
	return chunks, nil
}

func overviewChunk(n node, titles map[string]string) (Chunk, error) {
	lines := []string{n.Title}
	if n.Description != "" {
		lines = append(lines, n.Description)
	}
	if t := n.Theory; t != nil {
		if len(t.Spelling) > 0 {
			lines = append(lines, "Notes: "+strings.Join(t.Spelling, ", ")+".")
		}
		if len(t.Intervals) > 0 && len(t.Intervals) == len(t.Semitones) {
			parts := make([]string, len(t.Intervals))
			for i, iv := range t.Intervals {
				parts[i] = fmt.Sprintf("%s (%d semitones)", iv, t.Semitones[i])
			}
			lines = append(lines, "Intervals: "+strings.Join(parts, ", ")+".")
		}
	}
	if len(n.DegreeFunctions) > 0 {
		lines = append(lines, "Degree functions: "+strings.Join(n.DegreeFunctions, ", ")+".")
	}
	if body := cleanBody(n.Body); body != "" {
		lines = append(lines, "", body)
	}

	rels := []struct {
		label string
		slugs []string
	}{
		{"Relative major", nonEmpty(n.RelativeMajor)},
		{"Relative minor", nonEmpty(n.RelativeMinor)},
		{"Builds on", n.Prerequisites},
		{"Used by", n.UsedBy},
	}
	var relLines []string
	for _, r := range rels {
		if len(r.slugs) == 0 {
			continue
		}
		names, err := resolve(r.slugs, titles)
		if err != nil {
			return Chunk{}, fmt.Errorf("%s: %w", r.label, err)
		}
		relLines = append(relLines, fmt.Sprintf("%s: %s.", r.label, strings.Join(names, ", ")))
	}
	if len(relLines) > 0 {
		lines = append(lines, "")
		lines = append(lines, relLines...)
	}
	return Chunk{Section: "overview", Text: strings.Join(lines, "\n")}, nil
}

// cleanBody drops the title line and the "## Related" section (relationships
// come from frontmatter instead), reduces links to their text, and removes the
// generator's note about the render block.
func cleanBody(body string) string {
	var kept []string
	skipping := false
	for _, line := range strings.Split(strings.ReplaceAll(body, "\r\n", "\n"), "\n") {
		if strings.HasPrefix(line, "## ") {
			skipping = strings.TrimSpace(strings.TrimPrefix(line, "## ")) == "Related"
		}
		if skipping || strings.HasPrefix(line, "# ") {
			continue
		}
		kept = append(kept, line)
	}
	out := strings.Join(kept, "\n")
	out = linkRE.ReplaceAllString(out, "$1")
	out = renderNoteRE.ReplaceAllString(out, "")
	out = blankRunRE.ReplaceAllString(out, "\n\n")
	return strings.TrimSpace(out)
}

func patternChunk(title string, i int, f form) Chunk {
	span := fretSpan(f.StartFret, f.EndFret)
	head := fmt.Sprintf("%s — Pattern %d (%s, %s)", title, i, f.System, span)
	lines := append([]string{head}, stringLines(f.Notes, false)...)
	return Chunk{Section: fmt.Sprintf("pattern %d (%s)", i, span), Text: strings.Join(lines, "\n")}
}

func voicingChunk(title string, i int, p position) Chunk {
	t := tab(p.Notes)
	head := fmt.Sprintf("%s — Voicing %d: %s (%s", title, i, t, fretSpan(p.StartFret, p.EndFret))
	var tags []string
	for _, tag := range p.Tags {
		if tag != "computed" {
			tags = append(tags, tag)
		}
	}
	if len(tags) > 0 {
		head += "; " + strings.Join(tags, ", ")
	}
	head += ")"
	lines := append([]string{head}, stringLines(p.Notes, true)...)
	return Chunk{Section: fmt.Sprintf("voicing %d (%s)", i, t), Text: strings.Join(lines, "\n")}
}

func progressionChunk(title string, i int, f form, titles map[string]string) (Chunk, error) {
	lines := []string{title + " — Diatonic chords"}
	for _, e := range f.Events {
		names, err := resolve([]string{e.Triad, e.Seventh}, titles)
		if err != nil {
			return Chunk{}, fmt.Errorf("progression %s: %w", e.Roman, err)
		}
		lines = append(lines, fmt.Sprintf("%s: %s %s (%s; seventh: %s) — %s.",
			e.Roman, e.Root, e.Quality, names[0], names[1], e.Function))
	}
	section := "progression"
	if i > 1 {
		section = fmt.Sprintf("progression %d", i)
	}
	return Chunk{Section: section, Text: strings.Join(lines, "\n")}, nil
}

// stringLines lists notes per string from 6 (low E) to 1 (high E), frets
// ascending. In a voicing, a string with no note is muted and fingers are given.
func stringLines(notes []frettedNote, voicing bool) []string {
	byString := map[int][]frettedNote{}
	for _, n := range notes {
		byString[n.String] = append(byString[n.String], n)
	}
	var lines []string
	for s := 6; s >= 1; s-- {
		ns := byString[s]
		if len(ns) == 0 {
			if voicing {
				lines = append(lines, fmt.Sprintf("String %d (%s): muted.", s, stringNames[s]))
			}
			continue
		}
		sort.SliceStable(ns, func(i, j int) bool { return ns[i].Fret < ns[j].Fret })
		parts := make([]string, len(ns))
		for i, n := range ns {
			parts[i] = describeNote(n, voicing)
		}
		lines = append(lines, fmt.Sprintf("String %d (%s): %s.", s, stringNames[s], strings.Join(parts, ", ")))
	}
	return lines
}

func describeNote(n frettedNote, withFinger bool) string {
	label := n.Interval
	switch {
	case n.Role != "":
		label += ", " + n.Role
	case n.IsRoot:
		label += ", root"
	}
	if withFinger && n.Fret > 0 && n.Finger > 0 {
		label += fmt.Sprintf("; finger %d", n.Finger)
	}
	where := fmt.Sprintf("fret %d", n.Fret)
	if n.Fret == 0 {
		where = "open"
	}
	return fmt.Sprintf("%s %s (%s)", where, n.Note, label)
}

// tab renders a voicing low E → high E, "x" for strings without a note.
func tab(notes []frettedNote) string {
	frets := map[int]int{}
	for _, n := range notes {
		frets[n.String] = n.Fret
	}
	parts := make([]string, 0, 6)
	for s := 6; s >= 1; s-- {
		if f, ok := frets[s]; ok {
			parts = append(parts, strconv.Itoa(f))
		} else {
			parts = append(parts, "x")
		}
	}
	return strings.Join(parts, "-")
}

func fretSpan(lo, hi int) string {
	if lo == hi {
		return fmt.Sprintf("fret %d", lo)
	}
	return fmt.Sprintf("frets %d–%d", lo, hi)
}

func resolve(slugs []string, titles map[string]string) ([]string, error) {
	out := make([]string, len(slugs))
	for i, s := range slugs {
		t, ok := titles[s]
		if !ok {
			return nil, fmt.Errorf("unknown slug %q", s)
		}
		out[i] = t
	}
	return out, nil
}

func nonEmpty(s string) []string {
	if s == "" {
		return nil
	}
	return []string{s}
}
