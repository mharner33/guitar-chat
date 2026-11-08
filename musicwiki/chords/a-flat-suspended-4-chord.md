---
type: Concept
slug: /chords/a-flat-suspended-4-chord
title: Ab Suspended 4 Chord
description: 'The suspended 4 chord built on Ab: Ab, Db, Eb.'
tags:
- chord
- suspended-4
- a
parent: /chords
prerequisites:
- /triads/a-flat-major-triad
theory:
  category: chord
  quality: suspended-4
  root:
    pc: 8
    name: Ab
  intervals:
  - '1'
  - '4'
  - '5'
  semitones:
  - 0
  - 5
  - 7
  pitch_classes:
  - 8
  - 1
  - 3
  spelling:
  - Ab
  - Db
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
  - 1
  - 3
  emphasize_intervals:
  - '1'
  - '4'
  - '5'
  preferred_positions:
  - id: a-flat-suspended-4-voicing-1
    start_fret: 1
    end_fret: 4
    notes:
    - string: 6
      fret: 4
      pc: 8
      note: Ab
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 4
      pc: 1
      note: Db
      interval: '4'
      role: fourth
      finger: 3
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
      fret: 2
      pc: 1
      note: Db
      interval: '4'
      role: fourth
      finger: 2
    - string: 1
      fret: 4
      pc: 8
      note: Ab
      interval: '1'
      role: root
      finger: 3
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: a-flat-suspended-4-voicing-2
    start_fret: 4
    end_fret: 6
    notes:
    - string: 6
      fret: 4
      pc: 8
      note: Ab
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 4
      pc: 1
      note: Db
      interval: '4'
      role: fourth
      finger: 1
    - string: 4
      fret: 6
      pc: 8
      note: Ab
      interval: '1'
      role: root
      finger: 2
    - string: 3
      fret: 6
      pc: 1
      note: Db
      interval: '4'
      role: fourth
      finger: 2
    - string: 2
      fret: 4
      pc: 3
      note: Eb
      interval: '5'
      role: fifth
      finger: 1
    - string: 1
      fret: 4
      pc: 8
      note: Ab
      interval: '1'
      role: root
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: a-flat-suspended-4-voicing-3
    start_fret: 2
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
      fret: 6
      pc: 1
      note: Db
      interval: '4'
      role: fourth
      finger: 3
    - string: 2
      fret: 2
      pc: 1
      note: Db
      interval: '4'
      role: fourth
      finger: 1
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
playback:
  alphatex: \tempo 90 (4.6 4.5 1.4 1.3 2.2 4.1).2
  midi_root: 68
musical_forms:
- form: pitch_set
  id: a-flat-suspended-4-chord-pitch-set
  pitch_classes:
  - 8
  - 1
  - 3
  intervals:
  - '1'
  - '4'
  - '5'
  spelling:
  - Ab
  - Db
  - Eb
- form: voicing
  id: a-flat-suspended-4-voicing-1
  position_ref: a-flat-suspended-4-voicing-1
  start_fret: 1
  tab: 4-4-1-1-2-4
- form: voicing
  id: a-flat-suspended-4-voicing-2
  position_ref: a-flat-suspended-4-voicing-2
  start_fret: 4
  tab: 4-4-6-6-4-4
- form: voicing
  id: a-flat-suspended-4-voicing-3
  position_ref: a-flat-suspended-4-voicing-3
  start_fret: 2
  tab: 4-6-6-6-2-4
- form: tab_excerpt
  id: a-flat-suspended-4-chord-tab
  alphatex: \tempo 90 (4.6 4.5 1.4 1.3 2.2 4.1).2
---

# Ab Suspended 4 Chord

The Ab suspended 4 chord is a triad with the third replaced by the fourth — tense, wanting to fall to the third. Its tones are **Ab – Db – Eb**: Ab (root), Db (perfect fourth), Eb (perfect fifth).

| Interval | Note | Name |
|---|---|---|
| 1 | Ab | root |
| 4 | Db | perfect fourth |
| 5 | Eb | perfect fifth |

**Playing it:** the primary shape is `4-4-1-1-2-4` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Ab Major Triad](/triads/a-flat-major-triad.md)
