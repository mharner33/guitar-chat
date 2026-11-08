---
type: Concept
slug: /triads/e-flat-major-triad
title: Eb Major Triad
description: 'The major triad built on Eb: Eb, G, Bb.'
tags:
- triad
- major
- e
parent: /triads
prerequisites:
- /scales/e-flat-major-scale
used_by:
- /chords/e-flat-dominant-7th-chord
- /chords/e-flat-major-7th-chord
- /chords/e-flat-6th-chord
- /chords/e-flat-suspended-2-chord
- /chords/e-flat-suspended-4-chord
theory:
  category: triad
  quality: major
  root:
    pc: 3
    name: Eb
  intervals:
  - '1'
  - '3'
  - '5'
  semitones:
  - 0
  - 4
  - 7
  pitch_classes:
  - 3
  - 7
  - 10
  spelling:
  - Eb
  - G
  - Bb
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
  - 7
  - 10
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  preferred_positions:
  - id: e-flat-major-voicing-1
    start_fret: 1
    end_fret: 4
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
      finger: 2
    - string: 2
      fret: 4
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 3
    - string: 1
      fret: 3
      pc: 7
      note: G
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
  - id: e-flat-major-voicing-2
    start_fret: 3
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
      fret: 5
      pc: 7
      note: G
      interval: '3'
      role: third
      finger: 3
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '5'
      role: fifth
      finger: 1
    - string: 2
      fret: 4
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 2
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: '3'
      role: third
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: e-flat-major-voicing-3
    start_fret: 6
    end_fret: 8
    notes:
    - string: 5
      fret: 6
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 8
      pc: 10
      note: Bb
      interval: '5'
      role: fifth
      finger: 2
    - string: 3
      fret: 8
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 2
    - string: 2
      fret: 8
      pc: 7
      note: G
      interval: '3'
      role: third
      finger: 2
    - string: 1
      fret: 6
      pc: 10
      note: Bb
      interval: '5'
      role: fifth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (1.4 3.3 4.2 3.1).2
  midi_root: 63
musical_forms:
- form: pitch_set
  id: e-flat-major-triad-pitch-set
  pitch_classes:
  - 3
  - 7
  - 10
  intervals:
  - '1'
  - '3'
  - '5'
  spelling:
  - Eb
  - G
  - Bb
- form: voicing
  id: e-flat-major-voicing-1
  position_ref: e-flat-major-voicing-1
  start_fret: 1
  tab: x-x-1-3-4-3
- form: voicing
  id: e-flat-major-voicing-2
  position_ref: e-flat-major-voicing-2
  start_fret: 3
  tab: x-6-5-3-4-3
- form: voicing
  id: e-flat-major-voicing-3
  position_ref: e-flat-major-voicing-3
  start_fret: 6
  tab: x-6-8-8-8-6
- form: tab_excerpt
  id: e-flat-major-triad-tab
  alphatex: \tempo 90 (1.4 3.3 4.2 3.1).2
---

# Eb Major Triad

The Eb major triad is a bright, stable, consonant chord. Its tones are **Eb – G – Bb**.

**Construction:** stacked thirds — a major third (Eb–G) under a minor third (G–Bb).

| Interval | Note | Name |
|---|---|---|
| 1 | Eb | root |
| 3 | G | major third |
| 5 | Bb | perfect fifth |

**Playing it:** the primary shape is `x-x-1-3-4-3` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Eb Major Scale](/scales/e-flat-major-scale.md)
- **Used by:** [Eb Dominant 7th Chord](/chords/e-flat-dominant-7th-chord.md), [Eb Major 7th Chord](/chords/e-flat-major-7th-chord.md), [Eb 6th Chord](/chords/e-flat-6th-chord.md), [Eb Suspended 2 Chord](/chords/e-flat-suspended-2-chord.md), [Eb Suspended 4 Chord](/chords/e-flat-suspended-4-chord.md)
