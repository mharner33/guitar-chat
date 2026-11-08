---
type: Concept
slug: /chords/f-sharp-minor-7th-flat-5-chord
title: F# Minor 7th Flat 5 Chord
description: 'The minor 7th flat 5 chord built on F#: F#, A, C, E.'
tags:
- chord
- minor-7th-flat-5
- f
parent: /chords
prerequisites:
- /triads/f-sharp-diminished-triad
theory:
  category: chord
  quality: minor-7th-flat-5
  root:
    pc: 6
    name: F#
  intervals:
  - '1'
  - b3
  - b5
  - b7
  semitones:
  - 0
  - 3
  - 6
  - 10
  pitch_classes:
  - 6
  - 9
  - 0
  - 4
  spelling:
  - F#
  - A
  - C
  - E
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
  - 0
  - 4
  emphasize_intervals:
  - '1'
  - b3
  - b5
  - b7
  preferred_positions:
  - id: f-sharp-minor-7th-flat-5-voicing-1
    start_fret: 1
    end_fret: 2
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
      fret: 2
      pc: 4
      note: E
      interval: b7
      role: seventh
      finger: 2
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: b3
      role: third
      finger: 2
    - string: 2
      fret: 1
      pc: 0
      note: C
      interval: b5
      role: fifth
      finger: 1
    - string: 1
      fret: 0
      pc: 4
      note: E
      interval: b7
      role: seventh
      finger: 0
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: f-sharp-minor-7th-flat-5-voicing-2
    start_fret: 2
    end_fret: 5
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
      fret: 2
      pc: 4
      note: E
      interval: b7
      role: seventh
      finger: 1
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: b5
      role: fifth
      finger: 2
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: b7
      role: seventh
      finger: 2
    - string: 1
      fret: 0
      pc: 4
      note: E
      interval: b7
      role: seventh
      finger: 0
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: f-sharp-minor-7th-flat-5-voicing-3
    start_fret: 4
    end_fret: 5
    notes:
    - string: 4
      fret: 4
      pc: 6
      note: F#
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: b5
      role: fifth
      finger: 2
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: b7
      role: seventh
      finger: 2
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
  alphatex: \tempo 90 (2.6 0.5 2.4 2.3 1.2 0.1).2
  midi_root: 66
musical_forms:
- form: pitch_set
  id: f-sharp-minor-7th-flat-5-chord-pitch-set
  pitch_classes:
  - 6
  - 9
  - 0
  - 4
  intervals:
  - '1'
  - b3
  - b5
  - b7
  spelling:
  - F#
  - A
  - C
  - E
- form: voicing
  id: f-sharp-minor-7th-flat-5-voicing-1
  position_ref: f-sharp-minor-7th-flat-5-voicing-1
  start_fret: 1
  tab: 2-0-2-2-1-0
- form: voicing
  id: f-sharp-minor-7th-flat-5-voicing-2
  position_ref: f-sharp-minor-7th-flat-5-voicing-2
  start_fret: 2
  tab: 2-0-2-5-5-0
- form: voicing
  id: f-sharp-minor-7th-flat-5-voicing-3
  position_ref: f-sharp-minor-7th-flat-5-voicing-3
  start_fret: 4
  tab: x-x-4-5-5-5
- form: tab_excerpt
  id: f-sharp-minor-7th-flat-5-chord-tab
  alphatex: \tempo 90 (2.6 0.5 2.4 2.3 1.2 0.1).2
---

# F# Minor 7th Flat 5 Chord

The F# minor 7th flat 5 chord is the half-diminished chord: a diminished triad plus a minor seventh, the ii chord of a minor key. Its tones are **F# – A – C – E**: F# (root), A (minor third), C (diminished fifth), E (minor seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | F# | root |
| b3 | A | minor third |
| b5 | C | diminished fifth |
| b7 | E | minor seventh |

**Playing it:** the primary shape is `2-0-2-2-1-0` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [F# Diminished Triad](/triads/f-sharp-diminished-triad.md)
