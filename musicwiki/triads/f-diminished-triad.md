---
type: Concept
slug: /triads/f-diminished-triad
title: F Diminished Triad
description: 'The diminished triad built on F: F, Ab, Cb.'
tags:
- triad
- diminished
- f
parent: /triads
used_by:
- /chords/f-minor-7th-flat-5-chord
- /chords/f-diminished-7th-chord
theory:
  category: triad
  quality: diminished
  root:
    pc: 5
    name: F
  intervals:
  - '1'
  - b3
  - b5
  semitones:
  - 0
  - 3
  - 6
  pitch_classes:
  - 5
  - 8
  - 11
  spelling:
  - F
  - Ab
  - Cb
  enharmonic_policy: prefer_flats
  key_context: F
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
  - 5
  - 8
  - 11
  emphasize_intervals:
  - '1'
  - b3
  - b5
  preferred_positions:
  - id: f-diminished-voicing-1
    start_fret: 1
    end_fret: 3
    notes:
    - string: 6
      fret: 1
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 2
      pc: 11
      note: Cb
      interval: b5
      role: fifth
      finger: 2
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 3
    - string: 3
      fret: 1
      pc: 8
      note: Ab
      interval: b3
      role: third
      finger: 1
    - string: 2
      fret: 0
      pc: 11
      note: Cb
      interval: b5
      role: fifth
      finger: 0
    - string: 1
      fret: 1
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: f-diminished-voicing-2
    start_fret: 3
    end_fret: 6
    notes:
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 4
      pc: 11
      note: Cb
      interval: b5
      role: fifth
      finger: 2
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 3
    - string: 1
      fret: 4
      pc: 8
      note: Ab
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
  - id: f-diminished-voicing-3
    start_fret: 4
    end_fret: 8
    notes:
    - string: 5
      fret: 8
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 6
      pc: 8
      note: Ab
      interval: b3
      role: third
      finger: 2
    - string: 3
      fret: 4
      pc: 11
      note: Cb
      interval: b5
      role: fifth
      finger: 1
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 2
    - string: 1
      fret: 4
      pc: 8
      note: Ab
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
  alphatex: \tempo 90 (1.6 2.5 3.4 1.3 0.2 1.1).2
  midi_root: 65
musical_forms:
- form: pitch_set
  id: f-diminished-triad-pitch-set
  pitch_classes:
  - 5
  - 8
  - 11
  intervals:
  - '1'
  - b3
  - b5
  spelling:
  - F
  - Ab
  - Cb
- form: voicing
  id: f-diminished-voicing-1
  position_ref: f-diminished-voicing-1
  start_fret: 1
  tab: 1-2-3-1-0-1
- form: voicing
  id: f-diminished-voicing-2
  position_ref: f-diminished-voicing-2
  start_fret: 3
  tab: x-x-3-4-6-4
- form: voicing
  id: f-diminished-voicing-3
  position_ref: f-diminished-voicing-3
  start_fret: 4
  tab: x-8-6-4-6-4
- form: tab_excerpt
  id: f-diminished-triad-tab
  alphatex: \tempo 90 (1.6 2.5 3.4 1.3 0.2 1.1).2
---

# F Diminished Triad

The F diminished triad is a tense, unstable chord of two stacked minor thirds that wants to resolve. Its tones are **F – Ab – Cb**.

**Construction:** stacked thirds — a minor third (F–Ab) under a minor third (Ab–Cb).

| Interval | Note | Name |
|---|---|---|
| 1 | F | root |
| b3 | Ab | minor third |
| b5 | Cb | diminished fifth |

**Playing it:** the primary shape is `1-2-3-1-0-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Used by:** [F Minor 7th Flat 5 Chord](/chords/f-minor-7th-flat-5-chord.md), [F Diminished 7th Chord](/chords/f-diminished-7th-chord.md)
