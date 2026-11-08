---
type: Concept
slug: /chords/g-minor-6th-chord
title: G Minor 6th Chord
description: 'The minor 6th chord built on G: G, Bb, D, E.'
tags:
- chord
- minor-6th
- g
parent: /chords
prerequisites:
- /triads/g-minor-triad
theory:
  category: chord
  quality: minor-6th
  root:
    pc: 7
    name: G
  intervals:
  - '1'
  - b3
  - '5'
  - '6'
  semitones:
  - 0
  - 3
  - 7
  - 9
  pitch_classes:
  - 7
  - 10
  - 2
  - 4
  spelling:
  - G
  - Bb
  - D
  - E
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
  - 4
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  - '6'
  preferred_positions:
  - id: g-minor-6th-voicing-1
    start_fret: 1
    end_fret: 3
    notes:
    - string: 6
      fret: 3
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 2
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
      finger: 2
    - string: 1
      fret: 0
      pc: 4
      note: E
      interval: '6'
      role: sixth
      finger: 0
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: g-minor-6th-voicing-2
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
      finger: 2
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 2
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: b3
      role: third
      finger: 1
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '6'
      role: sixth
      finger: 2
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
  - id: g-minor-6th-voicing-3
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
      fret: 2
      pc: 4
      note: E
      interval: '6'
      role: sixth
      finger: 1
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
      fret: 0
      pc: 4
      note: E
      interval: '6'
      role: sixth
      finger: 0
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
playback:
  alphatex: \tempo 90 (3.6 1.5 0.4 0.3 3.2 0.1).2
  midi_root: 67
musical_forms:
- form: pitch_set
  id: g-minor-6th-chord-pitch-set
  pitch_classes:
  - 7
  - 10
  - 2
  - 4
  intervals:
  - '1'
  - b3
  - '5'
  - '6'
  spelling:
  - G
  - Bb
  - D
  - E
- form: voicing
  id: g-minor-6th-voicing-1
  position_ref: g-minor-6th-voicing-1
  start_fret: 1
  tab: 3-1-0-0-3-0
- form: voicing
  id: g-minor-6th-voicing-2
  position_ref: g-minor-6th-voicing-2
  start_fret: 3
  tab: 3-5-5-3-5-3
- form: voicing
  id: g-minor-6th-voicing-3
  position_ref: g-minor-6th-voicing-3
  start_fret: 2
  tab: 3-5-2-3-3-0
- form: tab_excerpt
  id: g-minor-6th-chord-tab
  alphatex: \tempo 90 (3.6 1.5 0.4 0.3 3.2 0.1).2
---

# G Minor 6th Chord

The G minor 6th chord is a minor triad plus a major sixth — a bright-edged minor color. Its tones are **G – Bb – D – E**: G (root), Bb (minor third), D (perfect fifth), E (major sixth).

| Interval | Note | Name |
|---|---|---|
| 1 | G | root |
| b3 | Bb | minor third |
| 5 | D | perfect fifth |
| 6 | E | major sixth |

**Playing it:** the primary shape is `3-1-0-0-3-0` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [G Minor Triad](/triads/g-minor-triad.md)
