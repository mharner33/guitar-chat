---
type: Concept
slug: /triads/a-flat-augmented-triad
title: Ab Augmented Triad
description: 'The augmented triad built on Ab: Ab, C, E.'
tags:
- triad
- augmented
- a
parent: /triads
theory:
  category: triad
  quality: augmented
  root:
    pc: 8
    name: Ab
  intervals:
  - '1'
  - '3'
  - '#5'
  semitones:
  - 0
  - 4
  - 8
  pitch_classes:
  - 8
  - 0
  - 4
  spelling:
  - Ab
  - C
  - E
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
  - 0
  - 4
  emphasize_intervals:
  - '1'
  - '3'
  - '#5'
  preferred_positions:
  - id: a-flat-augmented-voicing-1
    start_fret: 1
    end_fret: 4
    notes:
    - string: 6
      fret: 4
      pc: 8
      note: Ab
      interval: '1'
      role: root
      finger: 4
    - string: 5
      fret: 3
      pc: 0
      note: C
      interval: '3'
      role: third
      finger: 3
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: '#5'
      role: fifth
      finger: 2
    - string: 3
      fret: 1
      pc: 8
      note: Ab
      interval: '1'
      role: root
      finger: 1
    - string: 2
      fret: 1
      pc: 0
      note: C
      interval: '3'
      role: third
      finger: 1
    - string: 1
      fret: 0
      pc: 4
      note: E
      interval: '#5'
      role: fifth
      finger: 0
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: a-flat-augmented-voicing-2
    start_fret: 2
    end_fret: 5
    notes:
    - string: 6
      fret: 4
      pc: 8
      note: Ab
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 3
      pc: 0
      note: C
      interval: '3'
      role: third
      finger: 2
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: '#5'
      role: fifth
      finger: 1
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: '3'
      role: third
      finger: 4
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '#5'
      role: fifth
      finger: 4
    - string: 1
      fret: 0
      pc: 4
      note: E
      interval: '#5'
      role: fifth
      finger: 0
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: a-flat-augmented-voicing-3
    start_fret: 3
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
      fret: 3
      pc: 0
      note: C
      interval: '3'
      role: third
      finger: 1
    - string: 4
      fret: 6
      pc: 8
      note: Ab
      interval: '1'
      role: root
      finger: 4
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: '3'
      role: third
      finger: 3
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '#5'
      role: fifth
      finger: 3
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
  alphatex: \tempo 90 (4.6 3.5 2.4 1.3 1.2 0.1).2
  midi_root: 68
musical_forms:
- form: pitch_set
  id: a-flat-augmented-triad-pitch-set
  pitch_classes:
  - 8
  - 0
  - 4
  intervals:
  - '1'
  - '3'
  - '#5'
  spelling:
  - Ab
  - C
  - E
- form: voicing
  id: a-flat-augmented-voicing-1
  position_ref: a-flat-augmented-voicing-1
  start_fret: 1
  tab: 4-3-2-1-1-0
- form: voicing
  id: a-flat-augmented-voicing-2
  position_ref: a-flat-augmented-voicing-2
  start_fret: 2
  tab: 4-3-2-5-5-0
- form: voicing
  id: a-flat-augmented-voicing-3
  position_ref: a-flat-augmented-voicing-3
  start_fret: 3
  tab: 4-3-6-5-5-4
- form: tab_excerpt
  id: a-flat-augmented-triad-tab
  alphatex: \tempo 90 (4.6 3.5 2.4 1.3 1.2 0.1).2
---

# Ab Augmented Triad

The Ab augmented triad is a restless, symmetrical chord of two stacked major thirds. Its tones are **Ab – C – E**.

**Construction:** stacked thirds — a major third (Ab–C) under a major third (C–E).

| Interval | Note | Name |
|---|---|---|
| 1 | Ab | root |
| 3 | C | major third |
| #5 | E | augmented fifth |

**Playing it:** the primary shape is `4-3-2-1-1-0` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.
