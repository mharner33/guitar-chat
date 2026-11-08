---
type: Concept
slug: /chords/d-dominant-7th-chord
title: D Dominant 7th Chord
description: 'The dominant 7th chord built on D: D, F#, A, C.'
tags:
- chord
- dominant-7th
- d
parent: /chords
prerequisites:
- /triads/d-major-triad
used_by:
- /chords/d-9th-chord
- /chords/d-11th-chord
- /chords/d-13th-chord
theory:
  category: chord
  quality: dominant-7th
  root:
    pc: 2
    name: D
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  semitones:
  - 0
  - 4
  - 7
  - 10
  pitch_classes:
  - 2
  - 6
  - 9
  - 0
  spelling:
  - D
  - F#
  - A
  - C
  enharmonic_policy: prefer_sharps
  key_context: D
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
  - 2
  - 6
  - 9
  - 0
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - b7
  preferred_positions:
  - id: d-dominant-7th-voicing-1
    start_fret: 1
    end_fret: 2
    notes:
    - string: 4
      fret: 0
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 0
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: '5'
      role: fifth
      finger: 2
    - string: 2
      fret: 1
      pc: 0
      note: C
      interval: b7
      role: seventh
      finger: 1
    - string: 1
      fret: 2
      pc: 6
      note: F#
      interval: '3'
      role: third
      finger: 2
    mute_strings:
    - 5
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: d-dominant-7th-voicing-2
    start_fret: 3
    end_fret: 5
    notes:
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 4
      pc: 6
      note: F#
      interval: '3'
      role: third
      finger: 2
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: b7
      role: seventh
      finger: 3
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 1
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '5'
      role: fifth
      finger: 3
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: d-dominant-7th-voicing-3
    start_fret: 5
    end_fret: 7
    notes:
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 7
      pc: 9
      note: A
      interval: '5'
      role: fifth
      finger: 2
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: b7
      role: seventh
      finger: 1
    - string: 2
      fret: 7
      pc: 6
      note: F#
      interval: '3'
      role: third
      finger: 2
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '5'
      role: fifth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (0.4 2.3 1.2 2.1).2
  midi_root: 62
musical_forms:
- form: pitch_set
  id: d-dominant-7th-chord-pitch-set
  pitch_classes:
  - 2
  - 6
  - 9
  - 0
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  spelling:
  - D
  - F#
  - A
  - C
- form: voicing
  id: d-dominant-7th-voicing-1
  position_ref: d-dominant-7th-voicing-1
  start_fret: 1
  tab: x-x-0-2-1-2
- form: voicing
  id: d-dominant-7th-voicing-2
  position_ref: d-dominant-7th-voicing-2
  start_fret: 3
  tab: x-5-4-5-3-5
- form: voicing
  id: d-dominant-7th-voicing-3
  position_ref: d-dominant-7th-voicing-3
  start_fret: 5
  tab: x-5-7-5-7-5
- form: tab_excerpt
  id: d-dominant-7th-chord-tab
  alphatex: \tempo 90 (0.4 2.3 1.2 2.1).2
---

# D Dominant 7th Chord

The D dominant 7th chord is a major triad plus a minor seventh — the restless engine of tonal harmony, pulling toward the tonic a fourth above. Its tones are **D – F# – A – C**: D (root), F# (major third), A (perfect fifth), C (minor seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | D | root |
| 3 | F# | major third |
| 5 | A | perfect fifth |
| b7 | C | minor seventh |

**Playing it:** the primary shape is `x-x-0-2-1-2` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [D Major Triad](/triads/d-major-triad.md)
- **Used by:** [D 9th Chord](/chords/d-9th-chord.md), [D 11th Chord](/chords/d-11th-chord.md), [D 13th Chord](/chords/d-13th-chord.md)
