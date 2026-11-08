---
type: Concept
slug: /triads/a-sharp-diminished-triad
title: A# Diminished Triad
description: 'The diminished triad built on A#: A#, C#, E.'
tags:
- triad
- diminished
- a
parent: /triads
used_by:
- /chords/a-sharp-minor-7th-flat-5-chord
- /chords/a-sharp-diminished-7th-chord
theory:
  category: triad
  quality: diminished
  root:
    pc: 10
    name: A#
  intervals:
  - '1'
  - b3
  - b5
  semitones:
  - 0
  - 3
  - 6
  pitch_classes:
  - 10
  - 1
  - 4
  spelling:
  - A#
  - C#
  - E
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
  emphasize_intervals:
  - '1'
  - b3
  - b5
  preferred_positions:
  - id: a-sharp-diminished-voicing-1
    start_fret: 1
    end_fret: 3
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
      fret: 3
      pc: 10
      note: A#
      interval: '1'
      role: root
      finger: 3
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
  - id: a-sharp-diminished-voicing-2
    start_fret: 2
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
      finger: 3
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: b5
      role: fifth
      finger: 1
    - string: 3
      fret: 3
      pc: 10
      note: A#
      interval: '1'
      role: root
      finger: 2
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: b3
      role: third
      finger: 1
    - string: 1
      fret: 0
      pc: 4
      note: E
      interval: b5
      role: fifth
      finger: 0
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: a-sharp-diminished-voicing-3
    start_fret: 5
    end_fret: 8
    notes:
    - string: 6
      fret: 6
      pc: 10
      note: A#
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 7
      pc: 4
      note: E
      interval: b5
      role: fifth
      finger: 3
    - string: 4
      fret: 8
      pc: 10
      note: A#
      interval: '1'
      role: root
      finger: 4
    - string: 3
      fret: 6
      pc: 1
      note: C#
      interval: b3
      role: third
      finger: 2
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: b5
      role: fifth
      finger: 1
    - string: 1
      fret: 6
      pc: 10
      note: A#
      interval: '1'
      role: root
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (1.5 2.4 3.3 2.2 0.1).2
  midi_root: 70
musical_forms:
- form: pitch_set
  id: a-sharp-diminished-triad-pitch-set
  pitch_classes:
  - 10
  - 1
  - 4
  intervals:
  - '1'
  - b3
  - b5
  spelling:
  - A#
  - C#
  - E
- form: voicing
  id: a-sharp-diminished-voicing-1
  position_ref: a-sharp-diminished-voicing-1
  start_fret: 1
  tab: x-1-2-3-2-0
- form: voicing
  id: a-sharp-diminished-voicing-2
  position_ref: a-sharp-diminished-voicing-2
  start_fret: 2
  tab: 6-4-2-3-2-0
- form: voicing
  id: a-sharp-diminished-voicing-3
  position_ref: a-sharp-diminished-voicing-3
  start_fret: 5
  tab: 6-7-8-6-5-6
- form: tab_excerpt
  id: a-sharp-diminished-triad-tab
  alphatex: \tempo 90 (1.5 2.4 3.3 2.2 0.1).2
---

# A# Diminished Triad

The A# diminished triad is a tense, unstable chord of two stacked minor thirds that wants to resolve. Its tones are **A# – C# – E**.

**Construction:** stacked thirds — a minor third (A#–C#) under a minor third (C#–E).

| Interval | Note | Name |
|---|---|---|
| 1 | A# | root |
| b3 | C# | minor third |
| b5 | E | diminished fifth |

**Playing it:** the primary shape is `x-1-2-3-2-0` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Used by:** [A# Minor 7th Flat 5 Chord](/chords/a-sharp-minor-7th-flat-5-chord.md), [A# Diminished 7th Chord](/chords/a-sharp-diminished-7th-chord.md)
