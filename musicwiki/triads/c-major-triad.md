---
type: Concept
slug: /triads/c-major-triad
title: C Major Triad
description: 'The major triad built on C: C, E, G.'
tags:
- triad
- major
- c
parent: /triads
prerequisites:
- /scales/c-major-scale
used_by:
- /chords/c-dominant-7th-chord
- /chords/c-major-7th-chord
- /chords/c-6th-chord
- /chords/c-suspended-2-chord
- /chords/c-suspended-4-chord
theory:
  category: triad
  quality: major
  root:
    pc: 0
    name: C
  intervals:
  - '1'
  - '3'
  - '5'
  semitones:
  - 0
  - 4
  - 7
  pitch_classes:
  - 0
  - 4
  - 7
  spelling:
  - C
  - E
  - G
  enharmonic_policy: prefer_sharps
  key_context: C
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
  - 0
  - 4
  - 7
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  preferred_positions:
  - id: c-major-voicing-1
    start_fret: 1
    end_fret: 3
    notes:
    - string: 5
      fret: 3
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: '3'
      role: third
      finger: 2
    - string: 3
      fret: 0
      pc: 7
      note: G
      interval: '5'
      role: fifth
      finger: 0
    - string: 2
      fret: 1
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 1
    - string: 1
      fret: 0
      pc: 4
      note: E
      interval: '3'
      role: third
      finger: 0
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
  - id: c-major-voicing-2
    start_fret: 3
    end_fret: 5
    notes:
    - string: 5
      fret: 3
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: '5'
      role: fifth
      finger: 2
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 2
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '3'
      role: third
      finger: 2
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: '5'
      role: fifth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: c-major-voicing-3
    start_fret: 2
    end_fret: 5
    notes:
    - string: 5
      fret: 3
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: '3'
      role: third
      finger: 1
    - string: 3
      fret: 0
      pc: 7
      note: G
      interval: '5'
      role: fifth
      finger: 0
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '3'
      role: third
      finger: 3
    - string: 1
      fret: 0
      pc: 4
      note: E
      interval: '3'
      role: third
      finger: 0
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
playback:
  alphatex: \tempo 90 (3.5 2.4 0.3 1.2 0.1).2
  midi_root: 60
musical_forms:
- form: pitch_set
  id: c-major-triad-pitch-set
  pitch_classes:
  - 0
  - 4
  - 7
  intervals:
  - '1'
  - '3'
  - '5'
  spelling:
  - C
  - E
  - G
- form: voicing
  id: c-major-voicing-1
  position_ref: c-major-voicing-1
  start_fret: 1
  tab: x-3-2-0-1-0
- form: voicing
  id: c-major-voicing-2
  position_ref: c-major-voicing-2
  start_fret: 3
  tab: x-3-5-5-5-3
- form: voicing
  id: c-major-voicing-3
  position_ref: c-major-voicing-3
  start_fret: 2
  tab: x-3-2-0-5-0
- form: tab_excerpt
  id: c-major-triad-tab
  alphatex: \tempo 90 (3.5 2.4 0.3 1.2 0.1).2
---

# C Major Triad

The C major triad is a bright, stable, consonant chord. Its tones are **C – E – G**.

**Construction:** stacked thirds — a major third (C–E) under a minor third (E–G).

| Interval | Note | Name |
|---|---|---|
| 1 | C | root |
| 3 | E | major third |
| 5 | G | perfect fifth |

**Playing it:** the primary shape is `x-3-2-0-1-0` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [C Major Scale](/scales/c-major-scale.md)
- **Used by:** [C Dominant 7th Chord](/chords/c-dominant-7th-chord.md), [C Major 7th Chord](/chords/c-major-7th-chord.md), [C 6th Chord](/chords/c-6th-chord.md), [C Suspended 2 Chord](/chords/c-suspended-2-chord.md), [C Suspended 4 Chord](/chords/c-suspended-4-chord.md)
