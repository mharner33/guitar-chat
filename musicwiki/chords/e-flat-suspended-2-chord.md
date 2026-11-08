---
type: Concept
slug: /chords/e-flat-suspended-2-chord
title: Eb Suspended 2 Chord
description: 'The suspended 2 chord built on Eb: Eb, F, Bb.'
tags:
- chord
- suspended-2
- e
parent: /chords
prerequisites:
- /triads/e-flat-major-triad
theory:
  category: chord
  quality: suspended-2
  root:
    pc: 3
    name: Eb
  intervals:
  - '1'
  - '2'
  - '5'
  semitones:
  - 0
  - 2
  - 7
  pitch_classes:
  - 3
  - 5
  - 10
  spelling:
  - Eb
  - F
  - Bb
  enharmonic_policy: prefer_flats
  key_context: Eb
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
  - 3
  - 5
  - 10
  emphasize_intervals:
  - '1'
  - '2'
  - '5'
  preferred_positions:
  - id: e-flat-suspended-2-voicing-1
    start_fret: 1
    end_fret: 4
    notes:
    - string: 4
      fret: 1
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '5'
      role: fifth
      finger: 2
    - string: 2
      fret: 4
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 3
    - string: 1
      fret: 1
      pc: 5
      note: F
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
  - id: e-flat-suspended-2-voicing-2
    start_fret: 3
    end_fret: 6
    notes:
    - string: 5
      fret: 6
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '2'
      role: second
      finger: 1
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '5'
      role: fifth
      finger: 1
    - string: 2
      fret: 4
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 2
    - string: 1
      fret: 6
      pc: 10
      note: Bb
      interval: '5'
      role: fifth
      finger: 3
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: e-flat-suspended-2-voicing-3
    start_fret: 6
    end_fret: 8
    notes:
    - string: 5
      fret: 6
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 8
      pc: 10
      note: Bb
      interval: '5'
      role: fifth
      finger: 2
    - string: 3
      fret: 8
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 2
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '2'
      role: second
      finger: 1
    - string: 1
      fret: 6
      pc: 10
      note: Bb
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
  alphatex: \tempo 90 (1.4 3.3 4.2 1.1).2
  midi_root: 63
musical_forms:
- form: pitch_set
  id: e-flat-suspended-2-chord-pitch-set
  pitch_classes:
  - 3
  - 5
  - 10
  intervals:
  - '1'
  - '2'
  - '5'
  spelling:
  - Eb
  - F
  - Bb
- form: voicing
  id: e-flat-suspended-2-voicing-1
  position_ref: e-flat-suspended-2-voicing-1
  start_fret: 1
  tab: x-x-1-3-4-1
- form: voicing
  id: e-flat-suspended-2-voicing-2
  position_ref: e-flat-suspended-2-voicing-2
  start_fret: 3
  tab: x-6-3-3-4-6
- form: voicing
  id: e-flat-suspended-2-voicing-3
  position_ref: e-flat-suspended-2-voicing-3
  start_fret: 6
  tab: x-6-8-8-6-6
- form: tab_excerpt
  id: e-flat-suspended-2-chord-tab
  alphatex: \tempo 90 (1.4 3.3 4.2 1.1).2
---

# Eb Suspended 2 Chord

The Eb suspended 2 chord is a triad with the third replaced by the second — open and unresolved. Its tones are **Eb – F – Bb**: Eb (root), F (major second), Bb (perfect fifth).

| Interval | Note | Name |
|---|---|---|
| 1 | Eb | root |
| 2 | F | major second |
| 5 | Bb | perfect fifth |

**Playing it:** the primary shape is `x-x-1-3-4-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Eb Major Triad](/triads/e-flat-major-triad.md)
