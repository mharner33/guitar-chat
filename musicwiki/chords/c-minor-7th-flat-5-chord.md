---
type: Concept
slug: /chords/c-minor-7th-flat-5-chord
title: C Minor 7th Flat 5 Chord
description: 'The minor 7th flat 5 chord built on C: C, Eb, Gb, Bb.'
tags:
- chord
- minor-7th-flat-5
- c
parent: /chords
prerequisites:
- /triads/c-diminished-triad
theory:
  category: chord
  quality: minor-7th-flat-5
  root:
    pc: 0
    name: C
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
  - 0
  - 3
  - 6
  - 10
  spelling:
  - C
  - Eb
  - Gb
  - Bb
  enharmonic_policy: prefer_flats
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
  - 3
  - 6
  - 10
  emphasize_intervals:
  - '1'
  - b3
  - b5
  - b7
  preferred_positions:
  - id: c-minor-7th-flat-5-voicing-1
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
      fret: 1
      pc: 3
      note: Eb
      interval: b3
      role: third
      finger: 1
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: b7
      role: seventh
      finger: 3
    - string: 2
      fret: 1
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 1
    - string: 1
      fret: 2
      pc: 6
      note: Gb
      interval: b5
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: c-minor-7th-flat-5-voicing-2
    start_fret: 2
    end_fret: 4
    notes:
    - string: 5
      fret: 3
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 4
      pc: 6
      note: Gb
      interval: b5
      role: fifth
      finger: 3
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: b7
      role: seventh
      finger: 2
    - string: 2
      fret: 4
      pc: 3
      note: Eb
      interval: b3
      role: third
      finger: 3
    - string: 1
      fret: 2
      pc: 6
      note: Gb
      interval: b5
      role: fifth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: c-minor-7th-flat-5-voicing-3
    start_fret: 3
    end_fret: 4
    notes:
    - string: 5
      fret: 3
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 4
      pc: 6
      note: Gb
      interval: b5
      role: fifth
      finger: 2
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: b7
      role: seventh
      finger: 1
    - string: 2
      fret: 4
      pc: 3
      note: Eb
      interval: b3
      role: third
      finger: 2
    mute_strings:
    - 1
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (3.5 1.4 3.3 1.2 2.1).2
  midi_root: 60
musical_forms:
- form: pitch_set
  id: c-minor-7th-flat-5-chord-pitch-set
  pitch_classes:
  - 0
  - 3
  - 6
  - 10
  intervals:
  - '1'
  - b3
  - b5
  - b7
  spelling:
  - C
  - Eb
  - Gb
  - Bb
- form: voicing
  id: c-minor-7th-flat-5-voicing-1
  position_ref: c-minor-7th-flat-5-voicing-1
  start_fret: 1
  tab: x-3-1-3-1-2
- form: voicing
  id: c-minor-7th-flat-5-voicing-2
  position_ref: c-minor-7th-flat-5-voicing-2
  start_fret: 2
  tab: x-3-4-3-4-2
- form: voicing
  id: c-minor-7th-flat-5-voicing-3
  position_ref: c-minor-7th-flat-5-voicing-3
  start_fret: 3
  tab: x-3-4-3-4-x
- form: tab_excerpt
  id: c-minor-7th-flat-5-chord-tab
  alphatex: \tempo 90 (3.5 1.4 3.3 1.2 2.1).2
---

# C Minor 7th Flat 5 Chord

The C minor 7th flat 5 chord is the half-diminished chord: a diminished triad plus a minor seventh, the ii chord of a minor key. Its tones are **C – Eb – Gb – Bb**: C (root), Eb (minor third), Gb (diminished fifth), Bb (minor seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | C | root |
| b3 | Eb | minor third |
| b5 | Gb | diminished fifth |
| b7 | Bb | minor seventh |

**Playing it:** the primary shape is `x-3-1-3-1-2` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [C Diminished Triad](/triads/c-diminished-triad.md)
