---
type: Concept
slug: /chords/d-flat-suspended-2-chord
title: Db Suspended 2 Chord
description: 'The suspended 2 chord built on Db: Db, Eb, Ab.'
tags:
- chord
- suspended-2
- d
parent: /chords
prerequisites:
- /triads/d-flat-major-triad
theory:
  category: chord
  quality: suspended-2
  root:
    pc: 1
    name: Db
  intervals:
  - '1'
  - '2'
  - '5'
  semitones:
  - 0
  - 2
  - 7
  pitch_classes:
  - 1
  - 3
  - 8
  spelling:
  - Db
  - Eb
  - Ab
  enharmonic_policy: prefer_flats
  key_context: Db
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
  - 1
  - 3
  - 8
  emphasize_intervals:
  - '1'
  - '2'
  - '5'
  preferred_positions:
  - id: d-flat-suspended-2-voicing-1
    start_fret: 1
    end_fret: 4
    notes:
    - string: 5
      fret: 4
      pc: 1
      note: Db
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 1
      pc: 3
      note: Eb
      interval: '2'
      role: second
      finger: 1
    - string: 3
      fret: 1
      pc: 8
      note: Ab
      interval: '5'
      role: fifth
      finger: 1
    - string: 2
      fret: 2
      pc: 1
      note: Db
      interval: '1'
      role: root
      finger: 2
    - string: 1
      fret: 4
      pc: 8
      note: Ab
      interval: '5'
      role: fifth
      finger: 3
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: d-flat-suspended-2-voicing-2
    start_fret: 4
    end_fret: 6
    notes:
    - string: 5
      fret: 4
      pc: 1
      note: Db
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 6
      pc: 8
      note: Ab
      interval: '5'
      role: fifth
      finger: 2
    - string: 3
      fret: 6
      pc: 1
      note: Db
      interval: '1'
      role: root
      finger: 2
    - string: 2
      fret: 4
      pc: 3
      note: Eb
      interval: '2'
      role: second
      finger: 1
    - string: 1
      fret: 4
      pc: 8
      note: Ab
      interval: '5'
      role: fifth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: d-flat-suspended-2-voicing-3
    start_fret: 6
    end_fret: 9
    notes:
    - string: 6
      fret: 9
      pc: 1
      note: Db
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 6
      pc: 3
      note: Eb
      interval: '2'
      role: second
      finger: 1
    - string: 4
      fret: 6
      pc: 8
      note: Ab
      interval: '5'
      role: fifth
      finger: 1
    - string: 3
      fret: 6
      pc: 1
      note: Db
      interval: '1'
      role: root
      finger: 1
    - string: 2
      fret: 9
      pc: 8
      note: Ab
      interval: '5'
      role: fifth
      finger: 2
    - string: 1
      fret: 9
      pc: 1
      note: Db
      interval: '1'
      role: root
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (4.5 1.4 1.3 2.2 4.1).2
  midi_root: 61
musical_forms:
- form: pitch_set
  id: d-flat-suspended-2-chord-pitch-set
  pitch_classes:
  - 1
  - 3
  - 8
  intervals:
  - '1'
  - '2'
  - '5'
  spelling:
  - Db
  - Eb
  - Ab
- form: voicing
  id: d-flat-suspended-2-voicing-1
  position_ref: d-flat-suspended-2-voicing-1
  start_fret: 1
  tab: x-4-1-1-2-4
- form: voicing
  id: d-flat-suspended-2-voicing-2
  position_ref: d-flat-suspended-2-voicing-2
  start_fret: 4
  tab: x-4-6-6-4-4
- form: voicing
  id: d-flat-suspended-2-voicing-3
  position_ref: d-flat-suspended-2-voicing-3
  start_fret: 6
  tab: 9-6-6-6-9-9
- form: tab_excerpt
  id: d-flat-suspended-2-chord-tab
  alphatex: \tempo 90 (4.5 1.4 1.3 2.2 4.1).2
---

# Db Suspended 2 Chord

The Db suspended 2 chord is a triad with the third replaced by the second — open and unresolved. Its tones are **Db – Eb – Ab**: Db (root), Eb (major second), Ab (perfect fifth).

| Interval | Note | Name |
|---|---|---|
| 1 | Db | root |
| 2 | Eb | major second |
| 5 | Ab | perfect fifth |

**Playing it:** the primary shape is `x-4-1-1-2-4` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Db Major Triad](/triads/d-flat-major-triad.md)
