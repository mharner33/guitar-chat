---
type: Concept
slug: /chords/c-dominant-7th-chord
title: C Dominant 7th Chord
description: 'The dominant 7th chord built on C: C, E, G, Bb.'
tags:
- chord
- dominant-7th
- c
parent: /chords
prerequisites:
- /triads/c-major-triad
used_by:
- /chords/c-9th-chord
- /chords/c-11th-chord
- /chords/c-13th-chord
theory:
  category: chord
  quality: dominant-7th
  root:
    pc: 0
    name: C
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  semitones:
  - 0
  - 4
  - 7
  - 10
  pitch_classes:
  - 0
  - 4
  - 7
  - 10
  spelling:
  - C
  - E
  - G
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
  - 4
  - 7
  - 10
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - b7
  preferred_positions:
  - id: c-dominant-7th-voicing-1
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
      fret: 2
      pc: 4
      note: E
      interval: '3'
      role: third
      finger: 2
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
      fret: 3
      pc: 7
      note: G
      interval: '5'
      role: fifth
      finger: 3
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: c-dominant-7th-voicing-2
    start_fret: 3
    end_fret: 5
    notes:
    - string: 5
      fret: 3
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: '5'
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
      fret: 5
      pc: 4
      note: E
      interval: '3'
      role: third
      finger: 2
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: '5'
      role: fifth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: c-dominant-7th-voicing-3
    start_fret: 2
    end_fret: 5
    notes:
    - string: 5
      fret: 3
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: '3'
      role: third
      finger: 1
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: b7
      role: seventh
      finger: 2
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '3'
      role: third
      finger: 3
    - string: 1
      fret: 3
      pc: 7
      note: G
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
  alphatex: \tempo 90 (3.5 2.4 3.3 1.2 3.1).2
  midi_root: 60
musical_forms:
- form: pitch_set
  id: c-dominant-7th-chord-pitch-set
  pitch_classes:
  - 0
  - 4
  - 7
  - 10
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  spelling:
  - C
  - E
  - G
  - Bb
- form: voicing
  id: c-dominant-7th-voicing-1
  position_ref: c-dominant-7th-voicing-1
  start_fret: 1
  tab: x-3-2-3-1-3
- form: voicing
  id: c-dominant-7th-voicing-2
  position_ref: c-dominant-7th-voicing-2
  start_fret: 3
  tab: x-3-5-3-5-3
- form: voicing
  id: c-dominant-7th-voicing-3
  position_ref: c-dominant-7th-voicing-3
  start_fret: 2
  tab: x-3-2-3-5-3
- form: tab_excerpt
  id: c-dominant-7th-chord-tab
  alphatex: \tempo 90 (3.5 2.4 3.3 1.2 3.1).2
---

# C Dominant 7th Chord

The C dominant 7th chord is a major triad plus a minor seventh — the restless engine of tonal harmony, pulling toward the tonic a fourth above. Its tones are **C – E – G – Bb**: C (root), E (major third), G (perfect fifth), Bb (minor seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | C | root |
| 3 | E | major third |
| 5 | G | perfect fifth |
| b7 | Bb | minor seventh |

**Playing it:** the primary shape is `x-3-2-3-1-3` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [C Major Triad](/triads/c-major-triad.md)
- **Used by:** [C 9th Chord](/chords/c-9th-chord.md), [C 11th Chord](/chords/c-11th-chord.md), [C 13th Chord](/chords/c-13th-chord.md)
