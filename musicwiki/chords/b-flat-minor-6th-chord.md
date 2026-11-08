---
type: Concept
slug: /chords/b-flat-minor-6th-chord
title: Bb Minor 6th Chord
description: 'The minor 6th chord built on Bb: Bb, Db, F, G.'
tags:
- chord
- minor-6th
- b
parent: /chords
prerequisites:
- /triads/b-flat-minor-triad
theory:
  category: chord
  quality: minor-6th
  root:
    pc: 10
    name: Bb
  intervals:
  - '1'
  - b3
  - '5'
  - '6'
  semitones:
  - 0
  - 3
  - 7
  - 9
  pitch_classes:
  - 10
  - 1
  - 5
  - 7
  spelling:
  - Bb
  - Db
  - F
  - G
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
  - 7
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  - '6'
  preferred_positions:
  - id: b-flat-minor-6th-voicing-1
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
      fret: 0
      pc: 7
      note: G
      interval: '6'
      role: sixth
      finger: 0
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
    - open
    - barre
  - id: b-flat-minor-6th-voicing-2
    start_fret: 4
    end_fret: 6
    notes:
    - string: 6
      fret: 6
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 4
      pc: 1
      note: Db
      interval: b3
      role: third
      finger: 1
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: '6'
      role: sixth
      finger: 2
    - string: 3
      fret: 6
      pc: 1
      note: Db
      interval: b3
      role: third
      finger: 3
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '5'
      role: fifth
      finger: 3
    - string: 1
      fret: 6
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 3
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: b-flat-minor-6th-voicing-3
    start_fret: 3
    end_fret: 6
    notes:
    - string: 6
      fret: 6
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 3
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
      finger: 3
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: '6'
      role: sixth
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (1.5 3.4 0.3 2.2 1.1).2
  midi_root: 70
musical_forms:
- form: pitch_set
  id: b-flat-minor-6th-chord-pitch-set
  pitch_classes:
  - 10
  - 1
  - 5
  - 7
  intervals:
  - '1'
  - b3
  - '5'
  - '6'
  spelling:
  - Bb
  - Db
  - F
  - G
- form: voicing
  id: b-flat-minor-6th-voicing-1
  position_ref: b-flat-minor-6th-voicing-1
  start_fret: 1
  tab: x-1-3-0-2-1
- form: voicing
  id: b-flat-minor-6th-voicing-2
  position_ref: b-flat-minor-6th-voicing-2
  start_fret: 4
  tab: 6-4-5-6-6-6
- form: voicing
  id: b-flat-minor-6th-voicing-3
  position_ref: b-flat-minor-6th-voicing-3
  start_fret: 3
  tab: 6-4-3-3-6-3
- form: tab_excerpt
  id: b-flat-minor-6th-chord-tab
  alphatex: \tempo 90 (1.5 3.4 0.3 2.2 1.1).2
---

# Bb Minor 6th Chord

The Bb minor 6th chord is a minor triad plus a major sixth — a bright-edged minor color. Its tones are **Bb – Db – F – G**: Bb (root), Db (minor third), F (perfect fifth), G (major sixth).

| Interval | Note | Name |
|---|---|---|
| 1 | Bb | root |
| b3 | Db | minor third |
| 5 | F | perfect fifth |
| 6 | G | major sixth |

**Playing it:** the primary shape is `x-1-3-0-2-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Bb Minor Triad](/triads/b-flat-minor-triad.md)
