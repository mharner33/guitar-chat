---
type: Concept
slug: /chords/g-major-9th-chord
title: G Major 9th Chord
description: 'The major 9th chord built on G: G, B, D, F#, A.'
tags:
- chord
- major-9th
- g
parent: /chords
prerequisites:
- /chords/g-major-7th-chord
theory:
  category: chord
  quality: major-9th
  root:
    pc: 7
    name: G
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
  - 7
  - 11
  - 2
  - 6
  - 9
  spelling:
  - G
  - B
  - D
  - F#
  - A
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
  - 6
  - 9
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - '7'
  - '9'
  preferred_positions:
  - id: g-major-9th-voicing-1
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
      fret: 0
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 0
    - string: 2
      fret: 0
      pc: 11
      note: B
      interval: '3'
      role: third
      finger: 0
    - string: 1
      fret: 2
      pc: 6
      note: F#
      interval: '7'
      role: seventh
      finger: 1
    confidence: 0.85
    tags:
    - computed
    - open
  - id: g-major-9th-voicing-2
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
      fret: 4
      pc: 6
      note: F#
      interval: '7'
      role: seventh
      finger: 2
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '3'
      role: third
      finger: 2
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
      finger: 3
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: g-major-9th-voicing-3
    start_fret: 7
    end_fret: 10
    notes:
    - string: 5
      fret: 10
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 7
      pc: 9
      note: A
      interval: '9'
      role: ninth
      finger: 1
    - string: 3
      fret: 7
      pc: 2
      note: D
      interval: '5'
      role: fifth
      finger: 1
    - string: 2
      fret: 7
      pc: 6
      note: F#
      interval: '7'
      role: seventh
      finger: 1
    - string: 1
      fret: 7
      pc: 11
      note: B
      interval: '3'
      role: third
      finger: 1
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (3.6 0.5 0.4 0.3 0.2 2.1).2
  midi_root: 67
musical_forms:
- form: pitch_set
  id: g-major-9th-chord-pitch-set
  pitch_classes:
  - 7
  - 11
  - 2
  - 6
  - 9
  intervals:
  - '1'
  - '3'
  - '5'
  - '7'
  - '9'
  spelling:
  - G
  - B
  - D
  - F#
  - A
- form: voicing
  id: g-major-9th-voicing-1
  position_ref: g-major-9th-voicing-1
  start_fret: 2
  tab: 3-0-0-0-0-2
- form: voicing
  id: g-major-9th-voicing-2
  position_ref: g-major-9th-voicing-2
  start_fret: 3
  tab: 3-5-4-4-3-5
- form: voicing
  id: g-major-9th-voicing-3
  position_ref: g-major-9th-voicing-3
  start_fret: 7
  tab: x-10-7-7-7-7
- form: tab_excerpt
  id: g-major-9th-chord-tab
  alphatex: \tempo 90 (3.6 0.5 0.4 0.3 0.2 2.1).2
---

# G Major 9th Chord

The G major 9th chord is a major seventh extended with the ninth — expansive and warm. Its tones are **G – B – D – F# – A**: G (root), B (major third), D (perfect fifth), F# (major seventh), A (ninth).

| Interval | Note | Name |
|---|---|---|
| 1 | G | root |
| 3 | B | major third |
| 5 | D | perfect fifth |
| 7 | F# | major seventh |
| 9 | A | ninth |

**Playing it:** the primary shape is `3-0-0-0-0-2` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [G Major 7th Chord](/chords/g-major-7th-chord.md)
