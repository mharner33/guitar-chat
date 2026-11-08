---
type: Concept
slug: /chords/b-flat-6th-chord
title: Bb 6th Chord
description: 'The 6th chord built on Bb: Bb, D, F, G.'
tags:
- chord
- 6th
- b
parent: /chords
prerequisites:
- /triads/b-flat-major-triad
theory:
  category: chord
  quality: 6th
  root:
    pc: 10
    name: Bb
  intervals:
  - '1'
  - '3'
  - '5'
  - '6'
  semitones:
  - 0
  - 4
  - 7
  - 9
  pitch_classes:
  - 10
  - 2
  - 5
  - 7
  spelling:
  - Bb
  - D
  - F
  - G
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
  - 5
  - 7
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - '6'
  preferred_positions:
  - id: b-flat-6th-voicing-1
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
      fret: 0
      pc: 7
      note: G
      interval: '6'
      role: sixth
      finger: 0
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '3'
      role: third
      finger: 2
    - string: 1
      fret: 1
      pc: 5
      note: F
      interval: '5'
      role: fifth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: b-flat-6th-voicing-2
    start_fret: 3
    end_fret: 6
    notes:
    - string: 6
      fret: 6
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: '3'
      role: third
      finger: 2
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '5'
      role: fifth
      finger: 1
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
      fret: 3
      pc: 7
      note: G
      interval: '6'
      role: sixth
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: b-flat-6th-voicing-3
    start_fret: 5
    end_fret: 7
    notes:
    - string: 6
      fret: 6
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: '3'
      role: third
      finger: 1
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: '6'
      role: sixth
      finger: 1
    - string: 3
      fret: 7
      pc: 2
      note: D
      interval: '3'
      role: third
      finger: 3
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '5'
      role: fifth
      finger: 2
    - string: 1
      fret: 6
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (1.5 0.4 0.3 3.2 1.1).2
  midi_root: 70
musical_forms:
- form: pitch_set
  id: b-flat-6th-chord-pitch-set
  pitch_classes:
  - 10
  - 2
  - 5
  - 7
  intervals:
  - '1'
  - '3'
  - '5'
  - '6'
  spelling:
  - Bb
  - D
  - F
  - G
- form: voicing
  id: b-flat-6th-voicing-1
  position_ref: b-flat-6th-voicing-1
  start_fret: 1
  tab: x-1-0-0-3-1
- form: voicing
  id: b-flat-6th-voicing-2
  position_ref: b-flat-6th-voicing-2
  start_fret: 3
  tab: 6-5-3-3-3-3
- form: voicing
  id: b-flat-6th-voicing-3
  position_ref: b-flat-6th-voicing-3
  start_fret: 5
  tab: 6-5-5-7-6-6
- form: tab_excerpt
  id: b-flat-6th-chord-tab
  alphatex: \tempo 90 (1.5 0.4 0.3 3.2 1.1).2
---

# Bb 6th Chord

The Bb 6th chord is a major triad plus a major sixth — a sweet, vintage tonic substitute. Its tones are **Bb – D – F – G**: Bb (root), D (major third), F (perfect fifth), G (major sixth).

| Interval | Note | Name |
|---|---|---|
| 1 | Bb | root |
| 3 | D | major third |
| 5 | F | perfect fifth |
| 6 | G | major sixth |

**Playing it:** the primary shape is `x-1-0-0-3-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Bb Major Triad](/triads/b-flat-major-triad.md)
