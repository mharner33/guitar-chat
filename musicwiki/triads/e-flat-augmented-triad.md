---
type: Concept
slug: /triads/e-flat-augmented-triad
title: Eb Augmented Triad
description: 'The augmented triad built on Eb: Eb, G, B.'
tags:
- triad
- augmented
- e
parent: /triads
theory:
  category: triad
  quality: augmented
  root:
    pc: 3
    name: Eb
  intervals:
  - '1'
  - '3'
  - '#5'
  semitones:
  - 0
  - 4
  - 8
  pitch_classes:
  - 3
  - 7
  - 11
  spelling:
  - Eb
  - G
  - B
  enharmonic_policy: prefer_flats
  key_context: Eb
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
  - 3
  - 7
  - 11
  emphasize_intervals:
  - '1'
  - '3'
  - '#5'
  preferred_positions:
  - id: e-flat-augmented-voicing-1
    start_fret: 1
    end_fret: 3
    notes:
    - string: 4
      fret: 1
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 0
      pc: 7
      note: G
      interval: '3'
      role: third
      finger: 0
    - string: 2
      fret: 0
      pc: 11
      note: B
      interval: '#5'
      role: fifth
      finger: 0
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: '3'
      role: third
      finger: 2
    mute_strings:
    - 5
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
  - id: e-flat-augmented-voicing-2
    start_fret: 3
    end_fret: 6
    notes:
    - string: 5
      fret: 6
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 4
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: '3'
      role: third
      finger: 3
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '#5'
      role: fifth
      finger: 2
    - string: 2
      fret: 4
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 2
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: '3'
      role: third
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: e-flat-augmented-voicing-3
    start_fret: 4
    end_fret: 6
    notes:
    - string: 5
      fret: 6
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: '3'
      role: third
      finger: 2
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '#5'
      role: fifth
      finger: 1
    - string: 2
      fret: 4
      pc: 3
      note: Eb
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
playback:
  alphatex: \tempo 90 (1.4 0.3 0.2 3.1).2
  midi_root: 63
musical_forms:
- form: pitch_set
  id: e-flat-augmented-triad-pitch-set
  pitch_classes:
  - 3
  - 7
  - 11
  intervals:
  - '1'
  - '3'
  - '#5'
  spelling:
  - Eb
  - G
  - B
- form: voicing
  id: e-flat-augmented-voicing-1
  position_ref: e-flat-augmented-voicing-1
  start_fret: 1
  tab: x-x-1-0-0-3
- form: voicing
  id: e-flat-augmented-voicing-2
  position_ref: e-flat-augmented-voicing-2
  start_fret: 3
  tab: x-6-5-4-4-3
- form: voicing
  id: e-flat-augmented-voicing-3
  position_ref: e-flat-augmented-voicing-3
  start_fret: 4
  tab: x-6-5-4-4-x
- form: tab_excerpt
  id: e-flat-augmented-triad-tab
  alphatex: \tempo 90 (1.4 0.3 0.2 3.1).2
---

# Eb Augmented Triad

The Eb augmented triad is a restless, symmetrical chord of two stacked major thirds. Its tones are **Eb – G – B**.

**Construction:** stacked thirds — a major third (Eb–G) under a major third (G–B).

| Interval | Note | Name |
|---|---|---|
| 1 | Eb | root |
| 3 | G | major third |
| #5 | B | augmented fifth |

**Playing it:** the primary shape is `x-x-1-0-0-3` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.
