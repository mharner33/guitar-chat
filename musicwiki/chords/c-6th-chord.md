---
type: Concept
slug: /chords/c-6th-chord
title: C 6th Chord
description: 'The 6th chord built on C: C, E, G, A.'
tags:
- chord
- 6th
- c
parent: /chords
prerequisites:
- /triads/c-major-triad
theory:
  category: chord
  quality: 6th
  root:
    pc: 0
    name: C
  intervals:
  - '1'
  - '3'
  - '5'
  - '6'
  semitones:
  - 0
  - 4
  - 7
  - 9
  pitch_classes:
  - 0
  - 4
  - 7
  - 9
  spelling:
  - C
  - E
  - G
  - A
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
  - 9
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - '6'
  preferred_positions:
  - id: c-6th-voicing-1
    start_fret: 1
    end_fret: 3
    notes:
    - string: 5
      fret: 3
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: '3'
      role: third
      finger: 2
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: '6'
      role: sixth
      finger: 2
    - string: 2
      fret: 1
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 1
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: '5'
      role: fifth
      finger: 3
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: c-6th-voicing-2
    start_fret: 3
    end_fret: 5
    notes:
    - string: 5
      fret: 3
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: '5'
      role: fifth
      finger: 2
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 2
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '3'
      role: third
      finger: 2
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '6'
      role: sixth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: c-6th-voicing-3
    start_fret: 2
    end_fret: 5
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
      fret: 0
      pc: 7
      note: G
      interval: '5'
      role: fifth
      finger: 0
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '3'
      role: third
      finger: 3
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '6'
      role: sixth
      finger: 3
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
playback:
  alphatex: \tempo 90 (3.5 2.4 2.3 1.2 3.1).2
  midi_root: 60
musical_forms:
- form: pitch_set
  id: c-6th-chord-pitch-set
  pitch_classes:
  - 0
  - 4
  - 7
  - 9
  intervals:
  - '1'
  - '3'
  - '5'
  - '6'
  spelling:
  - C
  - E
  - G
  - A
- form: voicing
  id: c-6th-voicing-1
  position_ref: c-6th-voicing-1
  start_fret: 1
  tab: x-3-2-2-1-3
- form: voicing
  id: c-6th-voicing-2
  position_ref: c-6th-voicing-2
  start_fret: 3
  tab: x-3-5-5-5-5
- form: voicing
  id: c-6th-voicing-3
  position_ref: c-6th-voicing-3
  start_fret: 2
  tab: x-3-2-0-5-5
- form: tab_excerpt
  id: c-6th-chord-tab
  alphatex: \tempo 90 (3.5 2.4 2.3 1.2 3.1).2
---

# C 6th Chord

The C 6th chord is a major triad plus a major sixth — a sweet, vintage tonic substitute. Its tones are **C – E – G – A**: C (root), E (major third), G (perfect fifth), A (major sixth).

| Interval | Note | Name |
|---|---|---|
| 1 | C | root |
| 3 | E | major third |
| 5 | G | perfect fifth |
| 6 | A | major sixth |

**Playing it:** the primary shape is `x-3-2-2-1-3` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [C Major Triad](/triads/c-major-triad.md)
