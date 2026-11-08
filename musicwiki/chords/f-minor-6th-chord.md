---
type: Concept
slug: /chords/f-minor-6th-chord
title: F Minor 6th Chord
description: 'The minor 6th chord built on F: F, Ab, C, D.'
tags:
- chord
- minor-6th
- f
parent: /chords
prerequisites:
- /triads/f-minor-triad
theory:
  category: chord
  quality: minor-6th
  root:
    pc: 5
    name: F
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
  - 5
  - 8
  - 0
  - 2
  spelling:
  - F
  - Ab
  - C
  - D
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
  - 2
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  - '6'
  preferred_positions:
  - id: f-minor-6th-voicing-1
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
      finger: 2
    - string: 4
      fret: 0
      pc: 2
      note: D
      interval: '6'
      role: sixth
      finger: 0
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
      fret: 1
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: f-minor-6th-voicing-2
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
      fret: 3
      pc: 2
      note: D
      interval: '6'
      role: sixth
      finger: 1
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
  - id: f-minor-6th-voicing-3
    start_fret: 6
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
      finger: 1
    - string: 3
      fret: 7
      pc: 2
      note: D
      interval: '6'
      role: sixth
      finger: 2
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '1'
      role: root
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
  alphatex: \tempo 90 (1.6 3.5 0.4 1.3 1.2 1.1).2
  midi_root: 65
musical_forms:
- form: pitch_set
  id: f-minor-6th-chord-pitch-set
  pitch_classes:
  - 5
  - 8
  - 0
  - 2
  intervals:
  - '1'
  - b3
  - '5'
  - '6'
  spelling:
  - F
  - Ab
  - C
  - D
- form: voicing
  id: f-minor-6th-voicing-1
  position_ref: f-minor-6th-voicing-1
  start_fret: 1
  tab: 1-3-0-1-1-1
- form: voicing
  id: f-minor-6th-voicing-2
  position_ref: f-minor-6th-voicing-2
  start_fret: 3
  tab: x-x-3-5-3-4
- form: voicing
  id: f-minor-6th-voicing-3
  position_ref: f-minor-6th-voicing-3
  start_fret: 6
  tab: x-8-6-7-6-8
- form: tab_excerpt
  id: f-minor-6th-chord-tab
  alphatex: \tempo 90 (1.6 3.5 0.4 1.3 1.2 1.1).2
---

# F Minor 6th Chord

The F minor 6th chord is a minor triad plus a major sixth — a bright-edged minor color. Its tones are **F – Ab – C – D**: F (root), Ab (minor third), C (perfect fifth), D (major sixth).

| Interval | Note | Name |
|---|---|---|
| 1 | F | root |
| b3 | Ab | minor third |
| 5 | C | perfect fifth |
| 6 | D | major sixth |

**Playing it:** the primary shape is `1-3-0-1-1-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [F Minor Triad](/triads/f-minor-triad.md)
