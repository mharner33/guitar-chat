---
type: Concept
slug: /chords/g-minor-9th-chord
title: G Minor 9th Chord
description: 'The minor 9th chord built on G: G, Bb, D, F, A.'
tags:
- chord
- minor-9th
- g
parent: /chords
prerequisites:
- /chords/g-minor-7th-chord
theory:
  category: chord
  quality: minor-9th
  root:
    pc: 7
    name: G
  intervals:
  - '1'
  - b3
  - '5'
  - b7
  - '9'
  semitones:
  - 0
  - 3
  - 7
  - 10
  - 14
  pitch_classes:
  - 7
  - 10
  - 2
  - 5
  - 9
  spelling:
  - G
  - Bb
  - D
  - F
  - A
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
  - 5
  - 9
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  - b7
  - '9'
  preferred_positions:
  - id: g-minor-9th-voicing-1
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
      fret: 0
      pc: 9
      note: A
      interval: '9'
      role: ninth
      finger: 0
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
      fret: 1
      pc: 5
      note: F
      interval: b7
      role: seventh
      finger: 1
    confidence: 0.85
    tags:
    - computed
    - open
    - barre
  - id: g-minor-9th-voicing-2
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
      fret: 3
      pc: 5
      note: F
      interval: b7
      role: seventh
      finger: 1
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
      fret: 5
      pc: 9
      note: A
      interval: '9'
      role: ninth
      finger: 2
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: g-minor-9th-voicing-3
    start_fret: 2
    end_fret: 6
    notes:
    - string: 6
      fret: 3
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 0
      pc: 9
      note: A
      interval: '9'
      role: ninth
      finger: 0
    - string: 4
      fret: 0
      pc: 2
      note: D
      interval: '5'
      role: fifth
      finger: 0
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: '9'
      role: ninth
      finger: 1
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: b7
      role: seventh
      finger: 3
    - string: 1
      fret: 6
      pc: 10
      note: Bb
      interval: b3
      role: third
      finger: 3
    confidence: 0.85
    tags:
    - computed
    - open
    - barre
playback:
  alphatex: \tempo 90 (3.6 0.5 0.4 3.3 3.2 1.1).2
  midi_root: 67
musical_forms:
- form: pitch_set
  id: g-minor-9th-chord-pitch-set
  pitch_classes:
  - 7
  - 10
  - 2
  - 5
  - 9
  intervals:
  - '1'
  - b3
  - '5'
  - b7
  - '9'
  spelling:
  - G
  - Bb
  - D
  - F
  - A
- form: voicing
  id: g-minor-9th-voicing-1
  position_ref: g-minor-9th-voicing-1
  start_fret: 1
  tab: 3-0-0-3-3-1
- form: voicing
  id: g-minor-9th-voicing-2
  position_ref: g-minor-9th-voicing-2
  start_fret: 3
  tab: 3-5-3-3-3-5
- form: voicing
  id: g-minor-9th-voicing-3
  position_ref: g-minor-9th-voicing-3
  start_fret: 2
  tab: 3-0-0-2-6-6
- form: tab_excerpt
  id: g-minor-9th-chord-tab
  alphatex: \tempo 90 (3.6 0.5 0.4 3.3 3.2 1.1).2
---

# G Minor 9th Chord

The G minor 9th chord is a minor seventh extended with the ninth — smooth and modern. Its tones are **G – Bb – D – F – A**: G (root), Bb (minor third), D (perfect fifth), F (minor seventh), A (ninth).

| Interval | Note | Name |
|---|---|---|
| 1 | G | root |
| b3 | Bb | minor third |
| 5 | D | perfect fifth |
| b7 | F | minor seventh |
| 9 | A | ninth |

**Playing it:** the primary shape is `3-0-0-3-3-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [G Minor 7th Chord](/chords/g-minor-7th-chord.md)
