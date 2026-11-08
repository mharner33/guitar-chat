---
type: Concept
slug: /chords/b-flat-major-7th-chord
title: Bb Major 7th Chord
description: 'The major 7th chord built on Bb: Bb, D, F, A.'
tags:
- chord
- major-7th
- b
parent: /chords
prerequisites:
- /triads/b-flat-major-triad
used_by:
- /chords/b-flat-major-9th-chord
theory:
  category: chord
  quality: major-7th
  root:
    pc: 10
    name: Bb
  intervals:
  - '1'
  - '3'
  - '5'
  - '7'
  semitones:
  - 0
  - 4
  - 7
  - 11
  pitch_classes:
  - 10
  - 2
  - 5
  - 9
  spelling:
  - Bb
  - D
  - F
  - A
  enharmonic_policy: prefer_flats
  key_context: Bb
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
  - 2
  - 5
  - 9
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - '7'
  preferred_positions:
  - id: b-flat-major-7th-voicing-1
    start_fret: 1
    end_fret: 3
    notes:
    - string: 5
      fret: 1
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 0
      pc: 2
      note: D
      interval: '3'
      role: third
      finger: 0
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: '7'
      role: seventh
      finger: 2
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '3'
      role: third
      finger: 3
    - string: 1
      fret: 1
      pc: 5
      note: F
      interval: '5'
      role: fifth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: b-flat-major-7th-voicing-2
    start_fret: 3
    end_fret: 6
    notes:
    - string: 6
      fret: 6
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: '3'
      role: third
      finger: 2
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '5'
      role: fifth
      finger: 1
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 1
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '3'
      role: third
      finger: 1
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '7'
      role: seventh
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: b-flat-major-7th-voicing-3
    start_fret: 5
    end_fret: 7
    notes:
    - string: 6
      fret: 6
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: '3'
      role: third
      finger: 1
    - string: 4
      fret: 7
      pc: 9
      note: A
      interval: '7'
      role: seventh
      finger: 3
    - string: 3
      fret: 7
      pc: 2
      note: D
      interval: '3'
      role: third
      finger: 3
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '5'
      role: fifth
      finger: 2
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '7'
      role: seventh
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (1.5 0.4 2.3 3.2 1.1).2
  midi_root: 70
musical_forms:
- form: pitch_set
  id: b-flat-major-7th-chord-pitch-set
  pitch_classes:
  - 10
  - 2
  - 5
  - 9
  intervals:
  - '1'
  - '3'
  - '5'
  - '7'
  spelling:
  - Bb
  - D
  - F
  - A
- form: voicing
  id: b-flat-major-7th-voicing-1
  position_ref: b-flat-major-7th-voicing-1
  start_fret: 1
  tab: x-1-0-2-3-1
- form: voicing
  id: b-flat-major-7th-voicing-2
  position_ref: b-flat-major-7th-voicing-2
  start_fret: 3
  tab: 6-5-3-3-3-5
- form: voicing
  id: b-flat-major-7th-voicing-3
  position_ref: b-flat-major-7th-voicing-3
  start_fret: 5
  tab: 6-5-7-7-6-5
- form: tab_excerpt
  id: b-flat-major-7th-chord-tab
  alphatex: \tempo 90 (1.5 0.4 2.3 3.2 1.1).2
---

# Bb Major 7th Chord

The Bb major 7th chord is a major triad plus a major seventh — lush, warm, and at rest. Its tones are **Bb – D – F – A**: Bb (root), D (major third), F (perfect fifth), A (major seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | Bb | root |
| 3 | D | major third |
| 5 | F | perfect fifth |
| 7 | A | major seventh |

**Playing it:** the primary shape is `x-1-0-2-3-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Bb Major Triad](/triads/b-flat-major-triad.md)
- **Used by:** [Bb Major 9th Chord](/chords/b-flat-major-9th-chord.md)
