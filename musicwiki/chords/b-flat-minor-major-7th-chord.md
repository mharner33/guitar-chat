---
type: Concept
slug: /chords/b-flat-minor-major-7th-chord
title: Bb Minor Major 7th Chord
description: 'The minor major 7th chord built on Bb: Bb, Db, F, A.'
tags:
- chord
- minor-major-7th
- b
parent: /chords
prerequisites:
- /triads/b-flat-minor-triad
theory:
  category: chord
  quality: minor-major-7th
  root:
    pc: 10
    name: Bb
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
  - 10
  - 1
  - 5
  - 9
  spelling:
  - Bb
  - Db
  - F
  - A
  enharmonic_policy: prefer_flats
  key_context: Bb
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
  - 10
  - 1
  - 5
  - 9
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  - '7'
  preferred_positions:
  - id: b-flat-minor-major-7th-voicing-1
    start_fret: 1
    end_fret: 3
    notes:
    - string: 5
      fret: 1
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '5'
      role: fifth
      finger: 3
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: '7'
      role: seventh
      finger: 2
    - string: 2
      fret: 2
      pc: 1
      note: Db
      interval: b3
      role: third
      finger: 2
    - string: 1
      fret: 1
      pc: 5
      note: F
      interval: '5'
      role: fifth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: b-flat-minor-major-7th-voicing-2
    start_fret: 3
    end_fret: 6
    notes:
    - string: 6
      fret: 6
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 4
    - string: 5
      fret: 4
      pc: 1
      note: Db
      interval: b3
      role: third
      finger: 2
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '5'
      role: fifth
      finger: 1
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 1
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '5'
      role: fifth
      finger: 4
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '7'
      role: seventh
      finger: 3
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: b-flat-minor-major-7th-voicing-3
    start_fret: 2
    end_fret: 6
    notes:
    - string: 6
      fret: 6
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 4
    - string: 5
      fret: 0
      pc: 9
      note: A
      interval: '7'
      role: seventh
      finger: 0
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '5'
      role: fifth
      finger: 2
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: '7'
      role: seventh
      finger: 1
    - string: 2
      fret: 2
      pc: 1
      note: Db
      interval: b3
      role: third
      finger: 1
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '7'
      role: seventh
      finger: 3
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
playback:
  alphatex: \tempo 90 (1.5 3.4 2.3 2.2 1.1).2
  midi_root: 70
musical_forms:
- form: pitch_set
  id: b-flat-minor-major-7th-chord-pitch-set
  pitch_classes:
  - 10
  - 1
  - 5
  - 9
  intervals:
  - '1'
  - b3
  - '5'
  - '7'
  spelling:
  - Bb
  - Db
  - F
  - A
- form: voicing
  id: b-flat-minor-major-7th-voicing-1
  position_ref: b-flat-minor-major-7th-voicing-1
  start_fret: 1
  tab: x-1-3-2-2-1
- form: voicing
  id: b-flat-minor-major-7th-voicing-2
  position_ref: b-flat-minor-major-7th-voicing-2
  start_fret: 3
  tab: 6-4-3-3-6-5
- form: voicing
  id: b-flat-minor-major-7th-voicing-3
  position_ref: b-flat-minor-major-7th-voicing-3
  start_fret: 2
  tab: 6-0-3-2-2-5
- form: tab_excerpt
  id: b-flat-minor-major-7th-chord-tab
  alphatex: \tempo 90 (1.5 3.4 2.3 2.2 1.1).2
---

# Bb Minor Major 7th Chord

The Bb minor major 7th chord is a minor triad plus a major seventh — a haunting, unstable color. Its tones are **Bb – Db – F – A**: Bb (root), Db (minor third), F (perfect fifth), A (major seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | Bb | root |
| b3 | Db | minor third |
| 5 | F | perfect fifth |
| 7 | A | major seventh |

**Playing it:** the primary shape is `x-1-3-2-2-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Bb Minor Triad](/triads/b-flat-minor-triad.md)
