---
type: Concept
slug: /chords/a-minor-7th-chord
title: A Minor 7th Chord
description: 'The minor 7th chord built on A: A, C, E, G.'
tags:
- chord
- minor-7th
- a
parent: /chords
prerequisites:
- /triads/a-minor-triad
used_by:
- /chords/a-minor-9th-chord
theory:
  category: chord
  quality: minor-7th
  root:
    pc: 9
    name: A
  intervals:
  - '1'
  - b3
  - '5'
  - b7
  semitones:
  - 0
  - 3
  - 7
  - 10
  pitch_classes:
  - 9
  - 0
  - 4
  - 7
  spelling:
  - A
  - C
  - E
  - G
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
  - 7
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  - b7
  preferred_positions:
  - id: a-minor-7th-voicing-1
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
      fret: 0
      pc: 7
      note: G
      interval: b7
      role: seventh
      finger: 0
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
  - id: a-minor-7th-voicing-2
    start_fret: 3
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
      fret: 3
      pc: 0
      note: C
      interval: b3
      role: third
      finger: 1
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: b7
      role: seventh
      finger: 2
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
      fret: 3
      pc: 7
      note: G
      interval: b7
      role: seventh
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: a-minor-7th-voicing-3
    start_fret: 2
    end_fret: 5
    notes:
    - string: 6
      fret: 5
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 3
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
      finger: 3
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '5'
      role: fifth
      finger: 3
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: b7
      role: seventh
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
playback:
  alphatex: \tempo 90 (0.5 2.4 0.3 1.2 0.1).2
  midi_root: 69
musical_forms:
- form: pitch_set
  id: a-minor-7th-chord-pitch-set
  pitch_classes:
  - 9
  - 0
  - 4
  - 7
  intervals:
  - '1'
  - b3
  - '5'
  - b7
  spelling:
  - A
  - C
  - E
  - G
- form: voicing
  id: a-minor-7th-voicing-1
  position_ref: a-minor-7th-voicing-1
  start_fret: 1
  tab: x-0-2-0-1-0
- form: voicing
  id: a-minor-7th-voicing-2
  position_ref: a-minor-7th-voicing-2
  start_fret: 3
  tab: 5-3-5-5-5-3
- form: voicing
  id: a-minor-7th-voicing-3
  position_ref: a-minor-7th-voicing-3
  start_fret: 2
  tab: 5-0-2-5-5-3
- form: tab_excerpt
  id: a-minor-7th-chord-tab
  alphatex: \tempo 90 (0.5 2.4 0.3 1.2 0.1).2
---

# A Minor 7th Chord

The A minor 7th chord is a minor triad plus a minor seventh — smooth and mellow, the workhorse of jazz and soul. Its tones are **A – C – E – G**: A (root), C (minor third), E (perfect fifth), G (minor seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | A | root |
| b3 | C | minor third |
| 5 | E | perfect fifth |
| b7 | G | minor seventh |

**Playing it:** the primary shape is `x-0-2-0-1-0` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [A Minor Triad](/triads/a-minor-triad.md)
- **Used by:** [A Minor 9th Chord](/chords/a-minor-9th-chord.md)
