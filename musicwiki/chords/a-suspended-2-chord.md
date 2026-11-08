---
type: Concept
slug: /chords/a-suspended-2-chord
title: A Suspended 2 Chord
description: 'The suspended 2 chord built on A: A, B, E.'
tags:
- chord
- suspended-2
- a
parent: /chords
prerequisites:
- /triads/a-major-triad
theory:
  category: chord
  quality: suspended-2
  root:
    pc: 9
    name: A
  intervals:
  - '1'
  - '2'
  - '5'
  semitones:
  - 0
  - 2
  - 7
  pitch_classes:
  - 9
  - 11
  - 4
  spelling:
  - A
  - B
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
  - 11
  - 4
  emphasize_intervals:
  - '1'
  - '2'
  - '5'
  preferred_positions:
  - id: a-suspended-2-voicing-1
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
      fret: 0
      pc: 11
      note: B
      interval: '2'
      role: second
      finger: 0
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
  - id: a-suspended-2-voicing-2
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
      fret: 4
      pc: 11
      note: B
      interval: '2'
      role: second
      finger: 1
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
  - id: a-suspended-2-voicing-3
    start_fret: 5
    end_fret: 9
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
      fret: 9
      pc: 4
      note: E
      interval: '5'
      role: fifth
      finger: 3
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '5'
      role: fifth
      finger: 1
    - string: 1
      fret: 7
      pc: 11
      note: B
      interval: '2'
      role: second
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (0.5 2.4 2.3 0.2 0.1).2
  midi_root: 69
musical_forms:
- form: pitch_set
  id: a-suspended-2-chord-pitch-set
  pitch_classes:
  - 9
  - 11
  - 4
  intervals:
  - '1'
  - '2'
  - '5'
  spelling:
  - A
  - B
  - E
- form: voicing
  id: a-suspended-2-voicing-1
  position_ref: a-suspended-2-voicing-1
  start_fret: 2
  tab: x-0-2-2-0-0
- form: voicing
  id: a-suspended-2-voicing-2
  position_ref: a-suspended-2-voicing-2
  start_fret: 4
  tab: 5-7-7-4-5-5
- form: voicing
  id: a-suspended-2-voicing-3
  position_ref: a-suspended-2-voicing-3
  start_fret: 5
  tab: 5-7-7-9-5-7
- form: tab_excerpt
  id: a-suspended-2-chord-tab
  alphatex: \tempo 90 (0.5 2.4 2.3 0.2 0.1).2
---

# A Suspended 2 Chord

The A suspended 2 chord is a triad with the third replaced by the second — open and unresolved. Its tones are **A – B – E**: A (root), B (major second), E (perfect fifth).

| Interval | Note | Name |
|---|---|---|
| 1 | A | root |
| 2 | B | major second |
| 5 | E | perfect fifth |

**Playing it:** the primary shape is `x-0-2-2-0-0` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [A Major Triad](/triads/a-major-triad.md)
