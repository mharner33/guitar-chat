---
type: Concept
slug: /chords/a-suspended-4-chord
title: A Suspended 4 Chord
description: 'The suspended 4 chord built on A: A, D, E.'
tags:
- chord
- suspended-4
- a
parent: /chords
prerequisites:
- /triads/a-major-triad
theory:
  category: chord
  quality: suspended-4
  root:
    pc: 9
    name: A
  intervals:
  - '1'
  - '4'
  - '5'
  semitones:
  - 0
  - 5
  - 7
  pitch_classes:
  - 9
  - 2
  - 4
  spelling:
  - A
  - D
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
  - 2
  - 4
  emphasize_intervals:
  - '1'
  - '4'
  - '5'
  preferred_positions:
  - id: a-suspended-4-voicing-1
    start_fret: 2
    end_fret: 3
    notes:
    - string: 5
      fret: 0
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 0
    - string: 4
      fret: 0
      pc: 2
      note: D
      interval: '4'
      role: fourth
      finger: 0
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 1
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '4'
      role: fourth
      finger: 2
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
  - id: a-suspended-4-voicing-2
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
      fret: 5
      pc: 2
      note: D
      interval: '4'
      role: fourth
      finger: 1
    - string: 4
      fret: 7
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 2
    - string: 3
      fret: 7
      pc: 2
      note: D
      interval: '4'
      role: fourth
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
  - id: a-suspended-4-voicing-3
    start_fret: 3
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
      fret: 7
      pc: 2
      note: D
      interval: '4'
      role: fourth
      finger: 3
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '4'
      role: fourth
      finger: 1
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
  alphatex: \tempo 90 (0.5 0.4 2.3 3.2 0.1).2
  midi_root: 69
musical_forms:
- form: pitch_set
  id: a-suspended-4-chord-pitch-set
  pitch_classes:
  - 9
  - 2
  - 4
  intervals:
  - '1'
  - '4'
  - '5'
  spelling:
  - A
  - D
  - E
- form: voicing
  id: a-suspended-4-voicing-1
  position_ref: a-suspended-4-voicing-1
  start_fret: 2
  tab: x-0-0-2-3-0
- form: voicing
  id: a-suspended-4-voicing-2
  position_ref: a-suspended-4-voicing-2
  start_fret: 5
  tab: 5-5-7-7-5-5
- form: voicing
  id: a-suspended-4-voicing-3
  position_ref: a-suspended-4-voicing-3
  start_fret: 3
  tab: 5-7-7-7-3-5
- form: tab_excerpt
  id: a-suspended-4-chord-tab
  alphatex: \tempo 90 (0.5 0.4 2.3 3.2 0.1).2
---

# A Suspended 4 Chord

The A suspended 4 chord is a triad with the third replaced by the fourth — tense, wanting to fall to the third. Its tones are **A – D – E**: A (root), D (perfect fourth), E (perfect fifth).

| Interval | Note | Name |
|---|---|---|
| 1 | A | root |
| 4 | D | perfect fourth |
| 5 | E | perfect fifth |

**Playing it:** the primary shape is `x-0-0-2-3-0` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [A Major Triad](/triads/a-major-triad.md)
