# musicwiki

A generator for a guitar-focused **music-theory wiki**, emitted as an
[Open Knowledge Format (OKF)](https://github.com/GoogleCloudPlatform/knowledge-catalog/blob/main/okf/SPEC.md)
bundle. It produces a tree of self-contained Markdown files — one per concept
(scale, triad, or chord) — that a separate application (e.g. an AI-powered
lesson generator) consumes as its content and data layer.

This repository is the **generator**. The wiki it emits is a **separate bundle**
that your application reads. This README explains how to generate that bundle and
integrate it into an application.

---

## What you get

Running the generator writes an OKF bundle (default target `~/Apps/musicwiki/`):

```
musicwiki/
├── index.md          # bundle root listing (Intervals / Scales / Triads / Chords / Keys)
├── log.md            # changelog: one dated entry per generation run
├── intervals/        # 13 foundational interval nodes (unison … octave)
├── scales/           # 12 roots × 12 scales (9 modes + major/minor pentatonic + blues)
│   ├── index.md
│   ├── c-major-scale.md
│   ├── c-major-pentatonic-scale.md
│   ├── c-blues-scale.md
│   └── …
├── triads/           # 12 roots × 4 qualities
├── chords/           # 12 roots × 15 seventh/extended/suspended chords
└── keys/             # 24 keys (major + natural minor): diatonic chords & function
    ├── c-major.md
    ├── a-natural-minor.md
    └── …
```

Every concept file is UTF-8 Markdown with a YAML frontmatter block (the
structured data) followed by a prose body (the human/LLM-readable lesson).

## Generating the bundle

Requires Python 3.12+ and [`uv`](https://docs.astral.sh/uv/).

```bash
uv sync                                   # install dependencies
uv run musicwiki --corpus -o ./wiki       # generate the full corpus to ./wiki
uv run musicwiki --corpus --no-voicings   # faster: theory only, skip fretboard voicings
```

Regenerating into the same directory is safe and idempotent: node files are
overwritten in place, and each run **prepends** a dated entry to `log.md` rather
than clobbering history. **Slugs are stable across runs**, so database references
(below) keep pointing at the right concept.

---

## Integrating with an application

### 1. Reference concepts by `slug`

Every node has a `slug` — a bundle-relative path that mirrors the file location
(minus `.md`) and is the node's stable identifier:

```yaml
slug: /triads/c-major-triad
```

Store this slug as the foreign key in your database. To resolve a slug to its
file, strip the leading `/` and append `.md`:
`/triads/c-major-triad` → `<bundle>/triads/c-major-triad.md`.

The only frontmatter field OKF strictly requires is `type`; `slug` is the field
your integration should key on.

### 2. Load and parse a node

A node is plain Markdown-with-frontmatter — read it with any YAML + Markdown
tooling in any language. In Python:

```python
import yaml, pathlib

raw = pathlib.Path("wiki/triads/c-major-triad.md").read_text()
front, _, body = raw.partition("\n---\n")
data = yaml.safe_load(front.removeprefix("---\n"))   # structured fields
lesson_markdown = body                                # prose body
```

### 3. Node anatomy

Abbreviated `triads/c-major-triad.md`:

```yaml
type: Concept
slug: /triads/c-major-triad
title: C Major Triad
description: "The major triad built on C: C, E, G."
tags: [triad, major, c]
parent: /triads

# --- knowledge-graph edges (values are slugs of other nodes) ---
prerequisites: [/scales/c-major-scale]
used_by: [/chords/c-dominant-7th-chord, /chords/c-major-7th-chord, …]

# --- abstract theory (instrument-independent) ---
theory:
  category: triad
  quality: major
  root: {pc: 0, name: C}
  intervals: ["1", "3", "5"]
  semitones: [0, 4, 7]           # offsets from the root
  pitch_classes: [0, 4, 7]       # absolute pitch classes (0=C … 11=B)
  spelling: [C, E, G]            # letter-correct note names
  enharmonic_policy: prefer_sharps
  key_context: C

# --- concrete guitar realization ---
instrument:
  instrument_type: guitar
  tuning: [{string: 6, open_pc: 4, open_name: E2}, …]   # 1 = high E … 6 = low E
  capo: 0
render:
  display_mode: chord_tones
  highlight_pitch_classes: [0, 4, 7]   # for the "global map" view
  emphasize_intervals: ["1", "3", "5"]
  preferred_positions:                 # computed, ranked voicings
    - id: c-major-voicing-1
      start_fret: 1
      end_fret: 3
      notes:
        - {string: 5, fret: 3, pc: 0, note: C, interval: "1", role: root, finger: 3}
        - {string: 4, fret: 2, pc: 4, note: E, interval: "3", role: third, finger: 2}
        - …
      mute_strings: [6]
      confidence: 0.75
      tags: [computed, open]

# --- playback ---
playback:
  alphatex: "\\tempo 90 (3.5 2.4 0.3 1.2 0.1).2"   # AlphaTab (fret.string)
  midi_root: 60                                     # C4

# --- multiple musical payload types ---
musical_forms:
  - {form: pitch_set, id: c-major-triad-pitch-set, pitch_classes: [0, 4, 7], …}
  - {form: voicing, id: c-major-voicing-1, position_ref: c-major-voicing-1, tab: "x-3-2-0-1-0", …}
  - {form: tab_excerpt, id: c-major-triad-tab, alphatex: "\\tempo 90 (3.5 2.4 0.3 1.2 0.1).2"}
```

Scale nodes use `type: scale`, add `interval_formula` / `degree_functions`, and
their `musical_forms` include movable `pattern` boxes instead of voicings; they
omit `render`/`playback`/`instrument` (use `highlight_pitch_classes` from a
pattern or the pitch set to build neck maps).

### 4. Render the three granularities

The design stores **abstract pitch-class targets** and **concrete realizations**
in the same node, so you can render any of three views from one concept:

| View | Use for | Source fields |
|---|---|---|
| **Global neck map** | note-learning, scale-tone overlays across the whole neck | `render.highlight_pitch_classes` (map every fret whose pitch class matches) |
| **Pattern view** | boxed shapes — CAGED-style / 3-note-per-string scale boxes | scale `musical_forms` of `form: pattern` (each has `notes[]` with `string`/`fret`/`degree`/`is_root`) |
| **Voicing view** | a single concrete fingering | `render.preferred_positions[]` (per-string `fret`/`finger`/`role`, plus `mute_strings`) — or `musical_forms` of `form: voicing` |

Every fretted note carries `string` (1 = high E … 6 = low E), `fret`, `pc`,
`note` name, `interval`, `role`, and `finger` (0 = open). Voicings are ranked;
`preferred_positions[0]` is the best/primary shape, and each has a `confidence`
score and descriptive `tags` (`open`, `barre`, `computed`).

### 5. Play audio

- **Notation / playback:** `playback.alphatex` is [AlphaTab](https://alphatab.net)
  **alphaTex** using `fret.string` note syntax (string numbering matches the
  nodes: 1 = high E … 6 = low E). Feed it directly to an AlphaTab renderer/player.
- **Synth:** `playback.midi_root` is the root as a MIDI note number referenced to
  octave 4 (C = 60) — a register-independent root reference. For exact voicing
  pitches, compute per note: `midi = open_pc_of_string_in_octave + fret` using
  `instrument.tuning[].open_name` (e.g. `E2`, `A2`).

### 6. Traverse the knowledge graph

Relationship fields hold **slugs of other nodes**, forming a navigable graph your
app can walk (build a lesson path, suggest "next", show related concepts):

| Field | Meaning |
|---|---|
| `parent` | containing category (`/scales`, `/triads`, `/chords`, `/intervals`, `/keys`) |
| `relative_major` / `relative_minor` | scale ⇄ its relative key (and key ⇄ relative key) |
| `prerequisites` | concepts to learn first (mode → parent major scale; triad → its scale; 7th chord → its triad; key → its tonic scale) |
| `used_by` | concepts that build on this one (scale → its modes, tonic triad & key; triad → chords that extend it) |

**Key nodes** (`/keys/`) harmonize a scale into its diatonic chords: a `progression`
musical form carries the structured harmony (per-degree Roman numeral, root,
quality, harmonic function, and the slug of the actual triad/seventh chord node),
and the body renders the same as a Roman-numeral table. This is the bridge between
`/scales/` and `/triads/`/`/chords/`. **Interval nodes** (`/intervals/`) are the
foundational building blocks; their bodies link to their inversion.

Frontmatter relationship values are slugs **without** `.md` (they equal target
nodes' `slug` fields — match on those directly). The prose body additionally
contains standard Markdown links (`[Title](/scales/c-major-scale.md)`), which are
OKF's native way of asserting relationships if you prefer link-graph parsing.

`index.md` files give per-directory listings, and `log.md` records generation
history — both are reserved OKF filenames and are not concept nodes (no `slug`).

### 7. Drive an LLM lesson generator

Each node is designed to be handed to a language model as-is: the **prose body**
is a readable lesson (description, degree/chord-tone table, construction notes,
how to play the primary shape, and a Related section), and the **frontmatter** is
the exact structured data (pitch classes, intervals, fingerings, playback) the
model needs to stay accurate. A typical flow:

1. Your app maps a lesson request to one or more node slugs.
2. Load those nodes; follow `prerequisites` to assemble background context.
3. Pass the frontmatter (facts) + body (explanation) to the model as grounding.
4. Render the fretboard/audio from `render`, `musical_forms`, and `playback`.

---

## Programmatic use (Python)

If your application is in Python, you can build nodes directly instead of reading
files — the builders return the same structures the corpus writes:

```python
from musicwiki import build_triad, build_chord, build_scale, attach_render

triad = attach_render(build_triad("C", "major"))   # theory + voicings + playback + forms
chord = attach_render(build_chord("G", "dominant-7th"))
scale = build_scale("A", "natural_minor")
```

See `CLAUDE.md` for the generator's internal architecture and module reference.

## License

See repository.
