---
type: Concept
slug: /triads/e-augmented-triad
title: E Augmented Triad
description: 'The augmented triad built on E: E, G#, B#.'
tags:
- triad
- augmented
- e
parent: /triads
theory:
  category: triad
  quality: augmented
  root:
    pc: 4
    name: E
  intervals:
  - '1'
  - '3'
  - '#5'
  semitones:
  - 0
  - 4
  - 8
  pitch_classes:
  - 4
  - 8
  - 0
  spelling:
  - E
  - G#
  - B#
  enharmonic_policy: prefer_sharps
  key_context: E
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
  - 4
  - 8
  - 0
  emphasize_intervals:
  - '1'
  - '3'
  - '#5'
  preferred_positions:
  - id: e-augmented-voicing-1
    start_fret: 1
    end_fret: 3
    notes:
    - string: 6
      fret: 0
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 0
    - string: 5
      fret: 3
      pc: 0
      note: B#
      interval: '#5'
      role: fifth
      finger: 3
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 2
    - string: 3
      fret: 1
      pc: 8
      note: G#
      interval: '3'
      role: third
      finger: 1
    - string: 2
      fret: 1
      pc: 0
      note: B#
      interval: '#5'
      role: fifth
      finger: 1
    - string: 1
      fret: 0
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 0
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: e-augmented-voicing-2
    start_fret: 2
    end_fret: 5
    notes:
    - string: 6
      fret: 0
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 0
    - string: 5
      fret: 3
      pc: 0
      note: B#
      interval: '#5'
      role: fifth
      finger: 2
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 5
      pc: 0
      note: B#
      interval: '#5'
      role: fifth
      finger: 4
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 4
    - string: 1
      fret: 4
      pc: 8
      note: G#
      interval: '3'
      role: third
      finger: 3
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: e-augmented-voicing-3
    start_fret: 4
    end_fret: 7
    notes:
    - string: 5
      fret: 7
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 4
    - string: 4
      fret: 6
      pc: 8
      note: G#
      interval: '3'
      role: third
      finger: 3
    - string: 3
      fret: 5
      pc: 0
      note: B#
      interval: '#5'
      role: fifth
      finger: 2
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 2
    - string: 1
      fret: 4
      pc: 8
      note: G#
      interval: '3'
      role: third
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (0.6 3.5 2.4 1.3 1.2 0.1).2
  midi_root: 64
musical_forms:
- form: pitch_set
  id: e-augmented-triad-pitch-set
  pitch_classes:
  - 4
  - 8
  - 0
  intervals:
  - '1'
  - '3'
  - '#5'
  spelling:
  - E
  - G#
  - B#
- form: voicing
  id: e-augmented-voicing-1
  position_ref: e-augmented-voicing-1
  start_fret: 1
  tab: 0-3-2-1-1-0
- form: voicing
  id: e-augmented-voicing-2
  position_ref: e-augmented-voicing-2
  start_fret: 2
  tab: 0-3-2-5-5-4
- form: voicing
  id: e-augmented-voicing-3
  position_ref: e-augmented-voicing-3
  start_fret: 4
  tab: x-7-6-5-5-4
- form: tab_excerpt
  id: e-augmented-triad-tab
  alphatex: \tempo 90 (0.6 3.5 2.4 1.3 1.2 0.1).2
---

# E Augmented Triad

The E augmented triad is a restless, symmetrical chord of two stacked major thirds. Its tones are **E – G# – B#**.

**Construction:** stacked thirds — a major third (E–G#) under a major third (G#–B#).

| Interval | Note | Name |
|---|---|---|
| 1 | E | root |
| 3 | G# | major third |
| #5 | B# | augmented fifth |

**Playing it:** the primary shape is `0-3-2-1-1-0` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.
