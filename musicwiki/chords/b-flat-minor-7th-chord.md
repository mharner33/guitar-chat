---
type: Concept
slug: /chords/b-flat-minor-7th-chord
title: Bb Minor 7th Chord
description: 'The minor 7th chord built on Bb: Bb, Db, F, Ab.'
tags:
- chord
- minor-7th
- b
parent: /chords
prerequisites:
- /triads/b-flat-minor-triad
used_by:
- /chords/b-flat-minor-9th-chord
theory:
  category: chord
  quality: minor-7th
  root:
    pc: 10
    name: Bb
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
  - 10
  - 1
  - 5
  - 8
  spelling:
  - Bb
  - Db
  - F
  - Ab
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
  - 1
  - 5
  - 8
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  - b7
  preferred_positions:
  - id: b-flat-minor-7th-voicing-1
    start_fret: 1
    end_fret: 3
    notes:
    - string: 5
      fret: 1
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '5'
      role: fifth
      finger: 3
    - string: 3
      fret: 1
      pc: 8
      note: Ab
      interval: b7
      role: seventh
      finger: 1
    - string: 2
      fret: 2
      pc: 1
      note: Db
      interval: b3
      role: third
      finger: 2
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
  - id: b-flat-minor-7th-voicing-2
    start_fret: 4
    end_fret: 6
    notes:
    - string: 6
      fret: 6
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 4
      pc: 1
      note: Db
      interval: b3
      role: third
      finger: 1
    - string: 4
      fret: 6
      pc: 8
      note: Ab
      interval: b7
      role: seventh
      finger: 2
    - string: 3
      fret: 6
      pc: 1
      note: Db
      interval: b3
      role: third
      finger: 2
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '5'
      role: fifth
      finger: 2
    - string: 1
      fret: 4
      pc: 8
      note: Ab
      interval: b7
      role: seventh
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: b-flat-minor-7th-voicing-3
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
      fret: 4
      pc: 1
      note: Db
      interval: b3
      role: third
      finger: 2
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
      fret: 6
      pc: 5
      note: F
      interval: '5'
      role: fifth
      finger: 3
    - string: 1
      fret: 4
      pc: 8
      note: Ab
      interval: b7
      role: seventh
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (1.5 3.4 1.3 2.2 1.1).2
  midi_root: 70
musical_forms:
- form: pitch_set
  id: b-flat-minor-7th-chord-pitch-set
  pitch_classes:
  - 10
  - 1
  - 5
  - 8
  intervals:
  - '1'
  - b3
  - '5'
  - b7
  spelling:
  - Bb
  - Db
  - F
  - Ab
- form: voicing
  id: b-flat-minor-7th-voicing-1
  position_ref: b-flat-minor-7th-voicing-1
  start_fret: 1
  tab: x-1-3-1-2-1
- form: voicing
  id: b-flat-minor-7th-voicing-2
  position_ref: b-flat-minor-7th-voicing-2
  start_fret: 4
  tab: 6-4-6-6-6-4
- form: voicing
  id: b-flat-minor-7th-voicing-3
  position_ref: b-flat-minor-7th-voicing-3
  start_fret: 3
  tab: 6-4-3-3-6-4
- form: tab_excerpt
  id: b-flat-minor-7th-chord-tab
  alphatex: \tempo 90 (1.5 3.4 1.3 2.2 1.1).2
---

# Bb Minor 7th Chord

The Bb minor 7th chord is a minor triad plus a minor seventh — smooth and mellow, the workhorse of jazz and soul. Its tones are **Bb – Db – F – Ab**: Bb (root), Db (minor third), F (perfect fifth), Ab (minor seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | Bb | root |
| b3 | Db | minor third |
| 5 | F | perfect fifth |
| b7 | Ab | minor seventh |

**Playing it:** the primary shape is `x-1-3-1-2-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Bb Minor Triad](/triads/b-flat-minor-triad.md)
- **Used by:** [Bb Minor 9th Chord](/chords/b-flat-minor-9th-chord.md)
