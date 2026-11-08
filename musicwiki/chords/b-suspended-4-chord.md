---
type: Concept
slug: /chords/b-suspended-4-chord
title: B Suspended 4 Chord
description: 'The suspended 4 chord built on B: B, E, F#.'
tags:
- chord
- suspended-4
- b
parent: /chords
prerequisites:
- /triads/b-major-triad
theory:
  category: chord
  quality: suspended-4
  root:
    pc: 11
    name: B
  intervals:
  - '1'
  - '4'
  - '5'
  semitones:
  - 0
  - 5
  - 7
  pitch_classes:
  - 11
  - 4
  - 6
  spelling:
  - B
  - E
  - F#
  enharmonic_policy: prefer_sharps
  key_context: B
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
  - 11
  - 4
  - 6
  emphasize_intervals:
  - '1'
  - '4'
  - '5'
  preferred_positions:
  - id: b-suspended-4-voicing-1
    start_fret: 2
    end_fret: 4
    notes:
    - string: 5
      fret: 2
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: '4'
      role: fourth
      finger: 1
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 2
    - string: 2
      fret: 0
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 0
    - string: 1
      fret: 2
      pc: 6
      note: F#
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
  - id: b-suspended-4-voicing-2
    start_fret: 4
    end_fret: 7
    notes:
    - string: 6
      fret: 7
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 7
      pc: 4
      note: E
      interval: '4'
      role: fourth
      finger: 3
    - string: 4
      fret: 4
      pc: 6
      note: F#
      interval: '5'
      role: fifth
      finger: 1
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 1
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '4'
      role: fourth
      finger: 2
    - string: 1
      fret: 7
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 3
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: b-suspended-4-voicing-3
    start_fret: 7
    end_fret: 9
    notes:
    - string: 6
      fret: 7
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 7
      pc: 4
      note: E
      interval: '4'
      role: fourth
      finger: 1
    - string: 4
      fret: 9
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 2
    - string: 3
      fret: 9
      pc: 4
      note: E
      interval: '4'
      role: fourth
      finger: 2
    - string: 2
      fret: 7
      pc: 6
      note: F#
      interval: '5'
      role: fifth
      finger: 1
    - string: 1
      fret: 7
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (2.5 2.4 4.3 0.2 2.1).2
  midi_root: 71
musical_forms:
- form: pitch_set
  id: b-suspended-4-chord-pitch-set
  pitch_classes:
  - 11
  - 4
  - 6
  intervals:
  - '1'
  - '4'
  - '5'
  spelling:
  - B
  - E
  - F#
- form: voicing
  id: b-suspended-4-voicing-1
  position_ref: b-suspended-4-voicing-1
  start_fret: 2
  tab: x-2-2-4-0-2
- form: voicing
  id: b-suspended-4-voicing-2
  position_ref: b-suspended-4-voicing-2
  start_fret: 4
  tab: 7-7-4-4-5-7
- form: voicing
  id: b-suspended-4-voicing-3
  position_ref: b-suspended-4-voicing-3
  start_fret: 7
  tab: 7-7-9-9-7-7
- form: tab_excerpt
  id: b-suspended-4-chord-tab
  alphatex: \tempo 90 (2.5 2.4 4.3 0.2 2.1).2
---

# B Suspended 4 Chord

The B suspended 4 chord is a triad with the third replaced by the fourth — tense, wanting to fall to the third. Its tones are **B – E – F#**: B (root), E (perfect fourth), F# (perfect fifth).

| Interval | Note | Name |
|---|---|---|
| 1 | B | root |
| 4 | E | perfect fourth |
| 5 | F# | perfect fifth |

**Playing it:** the primary shape is `x-2-2-4-0-2` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [B Major Triad](/triads/b-major-triad.md)
