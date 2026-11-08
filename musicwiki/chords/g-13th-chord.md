---
type: Concept
slug: /chords/g-13th-chord
title: G 13th Chord
description: 'The 13th chord built on G: G, B, D, F, A, E.'
tags:
- chord
- 13th
- g
parent: /chords
prerequisites:
- /chords/g-dominant-7th-chord
theory:
  category: chord
  quality: 13th
  root:
    pc: 7
    name: G
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
  - 7
  - 11
  - 2
  - 5
  - 9
  - 4
  spelling:
  - G
  - B
  - D
  - F
  - A
  - E
  enharmonic_policy: prefer_sharps
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
  - 11
  - 2
  - 5
  - 9
  - 4
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '13'
  preferred_positions:
  - id: g-13th-voicing-1
    start_fret: 2
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
      fret: 2
      pc: 11
      note: B
      interval: '3'
      role: third
      finger: 1
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: b7
      role: seventh
      finger: 2
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: '9'
      role: ninth
      finger: 1
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
      interval: '13'
      role: thirteenth
      finger: 0
    confidence: 0.85
    tags:
    - computed
    - open
    - barre
  - id: g-13th-voicing-2
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
      fret: 2
      pc: 11
      note: B
      interval: '3'
      role: third
      finger: 2
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: '13'
      role: thirteenth
      finger: 2
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: '9'
      role: ninth
      finger: 2
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '5'
      role: fifth
      finger: 3
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
    - barre
  - id: g-13th-voicing-3
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
      fret: 3
      pc: 5
      note: F
      interval: b7
      role: seventh
      finger: 1
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '3'
      role: third
      finger: 2
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '13'
      role: thirteenth
      finger: 3
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '9'
      role: ninth
      finger: 3
    confidence: 0.85
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (3.6 2.5 3.4 2.3 3.2 0.1).2
  midi_root: 67
musical_forms:
- form: pitch_set
  id: g-13th-chord-pitch-set
  pitch_classes:
  - 7
  - 11
  - 2
  - 5
  - 9
  - 4
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '13'
  spelling:
  - G
  - B
  - D
  - F
  - A
  - E
- form: voicing
  id: g-13th-voicing-1
  position_ref: g-13th-voicing-1
  start_fret: 2
  tab: 3-2-3-2-3-0
- form: voicing
  id: g-13th-voicing-2
  position_ref: g-13th-voicing-2
  start_fret: 1
  tab: 3-2-2-2-3-1
- form: voicing
  id: g-13th-voicing-3
  position_ref: g-13th-voicing-3
  start_fret: 3
  tab: 3-5-3-4-5-5
- form: tab_excerpt
  id: g-13th-chord-tab
  alphatex: \tempo 90 (3.6 2.5 3.4 2.3 3.2 0.1).2
---

# G 13th Chord

The G 13th chord is a dominant chord extended through the thirteenth — the fullest dominant color. Its tones are **G – B – D – F – A – E**: G (root), B (major third), D (perfect fifth), F (minor seventh), A (ninth), E (thirteenth).

| Interval | Note | Name |
|---|---|---|
| 1 | G | root |
| 3 | B | major third |
| 5 | D | perfect fifth |
| b7 | F | minor seventh |
| 9 | A | ninth |
| 13 | E | thirteenth |

**Playing it:** the primary shape is `3-2-3-2-3-0` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [G Dominant 7th Chord](/chords/g-dominant-7th-chord.md)
