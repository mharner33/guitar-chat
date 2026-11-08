---
type: Concept
slug: /chords/d-6th-chord
title: D 6th Chord
description: 'The 6th chord built on D: D, F#, A, B.'
tags:
- chord
- 6th
- d
parent: /chords
prerequisites:
- /triads/d-major-triad
theory:
  category: chord
  quality: 6th
  root:
    pc: 2
    name: D
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
  - 2
  - 6
  - 9
  - 11
  spelling:
  - D
  - F#
  - A
  - B
  enharmonic_policy: prefer_sharps
  key_context: D
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
  - 2
  - 6
  - 9
  - 11
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - '6'
  preferred_positions:
  - id: d-6th-voicing-1
    start_fret: 2
    end_fret: 2
    notes:
    - string: 4
      fret: 0
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 0
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: '5'
      role: fifth
      finger: 1
    - string: 2
      fret: 0
      pc: 11
      note: B
      interval: '6'
      role: sixth
      finger: 0
    - string: 1
      fret: 2
      pc: 6
      note: F#
      interval: '3'
      role: third
      finger: 1
    mute_strings:
    - 5
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: d-6th-voicing-2
    start_fret: 3
    end_fret: 5
    notes:
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 4
      pc: 6
      note: F#
      interval: '3'
      role: third
      finger: 2
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '6'
      role: sixth
      finger: 2
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 1
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '5'
      role: fifth
      finger: 3
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: d-6th-voicing-3
    start_fret: 5
    end_fret: 7
    notes:
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 7
      pc: 9
      note: A
      interval: '5'
      role: fifth
      finger: 2
    - string: 3
      fret: 7
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 2
    - string: 2
      fret: 7
      pc: 6
      note: F#
      interval: '3'
      role: third
      finger: 2
    - string: 1
      fret: 7
      pc: 11
      note: B
      interval: '6'
      role: sixth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (0.4 2.3 0.2 2.1).2
  midi_root: 62
musical_forms:
- form: pitch_set
  id: d-6th-chord-pitch-set
  pitch_classes:
  - 2
  - 6
  - 9
  - 11
  intervals:
  - '1'
  - '3'
  - '5'
  - '6'
  spelling:
  - D
  - F#
  - A
  - B
- form: voicing
  id: d-6th-voicing-1
  position_ref: d-6th-voicing-1
  start_fret: 2
  tab: x-x-0-2-0-2
- form: voicing
  id: d-6th-voicing-2
  position_ref: d-6th-voicing-2
  start_fret: 3
  tab: x-5-4-4-3-5
- form: voicing
  id: d-6th-voicing-3
  position_ref: d-6th-voicing-3
  start_fret: 5
  tab: x-5-7-7-7-7
- form: tab_excerpt
  id: d-6th-chord-tab
  alphatex: \tempo 90 (0.4 2.3 0.2 2.1).2
---

# D 6th Chord

The D 6th chord is a major triad plus a major sixth — a sweet, vintage tonic substitute. Its tones are **D – F# – A – B**: D (root), F# (major third), A (perfect fifth), B (major sixth).

| Interval | Note | Name |
|---|---|---|
| 1 | D | root |
| 3 | F# | major third |
| 5 | A | perfect fifth |
| 6 | B | major sixth |

**Playing it:** the primary shape is `x-x-0-2-0-2` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [D Major Triad](/triads/d-major-triad.md)
