---
type: Concept
slug: /chords/d-sharp-minor-7th-flat-5-chord
title: D# Minor 7th Flat 5 Chord
description: 'The minor 7th flat 5 chord built on D#: D#, F#, A, C#.'
tags:
- chord
- minor-7th-flat-5
- d
parent: /chords
prerequisites:
- /triads/d-sharp-diminished-triad
theory:
  category: chord
  quality: minor-7th-flat-5
  root:
    pc: 3
    name: D#
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
  - 3
  - 6
  - 9
  - 1
  spelling:
  - D#
  - F#
  - A
  - C#
  enharmonic_policy: prefer_sharps
  key_context: D#
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
  - 3
  - 6
  - 9
  - 1
  emphasize_intervals:
  - '1'
  - b3
  - b5
  - b7
  preferred_positions:
  - id: d-sharp-minor-7th-flat-5-voicing-1
    start_fret: 1
    end_fret: 2
    notes:
    - string: 4
      fret: 1
      pc: 3
      note: D#
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: b5
      role: fifth
      finger: 2
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: b7
      role: seventh
      finger: 2
    - string: 1
      fret: 2
      pc: 6
      note: F#
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
  - id: d-sharp-minor-7th-flat-5-voicing-2
    start_fret: 4
    end_fret: 6
    notes:
    - string: 5
      fret: 6
      pc: 3
      note: D#
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 4
      pc: 6
      note: F#
      interval: b3
      role: third
      finger: 1
    - string: 3
      fret: 6
      pc: 1
      note: C#
      interval: b7
      role: seventh
      finger: 3
    - string: 2
      fret: 4
      pc: 3
      note: D#
      interval: '1'
      role: root
      finger: 1
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: b5
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: d-sharp-minor-7th-flat-5-voicing-3
    start_fret: 5
    end_fret: 7
    notes:
    - string: 5
      fret: 6
      pc: 3
      note: D#
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 7
      pc: 9
      note: A
      interval: b5
      role: fifth
      finger: 3
    - string: 3
      fret: 6
      pc: 1
      note: C#
      interval: b7
      role: seventh
      finger: 2
    - string: 2
      fret: 7
      pc: 6
      note: F#
      interval: b3
      role: third
      finger: 3
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: b5
      role: fifth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (1.4 2.3 2.2 2.1).2
  midi_root: 63
musical_forms:
- form: pitch_set
  id: d-sharp-minor-7th-flat-5-chord-pitch-set
  pitch_classes:
  - 3
  - 6
  - 9
  - 1
  intervals:
  - '1'
  - b3
  - b5
  - b7
  spelling:
  - D#
  - F#
  - A
  - C#
- form: voicing
  id: d-sharp-minor-7th-flat-5-voicing-1
  position_ref: d-sharp-minor-7th-flat-5-voicing-1
  start_fret: 1
  tab: x-x-1-2-2-2
- form: voicing
  id: d-sharp-minor-7th-flat-5-voicing-2
  position_ref: d-sharp-minor-7th-flat-5-voicing-2
  start_fret: 4
  tab: x-6-4-6-4-5
- form: voicing
  id: d-sharp-minor-7th-flat-5-voicing-3
  position_ref: d-sharp-minor-7th-flat-5-voicing-3
  start_fret: 5
  tab: x-6-7-6-7-5
- form: tab_excerpt
  id: d-sharp-minor-7th-flat-5-chord-tab
  alphatex: \tempo 90 (1.4 2.3 2.2 2.1).2
---

# D# Minor 7th Flat 5 Chord

The D# minor 7th flat 5 chord is the half-diminished chord: a diminished triad plus a minor seventh, the ii chord of a minor key. Its tones are **D# – F# – A – C#**: D# (root), F# (minor third), A (diminished fifth), C# (minor seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | D# | root |
| b3 | F# | minor third |
| b5 | A | diminished fifth |
| b7 | C# | minor seventh |

**Playing it:** the primary shape is `x-x-1-2-2-2` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [D# Diminished Triad](/triads/d-sharp-diminished-triad.md)
