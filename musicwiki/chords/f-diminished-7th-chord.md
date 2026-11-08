---
type: Concept
slug: /chords/f-diminished-7th-chord
title: F Diminished 7th Chord
description: 'The diminished 7th chord built on F: F, Ab, Cb, Ebb.'
tags:
- chord
- diminished-7th
- f
parent: /chords
prerequisites:
- /triads/f-diminished-triad
theory:
  category: chord
  quality: diminished-7th
  root:
    pc: 5
    name: F
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
  - 5
  - 8
  - 11
  - 2
  spelling:
  - F
  - Ab
  - Cb
  - Ebb
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
  - 2
  emphasize_intervals:
  - '1'
  - b3
  - b5
  - bb7
  preferred_positions:
  - id: f-diminished-7th-voicing-1
    start_fret: 1
    end_fret: 2
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
      fret: 0
      pc: 2
      note: Ebb
      interval: bb7
      role: seventh
      finger: 0
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
  - id: f-diminished-7th-voicing-2
    start_fret: 3
    end_fret: 4
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
      fret: 3
      pc: 2
      note: Ebb
      interval: bb7
      role: seventh
      finger: 1
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
  - id: f-diminished-7th-voicing-3
    start_fret: 6
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
      finger: 1
    - string: 3
      fret: 7
      pc: 2
      note: Ebb
      interval: bb7
      role: seventh
      finger: 2
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 1
    - string: 1
      fret: 7
      pc: 11
      note: Cb
      interval: b5
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (1.6 2.5 0.4 1.3 0.2 1.1).2
  midi_root: 65
musical_forms:
- form: pitch_set
  id: f-diminished-7th-chord-pitch-set
  pitch_classes:
  - 5
  - 8
  - 11
  - 2
  intervals:
  - '1'
  - b3
  - b5
  - bb7
  spelling:
  - F
  - Ab
  - Cb
  - Ebb
- form: voicing
  id: f-diminished-7th-voicing-1
  position_ref: f-diminished-7th-voicing-1
  start_fret: 1
  tab: 1-2-0-1-0-1
- form: voicing
  id: f-diminished-7th-voicing-2
  position_ref: f-diminished-7th-voicing-2
  start_fret: 3
  tab: x-x-3-4-3-4
- form: voicing
  id: f-diminished-7th-voicing-3
  position_ref: f-diminished-7th-voicing-3
  start_fret: 6
  tab: x-8-6-7-6-7
- form: tab_excerpt
  id: f-diminished-7th-chord-tab
  alphatex: \tempo 90 (1.6 2.5 0.4 1.3 0.2 1.1).2
---

# F Diminished 7th Chord

The F diminished 7th chord is four stacked minor thirds — symmetrical and maximally tense, able to resolve many directions. Its tones are **F – Ab – Cb – Ebb**: F (root), Ab (minor third), Cb (diminished fifth), Ebb (diminished seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | F | root |
| b3 | Ab | minor third |
| b5 | Cb | diminished fifth |
| bb7 | Ebb | diminished seventh |

**Playing it:** the primary shape is `1-2-0-1-0-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [F Diminished Triad](/triads/f-diminished-triad.md)
