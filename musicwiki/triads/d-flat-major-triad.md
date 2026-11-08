---
type: Concept
slug: /triads/d-flat-major-triad
title: Db Major Triad
description: 'The major triad built on Db: Db, F, Ab.'
tags:
- triad
- major
- d
parent: /triads
prerequisites:
- /scales/d-flat-major-scale
used_by:
- /chords/d-flat-dominant-7th-chord
- /chords/d-flat-major-7th-chord
- /chords/d-flat-6th-chord
- /chords/d-flat-suspended-2-chord
- /chords/d-flat-suspended-4-chord
theory:
  category: triad
  quality: major
  root:
    pc: 1
    name: Db
  intervals:
  - '1'
  - '3'
  - '5'
  semitones:
  - 0
  - 4
  - 7
  pitch_classes:
  - 1
  - 5
  - 8
  spelling:
  - Db
  - F
  - Ab
  enharmonic_policy: prefer_flats
  key_context: Db
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
  - 1
  - 5
  - 8
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  preferred_positions:
  - id: d-flat-major-voicing-1
    start_fret: 1
    end_fret: 4
    notes:
    - string: 5
      fret: 4
      pc: 1
      note: Db
      interval: '1'
      role: root
      finger: 4
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '3'
      role: third
      finger: 3
    - string: 3
      fret: 1
      pc: 8
      note: Ab
      interval: '5'
      role: fifth
      finger: 1
    - string: 2
      fret: 2
      pc: 1
      note: Db
      interval: '1'
      role: root
      finger: 2
    - string: 1
      fret: 1
      pc: 5
      note: F
      interval: '3'
      role: third
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: d-flat-major-voicing-2
    start_fret: 4
    end_fret: 6
    notes:
    - string: 5
      fret: 4
      pc: 1
      note: Db
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 6
      pc: 8
      note: Ab
      interval: '5'
      role: fifth
      finger: 2
    - string: 3
      fret: 6
      pc: 1
      note: Db
      interval: '1'
      role: root
      finger: 2
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '3'
      role: third
      finger: 2
    - string: 1
      fret: 4
      pc: 8
      note: Ab
      interval: '5'
      role: fifth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: d-flat-major-voicing-3
    start_fret: 3
    end_fret: 6
    notes:
    - string: 5
      fret: 4
      pc: 1
      note: Db
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '3'
      role: third
      finger: 1
    - string: 3
      fret: 6
      pc: 1
      note: Db
      interval: '1'
      role: root
      finger: 3
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '3'
      role: third
      finger: 3
    - string: 1
      fret: 4
      pc: 8
      note: Ab
      interval: '5'
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (4.5 3.4 1.3 2.2 1.1).2
  midi_root: 61
musical_forms:
- form: pitch_set
  id: d-flat-major-triad-pitch-set
  pitch_classes:
  - 1
  - 5
  - 8
  intervals:
  - '1'
  - '3'
  - '5'
  spelling:
  - Db
  - F
  - Ab
- form: voicing
  id: d-flat-major-voicing-1
  position_ref: d-flat-major-voicing-1
  start_fret: 1
  tab: x-4-3-1-2-1
- form: voicing
  id: d-flat-major-voicing-2
  position_ref: d-flat-major-voicing-2
  start_fret: 4
  tab: x-4-6-6-6-4
- form: voicing
  id: d-flat-major-voicing-3
  position_ref: d-flat-major-voicing-3
  start_fret: 3
  tab: x-4-3-6-6-4
- form: tab_excerpt
  id: d-flat-major-triad-tab
  alphatex: \tempo 90 (4.5 3.4 1.3 2.2 1.1).2
---

# Db Major Triad

The Db major triad is a bright, stable, consonant chord. Its tones are **Db – F – Ab**.

**Construction:** stacked thirds — a major third (Db–F) under a minor third (F–Ab).

| Interval | Note | Name |
|---|---|---|
| 1 | Db | root |
| 3 | F | major third |
| 5 | Ab | perfect fifth |

**Playing it:** the primary shape is `x-4-3-1-2-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Db Major Scale](/scales/d-flat-major-scale.md)
- **Used by:** [Db Dominant 7th Chord](/chords/d-flat-dominant-7th-chord.md), [Db Major 7th Chord](/chords/d-flat-major-7th-chord.md), [Db 6th Chord](/chords/d-flat-6th-chord.md), [Db Suspended 2 Chord](/chords/d-flat-suspended-2-chord.md), [Db Suspended 4 Chord](/chords/d-flat-suspended-4-chord.md)
