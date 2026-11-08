---
type: Concept
slug: /triads/b-diminished-triad
title: B Diminished Triad
description: 'The diminished triad built on B: B, D, F.'
tags:
- triad
- diminished
- b
parent: /triads
used_by:
- /chords/b-minor-7th-flat-5-chord
- /chords/b-diminished-7th-chord
theory:
  category: triad
  quality: diminished
  root:
    pc: 11
    name: B
  intervals:
  - '1'
  - b3
  - b5
  semitones:
  - 0
  - 3
  - 6
  pitch_classes:
  - 11
  - 2
  - 5
  spelling:
  - B
  - D
  - F
  enharmonic_policy: prefer_sharps
  key_context: B
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
  - 11
  - 2
  - 5
  emphasize_intervals:
  - '1'
  - b3
  - b5
  preferred_positions:
  - id: b-diminished-voicing-1
    start_fret: 1
    end_fret: 4
    notes:
    - string: 5
      fret: 2
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 0
      pc: 2
      note: D
      interval: b3
      role: third
      finger: 0
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 3
    - string: 2
      fret: 0
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 0
    - string: 1
      fret: 1
      pc: 5
      note: F
      interval: b5
      role: fifth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
  - id: b-diminished-voicing-2
    start_fret: 2
    end_fret: 4
    notes:
    - string: 5
      fret: 2
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: b5
      role: fifth
      finger: 2
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 3
    - string: 2
      fret: 3
      pc: 2
      note: D
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
  - id: b-diminished-voicing-3
    start_fret: 3
    end_fret: 7
    notes:
    - string: 6
      fret: 7
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 4
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: b3
      role: third
      finger: 3
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: b5
      role: fifth
      finger: 1
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 2
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: b3
      role: third
      finger: 1
    - string: 1
      fret: 7
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 4
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (2.5 0.4 4.3 0.2 1.1).2
  midi_root: 71
musical_forms:
- form: pitch_set
  id: b-diminished-triad-pitch-set
  pitch_classes:
  - 11
  - 2
  - 5
  intervals:
  - '1'
  - b3
  - b5
  spelling:
  - B
  - D
  - F
- form: voicing
  id: b-diminished-voicing-1
  position_ref: b-diminished-voicing-1
  start_fret: 1
  tab: x-2-0-4-0-1
- form: voicing
  id: b-diminished-voicing-2
  position_ref: b-diminished-voicing-2
  start_fret: 2
  tab: x-2-3-4-3-x
- form: voicing
  id: b-diminished-voicing-3
  position_ref: b-diminished-voicing-3
  start_fret: 3
  tab: 7-5-3-4-3-7
- form: tab_excerpt
  id: b-diminished-triad-tab
  alphatex: \tempo 90 (2.5 0.4 4.3 0.2 1.1).2
---

# B Diminished Triad

The B diminished triad is a tense, unstable chord of two stacked minor thirds that wants to resolve. Its tones are **B – D – F**.

**Construction:** stacked thirds — a minor third (B–D) under a minor third (D–F).

| Interval | Note | Name |
|---|---|---|
| 1 | B | root |
| b3 | D | minor third |
| b5 | F | diminished fifth |

**Playing it:** the primary shape is `x-2-0-4-0-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Used by:** [B Minor 7th Flat 5 Chord](/chords/b-minor-7th-flat-5-chord.md), [B Diminished 7th Chord](/chords/b-diminished-7th-chord.md)
