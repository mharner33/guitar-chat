---
type: Concept
slug: /chords/a-dominant-7th-chord
title: A Dominant 7th Chord
description: 'The dominant 7th chord built on A: A, C#, E, G.'
tags:
- chord
- dominant-7th
- a
parent: /chords
prerequisites:
- /triads/a-major-triad
used_by:
- /chords/a-9th-chord
- /chords/a-11th-chord
- /chords/a-13th-chord
theory:
  category: chord
  quality: dominant-7th
  root:
    pc: 9
    name: A
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
  - 9
  - 1
  - 4
  - 7
  spelling:
  - A
  - C#
  - E
  - G
  enharmonic_policy: prefer_sharps
  key_context: A
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
  - 9
  - 1
  - 4
  - 7
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - b7
  preferred_positions:
  - id: a-dominant-7th-voicing-1
    start_fret: 2
    end_fret: 2
    notes:
    - string: 5
      fret: 0
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 0
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: '5'
      role: fifth
      finger: 1
    - string: 3
      fret: 0
      pc: 7
      note: G
      interval: b7
      role: seventh
      finger: 0
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: '3'
      role: third
      finger: 1
    - string: 1
      fret: 0
      pc: 4
      note: E
      interval: '5'
      role: fifth
      finger: 0
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: a-dominant-7th-voicing-2
    start_fret: 4
    end_fret: 6
    notes:
    - string: 6
      fret: 5
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 4
      pc: 1
      note: C#
      interval: '3'
      role: third
      finger: 1
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: b7
      role: seventh
      finger: 2
    - string: 3
      fret: 6
      pc: 1
      note: C#
      interval: '3'
      role: third
      finger: 3
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '5'
      role: fifth
      finger: 2
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: a-dominant-7th-voicing-3
    start_fret: 3
    end_fret: 6
    notes:
    - string: 6
      fret: 5
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 4
      pc: 1
      note: C#
      interval: '3'
      role: third
      finger: 2
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: b7
      role: seventh
      finger: 3
    - string: 3
      fret: 6
      pc: 1
      note: C#
      interval: '3'
      role: third
      finger: 4
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '5'
      role: fifth
      finger: 3
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: b7
      role: seventh
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (0.5 2.4 0.3 2.2 0.1).2
  midi_root: 69
musical_forms:
- form: pitch_set
  id: a-dominant-7th-chord-pitch-set
  pitch_classes:
  - 9
  - 1
  - 4
  - 7
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  spelling:
  - A
  - C#
  - E
  - G
- form: voicing
  id: a-dominant-7th-voicing-1
  position_ref: a-dominant-7th-voicing-1
  start_fret: 2
  tab: x-0-2-0-2-0
- form: voicing
  id: a-dominant-7th-voicing-2
  position_ref: a-dominant-7th-voicing-2
  start_fret: 4
  tab: 5-4-5-6-5-5
- form: voicing
  id: a-dominant-7th-voicing-3
  position_ref: a-dominant-7th-voicing-3
  start_fret: 3
  tab: 5-4-5-6-5-3
- form: tab_excerpt
  id: a-dominant-7th-chord-tab
  alphatex: \tempo 90 (0.5 2.4 0.3 2.2 0.1).2
---

# A Dominant 7th Chord

The A dominant 7th chord is a major triad plus a minor seventh — the restless engine of tonal harmony, pulling toward the tonic a fourth above. Its tones are **A – C# – E – G**: A (root), C# (major third), E (perfect fifth), G (minor seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | A | root |
| 3 | C# | major third |
| 5 | E | perfect fifth |
| b7 | G | minor seventh |

**Playing it:** the primary shape is `x-0-2-0-2-0` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [A Major Triad](/triads/a-major-triad.md)
- **Used by:** [A 9th Chord](/chords/a-9th-chord.md), [A 11th Chord](/chords/a-11th-chord.md), [A 13th Chord](/chords/a-13th-chord.md)
