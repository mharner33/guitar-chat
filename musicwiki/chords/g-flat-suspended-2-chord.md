---
type: Concept
slug: /chords/g-flat-suspended-2-chord
title: Gb Suspended 2 Chord
description: 'The suspended 2 chord built on Gb: Gb, Ab, Db.'
tags:
- chord
- suspended-2
- g
parent: /chords
prerequisites:
- /triads/g-flat-major-triad
theory:
  category: chord
  quality: suspended-2
  root:
    pc: 6
    name: Gb
  intervals:
  - '1'
  - '2'
  - '5'
  semitones:
  - 0
  - 2
  - 7
  pitch_classes:
  - 6
  - 8
  - 1
  spelling:
  - Gb
  - Ab
  - Db
  enharmonic_policy: prefer_flats
  key_context: Gb
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
  - 6
  - 8
  - 1
  emphasize_intervals:
  - '1'
  - '2'
  - '5'
  preferred_positions:
  - id: g-flat-suspended-2-voicing-1
    start_fret: 1
    end_fret: 4
    notes:
    - string: 6
      fret: 2
      pc: 6
      note: Gb
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 4
      pc: 1
      note: Db
      interval: '5'
      role: fifth
      finger: 3
    - string: 4
      fret: 4
      pc: 6
      note: Gb
      interval: '1'
      role: root
      finger: 3
    - string: 3
      fret: 1
      pc: 8
      note: Ab
      interval: '2'
      role: second
      finger: 1
    - string: 2
      fret: 2
      pc: 1
      note: Db
      interval: '5'
      role: fifth
      finger: 2
    - string: 1
      fret: 2
      pc: 6
      note: Gb
      interval: '1'
      role: root
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: g-flat-suspended-2-voicing-2
    start_fret: 2
    end_fret: 6
    notes:
    - string: 6
      fret: 2
      pc: 6
      note: Gb
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 4
      pc: 1
      note: Db
      interval: '5'
      role: fifth
      finger: 2
    - string: 4
      fret: 4
      pc: 6
      note: Gb
      interval: '1'
      role: root
      finger: 2
    - string: 3
      fret: 6
      pc: 1
      note: Db
      interval: '5'
      role: fifth
      finger: 3
    - string: 2
      fret: 2
      pc: 1
      note: Db
      interval: '5'
      role: fifth
      finger: 1
    - string: 1
      fret: 4
      pc: 8
      note: Ab
      interval: '2'
      role: second
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: g-flat-suspended-2-voicing-3
    start_fret: 4
    end_fret: 7
    notes:
    - string: 4
      fret: 4
      pc: 6
      note: Gb
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 6
      pc: 1
      note: Db
      interval: '5'
      role: fifth
      finger: 2
    - string: 2
      fret: 7
      pc: 6
      note: Gb
      interval: '1'
      role: root
      finger: 3
    - string: 1
      fret: 4
      pc: 8
      note: Ab
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
playback:
  alphatex: \tempo 90 (2.6 4.5 4.4 1.3 2.2 2.1).2
  midi_root: 66
musical_forms:
- form: pitch_set
  id: g-flat-suspended-2-chord-pitch-set
  pitch_classes:
  - 6
  - 8
  - 1
  intervals:
  - '1'
  - '2'
  - '5'
  spelling:
  - Gb
  - Ab
  - Db
- form: voicing
  id: g-flat-suspended-2-voicing-1
  position_ref: g-flat-suspended-2-voicing-1
  start_fret: 1
  tab: 2-4-4-1-2-2
- form: voicing
  id: g-flat-suspended-2-voicing-2
  position_ref: g-flat-suspended-2-voicing-2
  start_fret: 2
  tab: 2-4-4-6-2-4
- form: voicing
  id: g-flat-suspended-2-voicing-3
  position_ref: g-flat-suspended-2-voicing-3
  start_fret: 4
  tab: x-x-4-6-7-4
- form: tab_excerpt
  id: g-flat-suspended-2-chord-tab
  alphatex: \tempo 90 (2.6 4.5 4.4 1.3 2.2 2.1).2
---

# Gb Suspended 2 Chord

The Gb suspended 2 chord is a triad with the third replaced by the second — open and unresolved. Its tones are **Gb – Ab – Db**: Gb (root), Ab (major second), Db (perfect fifth).

| Interval | Note | Name |
|---|---|---|
| 1 | Gb | root |
| 2 | Ab | major second |
| 5 | Db | perfect fifth |

**Playing it:** the primary shape is `2-4-4-1-2-2` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Gb Major Triad](/triads/g-flat-major-triad.md)
