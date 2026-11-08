---
type: Concept
slug: /triads/a-major-triad
title: A Major Triad
description: 'The major triad built on A: A, C#, E.'
tags:
- triad
- major
- a
parent: /triads
prerequisites:
- /scales/a-major-scale
used_by:
- /chords/a-dominant-7th-chord
- /chords/a-major-7th-chord
- /chords/a-6th-chord
- /chords/a-suspended-2-chord
- /chords/a-suspended-4-chord
theory:
  category: triad
  quality: major
  root:
    pc: 9
    name: A
  intervals:
  - '1'
  - '3'
  - '5'
  semitones:
  - 0
  - 4
  - 7
  pitch_classes:
  - 9
  - 1
  - 4
  spelling:
  - A
  - C#
  - E
  enharmonic_policy: prefer_sharps
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
  - 1
  - 4
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  preferred_positions:
  - id: a-major-voicing-1
    start_fret: 2
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
      fret: 2
      pc: 4
      note: E
      interval: '5'
      role: fifth
      finger: 1
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 1
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: '3'
      role: third
      finger: 1
    - string: 1
      fret: 0
      pc: 4
      note: E
      interval: '5'
      role: fifth
      finger: 0
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: a-major-voicing-2
    start_fret: 5
    end_fret: 7
    notes:
    - string: 6
      fret: 5
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 7
      pc: 4
      note: E
      interval: '5'
      role: fifth
      finger: 3
    - string: 4
      fret: 7
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 3
    - string: 3
      fret: 6
      pc: 1
      note: C#
      interval: '3'
      role: third
      finger: 2
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '5'
      role: fifth
      finger: 1
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: a-major-voicing-3
    start_fret: 4
    end_fret: 7
    notes:
    - string: 6
      fret: 5
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 4
      pc: 1
      note: C#
      interval: '3'
      role: third
      finger: 1
    - string: 4
      fret: 7
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 4
    - string: 3
      fret: 6
      pc: 1
      note: C#
      interval: '3'
      role: third
      finger: 3
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '5'
      role: fifth
      finger: 2
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (0.5 2.4 2.3 2.2 0.1).2
  midi_root: 69
musical_forms:
- form: pitch_set
  id: a-major-triad-pitch-set
  pitch_classes:
  - 9
  - 1
  - 4
  intervals:
  - '1'
  - '3'
  - '5'
  spelling:
  - A
  - C#
  - E
- form: voicing
  id: a-major-voicing-1
  position_ref: a-major-voicing-1
  start_fret: 2
  tab: x-0-2-2-2-0
- form: voicing
  id: a-major-voicing-2
  position_ref: a-major-voicing-2
  start_fret: 5
  tab: 5-7-7-6-5-5
- form: voicing
  id: a-major-voicing-3
  position_ref: a-major-voicing-3
  start_fret: 4
  tab: 5-4-7-6-5-5
- form: tab_excerpt
  id: a-major-triad-tab
  alphatex: \tempo 90 (0.5 2.4 2.3 2.2 0.1).2
---

# A Major Triad

The A major triad is a bright, stable, consonant chord. Its tones are **A – C# – E**.

**Construction:** stacked thirds — a major third (A–C#) under a minor third (C#–E).

| Interval | Note | Name |
|---|---|---|
| 1 | A | root |
| 3 | C# | major third |
| 5 | E | perfect fifth |

**Playing it:** the primary shape is `x-0-2-2-2-0` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [A Major Scale](/scales/a-major-scale.md)
- **Used by:** [A Dominant 7th Chord](/chords/a-dominant-7th-chord.md), [A Major 7th Chord](/chords/a-major-7th-chord.md), [A 6th Chord](/chords/a-6th-chord.md), [A Suspended 2 Chord](/chords/a-suspended-2-chord.md), [A Suspended 4 Chord](/chords/a-suspended-4-chord.md)
