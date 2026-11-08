---
type: Concept
slug: /chords/g-minor-7th-flat-5-chord
title: G Minor 7th Flat 5 Chord
description: 'The minor 7th flat 5 chord built on G: G, Bb, Db, F.'
tags:
- chord
- minor-7th-flat-5
- g
parent: /chords
prerequisites:
- /triads/g-diminished-triad
theory:
  category: chord
  quality: minor-7th-flat-5
  root:
    pc: 7
    name: G
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
  - 7
  - 10
  - 1
  - 5
  spelling:
  - G
  - Bb
  - Db
  - F
  enharmonic_policy: prefer_flats
  key_context: G
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
  - 7
  - 10
  - 1
  - 5
  emphasize_intervals:
  - '1'
  - b3
  - b5
  - b7
  preferred_positions:
  - id: g-minor-7th-flat-5-voicing-1
    start_fret: 1
    end_fret: 3
    notes:
    - string: 6
      fret: 3
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 1
      pc: 10
      note: Bb
      interval: b3
      role: third
      finger: 1
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: b7
      role: seventh
      finger: 3
    - string: 3
      fret: 0
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 0
    - string: 2
      fret: 2
      pc: 1
      note: Db
      interval: b5
      role: fifth
      finger: 2
    - string: 1
      fret: 1
      pc: 5
      note: F
      interval: b7
      role: seventh
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: g-minor-7th-flat-5-voicing-2
    start_fret: 2
    end_fret: 4
    notes:
    - string: 6
      fret: 3
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 4
      pc: 1
      note: Db
      interval: b5
      role: fifth
      finger: 3
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: b7
      role: seventh
      finger: 2
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: b3
      role: third
      finger: 2
    - string: 2
      fret: 2
      pc: 1
      note: Db
      interval: b5
      role: fifth
      finger: 1
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: g-minor-7th-flat-5-voicing-3
    start_fret: 3
    end_fret: 4
    notes:
    - string: 6
      fret: 3
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 4
      pc: 1
      note: Db
      interval: b5
      role: fifth
      finger: 2
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: b7
      role: seventh
      finger: 1
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: b3
      role: third
      finger: 1
    mute_strings:
    - 1
    - 2
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (3.6 1.5 3.4 0.3 2.2 1.1).2
  midi_root: 67
musical_forms:
- form: pitch_set
  id: g-minor-7th-flat-5-chord-pitch-set
  pitch_classes:
  - 7
  - 10
  - 1
  - 5
  intervals:
  - '1'
  - b3
  - b5
  - b7
  spelling:
  - G
  - Bb
  - Db
  - F
- form: voicing
  id: g-minor-7th-flat-5-voicing-1
  position_ref: g-minor-7th-flat-5-voicing-1
  start_fret: 1
  tab: 3-1-3-0-2-1
- form: voicing
  id: g-minor-7th-flat-5-voicing-2
  position_ref: g-minor-7th-flat-5-voicing-2
  start_fret: 2
  tab: 3-4-3-3-2-3
- form: voicing
  id: g-minor-7th-flat-5-voicing-3
  position_ref: g-minor-7th-flat-5-voicing-3
  start_fret: 3
  tab: 3-4-3-3-x-x
- form: tab_excerpt
  id: g-minor-7th-flat-5-chord-tab
  alphatex: \tempo 90 (3.6 1.5 3.4 0.3 2.2 1.1).2
---

# G Minor 7th Flat 5 Chord

The G minor 7th flat 5 chord is the half-diminished chord: a diminished triad plus a minor seventh, the ii chord of a minor key. Its tones are **G – Bb – Db – F**: G (root), Bb (minor third), Db (diminished fifth), F (minor seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | G | root |
| b3 | Bb | minor third |
| b5 | Db | diminished fifth |
| b7 | F | minor seventh |

**Playing it:** the primary shape is `3-1-3-0-2-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [G Diminished Triad](/triads/g-diminished-triad.md)
