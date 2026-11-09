# Corpus: musicwiki + hand-written guides (Design Spec)

Date: 2026-11-08
Status: Approved
Milestone: Phase 2 / Milestone 2 (`docs/plan.md` step 7 — ingestion). Supersedes the
corpus parts of `docs/design.md` §3 (corpus source, licenses), §4.1 (`corpus/` layout),
and §5.1 (ingestion input).

## 1. Goal & scope

Replace the planned Open Music Theory `corpus/` with two sources:

- **`musicwiki/`** — the generated Open Knowledge Format (OKF) bundle already in the
  repo (410 concept nodes: intervals, scales, triads, chords, keys). It is the
  **factual** layer: spellings, intervals, fretboard patterns, voicings, diatonic harmony.
- **`guides/`** — 8 hand-written explanation pages. They are the **conceptual** layer:
  the "why" questions the wiki cannot answer (why minor sounds different, which chords
  fit a mode, CAGED, …).

The wiki's value is mostly in its YAML frontmatter (a chord page is ~18 lines of prose
and ~250 lines of YAML), so ingest **flattens frontmatter into prose chunks in Go**.

### In scope
- Design-doc / agent-doc updates for the new corpus, layout, and licensing.
- 8 guide drafts in `guides/`, fact-checked against wiki data.
- `internal/corpus` package (source readers + wiki flattener) and its use from `cmd/ingest`.
- `CORPUS_DIRS` config variable.
- Draft golden-dataset rows (`evals/golden.jsonl`) covering both sources.

### Out of scope
- Lesson generation (`lesson-creator/` port) — separate spec later.
- Changes to the musicwiki generator (lives in another repo). `musicwiki/` is treated
  as read-only generated output, replaced wholesale, never hand-edited here.
- Reranking, hybrid search, per-source quotas (see §6 — only if evals demand it).
- Open Music Theory content. Dropped entirely.

## 2. Layout & licensing

```
guitar-chat/
├── musicwiki/              # generated OKF bundle (MIT, own LICENSE) — read-only here
│   ├── intervals/ scales/ triads/ chords/ keys/
│   └── index.md log.md README.md LICENSE .gitignore   # not ingested
├── guides/                 # hand-written explanation pages (MIT, root LICENSE)
│   ├── why-major-and-minor-sound-different.md
│   ├── modes-and-their-chords.md
│   ├── how-chords-are-built.md
│   ├── reading-the-fretboard.md
│   ├── the-caged-system.md
│   ├── pentatonic-positions.md
│   ├── diatonic-harmony-and-function.md
│   └── common-progressions.md
└── internal/corpus/        # source readers + flattener
```

- `corpus/`, `corpus/LICENSE` (CC BY-SA 4.0) and `corpus/ATTRIBUTION.md` are **removed
  from the design**. Nothing in the repo is adapted from CC BY-SA material.
- All corpus content is MIT: `musicwiki/LICENSE` covers the bundle; the root `LICENSE`
  covers `guides/`. If BY-SA material is ever adapted later, it goes in its own folder
  with its own LICENSE and ATTRIBUTION — never mixed into `guides/`.

### Guide file format

```markdown
---
title: Modes and Their Chords
description: What a mode is, how each one sounds, and which diatonic chords fit it.
tags: [modes, harmony]
related:
  - /scales/d-dorian-scale
  - /keys/c-major
---

# Modes and Their Chords

Intro paragraph…

## What a mode is
…
```

`title` is required; `description`, `tags`, `related` are optional. `related` values are
musicwiki slugs and must resolve to an existing node (ingest validates this — a typo is
an error, not a silent dead link).

## 3. Document IDs & reconcile

| Source | `documents.id` | `documents.title` |
|---|---|---|
| musicwiki | `wiki:` + node `slug` → `wiki:/scales/a-minor-pentatonic-scale` | frontmatter `title` |
| guides | `guide:` + filename stem → `guide:modes-and-their-chords` | frontmatter `title` |

