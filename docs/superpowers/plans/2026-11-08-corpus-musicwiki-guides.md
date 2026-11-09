# Corpus: musicwiki + guides Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Ingest the generated `musicwiki/` bundle (frontmatter flattened to prose) plus eight hand-written `guides/` pages into pgvector, idempotently and with per-source reconcile.

**Architecture:** `internal/corpus` reads each source directory into `[]Document{ID, Title, Chunks}` — an OKF reader that flattens YAML into prose chunks, and a guides reader that uses the heading-aware chunker in `internal/rag`. `cmd/ingest` loads and validates *all* sources before touching the database, embeds only chunks whose content-hash id is new (batches of 64), syncs each document in a transaction, then deletes documents of that source's id prefix that were not seen.

**Tech Stack:** Go 1.26, pgx/v5, Postgres 16 + pgvector, `go.yaml.in/yaml/v3` (new dependency), OpenAI-compatible `/embeddings` REST.

**Spec:** `docs/superpowers/specs/2026-11-08-corpus-musicwiki-guides-design.md`

## Global Constraints

- `musicwiki/` is read-only generated output — never edit files under it.
- Document ids: `wiki:` + node slug (`wiki:/scales/a-minor-pentatonic-scale`); `guide:` + filename stem (`guide:modes-and-their-chords`).
- Chunk id = `sha256(doc_id, section, text)` (fields NUL-separated), hex-encoded.
- `CORPUS_DIRS` default `musicwiki,guides`; reader chosen by directory basename (`musicwiki` → OKF reader, anything else → guides reader).
- Every wiki chunk's text starts with the node title.
- Ingest fails loudly (non-zero exit, nothing written or deleted) on: missing/empty source dir, unparseable YAML, unknown node `type`, missing `slug`/`title`/`type` (and `theory` on non-key nodes), unresolved slug, guide `related` slug not in musicwiki, guide without `title`.
- Embed batches of at most 64; every embed call gets a 60s deadline.
- All corpus content is MIT. No CC BY-SA, no `corpus/` directory, no `ATTRIBUTION.md`.
- No migration — the existing `documents`/`chunks` schema is used as-is.
- Unit tests make no network calls (`httptest` loopback is fine). Store tests skip when `DATABASE_URL` is unset.
- Checks before every commit: `go vet ./...`, `staticcheck ./...`, `go test ./...`, `gofmt -l .` (prints nothing).
- Commit messages end with `Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>`.

## Review Focus

1. **Two directories of the same kind in `CORPUS_DIRS`** (e.g. `guides,more-guides`) would share the `guide:` prefix, and each reconcile would delete the other's documents. Expect: `corpus.Load` rejects the config. → test in Task 8.
2. **Empty keep-list in SQL.** A nil Go slice becomes SQL `NULL`, and `NOT (id = ANY(NULL))` is `NULL`, so nothing is deleted. Expect: an empty keep list deletes every chunk of the doc / every document of the prefix, and only those. → test in Task 10.
3. **Embedding fails mid-ingest** (rate limit, bad key). Expect: no `SyncDocument` or delete runs for that source; the database is unchanged. → test in Task 11.
4. **Windows line endings in a hand-edited guide** (`\r\n`). Expect: frontmatter still parses and headings still split. → test in Task 5 and Task 4.
5. **`CORPUS_DIRS` written as `./musicwiki/` or an absolute path.** Expect: still detected as the wiki source by basename. → test in Task 8.

---

### Task 1: Doc updates and draft golden dataset

Docs-only; no tests. Brings every authoritative doc in line with the approved spec before code lands.

**Files:**
- Modify: `docs/design.md` (corpus-source + licenses paragraphs, §4.1 layout, §4.3 config table, §5.1, §9.5 step 6, §14)
- Modify: `CLAUDE.md` (ingestion paragraph, licensing guardrail)
- Modify: `agent.md` (ingestion paragraph line 69, licensing line 91)
- Modify: `docs/plan.md` (step 1 and step 7)
- Modify: `.env.example`
- Create: `evals/golden.jsonl`

**Interfaces:** none.

- [ ] **Step 1: `docs/design.md` — corpus source and licenses**

Replace the paragraph starting `Corpus source: seed content from openly licensed…` with:

```
Corpus source: two folders. musicwiki/ is a generated Open Knowledge Format bundle (intervals, scales, triads, chords, keys — one Markdown file per concept, structured YAML frontmatter plus a short prose body); it is produced by a separate generator repo and replaced wholesale, never hand-edited here. guides/ holds hand-written explanation pages (why major and minor sound different, modes and their chords, CAGED, chord construction, …) that answer the "why" questions the wiki's reference pages cannot. Ingestion is a separate command so the corpus can evolve without redeploying.
```

Replace the paragraph starting `Licenses: root LICENSE is MIT…` with:

```
Licenses: everything is MIT. The root LICENSE covers the Go code, web/, and guides/; musicwiki/LICENSE covers the generated bundle. Nothing in the repo is adapted from CC BY-SA material. If such material is ever adapted, it goes in its own folder with its own LICENSE and attribution file — never mixed into guides/.
```

- [ ] **Step 2: `docs/design.md` — §4.1 layout**

Replace the `corpus/` block (five lines from `├── corpus/` through `│   └── ...`) with:

```
├── musicwiki/                  # generated OKF bundle (MIT, own LICENSE) — read-only here
│   ├── intervals/ scales/ triads/ chords/ keys/
│   └── index.md log.md …       # not ingested
├── guides/                     # hand-written explanation pages (MIT)
│   ├── modes-and-their-chords.md
│   └── ...
```

and add under `internal/` (after the `rag/` block):

```
│   ├── corpus/                 # source readers: OKF flattener + guides reader
```

- [ ] **Step 3: `docs/design.md` — §4.3 config table**

After the three `EMBEDDING_DIMENSIONS` lines (variable / purpose / `1536`) insert:

```
CORPUS_DIRS
Comma-separated corpus source directories. A directory named musicwiki uses the OKF reader; any other uses the guides reader. At most one of each.
musicwiki,guides
```

- [ ] **Step 4: `docs/design.md` — §5.1 ingestion**

Replace `Read every corpus/*.md file. Skip corpus/LICENSE and corpus/ATTRIBUTION.md.` with:

```
Read every source in CORPUS_DIRS before writing anything. musicwiki: every *.md under intervals/ scales/ triads/ chords/ keys/ except index.md; frontmatter is flattened into prose chunks (overview, one per scale pattern, one per chord voicing, one per key progression), each starting with the node title. guides: every *.md in the folder, chunked by the heading-aware splitter below. Document ids carry a source prefix: wiki:<slug>, guide:<filename stem>. Ingest refuses to run if a source directory is missing or yields no documents, or if any file fails validation.
```

Replace the sentence starting `After the corpus walk, delete documents whose slugs are no longer in corpus/` with:

```
After each source is written, delete documents with that source's id prefix that were not in it, and their chunks. Reconcile never crosses sources: re-ingesting guides cannot delete wiki rows. Re-ingesting a degraded file or removing a file leaves only the current corpus searchable. This is what makes demo step 6 (groundedness drops after a degraded re-ingest) true. Unchanged chunks keep their content-hash ids and are not re-embedded.
```

- [ ] **Step 5: `docs/design.md` — §9.5 and §14**

In §9.5, replace `(6) re-ingest a degraded corpus (old chunks for that file are deleted)` with `(6) re-ingest a degraded guide (e.g. a wrong guides/modes-and-their-chords.md; its old chunks are deleted)`.

In §14, delete the line `Open Music Theory (corpus candidate, CC BY-SA)`.

- [ ] **Step 6: `CLAUDE.md` and `agent.md`**

`CLAUDE.md`, "Ingestion is idempotent and reconciling" paragraph — replace `For each corpus file: … After the walk, delete documents whose slugs are gone from \`corpus/\`.` with:

```
Sources are `musicwiki/` (generated OKF bundle, YAML frontmatter flattened to prose in `internal/corpus`) and `guides/` (hand-written), listed in `CORPUS_DIRS`. Document ids are `wiki:<slug>` / `guide:<stem>`. All sources are loaded and validated before any write. For each document: upsert it, insert new chunks, delete its chunk rows whose hashes weren't produced this run. After each source, delete documents with that source's prefix that weren't seen — reconcile never crosses sources. `musicwiki/` is generated: never hand-edit it.
```

