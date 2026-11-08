---
type: Concept
slug: /chords/d-major-9th-chord
title: D Major 9th Chord
description: 'The major 9th chord built on D: D, F#, A, C#, E.'
tags:
- chord
- major-9th
- d
parent: /chords
prerequisites:
- /chords/d-major-7th-chord
theory:
  category: chord
  quality: major-9th
  root:
    pc: 2
    name: D
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
  - 2
  - 6
  - 9
  - 1
  - 4
  spelling:
  - D
  - F#
  - A
  - C#
  - E
  enharmonic_policy: prefer_sharps
  key_context: D
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
  - 2
  - 6
  - 9
  - 1
  - 4
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - '7'
  - '9'
  preferred_positions:
  - id: d-major-9th-voicing-1
    start_fret: 2
    end_fret: 5
    notes:
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: '9'
      role: ninth
      finger: 1
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: '5'
      role: fifth
      finger: 1
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: '7'
      role: seventh
      finger: 1
    - string: 1
      fret: 2
      pc: 6
      note: F#
      interval: '3'
      role: third
      finger: 1
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: d-major-9th-voicing-2
    start_fret: 4
    end_fret: 6
    notes:
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 4
      pc: 6
      note: F#
      interval: '3'
      role: third
      finger: 1
    - string: 3
      fret: 6
      pc: 1
      note: C#
      interval: '7'
      role: seventh
      finger: 3
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '9'
      role: ninth
      finger: 2
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '5'
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: d-major-9th-voicing-3
    start_fret: 7
    end_fret: 10
    notes:
    - string: 6
      fret: 10
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 7
      pc: 4
      note: E
      interval: '9'
      role: ninth
      finger: 1
    - string: 4
      fret: 7
      pc: 9
      note: A
      interval: '5'
      role: fifth
      finger: 1
    - string: 3
      fret: 7
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 1
    - string: 2
      fret: 7
      pc: 6
      note: F#
      interval: '3'
      role: third
      finger: 1
    - string: 1
      fret: 9
      pc: 1
      note: C#
      interval: '7'
      role: seventh
      finger: 2
    confidence: 0.85
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (5.5 2.4 2.3 2.2 2.1).2
  midi_root: 62
musical_forms:
- form: pitch_set
  id: d-major-9th-chord-pitch-set
  pitch_classes:
  - 2
  - 6
  - 9
  - 1
  - 4
  intervals:
  - '1'
  - '3'
  - '5'
  - '7'
  - '9'
  spelling:
  - D
  - F#
  - A
  - C#
  - E
- form: voicing
  id: d-major-9th-voicing-1
  position_ref: d-major-9th-voicing-1
  start_fret: 2
  tab: x-5-2-2-2-2
- form: voicing
  id: d-major-9th-voicing-2
  position_ref: d-major-9th-voicing-2
  start_fret: 4
  tab: x-5-4-6-5-5
- form: voicing
  id: d-major-9th-voicing-3
  position_ref: d-major-9th-voicing-3
  start_fret: 7
  tab: 10-7-7-7-7-9
- form: tab_excerpt
  id: d-major-9th-chord-tab
  alphatex: \tempo 90 (5.5 2.4 2.3 2.2 2.1).2
---

# D Major 9th Chord

The D major 9th chord is a major seventh extended with the ninth — expansive and warm. Its tones are **D – F# – A – C# – E**: D (root), F# (major third), A (perfect fifth), C# (major seventh), E (ninth).

| Interval | Note | Name |
|---|---|---|
| 1 | D | root |
| 3 | F# | major third |
| 5 | A | perfect fifth |
| 7 | C# | major seventh |
| 9 | E | ninth |

**Playing it:** the primary shape is `x-5-2-2-2-2` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [D Major 7th Chord](/chords/d-major-7th-chord.md)
