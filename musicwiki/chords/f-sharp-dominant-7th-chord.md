---
type: Concept
slug: /chords/f-sharp-dominant-7th-chord
title: F# Dominant 7th Chord
description: 'The dominant 7th chord built on F#: F#, A#, C#, E.'
tags:
- chord
- dominant-7th
- f
parent: /chords
prerequisites:
- /triads/g-flat-major-triad
used_by:
- /chords/f-sharp-9th-chord
- /chords/f-sharp-11th-chord
- /chords/f-sharp-13th-chord
theory:
  category: chord
  quality: dominant-7th
  root:
    pc: 6
    name: F#
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
  - 6
  - 10
  - 1
  - 4
  spelling:
  - F#
  - A#
  - C#
  - E
  enharmonic_policy: prefer_sharps
  key_context: F#
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
  - 4
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - b7
  preferred_positions:
  - id: f-sharp-dominant-7th-voicing-1
    start_fret: 1
    end_fret: 3
    notes:
    - string: 6
      fret: 2
      pc: 6
      note: F#
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 1
      pc: 10
      note: A#
      interval: '3'
      role: third
      finger: 1
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: b7
      role: seventh
      finger: 2
    - string: 3
      fret: 3
      pc: 10
      note: A#
      interval: '3'
      role: third
      finger: 3
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: '5'
      role: fifth
      finger: 2
    - string: 1
      fret: 0
      pc: 4
      note: E
      interval: b7
      role: seventh
      finger: 0
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: f-sharp-dominant-7th-voicing-2
    start_fret: 2
    end_fret: 4
    notes:
    - string: 6
      fret: 2
      pc: 6
      note: F#
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 4
      pc: 1
      note: C#
      interval: '5'
      role: fifth
      finger: 3
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: b7
      role: seventh
      finger: 1
    - string: 3
      fret: 3
      pc: 10
      note: A#
      interval: '3'
      role: third
      finger: 2
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: '5'
      role: fifth
      finger: 1
    - string: 1
      fret: 0
      pc: 4
      note: E
      interval: b7
      role: seventh
      finger: 0
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: f-sharp-dominant-7th-voicing-3
    start_fret: 4
    end_fret: 6
    notes:
    - string: 4
      fret: 4
      pc: 6
      note: F#
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 6
      pc: 1
      note: C#
      interval: '5'
      role: fifth
      finger: 3
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: b7
      role: seventh
      finger: 2
    - string: 1
      fret: 6
      pc: 10
      note: A#
      interval: '3'
      role: third
      finger: 3
    mute_strings:
    - 5
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (2.6 1.5 2.4 3.3 2.2 0.1).2
  midi_root: 66
musical_forms:
- form: pitch_set
  id: f-sharp-dominant-7th-chord-pitch-set
  pitch_classes:
  - 6
  - 10
  - 1
  - 4
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  spelling:
  - F#
  - A#
  - C#
  - E
- form: voicing
  id: f-sharp-dominant-7th-voicing-1
  position_ref: f-sharp-dominant-7th-voicing-1
  start_fret: 1
  tab: 2-1-2-3-2-0
- form: voicing
  id: f-sharp-dominant-7th-voicing-2
  position_ref: f-sharp-dominant-7th-voicing-2
  start_fret: 2
  tab: 2-4-2-3-2-0
- form: voicing
  id: f-sharp-dominant-7th-voicing-3
  position_ref: f-sharp-dominant-7th-voicing-3
  start_fret: 4
  tab: x-x-4-6-5-6
- form: tab_excerpt
  id: f-sharp-dominant-7th-chord-tab
  alphatex: \tempo 90 (2.6 1.5 2.4 3.3 2.2 0.1).2
---

# F# Dominant 7th Chord

The F# dominant 7th chord is a major triad plus a minor seventh — the restless engine of tonal harmony, pulling toward the tonic a fourth above. Its tones are **F# – A# – C# – E**: F# (root), A# (major third), C# (perfect fifth), E (minor seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | F# | root |
| 3 | A# | major third |
| 5 | C# | perfect fifth |
| b7 | E | minor seventh |

**Playing it:** the primary shape is `2-1-2-3-2-0` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Gb Major Triad](/triads/g-flat-major-triad.md)
- **Used by:** [F# 9th Chord](/chords/f-sharp-9th-chord.md), [F# 11th Chord](/chords/f-sharp-11th-chord.md), [F# 13th Chord](/chords/f-sharp-13th-chord.md)
