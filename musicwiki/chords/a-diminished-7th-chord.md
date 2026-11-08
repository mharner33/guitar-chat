---
type: Concept
slug: /chords/a-diminished-7th-chord
title: A Diminished 7th Chord
description: 'The diminished 7th chord built on A: A, C, Eb, Gb.'
tags:
- chord
- diminished-7th
- a
parent: /chords
prerequisites:
- /triads/a-diminished-triad
theory:
  category: chord
  quality: diminished-7th
  root:
    pc: 9
    name: A
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
  - 9
  - 0
  - 3
  - 6
  spelling:
  - A
  - C
  - Eb
  - Gb
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
  - 6
  emphasize_intervals:
  - '1'
  - b3
  - b5
  - bb7
  preferred_positions:
  - id: a-diminished-7th-voicing-1
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
    - string: 1
      fret: 2
      pc: 6
      note: Gb
      interval: bb7
      role: seventh
      finger: 2
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: a-diminished-7th-voicing-2
    start_fret: 3
    end_fret: 5
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
      fret: 4
      pc: 6
      note: Gb
      interval: bb7
      role: seventh
      finger: 2
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
  - id: a-diminished-7th-voicing-3
    start_fret: 2
    end_fret: 5
    notes:
    - string: 6
      fret: 5
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 0
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 0
    - string: 4
      fret: 4
      pc: 6
      note: Gb
      interval: bb7
      role: seventh
      finger: 2
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
      fret: 2
      pc: 6
      note: Gb
      interval: bb7
      role: seventh
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
playback:
  alphatex: \tempo 90 (0.5 1.4 2.3 1.2 2.1).2
  midi_root: 69
musical_forms:
- form: pitch_set
  id: a-diminished-7th-chord-pitch-set
  pitch_classes:
  - 9
  - 0
  - 3
  - 6
  intervals:
  - '1'
  - b3
  - b5
  - bb7
  spelling:
  - A
  - C
  - Eb
  - Gb
- form: voicing
  id: a-diminished-7th-voicing-1
  position_ref: a-diminished-7th-voicing-1
  start_fret: 1
  tab: x-0-1-2-1-2
- form: voicing
  id: a-diminished-7th-voicing-2
  position_ref: a-diminished-7th-voicing-2
  start_fret: 3
  tab: 5-3-4-5-4-5
- form: voicing
  id: a-diminished-7th-voicing-3
  position_ref: a-diminished-7th-voicing-3
  start_fret: 2
  tab: 5-0-4-5-4-2
- form: tab_excerpt
  id: a-diminished-7th-chord-tab
  alphatex: \tempo 90 (0.5 1.4 2.3 1.2 2.1).2
---

# A Diminished 7th Chord

The A diminished 7th chord is four stacked minor thirds — symmetrical and maximally tense, able to resolve many directions. Its tones are **A – C – Eb – Gb**: A (root), C (minor third), Eb (diminished fifth), Gb (diminished seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | A | root |
| b3 | C | minor third |
| b5 | Eb | diminished fifth |
| bb7 | Gb | diminished seventh |

**Playing it:** the primary shape is `x-0-1-2-1-2` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [A Diminished Triad](/triads/a-diminished-triad.md)
