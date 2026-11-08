---
type: Concept
slug: /chords/e-6th-chord
title: E 6th Chord
description: 'The 6th chord built on E: E, G#, B, C#.'
tags:
- chord
- 6th
- e
parent: /chords
prerequisites:
- /triads/e-major-triad
theory:
  category: chord
  quality: 6th
  root:
    pc: 4
    name: E
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
  - 4
  - 8
  - 11
  - 1
  spelling:
  - E
  - G#
  - B
  - C#
  enharmonic_policy: prefer_sharps
  key_context: E
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
  - 4
  - 8
  - 11
  - 1
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - '6'
  preferred_positions:
  - id: e-6th-voicing-1
    start_fret: 1
    end_fret: 2
    notes:
    - string: 6
      fret: 0
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 0
    - string: 5
      fret: 2
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 2
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 2
    - string: 3
      fret: 1
      pc: 8
      note: G#
      interval: '3'
      role: third
      finger: 1
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: '6'
      role: sixth
      finger: 2
    - string: 1
      fret: 0
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 0
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: e-6th-voicing-2
    start_fret: 2
    end_fret: 4
    notes:
    - string: 6
      fret: 0
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 0
    - string: 5
      fret: 2
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 1
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 2
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: '6'
      role: sixth
      finger: 1
    - string: 1
      fret: 4
      pc: 8
      note: G#
      interval: '3'
      role: third
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: e-6th-voicing-3
    start_fret: 5
    end_fret: 7
    notes:
    - string: 5
      fret: 7
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 6
      pc: 8
      note: G#
      interval: '3'
      role: third
      finger: 2
    - string: 3
      fret: 6
      pc: 1
      note: C#
      interval: '6'
      role: sixth
      finger: 2
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 1
    - string: 1
      fret: 7
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 3
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (0.6 2.5 2.4 1.3 2.2 0.1).2
  midi_root: 64
musical_forms:
- form: pitch_set
  id: e-6th-chord-pitch-set
  pitch_classes:
  - 4
  - 8
  - 11
  - 1
  intervals:
  - '1'
  - '3'
  - '5'
  - '6'
  spelling:
  - E
  - G#
  - B
  - C#
- form: voicing
  id: e-6th-voicing-1
  position_ref: e-6th-voicing-1
  start_fret: 1
  tab: 0-2-2-1-2-0
- form: voicing
  id: e-6th-voicing-2
  position_ref: e-6th-voicing-2
  start_fret: 2
  tab: 0-2-2-4-2-4
- form: voicing
  id: e-6th-voicing-3
  position_ref: e-6th-voicing-3
  start_fret: 5
  tab: x-7-6-6-5-7
- form: tab_excerpt
  id: e-6th-chord-tab
  alphatex: \tempo 90 (0.6 2.5 2.4 1.3 2.2 0.1).2
---

# E 6th Chord

The E 6th chord is a major triad plus a major sixth — a sweet, vintage tonic substitute. Its tones are **E – G# – B – C#**: E (root), G# (major third), B (perfect fifth), C# (major sixth).

| Interval | Note | Name |
|---|---|---|
| 1 | E | root |
| 3 | G# | major third |
| 5 | B | perfect fifth |
| 6 | C# | major sixth |

**Playing it:** the primary shape is `0-2-2-1-2-0` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [E Major Triad](/triads/e-major-triad.md)