`CLAUDE.md`, Guardrails licensing bullet — replace with:

```
- Licensing: everything is MIT — root `LICENSE` covers code, `web/` and `guides/`; `musicwiki/LICENSE` covers the bundle. Do not add CC BY-SA or other share-alike text to `guides/`; adapted material needs its own folder, LICENSE and attribution.
```

`agent.md` line 69 — apply the same replacement as the CLAUDE.md ingestion paragraph (same wording). `agent.md` line 91 — replace with the same licensing sentence (without the leading `- `).

- [ ] **Step 7: `docs/plan.md` and `.env.example`**

`docs/plan.md` step 1: replace `` `corpus/`, `` with `` `guides/`, ``. Step 7: replace `Seed a few \`corpus/*.md\` pages + \`corpus/LICENSE\` + \`corpus/ATTRIBUTION.md\`.` with `Sources are the existing \`musicwiki/\` bundle and the hand-written \`guides/\` (see docs/superpowers/specs/2026-11-08-corpus-musicwiki-guides-design.md).`

`.env.example` — after the `EMBEDDING_DIMENSIONS=1536` line add:

```

# ── Corpus ───────────────────────────────────────────────────────────────────
# Source directories for cmd/ingest. A directory named musicwiki uses the OKF
# reader; any other uses the guides reader. At most one of each.
CORPUS_DIRS=musicwiki,guides
```

- [ ] **Step 8: Create `evals/golden.jsonl`**

One JSON object per line, exactly:

```jsonl
{"question": "Why does a major scale sound different from a minor scale?", "expected_points": ["the difference starts with the third: a major third is 4 semitones, a minor third is 3", "natural minor also lowers the 6th and 7th degrees", "minor is commonly heard as darker or sadder"], "rubric": "standard"}
{"question": "What chords work over D Dorian?", "expected_points": ["D Dorian uses the same notes as C major", "the diatonic chords are Dm7, Em7, Fmaj7, G7, Am7, Bm7b5, Cmaj7", "the natural 6th (B) is Dorian's characteristic note"], "rubric": "standard"}
{"question": "How is a dominant 7th chord built?", "expected_points": ["a major triad plus a minor seventh", "formula 1-3-5-b7"], "rubric": "standard"}
{"question": "What is the CAGED system?", "expected_points": ["five open-chord shapes: C, A, G, E, D", "the shapes are moved up the neck as movable forms", "together they cover the same chord in five places across the fretboard"], "rubric": "standard"}
{"question": "How do the minor pentatonic positions connect?", "expected_points": ["five overlapping positions (boxes)", "adjacent positions share notes", "A minor pentatonic position 1 sits at frets 5-8"], "rubric": "standard"}
{"question": "What does the V chord do in a key?", "expected_points": ["it has dominant function", "it pulls toward / resolves to the I (tonic) chord"], "rubric": "standard"}
{"question": "What is a ii-V-I in C major?", "expected_points": ["Dm7 - G7 - Cmaj7"], "rubric": "standard"}
{"question": "How does a 12-bar blues go?", "expected_points": ["built on the I, IV and V chords", "a standard 12-bar layout, e.g. I-I-I-I / IV-IV-I-I / V-IV-I-V"], "rubric": "standard"}
{"question": "What notes are in a G dominant 7th chord?", "expected_points": ["G, B, D, F"], "rubric": "standard"}
{"question": "How do I play an open C major 7th chord?", "expected_points": ["tab x-3-2-0-0-0 (low E muted)"], "rubric": "standard"}
{"question": "What notes are in the E blues scale?", "expected_points": ["E, G, A, Bb, B, D", "the b5 (Bb) is the added blue note"], "rubric": "standard"}
{"question": "Where is A minor pentatonic position 1 on the fretboard?", "expected_points": ["frets 5-8", "root A at string 6 (low E), fret 5"], "rubric": "standard"}
{"question": "What are the notes of an A minor 7 flat 5 chord?", "expected_points": ["A, C, Eb, G", "formula 1-b3-b5-b7"], "rubric": "standard"}
{"question": "How many semitones is a tritone?", "expected_points": ["6 semitones", "for example C to F#"], "rubric": "standard"}
{"question": "What is the relative major of A minor pentatonic?", "expected_points": ["C major pentatonic"], "rubric": "standard"}
{"question": "What are the diatonic chords in G major?", "expected_points": ["G, Am, Bm, C, D, Em, F#dim"], "rubric": "standard"}
{"question": "What's a good recipe for banana bread?", "expected_points": ["answers exactly \"I don't have enough context\" with no sources"], "rubric": "off_corpus"}
{"question": "How do I set up a Python virtual environment?", "expected_points": ["answers exactly \"I don't have enough context\" with no sources"], "rubric": "off_corpus"}
```

Validate it parses: `python3 -c "import json; [json.loads(l) for l in open('evals/golden.jsonl')]; print('ok')"` → `ok`.

- [ ] **Step 9: Commit**

```bash
git add docs/design.md CLAUDE.md agent.md docs/plan.md .env.example evals/golden.jsonl
git commit -m "docs: corpus is musicwiki + guides; drop Open Music Theory and CC BY-SA

Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>"
```

---

### Task 2: Draft the eight guides (owner review gate)

Content task. **Ends at an owner checkpoint** — do not start Task 8's real-guides test until the owner has reviewed (Tasks 3–7 may proceed in parallel).

**Files:**
- Create: `guides/why-major-and-minor-sound-different.md`
- Create: `guides/modes-and-their-chords.md`
- Create: `guides/how-chords-are-built.md`
- Create: `guides/reading-the-fretboard.md`
- Create: `guides/the-caged-system.md`
- Create: `guides/pentatonic-positions.md`
- Create: `guides/diatonic-harmony-and-function.md`
- Create: `guides/common-progressions.md`
- Scratch (not committed): checker script + report in the session scratchpad

**Interfaces:**
- Produces: guide files in the format below, consumed by `corpus.ReadGuides` (Task 8).

**Format rules (every guide):**

```markdown
---
title: <Title Case name>
description: <one sentence>
tags: [<lowercase>, ...]
related:
  - /scales/...        # musicwiki slugs only; each must exist
---

# <same title>

<1–2 paragraph introduction>

## <section>
...
```

- 600–1,500 words. Each `##` section under ~3,000 characters (one chunk).
- String numbering: 1 = high E … 6 = low E (matches the wiki). Tab written low → high (`x-3-2-0-1-0`).
- Use only keys/shapes that exist in musicwiki for concrete examples. Spell notes as the wiki does (e.g. `Bb`, `F#`).
- No text copied from any external source. Original prose only (MIT).

- [ ] **Step 1: Write the guides** — required content per guide:

