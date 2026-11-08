---
type: Concept
slug: /chords/d-flat-suspended-4-chord
title: Db Suspended 4 Chord
description: 'The suspended 4 chord built on Db: Db, Gb, Ab.'
tags:
- chord
- suspended-4
- d
parent: /chords
prerequisites:
- /triads/d-flat-major-triad
theory:
  category: chord
  quality: suspended-4
  root:
    pc: 1
    name: Db
  intervals:
  - '1'
  - '4'
  - '5'
  semitones:
  - 0
  - 5
  - 7
  pitch_classes:
  - 1
  - 6
  - 8
  spelling:
  - Db
  - Gb
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
  - 6
  - 8
  emphasize_intervals:
  - '1'
  - '4'
  - '5'
  preferred_positions:
  - id: d-flat-suspended-4-voicing-1
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
      fret: 4
      pc: 6
      note: Gb
      interval: '4'
      role: fourth
      finger: 3
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
      fret: 2
      pc: 6
      note: Gb
      interval: '4'
      role: fourth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: d-flat-suspended-4-voicing-2
    start_fret: 2
    end_fret: 6
    notes:
    - string: 5
      fret: 4
      pc: 1
      note: Db
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 4
      pc: 6
      note: Gb
      interval: '4'
      role: fourth
      finger: 2
    - string: 3
      fret: 6
      pc: 1
      note: Db
      interval: '1'
      role: root
      finger: 3
    - string: 2
      fret: 2
      pc: 1
      note: Db
      interval: '1'
      role: root
      finger: 1
    - string: 1
      fret: 4
      pc: 8
      note: Ab
      interval: '5'
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: d-flat-suspended-4-voicing-3
    start_fret: 4
    end_fret: 7
    notes:
    - string: 5
      fret: 4
      pc: 1
      note: Db
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 4
      pc: 6
      note: Gb
      interval: '4'
      role: fourth
      finger: 1
    - string: 3
      fret: 6
      pc: 1
      note: Db
      interval: '1'
      role: root
      finger: 2
    - string: 2
      fret: 7
      pc: 6
      note: Gb
      interval: '4'
      role: fourth
      finger: 3
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
playback:
  alphatex: \tempo 90 (4.5 4.4 1.3 2.2 2.1).2
  midi_root: 61
musical_forms:
- form: pitch_set
  id: d-flat-suspended-4-chord-pitch-set
  pitch_classes:
  - 1
  - 6
  - 8
  intervals:
  - '1'
  - '4'
  - '5'
  spelling:
  - Db
  - Gb
  - Ab
- form: voicing
  id: d-flat-suspended-4-voicing-1
  position_ref: d-flat-suspended-4-voicing-1
  start_fret: 1
  tab: x-4-4-1-2-2
- form: voicing
  id: d-flat-suspended-4-voicing-2
  position_ref: d-flat-suspended-4-voicing-2
  start_fret: 2
  tab: x-4-4-6-2-4
- form: voicing
  id: d-flat-suspended-4-voicing-3
  position_ref: d-flat-suspended-4-voicing-3
  start_fret: 4
  tab: x-4-4-6-7-4
- form: tab_excerpt
  id: d-flat-suspended-4-chord-tab
  alphatex: \tempo 90 (4.5 4.4 1.3 2.2 2.1).2
---

# Db Suspended 4 Chord

The Db suspended 4 chord is a triad with the third replaced by the fourth — tense, wanting to fall to the third. Its tones are **Db – Gb – Ab**: Db (root), Gb (perfect fourth), Ab (perfect fifth).

| Interval | Note | Name |
|---|---|---|
| 1 | Db | root |
| 4 | Gb | perfect fourth |
| 5 | Ab | perfect fifth |

**Playing it:** the primary shape is `x-4-4-1-2-2` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Db Major Triad](/triads/d-flat-major-triad.md)
