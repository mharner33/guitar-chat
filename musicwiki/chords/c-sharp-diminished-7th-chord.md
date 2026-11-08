---
type: Concept
slug: /chords/c-sharp-diminished-7th-chord
title: C# Diminished 7th Chord
description: 'The diminished 7th chord built on C#: C#, E, G, Bb.'
tags:
- chord
- diminished-7th
- c
parent: /chords
prerequisites:
- /triads/c-sharp-diminished-triad
theory:
  category: chord
  quality: diminished-7th
  root:
    pc: 1
    name: C#
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
  - 1
  - 4
  - 7
  - 10
  spelling:
  - C#
  - E
  - G
  - Bb
  enharmonic_policy: prefer_flats
  key_context: C#
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
  - 1
  - 4
  - 7
  - 10
  emphasize_intervals:
  - '1'
  - b3
  - b5
  - bb7
  preferred_positions:
  - id: c-sharp-diminished-7th-voicing-1
    start_fret: 2
    end_fret: 4
    notes:
    - string: 5
      fret: 4
      pc: 1
      note: C#
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: b3
      role: third
      finger: 1
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: bb7
      role: seventh
      finger: 2
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: '1'
      role: root
      finger: 1
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: b5
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: c-sharp-diminished-7th-voicing-2
    start_fret: 3
    end_fret: 5
    notes:
    - string: 5
      fret: 4
      pc: 1
      note: C#
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: b5
      role: fifth
      finger: 3
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: bb7
      role: seventh
      finger: 1
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: b3
      role: third
      finger: 3
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: b5
      role: fifth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: c-sharp-diminished-7th-voicing-3
    start_fret: 4
    end_fret: 6
    notes:
    - string: 5
      fret: 4
      pc: 1
      note: C#
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: b5
      role: fifth
      finger: 2
    - string: 3
      fret: 6
      pc: 1
      note: C#
      interval: '1'
      role: root
      finger: 3
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: b3
      role: third
      finger: 2
    - string: 1
      fret: 6
      pc: 10
      note: Bb
      interval: bb7
      role: seventh
      finger: 3
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (4.5 2.4 3.3 2.2 3.1).2
  midi_root: 61
musical_forms:
- form: pitch_set
  id: c-sharp-diminished-7th-chord-pitch-set
  pitch_classes:
  - 1
  - 4
  - 7
  - 10
  intervals:
  - '1'
  - b3
  - b5
  - bb7
  spelling:
  - C#
  - E
  - G
  - Bb
- form: voicing
  id: c-sharp-diminished-7th-voicing-1
  position_ref: c-sharp-diminished-7th-voicing-1
  start_fret: 2
  tab: x-4-2-3-2-3
- form: voicing
  id: c-sharp-diminished-7th-voicing-2
  position_ref: c-sharp-diminished-7th-voicing-2
  start_fret: 3
  tab: x-4-5-3-5-3
- form: voicing
  id: c-sharp-diminished-7th-voicing-3
  position_ref: c-sharp-diminished-7th-voicing-3
  start_fret: 4
  tab: x-4-5-6-5-6
- form: tab_excerpt
  id: c-sharp-diminished-7th-chord-tab
  alphatex: \tempo 90 (4.5 2.4 3.3 2.2 3.1).2
---

# C# Diminished 7th Chord

The C# diminished 7th chord is four stacked minor thirds — symmetrical and maximally tense, able to resolve many directions. Its tones are **C# – E – G – Bb**: C# (root), E (minor third), G (diminished fifth), Bb (diminished seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | C# | root |
| b3 | E | minor third |
| b5 | G | diminished fifth |
| bb7 | Bb | diminished seventh |

**Playing it:** the primary shape is `x-4-2-3-2-3` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [C# Diminished Triad](/triads/c-sharp-diminished-triad.md)
