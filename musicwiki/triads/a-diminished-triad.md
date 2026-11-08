---
type: Concept
slug: /triads/a-diminished-triad
title: A Diminished Triad
description: 'The diminished triad built on A: A, C, Eb.'
tags:
- triad
- diminished
- a
parent: /triads
used_by:
- /chords/a-minor-7th-flat-5-chord
- /chords/a-diminished-7th-chord
theory:
  category: triad
  quality: diminished
  root:
    pc: 9
    name: A
  intervals:
  - '1'
  - b3
  - b5
  semitones:
  - 0
  - 3
  - 6
  pitch_classes:
  - 9
  - 0
  - 3
  spelling:
  - A
  - C
  - Eb
  enharmonic_policy: prefer_flats
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
  - 0
  - 3
  emphasize_intervals:
  - '1'
  - b3
  - b5
  preferred_positions:
  - id: a-diminished-voicing-1
    start_fret: 1
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
      fret: 1
      pc: 3
      note: Eb
      interval: b5
      role: fifth
      finger: 1
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 2
    - string: 2
      fret: 1
      pc: 0
      note: C
      interval: b3
      role: third
      finger: 1
    mute_strings:
    - 1
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: a-diminished-voicing-2
    start_fret: 4
    end_fret: 7
    notes:
    - string: 6
      fret: 5
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 6
      pc: 3
      note: Eb
      interval: b5
      role: fifth
      finger: 3
    - string: 4
      fret: 7
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 4
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: b3
      role: third
      finger: 2
    - string: 2
      fret: 4
      pc: 3
      note: Eb
      interval: b5
      role: fifth
      finger: 1
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
  - id: a-diminished-voicing-3
    start_fret: 3
    end_fret: 7
    notes:
    - string: 6
      fret: 5
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 3
      pc: 0
      note: C
      interval: b3
      role: third
      finger: 1
    - string: 4
      fret: 7
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 4
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: b3
      role: third
      finger: 3
    - string: 2
      fret: 4
      pc: 3
      note: Eb
      interval: b5
      role: fifth
      finger: 2
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 3
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (0.5 1.4 2.3 1.2).2
  midi_root: 69
musical_forms:
- form: pitch_set
  id: a-diminished-triad-pitch-set
  pitch_classes:
  - 9
  - 0
  - 3
  intervals:
  - '1'
  - b3
  - b5
  spelling:
  - A
  - C
  - Eb
- form: voicing
  id: a-diminished-voicing-1
  position_ref: a-diminished-voicing-1
  start_fret: 1
  tab: x-0-1-2-1-x
- form: voicing
  id: a-diminished-voicing-2
  position_ref: a-diminished-voicing-2
  start_fret: 4
  tab: 5-6-7-5-4-5
- form: voicing
  id: a-diminished-voicing-3
  position_ref: a-diminished-voicing-3
  start_fret: 3
  tab: 5-3-7-5-4-5
- form: tab_excerpt
  id: a-diminished-triad-tab
  alphatex: \tempo 90 (0.5 1.4 2.3 1.2).2
---

# A Diminished Triad

The A diminished triad is a tense, unstable chord of two stacked minor thirds that wants to resolve. Its tones are **A – C – Eb**.

**Construction:** stacked thirds — a minor third (A–C) under a minor third (C–Eb).

| Interval | Note | Name |
|---|---|---|
| 1 | A | root |
| b3 | C | minor third |
| b5 | Eb | diminished fifth |

**Playing it:** the primary shape is `x-0-1-2-1-x` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Used by:** [A Minor 7th Flat 5 Chord](/chords/a-minor-7th-flat-5-chord.md), [A Diminished 7th Chord](/chords/a-diminished-7th-chord.md)
