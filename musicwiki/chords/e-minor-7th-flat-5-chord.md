---
type: Concept
slug: /chords/e-minor-7th-flat-5-chord
title: E Minor 7th Flat 5 Chord
description: 'The minor 7th flat 5 chord built on E: E, G, Bb, D.'
tags:
- chord
- minor-7th-flat-5
- e
parent: /chords
prerequisites:
- /triads/e-diminished-triad
theory:
  category: chord
  quality: minor-7th-flat-5
  root:
    pc: 4
    name: E
  intervals:
  - '1'
  - b3
  - b5
  - b7
  semitones:
  - 0
  - 3
  - 6
  - 10
  pitch_classes:
  - 4
  - 7
  - 10
  - 2
  spelling:
  - E
  - G
  - Bb
  - D
  enharmonic_policy: prefer_flats
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
  - 7
  - 10
  - 2
  emphasize_intervals:
  - '1'
  - b3
  - b5
  - b7
  preferred_positions:
  - id: e-minor-7th-flat-5-voicing-1
    start_fret: 1
    end_fret: 1
    notes:
    - string: 6
      fret: 0
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 0
    - string: 5
      fret: 1
      pc: 10
      note: Bb
      interval: b5
      role: fifth
      finger: 1
    - string: 4
      fret: 0
      pc: 2
      note: D
      interval: b7
      role: seventh
      finger: 0
    - string: 3
      fret: 0
      pc: 7
      note: G
      interval: b3
      role: third
      finger: 0
    mute_strings:
    - 1
    - 2
    confidence: 0.75
    tags:
    - computed
    - open
  - id: e-minor-7th-flat-5-voicing-2
    start_fret: 2
    end_fret: 3
    notes:
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: b5
      role: fifth
      finger: 2
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: b7
      role: seventh
      finger: 2
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: b3
      role: third
      finger: 2
    mute_strings:
    - 5
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: e-minor-7th-flat-5-voicing-3
    start_fret: 5
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
      fret: 5
      pc: 7
      note: G
      interval: b3
      role: third
      finger: 1
    - string: 3
      fret: 7
      pc: 2
      note: D
      interval: b7
      role: seventh
      finger: 3
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 1
    - string: 1
      fret: 6
      pc: 10
      note: Bb
      interval: b5
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (0.6 1.5 0.4 0.3).2
  midi_root: 64
musical_forms:
- form: pitch_set
  id: e-minor-7th-flat-5-chord-pitch-set
  pitch_classes:
  - 4
  - 7
  - 10
  - 2
  intervals:
  - '1'
  - b3
  - b5
  - b7
  spelling:
  - E
  - G
  - Bb
  - D
- form: voicing
  id: e-minor-7th-flat-5-voicing-1
  position_ref: e-minor-7th-flat-5-voicing-1
  start_fret: 1
  tab: 0-1-0-0-x-x
- form: voicing
  id: e-minor-7th-flat-5-voicing-2
  position_ref: e-minor-7th-flat-5-voicing-2
  start_fret: 2
  tab: x-x-2-3-3-3
- form: voicing
  id: e-minor-7th-flat-5-voicing-3
  position_ref: e-minor-7th-flat-5-voicing-3
  start_fret: 5
  tab: x-7-5-7-5-6
- form: tab_excerpt
  id: e-minor-7th-flat-5-chord-tab
  alphatex: \tempo 90 (0.6 1.5 0.4 0.3).2
---

# E Minor 7th Flat 5 Chord

The E minor 7th flat 5 chord is the half-diminished chord: a diminished triad plus a minor seventh, the ii chord of a minor key. Its tones are **E – G – Bb – D**: E (root), G (minor third), Bb (diminished fifth), D (minor seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | E | root |
| b3 | G | minor third |
| b5 | Bb | diminished fifth |
| b7 | D | minor seventh |

**Playing it:** the primary shape is `0-1-0-0-x-x` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [E Diminished Triad](/triads/e-diminished-triad.md)
