---
type: Concept
slug: /chords/d-minor-9th-chord
title: D Minor 9th Chord
description: 'The minor 9th chord built on D: D, F, A, C, E.'
tags:
- chord
- minor-9th
- d
parent: /chords
prerequisites:
- /chords/d-minor-7th-chord
theory:
  category: chord
  quality: minor-9th
  root:
    pc: 2
    name: D
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
  - 2
  - 5
  - 9
  - 0
  - 4
  spelling:
  - D
  - F
  - A
  - C
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
  - 5
  - 9
  - 0
  - 4
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  - b7
  - '9'
  preferred_positions:
  - id: d-minor-9th-voicing-1
    start_fret: 3
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
      fret: 3
      pc: 5
      note: F
      interval: b3
      role: third
      finger: 1
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
  - id: d-minor-9th-voicing-2
    start_fret: 1
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
      fret: 2
      pc: 4
      note: E
      interval: '9'
      role: ninth
      finger: 2
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: '5'
      role: fifth
      finger: 2
    - string: 2
      fret: 1
      pc: 0
      note: C
      interval: b7
      role: seventh
      finger: 1
    - string: 1
      fret: 1
      pc: 5
      note: F
      interval: b3
      role: third
      finger: 1
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: d-minor-9th-voicing-3
    start_fret: 2
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
      fret: 2
      pc: 4
      note: E
      interval: '9'
      role: ninth
      finger: 1
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: b7
      role: seventh
      finger: 2
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: b3
      role: third
      finger: 3
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
playback:
  alphatex: \tempo 90 (5.5 3.4 5.3 5.2 5.1).2
  midi_root: 62
musical_forms:
- form: pitch_set
  id: d-minor-9th-chord-pitch-set
  pitch_classes:
  - 2
  - 5
  - 9
  - 0
  - 4
  intervals:
  - '1'
  - b3
  - '5'
  - b7
  - '9'
  spelling:
  - D
  - F
  - A
  - C
  - E
- form: voicing
  id: d-minor-9th-voicing-1
  position_ref: d-minor-9th-voicing-1
  start_fret: 3
  tab: x-5-3-5-5-5
- form: voicing
  id: d-minor-9th-voicing-2
  position_ref: d-minor-9th-voicing-2
  start_fret: 1
  tab: x-5-2-2-1-1
- form: voicing
  id: d-minor-9th-voicing-3
  position_ref: d-minor-9th-voicing-3
  start_fret: 2
  tab: x-5-2-5-6-5
- form: tab_excerpt
  id: d-minor-9th-chord-tab
  alphatex: \tempo 90 (5.5 3.4 5.3 5.2 5.1).2
---

# D Minor 9th Chord

The D minor 9th chord is a minor seventh extended with the ninth — smooth and modern. Its tones are **D – F – A – C – E**: D (root), F (minor third), A (perfect fifth), C (minor seventh), E (ninth).

| Interval | Note | Name |
|---|---|---|
| 1 | D | root |
| b3 | F | minor third |
| 5 | A | perfect fifth |
| b7 | C | minor seventh |
| 9 | E | ninth |

**Playing it:** the primary shape is `x-5-3-5-5-5` (strings low→high, `x` = muted), around fret 3. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [D Minor 7th Chord](/chords/d-minor-7th-chord.md)
