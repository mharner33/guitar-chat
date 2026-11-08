---
type: Concept
slug: /triads/c-sharp-minor-triad
title: C# Minor Triad
description: 'The minor triad built on C#: C#, E, G#.'
tags:
- triad
- minor
- c
parent: /triads
prerequisites:
- /scales/c-sharp-natural-minor-scale
used_by:
- /chords/c-sharp-minor-7th-chord
- /chords/d-flat-minor-major-7th-chord
- /chords/c-sharp-minor-6th-chord
theory:
  category: triad
  quality: minor
  root:
    pc: 1
    name: C#
  intervals:
  - '1'
  - b3
  - '5'
  semitones:
  - 0
  - 3
  - 7
  pitch_classes:
  - 1
  - 4
  - 8
  spelling:
  - C#
  - E
  - G#
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
  - 8
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  preferred_positions:
  - id: c-sharp-minor-voicing-1
    start_fret: 1
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
      finger: 2
    - string: 3
      fret: 1
      pc: 8
      note: G#
      interval: '5'
      role: fifth
      finger: 1
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: '1'
      role: root
      finger: 2
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
  - id: c-sharp-minor-voicing-2
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
      fret: 6
      pc: 8
      note: G#
      interval: '5'
      role: fifth
      finger: 3
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
      fret: 4
      pc: 8
      note: G#
      interval: '5'
      role: fifth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: c-sharp-minor-voicing-3
    start_fret: 2
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
      fret: 2
      pc: 4
      note: E
      interval: b3
      role: third
      finger: 1
    - string: 3
      fret: 6
      pc: 1
      note: C#
      interval: '1'
      role: root
      finger: 3
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: '1'
      role: root
      finger: 1
    - string: 1
      fret: 4
      pc: 8
      note: G#
      interval: '5'
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (4.5 2.4 1.3 2.2 0.1).2
  midi_root: 61
musical_forms:
- form: pitch_set
  id: c-sharp-minor-triad-pitch-set
  pitch_classes:
  - 1
  - 4
  - 8
  intervals:
  - '1'
  - b3
  - '5'
  spelling:
  - C#
  - E
  - G#
- form: voicing
  id: c-sharp-minor-voicing-1
  position_ref: c-sharp-minor-voicing-1
  start_fret: 1
  tab: x-4-2-1-2-0
- form: voicing
  id: c-sharp-minor-voicing-2
  position_ref: c-sharp-minor-voicing-2
  start_fret: 4
  tab: x-4-6-6-5-4
- form: voicing
  id: c-sharp-minor-voicing-3
  position_ref: c-sharp-minor-voicing-3
  start_fret: 2
  tab: x-4-2-6-2-4
- form: tab_excerpt
  id: c-sharp-minor-triad-tab
  alphatex: \tempo 90 (4.5 2.4 1.3 2.2 0.1).2
---

# C# Minor Triad

The C# minor triad is a darker but still consonant chord. Its tones are **C# – E – G#**.

**Construction:** stacked thirds — a minor third (C#–E) under a major third (E–G#).

| Interval | Note | Name |
|---|---|---|
| 1 | C# | root |
| b3 | E | minor third |
| 5 | G# | perfect fifth |

**Playing it:** the primary shape is `x-4-2-1-2-0` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [C# Natural Minor Scale](/scales/c-sharp-natural-minor-scale.md)
- **Used by:** [C# Minor 7th Chord](/chords/c-sharp-minor-7th-chord.md), [Db Minor Major 7th Chord](/chords/d-flat-minor-major-7th-chord.md), [C# Minor 6th Chord](/chords/c-sharp-minor-6th-chord.md)
