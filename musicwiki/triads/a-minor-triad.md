---
type: Concept
slug: /triads/a-minor-triad
title: A Minor Triad
description: 'The minor triad built on A: A, C, E.'
tags:
- triad
- minor
- a
parent: /triads
prerequisites:
- /scales/a-natural-minor-scale
used_by:
- /chords/a-minor-7th-chord
- /chords/a-minor-major-7th-chord
- /chords/a-minor-6th-chord
theory:
  category: triad
  quality: minor
  root:
    pc: 9
    name: A
  intervals:
  - '1'
  - b3
  - '5'
  semitones:
  - 0
  - 3
  - 7
  pitch_classes:
  - 9
  - 0
  - 4
  spelling:
  - A
  - C
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
  - 0
  - 4
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  preferred_positions:
  - id: a-minor-voicing-1
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
      fret: 2
      pc: 4
      note: E
      interval: '5'
      role: fifth
      finger: 2
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
  - id: a-minor-voicing-2
    start_fret: 2
    end_fret: 5
    notes:
    - string: 6
      fret: 5
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 2
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
      fret: 5
      pc: 0
      note: C
      interval: b3
      role: third
      finger: 2
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '5'
      role: fifth
      finger: 2
    - string: 1
      fret: 0
      pc: 4
      note: E
      interval: '5'
      role: fifth
      finger: 0
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: a-minor-voicing-3
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
      finger: 2
    - string: 4
      fret: 7
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 2
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: b3
      role: third
      finger: 1
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
playback:
  alphatex: \tempo 90 (0.5 2.4 2.3 1.2 0.1).2
  midi_root: 69
musical_forms:
- form: pitch_set
  id: a-minor-triad-pitch-set
  pitch_classes:
  - 9
  - 0
  - 4
  intervals:
  - '1'
  - b3
  - '5'
  spelling:
  - A
  - C
  - E
- form: voicing
  id: a-minor-voicing-1
  position_ref: a-minor-voicing-1
  start_fret: 1
  tab: x-0-2-2-1-0
- form: voicing
  id: a-minor-voicing-2
  position_ref: a-minor-voicing-2
  start_fret: 2
  tab: 5-0-2-5-5-0
- form: voicing
  id: a-minor-voicing-3
  position_ref: a-minor-voicing-3
  start_fret: 5
  tab: 5-7-7-5-5-5
- form: tab_excerpt
  id: a-minor-triad-tab
  alphatex: \tempo 90 (0.5 2.4 2.3 1.2 0.1).2
---

# A Minor Triad

The A minor triad is a darker but still consonant chord. Its tones are **A – C – E**.

**Construction:** stacked thirds — a minor third (A–C) under a major third (C–E).

| Interval | Note | Name |
|---|---|---|
| 1 | A | root |
| b3 | C | minor third |
| 5 | E | perfect fifth |

**Playing it:** the primary shape is `x-0-2-2-1-0` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [A Natural Minor Scale](/scales/a-natural-minor-scale.md)
- **Used by:** [A Minor 7th Chord](/chords/a-minor-7th-chord.md), [A Minor Major 7th Chord](/chords/a-minor-major-7th-chord.md), [A Minor 6th Chord](/chords/a-minor-6th-chord.md)
