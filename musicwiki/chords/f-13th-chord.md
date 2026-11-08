---
type: Concept
slug: /chords/f-13th-chord
title: F 13th Chord
description: 'The 13th chord built on F: F, A, C, Eb, G, D.'
tags:
- chord
- 13th
- f
parent: /chords
prerequisites:
- /chords/f-dominant-7th-chord
theory:
  category: chord
  quality: 13th
  root:
    pc: 5
    name: F
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '13'
  semitones:
  - 0
  - 4
  - 7
  - 10
  - 14
  - 21
  pitch_classes:
  - 5
  - 9
  - 0
  - 3
  - 7
  - 2
  spelling:
  - F
  - A
  - C
  - Eb
  - G
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
  - 9
  - 0
  - 3
  - 7
  - 2
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '13'
  preferred_positions:
  - id: f-13th-voicing-1
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
      fret: 1
      pc: 3
      note: Eb
      interval: b7
      role: seventh
      finger: 1
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: '3'
      role: third
      finger: 2
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '13'
      role: thirteenth
      finger: 3
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: '9'
      role: ninth
      finger: 3
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: f-13th-voicing-2
    start_fret: 4
    end_fret: 8
    notes:
    - string: 5
      fret: 8
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 4
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: '9'
      role: ninth
      finger: 2
    - string: 3
      fret: 7
      pc: 2
      note: D
      interval: '13'
      role: thirteenth
      finger: 3
    - string: 2
      fret: 4
      pc: 3
      note: Eb
      interval: b7
      role: seventh
      finger: 1
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '3'
      role: third
      finger: 2
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: f-13th-voicing-3
    start_fret: 8
    end_fret: 10
    notes:
    - string: 5
      fret: 8
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 10
      pc: 0
      note: C
      interval: '5'
      role: fifth
      finger: 2
    - string: 3
      fret: 8
      pc: 3
      note: Eb
      interval: b7
      role: seventh
      finger: 1
    - string: 2
      fret: 10
      pc: 9
      note: A
      interval: '3'
      role: third
      finger: 2
    - string: 1
      fret: 10
      pc: 2
      note: D
      interval: '13'
      role: thirteenth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (1.6 3.5 1.4 2.3 3.2 3.1).2
  midi_root: 65
musical_forms:
- form: pitch_set
  id: f-13th-chord-pitch-set
  pitch_classes:
  - 5
  - 9
  - 0
  - 3
  - 7
  - 2
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '13'
  spelling:
  - F
  - A
  - C
  - Eb
  - G
  - D
- form: voicing
  id: f-13th-voicing-1
  position_ref: f-13th-voicing-1
  start_fret: 1
  tab: 1-3-1-2-3-3
- form: voicing
  id: f-13th-voicing-2
  position_ref: f-13th-voicing-2
  start_fret: 4
  tab: x-8-5-7-4-5
- form: voicing
  id: f-13th-voicing-3
  position_ref: f-13th-voicing-3
  start_fret: 8
  tab: x-8-10-8-10-10
- form: tab_excerpt
  id: f-13th-chord-tab
  alphatex: \tempo 90 (1.6 3.5 1.4 2.3 3.2 3.1).2
---

# F 13th Chord

The F 13th chord is a dominant chord extended through the thirteenth — the fullest dominant color. Its tones are **F – A – C – Eb – G – D**: F (root), A (major third), C (perfect fifth), Eb (minor seventh), G (ninth), D (thirteenth).

| Interval | Note | Name |
|---|---|---|
| 1 | F | root |
| 3 | A | major third |
| 5 | C | perfect fifth |
| b7 | Eb | minor seventh |
| 9 | G | ninth |
| 13 | D | thirteenth |

**Playing it:** the primary shape is `1-3-1-2-3-3` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [F Dominant 7th Chord](/chords/f-dominant-7th-chord.md)
