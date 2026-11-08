---
type: Concept
slug: /chords/d-suspended-2-chord
title: D Suspended 2 Chord
description: 'The suspended 2 chord built on D: D, E, A.'
tags:
- chord
- suspended-2
- d
parent: /chords
prerequisites:
- /triads/d-major-triad
theory:
  category: chord
  quality: suspended-2
  root:
    pc: 2
    name: D
  intervals:
  - '1'
  - '2'
  - '5'
  semitones:
  - 0
  - 2
  - 7
  pitch_classes:
  - 2
  - 4
  - 9
  spelling:
  - D
  - E
  - A
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
  - 4
  - 9
  emphasize_intervals:
  - '1'
  - '2'
  - '5'
  preferred_positions:
  - id: d-suspended-2-voicing-1
    start_fret: 2
    end_fret: 3
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
      fret: 3
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 2
    - string: 1
      fret: 0
      pc: 4
      note: E
      interval: '2'
      role: second
      finger: 0
    mute_strings:
    - 5
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
  - id: d-suspended-2-voicing-2
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
      fret: 5
      pc: 4
      note: E
      interval: '2'
      role: second
      finger: 1
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '5'
      role: fifth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: d-suspended-2-voicing-3
    start_fret: 7
    end_fret: 10
    notes:
    - string: 6
      fret: 10
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 7
      pc: 4
      note: E
      interval: '2'
      role: second
      finger: 1
    - string: 4
      fret: 7
      pc: 9
      note: A
      interval: '5'
      role: fifth
      finger: 1
    - string: 3
      fret: 7
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 1
    - string: 2
      fret: 10
      pc: 9
      note: A
      interval: '5'
      role: fifth
      finger: 2
    - string: 1
      fret: 10
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (0.4 2.3 3.2 0.1).2
  midi_root: 62
musical_forms:
- form: pitch_set
  id: d-suspended-2-chord-pitch-set
  pitch_classes:
  - 2
  - 4
  - 9
  intervals:
  - '1'
  - '2'
  - '5'
  spelling:
  - D
  - E
  - A
- form: voicing
  id: d-suspended-2-voicing-1
  position_ref: d-suspended-2-voicing-1
  start_fret: 2
  tab: x-x-0-2-3-0
- form: voicing
  id: d-suspended-2-voicing-2
  position_ref: d-suspended-2-voicing-2
  start_fret: 5
  tab: x-5-7-7-5-5
- form: voicing
  id: d-suspended-2-voicing-3
  position_ref: d-suspended-2-voicing-3
  start_fret: 7
  tab: 10-7-7-7-10-10
- form: tab_excerpt
  id: d-suspended-2-chord-tab
  alphatex: \tempo 90 (0.4 2.3 3.2 0.1).2
---

# D Suspended 2 Chord

The D suspended 2 chord is a triad with the third replaced by the second — open and unresolved. Its tones are **D – E – A**: D (root), E (major second), A (perfect fifth).

| Interval | Note | Name |
|---|---|---|
| 1 | D | root |
| 2 | E | major second |
| 5 | A | perfect fifth |

**Playing it:** the primary shape is `x-x-0-2-3-0` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [D Major Triad](/triads/d-major-triad.md)
