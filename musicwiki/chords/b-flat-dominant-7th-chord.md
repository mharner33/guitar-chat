---
type: Concept
slug: /chords/b-flat-dominant-7th-chord
title: Bb Dominant 7th Chord
description: 'The dominant 7th chord built on Bb: Bb, D, F, Ab.'
tags:
- chord
- dominant-7th
- b
parent: /chords
prerequisites:
- /triads/b-flat-major-triad
used_by:
- /chords/b-flat-9th-chord
- /chords/b-flat-11th-chord
- /chords/b-flat-13th-chord
theory:
  category: chord
  quality: dominant-7th
  root:
    pc: 10
    name: Bb
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
  - 10
  - 2
  - 5
  - 8
  spelling:
  - Bb
  - D
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
  - 2
  - 5
  - 8
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - b7
  preferred_positions:
  - id: b-flat-dominant-7th-voicing-1
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
      fret: 1
      pc: 8
      note: Ab
      interval: b7
      role: seventh
      finger: 1
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
  - id: b-flat-dominant-7th-voicing-2
    start_fret: 3
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
      fret: 5
      pc: 2
      note: D
      interval: '3'
      role: third
      finger: 3
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
  - id: b-flat-dominant-7th-voicing-3
    start_fret: 5
    end_fret: 7
    notes:
    - string: 6
      fret: 6
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: '3'
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
      fret: 7
      pc: 2
      note: D
      interval: '3'
      role: third
      finger: 3
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '5'
      role: fifth
      finger: 2
    - string: 1
      fret: 6
      pc: 10
      note: Bb
      interval: '1'
      role: root
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (1.5 0.4 1.3 3.2 1.1).2
  midi_root: 70
musical_forms:
- form: pitch_set
  id: b-flat-dominant-7th-chord-pitch-set
  pitch_classes:
  - 10
  - 2
  - 5
  - 8
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  spelling:
  - Bb
  - D
  - F
  - Ab
- form: voicing
  id: b-flat-dominant-7th-voicing-1
  position_ref: b-flat-dominant-7th-voicing-1
  start_fret: 1
  tab: x-1-0-1-3-1
- form: voicing
  id: b-flat-dominant-7th-voicing-2
  position_ref: b-flat-dominant-7th-voicing-2
  start_fret: 3
  tab: 6-5-3-3-3-4
- form: voicing
  id: b-flat-dominant-7th-voicing-3
  position_ref: b-flat-dominant-7th-voicing-3
  start_fret: 5
  tab: 6-5-6-7-6-6
- form: tab_excerpt
  id: b-flat-dominant-7th-chord-tab
  alphatex: \tempo 90 (1.5 0.4 1.3 3.2 1.1).2
---

# Bb Dominant 7th Chord

The Bb dominant 7th chord is a major triad plus a minor seventh — the restless engine of tonal harmony, pulling toward the tonic a fourth above. Its tones are **Bb – D – F – Ab**: Bb (root), D (major third), F (perfect fifth), Ab (minor seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | Bb | root |
| 3 | D | major third |
| 5 | F | perfect fifth |
| b7 | Ab | minor seventh |

**Playing it:** the primary shape is `x-1-0-1-3-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Bb Major Triad](/triads/b-flat-major-triad.md)
- **Used by:** [Bb 9th Chord](/chords/b-flat-9th-chord.md), [Bb 11th Chord](/chords/b-flat-11th-chord.md), [Bb 13th Chord](/chords/b-flat-13th-chord.md)
