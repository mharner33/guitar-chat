---
type: Concept
slug: /chords/a-flat-suspended-2-chord
title: Ab Suspended 2 Chord
description: 'The suspended 2 chord built on Ab: Ab, Bb, Eb.'
tags:
- chord
- suspended-2
- a
parent: /chords
prerequisites:
- /triads/a-flat-major-triad
theory:
  category: chord
  quality: suspended-2
  root:
    pc: 8
    name: Ab
  intervals:
  - '1'
  - '2'
  - '5'
  semitones:
  - 0
  - 2
  - 7
  pitch_classes:
  - 8
  - 10
  - 3
  spelling:
  - Ab
  - Bb
  - Eb
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
  - 10
  - 3
  emphasize_intervals:
  - '1'
  - '2'
  - '5'
  preferred_positions:
  - id: a-flat-suspended-2-voicing-1
    start_fret: 1
    end_fret: 4
    notes:
    - string: 6
      fret: 4
      pc: 8
      note: Ab
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 1
      pc: 10
      note: Bb
      interval: '2'
      role: second
      finger: 1
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
  - id: a-flat-suspended-2-voicing-2
    start_fret: 3
    end_fret: 6
    notes:
    - string: 6
      fret: 4
      pc: 8
      note: Ab
      interval: '1'
      role: root
      finger: 2
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
      fret: 3
      pc: 10
      note: Bb
      interval: '2'
      role: second
      finger: 1
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
  - id: a-flat-suspended-2-voicing-3
    start_fret: 4
    end_fret: 8
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
      finger: 2
    - string: 4
      fret: 6
      pc: 8
      note: Ab
      interval: '1'
      role: root
      finger: 2
    - string: 3
      fret: 8
      pc: 3
      note: Eb
      interval: '5'
      role: fifth
      finger: 3
    - string: 2
      fret: 4
      pc: 3
      note: Eb
      interval: '5'
      role: fifth
      finger: 1
    - string: 1
      fret: 6
      pc: 10
      note: Bb
      interval: '2'
      role: second
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (4.6 1.5 1.4 1.3 4.2 4.1).2
  midi_root: 68
musical_forms:
- form: pitch_set
  id: a-flat-suspended-2-chord-pitch-set
  pitch_classes:
  - 8
  - 10
  - 3
  intervals:
  - '1'
  - '2'
  - '5'
  spelling:
  - Ab
  - Bb
  - Eb
- form: voicing
  id: a-flat-suspended-2-voicing-1
  position_ref: a-flat-suspended-2-voicing-1
  start_fret: 1
  tab: 4-1-1-1-4-4
- form: voicing
  id: a-flat-suspended-2-voicing-2
  position_ref: a-flat-suspended-2-voicing-2
  start_fret: 3
  tab: 4-6-6-3-4-4
- form: voicing
  id: a-flat-suspended-2-voicing-3
  position_ref: a-flat-suspended-2-voicing-3
  start_fret: 4
  tab: 4-6-6-8-4-6
- form: tab_excerpt
  id: a-flat-suspended-2-chord-tab
  alphatex: \tempo 90 (4.6 1.5 1.4 1.3 4.2 4.1).2
---

# Ab Suspended 2 Chord

The Ab suspended 2 chord is a triad with the third replaced by the second — open and unresolved. Its tones are **Ab – Bb – Eb**: Ab (root), Bb (major second), Eb (perfect fifth).

| Interval | Note | Name |
|---|---|---|
| 1 | Ab | root |
| 2 | Bb | major second |
| 5 | Eb | perfect fifth |

**Playing it:** the primary shape is `4-1-1-1-4-4` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Ab Major Triad](/triads/a-flat-major-triad.md)
