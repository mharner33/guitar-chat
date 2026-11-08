---
type: Concept
slug: /chords/d-minor-7th-chord
title: D Minor 7th Chord
description: 'The minor 7th chord built on D: D, F, A, C.'
tags:
- chord
- minor-7th
- d
parent: /chords
prerequisites:
- /triads/d-minor-triad
used_by:
- /chords/d-minor-9th-chord
theory:
  category: chord
  quality: minor-7th
  root:
    pc: 2
    name: D
  intervals:
  - '1'
  - b3
  - '5'
  - b7
  semitones:
  - 0
  - 3
  - 7
  - 10
  pitch_classes:
  - 2
  - 5
  - 9
  - 0
  spelling:
  - D
  - F
  - A
  - C
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
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  - b7
  preferred_positions:
  - id: d-minor-7th-voicing-1
    start_fret: 1
    end_fret: 2
    notes:
    - string: 4
      fret: 0
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 0
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
    - 5
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: d-minor-7th-voicing-2
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
      fret: 3
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 1
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '5'
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: d-minor-7th-voicing-3
    start_fret: 5
    end_fret: 7
    notes:
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 7
      pc: 9
      note: A
      interval: '5'
      role: fifth
      finger: 3
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: b7
      role: seventh
      finger: 1
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: b3
      role: third
      finger: 2
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '5'
      role: fifth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (0.4 2.3 1.2 1.1).2
  midi_root: 62
musical_forms:
- form: pitch_set
  id: d-minor-7th-chord-pitch-set
  pitch_classes:
  - 2
  - 5
  - 9
  - 0
  intervals:
  - '1'
  - b3
  - '5'
  - b7
  spelling:
  - D
  - F
  - A
  - C
- form: voicing
  id: d-minor-7th-voicing-1
  position_ref: d-minor-7th-voicing-1
  start_fret: 1
  tab: x-x-0-2-1-1
- form: voicing
  id: d-minor-7th-voicing-2
  position_ref: d-minor-7th-voicing-2
  start_fret: 3
  tab: x-5-3-5-3-5
- form: voicing
  id: d-minor-7th-voicing-3
  position_ref: d-minor-7th-voicing-3
  start_fret: 5
  tab: x-5-7-5-6-5
- form: tab_excerpt
  id: d-minor-7th-chord-tab
  alphatex: \tempo 90 (0.4 2.3 1.2 1.1).2
---

# D Minor 7th Chord

The D minor 7th chord is a minor triad plus a minor seventh — smooth and mellow, the workhorse of jazz and soul. Its tones are **D – F – A – C**: D (root), F (minor third), A (perfect fifth), C (minor seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | D | root |
| b3 | F | minor third |
| 5 | A | perfect fifth |
| b7 | C | minor seventh |

**Playing it:** the primary shape is `x-x-0-2-1-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [D Minor Triad](/triads/d-minor-triad.md)
- **Used by:** [D Minor 9th Chord](/chords/d-minor-9th-chord.md)
