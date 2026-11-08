---
type: Concept
slug: /chords/d-suspended-4-chord
title: D Suspended 4 Chord
description: 'The suspended 4 chord built on D: D, G, A.'
tags:
- chord
- suspended-4
- d
parent: /chords
prerequisites:
- /triads/d-major-triad
theory:
  category: chord
  quality: suspended-4
  root:
    pc: 2
    name: D
  intervals:
  - '1'
  - '4'
  - '5'
  semitones:
  - 0
  - 5
  - 7
  pitch_classes:
  - 2
  - 7
  - 9
  spelling:
  - D
  - G
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
  - 7
  - 9
  emphasize_intervals:
  - '1'
  - '4'
  - '5'
  preferred_positions:
  - id: d-suspended-4-voicing-1
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
      fret: 3
      pc: 7
      note: G
      interval: '4'
      role: fourth
      finger: 2
    mute_strings:
    - 5
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: d-suspended-4-voicing-2
    start_fret: 3
    end_fret: 7
    notes:
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: '4'
      role: fourth
      finger: 2
    - string: 3
      fret: 7
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 3
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
      finger: 2
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: d-suspended-4-voicing-3
    start_fret: 5
    end_fret: 8
    notes:
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: '4'
      role: fourth
      finger: 1
    - string: 3
      fret: 7
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 2
    - string: 2
      fret: 8
      pc: 7
      note: G
      interval: '4'
      role: fourth
      finger: 3
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
playback:
  alphatex: \tempo 90 (0.4 2.3 3.2 3.1).2
  midi_root: 62
musical_forms:
- form: pitch_set
  id: d-suspended-4-chord-pitch-set
  pitch_classes:
  - 2
  - 7
  - 9
  intervals:
  - '1'
  - '4'
  - '5'
  spelling:
  - D
  - G
  - A
- form: voicing
  id: d-suspended-4-voicing-1
  position_ref: d-suspended-4-voicing-1
  start_fret: 2
  tab: x-x-0-2-3-3
- form: voicing
  id: d-suspended-4-voicing-2
  position_ref: d-suspended-4-voicing-2
  start_fret: 3
  tab: x-5-5-7-3-5
- form: voicing
  id: d-suspended-4-voicing-3
  position_ref: d-suspended-4-voicing-3
  start_fret: 5
  tab: x-5-5-7-8-5
- form: tab_excerpt
  id: d-suspended-4-chord-tab
  alphatex: \tempo 90 (0.4 2.3 3.2 3.1).2
---

# D Suspended 4 Chord

The D suspended 4 chord is a triad with the third replaced by the fourth — tense, wanting to fall to the third. Its tones are **D – G – A**: D (root), G (perfect fourth), A (perfect fifth).

| Interval | Note | Name |
|---|---|---|
| 1 | D | root |
| 4 | G | perfect fourth |
| 5 | A | perfect fifth |

**Playing it:** the primary shape is `x-x-0-2-3-3` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [D Major Triad](/triads/d-major-triad.md)
