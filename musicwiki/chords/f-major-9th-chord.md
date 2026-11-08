---
type: Concept
slug: /chords/f-major-9th-chord
title: F Major 9th Chord
description: 'The major 9th chord built on F: F, A, C, E, G.'
tags:
- chord
- major-9th
- f
parent: /chords
prerequisites:
- /chords/f-major-7th-chord
theory:
  category: chord
  quality: major-9th
  root:
    pc: 5
    name: F
  intervals:
  - '1'
  - '3'
  - '5'
  - '7'
  - '9'
  semitones:
  - 0
  - 4
  - 7
  - 11
  - 14
  pitch_classes:
  - 5
  - 9
  - 0
  - 4
  - 7
  spelling:
  - F
  - A
  - C
  - E
  - G
  enharmonic_policy: prefer_sharps
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
  - 4
  - 7
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - '7'
  - '9'
  preferred_positions:
  - id: f-major-9th-voicing-1
    start_fret: 1
    end_fret: 2
    notes:
    - string: 6
      fret: 1
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 0
      pc: 9
      note: A
      interval: '3'
      role: third
      finger: 0
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: '7'
      role: seventh
      finger: 2
    - string: 3
      fret: 0
      pc: 7
      note: G
      interval: '9'
      role: ninth
      finger: 0
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
    confidence: 0.85
    tags:
    - computed
    - open
    - barre
  - id: f-major-9th-voicing-2
    start_fret: 5
    end_fret: 8
    notes:
    - string: 5
      fret: 8
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: '9'
      role: ninth
      finger: 1
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
      fret: 5
      pc: 9
      note: A
      interval: '3'
      role: third
      finger: 1
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: f-major-9th-voicing-3
    start_fret: 7
    end_fret: 9
    notes:
    - string: 5
      fret: 8
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 7
      pc: 9
      note: A
      interval: '3'
      role: third
      finger: 1
    - string: 3
      fret: 9
      pc: 4
      note: E
      interval: '7'
      role: seventh
      finger: 3
    - string: 2
      fret: 8
      pc: 7
      note: G
      interval: '9'
      role: ninth
      finger: 2
    - string: 1
      fret: 8
      pc: 0
      note: C
      interval: '5'
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (1.6 0.5 2.4 0.3 1.2 0.1).2
  midi_root: 65
musical_forms:
- form: pitch_set
  id: f-major-9th-chord-pitch-set
  pitch_classes:
  - 5
  - 9
  - 0
  - 4
  - 7
  intervals:
  - '1'
  - '3'
  - '5'
  - '7'
  - '9'
  spelling:
  - F
  - A
  - C
  - E
  - G
- form: voicing
  id: f-major-9th-voicing-1
  position_ref: f-major-9th-voicing-1
  start_fret: 1
  tab: 1-0-2-0-1-0
- form: voicing
  id: f-major-9th-voicing-2
  position_ref: f-major-9th-voicing-2
  start_fret: 5
  tab: x-8-5-5-5-5
- form: voicing
  id: f-major-9th-voicing-3
  position_ref: f-major-9th-voicing-3
  start_fret: 7
  tab: x-8-7-9-8-8
- form: tab_excerpt
  id: f-major-9th-chord-tab
  alphatex: \tempo 90 (1.6 0.5 2.4 0.3 1.2 0.1).2
---

# F Major 9th Chord

The F major 9th chord is a major seventh extended with the ninth — expansive and warm. Its tones are **F – A – C – E – G**: F (root), A (major third), C (perfect fifth), E (major seventh), G (ninth).

| Interval | Note | Name |
|---|---|---|
| 1 | F | root |
| 3 | A | major third |
| 5 | C | perfect fifth |
| 7 | E | major seventh |
| 9 | G | ninth |

**Playing it:** the primary shape is `1-0-2-0-1-0` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [F Major 7th Chord](/chords/f-major-7th-chord.md)