- **No migration.** The existing `documents`/`chunks` schema fits; the prefix carries
  the source.
- **Reconcile is scoped per source.** After walking a source, delete only documents
  whose id has *that source's prefix* and was not seen. Ingesting guides can never
  delete wiki rows and vice versa. Per-document stale-chunk deletion is unchanged
  (`docs/design.md` §5.1).
- **Empty-source guard.** Ingest refuses to run (non-zero exit, nothing deleted) if a
  configured directory is missing or yields zero documents. Without this, a typo in
  `CORPUS_DIRS` would reconcile the whole source away.
- Source chips in the UI show `documents.title`; the id is not user-facing.

## 4. Configuration

| Variable | Default | Meaning |
|---|---|---|
| `CORPUS_DIRS` | `musicwiki,guides` | Comma-separated source directories, relative to the working directory. Each directory's reader is chosen by its name: `musicwiki` → OKF reader, anything else → guides reader. |

Read once in `internal/config`; added to `.env.example`. (Choosing the reader by name
keeps config to one variable; if a second OKF bundle ever appears, switch to
`name=kind` pairs.)

## 5. `internal/corpus`

```go
// Document is one source file, ready to embed.
type Document struct {
    ID     string  // "wiki:/scales/…" or "guide:…"
    Title  string
    Chunks []Chunk
}

type Chunk struct {
    Section string // "overview", "pattern 1 (frets 5–8)", "voicing 1 (x-3-2-0-0-0)", or a ## heading
    Text    string
}

func ReadMusicWiki(dir string) ([]Document, error)
func ReadGuides(dir string, wikiSlugs map[string]bool) ([]Document, error)
```

`cmd/ingest` calls the readers, then embeds (batches of 64) and upserts — it knows
nothing about OKF or guide formats. The heading-aware chunker stays at
`internal/rag/chunk.go` per `docs/design.md` §4.1; `ReadGuides` uses it.

### 5.1 Wiki flattening rules

Ingested folders: `intervals/ scales/ triads/ chords/ keys/`. Skipped: every `index.md`,
plus `log.md`, `README.md`, `LICENSE`, `.gitignore` at the bundle root.

**Every chunk's text starts with the node title** (`"A Minor Pentatonic Scale — Pattern 1 …"`).
Twelve roots produce near-identical text, so the title is what makes the root
retrievable.

| Node `type` | Chunks |
|---|---|
| all | **overview**: title; description; spelling; intervals with semitones; prose body with Markdown links reduced to their text and the `## Related` list turned into sentences; relationship fields as sentences ("Relative major: C Major Pentatonic Scale. Builds on: A Natural Minor Scale. Used by: A Blues Scale.") — slugs resolved to titles. |
| `scale` | + one **pattern** chunk per `musical_forms[form=pattern]`: system, fret range, then notes grouped by string from 6 (low E) to 1 (high E): `string 6 — fret 5 A (1, root), fret 8 C (b3)`. |
| `Concept` (triads, chords) | + one **voicing** chunk per `render.preferred_positions[]`: rank, tab (low→high, `x` = muted), fret range, tags (open/barre), then per string: fret, note, interval, role, finger (0 = open); muted strings listed. |
| `key` | + one **progression** chunk from `musical_forms[form=progression]`: each degree's Roman numeral, triad, seventh chord, and harmonic function. |
| `interval` | overview only. |

Rough volume: ~410 overview + ~430 pattern + ~680 voicing ≈ **1,500 chunks**.

**Determinism.** Output order and wording depend only on the input files (sorted file
walk, fixed field order). Same bundle ⇒ same chunk ids ⇒ re-ingest is a no-op.

**Errors fail loudly.** Unparseable YAML, an unknown `type`, or a missing required
field (`slug`, `title`, `type`, `theory` for non-key nodes) stops ingest with the file
path and field name. The bundle is generated; a schema surprise means the generator
changed and must not be skipped silently. Optional data that is absent (no
voicings in a `--no-voicings` build, no `relative_major`) simply produces fewer chunks
or sentences.

