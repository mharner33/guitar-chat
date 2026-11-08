---
type: scale
slug: /scales/e-locrian-scale
title: E Locrian Scale
description: The locrian scale built on E.
tags:
- scale
- locrian
- e
parent: /scales
prerequisites:
- /scales/f-major-scale
interval_formula:
- 1
- 2
- 2
- 1
- 2
- 2
- 2
degree_functions:
- tonic
- supertonic
- mediant
- subdominant
- dominant
- submediant
- subtonic
theory:
  category: scale
  root:
    pc: 4
    name: E
  intervals:
  - '1'
  - b2
  - b3
  - '4'
  - b5
  - b6
  - b7
  semitones:
  - 0
  - 1
  - 3
  - 5
  - 6
  - 8
  - 10
  pitch_classes:
  - 4
  - 5
  - 7
  - 9
  - 10
  - 0
  - 2
  spelling:
  - E
  - F
  - G
  - A
  - Bb
  - C
  - D
  enharmonic_policy: prefer_flats
  key_context: E
musical_forms:
- form: pitch_set
  id: e-locrian-scale-pitch-set
  pitch_classes:
  - 4
  - 5
  - 7
  - 9
  - 10
  - 0
  - 2
  intervals:
  - '1'
  - b2
  - b3
  - '4'
  - b5
  - b6
  - b7
  spelling:
  - E
  - F
  - G
  - A
  - Bb
  - C
  - D
- form: pattern
  id: e-locrian-scale-pattern-1
  system: 3-note-per-string
  degree: 1
  start_fret: 0
  end_fret: 5
  notes:
  - string: 6
    fret: 0
    degree: 1
    interval: '1'
    note: E
    pc: 4
    is_root: true
  - string: 6
    fret: 1
    degree: 2
    interval: b2
    note: F
    pc: 5
    is_root: false
  - string: 6
    fret: 3
    degree: 3
    interval: b3
    note: G
    pc: 7
    is_root: false
  - string: 5
    fret: 0
    degree: 4
    interval: '4'
    note: A
    pc: 9
    is_root: false
  - string: 5
    fret: 1
    degree: 5
    interval: b5
    note: Bb
    pc: 10
    is_root: false
  - string: 5
    fret: 3
    degree: 6
    interval: b6
    note: C
    pc: 0
    is_root: false
  - string: 4
    fret: 0
    degree: 7
    interval: b7
    note: D
    pc: 2
    is_root: false
  - string: 4
    fret: 2
    degree: 1
    interval: '1'
    note: E
    pc: 4
    is_root: true
  - string: 4
    fret: 3
    degree: 2
    interval: b2
    note: F
    pc: 5
    is_root: false
  - string: 3
    fret: 0
    degree: 3
    interval: b3
    note: G
    pc: 7
    is_root: false
  - string: 3
    fret: 2
    degree: 4
    interval: '4'
    note: A
    pc: 9
    is_root: false
  - string: 3
    fret: 3
    degree: 5
    interval: b5
    note: Bb
    pc: 10
    is_root: false
  - string: 2
    fret: 1
    degree: 6
    interval: b6
    note: C
    pc: 0
    is_root: false
  - string: 2
    fret: 3
    degree: 7
    interval: b7
    note: D
    pc: 2
    is_root: false
  - string: 2
    fret: 5
    degree: 1
    interval: '1'
    note: E
    pc: 4
    is_root: true
  - string: 1
    fret: 1
    degree: 2
    interval: b2
    note: F
    pc: 5
    is_root: false
  - string: 1
    fret: 3
    degree: 3
    interval: b3
    note: G
    pc: 7
    is_root: false
  - string: 1
    fret: 5
    degree: 4
    interval: '4'
    note: A
    pc: 9
    is_root: false
