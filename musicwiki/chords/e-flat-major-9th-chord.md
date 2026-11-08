---
type: Concept
slug: /chords/e-flat-major-9th-chord
title: Eb Major 9th Chord
description: 'The major 9th chord built on Eb: Eb, G, Bb, D, F.'
tags:
- chord
- major-9th
- e
parent: /chords
prerequisites:
- /chords/e-flat-major-7th-chord
theory:
  category: chord
  quality: major-9th
  root:
    pc: 3
    name: Eb
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
  - 3
  - 7
  - 10
  - 2
  - 5
  spelling:
  - Eb
  - G
  - Bb
  - D
  - F
  enharmonic_policy: prefer_flats
  key_context: Eb
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
  - 3
  - 7
  - 10
  - 2
  - 5
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - '7'
  - '9'
  preferred_positions:
  - id: e-flat-major-9th-voicing-1
    start_fret: 3
    end_fret: 6
    notes:
    - string: 5
      fret: 6
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '9'
      role: ninth
      finger: 1
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '5'
      role: fifth
      finger: 1
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '7'
      role: seventh
      finger: 1
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: '3'
      role: third
      finger: 1
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: e-flat-major-9th-voicing-2
    start_fret: 5
    end_fret: 7
    notes:
    - string: 5
      fret: 6
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: '3'
      role: third
      finger: 1
    - string: 3
      fret: 7
      pc: 2
      note: D
      interval: '7'
      role: seventh
      finger: 3
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '9'
      role: ninth
      finger: 2
    - string: 1
      fret: 6
      pc: 10
      note: Bb
      interval: '5'
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: e-flat-major-9th-voicing-3
    start_fret: 1
    end_fret: 3
    notes:
    - string: 4
      fret: 1
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 0
      pc: 7
      note: G
      interval: '3'
      role: third
      finger: 0
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '7'
      role: seventh
      finger: 2
    - string: 1
      fret: 1
      pc: 5
      note: F
      interval: '9'
      role: ninth
      finger: 1
    mute_strings:
    - 5
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
playback:
  alphatex: \tempo 90 (6.5 3.4 3.3 3.2 3.1).2
  midi_root: 63
musical_forms:
- form: pitch_set
  id: e-flat-major-9th-chord-pitch-set
  pitch_classes:
  - 3
  - 7
  - 10
  - 2
  - 5
  intervals:
  - '1'
  - '3'
  - '5'
  - '7'
  - '9'
  spelling:
  - Eb
  - G
  - Bb
  - D
  - F
- form: voicing
  id: e-flat-major-9th-voicing-1
  position_ref: e-flat-major-9th-voicing-1
  start_fret: 3
  tab: x-6-3-3-3-3
- form: voicing
  id: e-flat-major-9th-voicing-2
  position_ref: e-flat-major-9th-voicing-2
  start_fret: 5
  tab: x-6-5-7-6-6
- form: voicing
  id: e-flat-major-9th-voicing-3
  position_ref: e-flat-major-9th-voicing-3
  start_fret: 1
  tab: x-x-1-0-3-1
- form: tab_excerpt
  id: e-flat-major-9th-chord-tab
  alphatex: \tempo 90 (6.5 3.4 3.3 3.2 3.1).2
---

# Eb Major 9th Chord

The Eb major 9th chord is a major seventh extended with the ninth — expansive and warm. Its tones are **Eb – G – Bb – D – F**: Eb (root), G (major third), Bb (perfect fifth), D (major seventh), F (ninth).

| Interval | Note | Name |
|---|---|---|
| 1 | Eb | root |
| 3 | G | major third |
| 5 | Bb | perfect fifth |
| 7 | D | major seventh |
| 9 | F | ninth |

**Playing it:** the primary shape is `x-6-3-3-3-3` (strings low→high, `x` = muted), around fret 3. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Eb Major 7th Chord](/chords/e-flat-major-7th-chord.md)
