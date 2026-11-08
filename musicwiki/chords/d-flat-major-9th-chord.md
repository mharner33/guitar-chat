---
type: Concept
slug: /chords/d-flat-major-9th-chord
title: Db Major 9th Chord
description: 'The major 9th chord built on Db: Db, F, Ab, C, Eb.'
tags:
- chord
- major-9th
- d
parent: /chords
prerequisites:
- /chords/d-flat-major-7th-chord
theory:
  category: chord
  quality: major-9th
  root:
    pc: 1
    name: Db
  intervals:
  - '1'
  - '3'
  - '5'
  - '7'
  - '9'
  semitones:
  - 0
  - 4
  - 7
  - 11
  - 14
  pitch_classes:
  - 1
  - 5
  - 8
  - 0
  - 3
  spelling:
  - Db
  - F
  - Ab
  - C
  - Eb
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
  - 0
  - 3
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - '7'
  - '9'
  preferred_positions:
  - id: d-flat-major-9th-voicing-1
    start_fret: 1
    end_fret: 4
    notes:
    - string: 5
      fret: 4
      pc: 1
      note: Db
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 1
      pc: 3
      note: Eb
      interval: '9'
      role: ninth
      finger: 1
    - string: 3
      fret: 1
      pc: 8
      note: Ab
      interval: '5'
      role: fifth
      finger: 1
    - string: 2
      fret: 1
      pc: 0
      note: C
      interval: '7'
      role: seventh
      finger: 1
    - string: 1
      fret: 1
      pc: 5
      note: F
      interval: '3'
      role: third
      finger: 1
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: d-flat-major-9th-voicing-2
    start_fret: 3
    end_fret: 5
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
      fret: 5
      pc: 0
      note: C
      interval: '7'
      role: seventh
      finger: 3
    - string: 2
      fret: 4
      pc: 3
      note: Eb
      interval: '9'
      role: ninth
      finger: 2
    - string: 1
      fret: 4
      pc: 8
      note: Ab
      interval: '5'
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: d-flat-major-9th-voicing-3
    start_fret: 6
    end_fret: 9
    notes:
    - string: 6
      fret: 9
      pc: 1
      note: Db
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 6
      pc: 3
      note: Eb
      interval: '9'
      role: ninth
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
      fret: 6
      pc: 5
      note: F
      interval: '3'
      role: third
      finger: 1
    - string: 1
      fret: 8
      pc: 0
      note: C
      interval: '7'
      role: seventh
      finger: 2
    confidence: 0.85
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (4.5 1.4 1.3 1.2 1.1).2
  midi_root: 61
musical_forms:
- form: pitch_set
  id: d-flat-major-9th-chord-pitch-set
  pitch_classes:
  - 1
  - 5
  - 8
  - 0
  - 3
  intervals:
  - '1'
  - '3'
  - '5'
  - '7'
  - '9'
  spelling:
  - Db
  - F
  - Ab
  - C
  - Eb
- form: voicing
  id: d-flat-major-9th-voicing-1
  position_ref: d-flat-major-9th-voicing-1
  start_fret: 1
  tab: x-4-1-1-1-1
- form: voicing
  id: d-flat-major-9th-voicing-2
  position_ref: d-flat-major-9th-voicing-2
  start_fret: 3
  tab: x-4-3-5-4-4
- form: voicing
  id: d-flat-major-9th-voicing-3
  position_ref: d-flat-major-9th-voicing-3
  start_fret: 6
  tab: 9-6-6-6-6-8
- form: tab_excerpt
  id: d-flat-major-9th-chord-tab
  alphatex: \tempo 90 (4.5 1.4 1.3 1.2 1.1).2
---

# Db Major 9th Chord

The Db major 9th chord is a major seventh extended with the ninth — expansive and warm. Its tones are **Db – F – Ab – C – Eb**: Db (root), F (major third), Ab (perfect fifth), C (major seventh), Eb (ninth).

| Interval | Note | Name |
|---|---|---|
| 1 | Db | root |
| 3 | F | major third |
| 5 | Ab | perfect fifth |
| 7 | C | major seventh |
| 9 | Eb | ninth |

**Playing it:** the primary shape is `x-4-1-1-1-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Db Major 7th Chord](/chords/d-flat-major-7th-chord.md)
