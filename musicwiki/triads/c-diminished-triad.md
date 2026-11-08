---
type: Concept
slug: /triads/c-diminished-triad
title: C Diminished Triad
description: 'The diminished triad built on C: C, Eb, Gb.'
tags:
- triad
- diminished
- c
parent: /triads
used_by:
- /chords/c-minor-7th-flat-5-chord
- /chords/c-diminished-7th-chord
theory:
  category: triad
  quality: diminished
  root:
    pc: 0
    name: C
  intervals:
  - '1'
  - b3
  - b5
  semitones:
  - 0
  - 3
  - 6
  pitch_classes:
  - 0
  - 3
  - 6
  spelling:
  - C
  - Eb
  - Gb
  enharmonic_policy: prefer_flats
  key_context: C
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
  - 0
  - 3
  - 6
  emphasize_intervals:
  - '1'
  - b3
  - b5
  preferred_positions:
  - id: c-diminished-voicing-1
    start_fret: 2
    end_fret: 5
    notes:
    - string: 5
      fret: 3
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 4
      pc: 6
      note: Gb
      interval: b5
      role: fifth
      finger: 3
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 4
    - string: 2
      fret: 4
      pc: 3
      note: Eb
      interval: b3
      role: third
      finger: 3
    - string: 1
      fret: 2
      pc: 6
      note: Gb
      interval: b5
      role: fifth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: c-diminished-voicing-2
    start_fret: 1
    end_fret: 5
    notes:
    - string: 5
      fret: 3
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 1
      pc: 3
      note: Eb
      interval: b3
      role: third
      finger: 1
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 4
    - string: 2
      fret: 1
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 1
    - string: 1
      fret: 2
      pc: 6
      note: Gb
      interval: b5
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: c-diminished-voicing-3
    start_fret: 3
    end_fret: 5
    notes:
    - string: 5
      fret: 3
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 4
      pc: 6
      note: Gb
      interval: b5
      role: fifth
      finger: 2
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 3
    - string: 2
      fret: 4
      pc: 3
      note: Eb
      interval: b3
      role: third
      finger: 2
    mute_strings:
    - 1
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (3.5 4.4 5.3 4.2 2.1).2
  midi_root: 60
musical_forms:
- form: pitch_set
  id: c-diminished-triad-pitch-set
  pitch_classes:
  - 0
  - 3
  - 6
  intervals:
  - '1'
  - b3
  - b5
  spelling:
  - C
  - Eb
  - Gb
- form: voicing
  id: c-diminished-voicing-1
  position_ref: c-diminished-voicing-1
  start_fret: 2
  tab: x-3-4-5-4-2
- form: voicing
  id: c-diminished-voicing-2
  position_ref: c-diminished-voicing-2
  start_fret: 1
  tab: x-3-1-5-1-2
- form: voicing
  id: c-diminished-voicing-3
  position_ref: c-diminished-voicing-3
  start_fret: 3
  tab: x-3-4-5-4-x
- form: tab_excerpt
  id: c-diminished-triad-tab
  alphatex: \tempo 90 (3.5 4.4 5.3 4.2 2.1).2
---

# C Diminished Triad

The C diminished triad is a tense, unstable chord of two stacked minor thirds that wants to resolve. Its tones are **C – Eb – Gb**.

**Construction:** stacked thirds — a minor third (C–Eb) under a minor third (Eb–Gb).

| Interval | Note | Name |
|---|---|---|
| 1 | C | root |
| b3 | Eb | minor third |
| b5 | Gb | diminished fifth |

**Playing it:** the primary shape is `x-3-4-5-4-2` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Used by:** [C Minor 7th Flat 5 Chord](/chords/c-minor-7th-flat-5-chord.md), [C Diminished 7th Chord](/chords/c-diminished-7th-chord.md)
