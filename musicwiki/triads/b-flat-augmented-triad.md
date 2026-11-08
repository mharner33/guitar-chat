---
type: Concept
slug: /triads/b-flat-augmented-triad
title: Bb Augmented Triad
description: 'The augmented triad built on Bb: Bb, D, F#.'
tags:
- triad
- augmented
- b
parent: /triads
theory:
  category: triad
  quality: augmented
  root:
    pc: 10
    name: Bb
  intervals:
  - '1'
  - '3'
  - '#5'
  semitones:
  - 0
  - 4
  - 8
  pitch_classes:
  - 10
  - 2
  - 6
  spelling:
  - Bb
  - D
  - F#
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
  - 2
  - 6
  emphasize_intervals:
  - '1'
  - '3'
  - '#5'
  preferred_positions:
  - id: b-flat-augmented-voicing-1
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
      fret: 0
      pc: 2
      note: D
      interval: '3'
      role: third
      finger: 0
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 3
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '3'
      role: third
      finger: 3
    - string: 1
      fret: 2
      pc: 6
      note: F#
      interval: '#5'
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: b-flat-augmented-voicing-2
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
      fret: 5
      pc: 2
      note: D
      interval: '3'
      role: third
      finger: 3
    - string: 4
      fret: 4
      pc: 6
      note: F#
      interval: '#5'
      role: fifth
      finger: 2
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 1
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '3'
      role: third
      finger: 1
    - string: 1
      fret: 6
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 4
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: b-flat-augmented-voicing-3
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
      fret: 5
      pc: 2
      note: D
      interval: '3'
      role: third
      finger: 3
    - string: 4
      fret: 0
      pc: 2
      note: D
      interval: '3'
      role: third
      finger: 0
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 2
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '3'
      role: third
      finger: 2
    - string: 1
      fret: 2
      pc: 6
      note: F#
      interval: '#5'
      role: fifth
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
playback:
  alphatex: \tempo 90 (1.5 0.4 3.3 3.2 2.1).2
  midi_root: 70
musical_forms:
- form: pitch_set
  id: b-flat-augmented-triad-pitch-set
  pitch_classes:
  - 10
  - 2
  - 6
  intervals:
  - '1'
  - '3'
  - '#5'
  spelling:
  - Bb
  - D
  - F#
- form: voicing
  id: b-flat-augmented-voicing-1
  position_ref: b-flat-augmented-voicing-1
  start_fret: 1
  tab: x-1-0-3-3-2
- form: voicing
  id: b-flat-augmented-voicing-2
  position_ref: b-flat-augmented-voicing-2
  start_fret: 3
  tab: 6-5-4-3-3-6
- form: voicing
  id: b-flat-augmented-voicing-3
  position_ref: b-flat-augmented-voicing-3
  start_fret: 2
  tab: 6-5-0-3-3-2
- form: tab_excerpt
  id: b-flat-augmented-triad-tab
  alphatex: \tempo 90 (1.5 0.4 3.3 3.2 2.1).2
---

# Bb Augmented Triad

The Bb augmented triad is a restless, symmetrical chord of two stacked major thirds. Its tones are **Bb – D – F#**.

**Construction:** stacked thirds — a major third (Bb–D) under a major third (D–F#).

| Interval | Note | Name |
|---|---|---|
| 1 | Bb | root |
| 3 | D | major third |
| #5 | F# | augmented fifth |

**Playing it:** the primary shape is `x-1-0-3-3-2` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.
