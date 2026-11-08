---
type: Concept
slug: /chords/f-sharp-diminished-7th-chord
title: F# Diminished 7th Chord
description: 'The diminished 7th chord built on F#: F#, A, C, Eb.'
tags:
- chord
- diminished-7th
- f
parent: /chords
prerequisites:
- /triads/f-sharp-diminished-triad
theory:
  category: chord
  quality: diminished-7th
  root:
    pc: 6
    name: F#
  intervals:
  - '1'
  - b3
  - b5
  - bb7
  semitones:
  - 0
  - 3
  - 6
  - 9
  pitch_classes:
  - 6
  - 9
  - 0
  - 3
  spelling:
  - F#
  - A
  - C
  - Eb
  enharmonic_policy: prefer_flats
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
  - 3
  emphasize_intervals:
  - '1'
  - b3
  - b5
  - bb7
  preferred_positions:
  - id: f-sharp-diminished-7th-voicing-1
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
      fret: 1
      pc: 3
      note: Eb
      interval: bb7
      role: seventh
      finger: 1
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
      fret: 2
      pc: 6
      note: F#
      interval: '1'
      role: root
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: f-sharp-diminished-7th-voicing-2
    start_fret: 2
    end_fret: 4
    notes:
    - string: 6
      fret: 2
      pc: 6
      note: F#
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 3
      pc: 0
      note: C
      interval: b5
      role: fifth
      finger: 2
    - string: 4
      fret: 4
      pc: 6
      note: F#
      interval: '1'
      role: root
      finger: 3
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: b3
      role: third
      finger: 1
    - string: 2
      fret: 4
      pc: 3
      note: Eb
      interval: bb7
      role: seventh
      finger: 3
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
    - barre
  - id: f-sharp-diminished-7th-voicing-3
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
      fret: 4
      pc: 3
      note: Eb
      interval: bb7
      role: seventh
      finger: 1
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
  alphatex: \tempo 90 (2.6 0.5 1.4 2.3 1.2 2.1).2
  midi_root: 66
musical_forms:
- form: pitch_set
  id: f-sharp-diminished-7th-chord-pitch-set
  pitch_classes:
  - 6
  - 9
  - 0
  - 3
  intervals:
  - '1'
  - b3
  - b5
  - bb7
  spelling:
  - F#
  - A
  - C
  - Eb
- form: voicing
  id: f-sharp-diminished-7th-voicing-1
  position_ref: f-sharp-diminished-7th-voicing-1
  start_fret: 1
  tab: 2-0-1-2-1-2
- form: voicing
  id: f-sharp-diminished-7th-voicing-2
  position_ref: f-sharp-diminished-7th-voicing-2
  start_fret: 2
  tab: 2-3-4-2-4-2
- form: voicing
  id: f-sharp-diminished-7th-voicing-3
  position_ref: f-sharp-diminished-7th-voicing-3
  start_fret: 4
  tab: x-x-4-5-4-5
- form: tab_excerpt
  id: f-sharp-diminished-7th-chord-tab
  alphatex: \tempo 90 (2.6 0.5 1.4 2.3 1.2 2.1).2
---

# F# Diminished 7th Chord

The F# diminished 7th chord is four stacked minor thirds — symmetrical and maximally tense, able to resolve many directions. Its tones are **F# – A – C – Eb**: F# (root), A (minor third), C (diminished fifth), Eb (diminished seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | F# | root |
| b3 | A | minor third |
| b5 | C | diminished fifth |
| bb7 | Eb | diminished seventh |

**Playing it:** the primary shape is `2-0-1-2-1-2` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [F# Diminished Triad](/triads/f-sharp-diminished-triad.md)
