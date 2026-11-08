---
type: Concept
slug: /triads/e-diminished-triad
title: E Diminished Triad
description: 'The diminished triad built on E: E, G, Bb.'
tags:
- triad
- diminished
- e
parent: /triads
used_by:
- /chords/e-minor-7th-flat-5-chord
- /chords/e-diminished-7th-chord
theory:
  category: triad
  quality: diminished
  root:
    pc: 4
    name: E
  intervals:
  - '1'
  - b3
  - b5
  semitones:
  - 0
  - 3
  - 6
  pitch_classes:
  - 4
  - 7
  - 10
  spelling:
  - E
  - G
  - Bb
  enharmonic_policy: prefer_flats
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
  - 7
  - 10
  emphasize_intervals:
  - '1'
  - b3
  - b5
  preferred_positions:
  - id: e-diminished-voicing-1
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
      fret: 1
      pc: 10
      note: Bb
      interval: b5
      role: fifth
      finger: 1
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 2
    - string: 3
      fret: 0
      pc: 7
      note: G
      interval: b3
      role: third
      finger: 0
    mute_strings:
    - 1
    - 2
    confidence: 0.75
    tags:
    - computed
    - open
  - id: e-diminished-voicing-2
    start_fret: 2
    end_fret: 5
    notes:
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: b5
      role: fifth
      finger: 2
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 3
    - string: 1
      fret: 3
      pc: 7
      note: G
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
  - id: e-diminished-voicing-3
    start_fret: 3
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
      fret: 5
      pc: 7
      note: G
      interval: b3
      role: third
      finger: 2
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: b5
      role: fifth
      finger: 1
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 2
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: b3
      role: third
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (0.6 1.5 2.4 0.3).2
  midi_root: 64
musical_forms:
- form: pitch_set
  id: e-diminished-triad-pitch-set
  pitch_classes:
  - 4
  - 7
  - 10
  intervals:
  - '1'
  - b3
  - b5
  spelling:
  - E
  - G
  - Bb
- form: voicing
  id: e-diminished-voicing-1
  position_ref: e-diminished-voicing-1
  start_fret: 1
  tab: 0-1-2-0-x-x
- form: voicing
  id: e-diminished-voicing-2
  position_ref: e-diminished-voicing-2
  start_fret: 2
  tab: x-x-2-3-5-3
- form: voicing
  id: e-diminished-voicing-3
  position_ref: e-diminished-voicing-3
  start_fret: 3
  tab: x-7-5-3-5-3
- form: tab_excerpt
  id: e-diminished-triad-tab
  alphatex: \tempo 90 (0.6 1.5 2.4 0.3).2
---

# E Diminished Triad

The E diminished triad is a tense, unstable chord of two stacked minor thirds that wants to resolve. Its tones are **E – G – Bb**.

**Construction:** stacked thirds — a minor third (E–G) under a minor third (G–Bb).

| Interval | Note | Name |
|---|---|---|
| 1 | E | root |
| b3 | G | minor third |
| b5 | Bb | diminished fifth |

**Playing it:** the primary shape is `0-1-2-0-x-x` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Used by:** [E Minor 7th Flat 5 Chord](/chords/e-minor-7th-flat-5-chord.md), [E Diminished 7th Chord](/chords/e-diminished-7th-chord.md)
