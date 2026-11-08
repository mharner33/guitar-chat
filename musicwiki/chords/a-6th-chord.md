---
type: Concept
slug: /chords/a-6th-chord
title: A 6th Chord
description: 'The 6th chord built on A: A, C#, E, F#.'
tags:
- chord
- 6th
- a
parent: /chords
prerequisites:
- /triads/a-major-triad
theory:
  category: chord
  quality: 6th
  root:
    pc: 9
    name: A
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
  - 9
  - 1
  - 4
  - 6
  spelling:
  - A
  - C#
  - E
  - F#
  enharmonic_policy: prefer_sharps
  key_context: A
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
  - 9
  - 1
  - 4
  - 6
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - '6'
  preferred_positions:
  - id: a-6th-voicing-1
    start_fret: 2
    end_fret: 2
    notes:
    - string: 5
      fret: 0
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 0
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: '5'
      role: fifth
      finger: 1
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 1
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: '3'
      role: third
      finger: 1
    - string: 1
      fret: 2
      pc: 6
      note: F#
      interval: '6'
      role: sixth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: a-6th-voicing-2
    start_fret: 4
    end_fret: 6
    notes:
    - string: 6
      fret: 5
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 4
      pc: 1
      note: C#
      interval: '3'
      role: third
      finger: 1
    - string: 4
      fret: 4
      pc: 6
      note: F#
      interval: '6'
      role: sixth
      finger: 1
    - string: 3
      fret: 6
      pc: 1
      note: C#
      interval: '3'
      role: third
      finger: 3
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '5'
      role: fifth
      finger: 2
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: a-6th-voicing-3
    start_fret: 5
    end_fret: 7
    notes:
    - string: 6
      fret: 5
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 7
      pc: 4
      note: E
      interval: '5'
      role: fifth
      finger: 3
    - string: 4
      fret: 7
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 3
    - string: 3
      fret: 6
      pc: 1
      note: C#
      interval: '3'
      role: third
      finger: 2
    - string: 2
      fret: 7
      pc: 6
      note: F#
      interval: '6'
      role: sixth
      finger: 3
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (0.5 2.4 2.3 2.2 2.1).2
  midi_root: 69
musical_forms:
- form: pitch_set
  id: a-6th-chord-pitch-set
  pitch_classes:
  - 9
  - 1
  - 4
  - 6
  intervals:
  - '1'
  - '3'
  - '5'
  - '6'
  spelling:
  - A
  - C#
  - E
  - F#
- form: voicing
  id: a-6th-voicing-1
  position_ref: a-6th-voicing-1
  start_fret: 2
  tab: x-0-2-2-2-2
- form: voicing
  id: a-6th-voicing-2
  position_ref: a-6th-voicing-2
  start_fret: 4
  tab: 5-4-4-6-5-5
- form: voicing
  id: a-6th-voicing-3
  position_ref: a-6th-voicing-3
  start_fret: 5
  tab: 5-7-7-6-7-5
- form: tab_excerpt
  id: a-6th-chord-tab
  alphatex: \tempo 90 (0.5 2.4 2.3 2.2 2.1).2
---

# A 6th Chord

The A 6th chord is a major triad plus a major sixth — a sweet, vintage tonic substitute. Its tones are **A – C# – E – F#**: A (root), C# (major third), E (perfect fifth), F# (major sixth).

| Interval | Note | Name |
|---|---|---|
| 1 | A | root |
| 3 | C# | major third |
| 5 | E | perfect fifth |
| 6 | F# | major sixth |

**Playing it:** the primary shape is `x-0-2-2-2-2` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [A Major Triad](/triads/a-major-triad.md)
