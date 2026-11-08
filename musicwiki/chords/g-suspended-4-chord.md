---
type: Concept
slug: /chords/g-suspended-4-chord
title: G Suspended 4 Chord
description: 'The suspended 4 chord built on G: G, C, D.'
tags:
- chord
- suspended-4
- g
parent: /chords
prerequisites:
- /triads/g-major-triad
theory:
  category: chord
  quality: suspended-4
  root:
    pc: 7
    name: G
  intervals:
  - '1'
  - '4'
  - '5'
  semitones:
  - 0
  - 5
  - 7
  pitch_classes:
  - 7
  - 0
  - 2
  spelling:
  - G
  - C
  - D
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
  - 0
  - 2
  emphasize_intervals:
  - '1'
  - '4'
  - '5'
  preferred_positions:
  - id: g-suspended-4-voicing-1
    start_fret: 1
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
      fret: 3
      pc: 0
      note: C
      interval: '4'
      role: fourth
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
      fret: 1
      pc: 0
      note: C
      interval: '4'
      role: fourth
      finger: 1
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
  - id: g-suspended-4-voicing-2
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
      fret: 3
      pc: 0
      note: C
      interval: '4'
      role: fourth
      finger: 1
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 2
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: '4'
      role: fourth
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
  - id: g-suspended-4-voicing-3
    start_fret: 5
    end_fret: 8
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
      fret: 8
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 3
    - string: 1
      fret: 8
      pc: 0
      note: C
      interval: '4'
      role: fourth
      finger: 3
    mute_strings:
    - 5
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (3.6 3.5 0.4 0.3 1.2 3.1).2
  midi_root: 67
musical_forms:
- form: pitch_set
  id: g-suspended-4-chord-pitch-set
  pitch_classes:
  - 7
  - 0
  - 2
  intervals:
  - '1'
  - '4'
  - '5'
  spelling:
  - G
  - C
  - D
- form: voicing
  id: g-suspended-4-voicing-1
  position_ref: g-suspended-4-voicing-1
  start_fret: 1
  tab: 3-3-0-0-1-3
- form: voicing
  id: g-suspended-4-voicing-2
  position_ref: g-suspended-4-voicing-2
  start_fret: 3
  tab: 3-3-5-5-3-3
- form: voicing
  id: g-suspended-4-voicing-3
  position_ref: g-suspended-4-voicing-3
  start_fret: 5
  tab: x-x-5-7-8-8
- form: tab_excerpt
  id: g-suspended-4-chord-tab
  alphatex: \tempo 90 (3.6 3.5 0.4 0.3 1.2 3.1).2
---

# G Suspended 4 Chord

The G suspended 4 chord is a triad with the third replaced by the fourth — tense, wanting to fall to the third. Its tones are **G – C – D**: G (root), C (perfect fourth), D (perfect fifth).

| Interval | Note | Name |
|---|---|---|
| 1 | G | root |
| 4 | C | perfect fourth |
| 5 | D | perfect fifth |

**Playing it:** the primary shape is `3-3-0-0-1-3` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [G Major Triad](/triads/g-major-triad.md)
