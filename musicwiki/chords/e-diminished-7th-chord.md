---
type: Concept
slug: /chords/e-diminished-7th-chord
title: E Diminished 7th Chord
description: 'The diminished 7th chord built on E: E, G, Bb, Db.'
tags:
- chord
- diminished-7th
- e
parent: /chords
prerequisites:
- /triads/e-diminished-triad
theory:
  category: chord
  quality: diminished-7th
  root:
    pc: 4
    name: E
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
  - 4
  - 7
  - 10
  - 1
  spelling:
  - E
  - G
  - Bb
  - Db
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
  - 1
  emphasize_intervals:
  - '1'
  - b3
  - b5
  - bb7
  preferred_positions:
  - id: e-diminished-7th-voicing-1
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
    - string: 2
      fret: 2
      pc: 1
      note: Db
      interval: bb7
      role: seventh
      finger: 2
    - string: 1
      fret: 0
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 0
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: e-diminished-7th-voicing-2
    start_fret: 2
    end_fret: 4
    notes:
    - string: 6
      fret: 0
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 0
    - string: 5
      fret: 4
      pc: 1
      note: Db
      interval: bb7
      role: seventh
      finger: 3
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
      fret: 2
      pc: 1
      note: Db
      interval: bb7
      role: seventh
      finger: 1
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: b3
      role: third
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: e-diminished-7th-voicing-3
    start_fret: 5
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
      finger: 1
    - string: 3
      fret: 6
      pc: 1
      note: Db
      interval: bb7
      role: seventh
      finger: 2
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 1
    - string: 1
      fret: 6
      pc: 10
      note: Bb
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
  alphatex: \tempo 90 (0.6 1.5 2.4 0.3 2.2 0.1).2
  midi_root: 64
musical_forms:
- form: pitch_set
  id: e-diminished-7th-chord-pitch-set
  pitch_classes:
  - 4
  - 7
  - 10
  - 1
  intervals:
  - '1'
  - b3
  - b5
  - bb7
  spelling:
  - E
  - G
  - Bb
  - Db
- form: voicing
  id: e-diminished-7th-voicing-1
  position_ref: e-diminished-7th-voicing-1
  start_fret: 1
  tab: 0-1-2-0-2-0
- form: voicing
  id: e-diminished-7th-voicing-2
  position_ref: e-diminished-7th-voicing-2
  start_fret: 2
  tab: 0-4-2-3-2-3
- form: voicing
  id: e-diminished-7th-voicing-3
  position_ref: e-diminished-7th-voicing-3
  start_fret: 5
  tab: x-7-5-6-5-6
- form: tab_excerpt
  id: e-diminished-7th-chord-tab
  alphatex: \tempo 90 (0.6 1.5 2.4 0.3 2.2 0.1).2
---

# E Diminished 7th Chord

The E diminished 7th chord is four stacked minor thirds — symmetrical and maximally tense, able to resolve many directions. Its tones are **E – G – Bb – Db**: E (root), G (minor third), Bb (diminished fifth), Db (diminished seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | E | root |
| b3 | G | minor third |
| b5 | Bb | diminished fifth |
| bb7 | Db | diminished seventh |

**Playing it:** the primary shape is `0-1-2-0-2-0` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [E Diminished Triad](/triads/e-diminished-triad.md)
