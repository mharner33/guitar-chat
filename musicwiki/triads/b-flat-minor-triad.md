---
type: Concept
slug: /triads/b-flat-minor-triad
title: Bb Minor Triad
description: 'The minor triad built on Bb: Bb, Db, F.'
tags:
- triad
- minor
- b
parent: /triads
prerequisites:
- /scales/b-flat-natural-minor-scale
used_by:
- /chords/b-flat-minor-7th-chord
- /chords/b-flat-minor-major-7th-chord
- /chords/b-flat-minor-6th-chord
theory:
  category: triad
  quality: minor
  root:
    pc: 10
    name: Bb
  intervals:
  - '1'
  - b3
  - '5'
  semitones:
  - 0
  - 3
  - 7
  pitch_classes:
  - 10
  - 1
  - 5
  spelling:
  - Bb
  - Db
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
  - 1
  - 5
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  preferred_positions:
  - id: b-flat-minor-voicing-1
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
      fret: 3
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 3
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
  - id: b-flat-minor-voicing-2
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
  - id: b-flat-minor-voicing-3
    start_fret: 2
    end_fret: 6
    notes:
    - string: 6
      fret: 6
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 4
    - string: 5
      fret: 4
      pc: 1
      note: Db
      interval: b3
      role: third
      finger: 3
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '5'
      role: fifth
      finger: 2
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 2
    - string: 2
      fret: 2
      pc: 1
      note: Db
      interval: b3
      role: third
      finger: 1
    - string: 1
      fret: 6
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 4
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (1.5 3.4 3.3 2.2 1.1).2
  midi_root: 70
musical_forms:
- form: pitch_set
  id: b-flat-minor-triad-pitch-set
  pitch_classes:
  - 10
  - 1
  - 5
  intervals:
  - '1'
  - b3
  - '5'
  spelling:
  - Bb
  - Db
  - F
- form: voicing
  id: b-flat-minor-voicing-1
  position_ref: b-flat-minor-voicing-1
  start_fret: 1
  tab: x-1-3-3-2-1
- form: voicing
  id: b-flat-minor-voicing-2
  position_ref: b-flat-minor-voicing-2
  start_fret: 3
  tab: 6-4-3-3-6-6
- form: voicing
  id: b-flat-minor-voicing-3
  position_ref: b-flat-minor-voicing-3
  start_fret: 2
  tab: 6-4-3-3-2-6
- form: tab_excerpt
  id: b-flat-minor-triad-tab
  alphatex: \tempo 90 (1.5 3.4 3.3 2.2 1.1).2
---

# Bb Minor Triad

The Bb minor triad is a darker but still consonant chord. Its tones are **Bb – Db – F**.

**Construction:** stacked thirds — a minor third (Bb–Db) under a major third (Db–F).

| Interval | Note | Name |
|---|---|---|
| 1 | Bb | root |
| b3 | Db | minor third |
| 5 | F | perfect fifth |

**Playing it:** the primary shape is `x-1-3-3-2-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Bb Natural Minor Scale](/scales/b-flat-natural-minor-scale.md)
- **Used by:** [Bb Minor 7th Chord](/chords/b-flat-minor-7th-chord.md), [Bb Minor Major 7th Chord](/chords/b-flat-minor-major-7th-chord.md), [Bb Minor 6th Chord](/chords/b-flat-minor-6th-chord.md)
