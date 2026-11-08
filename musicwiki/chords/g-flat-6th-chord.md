---
type: Concept
slug: /chords/g-flat-6th-chord
title: Gb 6th Chord
description: 'The 6th chord built on Gb: Gb, Bb, Db, Eb.'
tags:
- chord
- 6th
- g
parent: /chords
prerequisites:
- /triads/g-flat-major-triad
theory:
  category: chord
  quality: 6th
  root:
    pc: 6
    name: Gb
  intervals:
  - '1'
  - '3'
  - '5'
  - '6'
  semitones:
  - 0
  - 4
  - 7
  - 9
  pitch_classes:
  - 6
  - 10
  - 1
  - 3
  spelling:
  - Gb
  - Bb
  - Db
  - Eb
  enharmonic_policy: prefer_flats
  key_context: Gb
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
  - 6
  - 10
  - 1
  - 3
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - '6'
  preferred_positions:
  - id: g-flat-6th-voicing-1
    start_fret: 1
    end_fret: 3
    notes:
    - string: 6
      fret: 2
      pc: 6
      note: Gb
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 1
      pc: 10
      note: Bb
      interval: '3'
      role: third
      finger: 1
    - string: 4
      fret: 1
      pc: 3
      note: Eb
      interval: '6'
      role: sixth
      finger: 1
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '3'
      role: third
      finger: 3
    - string: 2
      fret: 2
      pc: 1
      note: Db
      interval: '5'
      role: fifth
      finger: 2
    - string: 1
      fret: 2
      pc: 6
      note: Gb
      interval: '1'
      role: root
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: g-flat-6th-voicing-2
    start_fret: 2
    end_fret: 4
    notes:
    - string: 6
      fret: 2
      pc: 6
      note: Gb
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 4
      pc: 1
      note: Db
      interval: '5'
      role: fifth
      finger: 3
    - string: 4
      fret: 4
      pc: 6
      note: Gb
      interval: '1'
      role: root
      finger: 3
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '3'
      role: third
      finger: 2
    - string: 2
      fret: 4
      pc: 3
      note: Eb
      interval: '6'
      role: sixth
      finger: 3
    - string: 1
      fret: 2
      pc: 6
      note: Gb
      interval: '1'
      role: root
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: g-flat-6th-voicing-3
    start_fret: 4
    end_fret: 6
    notes:
    - string: 4
      fret: 4
      pc: 6
      note: Gb
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 6
      pc: 1
      note: Db
      interval: '5'
      role: fifth
      finger: 2
    - string: 2
      fret: 4
      pc: 3
      note: Eb
      interval: '6'
      role: sixth
      finger: 1
    - string: 1
      fret: 6
      pc: 10
      note: Bb
      interval: '3'
      role: third
      finger: 2
    mute_strings:
    - 5
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (2.6 1.5 1.4 3.3 2.2 2.1).2
  midi_root: 66
musical_forms:
- form: pitch_set
  id: g-flat-6th-chord-pitch-set
  pitch_classes:
  - 6
  - 10
  - 1
  - 3
  intervals:
  - '1'
  - '3'
  - '5'
  - '6'
  spelling:
  - Gb
  - Bb
  - Db
  - Eb
- form: voicing
  id: g-flat-6th-voicing-1
  position_ref: g-flat-6th-voicing-1
  start_fret: 1
  tab: 2-1-1-3-2-2
- form: voicing
  id: g-flat-6th-voicing-2
  position_ref: g-flat-6th-voicing-2
  start_fret: 2
  tab: 2-4-4-3-4-2
- form: voicing
  id: g-flat-6th-voicing-3
  position_ref: g-flat-6th-voicing-3
  start_fret: 4
  tab: x-x-4-6-4-6
- form: tab_excerpt
  id: g-flat-6th-chord-tab
  alphatex: \tempo 90 (2.6 1.5 1.4 3.3 2.2 2.1).2
---

# Gb 6th Chord

The Gb 6th chord is a major triad plus a major sixth — a sweet, vintage tonic substitute. Its tones are **Gb – Bb – Db – Eb**: Gb (root), Bb (major third), Db (perfect fifth), Eb (major sixth).

| Interval | Note | Name |
|---|---|---|
| 1 | Gb | root |
| 3 | Bb | major third |
| 5 | Db | perfect fifth |
| 6 | Eb | major sixth |

**Playing it:** the primary shape is `2-1-1-3-2-2` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Gb Major Triad](/triads/g-flat-major-triad.md)
