---
type: Concept
slug: /chords/g-dominant-7th-chord
title: G Dominant 7th Chord
description: 'The dominant 7th chord built on G: G, B, D, F.'
tags:
- chord
- dominant-7th
- g
parent: /chords
prerequisites:
- /triads/g-major-triad
used_by:
- /chords/g-9th-chord
- /chords/g-11th-chord
- /chords/g-13th-chord
theory:
  category: chord
  quality: dominant-7th
  root:
    pc: 7
    name: G
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  semitones:
  - 0
  - 4
  - 7
  - 10
  pitch_classes:
  - 7
  - 11
  - 2
  - 5
  spelling:
  - G
  - B
  - D
  - F
  enharmonic_policy: prefer_sharps
  key_context: G
instrument:
  instrument_type: guitar
  tuning:
  - string: 6
    open_pc: 4
    open_name: E2
  - string: 5
    open_pc: 9
    open_name: A2
  - string: 4
    open_pc: 2
    open_name: D3
  - string: 3
    open_pc: 7
    open_name: G3
  - string: 2
    open_pc: 11
    open_name: B3
  - string: 1
    open_pc: 4
    open_name: E4
  capo: 0
render:
  display_mode: chord_tones
  highlight_pitch_classes:
  - 7
  - 11
  - 2
  - 5
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - b7
  preferred_positions:
  - id: g-dominant-7th-voicing-1
    start_fret: 2
    end_fret: 3
    notes:
    - string: 6
      fret: 3
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 2
      pc: 11
      note: B
      interval: '3'
      role: third
      finger: 1
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: b7
      role: seventh
      finger: 2
    - string: 3
      fret: 0
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 0
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '5'
      role: fifth
      finger: 2
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: g-dominant-7th-voicing-2
    start_fret: 1
    end_fret: 3
    notes:
    - string: 6
      fret: 3
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 2
      pc: 11
      note: B
      interval: '3'
      role: third
      finger: 2
    - string: 4
      fret: 0
      pc: 2
      note: D
      interval: '5'
      role: fifth
      finger: 0
    - string: 3
      fret: 0
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 0
    - string: 2
      fret: 0
      pc: 11
      note: B
      interval: '3'
      role: third
      finger: 0
    - string: 1
      fret: 1
      pc: 5
      note: F
      interval: b7
      role: seventh
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - open
  - id: g-dominant-7th-voicing-3
    start_fret: 3
    end_fret: 5
    notes:
    - string: 6
      fret: 3
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: '5'
      role: fifth
      finger: 3
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: b7
      role: seventh
      finger: 1
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '3'
      role: third
      finger: 2
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '5'
      role: fifth
      finger: 1
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (3.6 2.5 3.4 0.3 3.2 3.1).2
  midi_root: 67
musical_forms:
- form: pitch_set
  id: g-dominant-7th-chord-pitch-set
  pitch_classes:
  - 7
  - 11
  - 2
  - 5
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  spelling:
  - G
  - B
  - D
  - F
- form: voicing
  id: g-dominant-7th-voicing-1
  position_ref: g-dominant-7th-voicing-1
  start_fret: 2
  tab: 3-2-3-0-3-3
- form: voicing
  id: g-dominant-7th-voicing-2
  position_ref: g-dominant-7th-voicing-2
  start_fret: 1
  tab: 3-2-0-0-0-1
- form: voicing
  id: g-dominant-7th-voicing-3
  position_ref: g-dominant-7th-voicing-3
  start_fret: 3
  tab: 3-5-3-4-3-3
- form: tab_excerpt
  id: g-dominant-7th-chord-tab
  alphatex: \tempo 90 (3.6 2.5 3.4 0.3 3.2 3.1).2
---

# G Dominant 7th Chord

The G dominant 7th chord is a major triad plus a minor seventh — the restless engine of tonal harmony, pulling toward the tonic a fourth above. Its tones are **G – B – D – F**: G (root), B (major third), D (perfect fifth), F (minor seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | G | root |
| 3 | B | major third |
| 5 | D | perfect fifth |
| b7 | F | minor seventh |

**Playing it:** the primary shape is `3-2-3-0-3-3` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [G Major Triad](/triads/g-major-triad.md)
- **Used by:** [G 9th Chord](/chords/g-9th-chord.md), [G 11th Chord](/chords/g-11th-chord.md), [G 13th Chord](/chords/g-13th-chord.md)
