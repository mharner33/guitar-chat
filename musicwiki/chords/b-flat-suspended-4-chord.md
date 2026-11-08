---
type: Concept
slug: /chords/b-flat-suspended-4-chord
title: Bb Suspended 4 Chord
description: 'The suspended 4 chord built on Bb: Bb, Eb, F.'
tags:
- chord
- suspended-4
- b
parent: /chords
prerequisites:
- /triads/b-flat-major-triad
theory:
  category: chord
  quality: suspended-4
  root:
    pc: 10
    name: Bb
  intervals:
  - '1'
  - '4'
  - '5'
  semitones:
  - 0
  - 5
  - 7
  pitch_classes:
  - 10
  - 3
  - 5
  spelling:
  - Bb
  - Eb
  - F
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
  - 3
  - 5
  emphasize_intervals:
  - '1'
  - '4'
  - '5'
  preferred_positions:
  - id: b-flat-suspended-4-voicing-1
    start_fret: 1
    end_fret: 4
    notes:
    - string: 5
      fret: 1
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 1
      pc: 3
      note: Eb
      interval: '4'
      role: fourth
      finger: 1
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 2
    - string: 2
      fret: 4
      pc: 3
      note: Eb
      interval: '4'
      role: fourth
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
    - barre
  - id: b-flat-suspended-4-voicing-2
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
      fret: 6
      pc: 3
      note: Eb
      interval: '4'
      role: fourth
      finger: 3
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
      fret: 4
      pc: 3
      note: Eb
      interval: '4'
      role: fourth
      finger: 2
    - string: 1
      fret: 6
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 3
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: b-flat-suspended-4-voicing-3
    start_fret: 6
    end_fret: 8
    notes:
    - string: 6
      fret: 6
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 6
      pc: 3
      note: Eb
      interval: '4'
      role: fourth
      finger: 1
    - string: 4
      fret: 8
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 2
    - string: 3
      fret: 8
      pc: 3
      note: Eb
      interval: '4'
      role: fourth
      finger: 2
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '5'
      role: fifth
      finger: 1
    - string: 1
      fret: 6
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (1.5 1.4 3.3 4.2 1.1).2
  midi_root: 70
musical_forms:
- form: pitch_set
  id: b-flat-suspended-4-chord-pitch-set
  pitch_classes:
  - 10
  - 3
  - 5
  intervals:
  - '1'
  - '4'
  - '5'
  spelling:
  - Bb
  - Eb
  - F
- form: voicing
  id: b-flat-suspended-4-voicing-1
  position_ref: b-flat-suspended-4-voicing-1
  start_fret: 1
  tab: x-1-1-3-4-1
- form: voicing
  id: b-flat-suspended-4-voicing-2
  position_ref: b-flat-suspended-4-voicing-2
  start_fret: 3
  tab: 6-6-3-3-4-6
- form: voicing
  id: b-flat-suspended-4-voicing-3
  position_ref: b-flat-suspended-4-voicing-3
  start_fret: 6
  tab: 6-6-8-8-6-6
- form: tab_excerpt
  id: b-flat-suspended-4-chord-tab
  alphatex: \tempo 90 (1.5 1.4 3.3 4.2 1.1).2
---

# Bb Suspended 4 Chord

The Bb suspended 4 chord is a triad with the third replaced by the fourth — tense, wanting to fall to the third. Its tones are **Bb – Eb – F**: Bb (root), Eb (perfect fourth), F (perfect fifth).

| Interval | Note | Name |
|---|---|---|
| 1 | Bb | root |
| 4 | Eb | perfect fourth |
| 5 | F | perfect fifth |

**Playing it:** the primary shape is `x-1-1-3-4-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Bb Major Triad](/triads/b-flat-major-triad.md)
