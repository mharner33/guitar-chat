---
type: Concept
slug: /triads/b-augmented-triad
title: B Augmented Triad
description: 'The augmented triad built on B: B, D#, F##.'
tags:
- triad
- augmented
- b
parent: /triads
theory:
  category: triad
  quality: augmented
  root:
    pc: 11
    name: B
  intervals:
  - '1'
  - '3'
  - '#5'
  semitones:
  - 0
  - 4
  - 8
  pitch_classes:
  - 11
  - 3
  - 7
  spelling:
  - B
  - D#
  - F##
  enharmonic_policy: prefer_sharps
  key_context: B
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
  - 11
  - 3
  - 7
  emphasize_intervals:
  - '1'
  - '3'
  - '#5'
  preferred_positions:
  - id: b-augmented-voicing-1
    start_fret: 1
    end_fret: 2
    notes:
    - string: 5
      fret: 2
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 1
      pc: 3
      note: D#
      interval: '3'
      role: third
      finger: 1
    - string: 3
      fret: 0
      pc: 7
      note: F##
      interval: '#5'
      role: fifth
      finger: 0
    - string: 2
      fret: 0
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 0
    mute_strings:
    - 1
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
  - id: b-augmented-voicing-2
    start_fret: 2
    end_fret: 5
    notes:
    - string: 5
      fret: 2
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 5
      pc: 7
      note: F##
      interval: '#5'
      role: fifth
      finger: 4
    - string: 3
      fret: 0
      pc: 7
      note: F##
      interval: '#5'
      role: fifth
      finger: 0
    - string: 2
      fret: 4
      pc: 3
      note: D#
      interval: '3'
      role: third
      finger: 3
    - string: 1
      fret: 3
      pc: 7
      note: F##
      interval: '#5'
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
  - id: b-augmented-voicing-3
    start_fret: 4
    end_fret: 7
    notes:
    - string: 6
      fret: 7
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 4
    - string: 5
      fret: 6
      pc: 3
      note: D#
      interval: '3'
      role: third
      finger: 3
    - string: 4
      fret: 5
      pc: 7
      note: F##
      interval: '#5'
      role: fifth
      finger: 2
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 1
    - string: 2
      fret: 4
      pc: 3
      note: D#
      interval: '3'
      role: third
      finger: 1
    - string: 1
      fret: 7
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 4
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (2.5 1.4 0.3 0.2).2
  midi_root: 71
musical_forms:
- form: pitch_set
  id: b-augmented-triad-pitch-set
  pitch_classes:
  - 11
  - 3
  - 7
  intervals:
  - '1'
  - '3'
  - '#5'
  spelling:
  - B
  - D#
  - F##
- form: voicing
  id: b-augmented-voicing-1
  position_ref: b-augmented-voicing-1
  start_fret: 1
  tab: x-2-1-0-0-x
- form: voicing
  id: b-augmented-voicing-2
  position_ref: b-augmented-voicing-2
  start_fret: 2
  tab: x-2-5-0-4-3
- form: voicing
  id: b-augmented-voicing-3
  position_ref: b-augmented-voicing-3
  start_fret: 4
  tab: 7-6-5-4-4-7
- form: tab_excerpt
  id: b-augmented-triad-tab
  alphatex: \tempo 90 (2.5 1.4 0.3 0.2).2
---

# B Augmented Triad

The B augmented triad is a restless, symmetrical chord of two stacked major thirds. Its tones are **B – D# – F##**.

**Construction:** stacked thirds — a major third (B–D#) under a major third (D#–F##).

| Interval | Note | Name |
|---|---|---|
| 1 | B | root |
| 3 | D# | major third |
| #5 | F## | augmented fifth |

**Playing it:** the primary shape is `x-2-1-0-0-x` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.
