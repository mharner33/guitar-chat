---
type: Concept
slug: /chords/c-sharp-minor-7th-flat-5-chord
title: C# Minor 7th Flat 5 Chord
description: 'The minor 7th flat 5 chord built on C#: C#, E, G, B.'
tags:
- chord
- minor-7th-flat-5
- c
parent: /chords
prerequisites:
- /triads/c-sharp-diminished-triad
theory:
  category: chord
  quality: minor-7th-flat-5
  root:
    pc: 1
    name: C#
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
  - 1
  - 4
  - 7
  - 11
  spelling:
  - C#
  - E
  - G
  - B
  enharmonic_policy: prefer_sharps
  key_context: C#
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
  - 4
  - 7
  - 11
  emphasize_intervals:
  - '1'
  - b3
  - b5
  - b7
  preferred_positions:
  - id: c-sharp-minor-7th-flat-5-voicing-1
    start_fret: 2
    end_fret: 4
    notes:
    - string: 5
      fret: 4
      pc: 1
      note: C#
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: b3
      role: third
      finger: 1
    - string: 3
      fret: 0
      pc: 7
      note: G
      interval: b5
      role: fifth
      finger: 0
    - string: 2
      fret: 0
      pc: 11
      note: B
      interval: b7
      role: seventh
      finger: 0
    - string: 1
      fret: 0
      pc: 4
      note: E
      interval: b3
      role: third
      finger: 0
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
  - id: c-sharp-minor-7th-flat-5-voicing-2
    start_fret: 3
    end_fret: 5
    notes:
    - string: 5
      fret: 4
      pc: 1
      note: C#
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: b5
      role: fifth
      finger: 3
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: b7
      role: seventh
      finger: 2
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: b3
      role: third
      finger: 3
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: b5
      role: fifth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: c-sharp-minor-7th-flat-5-voicing-3
    start_fret: 4
    end_fret: 5
    notes:
    - string: 5
      fret: 4
      pc: 1
      note: C#
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: b5
      role: fifth
      finger: 2
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: b7
      role: seventh
      finger: 1
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: b3
      role: third
      finger: 2
    mute_strings:
    - 1
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (4.5 2.4 0.3 0.2 0.1).2
  midi_root: 61
musical_forms:
- form: pitch_set
  id: c-sharp-minor-7th-flat-5-chord-pitch-set
  pitch_classes:
  - 1
  - 4
  - 7
  - 11
  intervals:
  - '1'
  - b3
  - b5
  - b7
  spelling:
  - C#
  - E
  - G
  - B
- form: voicing
  id: c-sharp-minor-7th-flat-5-voicing-1
  position_ref: c-sharp-minor-7th-flat-5-voicing-1
  start_fret: 2
  tab: x-4-2-0-0-0
- form: voicing
  id: c-sharp-minor-7th-flat-5-voicing-2
  position_ref: c-sharp-minor-7th-flat-5-voicing-2
  start_fret: 3
  tab: x-4-5-4-5-3
- form: voicing
  id: c-sharp-minor-7th-flat-5-voicing-3
  position_ref: c-sharp-minor-7th-flat-5-voicing-3
  start_fret: 4
  tab: x-4-5-4-5-x
- form: tab_excerpt
  id: c-sharp-minor-7th-flat-5-chord-tab
  alphatex: \tempo 90 (4.5 2.4 0.3 0.2 0.1).2
---

# C# Minor 7th Flat 5 Chord

The C# minor 7th flat 5 chord is the half-diminished chord: a diminished triad plus a minor seventh, the ii chord of a minor key. Its tones are **C# – E – G – B**: C# (root), E (minor third), G (diminished fifth), B (minor seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | C# | root |
| b3 | E | minor third |
| b5 | G | diminished fifth |
| b7 | B | minor seventh |

**Playing it:** the primary shape is `x-4-2-0-0-0` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [C# Diminished Triad](/triads/c-sharp-diminished-triad.md)
