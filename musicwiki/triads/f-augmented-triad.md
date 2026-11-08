---
type: Concept
slug: /triads/f-augmented-triad
title: F Augmented Triad
description: 'The augmented triad built on F: F, A, C#.'
tags:
- triad
- augmented
- f
parent: /triads
theory:
  category: triad
  quality: augmented
  root:
    pc: 5
    name: F
  intervals:
  - '1'
  - '3'
  - '#5'
  semitones:
  - 0
  - 4
  - 8
  pitch_classes:
  - 5
  - 9
  - 1
  spelling:
  - F
  - A
  - C#
  enharmonic_policy: prefer_sharps
  key_context: F
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
  - 5
  - 9
  - 1
  emphasize_intervals:
  - '1'
  - '3'
  - '#5'
  preferred_positions:
  - id: f-augmented-voicing-1
    start_fret: 1
    end_fret: 3
    notes:
    - string: 6
      fret: 1
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 0
      pc: 9
      note: A
      interval: '3'
      role: third
      finger: 0
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 3
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: '3'
      role: third
      finger: 2
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: '#5'
      role: fifth
      finger: 2
    - string: 1
      fret: 1
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: f-augmented-voicing-2
    start_fret: 2
    end_fret: 5
    notes:
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 2
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: '3'
      role: third
      finger: 1
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: '#5'
      role: fifth
      finger: 1
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '3'
      role: third
      finger: 3
    mute_strings:
    - 5
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: f-augmented-voicing-3
    start_fret: 3
    end_fret: 6
    notes:
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 6
      pc: 1
      note: C#
      interval: '#5'
      role: fifth
      finger: 3
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 3
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '3'
      role: third
      finger: 2
    mute_strings:
    - 5
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (1.6 0.5 3.4 2.3 2.2 1.1).2
  midi_root: 65
musical_forms:
- form: pitch_set
  id: f-augmented-triad-pitch-set
  pitch_classes:
  - 5
  - 9
  - 1
  intervals:
  - '1'
  - '3'
  - '#5'
  spelling:
  - F
  - A
  - C#
- form: voicing
  id: f-augmented-voicing-1
  position_ref: f-augmented-voicing-1
  start_fret: 1
  tab: 1-0-3-2-2-1
- form: voicing
  id: f-augmented-voicing-2
  position_ref: f-augmented-voicing-2
  start_fret: 2
  tab: x-x-3-2-2-5
- form: voicing
  id: f-augmented-voicing-3
  position_ref: f-augmented-voicing-3
  start_fret: 3
  tab: x-x-3-6-6-5
- form: tab_excerpt
  id: f-augmented-triad-tab
  alphatex: \tempo 90 (1.6 0.5 3.4 2.3 2.2 1.1).2
---

# F Augmented Triad

The F augmented triad is a restless, symmetrical chord of two stacked major thirds. Its tones are **F – A – C#**.

**Construction:** stacked thirds — a major third (F–A) under a major third (A–C#).

| Interval | Note | Name |
|---|---|---|
| 1 | F | root |
| 3 | A | major third |
| #5 | C# | augmented fifth |

**Playing it:** the primary shape is `1-0-3-2-2-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.
