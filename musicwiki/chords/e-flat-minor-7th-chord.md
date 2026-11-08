---
type: Concept
slug: /chords/e-flat-minor-7th-chord
title: Eb Minor 7th Chord
description: 'The minor 7th chord built on Eb: Eb, Gb, Bb, Db.'
tags:
- chord
- minor-7th
- e
parent: /chords
prerequisites:
- /triads/e-flat-minor-triad
used_by:
- /chords/e-flat-minor-9th-chord
theory:
  category: chord
  quality: minor-7th
  root:
    pc: 3
    name: Eb
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
  - 3
  - 6
  - 10
  - 1
  spelling:
  - Eb
  - Gb
  - Bb
  - Db
  enharmonic_policy: prefer_flats
  key_context: Eb
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
  - 3
  - 6
  - 10
  - 1
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  - b7
  preferred_positions:
  - id: e-flat-minor-7th-voicing-1
    start_fret: 1
    end_fret: 3
    notes:
    - string: 4
      fret: 1
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '5'
      role: fifth
      finger: 3
    - string: 2
      fret: 2
      pc: 1
      note: Db
      interval: b7
      role: seventh
      finger: 2
    - string: 1
      fret: 2
      pc: 6
      note: Gb
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
  - id: e-flat-minor-7th-voicing-2
    start_fret: 4
    end_fret: 6
    notes:
    - string: 5
      fret: 6
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 4
      pc: 6
      note: Gb
      interval: b3
      role: third
      finger: 1
    - string: 3
      fret: 6
      pc: 1
      note: Db
      interval: b7
      role: seventh
      finger: 2
    - string: 2
      fret: 4
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 1
    - string: 1
      fret: 6
      pc: 10
      note: Bb
      interval: '5'
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: e-flat-minor-7th-voicing-3
    start_fret: 2
    end_fret: 6
    notes:
    - string: 5
      fret: 6
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 4
    - string: 4
      fret: 4
      pc: 6
      note: Gb
      interval: b3
      role: third
      finger: 3
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '5'
      role: fifth
      finger: 2
    - string: 2
      fret: 2
      pc: 1
      note: Db
      interval: b7
      role: seventh
      finger: 1
    - string: 1
      fret: 2
      pc: 6
      note: Gb
      interval: b3
      role: third
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (1.4 3.3 2.2 2.1).2
  midi_root: 63
musical_forms:
- form: pitch_set
  id: e-flat-minor-7th-chord-pitch-set
  pitch_classes:
  - 3
  - 6
  - 10
  - 1
  intervals:
  - '1'
  - b3
  - '5'
  - b7
  spelling:
  - Eb
  - Gb
  - Bb
  - Db
- form: voicing
  id: e-flat-minor-7th-voicing-1
  position_ref: e-flat-minor-7th-voicing-1
  start_fret: 1
  tab: x-x-1-3-2-2
- form: voicing
  id: e-flat-minor-7th-voicing-2
  position_ref: e-flat-minor-7th-voicing-2
  start_fret: 4
  tab: x-6-4-6-4-6
- form: voicing
  id: e-flat-minor-7th-voicing-3
  position_ref: e-flat-minor-7th-voicing-3
  start_fret: 2
  tab: x-6-4-3-2-2
- form: tab_excerpt
  id: e-flat-minor-7th-chord-tab
  alphatex: \tempo 90 (1.4 3.3 2.2 2.1).2
---

# Eb Minor 7th Chord

The Eb minor 7th chord is a minor triad plus a minor seventh — smooth and mellow, the workhorse of jazz and soul. Its tones are **Eb – Gb – Bb – Db**: Eb (root), Gb (minor third), Bb (perfect fifth), Db (minor seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | Eb | root |
| b3 | Gb | minor third |
| 5 | Bb | perfect fifth |
| b7 | Db | minor seventh |

**Playing it:** the primary shape is `x-x-1-3-2-2` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Eb Minor Triad](/triads/e-flat-minor-triad.md)
- **Used by:** [Eb Minor 9th Chord](/chords/e-flat-minor-9th-chord.md)