### 5.2 Testing

- Table-driven golden-text tests over small hand-picked fixtures in
  `internal/corpus/testdata/` — one node per type (interval, scale, triad, chord, key)
  plus one guide. Expected chunk text is checked in next to each fixture.
- Error cases: bad YAML, unknown type, missing `slug`, guide `related` slug that does
  not resolve.
- Reconcile scoping: a store-level test that ingesting source A leaves source B's
  documents untouched, and that the empty-source guard deletes nothing.
- Cross-check: the flattened A minor pentatonic pattern chunks contain exactly the
  frets in `lesson-creator/fixtures.py` for the patterns the wiki has (1–3 today).
- No network in unit tests (fakes for `Embedder`, per `agent.md`).

## 6. Guides content

Eight pages, 600–1,500 words each, `##` sections sized to fit one chunk (< ~800 tokens):

1. **Why major and minor sound different** — the third, interval quality, consonance/tension.
2. **Modes and their chords** — modes as rotations of the major scale; per-mode
   characteristic note; the diatonic chords that fit each mode (fills the gap that wiki
   key nodes only harmonize major and natural minor).
3. **How chords are built** — stacking thirds; 6ths, 7ths, 9ths, 11ths, 13ths; sus chords.
4. **Reading the fretboard** — string/fret numbering (1 = high E … 6 = low E, matching
   the wiki), octave shapes, intervals as shapes.
5. **The CAGED system** — the five open shapes moved up the neck.
6. **Pentatonic positions** — the five-position system and how positions connect.
7. **Diatonic harmony and function** — tonic / predominant / dominant; Roman numerals.
8. **Common progressions** — I–IV–V, ii–V–I, I–V–vi–IV, 12-bar blues.

**Accuracy process.** Examples use keys and shapes that exist in the wiki; each guide's
`related` lists the wiki nodes it leans on. Before review, a throwaway checker
(not committed) extracts every concrete note/spelling/fret claim from the drafts and
compares it to wiki YAML. The owner receives the drafts plus the checker's report;
claims the wiki cannot verify (CAGED shape names, progression conventions, positions 4–5
of the pentatonic) are flagged for owner review rather than presented as checked.

**Known wiki gap.** Scale nodes carry ~3 patterns each (429 across 144 scales); the
minor pentatonic has positions 1–3 only. Guide 6 describes all five positions; the
fret details for 4–5 come from the hand-checked `lesson-creator/fixtures.py` and are
flagged as such. Fixing the generator is out of scope here.

## 7. Retrieval balance & evals

With `RETRIEVAL_TOP_K=5` over ~1,500 chunks, the risk is near-duplicate wiki chunks
(other roots, other voicings) crowding out the one relevant guide section. **No
reranking or quota is added up front.** Instead:

- The golden set below includes guide-dependent and wiki-dependent rows, so crowding
  shows up as a correctness/groundedness failure.
- If it does, the first remedy is a per-document cap (max 2 chunks per document) in the
  retrieval SQL — a one-line change, decided from eval data.
- `RETRIEVAL_MIN_SCORE=0.5` predates flattening. After the first real ingest, measure
  score distributions for on- and off-corpus questions and record the chosen floor in
  `docs/design.md` §4.3.

**Demo step 6** (degraded re-ingest ⇒ groundedness drops) uses a guide: replace
`guides/modes-and-their-chords.md` with a wrong version, re-ingest, run evals, restore.

### 7.1 Draft golden rows

Committed now as `evals/golden.jsonl` (`{question, expected_points[], rubric}` per
`docs/design.md` §4.1); consumed unchanged when Milestone 4 builds the harness.
Wiki-fact values below were read from the bundle on 2026-11-08.

