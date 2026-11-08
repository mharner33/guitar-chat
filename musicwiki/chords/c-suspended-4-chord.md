---
type: Concept
slug: /chords/c-suspended-4-chord
title: C Suspended 4 Chord
description: 'The suspended 4 chord built on C: C, F, G.'
tags:
- chord
- suspended-4
- c
parent: /chords
prerequisites:
- /triads/c-major-triad
theory:
  category: chord
  quality: suspended-4
  root:
    pc: 0
    name: C
  intervals:
  - '1'
  - '4'
  - '5'
  semitones:
  - 0
  - 5
  - 7
  pitch_classes:
  - 0
  - 5
  - 7
  spelling:
  - C
  - F
  - G
  enharmonic_policy: prefer_sharps
  key_context: C
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
  - 0
  - 5
  - 7
  emphasize_intervals:
  - '1'
  - '4'
  - '5'
  preferred_positions:
  - id: c-suspended-4-voicing-1
    start_fret: 1
    end_fret: 3
    notes:
    - string: 5
      fret: 3
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '4'
      role: fourth
      finger: 2
    - string: 3
      fret: 0
      pc: 7
      note: G
      interval: '5'
      role: fifth
      finger: 0
    - string: 2
      fret: 1
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 1
    - string: 1
      fret: 1
      pc: 5
      note: F
      interval: '4'
      role: fourth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: c-suspended-4-voicing-2
    start_fret: 3
    end_fret: 6
    notes:
    - string: 5
      fret: 3
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '4'
      role: fourth
      finger: 1
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 2
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '4'
      role: fourth
      finger: 3
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: '5'
      role: fifth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: c-suspended-4-voicing-3
    start_fret: 5
    end_fret: 8
    notes:
    - string: 6
      fret: 8
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 8
      pc: 5
      note: F
      interval: '4'
      role: fourth
      finger: 3
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: '5'
      role: fifth
      finger: 1
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 1
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '4'
      role: fourth
      finger: 2
    - string: 1
      fret: 8
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 3
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (3.5 3.4 0.3 1.2 1.1).2
  midi_root: 60
musical_forms:
- form: pitch_set
  id: c-suspended-4-chord-pitch-set
  pitch_classes:
  - 0
  - 5
  - 7
  intervals:
  - '1'
  - '4'
  - '5'
  spelling:
  - C
  - F
  - G
- form: voicing
  id: c-suspended-4-voicing-1
  position_ref: c-suspended-4-voicing-1
  start_fret: 1
  tab: x-3-3-0-1-1
- form: voicing
  id: c-suspended-4-voicing-2
  position_ref: c-suspended-4-voicing-2
  start_fret: 3
  tab: x-3-3-5-6-3
- form: voicing
  id: c-suspended-4-voicing-3
  position_ref: c-suspended-4-voicing-3
  start_fret: 5
  tab: 8-8-5-5-6-8
- form: tab_excerpt
  id: c-suspended-4-chord-tab
  alphatex: \tempo 90 (3.5 3.4 0.3 1.2 1.1).2
---

# C Suspended 4 Chord

The C suspended 4 chord is a triad with the third replaced by the fourth — tense, wanting to fall to the third. Its tones are **C – F – G**: C (root), F (perfect fourth), G (perfect fifth).

| Interval | Note | Name |
|---|---|---|
| 1 | C | root |
| 4 | F | perfect fourth |
| 5 | G | perfect fifth |

**Playing it:** the primary shape is `x-3-3-0-1-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [C Major Triad](/triads/c-major-triad.md)
