---
type: Concept
slug: /triads/g-flat-major-triad
title: Gb Major Triad
description: 'The major triad built on Gb: Gb, Bb, Db.'
tags:
- triad
- major
- g
parent: /triads
prerequisites:
- /scales/g-flat-major-scale
used_by:
- /chords/f-sharp-dominant-7th-chord
- /chords/g-flat-major-7th-chord
- /chords/g-flat-6th-chord
- /chords/g-flat-suspended-2-chord
- /chords/f-sharp-suspended-4-chord
theory:
  category: triad
  quality: major
  root:
    pc: 6
    name: Gb
  intervals:
  - '1'
  - '3'
  - '5'
  semitones:
  - 0
  - 4
  - 7
  pitch_classes:
  - 6
  - 10
  - 1
  spelling:
  - Gb
  - Bb
  - Db
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
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  preferred_positions:
  - id: g-flat-major-voicing-1
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
      fret: 2
      pc: 1
      note: Db
      interval: '5'
      role: fifth
      finger: 1
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
  - id: g-flat-major-voicing-2
    start_fret: 1
    end_fret: 4
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
      fret: 4
      pc: 6
      note: Gb
      interval: '1'
      role: root
      finger: 4
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
  - id: g-flat-major-voicing-3
    start_fret: 4
    end_fret: 7
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
      fret: 7
      pc: 6
      note: Gb
      interval: '1'
      role: root
      finger: 3
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
  alphatex: \tempo 90 (2.6 4.5 4.4 3.3 2.2 2.1).2
  midi_root: 66
musical_forms:
- form: pitch_set
  id: g-flat-major-triad-pitch-set
  pitch_classes:
  - 6
  - 10
  - 1
  intervals:
  - '1'
  - '3'
  - '5'
  spelling:
  - Gb
  - Bb
  - Db
- form: voicing
  id: g-flat-major-voicing-1
  position_ref: g-flat-major-voicing-1
  start_fret: 2
  tab: 2-4-4-3-2-2
- form: voicing
  id: g-flat-major-voicing-2
  position_ref: g-flat-major-voicing-2
  start_fret: 1
  tab: 2-1-4-3-2-2
- form: voicing
  id: g-flat-major-voicing-3
  position_ref: g-flat-major-voicing-3
  start_fret: 4
  tab: x-x-4-6-7-6
- form: tab_excerpt
  id: g-flat-major-triad-tab
  alphatex: \tempo 90 (2.6 4.5 4.4 3.3 2.2 2.1).2
---

# Gb Major Triad

The Gb major triad is a bright, stable, consonant chord. Its tones are **Gb – Bb – Db**.

**Construction:** stacked thirds — a major third (Gb–Bb) under a minor third (Bb–Db).

| Interval | Note | Name |
|---|---|---|
| 1 | Gb | root |
| 3 | Bb | major third |
| 5 | Db | perfect fifth |

**Playing it:** the primary shape is `2-4-4-3-2-2` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Gb Major Scale](/scales/g-flat-major-scale.md)
- **Used by:** [F# Dominant 7th Chord](/chords/f-sharp-dominant-7th-chord.md), [Gb Major 7th Chord](/chords/g-flat-major-7th-chord.md), [Gb 6th Chord](/chords/g-flat-6th-chord.md), [Gb Suspended 2 Chord](/chords/g-flat-suspended-2-chord.md), [F# Suspended 4 Chord](/chords/f-sharp-suspended-4-chord.md)
