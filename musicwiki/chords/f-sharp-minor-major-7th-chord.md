---
type: Concept
slug: /chords/f-sharp-minor-major-7th-chord
title: F# Minor Major 7th Chord
description: 'The minor major 7th chord built on F#: F#, A, C#, E#.'
tags:
- chord
- minor-major-7th
- f
parent: /chords
prerequisites:
- /triads/f-sharp-minor-triad
theory:
  category: chord
  quality: minor-major-7th
  root:
    pc: 6
    name: F#
  intervals:
  - '1'
  - b3
  - '5'
  - '7'
  semitones:
  - 0
  - 3
  - 7
  - 11
  pitch_classes:
  - 6
  - 9
  - 1
  - 5
  spelling:
  - F#
  - A
  - C#
  - E#
  enharmonic_policy: prefer_sharps
  key_context: F#
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
  - 6
  - 9
  - 1
  - 5
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  - '7'
  preferred_positions:
  - id: f-sharp-minor-major-7th-voicing-1
    start_fret: 2
    end_fret: 3
    notes:
    - string: 6
      fret: 2
      pc: 6
      note: F#
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 0
      pc: 9
      note: A
      interval: b3
      role: third
      finger: 0
    - string: 4
      fret: 3
      pc: 5
      note: E#
      interval: '7'
      role: seventh
      finger: 2
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: b3
      role: third
      finger: 1
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: '5'
      role: fifth
      finger: 1
    - string: 1
      fret: 2
      pc: 6
      note: F#
      interval: '1'
      role: root
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: f-sharp-minor-major-7th-voicing-2
    start_fret: 1
    end_fret: 3
    notes:
    - string: 6
      fret: 2
      pc: 6
      note: F#
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 0
      pc: 9
      note: A
      interval: b3
      role: third
      finger: 0
    - string: 4
      fret: 3
      pc: 5
      note: E#
      interval: '7'
      role: seventh
      finger: 3
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: b3
      role: third
      finger: 2
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: '5'
      role: fifth
      finger: 2
    - string: 1
      fret: 1
      pc: 5
      note: E#
      interval: '7'
      role: seventh
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: f-sharp-minor-major-7th-voicing-3
    start_fret: 4
    end_fret: 6
    notes:
    - string: 4
      fret: 4
      pc: 6
      note: F#
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 6
      pc: 1
      note: C#
      interval: '5'
      role: fifth
      finger: 3
    - string: 2
      fret: 6
      pc: 5
      note: E#
      interval: '7'
      role: seventh
      finger: 3
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: b3
      role: third
      finger: 2
    mute_strings:
    - 5
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (2.6 0.5 3.4 2.3 2.2 2.1).2
  midi_root: 66
musical_forms:
- form: pitch_set
  id: f-sharp-minor-major-7th-chord-pitch-set
  pitch_classes:
  - 6
  - 9
  - 1
  - 5
  intervals:
  - '1'
  - b3
  - '5'
  - '7'
  spelling:
  - F#
  - A
  - C#
  - E#
- form: voicing
  id: f-sharp-minor-major-7th-voicing-1
  position_ref: f-sharp-minor-major-7th-voicing-1
  start_fret: 2
  tab: 2-0-3-2-2-2
- form: voicing
  id: f-sharp-minor-major-7th-voicing-2
  position_ref: f-sharp-minor-major-7th-voicing-2
  start_fret: 1
  tab: 2-0-3-2-2-1
- form: voicing
  id: f-sharp-minor-major-7th-voicing-3
  position_ref: f-sharp-minor-major-7th-voicing-3
  start_fret: 4
  tab: x-x-4-6-6-5
- form: tab_excerpt
  id: f-sharp-minor-major-7th-chord-tab
  alphatex: \tempo 90 (2.6 0.5 3.4 2.3 2.2 2.1).2
---

# F# Minor Major 7th Chord

The F# minor major 7th chord is a minor triad plus a major seventh — a haunting, unstable color. Its tones are **F# – A – C# – E#**: F# (root), A (minor third), C# (perfect fifth), E# (major seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | F# | root |
| b3 | A | minor third |
| 5 | C# | perfect fifth |
| 7 | E# | major seventh |

**Playing it:** the primary shape is `2-0-3-2-2-2` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [F# Minor Triad](/triads/f-sharp-minor-triad.md)
