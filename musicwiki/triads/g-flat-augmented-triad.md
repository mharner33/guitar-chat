---
type: Concept
slug: /triads/g-flat-augmented-triad
title: Gb Augmented Triad
description: 'The augmented triad built on Gb: Gb, Bb, D.'
tags:
- triad
- augmented
- g
parent: /triads
theory:
  category: triad
  quality: augmented
  root:
    pc: 6
    name: Gb
  intervals:
  - '1'
  - '3'
  - '#5'
  semitones:
  - 0
  - 4
  - 8
  pitch_classes:
  - 6
  - 10
  - 2
  spelling:
  - Gb
  - Bb
  - D
  enharmonic_policy: prefer_flats
  key_context: Gb
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
  - 6
  - 10
  - 2
  emphasize_intervals:
  - '1'
  - '3'
  - '#5'
  preferred_positions:
  - id: g-flat-augmented-voicing-1
    start_fret: 1
    end_fret: 3
    notes:
    - string: 6
      fret: 2
      pc: 6
      note: Gb
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 1
      pc: 10
      note: Bb
      interval: '3'
      role: third
      finger: 1
    - string: 4
      fret: 0
      pc: 2
      note: D
      interval: '#5'
      role: fifth
      finger: 0
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '3'
      role: third
      finger: 3
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '#5'
      role: fifth
      finger: 3
    - string: 1
      fret: 2
      pc: 6
      note: Gb
      interval: '1'
      role: root
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: g-flat-augmented-voicing-2
    start_fret: 2
    end_fret: 5
    notes:
    - string: 6
      fret: 2
      pc: 6
      note: Gb
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: '#5'
      role: fifth
      finger: 3
    - string: 4
      fret: 0
      pc: 2
      note: D
      interval: '#5'
      role: fifth
      finger: 0
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '3'
      role: third
      finger: 2
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '#5'
      role: fifth
      finger: 2
    - string: 1
      fret: 2
      pc: 6
      note: Gb
      interval: '1'
      role: root
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: g-flat-augmented-voicing-3
    start_fret: 3
    end_fret: 6
    notes:
    - string: 4
      fret: 4
      pc: 6
      note: Gb
      interval: '1'
      role: root
      finger: 2
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '3'
      role: third
      finger: 1
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '#5'
      role: fifth
      finger: 1
    - string: 1
      fret: 6
      pc: 10
      note: Bb
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
playback:
  alphatex: \tempo 90 (2.6 1.5 0.4 3.3 3.2 2.1).2
  midi_root: 66
musical_forms:
- form: pitch_set
  id: g-flat-augmented-triad-pitch-set
  pitch_classes:
  - 6
  - 10
  - 2
  intervals:
  - '1'
  - '3'
  - '#5'
  spelling:
  - Gb
  - Bb
  - D
- form: voicing
  id: g-flat-augmented-voicing-1
  position_ref: g-flat-augmented-voicing-1
  start_fret: 1
  tab: 2-1-0-3-3-2
- form: voicing
  id: g-flat-augmented-voicing-2
  position_ref: g-flat-augmented-voicing-2
  start_fret: 2
  tab: 2-5-0-3-3-2
- form: voicing
  id: g-flat-augmented-voicing-3
  position_ref: g-flat-augmented-voicing-3
  start_fret: 3
  tab: x-x-4-3-3-6
- form: tab_excerpt
  id: g-flat-augmented-triad-tab
  alphatex: \tempo 90 (2.6 1.5 0.4 3.3 3.2 2.1).2
---

# Gb Augmented Triad

The Gb augmented triad is a restless, symmetrical chord of two stacked major thirds. Its tones are **Gb – Bb – D**.

**Construction:** stacked thirds — a major third (Gb–Bb) under a major third (Bb–D).

| Interval | Note | Name |
|---|---|---|
| 1 | Gb | root |
| 3 | Bb | major third |
| #5 | D | augmented fifth |

**Playing it:** the primary shape is `2-1-0-3-3-2` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.