| File | Required sections / facts | `related` (minimum) |
|---|---|---|
| why-major-and-minor-sound-different | The third (M3 = 4 semitones, m3 = 3); C major vs C natural minor spelled side by side (C D E F G A B / C D Eb F G Ab Bb); the b6 and b7; consonance/tension; "sad" is cultural association, not physics | `/scales/c-major-scale`, `/scales/c-natural-minor-scale`, `/intervals/major-third`, `/intervals/minor-third` |
| modes-and-their-chords | Modes as rotations of C major (C Ionian … B Locrian); each mode's formula and characteristic note (Dorian natural 6, Phrygian b2, Lydian #4, Mixolydian b7, Aeolian b6, Locrian b5); the seven diatonic 7th chords shared by C major's modes (Cmaj7 Dm7 Em7 Fmaj7 G7 Am7 Bm7b5); which chord is the "home" chord per mode (D Dorian → Dm7, G Mixolydian → G7, …) | `/scales/d-dorian-scale`, `/scales/g-mixolydian-scale`, `/keys/c-major` |
| how-chords-are-built | Stacking thirds; four triad qualities; 7th chords (maj7, 7, m7, m7b5, dim7); 6th chords; extensions 9/11/13 (C9 = C E G Bb D); sus2/sus4 replace the third | `/triads/c-major-triad`, `/chords/c-dominant-7th-chord`, `/chords/c-9th-chord`, `/chords/c-suspended-4-chord` |
| reading-the-fretboard | String numbering and open-string names; fret = semitone; octave shapes; finding a root on strings 6 and 5; intervals as fretboard shapes (M3 = one string up, one fret back, except across G–B) | `/intervals/perfect-octave`, `/intervals/major-third` |
| the-caged-system | The five open shapes C A G E D; each as a movable barre form; their order up the neck for one chord (C: C-shape open, A-shape 3rd fret, G-shape, E-shape 8th fret, D-shape); connecting shapes; link to scale positions | `/triads/c-major-triad` |
| pentatonic-positions | Minor pentatonic formula 1 b3 4 5 b7; A minor = A C D E G; five positions with fret ranges 5–8, 7–10, 9–13, 12–15, 14–17 and per-string frets (from `lesson-creator/fixtures.py`); shared notes between neighbours; relative major (C major pentatonic, same notes); blues scale adds b5 | `/scales/a-minor-pentatonic-scale`, `/scales/c-major-pentatonic-scale`, `/scales/a-blues-scale` |
| diatonic-harmony-and-function | Harmonizing the major scale; Roman numerals (upper/lower/°); C major's chords I ii iii IV V vi vii°; functions tonic / predominant / dominant as the wiki labels them; V→I resolution and the tritone in V7 (G7: B–F) | `/keys/c-major`, `/keys/a-natural-minor`, `/chords/g-dominant-7th-chord` |
| common-progressions | I–IV–V (G: G C D); ii–V–I (C: Dm7 G7 Cmaj7); I–V–vi–IV (C: C G Am F); 12-bar blues in A with the bar layout I I I I / IV IV I I / V IV I V (A7 D7 E7); minor i–iv–v | `/keys/g-major`, `/keys/c-major`, `/chords/a-dominant-7th-chord` |

- [ ] **Step 2: Write the throwaway checker** (scratchpad, Python, run with `python3 -I`)

The checker must:
1. Load every musicwiki node's frontmatter (`yaml.safe_load` on the text between the first `---\n` and `\n---\n`), building `slug → node`.
2. For each guide: verify every `related` slug exists; flag any `##` section over 3,000 chars; count words (600–1,500).
3. Extract claims with regexes and check each against the wiki:
   - Spellings like `C E G Bb D` or `C – E – G` following a chord/scale name → compare with `theory.spelling` of the matching node (match by title, case-insensitive).
   - Tabs like `x-3-2-0-0-0` near a chord name → must equal one of that node's `musical_forms[form=voicing].tab`.
   - Semitone claims (`N semitones`) next to an interval name → compare with `/intervals/<name>` `theory.semitones[1]`.
   - Key chord lists (`I ii iii …` or chord names for a key) → compare with `/keys/<key>` progression events.
4. Print a report: per guide, `VERIFIED`, `MISMATCH` (with wiki value), and `UNVERIFIABLE` (claims with no matching wiki data — e.g. CAGED shape order, pentatonic positions 4–5, 12-bar layout).

Run it; fix every `MISMATCH` in the guides; re-run until zero mismatches.

- [ ] **Step 3: Commit the guides**

```bash
git add guides/
git commit -m "docs(guides): draft eight explanation pages

Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>"
```

- [ ] **Step 4: Owner checkpoint** — send the owner the list of guide files and the checker report (verified counts per guide + every `UNVERIFIABLE` claim verbatim). Wait for review; apply edits; commit as `docs(guides): apply owner review`.

---

### Task 3: `CORPUS_DIRS` config

**Files:**
- Modify: `internal/config/config.go`
- Test: `internal/config/config_test.go`

**Interfaces:**
- Produces: `Config.CorpusDirs []string` — trimmed, non-empty entries, in config order.

- [ ] **Step 1: Write the failing test** (append to `config_test.go`)

```go
func TestCorpusDirs(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    []string
		wantErr bool
	}{
		{name: "unset uses default", value: "", want: []string{"musicwiki", "guides"}},
		{name: "trims and drops empties", value: " musicwiki , ,guides,", want: []string{"musicwiki", "guides"}},
		{name: "single dir", value: "guides", want: []string{"guides"}},
		{name: "only separators is an error", value: " , ,", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setEnv(t, withOverride(minimalEnv(), "CORPUS_DIRS", tt.value))
			c, err := Load()
			if tt.wantErr {
				if err == nil {
					t.Fatal("Load() succeeded, want error")
				}
				return
			}
			if err != nil {
				t.Fatalf("Load() error: %v", err)
			}
			if !slices.Equal(c.CorpusDirs, tt.want) {
				t.Errorf("CorpusDirs = %q, want %q", c.CorpusDirs, tt.want)
			}
		})
	}
}
```

Add `"slices"` to the test file's imports. (`withOverride` already exists at `config_test.go:137`; setting `CORPUS_DIRS` to `""` via `t.Setenv` exercises the default even if the developer's shell exports it.)

- [ ] **Step 2: Run to verify it fails**

Run: `go test ./internal/config -run TestCorpusDirs`
Expected: FAIL — `c.CorpusDirs undefined`.

- [ ] **Step 3: Implement**

In the `Config` struct, after `EmbeddingDimensions`:

```go
	// CorpusDirs lists the ingest source directories. A directory whose base
	// name is "musicwiki" is read as an OKF bundle; any other as guides.
	CorpusDirs []string
```

In `Load`, add a helper beside `optStr`:

```go
	optList := func(key, def string) []string {
		v := strings.TrimSpace(os.Getenv(key))
		if v == "" {
			v = def
		}
		var out []string
		for _, part := range strings.Split(v, ",") {
			if p := strings.TrimSpace(part); p != "" {
				out = append(out, p)
			}
		}
		if len(out) == 0 {
			errs = append(errs, fmt.Sprintf("%s must list at least one directory", key))
		}
		return out
	}
```

and in the `Config{…}` literal after `EmbeddingDimensions: reqInt("EMBEDDING_DIMENSIONS"),`:

```go
		CorpusDirs: optList("CORPUS_DIRS", "musicwiki,guides"),
```

- [ ] **Step 4: Run tests**

Run: `go test ./internal/config`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
gofmt -l . && go vet ./... && staticcheck ./...
git add internal/config/
git commit -m "feat(config): CORPUS_DIRS lists ingest source directories

Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>"
```

---

### Task 4: Heading-aware markdown chunker

**Files:**
- Create: `internal/rag/chunk.go`
- Test: `internal/rag/chunk_test.go`

**Interfaces:**
- Produces:
  - `type Chunk struct { Section string; Text string }`
  - `func ChunkMarkdown(body string) []Chunk` — splits on `## ` headings (heading text → `Section`, heading line not in `Text`); text before the first `##` gets `Section: "Introduction"`; `# ` lines dropped; empty sections skipped; a section longer than 3,200 chars (~800 tokens) is split at paragraph boundaries into `"<Section>"`, `"<Section> (part 2)"`, …, each later part starting with trailing paragraphs of the previous part totalling ≤ 400 chars (~100 tokens). A single paragraph longer than the max is kept whole.

- [ ] **Step 1: Write the failing test**

```go
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
```

- [ ] **Step 2: Run to verify it fails**

Run: `go test ./internal/rag`
Expected: FAIL — `undefined: ChunkMarkdown`.

- [ ] **Step 3: Implement `internal/rag/chunk.go`**

