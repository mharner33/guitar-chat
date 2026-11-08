---
type: Concept
slug: /triads/d-augmented-triad
title: D Augmented Triad
description: 'The augmented triad built on D: D, F#, A#.'
tags:
- triad
- augmented
- d
parent: /triads
theory:
  category: triad
  quality: augmented
  root:
    pc: 2
    name: D
  intervals:
  - '1'
  - '3'
  - '#5'
  semitones:
  - 0
  - 4
  - 8
  pitch_classes:
  - 2
  - 6
  - 10
  spelling:
  - D
  - F#
  - A#
  enharmonic_policy: prefer_sharps
  key_context: D
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
  - 2
  - 6
  - 10
  emphasize_intervals:
  - '1'
  - '3'
  - '#5'
  preferred_positions:
  - id: d-augmented-voicing-1
    start_fret: 2
    end_fret: 3
    notes:
    - string: 4
      fret: 0
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 0
    - string: 3
      fret: 3
      pc: 10
      note: A#
      interval: '#5'
      role: fifth
      finger: 2
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 2
    - string: 1
      fret: 2
      pc: 6
      note: F#
      interval: '3'
      role: third
      finger: 1
    mute_strings:
    - 5
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: d-augmented-voicing-2
    start_fret: 3
    end_fret: 5
    notes:
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 4
      pc: 6
      note: F#
      interval: '3'
      role: third
      finger: 2
    - string: 3
      fret: 3
      pc: 10
      note: A#
      interval: '#5'
      role: fifth
      finger: 1
    - string: 2
      fret: 3
      pc: 2
      note: D
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
  - id: d-augmented-voicing-3
    start_fret: 4
    end_fret: 7
    notes:
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 4
      pc: 6
      note: F#
      interval: '3'
      role: third
      finger: 1
    - string: 3
      fret: 7
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 4
    - string: 2
      fret: 7
      pc: 6
      note: F#
      interval: '3'
      role: third
      finger: 4
    - string: 1
      fret: 6
      pc: 10
      note: A#
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
  alphatex: \tempo 90 (0.4 3.3 3.2 2.1).2
  midi_root: 62
musical_forms:
- form: pitch_set
  id: d-augmented-triad-pitch-set
  pitch_classes:
  - 2
  - 6
  - 10
  intervals:
  - '1'
  - '3'
  - '#5'
  spelling:
  - D
  - F#
  - A#
- form: voicing
  id: d-augmented-voicing-1
  position_ref: d-augmented-voicing-1
  start_fret: 2
  tab: x-x-0-3-3-2
- form: voicing
  id: d-augmented-voicing-2
  position_ref: d-augmented-voicing-2
  start_fret: 3
  tab: x-5-4-3-3-x
- form: voicing
  id: d-augmented-voicing-3
  position_ref: d-augmented-voicing-3
  start_fret: 4
  tab: x-5-4-7-7-6
- form: tab_excerpt
  id: d-augmented-triad-tab
  alphatex: \tempo 90 (0.4 3.3 3.2 2.1).2
---

# D Augmented Triad

The D augmented triad is a restless, symmetrical chord of two stacked major thirds. Its tones are **D – F# – A#**.

**Construction:** stacked thirds — a major third (D–F#) under a major third (F#–A#).

| Interval | Note | Name |
|---|---|---|
| 1 | D | root |
| 3 | F# | major third |
| #5 | A# | augmented fifth |

**Playing it:** the primary shape is `x-x-0-3-3-2` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.
