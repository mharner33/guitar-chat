---
type: Concept
slug: /triads/a-augmented-triad
title: A Augmented Triad
description: 'The augmented triad built on A: A, C#, E#.'
tags:
- triad
- augmented
- a
parent: /triads
theory:
  category: triad
  quality: augmented
  root:
    pc: 9
    name: A
  intervals:
  - '1'
  - '3'
  - '#5'
  semitones:
  - 0
  - 4
  - 8
  pitch_classes:
  - 9
  - 1
  - 5
  spelling:
  - A
  - C#
  - E#
  enharmonic_policy: prefer_sharps
  key_context: A
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
  - 9
  - 1
  - 5
  emphasize_intervals:
  - '1'
  - '3'
  - '#5'
  preferred_positions:
  - id: a-augmented-voicing-1
    start_fret: 1
    end_fret: 3
    notes:
    - string: 5
      fret: 0
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 0
    - string: 4
      fret: 3
      pc: 5
      note: E#
      interval: '#5'
      role: fifth
      finger: 3
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 2
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: '3'
      role: third
      finger: 2
    - string: 1
      fret: 1
      pc: 5
      note: E#
      interval: '#5'
      role: fifth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: a-augmented-voicing-2
    start_fret: 2
    end_fret: 3
    notes:
    - string: 5
      fret: 0
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 0
    - string: 4
      fret: 3
      pc: 5
      note: E#
      interval: '#5'
      role: fifth
      finger: 2
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 1
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: '3'
      role: third
      finger: 1
    mute_strings:
    - 1
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: a-augmented-voicing-3
    start_fret: 3
    end_fret: 6
    notes:
    - string: 6
      fret: 5
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 4
      pc: 1
      note: C#
      interval: '3'
      role: third
      finger: 2
    - string: 4
      fret: 3
      pc: 5
      note: E#
      interval: '#5'
      role: fifth
      finger: 1
    - string: 3
      fret: 6
      pc: 1
      note: C#
      interval: '3'
      role: third
      finger: 4
    - string: 2
      fret: 6
      pc: 5
      note: E#
      interval: '#5'
      role: fifth
      finger: 4
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 3
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (0.5 3.4 2.3 2.2 1.1).2
  midi_root: 69
musical_forms:
- form: pitch_set
  id: a-augmented-triad-pitch-set
  pitch_classes:
  - 9
  - 1
  - 5
  intervals:
  - '1'
  - '3'
  - '#5'
  spelling:
  - A
  - C#
  - E#
- form: voicing
  id: a-augmented-voicing-1
  position_ref: a-augmented-voicing-1
  start_fret: 1
  tab: x-0-3-2-2-1
- form: voicing
  id: a-augmented-voicing-2
  position_ref: a-augmented-voicing-2
  start_fret: 2
  tab: x-0-3-2-2-x
- form: voicing
  id: a-augmented-voicing-3
  position_ref: a-augmented-voicing-3
  start_fret: 3
  tab: 5-4-3-6-6-5
- form: tab_excerpt
  id: a-augmented-triad-tab
  alphatex: \tempo 90 (0.5 3.4 2.3 2.2 1.1).2
---

# A Augmented Triad

The A augmented triad is a restless, symmetrical chord of two stacked major thirds. Its tones are **A – C# – E#**.

**Construction:** stacked thirds — a major third (A–C#) under a major third (C#–E#).

| Interval | Note | Name |
|---|---|---|
| 1 | A | root |
| 3 | C# | major third |
| #5 | E# | augmented fifth |

**Playing it:** the primary shape is `x-0-3-2-2-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.
