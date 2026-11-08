---
type: scale
slug: /scales/c-sharp-dorian-scale
title: C# Dorian Scale
description: The dorian scale built on C#.
tags:
- scale
- dorian
- c
parent: /scales
prerequisites:
- /scales/b-major-scale
interval_formula:
- 2
- 1
- 2
- 2
- 2
- 1
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
    pc: 1
    name: C#
  intervals:
  - '1'
  - '2'
  - b3
  - '4'
  - '5'
  - '6'
  - b7
  semitones:
  - 0
  - 2
  - 3
  - 5
  - 7
  - 9
  - 10
  pitch_classes:
  - 1
  - 3
  - 4
  - 6
  - 8
  - 10
  - 11
  spelling:
  - C#
  - D#
  - E
  - F#
  - G#
  - A#
  - B
  enharmonic_policy: prefer_sharps
  key_context: C#
musical_forms:
- form: pitch_set
  id: c-sharp-dorian-scale-pitch-set
  pitch_classes:
  - 1
  - 3
  - 4
  - 6
  - 8
  - 10
  - 11
  intervals:
  - '1'
  - '2'
  - b3
  - '4'
  - '5'
  - '6'
  - b7
  spelling:
  - C#
  - D#
  - E
  - F#
  - G#
  - A#
  - B
- form: pattern
  id: c-sharp-dorian-scale-pattern-1
  system: 3-note-per-string
  degree: 1
  start_fret: 9
  end_fret: 14
  notes:
  - string: 6
    fret: 9
    degree: 1
    interval: '1'
    note: C#
    pc: 1
    is_root: true
  - string: 6
    fret: 11
    degree: 2
    interval: '2'
    note: D#
    pc: 3
    is_root: false
  - string: 6
    fret: 12
    degree: 3
    interval: b3
    note: E
    pc: 4
    is_root: false
  - string: 5
    fret: 9
    degree: 4
    interval: '4'
    note: F#
    pc: 6
    is_root: false
  - string: 5
    fret: 11
    degree: 5
    interval: '5'
    note: G#
    pc: 8
    is_root: false
  - string: 5
    fret: 13
    degree: 6
    interval: '6'
    note: A#
    pc: 10
    is_root: false
  - string: 4
    fret: 9
    degree: 7
    interval: b7
    note: B
    pc: 11
    is_root: false
  - string: 4
    fret: 11
    degree: 1
    interval: '1'
    note: C#
    pc: 1
    is_root: true
  - string: 4
    fret: 13
    degree: 2
    interval: '2'
    note: D#
    pc: 3
    is_root: false
  - string: 3
    fret: 9
    degree: 3
    interval: b3
    note: E
    pc: 4
    is_root: false
  - string: 3
    fret: 11
    degree: 4
    interval: '4'
    note: F#
    pc: 6
    is_root: false
  - string: 3
    fret: 13
    degree: 5
    interval: '5'
    note: G#
    pc: 8
    is_root: false
  - string: 2
    fret: 11
    degree: 6
    interval: '6'
    note: A#
    pc: 10
    is_root: false
  - string: 2
    fret: 12
    degree: 7
    interval: b7
    note: B
    pc: 11
    is_root: false
  - string: 2
    fret: 14
    degree: 1
    interval: '1'
    note: C#
    pc: 1
    is_root: true
  - string: 1
    fret: 11
    degree: 2
    interval: '2'
    note: D#
    pc: 3
    is_root: false
  - string: 1
    fret: 12
    degree: 3
    interval: b3
    note: E
    pc: 4
    is_root: false
  - string: 1
    fret: 14
    degree: 4
    interval: '4'
    note: F#
    pc: 6
    is_root: false
