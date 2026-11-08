---
type: Concept
slug: /chords/g-6th-chord
title: G 6th Chord
description: 'The 6th chord built on G: G, B, D, E.'
tags:
- chord
- 6th
- g
parent: /chords
prerequisites:
- /triads/g-major-triad
theory:
  category: chord
  quality: 6th
  root:
    pc: 7
    name: G
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
  - 7
  - 11
  - 2
  - 4
  spelling:
  - G
  - B
  - D
  - E
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
  - 4
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - '6'
  preferred_positions:
  - id: g-6th-voicing-1
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
      fret: 0
      pc: 4
      note: E
      interval: '6'
      role: sixth
      finger: 0
    confidence: 0.75
    tags:
    - computed
    - open
  - id: g-6th-voicing-2
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
      fret: 5
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 3
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '3'
      role: third
      finger: 2
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '6'
      role: sixth
      finger: 3
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
  - id: g-6th-voicing-3
    start_fret: 5
    end_fret: 7
    notes:
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 7
      pc: 2
      note: D
      interval: '5'
      role: fifth
      finger: 2
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '6'
      role: sixth
      finger: 1
    - string: 1
      fret: 7
      pc: 11
      note: B
      interval: '3'
      role: third
      finger: 2
    mute_strings:
    - 5
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (3.6 2.5 0.4 0.3 0.2 0.1).2
  midi_root: 67
musical_forms:
- form: pitch_set
  id: g-6th-chord-pitch-set
  pitch_classes:
  - 7
  - 11
  - 2
  - 4
  intervals:
  - '1'
  - '3'
  - '5'
  - '6'
  spelling:
  - G
  - B
  - D
  - E
- form: voicing
  id: g-6th-voicing-1
  position_ref: g-6th-voicing-1
  start_fret: 2
  tab: 3-2-0-0-0-0
- form: voicing
  id: g-6th-voicing-2
  position_ref: g-6th-voicing-2
  start_fret: 3
  tab: 3-5-5-4-5-3
- form: voicing
  id: g-6th-voicing-3
  position_ref: g-6th-voicing-3
  start_fret: 5
  tab: x-x-5-7-5-7
- form: tab_excerpt
  id: g-6th-chord-tab
  alphatex: \tempo 90 (3.6 2.5 0.4 0.3 0.2 0.1).2
---

# G 6th Chord

The G 6th chord is a major triad plus a major sixth — a sweet, vintage tonic substitute. Its tones are **G – B – D – E**: G (root), B (major third), D (perfect fifth), E (major sixth).

| Interval | Note | Name |
|---|---|---|
| 1 | G | root |
| 3 | B | major third |
| 5 | D | perfect fifth |
| 6 | E | major sixth |

**Playing it:** the primary shape is `3-2-0-0-0-0` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [G Major Triad](/triads/g-major-triad.md)
