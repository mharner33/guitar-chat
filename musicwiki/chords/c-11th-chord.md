---
type: Concept
slug: /chords/c-11th-chord
title: C 11th Chord
description: 'The 11th chord built on C: C, E, G, Bb, D, F.'
tags:
- chord
- 11th
- c
parent: /chords
prerequisites:
- /chords/c-dominant-7th-chord
theory:
  category: chord
  quality: 11th
  root:
    pc: 0
    name: C
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '11'
  semitones:
  - 0
  - 4
  - 7
  - 10
  - 14
  - 17
  pitch_classes:
  - 0
  - 4
  - 7
  - 10
  - 2
  - 5
  spelling:
  - C
  - E
  - G
  - Bb
  - D
  - F
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
  - 5
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '11'
  preferred_positions:
  - id: c-11th-voicing-1
    start_fret: 1
    end_fret: 3
    notes:
    - string: 5
      fret: 3
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: '3'
      role: third
      finger: 2
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: b7
      role: seventh
      finger: 3
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '9'
      role: ninth
      finger: 3
    - string: 1
      fret: 1
      pc: 5
      note: F
      interval: '11'
      role: eleventh
      finger: 1
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: c-11th-voicing-2
    start_fret: 5
    end_fret: 8
    notes:
    - string: 6
      fret: 8
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 4
    - string: 5
      fret: 7
      pc: 4
      note: E
      interval: '3'
      role: third
      finger: 3
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: '5'
      role: fifth
      finger: 1
    - string: 3
      fret: 7
      pc: 2
      note: D
      interval: '9'
      role: ninth
      finger: 3
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '11'
      role: eleventh
      finger: 2
    - string: 1
      fret: 6
      pc: 10
      note: Bb
      interval: b7
      role: seventh
      finger: 2
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: c-11th-voicing-3
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
      fret: 8
      pc: 5
      note: F
      interval: '11'
      role: eleventh
      finger: 1
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
      fret: 8
      pc: 7
      note: G
      interval: '5'
      role: fifth
      finger: 1
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
  alphatex: \tempo 90 (3.5 2.4 3.3 3.2 1.1).2
  midi_root: 60
musical_forms:
- form: pitch_set
  id: c-11th-chord-pitch-set
  pitch_classes:
  - 0
  - 4
  - 7
  - 10
  - 2
  - 5
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '11'
  spelling:
  - C
  - E
  - G
  - Bb
  - D
  - F
- form: voicing
  id: c-11th-voicing-1
  position_ref: c-11th-voicing-1
  start_fret: 1
  tab: x-3-2-3-3-1
- form: voicing
  id: c-11th-voicing-2
  position_ref: c-11th-voicing-2
  start_fret: 5
  tab: 8-7-5-7-6-6
- form: voicing
  id: c-11th-voicing-3
  position_ref: c-11th-voicing-3
  start_fret: 8
  tab: 8-8-8-9-8-10
- form: tab_excerpt
  id: c-11th-chord-tab
  alphatex: \tempo 90 (3.5 2.4 3.3 3.2 1.1).2
---

# C 11th Chord

The C 11th chord is a dominant chord extended through the eleventh. Its tones are **C – E – G – Bb – D – F**: C (root), E (major third), G (perfect fifth), Bb (minor seventh), D (ninth), F (eleventh).

| Interval | Note | Name |
|---|---|---|
| 1 | C | root |
| 3 | E | major third |
| 5 | G | perfect fifth |
| b7 | Bb | minor seventh |
| 9 | D | ninth |
| 11 | F | eleventh |

**Playing it:** the primary shape is `x-3-2-3-3-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [C Dominant 7th Chord](/chords/c-dominant-7th-chord.md)