- form: pattern
  id: c-sharp-dorian-scale-pattern-2
  system: 3-note-per-string
  degree: 2
  start_fret: 11
  end_fret: 16
  notes:
  - string: 6
    fret: 11
    degree: 2
    interval: '2'
    note: D#
    pc: 3
    is_root: false
  - string: 6
    fret: 12
    degree: 3
    interval: b3
    note: E
    pc: 4
    is_root: false
  - string: 6
    fret: 14
    degree: 4
    interval: '4'
    note: F#
    pc: 6
    is_root: false
  - string: 5
    fret: 11
    degree: 5
    interval: '5'
    note: G#
    pc: 8
    is_root: false
  - string: 5
    fret: 13
    degree: 6
    interval: '6'
    note: A#
    pc: 10
    is_root: false
  - string: 5
    fret: 14
    degree: 7
    interval: b7
    note: B
    pc: 11
    is_root: false
  - string: 4
    fret: 11
    degree: 1
    interval: '1'
    note: C#
    pc: 1
    is_root: true
  - string: 4
    fret: 13
    degree: 2
    interval: '2'
    note: D#
    pc: 3
    is_root: false
  - string: 4
    fret: 14
    degree: 3
    interval: b3
    note: E
    pc: 4
    is_root: false
  - string: 3
    fret: 11
    degree: 4
    interval: '4'
    note: F#
    pc: 6
    is_root: false
  - string: 3
    fret: 13
    degree: 5
    interval: '5'
    note: G#
    pc: 8
    is_root: false
  - string: 3
    fret: 15
    degree: 6
    interval: '6'
    note: A#
    pc: 10
    is_root: false
  - string: 2
    fret: 12
    degree: 7
    interval: b7
    note: B
    pc: 11
    is_root: false
  - string: 2
    fret: 14
    degree: 1
    interval: '1'
    note: C#
    pc: 1
    is_root: true
  - string: 2
    fret: 16
    degree: 2
    interval: '2'
    note: D#
    pc: 3
    is_root: false
  - string: 1
    fret: 12
    degree: 3
    interval: b3
    note: E
    pc: 4
    is_root: false
  - string: 1
    fret: 14
    degree: 4
    interval: '4'
    note: F#
    pc: 6
    is_root: false
  - string: 1
    fret: 16
    degree: 5
    interval: '5'
    note: G#
    pc: 8
    is_root: false
- form: pattern
  id: c-sharp-dorian-scale-pattern-3
  system: 3-note-per-string
  degree: 3
  start_fret: 0
  end_fret: 6
  notes:
  - string: 6
    fret: 0
    degree: 3
    interval: b3
    note: E
    pc: 4
    is_root: false
  - string: 6
    fret: 2
    degree: 4
    interval: '4'
    note: F#
    pc: 6
    is_root: false
  - string: 6
    fret: 4
    degree: 5
    interval: '5'
    note: G#
    pc: 8
    is_root: false
  - string: 5
    fret: 1
    degree: 6
    interval: '6'
    note: A#
    pc: 10
    is_root: false
  - string: 5
    fret: 2
    degree: 7
    interval: b7
    note: B
    pc: 11
    is_root: false
  - string: 5
    fret: 4
    degree: 1
    interval: '1'
    note: C#
    pc: 1
    is_root: true
  - string: 4
    fret: 1
    degree: 2
    interval: '2'
    note: D#
    pc: 3
    is_root: false
  - string: 4
    fret: 2
    degree: 3
    interval: b3
    note: E
    pc: 4
    is_root: false
  - string: 4
    fret: 4
    degree: 4
    interval: '4'
    note: F#
    pc: 6
    is_root: false
  - string: 3
    fret: 1
    degree: 5
    interval: '5'
    note: G#
    pc: 8
    is_root: false
  - string: 3
    fret: 3
    degree: 6
    interval: '6'
    note: A#
    pc: 10
    is_root: false
  - string: 3
    fret: 4
    degree: 7
    interval: b7
    note: B
    pc: 11
    is_root: false
  - string: 2
    fret: 2
    degree: 1
    interval: '1'
    note: C#
    pc: 1
    is_root: true
  - string: 2
    fret: 4
    degree: 2
    interval: '2'
    note: D#
    pc: 3
    is_root: false
  - string: 2
    fret: 5
    degree: 3
    interval: b3
    note: E
    pc: 4
    is_root: false
  - string: 1
    fret: 2
    degree: 4
    interval: '4'
    note: F#
    pc: 6
    is_root: false
  - string: 1
    fret: 4
    degree: 5
    interval: '5'
    note: G#
    pc: 8
    is_root: false
  - string: 1
    fret: 6
    degree: 6
    interval: '6'
    note: A#
    pc: 10
    is_root: false
---

# C# Dorian Scale

The C# dorian scale is a minor mode with a raised sixth — a staple of jazz, folk, and modal rock. It is spelled **C# – D# – E – F# – G# – A# – B**.

**Step pattern:** W–H–W–W–W–H–W  *(W = whole, H = half, W½ = one-and-a-half steps)*

| Degree | Note | Interval | Function |
|---|---|---|---|
| 1 | C# | 1 | tonic |
| 2 | D# | 2 | supertonic |
| 3 | E | b3 | mediant |
| 4 | F# | 4 | subdominant |
| 5 | G# | 5 | dominant |
| 6 | A# | 6 | submediant |
| 7 | B | b7 | subtonic |

## Related

- **Build on:** [B Major Scale](/scales/b-major-scale.md)
