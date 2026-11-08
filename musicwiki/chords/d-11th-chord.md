---
type: Concept
slug: /chords/d-11th-chord
title: D 11th Chord
description: 'The 11th chord built on D: D, F#, A, C, E, G.'
tags:
- chord
- 11th
- d
parent: /chords
prerequisites:
- /chords/d-dominant-7th-chord
theory:
  category: chord
  quality: 11th
  root:
    pc: 2
    name: D
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
  - 2
  - 6
  - 9
  - 0
  - 4
  - 7
  spelling:
  - D
  - F#
  - A
  - C
  - E
  - G
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
  - 0
  - 4
  - 7
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '11'
  preferred_positions:
  - id: d-11th-voicing-1
    start_fret: 3
    end_fret: 5
    notes:
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 4
      pc: 6
      note: F#
      interval: '3'
      role: third
      finger: 2
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: b7
      role: seventh
      finger: 3
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '9'
      role: ninth
      finger: 3
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: '11'
      role: eleventh
      finger: 1
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: d-11th-voicing-2
    start_fret: 7
    end_fret: 10
    notes:
    - string: 6
      fret: 10
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 4
    - string: 5
      fret: 9
      pc: 6
      note: F#
      interval: '3'
      role: third
      finger: 3
    - string: 4
      fret: 7
      pc: 9
      note: A
      interval: '5'
      role: fifth
      finger: 1
    - string: 3
      fret: 9
      pc: 4
      note: E
      interval: '9'
      role: ninth
      finger: 3
    - string: 2
      fret: 8
      pc: 7
      note: G
      interval: '11'
      role: eleventh
      finger: 2
    - string: 1
      fret: 8
      pc: 0
      note: C
      interval: b7
      role: seventh
      finger: 2
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: d-11th-voicing-3
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
      fret: 5
      pc: 7
      note: G
      interval: '11'
      role: eleventh
      finger: 2
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: b7
      role: seventh
      finger: 2
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '9'
      role: ninth
      finger: 2
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
playback:
  alphatex: \tempo 90 (5.5 4.4 5.3 5.2 3.1).2
  midi_root: 62
musical_forms:
- form: pitch_set
  id: d-11th-chord-pitch-set
  pitch_classes:
  - 2
  - 6
  - 9
  - 0
  - 4
  - 7
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '11'
  spelling:
  - D
  - F#
  - A
  - C
  - E
  - G
- form: voicing
  id: d-11th-voicing-1
  position_ref: d-11th-voicing-1
  start_fret: 3
  tab: x-5-4-5-5-3
- form: voicing
  id: d-11th-voicing-2
  position_ref: d-11th-voicing-2
  start_fret: 7
  tab: 10-9-7-9-8-8
- form: voicing
  id: d-11th-voicing-3
  position_ref: d-11th-voicing-3
  start_fret: 2
  tab: x-5-5-5-5-2
- form: tab_excerpt
  id: d-11th-chord-tab
  alphatex: \tempo 90 (5.5 4.4 5.3 5.2 3.1).2
---

# D 11th Chord

The D 11th chord is a dominant chord extended through the eleventh. Its tones are **D – F# – A – C – E – G**: D (root), F# (major third), A (perfect fifth), C (minor seventh), E (ninth), G (eleventh).

| Interval | Note | Name |
|---|---|---|
| 1 | D | root |
| 3 | F# | major third |
| 5 | A | perfect fifth |
| b7 | C | minor seventh |
| 9 | E | ninth |
| 11 | G | eleventh |

**Playing it:** the primary shape is `x-5-4-5-5-3` (strings low→high, `x` = muted), around fret 3. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [D Dominant 7th Chord](/chords/d-dominant-7th-chord.md)
