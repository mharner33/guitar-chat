---
type: Concept
slug: /chords/d-diminished-7th-chord
title: D Diminished 7th Chord
description: 'The diminished 7th chord built on D: D, F, Ab, Cb.'
tags:
- chord
- diminished-7th
- d
parent: /chords
prerequisites:
- /triads/d-diminished-triad
theory:
  category: chord
  quality: diminished-7th
  root:
    pc: 2
    name: D
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
  - 2
  - 5
  - 8
  - 11
  spelling:
  - D
  - F
  - Ab
  - Cb
  enharmonic_policy: prefer_flats
  key_context: D
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
  - 2
  - 5
  - 8
  - 11
  emphasize_intervals:
  - '1'
  - b3
  - b5
  - bb7
  preferred_positions:
  - id: d-diminished-7th-voicing-1
    start_fret: 1
    end_fret: 1
    notes:
    - string: 4
      fret: 0
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 0
    - string: 3
      fret: 1
      pc: 8
      note: Ab
      interval: b5
      role: fifth
      finger: 1
    - string: 2
      fret: 0
      pc: 11
      note: Cb
      interval: bb7
      role: seventh
      finger: 0
    - string: 1
      fret: 1
      pc: 5
      note: F
      interval: b3
      role: third
      finger: 1
    mute_strings:
    - 5
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: d-diminished-7th-voicing-2
    start_fret: 3
    end_fret: 5
    notes:
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: b3
      role: third
      finger: 1
    - string: 3
      fret: 4
      pc: 11
      note: Cb
      interval: bb7
      role: seventh
      finger: 2
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 1
    - string: 1
      fret: 4
      pc: 8
      note: Ab
      interval: b5
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: d-diminished-7th-voicing-3
    start_fret: 4
    end_fret: 6
    notes:
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 6
      pc: 8
      note: Ab
      interval: b5
      role: fifth
      finger: 3
    - string: 3
      fret: 4
      pc: 11
      note: Cb
      interval: bb7
      role: seventh
      finger: 1
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: b3
      role: third
      finger: 3
    - string: 1
      fret: 4
      pc: 8
      note: Ab
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
  alphatex: \tempo 90 (0.4 1.3 0.2 1.1).2
  midi_root: 62
musical_forms:
- form: pitch_set
  id: d-diminished-7th-chord-pitch-set
  pitch_classes:
  - 2
  - 5
  - 8
  - 11
  intervals:
  - '1'
  - b3
  - b5
  - bb7
  spelling:
  - D
  - F
  - Ab
  - Cb
- form: voicing
  id: d-diminished-7th-voicing-1
  position_ref: d-diminished-7th-voicing-1
  start_fret: 1
  tab: x-x-0-1-0-1
- form: voicing
  id: d-diminished-7th-voicing-2
  position_ref: d-diminished-7th-voicing-2
  start_fret: 3
  tab: x-5-3-4-3-4
- form: voicing
  id: d-diminished-7th-voicing-3
  position_ref: d-diminished-7th-voicing-3
  start_fret: 4
  tab: x-5-6-4-6-4
- form: tab_excerpt
  id: d-diminished-7th-chord-tab
  alphatex: \tempo 90 (0.4 1.3 0.2 1.1).2
---

# D Diminished 7th Chord

The D diminished 7th chord is four stacked minor thirds — symmetrical and maximally tense, able to resolve many directions. Its tones are **D – F – Ab – Cb**: D (root), F (minor third), Ab (diminished fifth), Cb (diminished seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | D | root |
| b3 | F | minor third |
| b5 | Ab | diminished fifth |
| bb7 | Cb | diminished seventh |

**Playing it:** the primary shape is `x-x-0-1-0-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [D Diminished Triad](/triads/d-diminished-triad.md)
