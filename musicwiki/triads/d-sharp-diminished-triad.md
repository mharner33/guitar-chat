---
type: Concept
slug: /triads/d-sharp-diminished-triad
title: D# Diminished Triad
description: 'The diminished triad built on D#: D#, F#, A.'
tags:
- triad
- diminished
- d
parent: /triads
used_by:
- /chords/d-sharp-minor-7th-flat-5-chord
- /chords/d-sharp-diminished-7th-chord
theory:
  category: triad
  quality: diminished
  root:
    pc: 3
    name: D#
  intervals:
  - '1'
  - b3
  - b5
  semitones:
  - 0
  - 3
  - 6
  pitch_classes:
  - 3
  - 6
  - 9
  spelling:
  - D#
  - F#
  - A
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
  emphasize_intervals:
  - '1'
  - b3
  - b5
  preferred_positions:
  - id: d-sharp-diminished-voicing-1
    start_fret: 1
    end_fret: 4
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
      fret: 4
      pc: 3
      note: D#
      interval: '1'
      role: root
      finger: 3
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
  - id: d-sharp-diminished-voicing-2
    start_fret: 2
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
      finger: 2
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: b5
      role: fifth
      finger: 1
    - string: 2
      fret: 4
      pc: 3
      note: D#
      interval: '1'
      role: root
      finger: 2
    - string: 1
      fret: 2
      pc: 6
      note: F#
      interval: b3
      role: third
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: d-sharp-diminished-voicing-3
    start_fret: 5
    end_fret: 8
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
      fret: 8
      pc: 3
      note: D#
      interval: '1'
      role: root
      finger: 4
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
  alphatex: \tempo 90 (1.4 2.3 4.2 2.1).2
  midi_root: 63
musical_forms:
- form: pitch_set
  id: d-sharp-diminished-triad-pitch-set
  pitch_classes:
  - 3
  - 6
  - 9
  intervals:
  - '1'
  - b3
  - b5
  spelling:
  - D#
  - F#
  - A
- form: voicing
  id: d-sharp-diminished-voicing-1
  position_ref: d-sharp-diminished-voicing-1
  start_fret: 1
  tab: x-x-1-2-4-2
- form: voicing
  id: d-sharp-diminished-voicing-2
  position_ref: d-sharp-diminished-voicing-2
  start_fret: 2
  tab: x-6-4-2-4-2
- form: voicing
  id: d-sharp-diminished-voicing-3
  position_ref: d-sharp-diminished-voicing-3
  start_fret: 5
  tab: x-6-7-8-7-5
- form: tab_excerpt
  id: d-sharp-diminished-triad-tab
  alphatex: \tempo 90 (1.4 2.3 4.2 2.1).2
---

# D# Diminished Triad

The D# diminished triad is a tense, unstable chord of two stacked minor thirds that wants to resolve. Its tones are **D# – F# – A**.

**Construction:** stacked thirds — a minor third (D#–F#) under a minor third (F#–A).

| Interval | Note | Name |
|---|---|---|
| 1 | D# | root |
| b3 | F# | minor third |
| b5 | A | diminished fifth |

**Playing it:** the primary shape is `x-x-1-2-4-2` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Used by:** [D# Minor 7th Flat 5 Chord](/chords/d-sharp-minor-7th-flat-5-chord.md), [D# Diminished 7th Chord](/chords/d-sharp-diminished-7th-chord.md)
