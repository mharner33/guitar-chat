---
type: Concept
slug: /triads/c-sharp-diminished-triad
title: C# Diminished Triad
description: 'The diminished triad built on C#: C#, E, G.'
tags:
- triad
- diminished
- c
parent: /triads
used_by:
- /chords/c-sharp-minor-7th-flat-5-chord
- /chords/c-sharp-diminished-7th-chord
theory:
  category: triad
  quality: diminished
  root:
    pc: 1
    name: C#
  intervals:
  - '1'
  - b3
  - b5
  semitones:
  - 0
  - 3
  - 6
  pitch_classes:
  - 1
  - 4
  - 7
  spelling:
  - C#
  - E
  - G
  enharmonic_policy: prefer_sharps
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
  emphasize_intervals:
  - '1'
  - b3
  - b5
  preferred_positions:
  - id: c-sharp-diminished-voicing-1
    start_fret: 2
    end_fret: 4
    notes:
    - string: 5
      fret: 4
      pc: 1
      note: C#
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: b3
      role: third
      finger: 1
    - string: 3
      fret: 0
      pc: 7
      note: G
      interval: b5
      role: fifth
      finger: 0
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: '1'
      role: root
      finger: 1
    - string: 1
      fret: 0
      pc: 4
      note: E
      interval: b3
      role: third
      finger: 0
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: c-sharp-diminished-voicing-2
    start_fret: 3
    end_fret: 6
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
      fret: 6
      pc: 1
      note: C#
      interval: '1'
      role: root
      finger: 4
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
  - id: c-sharp-diminished-voicing-3
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
    mute_strings:
    - 1
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (4.5 2.4 0.3 2.2 0.1).2
  midi_root: 61
musical_forms:
- form: pitch_set
  id: c-sharp-diminished-triad-pitch-set
  pitch_classes:
  - 1
  - 4
  - 7
  intervals:
  - '1'
  - b3
  - b5
  spelling:
  - C#
  - E
  - G
- form: voicing
  id: c-sharp-diminished-voicing-1
  position_ref: c-sharp-diminished-voicing-1
  start_fret: 2
  tab: x-4-2-0-2-0
- form: voicing
  id: c-sharp-diminished-voicing-2
  position_ref: c-sharp-diminished-voicing-2
  start_fret: 3
  tab: x-4-5-6-5-3
- form: voicing
  id: c-sharp-diminished-voicing-3
  position_ref: c-sharp-diminished-voicing-3
  start_fret: 4
  tab: x-4-5-6-5-x
- form: tab_excerpt
  id: c-sharp-diminished-triad-tab
  alphatex: \tempo 90 (4.5 2.4 0.3 2.2 0.1).2
---

# C# Diminished Triad

The C# diminished triad is a tense, unstable chord of two stacked minor thirds that wants to resolve. Its tones are **C# – E – G**.

**Construction:** stacked thirds — a minor third (C#–E) under a minor third (E–G).

| Interval | Note | Name |
|---|---|---|
| 1 | C# | root |
| b3 | E | minor third |
| b5 | G | diminished fifth |

**Playing it:** the primary shape is `x-4-2-0-2-0` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Used by:** [C# Minor 7th Flat 5 Chord](/chords/c-sharp-minor-7th-flat-5-chord.md), [C# Diminished 7th Chord](/chords/c-sharp-diminished-7th-chord.md)
