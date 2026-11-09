package corpus

import (
	"strings"
	"testing"
)

var testTitles = map[string]string{
	"/scales/a-natural-minor-scale":    "A Natural Minor Scale",
	"/scales/c-major-pentatonic-scale": "C Major Pentatonic Scale",
	"/scales/a-blues-scale":            "A Blues Scale",
	"/triads/c-major-triad":            "C Major Triad",
	"/chords/c-major-7th-chord":        "C Major 7th Chord",
}

func TestFlattenScaleOverviewAndPattern(t *testing.T) {
	n := node{
		Type: "scale", Slug: "/scales/a-minor-pentatonic-scale", Title: "A Minor Pentatonic Scale",
		Description:   "The minor pentatonic scale built on A.",
		RelativeMajor: "/scales/c-major-pentatonic-scale",
		Prerequisites: []string{"/scales/a-natural-minor-scale"},
		UsedBy:        []string{"/scales/a-blues-scale"},
		Theory: &theory{
			Intervals: []string{"1", "b3", "4", "5", "b7"},
			Semitones: []int{0, 3, 5, 7, 10},
			Spelling:  []string{"A", "C", "D", "E", "G"},
		},
		Body: "\n# A Minor Pentatonic Scale\n\nThe [A minor pentatonic](/scales/a-minor-pentatonic-scale.md) is spelled **A – C – D – E – G**.\n\n## Related\n\n- **Build on:** [A Natural Minor Scale](/scales/a-natural-minor-scale.md)\n",
		MusicalForms: []form{
			{Form: "pitch_set"},
			{Form: "pattern", System: "2-note-per-string", StartFret: 5, EndFret: 8, Notes: []frettedNote{
				{String: 1, Fret: 8, Note: "C", Interval: "b3"},
				{String: 1, Fret: 5, Note: "A", Interval: "1", IsRoot: true},
				{String: 6, Fret: 8, Note: "C", Interval: "b3"},
				{String: 6, Fret: 5, Note: "A", Interval: "1", IsRoot: true},
			}},
		},
	}

	got, err := flatten(n, testTitles)
	if err != nil {
		t.Fatalf("flatten: %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("got %d chunks, want overview + 1 pattern", len(got))
	}

	wantOverview := strings.Join([]string{
		"A Minor Pentatonic Scale",
		"The minor pentatonic scale built on A.",
		"Notes: A, C, D, E, G.",
		"Intervals: 1 (0 semitones), b3 (3 semitones), 4 (5 semitones), 5 (7 semitones), b7 (10 semitones).",
		"",
		"The A minor pentatonic is spelled **A – C – D – E – G**.",
		"",
		"Relative major: C Major Pentatonic Scale.",
		"Builds on: A Natural Minor Scale.",
		"Used by: A Blues Scale.",
	}, "\n")
	if got[0].Section != "overview" || got[0].Text != wantOverview {
		t.Errorf("overview =\n%s\n--- want ---\n%s", got[0].Text, wantOverview)
	}

	wantPattern := strings.Join([]string{
		"A Minor Pentatonic Scale — Pattern 1 (2-note-per-string, frets 5–8)",
		"String 6 (low E): fret 5 A (1, root), fret 8 C (b3).",
		"String 1 (high E): fret 5 A (1, root), fret 8 C (b3).",
	}, "\n")
	if got[1].Section != "pattern 1 (frets 5–8)" || got[1].Text != wantPattern {
		t.Errorf("pattern = %q / %q\nwant %q", got[1].Section, got[1].Text, wantPattern)
	}
}

func TestFlattenChordVoicing(t *testing.T) {
	n := node{
		Type: "Concept", Slug: "/chords/c-major-7th-chord", Title: "C Major 7th Chord",
		Theory: &theory{Spelling: []string{"C", "E", "G", "B"}},
		Body:   "\n# C Major 7th Chord\n\n**Playing it:** the primary shape is `x-3-2-0-0-0`, around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.\n",
		Render: &render{PreferredPositions: []position{{
			StartFret: 2, EndFret: 3, Tags: []string{"computed", "open"},
			Notes: []frettedNote{
				{String: 5, Fret: 3, Note: "C", Interval: "1", Role: "root", Finger: 2},
				{String: 4, Fret: 2, Note: "E", Interval: "3", Role: "third", Finger: 1},
				{String: 3, Fret: 0, Note: "G", Interval: "5", Role: "fifth"},
				{String: 2, Fret: 0, Note: "B", Interval: "7", Role: "seventh"},
				{String: 1, Fret: 0, Note: "E", Interval: "3", Role: "third"},
			},
		}}},
	}

	got, err := flatten(n, testTitles)
	if err != nil {
		t.Fatalf("flatten: %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("got %d chunks, want overview + 1 voicing", len(got))
	}
	if strings.Contains(got[0].Text, "render block") {
		t.Errorf("overview kept generator note: %q", got[0].Text)
	}
	if !strings.Contains(got[0].Text, "the primary shape is `x-3-2-0-0-0`, around fret 2.") {
		t.Errorf("overview lost playing note: %q", got[0].Text)
	}
	wantVoicing := strings.Join([]string{
		"C Major 7th Chord — Voicing 1: x-3-2-0-0-0 (frets 2–3; open)",
		"String 6 (low E): muted.",
		"String 5 (A): fret 3 C (1, root; finger 2).",
		"String 4 (D): fret 2 E (3, third; finger 1).",
		"String 3 (G): open G (5, fifth).",
		"String 2 (B): open B (7, seventh).",
		"String 1 (high E): open E (3, third).",
	}, "\n")
	if got[1].Section != "voicing 1 (x-3-2-0-0-0)" || got[1].Text != wantVoicing {
		t.Errorf("voicing = %q /\n%s\nwant\n%s", got[1].Section, got[1].Text, wantVoicing)
	}
}

func TestFlattenKeyProgression(t *testing.T) {
	n := node{
		Type: "key", Slug: "/keys/c-major", Title: "C Major",
		MusicalForms: []form{{Form: "progression", Events: []degreeEvent{
			{Roman: "I", Root: "C", Quality: "major", Function: "tonic", Triad: "/triads/c-major-triad", Seventh: "/chords/c-major-7th-chord"},
		}}},
	}
	got, err := flatten(n, testTitles)
	if err != nil {
		t.Fatalf("flatten: %v", err)
	}
	want := "C Major — Diatonic chords\nI: C major (C Major Triad; seventh: C Major 7th Chord) — tonic."
	if len(got) != 2 || got[1].Section != "progression" || got[1].Text != want {
		t.Fatalf("got %q, want progression chunk %q", got, want)
	}
}

func TestFlattenIntervalIsOverviewOnly(t *testing.T) {
	n := node{
		Type: "interval", Slug: "/intervals/minor-third", Title: "Minor Third",
		Description: "The interval of 3 semitones (m3).",
		Theory:      &theory{Intervals: []string{"1", "m3"}, Semitones: []int{0, 3}, Spelling: []string{"C", "Eb"}},
		Body:        "\n# Minor Third\n\nInverting it gives the [Major Sixth](/intervals/major-sixth.md).\n",
	}
	got, err := flatten(n, testTitles)
	if err != nil {
		t.Fatalf("flatten: %v", err)
	}
	want := "Minor Third\nThe interval of 3 semitones (m3).\nNotes: C, Eb.\nIntervals: 1 (0 semitones), m3 (3 semitones).\n\nInverting it gives the Major Sixth."
	if len(got) != 1 || got[0].Text != want {
		t.Fatalf("got %q\nwant %q", got, want)
	}
}

func TestFlattenUnresolvedSlug(t *testing.T) {
	n := node{Type: "scale", Slug: "/s", Title: "S", Theory: &theory{}, UsedBy: []string{"/scales/nope"}}
	_, err := flatten(n, testTitles)
	if err == nil || !strings.Contains(err.Error(), `unknown slug "/scales/nope"`) {
		t.Fatalf("err = %v, want unknown slug", err)
	}
}
