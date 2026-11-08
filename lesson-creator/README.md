# Guitar Lesson-Building Harness

A pipeline that takes a prompt like *"create a lesson on how to use the minor pentatonic scale and the various positions"* and produces a complete, validated lesson in a fixed JSON format, with deterministic SVG chord/note diagrams, practice steps, tips, and related-topic links.

This Project preserves the original design conversation plus a working MVP implementation.

## Original design (preserved)

The core principle: **the LLM never freehands musical facts.** Facts (notes, frets, roots) come from a deterministic theory engine and independently checked reference fixtures; the LLM handles pedagogy, explanations, and practice design.

The pipeline is a single sequential flow, not a multi-agent system:

```
prompt → brief → musical spec → content → deterministic diagrams → validation → preview
```

1. **Lesson schema (lesson-schema.json)** — versioned contract every lesson satisfies. Includes audience_level enum, diagram refs, topic IDs.
2. **Prompt → explicit brief** — resolve scope (key, tuning, level, position convention) with visible defaults.
3. **Canonical musical specification** — structured note/fret data from the theory engine.
4. **Instructional content** — sections, practice steps with goal/action/duration/success criterion, tips.
5. **Deterministic diagram rendering** — SVG from structured data.
6. **Validation** — hard gates (schema, note mapping, fixture comparison, ref resolution, topic catalog) vs. quality flags.
7. **Preview** — render and inspect the actual lesson page.

## Files

| File | Purpose |
|------|---------|
| `lesson-schema.json` | JSON Schema (draft 2020-12) for the lesson contract |
| `theory.py` | Music theory engine: chromatic notes, standard tuning, scale computation, position shapes transposed to a root |
| `fixtures.py` | Reference shapes for the 5-position minor pentatonic system, verified against [MuseDrill's pentatonic scale guide](https://musedrill.com/eng-pentatonic-scale) (independent source) |
| `validate.py` | JSON Schema enforcement via `jsonschema` + semantic checks against fixtures + topic catalog |
| `render.py` | Deterministic SVG diagram renderer (string 6 leftmost / 1 rightmost, conventional numbering) |
| `generate.py` | Takes a prompt, emits a lesson JSON following the schema. Currently restricted to A minor pentatonic |
| `preview.py` | Renders lesson JSON as readable Markdown for human review |
| `test_negative.py` | 10 mutation tests proving the validator catches real errors (the missing-destination case runs separately in main) |
| `catalog.json` | Demo related-topic catalog |
| `topics/` | Stub Markdown pages for each catalog entry (real destinations, not just IDs) |
| `lessons/` | Generated lesson JSON + rendered diagrams + Markdown preview |

## Commands

```bash
pip install jsonschema cairosvg   # deps

# generate the A minor pentatonic prototype
python generate.py "create a lesson on how to use the minor pentatonic scale and the various positions"

# validate against schema + fixtures + catalog
python validate.py lessons/a-minor-pentatonic-positions.json

# render SVG diagrams
python render.py lessons/a-minor-pentatonic-positions.json

# render readable Markdown preview
python preview.py lessons/a-minor-pentatonic-positions.json

# run negative mutation tests
python test_negative.py

# rasterize for preview (CairoSVG; ImageMagick's convert mangles these SVGs)
python -c "import cairosvg, pathlib; [cairosvg.svg2png(url=str(f), write_to=str(f.with_suffix('.png'))) for f in pathlib.Path('lessons/diagrams').glob('*.svg')]"
```

## Verification

- **Positive test:** `validate.py` passes against the A minor pentatonic prototype.
- **Negative tests (10/10 caught by the validator):**
  - Root flag flipped to False → caught by fixture comparison
  - Extra out-of-scale note → caught by pitch-class check
  - Note removed → caught by fixture set-difference
  - Related topic ID not in catalog → caught by catalog lookup
  - Broken diagram ref → caught by ref resolution
  - Invalid audience enum → caught by JSON Schema
  - Declared fret bounds exclude notes (low=100, high=103) → caught by bounds check
  - Declared fret range reversed (low > high) → caught by bounds check
  - Unsupported root (E minor) → caught by scope restriction
  - Related topic with a catalog entry but missing destination file → caught by destination-file check (validated for the specific missing-destination error, not just any failure)
- **Fixture verification:** all five positions verified against an independent published source ([MuseDrill](https://musedrill.com/eng-pentatonic-scale)) — 60/60 notes match, 0 mismatches.
- **Related topics:** each ID has a real destination file in `topics/`; the validator rejects IDs with no file.

## Scope

This prototype supports **A minor pentatonic, standard tuning, 5-position system only.** The generator and validator both explicitly reject other roots and scales until per-root fixtures are added.

**Prompt-parser limitation:** the `in <key>` regex only matches lowercase note letters after the word "in," so a request like "D minor pentatonic" is silently treated as the default A minor rather than rejected or detected. Treat this as a prototype for the supported default-A prompt (the canonical example above), not a general-purpose generator. Detecting arbitrary keys requires a more robust parser or an explicit key parameter.

## Known limitations / next steps

- Only minor pentatonic + 5-position system is implemented. Add major pentatonic, blues scale, triads, chords — each requires hand-checked fixtures.
- Prose is template-based. Next step: pass the canonical musical spec to an LLM for pedagogical content, keeping diagrams/validation deterministic.
- `catalog.json` and `topics/` are demos; the real catalog should live in the app's database with real routes.
- ImageMagick's `convert` rasterizes these SVGs incorrectly (drops white-filled circles and fragments the grid). Use CairoSVG (or the browser) for previews.
