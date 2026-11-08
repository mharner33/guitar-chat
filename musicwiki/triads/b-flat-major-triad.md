---
type: Concept
slug: /triads/b-flat-major-triad
title: Bb Major Triad
description: 'The major triad built on Bb: Bb, D, F.'
tags:
- triad
- major
- b
parent: /triads
prerequisites:
- /scales/b-flat-major-scale
used_by:
- /chords/b-flat-dominant-7th-chord
- /chords/b-flat-major-7th-chord
- /chords/b-flat-6th-chord
- /chords/b-flat-suspended-2-chord
- /chords/b-flat-suspended-4-chord
theory:
  category: triad
  quality: major
  root:
    pc: 10
    name: Bb
  intervals:
  - '1'
  - '3'
  - '5'
  semitones:
  - 0
  - 4
  - 7
  pitch_classes:
  - 10
  - 2
  - 5
  spelling:
  - Bb
  - D
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
  - 2
  - 5
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  preferred_positions:
  - id: b-flat-major-voicing-1
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
      fret: 0
      pc: 2
      note: D
      interval: '3'
      role: third
      finger: 0
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 2
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '3'
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
    - open
    - barre
  - id: b-flat-major-voicing-2
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
      fret: 5
      pc: 2
      note: D
      interval: '3'
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
      fret: 3
      pc: 2
      note: D
      interval: '3'
      role: third
      finger: 1
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
  - id: b-flat-major-voicing-3
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
      fret: 8
      pc: 5
      note: F
      interval: '5'
      role: fifth
      finger: 3
    - string: 4
      fret: 8
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 3
    - string: 3
      fret: 7
      pc: 2
      note: D
      interval: '3'
      role: third
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
  alphatex: \tempo 90 (1.5 0.4 3.3 3.2 1.1).2
  midi_root: 70
musical_forms:
- form: pitch_set
  id: b-flat-major-triad-pitch-set
  pitch_classes:
  - 10
  - 2
  - 5
  intervals:
  - '1'
  - '3'
  - '5'
  spelling:
  - Bb
  - D
  - F
- form: voicing
  id: b-flat-major-voicing-1
  position_ref: b-flat-major-voicing-1
  start_fret: 1
  tab: x-1-0-3-3-1
- form: voicing
  id: b-flat-major-voicing-2
  position_ref: b-flat-major-voicing-2
  start_fret: 3
  tab: 6-5-3-3-3-6
- form: voicing
  id: b-flat-major-voicing-3
  position_ref: b-flat-major-voicing-3
  start_fret: 6
  tab: 6-8-8-7-6-6
- form: tab_excerpt
  id: b-flat-major-triad-tab
  alphatex: \tempo 90 (1.5 0.4 3.3 3.2 1.1).2
---

# Bb Major Triad

The Bb major triad is a bright, stable, consonant chord. Its tones are **Bb – D – F**.

**Construction:** stacked thirds — a major third (Bb–D) under a minor third (D–F).

| Interval | Note | Name |
|---|---|---|
| 1 | Bb | root |
| 3 | D | major third |
| 5 | F | perfect fifth |

**Playing it:** the primary shape is `x-1-0-3-3-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Bb Major Scale](/scales/b-flat-major-scale.md)
- **Used by:** [Bb Dominant 7th Chord](/chords/b-flat-dominant-7th-chord.md), [Bb Major 7th Chord](/chords/b-flat-major-7th-chord.md), [Bb 6th Chord](/chords/b-flat-6th-chord.md), [Bb Suspended 2 Chord](/chords/b-flat-suspended-2-chord.md), [Bb Suspended 4 Chord](/chords/b-flat-suspended-4-chord.md)
