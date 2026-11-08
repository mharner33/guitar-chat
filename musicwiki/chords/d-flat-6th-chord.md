---
type: Concept
slug: /chords/d-flat-6th-chord
title: Db 6th Chord
description: 'The 6th chord built on Db: Db, F, Ab, Bb.'
tags:
- chord
- 6th
- d
parent: /chords
prerequisites:
- /triads/d-flat-major-triad
theory:
  category: chord
  quality: 6th
  root:
    pc: 1
    name: Db
  intervals:
  - '1'
  - '3'
  - '5'
  - '6'
  semitones:
  - 0
  - 4
  - 7
  - 9
  pitch_classes:
  - 1
  - 5
  - 8
  - 10
  spelling:
  - Db
  - F
  - Ab
  - Bb
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
  - 5
  - 8
  - 10
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - '6'
  preferred_positions:
  - id: d-flat-6th-voicing-1
    start_fret: 2
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
      fret: 3
      pc: 5
      note: F
      interval: '3'
      role: third
      finger: 2
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '6'
      role: sixth
      finger: 2
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
      finger: 3
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: d-flat-6th-voicing-2
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
      fret: 6
      pc: 5
      note: F
      interval: '3'
      role: third
      finger: 2
    - string: 1
      fret: 6
      pc: 10
      note: Bb
      interval: '6'
      role: sixth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: d-flat-6th-voicing-3
    start_fret: 3
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
      fret: 3
      pc: 5
      note: F
      interval: '3'
      role: third
      finger: 1
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '6'
      role: sixth
      finger: 1
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '3'
      role: third
      finger: 3
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
playback:
  alphatex: \tempo 90 (4.5 3.4 3.3 2.2 4.1).2
  midi_root: 61
musical_forms:
- form: pitch_set
  id: d-flat-6th-chord-pitch-set
  pitch_classes:
  - 1
  - 5
  - 8
  - 10
  intervals:
  - '1'
  - '3'
  - '5'
  - '6'
  spelling:
  - Db
  - F
  - Ab
  - Bb
- form: voicing
  id: d-flat-6th-voicing-1
  position_ref: d-flat-6th-voicing-1
  start_fret: 2
  tab: x-4-3-3-2-4
- form: voicing
  id: d-flat-6th-voicing-2
  position_ref: d-flat-6th-voicing-2
  start_fret: 4
  tab: x-4-6-6-6-6
- form: voicing
  id: d-flat-6th-voicing-3
  position_ref: d-flat-6th-voicing-3
  start_fret: 3
  tab: x-4-3-3-6-4
- form: tab_excerpt
  id: d-flat-6th-chord-tab
  alphatex: \tempo 90 (4.5 3.4 3.3 2.2 4.1).2
---

# Db 6th Chord

The Db 6th chord is a major triad plus a major sixth — a sweet, vintage tonic substitute. Its tones are **Db – F – Ab – Bb**: Db (root), F (major third), Ab (perfect fifth), Bb (major sixth).

| Interval | Note | Name |
|---|---|---|
| 1 | Db | root |
| 3 | F | major third |
| 5 | Ab | perfect fifth |
| 6 | Bb | major sixth |

**Playing it:** the primary shape is `x-4-3-3-2-4` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Db Major Triad](/triads/d-flat-major-triad.md)
