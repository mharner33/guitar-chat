---
type: Concept
slug: /chords/a-flat-6th-chord
title: Ab 6th Chord
description: 'The 6th chord built on Ab: Ab, C, Eb, F.'
tags:
- chord
- 6th
- a
parent: /chords
prerequisites:
- /triads/a-flat-major-triad
theory:
  category: chord
  quality: 6th
  root:
    pc: 8
    name: Ab
  intervals:
  - '1'
  - '3'
  - '5'
  - '6'
  semitones:
  - 0
  - 4
  - 7
  - 9
  pitch_classes:
  - 8
  - 0
  - 3
  - 5
  spelling:
  - Ab
  - C
  - Eb
  - F
  enharmonic_policy: prefer_flats
  key_context: Ab
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
  - 8
  - 0
  - 3
  - 5
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - '6'
  preferred_positions:
  - id: a-flat-6th-voicing-1
    start_fret: 1
    end_fret: 4
    notes:
    - string: 6
      fret: 4
      pc: 8
      note: Ab
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 3
      pc: 0
      note: C
      interval: '3'
      role: third
      finger: 2
    - string: 4
      fret: 1
      pc: 3
      note: Eb
      interval: '5'
      role: fifth
      finger: 1
    - string: 3
      fret: 1
      pc: 8
      note: Ab
      interval: '1'
      role: root
      finger: 1
    - string: 2
      fret: 1
      pc: 0
      note: C
      interval: '3'
      role: third
      finger: 1
    - string: 1
      fret: 1
      pc: 5
      note: F
      interval: '6'
      role: sixth
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: a-flat-6th-voicing-2
    start_fret: 3
    end_fret: 5
    notes:
    - string: 6
      fret: 4
      pc: 8
      note: Ab
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 3
      pc: 0
      note: C
      interval: '3'
      role: third
      finger: 1
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '6'
      role: sixth
      finger: 1
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: '3'
      role: third
      finger: 3
    - string: 2
      fret: 4
      pc: 3
      note: Eb
      interval: '5'
      role: fifth
      finger: 2
    - string: 1
      fret: 4
      pc: 8
      note: Ab
      interval: '1'
      role: root
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: a-flat-6th-voicing-3
    start_fret: 4
    end_fret: 6
    notes:
    - string: 6
      fret: 4
      pc: 8
      note: Ab
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 6
      pc: 3
      note: Eb
      interval: '5'
      role: fifth
      finger: 3
    - string: 4
      fret: 6
      pc: 8
      note: Ab
      interval: '1'
      role: root
      finger: 3
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: '3'
      role: third
      finger: 2
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '6'
      role: sixth
      finger: 3
    - string: 1
      fret: 4
      pc: 8
      note: Ab
      interval: '1'
      role: root
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (4.6 3.5 1.4 1.3 1.2 1.1).2
  midi_root: 68
musical_forms:
- form: pitch_set
  id: a-flat-6th-chord-pitch-set
  pitch_classes:
  - 8
  - 0
  - 3
  - 5
  intervals:
  - '1'
  - '3'
  - '5'
  - '6'
  spelling:
  - Ab
  - C
  - Eb
  - F
- form: voicing
  id: a-flat-6th-voicing-1
  position_ref: a-flat-6th-voicing-1
  start_fret: 1
  tab: 4-3-1-1-1-1
- form: voicing
  id: a-flat-6th-voicing-2
  position_ref: a-flat-6th-voicing-2
  start_fret: 3
  tab: 4-3-3-5-4-4
- form: voicing
  id: a-flat-6th-voicing-3
  position_ref: a-flat-6th-voicing-3
  start_fret: 4
  tab: 4-6-6-5-6-4
- form: tab_excerpt
  id: a-flat-6th-chord-tab
  alphatex: \tempo 90 (4.6 3.5 1.4 1.3 1.2 1.1).2
---

# Ab 6th Chord

The Ab 6th chord is a major triad plus a major sixth — a sweet, vintage tonic substitute. Its tones are **Ab – C – Eb – F**: Ab (root), C (major third), Eb (perfect fifth), F (major sixth).

| Interval | Note | Name |
|---|---|---|
| 1 | Ab | root |
| 3 | C | major third |
| 5 | Eb | perfect fifth |
| 6 | F | major sixth |

**Playing it:** the primary shape is `4-3-1-1-1-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Ab Major Triad](/triads/a-flat-major-triad.md)
