---
type: Concept
slug: /chords/c-minor-9th-chord
title: C Minor 9th Chord
description: 'The minor 9th chord built on C: C, Eb, G, Bb, D.'
tags:
- chord
- minor-9th
- c
parent: /chords
prerequisites:
- /chords/c-minor-7th-chord
theory:
  category: chord
  quality: minor-9th
  root:
    pc: 0
    name: C
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
  - 0
  - 3
  - 7
  - 10
  - 2
  spelling:
  - C
  - Eb
  - G
  - Bb
  - D
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
  - 3
  - 7
  - 10
  - 2
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  - b7
  - '9'
  preferred_positions:
  - id: c-minor-9th-voicing-1
    start_fret: 1
    end_fret: 3
    notes:
    - string: 5
      fret: 3
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 1
      pc: 3
      note: Eb
      interval: b3
      role: third
      finger: 1
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: b7
      role: seventh
      finger: 2
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
  - id: c-minor-9th-voicing-2
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
      fret: 6
      pc: 3
      note: Eb
      interval: b3
      role: third
      finger: 1
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
  - id: c-minor-9th-voicing-3
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
      fret: 8
      pc: 3
      note: Eb
      interval: b3
      role: third
      finger: 3
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
      finger: 2
    confidence: 0.85
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (3.5 1.4 3.3 3.2 3.1).2
  midi_root: 60
musical_forms:
- form: pitch_set
  id: c-minor-9th-chord-pitch-set
  pitch_classes:
  - 0
  - 3
  - 7
  - 10
  - 2
  intervals:
  - '1'
  - b3
  - '5'
  - b7
  - '9'
  spelling:
  - C
  - Eb
  - G
  - Bb
  - D
- form: voicing
  id: c-minor-9th-voicing-1
  position_ref: c-minor-9th-voicing-1
  start_fret: 1
  tab: x-3-1-3-3-3
- form: voicing
  id: c-minor-9th-voicing-2
  position_ref: c-minor-9th-voicing-2
  start_fret: 6
  tab: 8-6-8-7-8-6
- form: voicing
  id: c-minor-9th-voicing-3
  position_ref: c-minor-9th-voicing-3
  start_fret: 5
  tab: 8-5-5-8-8-6
- form: tab_excerpt
  id: c-minor-9th-chord-tab
  alphatex: \tempo 90 (3.5 1.4 3.3 3.2 3.1).2
---

# C Minor 9th Chord

The C minor 9th chord is a minor seventh extended with the ninth — smooth and modern. Its tones are **C – Eb – G – Bb – D**: C (root), Eb (minor third), G (perfect fifth), Bb (minor seventh), D (ninth).

| Interval | Note | Name |
|---|---|---|
| 1 | C | root |
| b3 | Eb | minor third |
| 5 | G | perfect fifth |
| b7 | Bb | minor seventh |
| 9 | D | ninth |

**Playing it:** the primary shape is `x-3-1-3-3-3` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [C Minor 7th Chord](/chords/c-minor-7th-chord.md)
