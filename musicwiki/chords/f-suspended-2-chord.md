---
type: Concept
slug: /chords/f-suspended-2-chord
title: F Suspended 2 Chord
description: 'The suspended 2 chord built on F: F, G, C.'
tags:
- chord
- suspended-2
- f
parent: /chords
prerequisites:
- /triads/f-major-triad
theory:
  category: chord
  quality: suspended-2
  root:
    pc: 5
    name: F
  intervals:
  - '1'
  - '2'
  - '5'
  semitones:
  - 0
  - 2
  - 7
  pitch_classes:
  - 5
  - 7
  - 0
  spelling:
  - F
  - G
  - C
  enharmonic_policy: prefer_sharps
  key_context: F
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
  - 5
  - 7
  - 0
  emphasize_intervals:
  - '1'
  - '2'
  - '5'
  preferred_positions:
  - id: f-suspended-2-voicing-1
    start_fret: 1
    end_fret: 3
    notes:
    - string: 6
      fret: 1
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 3
      pc: 0
      note: C
      interval: '5'
      role: fifth
      finger: 2
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 2
    - string: 3
      fret: 0
      pc: 7
      note: G
      interval: '2'
      role: second
      finger: 0
    - string: 2
      fret: 1
      pc: 0
      note: C
      interval: '5'
      role: fifth
      finger: 1
    - string: 1
      fret: 1
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: f-suspended-2-voicing-2
    start_fret: 3
    end_fret: 6
    notes:
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: '5'
      role: fifth
      finger: 2
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 3
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: '2'
      role: second
      finger: 1
    mute_strings:
    - 5
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: f-suspended-2-voicing-3
    start_fret: 5
    end_fret: 8
    notes:
    - string: 5
      fret: 8
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: '2'
      role: second
      finger: 1
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: '5'
      role: fifth
      finger: 1
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 2
    - string: 1
      fret: 8
      pc: 0
      note: C
      interval: '5'
      role: fifth
      finger: 3
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (1.6 3.5 3.4 0.3 1.2 1.1).2
  midi_root: 65
musical_forms:
- form: pitch_set
  id: f-suspended-2-chord-pitch-set
  pitch_classes:
  - 5
  - 7
  - 0
  intervals:
  - '1'
  - '2'
  - '5'
  spelling:
  - F
  - G
  - C
- form: voicing
  id: f-suspended-2-voicing-1
  position_ref: f-suspended-2-voicing-1
  start_fret: 1
  tab: 1-3-3-0-1-1
- form: voicing
  id: f-suspended-2-voicing-2
  position_ref: f-suspended-2-voicing-2
  start_fret: 3
  tab: x-x-3-5-6-3
- form: voicing
  id: f-suspended-2-voicing-3
  position_ref: f-suspended-2-voicing-3
  start_fret: 5
  tab: x-8-5-5-6-8
- form: tab_excerpt
  id: f-suspended-2-chord-tab
  alphatex: \tempo 90 (1.6 3.5 3.4 0.3 1.2 1.1).2
---

# F Suspended 2 Chord

The F suspended 2 chord is a triad with the third replaced by the second — open and unresolved. Its tones are **F – G – C**: F (root), G (major second), C (perfect fifth).

| Interval | Note | Name |
|---|---|---|
| 1 | F | root |
| 2 | G | major second |
| 5 | C | perfect fifth |

**Playing it:** the primary shape is `1-3-3-0-1-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [F Major Triad](/triads/f-major-triad.md)
