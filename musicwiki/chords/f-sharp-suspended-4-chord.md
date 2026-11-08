---
type: Concept
slug: /chords/f-sharp-suspended-4-chord
title: F# Suspended 4 Chord
description: 'The suspended 4 chord built on F#: F#, B, C#.'
tags:
- chord
- suspended-4
- f
parent: /chords
prerequisites:
- /triads/g-flat-major-triad
theory:
  category: chord
  quality: suspended-4
  root:
    pc: 6
    name: F#
  intervals:
  - '1'
  - '4'
  - '5'
  semitones:
  - 0
  - 5
  - 7
  pitch_classes:
  - 6
  - 11
  - 1
  spelling:
  - F#
  - B
  - C#
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
  - 11
  - 1
  emphasize_intervals:
  - '1'
  - '4'
  - '5'
  preferred_positions:
  - id: f-sharp-suspended-4-voicing-1
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
      fret: 2
      pc: 11
      note: B
      interval: '4'
      role: fourth
      finger: 1
    - string: 4
      fret: 4
      pc: 6
      note: F#
      interval: '1'
      role: root
      finger: 2
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '4'
      role: fourth
      finger: 2
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: '5'
      role: fifth
      finger: 1
    - string: 1
      fret: 2
      pc: 6
      note: F#
      interval: '1'
      role: root
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: f-sharp-suspended-4-voicing-2
    start_fret: 4
    end_fret: 7
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
      finger: 2
    - string: 2
      fret: 7
      pc: 6
      note: F#
      interval: '1'
      role: root
      finger: 3
    - string: 1
      fret: 7
      pc: 11
      note: B
      interval: '4'
      role: fourth
      finger: 3
    mute_strings:
    - 5
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: f-sharp-suspended-4-voicing-3
    start_fret: 6
    end_fret: 9
    notes:
    - string: 5
      fret: 9
      pc: 6
      note: F#
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 9
      pc: 11
      note: B
      interval: '4'
      role: fourth
      finger: 3
    - string: 3
      fret: 6
      pc: 1
      note: C#
      interval: '5'
      role: fifth
      finger: 1
    - string: 2
      fret: 7
      pc: 6
      note: F#
      interval: '1'
      role: root
      finger: 2
    - string: 1
      fret: 7
      pc: 11
      note: B
      interval: '4'
      role: fourth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (2.6 2.5 4.4 4.3 2.2 2.1).2
  midi_root: 66
musical_forms:
- form: pitch_set
  id: f-sharp-suspended-4-chord-pitch-set
  pitch_classes:
  - 6
  - 11
  - 1
  intervals:
  - '1'
  - '4'
  - '5'
  spelling:
  - F#
  - B
  - C#
- form: voicing
  id: f-sharp-suspended-4-voicing-1
  position_ref: f-sharp-suspended-4-voicing-1
  start_fret: 2
  tab: 2-2-4-4-2-2
- form: voicing
  id: f-sharp-suspended-4-voicing-2
  position_ref: f-sharp-suspended-4-voicing-2
  start_fret: 4
  tab: x-x-4-6-7-7
- form: voicing
  id: f-sharp-suspended-4-voicing-3
  position_ref: f-sharp-suspended-4-voicing-3
  start_fret: 6
  tab: x-9-9-6-7-7
- form: tab_excerpt
  id: f-sharp-suspended-4-chord-tab
  alphatex: \tempo 90 (2.6 2.5 4.4 4.3 2.2 2.1).2
---

# F# Suspended 4 Chord

The F# suspended 4 chord is a triad with the third replaced by the fourth — tense, wanting to fall to the third. Its tones are **F# – B – C#**: F# (root), B (perfect fourth), C# (perfect fifth).

| Interval | Note | Name |
|---|---|---|
| 1 | F# | root |
| 4 | B | perfect fourth |
| 5 | C# | perfect fifth |

**Playing it:** the primary shape is `2-2-4-4-2-2` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Gb Major Triad](/triads/g-flat-major-triad.md)