| # | Question | Expected points | Source |
|---|---|---|---|
| 1 | Why does a major scale sound different from a minor scale? | the 3rd: major 3rd (4 semitones) vs minor 3rd (3); minor also lowers the 6th and 7th; minor heard as darker | guide 1 |
| 2 | What chords work over D Dorian? | D Dorian = notes of C major; Dm7, Em7, Fmaj7, G7, Am7, Bm7b5, Cmaj7; the natural 6 (B) is the characteristic note | guide 2 + wiki |
| 3 | How is a dominant 7th chord built? | major triad + minor 7th; 1–3–5–b7 | guide 3 + wiki |
| 4 | What is the CAGED system? | five open shapes C, A, G, E, D; movable up the neck; same chord in five places | guide 5 |
| 5 | How do the pentatonic positions connect? | five overlapping positions; adjacent positions share notes; position 1 of A minor at frets 5–8 | guide 6 + wiki |
| 6 | What does the V chord do in a key? | dominant function; pulls/resolves to I (tonic) | guide 7 + wiki |
| 7 | What is a ii–V–I in C major? | Dm7 – G7 – Cmaj7 | guide 8 + wiki |
| 8 | How does a 12-bar blues go? | I–IV–V chords; standard 12-bar layout | guide 8 |
| 9 | What notes are in a G dominant 7th chord? | G, B, D, F | wiki |
| 10 | How do I play an open C major 7th chord? | tab x-3-2-0-0-0 | wiki |
| 11 | What notes are in the E blues scale? | E, G, A, Bb, B, D; the b5 (Bb) is the blue note | wiki |
| 12 | Where is A minor pentatonic position 1 on the fretboard? | frets 5–8; root A at string 6 fret 5 | wiki |
| 13 | What are the notes of an A minor 7 flat 5 chord? | A, C, Eb, G; 1–b3–b5–b7 | wiki |
| 14 | How many semitones is a tritone? | 6 semitones; e.g. C to F# | wiki |
| 15 | What scale is the relative major of A minor pentatonic? | C major pentatonic | wiki |
| 16 | What are the diatonic chords in G major? | G, Am, Bm, C, D, Em, F#dim | wiki |
| 17 | What's a good recipe for banana bread? | canned "I don't have enough context"; `sources: []` | off-corpus |
| 18 | How do I set up a Python virtual environment? | canned "I don't have enough context"; `sources: []` | off-corpus |

Rubric (all rows): correctness against expected points; groundedness in the cited
chunks; off-corpus rows pass only on the exact canned string.

## 8. Doc updates (same change as this spec's implementation)

- `docs/design.md`: §3 corpus source + licenses paragraph; §4.1 layout (`musicwiki/`,
  `guides/`, `internal/corpus/`; drop `corpus/`); §4.3 add `CORPUS_DIRS`; §5.1 ingestion
  input, per-source reconcile, empty-source guard; §9.5 demo step 6 uses a guide; §14
  drop the Open Music Theory reference.
- `CLAUDE.md`: Guardrails licensing bullet; "Ingestion is idempotent" paragraph;
  layout list.
- `agent.md`: any licensing / corpus references.
- `docs/plan.md` step 7: seed content is `guides/` + existing `musicwiki/`.
- `.env.example`: `CORPUS_DIRS`.

## 9. Work order

1. Doc updates (§8) and `evals/golden.jsonl`.
2. Eight guide drafts + checker report → owner review.
3. `internal/corpus` (TDD per §5.2), `CORPUS_DIRS` config, `cmd/ingest` wiring
   (Milestone 2, plan step 7).
4. First real ingest: measure score floor (§7), record in design doc.

## 10. Acceptance

- `go run ./cmd/ingest` twice ⇒ identical row counts, zero duplicates.
- Removing a guide file and re-ingesting deletes that guide's document and chunks and
  nothing else; wiki row count unchanged.
- `CORPUS_DIRS=musicwiki,gudies` (typo) ⇒ non-zero exit, no rows deleted.
- Flattener golden-text tests pass for all five node types and a guide.
- `go vet`, `staticcheck`, `go test ./...`, `gofmt -l .` clean.
