---
type: Concept
slug: /chords/a-9th-chord
title: A 9th Chord
description: 'The 9th chord built on A: A, C#, E, G, B.'
tags:
- chord
- 9th
- a
parent: /chords
prerequisites:
- /chords/a-dominant-7th-chord
theory:
  category: chord
  quality: 9th
  root:
    pc: 9
    name: A
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  semitones:
  - 0
  - 4
  - 7
  - 10
  - 14
  pitch_classes:
  - 9
  - 1
  - 4
  - 7
  - 11
  spelling:
  - A
  - C#
  - E
  - G
  - B
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
  - 11
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  preferred_positions:
  - id: a-9th-voicing-1
    start_fret: 4
    end_fret: 5
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
      fret: 4
      pc: 11
      note: B
      interval: '9'
      role: ninth
      finger: 1
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
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: a-9th-voicing-2
    start_fret: 3
    end_fret: 5
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
      fret: 4
      pc: 11
      note: B
      interval: '9'
      role: ninth
      finger: 2
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
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: a-9th-voicing-3
    start_fret: 2
    end_fret: 4
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
      fret: 4
      pc: 11
      note: B
      interval: '9'
      role: ninth
      finger: 3
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: '3'
      role: third
      finger: 1
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: b7
      role: seventh
      finger: 2
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - open
    - barre
playback:
  alphatex: \tempo 90 (5.6 4.5 5.4 4.3 5.2 5.1).2
  midi_root: 69
musical_forms:
- form: pitch_set
  id: a-9th-chord-pitch-set
  pitch_classes:
  - 9
  - 1
  - 4
  - 7
  - 11
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  spelling:
  - A
  - C#
  - E
  - G
  - B
- form: voicing
  id: a-9th-voicing-1
  position_ref: a-9th-voicing-1
  start_fret: 4
  tab: 5-4-5-4-5-5
- form: voicing
  id: a-9th-voicing-2
  position_ref: a-9th-voicing-2
  start_fret: 3
  tab: 5-4-5-4-5-3
- form: voicing
  id: a-9th-voicing-3
  position_ref: a-9th-voicing-3
  start_fret: 2
  tab: x-0-2-4-2-3
- form: tab_excerpt
  id: a-9th-chord-tab
  alphatex: \tempo 90 (5.6 4.5 5.4 4.3 5.2 5.1).2
---

# A 9th Chord

The A 9th chord is a dominant seventh extended with the ninth — a fuller dominant color. Its tones are **A – C# – E – G – B**: A (root), C# (major third), E (perfect fifth), G (minor seventh), B (ninth).

| Interval | Note | Name |
|---|---|---|
| 1 | A | root |
| 3 | C# | major third |
| 5 | E | perfect fifth |
| b7 | G | minor seventh |
| 9 | B | ninth |

**Playing it:** the primary shape is `5-4-5-4-5-5` (strings low→high, `x` = muted), around fret 4. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [A Dominant 7th Chord](/chords/a-dominant-7th-chord.md)
