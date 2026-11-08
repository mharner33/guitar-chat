---
type: Concept
slug: /chords/f-11th-chord
title: F 11th Chord
description: 'The 11th chord built on F: F, A, C, Eb, G, Bb.'
tags:
- chord
- 11th
- f
parent: /chords
prerequisites:
- /chords/f-dominant-7th-chord
theory:
  category: chord
  quality: 11th
  root:
    pc: 5
    name: F
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
  - 5
  - 9
  - 0
  - 3
  - 7
  - 10
  spelling:
  - F
  - A
  - C
  - Eb
  - G
  - Bb
  enharmonic_policy: prefer_flats
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
  - 3
  - 7
  - 10
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '11'
  preferred_positions:
  - id: f-11th-voicing-1
    start_fret: 1
    end_fret: 3
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
      fret: 1
      pc: 3
      note: Eb
      interval: b7
      role: seventh
      finger: 1
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '11'
      role: eleventh
      finger: 2
    - string: 2
      fret: 1
      pc: 0
      note: C
      interval: '5'
      role: fifth
      finger: 1
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: '9'
      role: ninth
      finger: 2
    confidence: 0.85
    tags:
    - computed
    - open
    - barre
  - id: f-11th-voicing-2
    start_fret: 6
    end_fret: 8
    notes:
    - string: 5
      fret: 8
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 7
      pc: 9
      note: A
      interval: '3'
      role: third
      finger: 2
    - string: 3
      fret: 8
      pc: 3
      note: Eb
      interval: b7
      role: seventh
      finger: 3
    - string: 2
      fret: 8
      pc: 7
      note: G
      interval: '9'
      role: ninth
      finger: 3
    - string: 1
      fret: 6
      pc: 10
      note: Bb
      interval: '11'
      role: eleventh
      finger: 1
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: f-11th-voicing-3
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
      fret: 8
      pc: 10
      note: Bb
      interval: '11'
      role: eleventh
      finger: 2
    - string: 3
      fret: 8
      pc: 3
      note: Eb
      interval: b7
      role: seventh
      finger: 2
    - string: 2
      fret: 8
      pc: 7
      note: G
      interval: '9'
      role: ninth
      finger: 2
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
playback:
  alphatex: \tempo 90 (1.6 0.5 1.4 3.3 1.2 3.1).2
  midi_root: 65
musical_forms:
- form: pitch_set
  id: f-11th-chord-pitch-set
  pitch_classes:
  - 5
  - 9
  - 0
  - 3
  - 7
  - 10
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '11'
  spelling:
  - F
  - A
  - C
  - Eb
  - G
  - Bb
- form: voicing
  id: f-11th-voicing-1
  position_ref: f-11th-voicing-1
  start_fret: 1
  tab: 1-0-1-3-1-3
- form: voicing
  id: f-11th-voicing-2
  position_ref: f-11th-voicing-2
  start_fret: 6
  tab: x-8-7-8-8-6
- form: voicing
  id: f-11th-voicing-3
  position_ref: f-11th-voicing-3
  start_fret: 5
  tab: x-8-8-8-8-5
- form: tab_excerpt
  id: f-11th-chord-tab
  alphatex: \tempo 90 (1.6 0.5 1.4 3.3 1.2 3.1).2
---

# F 11th Chord

The F 11th chord is a dominant chord extended through the eleventh. Its tones are **F – A – C – Eb – G – Bb**: F (root), A (major third), C (perfect fifth), Eb (minor seventh), G (ninth), Bb (eleventh).

| Interval | Note | Name |
|---|---|---|
| 1 | F | root |
| 3 | A | major third |
| 5 | C | perfect fifth |
| b7 | Eb | minor seventh |
| 9 | G | ninth |
| 11 | Bb | eleventh |

**Playing it:** the primary shape is `1-0-1-3-1-3` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [F Dominant 7th Chord](/chords/f-dominant-7th-chord.md)
