---
type: Concept
slug: /chords/a-sharp-diminished-7th-chord
title: A# Diminished 7th Chord
description: 'The diminished 7th chord built on A#: A#, C#, E, G.'
tags:
- chord
- diminished-7th
- a
parent: /chords
prerequisites:
- /triads/a-sharp-diminished-triad
theory:
  category: chord
  quality: diminished-7th
  root:
    pc: 10
    name: A#
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
  - 10
  - 1
  - 4
  - 7
  spelling:
  - A#
  - C#
  - E
  - G
  enharmonic_policy: prefer_sharps
  key_context: A#
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
  - 10
  - 1
  - 4
  - 7
  emphasize_intervals:
  - '1'
  - b3
  - b5
  - bb7
  preferred_positions:
  - id: a-sharp-diminished-7th-voicing-1
    start_fret: 1
    end_fret: 2
    notes:
    - string: 5
      fret: 1
      pc: 10
      note: A#
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: b5
      role: fifth
      finger: 2
    - string: 3
      fret: 0
      pc: 7
      note: G
      interval: bb7
      role: seventh
      finger: 0
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: b3
      role: third
      finger: 2
    - string: 1
      fret: 0
      pc: 4
      note: E
      interval: b5
      role: fifth
      finger: 0
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: a-sharp-diminished-7th-voicing-2
    start_fret: 4
    end_fret: 6
    notes:
    - string: 6
      fret: 6
      pc: 10
      note: A#
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 4
      pc: 1
      note: C#
      interval: b3
      role: third
      finger: 1
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: bb7
      role: seventh
      finger: 2
    - string: 3
      fret: 6
      pc: 1
      note: C#
      interval: b3
      role: third
      finger: 3
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: b5
      role: fifth
      finger: 2
    - string: 1
      fret: 6
      pc: 10
      note: A#
      interval: '1'
      role: root
      finger: 3
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: a-sharp-diminished-7th-voicing-3
    start_fret: 3
    end_fret: 6
    notes:
    - string: 6
      fret: 6
      pc: 10
      note: A#
      interval: '1'
      role: root
      finger: 4
    - string: 5
      fret: 4
      pc: 1
      note: C#
      interval: b3
      role: third
      finger: 2
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: bb7
      role: seventh
      finger: 3
    - string: 3
      fret: 3
      pc: 10
      note: A#
      interval: '1'
      role: root
      finger: 1
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: b5
      role: fifth
      finger: 3
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: bb7
      role: seventh
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (1.5 2.4 0.3 2.2 0.1).2
  midi_root: 70
musical_forms:
- form: pitch_set
  id: a-sharp-diminished-7th-chord-pitch-set
  pitch_classes:
  - 10
  - 1
  - 4
  - 7
  intervals:
  - '1'
  - b3
  - b5
  - bb7
  spelling:
  - A#
  - C#
  - E
  - G
- form: voicing
  id: a-sharp-diminished-7th-voicing-1
  position_ref: a-sharp-diminished-7th-voicing-1
  start_fret: 1
  tab: x-1-2-0-2-0
- form: voicing
  id: a-sharp-diminished-7th-voicing-2
  position_ref: a-sharp-diminished-7th-voicing-2
  start_fret: 4
  tab: 6-4-5-6-5-6
- form: voicing
  id: a-sharp-diminished-7th-voicing-3
  position_ref: a-sharp-diminished-7th-voicing-3
  start_fret: 3
  tab: 6-4-5-3-5-3
- form: tab_excerpt
  id: a-sharp-diminished-7th-chord-tab
  alphatex: \tempo 90 (1.5 2.4 0.3 2.2 0.1).2
---

# A# Diminished 7th Chord

The A# diminished 7th chord is four stacked minor thirds — symmetrical and maximally tense, able to resolve many directions. Its tones are **A# – C# – E – G**: A# (root), C# (minor third), E (diminished fifth), G (diminished seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | A# | root |
| b3 | C# | minor third |
| b5 | E | diminished fifth |
| bb7 | G | diminished seventh |

**Playing it:** the primary shape is `x-1-2-0-2-0` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [A# Diminished Triad](/triads/a-sharp-diminished-triad.md)
