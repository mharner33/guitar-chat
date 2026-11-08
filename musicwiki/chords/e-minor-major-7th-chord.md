---
type: Concept
slug: /chords/e-minor-major-7th-chord
title: E Minor Major 7th Chord
description: 'The minor major 7th chord built on E: E, G, B, D#.'
tags:
- chord
- minor-major-7th
- e
parent: /chords
prerequisites:
- /triads/e-minor-triad
theory:
  category: chord
  quality: minor-major-7th
  root:
    pc: 4
    name: E
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
  - 4
  - 7
  - 11
  - 3
  spelling:
  - E
  - G
  - B
  - D#
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
  - 7
  - 11
  - 3
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  - '7'
  preferred_positions:
  - id: e-minor-major-7th-voicing-1
    start_fret: 1
    end_fret: 2
    notes:
    - string: 6
      fret: 0
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 0
    - string: 5
      fret: 2
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 2
    - string: 4
      fret: 1
      pc: 3
      note: D#
      interval: '7'
      role: seventh
      finger: 1
    - string: 3
      fret: 0
      pc: 7
      note: G
      interval: b3
      role: third
      finger: 0
    - string: 2
      fret: 0
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 0
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
  - id: e-minor-major-7th-voicing-2
    start_fret: 2
    end_fret: 4
    notes:
    - string: 6
      fret: 0
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 0
    - string: 5
      fret: 2
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 1
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 0
      pc: 7
      note: G
      interval: b3
      role: third
      finger: 0
    - string: 2
      fret: 4
      pc: 3
      note: D#
      interval: '7'
      role: seventh
      finger: 2
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
  - id: e-minor-major-7th-voicing-3
    start_fret: 4
    end_fret: 7
    notes:
    - string: 5
      fret: 7
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: b3
      role: third
      finger: 2
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 1
    - string: 2
      fret: 4
      pc: 3
      note: D#
      interval: '7'
      role: seventh
      finger: 1
    - string: 1
      fret: 7
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 3
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (0.6 2.5 1.4 0.3 0.2 0.1).2
  midi_root: 64
musical_forms:
- form: pitch_set
  id: e-minor-major-7th-chord-pitch-set
  pitch_classes:
  - 4
  - 7
  - 11
  - 3
  intervals:
  - '1'
  - b3
  - '5'
  - '7'
  spelling:
  - E
  - G
  - B
  - D#
- form: voicing
  id: e-minor-major-7th-voicing-1
  position_ref: e-minor-major-7th-voicing-1
  start_fret: 1
  tab: 0-2-1-0-0-0
- form: voicing
  id: e-minor-major-7th-voicing-2
  position_ref: e-minor-major-7th-voicing-2
  start_fret: 2
  tab: 0-2-2-0-4-0
- form: voicing
  id: e-minor-major-7th-voicing-3
  position_ref: e-minor-major-7th-voicing-3
  start_fret: 4
  tab: x-7-5-4-4-7
- form: tab_excerpt
  id: e-minor-major-7th-chord-tab
  alphatex: \tempo 90 (0.6 2.5 1.4 0.3 0.2 0.1).2
---

# E Minor Major 7th Chord

The E minor major 7th chord is a minor triad plus a major seventh — a haunting, unstable color. Its tones are **E – G – B – D#**: E (root), G (minor third), B (perfect fifth), D# (major seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | E | root |
| b3 | G | minor third |
| 5 | B | perfect fifth |
| 7 | D# | major seventh |

**Playing it:** the primary shape is `0-2-1-0-0-0` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [E Minor Triad](/triads/e-minor-triad.md)