```go
// Package rag holds the retrieval-augmented generation pipeline: chunking,
// embedding, retrieval, and prompt assembly.
package rag

import (
	"fmt"
	"regexp"
	"strings"
)

// Chunk is one embeddable piece of a document.
type Chunk struct {
	Section string
	Text    string
}

// Size limits approximate tokens as ~4 characters each, so no tokenizer
// dependency is needed: ~800 tokens per chunk, ~100 tokens of overlap.
const (
	maxChunkChars     = 3200
	overlapChunkChars = 400
)

var paragraphBreak = regexp.MustCompile(`\n[ \t]*\n`)

// ChunkMarkdown splits a Markdown body on "## " headings. Text before the first
// heading becomes the "Introduction" section; "# " title lines are dropped
// because the document title is carried separately. Sections longer than
// maxChunkChars are split at paragraph boundaries with a small overlap.
func ChunkMarkdown(body string) []Chunk {
	var out []Chunk
	section := "Introduction"
	var buf []string
	flush := func() {
		text := strings.TrimSpace(strings.Join(buf, "\n"))
		buf = buf[:0]
		if text != "" {
			out = append(out, splitSection(section, text)...)
		}
	}
	for _, line := range strings.Split(strings.ReplaceAll(body, "\r\n", "\n"), "\n") {
		switch {
		case strings.HasPrefix(line, "## "):
			flush()
			section = strings.TrimSpace(strings.TrimPrefix(line, "## "))
		case strings.HasPrefix(line, "# "):
			// Title line: the document title is stored on the document.
		default:
			buf = append(buf, line)
		}
	}
	flush()
	return out
}

func splitSection(section, text string) []Chunk {
	if len(text) <= maxChunkChars {
		return []Chunk{{Section: section, Text: text}}
	}
	var paras []string
	for _, p := range paragraphBreak.Split(text, -1) {
		if p = strings.TrimSpace(p); p != "" {
			paras = append(paras, p)
		}
	}

	var chunks []Chunk
	var cur []string
	curLen := 0
	emit := func() {
		name := section
		if len(chunks) > 0 {
			name = fmt.Sprintf("%s (part %d)", section, len(chunks)+1)
		}
		chunks = append(chunks, Chunk{Section: name, Text: strings.Join(cur, "\n\n")})
	}
	for _, p := range paras {
		if curLen > 0 && curLen+len(p)+2 > maxChunkChars {
			emit()
			cur, curLen = overlapTail(cur)
		}
		cur = append(cur, p)
		curLen += len(p) + 2
	}
	emit()
	return chunks
}

// overlapTail returns the trailing paragraphs of paras whose combined length
// fits in overlapChunkChars, as a fresh slice, plus that length.
func overlapTail(paras []string) ([]string, int) {
	n, total := 0, 0
	for i := len(paras) - 1; i >= 0; i-- {
		if total+len(paras[i])+2 > overlapChunkChars {
			break
		}
		total += len(paras[i]) + 2
		n++
	}
	return append([]string(nil), paras[len(paras)-n:]...), total
}
```

- [ ] **Step 4: Run tests**

Run: `go test ./internal/rag`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
gofmt -l . && go vet ./... && staticcheck ./...
git add internal/rag/
git commit -m "feat(rag): heading-aware markdown chunker

Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>"
```

---

### Task 5: OKF node parsing and validation

**Files:**
- Create: `internal/corpus/corpus.go`
- Create: `internal/corpus/okf.go`
- Test: `internal/corpus/okf_test.go`
- Modify: `go.mod`, `go.sum` (add `go.yaml.in/yaml/v3`)

**Interfaces:**
- Produces (package `corpus`):
  - `const WikiPrefix = "wiki:"`, `const GuidePrefix = "guide:"`
  - `type Chunk = rag.Chunk`
  - `type Document struct { ID string; Title string; Chunks []Chunk }`
  - unexported: `splitFrontmatter(raw []byte) (front []byte, body string, err error)`, `parseNode(path string, raw []byte) (node, error)`, and the `node`/`theory`/`render`/`position`/`frettedNote`/`form`/`degreeEvent` types below (used by Tasks 6–8).

- [ ] **Step 1: Add the dependency**

Run: `go get go.yaml.in/yaml/v3@v3.0.4`
Justification (agent.md requires one): YAML frontmatter is the bundle's data format and the standard library has no YAML parser. `go.yaml.in/yaml/v3` is the maintained continuation of `gopkg.in/yaml.v3` (archived).

- [ ] **Step 2: Write the failing test** (`internal/corpus/okf_test.go`)

```go
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
```

- [ ] **Step 3: Run to verify it fails**

Run: `go test ./internal/corpus`
Expected: FAIL — `undefined: parseNode`.

- [ ] **Step 4: Implement `internal/corpus/corpus.go`**

```go
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
```

- [ ] **Step 5: Implement `internal/corpus/okf.go`**

```go
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
```

- [ ] **Step 6: Run tests**

Run: `go test ./internal/corpus`
Expected: PASS.

- [ ] **Step 7: Commit**

```bash
go mod tidy && gofmt -l . && go vet ./... && staticcheck ./...
git add go.mod go.sum internal/corpus/
git commit -m "feat(corpus): parse and validate musicwiki OKF nodes

Adds go.yaml.in/yaml/v3 (maintained yaml.v3): the bundle's data is YAML
frontmatter and the standard library has no YAML parser.

Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>"
```

---

### Task 6: Flatten nodes into prose chunks

**Files:**
- Create: `internal/corpus/flatten.go`
- Test: `internal/corpus/flatten_test.go`

**Interfaces:**
- Consumes: `node`, `form`, `position`, `frettedNote`, `degreeEvent` (Task 5).
- Produces: `func flatten(n node, titles map[string]string) ([]Chunk, error)` — `titles` maps every bundle slug to its title; an unresolved slug is an error. Section names: `overview`, `pattern N (frets a–b)`, `voicing N (<tab>)`, `progression` (`progression N` for N > 1).

- [ ] **Step 1: Write the failing test** (`internal/corpus/flatten_test.go`)

```go
package corpus

import (
	"strings"
	"testing"
)

