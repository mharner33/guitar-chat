---
type: Concept
slug: /chords/d-minor-major-7th-chord
title: D Minor Major 7th Chord
description: 'The minor major 7th chord built on D: D, F, A, C#.'
tags:
- chord
- minor-major-7th
- d
parent: /chords
prerequisites:
- /triads/d-minor-triad
theory:
  category: chord
  quality: minor-major-7th
  root:
    pc: 2
    name: D
  intervals:
  - '1'
  - b3
  - '5'
  - '7'
  semitones:
  - 0
  - 3
  - 7
  - 11
  pitch_classes:
  - 2
  - 5
  - 9
  - 1
  spelling:
  - D
  - F
  - A
  - C#
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
  - 5
  - 9
  - 1
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  - '7'
  preferred_positions:
  - id: d-minor-major-7th-voicing-1
    start_fret: 1
    end_fret: 2
    notes:
    - string: 4
      fret: 0
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 0
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: '5'
      role: fifth
      finger: 2
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: '7'
      role: seventh
      finger: 2
    - string: 1
      fret: 1
      pc: 5
      note: F
      interval: b3
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
  - id: d-minor-major-7th-voicing-2
    start_fret: 2
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
      fret: 3
      pc: 5
      note: F
      interval: b3
      role: third
      finger: 2
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: '5'
      role: fifth
      finger: 1
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: '7'
      role: seventh
      finger: 1
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '5'
      role: fifth
      finger: 3
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: d-minor-major-7th-voicing-3
    start_fret: 3
    end_fret: 6
    notes:
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: b3
      role: third
      finger: 1
    - string: 3
      fret: 6
      pc: 1
      note: C#
      interval: '7'
      role: seventh
      finger: 3
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 1
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '5'
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (0.4 2.3 2.2 1.1).2
  midi_root: 62
musical_forms:
- form: pitch_set
  id: d-minor-major-7th-chord-pitch-set
  pitch_classes:
  - 2
  - 5
  - 9
  - 1
  intervals:
  - '1'
  - b3
  - '5'
  - '7'
  spelling:
  - D
  - F
  - A
  - C#
- form: voicing
  id: d-minor-major-7th-voicing-1
  position_ref: d-minor-major-7th-voicing-1
  start_fret: 1
  tab: x-x-0-2-2-1
- form: voicing
  id: d-minor-major-7th-voicing-2
  position_ref: d-minor-major-7th-voicing-2
  start_fret: 2
  tab: x-5-3-2-2-5
- form: voicing
  id: d-minor-major-7th-voicing-3
  position_ref: d-minor-major-7th-voicing-3
  start_fret: 3
  tab: x-5-3-6-3-5
- form: tab_excerpt
  id: d-minor-major-7th-chord-tab
  alphatex: \tempo 90 (0.4 2.3 2.2 1.1).2
---

# D Minor Major 7th Chord

The D minor major 7th chord is a minor triad plus a major seventh — a haunting, unstable color. Its tones are **D – F – A – C#**: D (root), F (minor third), A (perfect fifth), C# (major seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | D | root |
| b3 | F | minor third |
| 5 | A | perfect fifth |
| 7 | C# | major seventh |

**Playing it:** the primary shape is `x-x-0-2-2-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [D Minor Triad](/triads/d-minor-triad.md)