- form: pattern
  id: e-locrian-scale-pattern-2
  system: 3-note-per-string
  degree: 2
  start_fret: 1
  end_fret: 6
  notes:
  - string: 6
    fret: 1
    degree: 2
    interval: b2
    note: F
    pc: 5
    is_root: false
  - string: 6
    fret: 3
    degree: 3
    interval: b3
    note: G
    pc: 7
    is_root: false
  - string: 6
    fret: 5
    degree: 4
    interval: '4'
    note: A
    pc: 9
    is_root: false
  - string: 5
    fret: 1
    degree: 5
    interval: b5
    note: Bb
    pc: 10
    is_root: false
  - string: 5
    fret: 3
    degree: 6
    interval: b6
    note: C
    pc: 0
    is_root: false
  - string: 5
    fret: 5
    degree: 7
    interval: b7
    note: D
    pc: 2
    is_root: false
  - string: 4
    fret: 2
    degree: 1
    interval: '1'
    note: E
    pc: 4
    is_root: true
  - string: 4
    fret: 3
    degree: 2
    interval: b2
    note: F
    pc: 5
    is_root: false
  - string: 4
    fret: 5
    degree: 3
    interval: b3
    note: G
    pc: 7
    is_root: false
  - string: 3
    fret: 2
    degree: 4
    interval: '4'
    note: A
    pc: 9
    is_root: false
  - string: 3
    fret: 3
    degree: 5
    interval: b5
    note: Bb
    pc: 10
    is_root: false
  - string: 3
    fret: 5
    degree: 6
    interval: b6
    note: C
    pc: 0
    is_root: false
  - string: 2
    fret: 3
    degree: 7
    interval: b7
    note: D
    pc: 2
    is_root: false
  - string: 2
    fret: 5
    degree: 1
    interval: '1'
    note: E
    pc: 4
    is_root: true
  - string: 2
    fret: 6
    degree: 2
    interval: b2
    note: F
    pc: 5
    is_root: false
  - string: 1
    fret: 3
    degree: 3
    interval: b3
    note: G
    pc: 7
    is_root: false
  - string: 1
    fret: 5
    degree: 4
    interval: '4'
    note: A
    pc: 9
    is_root: false
  - string: 1
    fret: 6
    degree: 5
    interval: b5
    note: Bb
    pc: 10
    is_root: false
- form: pattern
  id: e-locrian-scale-pattern-3
  system: 3-note-per-string
  degree: 3
  start_fret: 3
  end_fret: 8
  notes:
  - string: 6
    fret: 3
    degree: 3
    interval: b3
    note: G
    pc: 7
    is_root: false
  - string: 6
    fret: 5
    degree: 4
    interval: '4'
    note: A
    pc: 9
    is_root: false
  - string: 6
    fret: 6
    degree: 5
    interval: b5
    note: Bb
    pc: 10
    is_root: false
  - string: 5
    fret: 3
    degree: 6
    interval: b6
    note: C
    pc: 0
    is_root: false
  - string: 5
    fret: 5
    degree: 7
    interval: b7
    note: D
    pc: 2
    is_root: false
  - string: 5
    fret: 7
    degree: 1
    interval: '1'
    note: E
    pc: 4
    is_root: true
  - string: 4
    fret: 3
    degree: 2
    interval: b2
    note: F
    pc: 5
    is_root: false
  - string: 4
    fret: 5
    degree: 3
    interval: b3
    note: G
    pc: 7
    is_root: false
  - string: 4
    fret: 7
    degree: 4
    interval: '4'
    note: A
    pc: 9
    is_root: false
  - string: 3
    fret: 3
    degree: 5
    interval: b5
    note: Bb
    pc: 10
    is_root: false
  - string: 3
    fret: 5
    degree: 6
    interval: b6
    note: C
    pc: 0
    is_root: false
  - string: 3
    fret: 7
    degree: 7
    interval: b7
    note: D
    pc: 2
    is_root: false
  - string: 2
    fret: 5
    degree: 1
    interval: '1'
    note: E
    pc: 4
    is_root: true
  - string: 2
    fret: 6
    degree: 2
    interval: b2
    note: F
    pc: 5
    is_root: false
  - string: 2
    fret: 8
    degree: 3
    interval: b3
    note: G
    pc: 7
    is_root: false
  - string: 1
    fret: 5
    degree: 4
    interval: '4'
    note: A
    pc: 9
    is_root: false
  - string: 1
    fret: 6
    degree: 5
    interval: b5
    note: Bb
    pc: 10
    is_root: false
  - string: 1
    fret: 8
    degree: 6
    interval: b6
    note: C
    pc: 0
    is_root: false
---

# E Locrian Scale

The E locrian scale is the darkest mode, built on a diminished tonic — rarely used as a tonal center. It is spelled **E – F – G – A – Bb – C – D**.

**Step pattern:** H–W–W–H–W–W–W  *(W = whole, H = half, W½ = one-and-a-half steps)*

| Degree | Note | Interval | Function |
|---|---|---|---|
| 1 | E | 1 | tonic |
| 2 | F | b2 | supertonic |
| 3 | G | b3 | mediant |
| 4 | A | 4 | subdominant |
| 5 | Bb | b5 | dominant |
| 6 | C | b6 | submediant |
| 7 | D | b7 | subtonic |

## Related

- **Build on:** [F Major Scale](/scales/f-major-scale.md)