var testTitles = map[string]string{
	"/scales/a-natural-minor-scale":     "A Natural Minor Scale",
	"/scales/c-major-pentatonic-scale":  "C Major Pentatonic Scale",
	"/scales/a-blues-scale":             "A Blues Scale",
	"/triads/c-major-triad":             "C Major Triad",
	"/chords/c-major-7th-chord":         "C Major 7th Chord",
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
```

- [ ] **Step 2: Run to verify it fails**

Run: `go test ./internal/corpus -run TestFlatten`
Expected: FAIL — `undefined: flatten`.

- [ ] **Step 3: Implement `internal/corpus/flatten.go`**

```go
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
```

- [ ] **Step 4: Run tests**

Run: `go test ./internal/corpus`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
gofmt -l . && go vet ./... && staticcheck ./...
git add internal/corpus/
git commit -m "feat(corpus): flatten OKF frontmatter into prose chunks

Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>"
```

---

### Task 7: `ReadMusicWiki` over the real bundle

**Files:**
- Create: `internal/corpus/musicwiki.go`
- Test: `internal/corpus/musicwiki_test.go`

**Interfaces:**
- Consumes: `parseNode` (Task 5), `flatten` (Task 6).
- Produces: `func ReadMusicWiki(dir string) ([]Document, error)` — documents in folder order `intervals, scales, triads, chords, keys`, files sorted by name; ids `WikiPrefix + slug`. Missing subfolders are skipped (the zero-document guard lives in `Load`, Task 8); duplicate slugs are an error.

- [ ] **Step 1: Write the failing test** (`internal/corpus/musicwiki_test.go`)

```go
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
```

- [ ] **Step 2: Run to verify it fails**

Run: `go test ./internal/corpus -run 'ReadMusicWiki|MinorPentatonic'`
Expected: FAIL — `undefined: ReadMusicWiki`.

- [ ] **Step 3: Implement `internal/corpus/musicwiki.go`**

```go
package corpus

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// wikiFolders are the OKF concept folders, in ingest order. Bundle-root files
// (index.md, log.md, README.md, LICENSE) and every index.md are not concepts.
var wikiFolders = []string{"intervals", "scales", "triads", "chords", "keys"}

// ReadMusicWiki reads and flattens every concept node in an OKF bundle. All
// nodes are parsed before any is flattened, so relationship slugs resolve to
// titles across the whole bundle.
func ReadMusicWiki(dir string) ([]Document, error) {
	var nodes []node
	titles := map[string]string{}
	for _, sub := range wikiFolders {
		entries, err := os.ReadDir(filepath.Join(dir, sub))
		if errors.Is(err, fs.ErrNotExist) {
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", filepath.Join(dir, sub), err)
		}
		for _, e := range entries {
			if e.IsDir() || filepath.Ext(e.Name()) != ".md" || e.Name() == "index.md" {
				continue
			}
			path := filepath.Join(dir, sub, e.Name())
			raw, err := os.ReadFile(path)
			if err != nil {
				return nil, fmt.Errorf("read %s: %w", path, err)
			}
			n, err := parseNode(path, raw)
			if err != nil {
				return nil, err
			}
			if _, dup := titles[n.Slug]; dup {
				return nil, fmt.Errorf("%s: duplicate slug %q", path, n.Slug)
			}
			titles[n.Slug] = n.Title
			nodes = append(nodes, n)
		}
	}

	docs := make([]Document, 0, len(nodes))
	for _, n := range nodes {
		chunks, err := flatten(n, titles)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", n.Path, err)
		}
		docs = append(docs, Document{ID: WikiPrefix + n.Slug, Title: n.Title, Chunks: chunks})
	}
	return docs, nil
}
```

- [ ] **Step 4: Run tests**

Run: `go test ./internal/corpus`
Expected: PASS. If the real-bundle counts differ, print `kinds` and compare with the bundle (`grep -c` on the YAML) before changing the expected numbers — the numbers above were measured on 2026-11-08 (409 nodes, 429 patterns, 684 voicings, 24 progressions).

- [ ] **Step 5: Commit**

```bash
gofmt -l . && go vet ./... && staticcheck ./...
git add internal/corpus/
git commit -m "feat(corpus): read the musicwiki bundle into documents

Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>"
```

---

### Task 8: `ReadGuides` and `Load` (source guards)

**Files:**
- Create: `internal/corpus/guides.go`
- Create: `internal/corpus/load.go`
- Test: `internal/corpus/guides_test.go`, `internal/corpus/load_test.go`

**Interfaces:**
- Consumes: `splitFrontmatter` (Task 5), `rag.ChunkMarkdown` (Task 4), `ReadMusicWiki` (Task 7).
- Produces:
  - `func ReadGuides(dir string, wikiSlugs map[string]bool) ([]Document, error)` — every `*.md` directly in `dir` except `README.md`, sorted; id `GuidePrefix + stem`; chunk text `"<Title> — <Section>\n<text>"`.
  - `type Source struct { Dir string; Prefix string; Docs []Document }`
  - `func Load(dirs []string) ([]Source, error)` — wiki sources first, then guides; errors on missing dir, non-directory, zero documents, or more than one dir of a kind.

- [ ] **Step 1: Write the failing tests**

`internal/corpus/guides_test.go`:

```go
package corpus

import (
	"strings"
	"testing"
)

func TestReadGuides(t *testing.T) {
	dir := writeTree(t, map[string]string{
		"modes.md": "---\ntitle: Modes\nrelated: [/scales/d-dorian-scale]\n---\n\n# Modes\n\nIntro.\n\n## Dorian\nNatural 6.\n",
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
```

`internal/corpus/load_test.go`:

```go
package corpus

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestLoad(t *testing.T) {
	root := writeTree(t, map[string]string{
		"musicwiki/intervals/u.md": "---\ntype: interval\nslug: /intervals/u\ntitle: U\ntheory: {}\n---\n",
		"guides/g.md":              "---\ntitle: G\nrelated: [/intervals/u]\n---\n\n## A\nB.\n",
		"empty/.keep":              "",
		"notadir":                  "file",
	})
	wiki := filepath.Join(root, "musicwiki")
	guides := filepath.Join(root, "guides")

	t.Run("guides listed first still resolve wiki slugs", func(t *testing.T) {
		got, err := Load([]string{guides, wiki})
		if err != nil {
			t.Fatalf("Load: %v", err)
		}
		if len(got) != 2 || got[0].Prefix != WikiPrefix || got[1].Prefix != GuidePrefix {
			t.Fatalf("sources = %+v, want wiki then guides", got)
		}
	})
	t.Run("trailing slash and ./ still detect musicwiki", func(t *testing.T) {
		got, err := Load([]string{wiki + "/", guides})
		if err != nil || got[0].Prefix != WikiPrefix {
			t.Fatalf("Load = %+v, %v; want wiki source", got, err)
		}
	})

	errs := []struct {
		name    string
		dirs    []string
		wantErr string
	}{
		{"missing dir", []string{wiki, filepath.Join(root, "gudies")}, "gudies"},
		{"not a directory", []string{filepath.Join(root, "notadir")}, "not a directory"},
		{"empty guides dir", []string{wiki, filepath.Join(root, "empty")}, "no documents"},
		{"two guide dirs share a prefix", []string{guides, filepath.Join(root, "empty")}, "at most one"},
		{"two wiki dirs", []string{wiki, wiki}, "at most one"},
		{"guide related needs the wiki", []string{guides}, `related slug "/intervals/u"`},
	}
	for _, tt := range errs {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Load(tt.dirs)
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("err = %v, want containing %q", err, tt.wantErr)
			}
		})
	}
}
```

- [ ] **Step 2: Run to verify they fail**

Run: `go test ./internal/corpus -run 'Guides|Load'`
Expected: FAIL — `undefined: ReadGuides`, `undefined: Load`.

- [ ] **Step 3: Implement `internal/corpus/guides.go`**

```go
package corpus

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go.yaml.in/yaml/v3"

	"github.com/mharner33/guitar-chat/internal/rag"
)

type guideFront struct {
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
	Tags        []string `yaml:"tags"`
	Related     []string `yaml:"related"`
}

// ReadGuides reads every hand-written guide in dir. Each guide's related slugs
// must name musicwiki nodes, so a typo fails ingest instead of leaving a dead
// link.
func ReadGuides(dir string, wikiSlugs map[string]bool) ([]Document, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", dir, err)
	}
	var docs []Document
	for _, e := range entries {
		if e.IsDir() || filepath.Ext(e.Name()) != ".md" || e.Name() == "README.md" {
			continue
		}
		path := filepath.Join(dir, e.Name())
		doc, err := readGuide(path, wikiSlugs)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", path, err)
		}
		docs = append(docs, doc)
	}
	return docs, nil
}

func readGuide(path string, wikiSlugs map[string]bool) (Document, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return Document{}, err
	}
	front, body, err := splitFrontmatter(raw)
	if err != nil {
		return Document{}, err
	}
	var g guideFront
	if err := yaml.Unmarshal(front, &g); err != nil {
		return Document{}, fmt.Errorf("parse frontmatter: %w", err)
	}
	if strings.TrimSpace(g.Title) == "" {
		return Document{}, fmt.Errorf("missing required field %q", "title")
	}
	for _, s := range g.Related {
		if !wikiSlugs[s] {
			return Document{}, fmt.Errorf("related slug %q is not a musicwiki node", s)
		}
	}
	sections := rag.ChunkMarkdown(body)
	if len(sections) == 0 {
		return Document{}, fmt.Errorf("guide has no content")
	}
	chunks := make([]Chunk, len(sections))
	for i, s := range sections {
		chunks[i] = Chunk{Section: s.Section, Text: g.Title + " — " + s.Section + "\n" + s.Text}
	}
	stem := strings.TrimSuffix(filepath.Base(path), ".md")
	return Document{ID: GuidePrefix + stem, Title: g.Title, Chunks: chunks}, nil
}
```

- [ ] **Step 4: Implement `internal/corpus/load.go`**

```go
package corpus

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Source is one ingest source: its directory, the id prefix it owns, and its
// documents. Reconcile deletes only documents with this prefix.
type Source struct {
	Dir    string
	Prefix string
	Docs   []Document
}

// Load reads every configured source before anything is written. A directory
// whose base name is "musicwiki" is read as an OKF bundle; any other as
// guides. Wiki sources are read first so guides can validate related slugs.
// It refuses missing or empty directories (a typo must not reconcile a whole
// source away) and more than one directory per kind (they would share a
// prefix and delete each other's documents).
func Load(dirs []string) ([]Source, error) {
	var wikiDirs, guideDirs []string
	for _, d := range dirs {
		if filepath.Base(filepath.Clean(d)) == "musicwiki" {
			wikiDirs = append(wikiDirs, d)
		} else {
			guideDirs = append(guideDirs, d)
		}
	}
	if len(wikiDirs) > 1 || len(guideDirs) > 1 {
		return nil, fmt.Errorf("CORPUS_DIRS %q: at most one musicwiki and one guides directory (each kind owns one id prefix)", strings.Join(dirs, ","))
	}

	var sources []Source
	wikiSlugs := map[string]bool{}
	for _, d := range wikiDirs {
		if err := requireDir(d); err != nil {
			return nil, err
		}
		docs, err := ReadMusicWiki(d)
		if err != nil {
			return nil, err
		}
		if len(docs) == 0 {
			return nil, fmt.Errorf("corpus dir %q: no documents", d)
		}
		for _, doc := range docs {
			wikiSlugs[strings.TrimPrefix(doc.ID, WikiPrefix)] = true
		}
		sources = append(sources, Source{Dir: d, Prefix: WikiPrefix, Docs: docs})
	}
	for _, d := range guideDirs {
		if err := requireDir(d); err != nil {
			return nil, err
		}
		docs, err := ReadGuides(d, wikiSlugs)
		if err != nil {
			return nil, err
		}
		if len(docs) == 0 {
			return nil, fmt.Errorf("corpus dir %q: no documents", d)
		}
		sources = append(sources, Source{Dir: d, Prefix: GuidePrefix, Docs: docs})
	}
	return sources, nil
}

func requireDir(d string) error {
	fi, err := os.Stat(d)
	if err != nil {
		return fmt.Errorf("corpus dir %q: %w", d, err)
	}
	if !fi.IsDir() {
		return fmt.Errorf("corpus dir %q: not a directory", d)
	}
	return nil
}
```

- [ ] **Step 5: Run tests**

Run: `go test ./internal/corpus`
Expected: PASS. (`TestRealGuidesLoad` needs Task 2's guides; if Task 2 is not merged yet, it fails with "no such file" — finish Task 2 first.)

- [ ] **Step 6: Commit**

```bash
gofmt -l . && go vet ./... && staticcheck ./...
git add internal/corpus/
git commit -m "feat(corpus): read guides and load sources with prefix guards

Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>"
```

---

### Task 9: `Embedder` interface and HTTP client

**Files:**
- Create: `internal/rag/embed.go`
- Test: `internal/rag/embed_test.go`

**Interfaces:**
- Produces:
  - `type Embedder interface { Embed(ctx context.Context, texts []string) ([][]float32, error) }`
  - `type HTTPEmbedder struct { BaseURL, APIKey, Model string; Dimensions int; Client *http.Client }` with `Embed` — POSTs `{"model", "input"}` to `BaseURL + "/embeddings"`; returns vectors in input order; errors on non-2xx (status + first 512 bytes of body), wrong count, bad/duplicate index, or wrong dimension.

- [ ] **Step 1: Write the failing test**

```go
package rag

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func embedServer(t *testing.T, handler func(w http.ResponseWriter, input []string)) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/embeddings" || r.Header.Get("Authorization") != "Bearer k" {
			t.Errorf("got %s %s auth=%q", r.Method, r.URL.Path, r.Header.Get("Authorization"))
		}
		var req struct {
			Model string   `json:"model"`
			Input []string `json:"input"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Model != "m" {
			t.Errorf("bad request body: %v model=%q", err, req.Model)
		}
		handler(w, req.Input)
	}))
	t.Cleanup(srv.Close)
	return srv
}

