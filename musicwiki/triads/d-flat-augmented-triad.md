---
type: Concept
slug: /triads/d-flat-augmented-triad
title: Db Augmented Triad
description: 'The augmented triad built on Db: Db, F, A.'
tags:
- triad
- augmented
- d
parent: /triads
theory:
  category: triad
  quality: augmented
  root:
    pc: 1
    name: Db
  intervals:
  - '1'
  - '3'
  - '#5'
  semitones:
  - 0
  - 4
  - 8
  pitch_classes:
  - 1
  - 5
  - 9
  spelling:
  - Db
  - F
  - A
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
  - 9
  emphasize_intervals:
  - '1'
  - '3'
  - '#5'
  preferred_positions:
  - id: d-flat-augmented-voicing-1
    start_fret: 1
    end_fret: 4
    notes:
    - string: 5
      fret: 4
      pc: 1
      note: Db
      interval: '1'
      role: root
      finger: 4
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '3'
      role: third
      finger: 3
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: '#5'
      role: fifth
      finger: 2
    - string: 2
      fret: 2
      pc: 1
      note: Db
      interval: '1'
      role: root
      finger: 2
    - string: 1
      fret: 1
      pc: 5
      note: F
      interval: '3'
      role: third
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: d-flat-augmented-voicing-2
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
      fret: 2
      pc: 9
      note: A
      interval: '#5'
      role: fifth
      finger: 1
    - string: 2
      fret: 2
      pc: 1
      note: Db
      interval: '1'
      role: root
      finger: 1
    mute_strings:
    - 1
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: d-flat-augmented-voicing-3
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
      fret: 6
      pc: 1
      note: Db
      interval: '1'
      role: root
      finger: 4
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '3'
      role: third
      finger: 4
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '#5'
      role: fifth
      finger: 3
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (4.5 3.4 2.3 2.2 1.1).2
  midi_root: 61
musical_forms:
- form: pitch_set
  id: d-flat-augmented-triad-pitch-set
  pitch_classes:
  - 1
  - 5
  - 9
  intervals:
  - '1'
  - '3'
  - '#5'
  spelling:
  - Db
  - F
  - A
- form: voicing
  id: d-flat-augmented-voicing-1
  position_ref: d-flat-augmented-voicing-1
  start_fret: 1
  tab: x-4-3-2-2-1
- form: voicing
  id: d-flat-augmented-voicing-2
  position_ref: d-flat-augmented-voicing-2
  start_fret: 2
  tab: x-4-3-2-2-x
- form: voicing
  id: d-flat-augmented-voicing-3
  position_ref: d-flat-augmented-voicing-3
  start_fret: 3
  tab: x-4-3-6-6-5
- form: tab_excerpt
  id: d-flat-augmented-triad-tab
  alphatex: \tempo 90 (4.5 3.4 2.3 2.2 1.1).2
---

# Db Augmented Triad

The Db augmented triad is a restless, symmetrical chord of two stacked major thirds. Its tones are **Db – F – A**.

**Construction:** stacked thirds — a major third (Db–F) under a major third (F–A).

| Interval | Note | Name |
|---|---|---|
| 1 | Db | root |
| 3 | F | major third |
| #5 | A | augmented fifth |

**Playing it:** the primary shape is `x-4-3-2-2-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.
