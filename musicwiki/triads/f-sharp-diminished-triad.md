---
type: Concept
slug: /triads/f-sharp-diminished-triad
title: F# Diminished Triad
description: 'The diminished triad built on F#: F#, A, C.'
tags:
- triad
- diminished
- f
parent: /triads
used_by:
- /chords/f-sharp-minor-7th-flat-5-chord
- /chords/f-sharp-diminished-7th-chord
theory:
  category: triad
  quality: diminished
  root:
    pc: 6
    name: F#
  intervals:
  - '1'
  - b3
  - b5
  semitones:
  - 0
  - 3
  - 6
  pitch_classes:
  - 6
  - 9
  - 0
  spelling:
  - F#
  - A
  - C
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
  emphasize_intervals:
  - '1'
  - b3
  - b5
  preferred_positions:
  - id: f-sharp-diminished-voicing-1
    start_fret: 1
    end_fret: 4
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
  - id: f-sharp-diminished-voicing-2
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
    mute_strings:
    - 1
    - 2
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: f-sharp-diminished-voicing-3
    start_fret: 4
    end_fret: 7
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
      fret: 7
      pc: 6
      note: F#
      interval: '1'
      role: root
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
  alphatex: \tempo 90 (2.6 0.5 4.4 2.3 1.2 2.1).2
  midi_root: 66
musical_forms:
- form: pitch_set
  id: f-sharp-diminished-triad-pitch-set
  pitch_classes:
  - 6
  - 9
  - 0
  intervals:
  - '1'
  - b3
  - b5
  spelling:
  - F#
  - A
  - C
- form: voicing
  id: f-sharp-diminished-voicing-1
  position_ref: f-sharp-diminished-voicing-1
  start_fret: 1
  tab: 2-0-4-2-1-2
- form: voicing
  id: f-sharp-diminished-voicing-2
  position_ref: f-sharp-diminished-voicing-2
  start_fret: 2
  tab: 2-3-4-2-x-x
- form: voicing
  id: f-sharp-diminished-voicing-3
  position_ref: f-sharp-diminished-voicing-3
  start_fret: 4
  tab: x-x-4-5-7-5
- form: tab_excerpt
  id: f-sharp-diminished-triad-tab
  alphatex: \tempo 90 (2.6 0.5 4.4 2.3 1.2 2.1).2
---

# F# Diminished Triad

The F# diminished triad is a tense, unstable chord of two stacked minor thirds that wants to resolve. Its tones are **F# – A – C**.

**Construction:** stacked thirds — a minor third (F#–A) under a minor third (A–C).

| Interval | Note | Name |
|---|---|---|
| 1 | F# | root |
| b3 | A | minor third |
| b5 | C | diminished fifth |

**Playing it:** the primary shape is `2-0-4-2-1-2` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Used by:** [F# Minor 7th Flat 5 Chord](/chords/f-sharp-minor-7th-flat-5-chord.md), [F# Diminished 7th Chord](/chords/f-sharp-diminished-7th-chord.md)
