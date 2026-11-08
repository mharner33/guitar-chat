---
type: Concept
slug: /chords/d-flat-minor-major-7th-chord
title: Db Minor Major 7th Chord
description: 'The minor major 7th chord built on Db: Db, Fb, Ab, C.'
tags:
- chord
- minor-major-7th
- d
parent: /chords
prerequisites:
- /triads/c-sharp-minor-triad
theory:
  category: chord
  quality: minor-major-7th
  root:
    pc: 1
    name: Db
  intervals:
  - '1'
  - b3
  - '5'
  - '7'
  semitones:
  - 0
  - 3
  - 7
  - 11
  pitch_classes:
  - 1
  - 4
  - 8
  - 0
  spelling:
  - Db
  - Fb
  - Ab
  - C
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
  - 4
  - 8
  - 0
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  - '7'
  preferred_positions:
  - id: d-flat-minor-major-7th-voicing-1
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
      fret: 2
      pc: 4
      note: Fb
      interval: b3
      role: third
      finger: 2
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
      fret: 0
      pc: 4
      note: Fb
      interval: b3
      role: third
      finger: 0
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: d-flat-minor-major-7th-voicing-2
    start_fret: 2
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
      fret: 2
      pc: 4
      note: Fb
      interval: b3
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
  - id: d-flat-minor-major-7th-voicing-3
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
      finger: 3
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: '7'
      role: seventh
      finger: 2
    - string: 2
      fret: 5
      pc: 4
      note: Fb
      interval: b3
      role: third
      finger: 2
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
  alphatex: \tempo 90 (4.5 2.4 1.3 1.2 0.1).2
  midi_root: 61
musical_forms:
- form: pitch_set
  id: d-flat-minor-major-7th-chord-pitch-set
  pitch_classes:
  - 1
  - 4
  - 8
  - 0
  intervals:
  - '1'
  - b3
  - '5'
  - '7'
  spelling:
  - Db
  - Fb
  - Ab
  - C
- form: voicing
  id: d-flat-minor-major-7th-voicing-1
  position_ref: d-flat-minor-major-7th-voicing-1
  start_fret: 1
  tab: x-4-2-1-1-0
- form: voicing
  id: d-flat-minor-major-7th-voicing-2
  position_ref: d-flat-minor-major-7th-voicing-2
  start_fret: 2
  tab: x-4-2-5-2-4
- form: voicing
  id: d-flat-minor-major-7th-voicing-3
  position_ref: d-flat-minor-major-7th-voicing-3
  start_fret: 4
  tab: x-4-6-5-5-4
- form: tab_excerpt
  id: d-flat-minor-major-7th-chord-tab
  alphatex: \tempo 90 (4.5 2.4 1.3 1.2 0.1).2
---

# Db Minor Major 7th Chord

The Db minor major 7th chord is a minor triad plus a major seventh — a haunting, unstable color. Its tones are **Db – Fb – Ab – C**: Db (root), Fb (minor third), Ab (perfect fifth), C (major seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | Db | root |
| b3 | Fb | minor third |
| 5 | Ab | perfect fifth |
| 7 | C | major seventh |

**Playing it:** the primary shape is `x-4-2-1-1-0` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [C# Minor Triad](/triads/c-sharp-minor-triad.md)
