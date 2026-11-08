---
type: Concept
slug: /chords/e-flat-minor-major-7th-chord
title: Eb Minor Major 7th Chord
description: 'The minor major 7th chord built on Eb: Eb, Gb, Bb, D.'
tags:
- chord
- minor-major-7th
- e
parent: /chords
prerequisites:
- /triads/e-flat-minor-triad
theory:
  category: chord
  quality: minor-major-7th
  root:
    pc: 3
    name: Eb
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
  - 3
  - 6
  - 10
  - 2
  spelling:
  - Eb
  - Gb
  - Bb
  - D
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
  - 6
  - 10
  - 2
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  - '7'
  preferred_positions:
  - id: e-flat-minor-major-7th-voicing-1
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
      fret: 3
      pc: 10
      note: Bb
      interval: '5'
      role: fifth
      finger: 3
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '7'
      role: seventh
      finger: 3
    - string: 1
      fret: 2
      pc: 6
      note: Gb
      interval: b3
      role: third
      finger: 2
    mute_strings:
    - 5
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: e-flat-minor-major-7th-voicing-2
    start_fret: 3
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
      fret: 4
      pc: 6
      note: Gb
      interval: b3
      role: third
      finger: 2
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '5'
      role: fifth
      finger: 1
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '7'
      role: seventh
      finger: 1
    - string: 1
      fret: 6
      pc: 10
      note: Bb
      interval: '5'
      role: fifth
      finger: 3
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: e-flat-minor-major-7th-voicing-3
    start_fret: 2
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
      fret: 0
      pc: 2
      note: D
      interval: '7'
      role: seventh
      finger: 0
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '5'
      role: fifth
      finger: 2
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '7'
      role: seventh
      finger: 2
    - string: 1
      fret: 2
      pc: 6
      note: Gb
      interval: b3
      role: third
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
playback:
  alphatex: \tempo 90 (1.4 3.3 3.2 2.1).2
  midi_root: 63
musical_forms:
- form: pitch_set
  id: e-flat-minor-major-7th-chord-pitch-set
  pitch_classes:
  - 3
  - 6
  - 10
  - 2
  intervals:
  - '1'
  - b3
  - '5'
  - '7'
  spelling:
  - Eb
  - Gb
  - Bb
  - D
- form: voicing
  id: e-flat-minor-major-7th-voicing-1
  position_ref: e-flat-minor-major-7th-voicing-1
  start_fret: 1
  tab: x-x-1-3-3-2
- form: voicing
  id: e-flat-minor-major-7th-voicing-2
  position_ref: e-flat-minor-major-7th-voicing-2
  start_fret: 3
  tab: x-6-4-3-3-6
- form: voicing
  id: e-flat-minor-major-7th-voicing-3
  position_ref: e-flat-minor-major-7th-voicing-3
  start_fret: 2
  tab: x-6-0-3-3-2
- form: tab_excerpt
  id: e-flat-minor-major-7th-chord-tab
  alphatex: \tempo 90 (1.4 3.3 3.2 2.1).2
---

# Eb Minor Major 7th Chord

The Eb minor major 7th chord is a minor triad plus a major seventh — a haunting, unstable color. Its tones are **Eb – Gb – Bb – D**: Eb (root), Gb (minor third), Bb (perfect fifth), D (major seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | Eb | root |
| b3 | Gb | minor third |
| 5 | Bb | perfect fifth |
| 7 | D | major seventh |

**Playing it:** the primary shape is `x-x-1-3-3-2` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Eb Minor Triad](/triads/e-flat-minor-triad.md)
