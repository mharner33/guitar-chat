---
type: Concept
slug: /chords/e-suspended-2-chord
title: E Suspended 2 Chord
description: 'The suspended 2 chord built on E: E, F#, B.'
tags:
- chord
- suspended-2
- e
parent: /chords
prerequisites:
- /triads/e-major-triad
theory:
  category: chord
  quality: suspended-2
  root:
    pc: 4
    name: E
  intervals:
  - '1'
  - '2'
  - '5'
  semitones:
  - 0
  - 2
  - 7
  pitch_classes:
  - 4
  - 6
  - 11
  spelling:
  - E
  - F#
  - B
  enharmonic_policy: prefer_sharps
  key_context: E
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
  - 4
  - 6
  - 11
  emphasize_intervals:
  - '1'
  - '2'
  - '5'
  preferred_positions:
  - id: e-suspended-2-voicing-1
    start_fret: 2
    end_fret: 4
    notes:
    - string: 6
      fret: 0
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 0
    - string: 5
      fret: 2
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 1
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 2
    - string: 2
      fret: 0
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 0
    - string: 1
      fret: 2
      pc: 6
      note: F#
      interval: '2'
      role: second
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: e-suspended-2-voicing-2
    start_fret: 4
    end_fret: 7
    notes:
    - string: 5
      fret: 7
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 4
      pc: 6
      note: F#
      interval: '2'
      role: second
      finger: 1
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 1
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 2
    - string: 1
      fret: 7
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 3
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: e-suspended-2-voicing-3
    start_fret: 7
    end_fret: 9
    notes:
    - string: 5
      fret: 7
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 9
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 2
    - string: 3
      fret: 9
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 2
    - string: 2
      fret: 7
      pc: 6
      note: F#
      interval: '2'
      role: second
      finger: 1
    - string: 1
      fret: 7
      pc: 11
      note: B
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
  alphatex: \tempo 90 (0.6 2.5 2.4 4.3 0.2 2.1).2
  midi_root: 64
musical_forms:
- form: pitch_set
  id: e-suspended-2-chord-pitch-set
  pitch_classes:
  - 4
  - 6
  - 11
  intervals:
  - '1'
  - '2'
  - '5'
  spelling:
  - E
  - F#
  - B
- form: voicing
  id: e-suspended-2-voicing-1
  position_ref: e-suspended-2-voicing-1
  start_fret: 2
  tab: 0-2-2-4-0-2
- form: voicing
  id: e-suspended-2-voicing-2
  position_ref: e-suspended-2-voicing-2
  start_fret: 4
  tab: x-7-4-4-5-7
- form: voicing
  id: e-suspended-2-voicing-3
  position_ref: e-suspended-2-voicing-3
  start_fret: 7
  tab: x-7-9-9-7-7
- form: tab_excerpt
  id: e-suspended-2-chord-tab
  alphatex: \tempo 90 (0.6 2.5 2.4 4.3 0.2 2.1).2
---

# E Suspended 2 Chord

The E suspended 2 chord is a triad with the third replaced by the second — open and unresolved. Its tones are **E – F# – B**: E (root), F# (major second), B (perfect fifth).

| Interval | Note | Name |
|---|---|---|
| 1 | E | root |
| 2 | F# | major second |
| 5 | B | perfect fifth |

**Playing it:** the primary shape is `0-2-2-4-0-2` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [E Major Triad](/triads/e-major-triad.md)
