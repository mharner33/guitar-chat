---
type: Concept
slug: /chords/c-major-9th-chord
title: C Major 9th Chord
description: 'The major 9th chord built on C: C, E, G, B, D.'
tags:
- chord
- major-9th
- c
parent: /chords
prerequisites:
- /chords/c-major-7th-chord
theory:
  category: chord
  quality: major-9th
  root:
    pc: 0
    name: C
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
  - 0
  - 4
  - 7
  - 11
  - 2
  spelling:
  - C
  - E
  - G
  - B
  - D
  enharmonic_policy: prefer_sharps
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
  - 11
  - 2
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - '7'
  - '9'
  preferred_positions:
  - id: c-major-9th-voicing-1
    start_fret: 2
    end_fret: 4
    notes:
    - string: 5
      fret: 3
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: '3'
      role: third
      finger: 1
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '7'
      role: seventh
      finger: 3
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '9'
      role: ninth
      finger: 2
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: '5'
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: c-major-9th-voicing-2
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
      fret: 5
      pc: 2
      note: D
      interval: '9'
      role: ninth
      finger: 1
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: '5'
      role: fifth
      finger: 1
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 1
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '3'
      role: third
      finger: 1
    - string: 1
      fret: 7
      pc: 11
      note: B
      interval: '7'
      role: seventh
      finger: 2
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: c-major-9th-voicing-3
    start_fret: 7
    end_fret: 9
    notes:
    - string: 6
      fret: 8
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 7
      pc: 4
      note: E
      interval: '3'
      role: third
      finger: 1
    - string: 4
      fret: 9
      pc: 11
      note: B
      interval: '7'
      role: seventh
      finger: 3
    - string: 3
      fret: 7
      pc: 2
      note: D
      interval: '9'
      role: ninth
      finger: 1
    - string: 2
      fret: 8
      pc: 7
      note: G
      interval: '5'
      role: fifth
      finger: 2
    - string: 1
      fret: 7
      pc: 11
      note: B
      interval: '7'
      role: seventh
      finger: 1
    confidence: 0.85
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (3.5 2.4 4.3 3.2 3.1).2
  midi_root: 60
musical_forms:
- form: pitch_set
  id: c-major-9th-chord-pitch-set
  pitch_classes:
  - 0
  - 4
  - 7
  - 11
  - 2
  intervals:
  - '1'
  - '3'
  - '5'
  - '7'
  - '9'
  spelling:
  - C
  - E
  - G
  - B
  - D
- form: voicing
  id: c-major-9th-voicing-1
  position_ref: c-major-9th-voicing-1
  start_fret: 2
  tab: x-3-2-4-3-3
- form: voicing
  id: c-major-9th-voicing-2
  position_ref: c-major-9th-voicing-2
  start_fret: 5
  tab: 8-5-5-5-5-7
- form: voicing
  id: c-major-9th-voicing-3
  position_ref: c-major-9th-voicing-3
  start_fret: 7
  tab: 8-7-9-7-8-7
- form: tab_excerpt
  id: c-major-9th-chord-tab
  alphatex: \tempo 90 (3.5 2.4 4.3 3.2 3.1).2
---

# C Major 9th Chord

The C major 9th chord is a major seventh extended with the ninth — expansive and warm. Its tones are **C – E – G – B – D**: C (root), E (major third), G (perfect fifth), B (major seventh), D (ninth).

| Interval | Note | Name |
|---|---|---|
| 1 | C | root |
| 3 | E | major third |
| 5 | G | perfect fifth |
| 7 | B | major seventh |
| 9 | D | ninth |

**Playing it:** the primary shape is `x-3-2-4-3-3` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [C Major 7th Chord](/chords/c-major-7th-chord.md)
