---
type: Concept
slug: /chords/b-flat-9th-chord
title: Bb 9th Chord
description: 'The 9th chord built on Bb: Bb, D, F, Ab, C.'
tags:
- chord
- 9th
- b
parent: /chords
prerequisites:
- /chords/b-flat-dominant-7th-chord
theory:
  category: chord
  quality: 9th
  root:
    pc: 10
    name: Bb
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  semitones:
  - 0
  - 4
  - 7
  - 10
  - 14
  pitch_classes:
  - 10
  - 2
  - 5
  - 8
  - 0
  spelling:
  - Bb
  - D
  - F
  - Ab
  - C
  enharmonic_policy: prefer_flats
  key_context: Bb
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
  - 10
  - 2
  - 5
  - 8
  - 0
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  preferred_positions:
  - id: b-flat-9th-voicing-1
    start_fret: 1
    end_fret: 1
    notes:
    - string: 5
      fret: 1
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 0
      pc: 2
      note: D
      interval: '3'
      role: third
      finger: 0
    - string: 3
      fret: 1
      pc: 8
      note: Ab
      interval: b7
      role: seventh
      finger: 1
    - string: 2
      fret: 1
      pc: 0
      note: C
      interval: '9'
      role: ninth
      finger: 1
    - string: 1
      fret: 1
      pc: 5
      note: F
      interval: '5'
      role: fifth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - open
    - barre
  - id: b-flat-9th-voicing-2
    start_fret: 5
    end_fret: 6
    notes:
    - string: 6
      fret: 6
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: '3'
      role: third
      finger: 1
    - string: 4
      fret: 6
      pc: 8
      note: Ab
      interval: b7
      role: seventh
      finger: 2
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: '9'
      role: ninth
      finger: 1
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '5'
      role: fifth
      finger: 2
    - string: 1
      fret: 6
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 2
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: b-flat-9th-voicing-3
    start_fret: 4
    end_fret: 6
    notes:
    - string: 6
      fret: 6
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: '3'
      role: third
      finger: 2
    - string: 4
      fret: 6
      pc: 8
      note: Ab
      interval: b7
      role: seventh
      finger: 3
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: '9'
      role: ninth
      finger: 2
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '5'
      role: fifth
      finger: 3
    - string: 1
      fret: 4
      pc: 8
      note: Ab
      interval: b7
      role: seventh
      finger: 1
    confidence: 0.85
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (1.5 0.4 1.3 1.2 1.1).2
  midi_root: 70
musical_forms:
- form: pitch_set
  id: b-flat-9th-chord-pitch-set
  pitch_classes:
  - 10
  - 2
  - 5
  - 8
  - 0
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  spelling:
  - Bb
  - D
  - F
  - Ab
  - C
- form: voicing
  id: b-flat-9th-voicing-1
  position_ref: b-flat-9th-voicing-1
  start_fret: 1
  tab: x-1-0-1-1-1
- form: voicing
  id: b-flat-9th-voicing-2
  position_ref: b-flat-9th-voicing-2
  start_fret: 5
  tab: 6-5-6-5-6-6
- form: voicing
  id: b-flat-9th-voicing-3
  position_ref: b-flat-9th-voicing-3
  start_fret: 4
  tab: 6-5-6-5-6-4
- form: tab_excerpt
  id: b-flat-9th-chord-tab
  alphatex: \tempo 90 (1.5 0.4 1.3 1.2 1.1).2
---

# Bb 9th Chord

The Bb 9th chord is a dominant seventh extended with the ninth — a fuller dominant color. Its tones are **Bb – D – F – Ab – C**: Bb (root), D (major third), F (perfect fifth), Ab (minor seventh), C (ninth).

| Interval | Note | Name |
|---|---|---|
| 1 | Bb | root |
| 3 | D | major third |
| 5 | F | perfect fifth |
| b7 | Ab | minor seventh |
| 9 | C | ninth |

**Playing it:** the primary shape is `x-1-0-1-1-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Bb Dominant 7th Chord](/chords/b-flat-dominant-7th-chord.md)