func TestHTTPEmbedder(t *testing.T) {
	tests := []struct {
		name    string
		reply   string
		status  int
		wantErr string
		want    [][]float32
	}{
		{
			name:  "reorders by index",
			reply: `{"data":[{"index":1,"embedding":[3,4]},{"index":0,"embedding":[1,2]}]}`,
			want:  [][]float32{{1, 2}, {3, 4}},
		},
		{name: "wrong dimension", reply: `{"data":[{"index":0,"embedding":[1]},{"index":1,"embedding":[3,4]}]}`, wantErr: "got 1 dimensions, want 2"},
		{name: "wrong count", reply: `{"data":[{"index":0,"embedding":[1,2]}]}`, wantErr: "1 embeddings for 2 inputs"},
		{name: "duplicate index", reply: `{"data":[{"index":0,"embedding":[1,2]},{"index":0,"embedding":[1,2]}]}`, wantErr: "bad index 0"},
		{name: "http error", status: 429, reply: `{"error":"rate limited"}`, wantErr: "429"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := embedServer(t, func(w http.ResponseWriter, _ []string) {
				if tt.status != 0 {
					w.WriteHeader(tt.status)
				}
				_, _ = w.Write([]byte(tt.reply))
			})
			e := &HTTPEmbedder{BaseURL: srv.URL + "/v1/", APIKey: "k", Model: "m", Dimensions: 2}
			got, err := e.Embed(context.Background(), []string{"a", "b"})
			if tt.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("err = %v, want containing %q", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("Embed: %v", err)
			}
			for i := range tt.want {
				if got[i][0] != tt.want[i][0] || got[i][1] != tt.want[i][1] {
					t.Errorf("vec %d = %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestHTTPEmbedderEmptyInputMakesNoRequest(t *testing.T) {
	e := &HTTPEmbedder{BaseURL: "http://127.0.0.1:1", Model: "m", Dimensions: 2}
	got, err := e.Embed(context.Background(), nil)
	if err != nil || got != nil {
		t.Fatalf("Embed(nil) = %v, %v; want nil, nil", got, err)
	}
}
```

- [ ] **Step 2: Run to verify it fails**

Run: `go test ./internal/rag -run Embedder`
Expected: FAIL — `undefined: HTTPEmbedder`.

- [ ] **Step 3: Implement `internal/rag/embed.go`**

```go
package rag

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Embedder turns texts into vectors. It is defined here, at its consumers
// (ingest and the answer pipeline), so tests can fake it.
type Embedder interface {
	Embed(ctx context.Context, texts []string) ([][]float32, error)
}

// HTTPEmbedder calls an OpenAI-compatible /embeddings endpoint.
type HTTPEmbedder struct {
	BaseURL    string // e.g. https://api.openai.com/v1
	APIKey     string
	Model      string
	Dimensions int // expected vector width (EMBEDDING_DIMENSIONS)
	Client     *http.Client
}

// Embed returns one vector per text, in input order. The caller's context
// carries the deadline.
func (e *HTTPEmbedder) Embed(ctx context.Context, texts []string) ([][]float32, error) {
	if len(texts) == 0 {
		return nil, nil
	}
	body, err := json.Marshal(map[string]any{"model": e.Model, "input": texts})
	if err != nil {
		return nil, fmt.Errorf("embeddings: encode request: %w", err)
	}
	url := strings.TrimRight(e.BaseURL, "/") + "/embeddings"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("embeddings: build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+e.APIKey)

	client := e.Client
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("embeddings request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		msg, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		return nil, fmt.Errorf("embeddings: %s: %s", resp.Status, strings.TrimSpace(string(msg)))
	}

	var out struct {
		Data []struct {
			Index     int       `json:"index"`
			Embedding []float32 `json:"embedding"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, fmt.Errorf("embeddings: decode response: %w", err)
	}
	if len(out.Data) != len(texts) {
		return nil, fmt.Errorf("embeddings: got %d embeddings for %d inputs", len(out.Data), len(texts))
	}
	vecs := make([][]float32, len(texts))
	for _, d := range out.Data {
		if d.Index < 0 || d.Index >= len(texts) || vecs[d.Index] != nil {
			return nil, fmt.Errorf("embeddings: bad index %d", d.Index)
		}
		if len(d.Embedding) != e.Dimensions {
			return nil, fmt.Errorf("embeddings: got %d dimensions, want %d (EMBEDDING_DIMENSIONS)", len(d.Embedding), e.Dimensions)
		}
		vecs[d.Index] = d.Embedding
	}
	return vecs, nil
}
```

- [ ] **Step 4: Run tests**

Run: `go test ./internal/rag`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
gofmt -l . && go vet ./... && staticcheck ./...
git add internal/rag/
git commit -m "feat(rag): Embedder interface and OpenAI-compatible HTTP client

Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>"
```

---

### Task 10: Store — chunk ids, document sync, prefix reconcile

**Files:**
- Create: `internal/store/corpus.go`
- Test: `internal/store/corpus_test.go`

**Interfaces:**
- Produces:
  - `func ChunkID(docID, section, text string) string` — hex sha256 of the three fields, each followed by a NUL byte.
  - `type NewChunk struct { ID, Section, Text string; Embedding []float32 }`
  - `func (s *Store) ChunkIDs(ctx context.Context, docID string) (map[string]bool, error)`
  - `func (s *Store) SyncDocument(ctx context.Context, docID, title string, keepIDs []string, added []NewChunk) error` — one transaction: upsert document, insert `added` (`ON CONFLICT DO NOTHING`), delete the document's chunks not in `keepIDs`. Empty `keepIDs` deletes all of the document's chunks.
  - `func (s *Store) DeleteMissingDocuments(ctx context.Context, prefix string, keep []string) (int64, error)` — deletes documents whose id starts with `prefix` and is not in `keep` (chunks cascade). Refuses an empty prefix.

- [ ] **Step 1: Write the failing tests**

```go
package store

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"os"
	"testing"
	"time"
)

func TestChunkID(t *testing.T) {
	a := ChunkID("wiki:/x", "overview", "text")
	if len(a) != 64 {
		t.Fatalf("len = %d, want 64 hex chars", len(a))
	}
	if a != ChunkID("wiki:/x", "overview", "text") {
		t.Error("ChunkID not deterministic")
	}
	// Field boundaries matter: moving characters between fields changes the id.
	if ChunkID("ab", "c", "d") == ChunkID("a", "bc", "d") {
		t.Error("ChunkID ignores field boundaries")
	}
}

func TestVectorLiteral(t *testing.T) {
	if got := vectorLiteral([]float32{0.5, -1, 2.25}); got != "[0.5,-1,2.25]" {
		t.Errorf("vectorLiteral = %q", got)
	}
}

// Integration: requires DATABASE_URL with migrations applied (vector(1536)).
func TestSyncAndReconcile(t *testing.T) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Skip("DATABASE_URL not set; skipping store integration test")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	st, err := Open(ctx, dsn)
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	defer st.Close()

	// Unique prefixes isolate this run from real data and other runs.
	pa, pb := testPrefix(t), testPrefix(t)
	t.Cleanup(func() {
		for _, p := range []string{pa, pb} {
			_, _ = st.Pool().Exec(context.Background(), `DELETE FROM documents WHERE starts_with(id, $1)`, p)
		}
	})
	vec := make([]float32, 1536)
	mk := func(doc, text string) NewChunk {
		return NewChunk{ID: ChunkID(doc, "s", text), Section: "s", Text: text, Embedding: vec}
	}

	docA, docB := pa+"one", pb+"one"
	c1, c2 := mk(docA, "first"), mk(docA, "second")
	if err := st.SyncDocument(ctx, docA, "One", []string{c1.ID, c2.ID}, []NewChunk{c1, c2}); err != nil {
		t.Fatalf("SyncDocument A: %v", err)
	}
	b1 := mk(docB, "b")
	if err := st.SyncDocument(ctx, docB, "B", []string{b1.ID}, []NewChunk{b1}); err != nil {
		t.Fatalf("SyncDocument B: %v", err)
	}

	// Re-sync keeping only c1: c2 is deleted, nothing re-inserted.
	if err := st.SyncDocument(ctx, docA, "One", []string{c1.ID}, nil); err != nil {
		t.Fatalf("resync: %v", err)
	}
	ids, err := st.ChunkIDs(ctx, docA)
	if err != nil || len(ids) != 1 || !ids[c1.ID] {
		t.Fatalf("ChunkIDs after resync = %v, %v; want only c1", ids, err)
	}

	// Empty keep list removes every chunk of that document (not none).
	if err := st.SyncDocument(ctx, docA, "One", nil, nil); err != nil {
		t.Fatalf("empty resync: %v", err)
	}
	if ids, _ := st.ChunkIDs(ctx, docA); len(ids) != 0 {
		t.Fatalf("ChunkIDs after empty keep = %v, want none", ids)
	}

	// Reconcile prefix A with nothing kept: deletes A's doc, not B's.
	n, err := st.DeleteMissingDocuments(ctx, pa, nil)
	if err != nil || n != 1 {
		t.Fatalf("DeleteMissingDocuments(A) = %d, %v; want 1", n, err)
	}
	if ids, _ := st.ChunkIDs(ctx, docB); len(ids) != 1 {
		t.Fatalf("source B chunks = %v, want untouched", ids)
	}
	if _, err := st.DeleteMissingDocuments(ctx, "", nil); err == nil {
		t.Fatal("DeleteMissingDocuments with empty prefix succeeded, want error")
	}
}

func testPrefix(t *testing.T) string {
	t.Helper()
	b := make([]byte, 6)
	if _, err := rand.Read(b); err != nil {
		t.Fatal(err)
	}
	return "test-" + hex.EncodeToString(b) + ":"
}
```

- [ ] **Step 2: Run to verify it fails**

Run: `go test ./internal/store`
Expected: FAIL — `undefined: ChunkID`.

- [ ] **Step 3: Implement `internal/store/corpus.go`**

```go
package store

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
)

// ChunkID is a chunk's content hash: sha256 over doc id, section and text,
// each NUL-terminated so field boundaries count. Identical content keeps the
// same id across runs, which is what makes ingest idempotent.
func ChunkID(docID, section, text string) string {
	h := sha256.New()
	for _, part := range []string{docID, section, text} {
		h.Write([]byte(part))
		h.Write([]byte{0})
	}
	return hex.EncodeToString(h.Sum(nil))
}

// NewChunk is a chunk to insert, with its embedding.
type NewChunk struct {
	ID        string
	Section   string
	Text      string
	Embedding []float32
}

// ChunkIDs returns the ids of a document's stored chunks.
func (s *Store) ChunkIDs(ctx context.Context, docID string) (map[string]bool, error) {
	rows, err := s.pool.Query(ctx, `SELECT id FROM chunks WHERE document_id = $1`, docID)
	if err != nil {
		return nil, fmt.Errorf("list chunks of %s: %w", docID, err)
	}
	ids, err := pgx.CollectRows(rows, pgx.RowTo[string])
	if err != nil {
		return nil, fmt.Errorf("list chunks of %s: %w", docID, err)
	}
	out := make(map[string]bool, len(ids))
	for _, id := range ids {
		out[id] = true
	}
	return out, nil
}

// SyncDocument makes a document's stored chunks exactly keepIDs, in one
// transaction: upsert the document, insert added chunks, delete the rest.
func (s *Store) SyncDocument(ctx context.Context, docID, title string, keepIDs []string, added []NewChunk) error {
	if keepIDs == nil {
		keepIDs = []string{} // nil would encode as NULL, and NOT (id = ANY(NULL)) deletes nothing
	}
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("sync %s: begin: %w", docID, err)
	}
	defer tx.Rollback(ctx) //nolint:errcheck // no-op after a successful Commit

	if _, err := tx.Exec(ctx, `
		INSERT INTO documents (id, title) VALUES ($1, $2)
		ON CONFLICT (id) DO UPDATE SET title = EXCLUDED.title, updated_at = now()`,
		docID, title); err != nil {
		return fmt.Errorf("sync %s: upsert document: %w", docID, err)
	}
	for _, c := range added {
		if _, err := tx.Exec(ctx, `
			INSERT INTO chunks (id, document_id, section, text, embedding)
			VALUES ($1, $2, $3, $4, $5::text::vector)
			ON CONFLICT (id) DO NOTHING`,
			c.ID, docID, c.Section, c.Text, vectorLiteral(c.Embedding)); err != nil {
			return fmt.Errorf("sync %s: insert chunk %q: %w", docID, c.Section, err)
		}
	}
	if _, err := tx.Exec(ctx,
		`DELETE FROM chunks WHERE document_id = $1 AND NOT (id = ANY($2))`,
		docID, keepIDs); err != nil {
		return fmt.Errorf("sync %s: delete stale chunks: %w", docID, err)
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("sync %s: commit: %w", docID, err)
	}
	return nil
}

// DeleteMissingDocuments deletes documents whose id starts with prefix and is
// not in keep; their chunks cascade. Scoping by prefix keeps one source's
// reconcile from touching another's documents.
func (s *Store) DeleteMissingDocuments(ctx context.Context, prefix string, keep []string) (int64, error) {
	if prefix == "" {
		return 0, errors.New("delete missing documents: empty prefix would match every document")
	}
	if keep == nil {
		keep = []string{}
	}
	tag, err := s.pool.Exec(ctx,
		`DELETE FROM documents WHERE starts_with(id, $1) AND NOT (id = ANY($2))`,
		prefix, keep)
	if err != nil {
		return 0, fmt.Errorf("delete missing %s documents: %w", prefix, err)
	}
	return tag.RowsAffected(), nil
}

// vectorLiteral formats v as a pgvector text literal, e.g. "[0.5,-1]".
func vectorLiteral(v []float32) string {
	var b strings.Builder
	b.WriteByte('[')
	for i, f := range v {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteString(strconv.FormatFloat(float64(f), 'g', -1, 32))
	}
	b.WriteByte(']')
	return b.String()
}
```

- [ ] **Step 4: Run tests**

Run: `go test ./internal/store` → PASS (integration test skipped without a DB).
Then against a local DB: `docker compose --profile local up -d db`, apply migrations by running the current ingest stub once with the required env exported (`set -a; . ./.env; set +a; DATABASE_URL=postgres://app:app@localhost:5432/theory?sslmode=disable go run ./cmd/ingest` — the stub migrates, checks dimensions, and exits 0), then
`DATABASE_URL=postgres://app:app@localhost:5432/theory?sslmode=disable go test ./internal/store -run TestSyncAndReconcile -v` → PASS.

- [ ] **Step 5: Commit**

```bash
gofmt -l . && go vet ./... && staticcheck ./...
git add internal/store/
git commit -m "feat(store): content-hash chunk ids, document sync, prefix-scoped reconcile

Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>"
```

---

### Task 11: Wire `cmd/ingest`, ship the corpus in the image, acceptance run

**Files:**
- Create: `cmd/ingest/ingest.go`
- Test: `cmd/ingest/ingest_test.go`
- Modify: `cmd/ingest/main.go`
- Modify: `Dockerfile`

**Interfaces:**
- Consumes: `corpus.Load`, `corpus.Source`, `corpus.Document` (Task 8); `rag.Embedder`, `rag.HTTPEmbedder` (Task 9); `store.ChunkID`, `store.NewChunk`, `(*store.Store).ChunkIDs/SyncDocument/DeleteMissingDocuments` (Task 10); `config.Config.CorpusDirs` (Task 3).
- Produces: `func ingest(ctx context.Context, st docStore, emb rag.Embedder, sources []corpus.Source) (stats, error)`.

- [ ] **Step 1: Write the failing test** (`cmd/ingest/ingest_test.go`)

```go
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
```

- [ ] **Step 2: Run to verify it fails**

Run: `go test ./cmd/ingest`
Expected: FAIL — `undefined: ingest`.

- [ ] **Step 3: Implement `cmd/ingest/ingest.go`**

```go
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
```

- [ ] **Step 4: Run tests**

Run: `go test ./cmd/ingest`
Expected: PASS.

- [ ] **Step 5: Wire `cmd/ingest/main.go`**

Replace the package comment and `run` with:

```go
// Command ingest chunks and embeds the corpus sources listed in CORPUS_DIRS
// (the musicwiki bundle and the hand-written guides) into pgvector. It is
// idempotent and reconciles each source's documents to what is on disk.
package main
```

```go
func run() error {
	ctx := context.Background()

	// Local dev convenience, same as cmd/api: pick up ./.env if present. Real
	// environment variables take precedence; in containers the file is absent.
	loaded, err := config.LoadDotEnv(".env")
	if err != nil {
		return err
	}
	if loaded {
		slog.Info("loaded .env")
	}

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	// Read and validate every source before touching the database, so a bad
	// file or a mistyped directory cannot delete anything.
	sources, err := corpus.Load(cfg.CorpusDirs)
	if err != nil {
		return err
	}

	st, err := store.Open(ctx, cfg.DatabaseURL)
	if err != nil {
		return err
	}
	defer st.Close()

	if err := st.Migrate(ctx, migrations.FS); err != nil {
		return err
	}
	if err := st.CheckEmbeddingDimensions(ctx, cfg.EmbeddingDimensions); err != nil {
		return err
	}

	emb := &rag.HTTPEmbedder{
		BaseURL:    cfg.EmbedBaseURL(),
		APIKey:     cfg.LLMAPIKey,
		Model:      cfg.LLMEmbedModel,
		Dimensions: cfg.EmbeddingDimensions,
		Client:     &http.Client{Timeout: embedTimeout},
	}
	s, err := ingest(ctx, st, emb, sources)
	if err != nil {
		return err
	}
	slog.Info("ingest complete",
		"documents", s.Documents, "chunks", s.Chunks, "embedded", s.Embedded, "deleted_documents", s.Deleted)
	return nil
}
```

Update imports to: `context`, `log/slog`, `net/http`, `os`, and `internal/config`, `internal/corpus`, `internal/rag`, `internal/store`, `migrations`.

- [ ] **Step 6: Ship the corpus in the image** — in `Dockerfile`, after `COPY --from=build /out/ingest /app/ingest` add:

```dockerfile
# Corpus sources for the one-shot ingest service (CORPUS_DIRS defaults to
# musicwiki,guides, relative to WORKDIR /app).
COPY musicwiki /app/musicwiki
COPY guides    /app/guides
```

and update the header comment's last sentence to: `The web/ UI is embedded into the api binary via go:embed; musicwiki/ and guides/ are copied for the ingest service.`

- [ ] **Step 7: Full checks**

Run: `gofmt -l . && go vet ./... && staticcheck ./... && go test ./...`
Expected: no gofmt output; all PASS.

- [ ] **Step 8: Acceptance run against a real database** (needs `.env` with a working `LLM_API_KEY` for embeddings)

```bash
docker compose --profile local up -d db
export DATABASE_URL='postgres://app:app@localhost:5432/theory?sslmode=disable'
q() { psql "$DATABASE_URL" -Atc "$1"; }

go run ./cmd/ingest                       # expect: documents=417 embedded≈chunks
q "SELECT count(*) FROM documents"         # 417 (409 wiki + 8 guides)
q "SELECT count(*) FROM chunks"            # note N

go run ./cmd/ingest                       # expect: embedded=0 deleted_documents=0
q "SELECT count(*) FROM chunks"            # still N — no duplicates

hold=$(mktemp -d)
mv guides/common-progressions.md "$hold/"
go run ./cmd/ingest                       # expect: deleted_documents=1
q "SELECT count(*) FROM documents WHERE id LIKE 'wiki:%'"   # still 409
mv "$hold/common-progressions.md" guides/
go run ./cmd/ingest                       # guide restored

CORPUS_DIRS=musicwiki,gudies go run ./cmd/ingest; echo "exit=$?"   # non-zero, "gudies" in error
q "SELECT count(*) FROM documents"         # still 417
```

Record the actual counts in the commit message body.

- [ ] **Step 9: Commit**

```bash
git add cmd/ingest/ Dockerfile
git commit -m "feat(ingest): ingest musicwiki + guides with per-source reconcile

Acceptance (local pgvector): <documents>/<chunks> after first run; second
run embedded 0; removing a guide deleted 1 document, wiki untouched;
CORPUS_DIRS typo exited non-zero with no deletes.

Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>"
```

---

## Not in this plan (deliberately)

- **Score-floor measurement (spec §9 step 4)** needs the retriever (`docs/plan.md` step 8), which does not exist yet. Do it as the first step of that work: embed the 16 on-corpus and 2 off-corpus golden questions, record top-5 similarity distributions, and set `RETRIEVAL_MIN_SCORE` in `docs/design.md` §4.3.
- Per-document result cap — only if Milestone 4 evals show crowding (spec §7).
