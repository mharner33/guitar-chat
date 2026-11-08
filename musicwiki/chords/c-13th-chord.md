---
type: Concept
slug: /chords/c-13th-chord
title: C 13th Chord
description: 'The 13th chord built on C: C, E, G, Bb, D, A.'
tags:
- chord
- 13th
- c
parent: /chords
prerequisites:
- /chords/c-dominant-7th-chord
theory:
  category: chord
  quality: 13th
  root:
    pc: 0
    name: C
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
  - 0
  - 4
  - 7
  - 10
  - 2
  - 9
  spelling:
  - C
  - E
  - G
  - Bb
  - D
  - A
  enharmonic_policy: prefer_flats
  key_context: C
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
  - 0
  - 4
  - 7
  - 10
  - 2
  - 9
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '13'
  preferred_positions:
  - id: c-13th-voicing-1
    start_fret: 6
    end_fret: 8
    notes:
    - string: 6
      fret: 8
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 7
      pc: 4
      note: E
      interval: '3'
      role: third
      finger: 2
    - string: 4
      fret: 7
      pc: 9
      note: A
      interval: '13'
      role: thirteenth
      finger: 2
    - string: 3
      fret: 7
      pc: 2
      note: D
      interval: '9'
      role: ninth
      finger: 2
    - string: 2
      fret: 8
      pc: 7
      note: G
      interval: '5'
      role: fifth
      finger: 3
    - string: 1
      fret: 6
      pc: 10
      note: Bb
      interval: b7
      role: seventh
      finger: 1
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: c-13th-voicing-2
    start_fret: 5
    end_fret: 8
    notes:
    - string: 6
      fret: 8
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 7
      pc: 4
      note: E
      interval: '3'
      role: third
      finger: 2
    - string: 4
      fret: 8
      pc: 10
      note: Bb
      interval: b7
      role: seventh
      finger: 3
    - string: 3
      fret: 7
      pc: 2
      note: D
      interval: '9'
      role: ninth
      finger: 2
    - string: 2
      fret: 8
      pc: 7
      note: G
      interval: '5'
      role: fifth
      finger: 3
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '13'
      role: thirteenth
      finger: 1
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: c-13th-voicing-3
    start_fret: 8
    end_fret: 10
    notes:
    - string: 6
      fret: 8
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 10
      pc: 7
      note: G
      interval: '5'
      role: fifth
      finger: 3
    - string: 4
      fret: 8
      pc: 10
      note: Bb
      interval: b7
      role: seventh
      finger: 1
    - string: 3
      fret: 9
      pc: 4
      note: E
      interval: '3'
      role: third
      finger: 2
    - string: 2
      fret: 10
      pc: 9
      note: A
      interval: '13'
      role: thirteenth
      finger: 3
    - string: 1
      fret: 10
      pc: 2
      note: D
      interval: '9'
      role: ninth
      finger: 3
    confidence: 0.85
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (8.6 7.5 7.4 7.3 8.2 6.1).2
  midi_root: 60
musical_forms:
- form: pitch_set
  id: c-13th-chord-pitch-set
  pitch_classes:
  - 0
  - 4
  - 7
  - 10
  - 2
  - 9
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '13'
  spelling:
  - C
  - E
  - G
  - Bb
  - D
  - A
- form: voicing
  id: c-13th-voicing-1
  position_ref: c-13th-voicing-1
  start_fret: 6
  tab: 8-7-7-7-8-6
- form: voicing
  id: c-13th-voicing-2
  position_ref: c-13th-voicing-2
  start_fret: 5
  tab: 8-7-8-7-8-5
- form: voicing
  id: c-13th-voicing-3
  position_ref: c-13th-voicing-3
  start_fret: 8
  tab: 8-10-8-9-10-10
- form: tab_excerpt
  id: c-13th-chord-tab
  alphatex: \tempo 90 (8.6 7.5 7.4 7.3 8.2 6.1).2
---

# C 13th Chord

The C 13th chord is a dominant chord extended through the thirteenth — the fullest dominant color. Its tones are **C – E – G – Bb – D – A**: C (root), E (major third), G (perfect fifth), Bb (minor seventh), D (ninth), A (thirteenth).

| Interval | Note | Name |
|---|---|---|
| 1 | C | root |
| 3 | E | major third |
| 5 | G | perfect fifth |
| b7 | Bb | minor seventh |
| 9 | D | ninth |
| 13 | A | thirteenth |

**Playing it:** the primary shape is `8-7-7-7-8-6` (strings low→high, `x` = muted), around fret 6. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [C Dominant 7th Chord](/chords/c-dominant-7th-chord.md)
