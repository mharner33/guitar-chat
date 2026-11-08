---
type: Concept
slug: /chords/g-minor-major-7th-chord
title: G Minor Major 7th Chord
description: 'The minor major 7th chord built on G: G, Bb, D, F#.'
tags:
- chord
- minor-major-7th
- g
parent: /chords
prerequisites:
- /triads/g-minor-triad
theory:
  category: chord
  quality: minor-major-7th
  root:
    pc: 7
    name: G
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
  - 7
  - 10
  - 2
  - 6
  spelling:
  - G
  - Bb
  - D
  - F#
  enharmonic_policy: prefer_flats
  key_context: G
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
  - 7
  - 10
  - 2
  - 6
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  - '7'
  preferred_positions:
  - id: g-minor-major-7th-voicing-1
    start_fret: 1
    end_fret: 3
    notes:
    - string: 6
      fret: 3
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 1
      pc: 10
      note: Bb
      interval: b3
      role: third
      finger: 1
    - string: 4
      fret: 0
      pc: 2
      note: D
      interval: '5'
      role: fifth
      finger: 0
    - string: 3
      fret: 0
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 0
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '5'
      role: fifth
      finger: 3
    - string: 1
      fret: 2
      pc: 6
      note: F#
      interval: '7'
      role: seventh
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: g-minor-major-7th-voicing-2
    start_fret: 3
    end_fret: 5
    notes:
    - string: 6
      fret: 3
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: '5'
      role: fifth
      finger: 3
    - string: 4
      fret: 4
      pc: 6
      note: F#
      interval: '7'
      role: seventh
      finger: 2
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: b3
      role: third
      finger: 1
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '5'
      role: fifth
      finger: 1
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: g-minor-major-7th-voicing-3
    start_fret: 2
    end_fret: 5
    notes:
    - string: 6
      fret: 3
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: '5'
      role: fifth
      finger: 3
    - string: 4
      fret: 0
      pc: 2
      note: D
      interval: '5'
      role: fifth
      finger: 0
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: b3
      role: third
      finger: 2
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '5'
      role: fifth
      finger: 2
    - string: 1
      fret: 2
      pc: 6
      note: F#
      interval: '7'
      role: seventh
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
playback:
  alphatex: \tempo 90 (3.6 1.5 0.4 0.3 3.2 2.1).2
  midi_root: 67
musical_forms:
- form: pitch_set
  id: g-minor-major-7th-chord-pitch-set
  pitch_classes:
  - 7
  - 10
  - 2
  - 6
  intervals:
  - '1'
  - b3
  - '5'
  - '7'
  spelling:
  - G
  - Bb
  - D
  - F#
- form: voicing
  id: g-minor-major-7th-voicing-1
  position_ref: g-minor-major-7th-voicing-1
  start_fret: 1
  tab: 3-1-0-0-3-2
- form: voicing
  id: g-minor-major-7th-voicing-2
  position_ref: g-minor-major-7th-voicing-2
  start_fret: 3
  tab: 3-5-4-3-3-3
- form: voicing
  id: g-minor-major-7th-voicing-3
  position_ref: g-minor-major-7th-voicing-3
  start_fret: 2
  tab: 3-5-0-3-3-2
- form: tab_excerpt
  id: g-minor-major-7th-chord-tab
  alphatex: \tempo 90 (3.6 1.5 0.4 0.3 3.2 2.1).2
---

# G Minor Major 7th Chord

The G minor major 7th chord is a minor triad plus a major seventh — a haunting, unstable color. Its tones are **G – Bb – D – F#**: G (root), Bb (minor third), D (perfect fifth), F# (major seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | G | root |
| b3 | Bb | minor third |
| 5 | D | perfect fifth |
| 7 | F# | major seventh |

**Playing it:** the primary shape is `3-1-0-0-3-2` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [G Minor Triad](/triads/g-minor-triad.md)
