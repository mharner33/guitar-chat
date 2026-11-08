---
type: Concept
slug: /chords/f-minor-major-7th-chord
title: F Minor Major 7th Chord
description: 'The minor major 7th chord built on F: F, Ab, C, E.'
tags:
- chord
- minor-major-7th
- f
parent: /chords
prerequisites:
- /triads/f-minor-triad
theory:
  category: chord
  quality: minor-major-7th
  root:
    pc: 5
    name: F
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
  - 5
  - 8
  - 0
  - 4
  spelling:
  - F
  - Ab
  - C
  - E
  enharmonic_policy: prefer_flats
  key_context: F
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
  - 5
  - 8
  - 0
  - 4
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  - '7'
  preferred_positions:
  - id: f-minor-major-7th-voicing-1
    start_fret: 1
    end_fret: 3
    notes:
    - string: 6
      fret: 1
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 3
      pc: 0
      note: C
      interval: '5'
      role: fifth
      finger: 3
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: '7'
      role: seventh
      finger: 2
    - string: 3
      fret: 1
      pc: 8
      note: Ab
      interval: b3
      role: third
      finger: 1
    - string: 2
      fret: 1
      pc: 0
      note: C
      interval: '5'
      role: fifth
      finger: 1
    - string: 1
      fret: 0
      pc: 4
      note: E
      interval: '7'
      role: seventh
      finger: 0
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: f-minor-major-7th-voicing-2
    start_fret: 3
    end_fret: 5
    notes:
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: '5'
      role: fifth
      finger: 3
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '7'
      role: seventh
      finger: 3
    - string: 1
      fret: 4
      pc: 8
      note: Ab
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
  - id: f-minor-major-7th-voicing-3
    start_fret: 5
    end_fret: 8
    notes:
    - string: 5
      fret: 8
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 6
      pc: 8
      note: Ab
      interval: b3
      role: third
      finger: 2
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: '5'
      role: fifth
      finger: 1
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '7'
      role: seventh
      finger: 1
    - string: 1
      fret: 8
      pc: 0
      note: C
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
  alphatex: \tempo 90 (1.6 3.5 2.4 1.3 1.2 0.1).2
  midi_root: 65
musical_forms:
- form: pitch_set
  id: f-minor-major-7th-chord-pitch-set
  pitch_classes:
  - 5
  - 8
  - 0
  - 4
  intervals:
  - '1'
  - b3
  - '5'
  - '7'
  spelling:
  - F
  - Ab
  - C
  - E
- form: voicing
  id: f-minor-major-7th-voicing-1
  position_ref: f-minor-major-7th-voicing-1
  start_fret: 1
  tab: 1-3-2-1-1-0
- form: voicing
  id: f-minor-major-7th-voicing-2
  position_ref: f-minor-major-7th-voicing-2
  start_fret: 3
  tab: x-x-3-5-5-4
- form: voicing
  id: f-minor-major-7th-voicing-3
  position_ref: f-minor-major-7th-voicing-3
  start_fret: 5
  tab: x-8-6-5-5-8
- form: tab_excerpt
  id: f-minor-major-7th-chord-tab
  alphatex: \tempo 90 (1.6 3.5 2.4 1.3 1.2 0.1).2
---

# F Minor Major 7th Chord

The F minor major 7th chord is a minor triad plus a major seventh — a haunting, unstable color. Its tones are **F – Ab – C – E**: F (root), Ab (minor third), C (perfect fifth), E (major seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | F | root |
| b3 | Ab | minor third |
| 5 | C | perfect fifth |
| 7 | E | major seventh |

**Playing it:** the primary shape is `1-3-2-1-1-0` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [F Minor Triad](/triads/f-minor-triad.md)
