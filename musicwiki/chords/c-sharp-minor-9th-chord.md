---
type: Concept
slug: /chords/c-sharp-minor-9th-chord
title: C# Minor 9th Chord
description: 'The minor 9th chord built on C#: C#, E, G#, B, D#.'
tags:
- chord
- minor-9th
- c
parent: /chords
prerequisites:
- /chords/c-sharp-minor-7th-chord
theory:
  category: chord
  quality: minor-9th
  root:
    pc: 1
    name: C#
  intervals:
  - '1'
  - b3
  - '5'
  - b7
  - '9'
  semitones:
  - 0
  - 3
  - 7
  - 10
  - 14
  pitch_classes:
  - 1
  - 4
  - 8
  - 11
  - 3
  spelling:
  - C#
  - E
  - G#
  - B
  - D#
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
  - 8
  - 11
  - 3
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  - b7
  - '9'
  preferred_positions:
  - id: c-sharp-minor-9th-voicing-1
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
      fret: 4
      pc: 11
      note: B
      interval: b7
      role: seventh
      finger: 2
    - string: 2
      fret: 4
      pc: 3
      note: D#
      interval: '9'
      role: ninth
      finger: 2
    - string: 1
      fret: 4
      pc: 8
      note: G#
      interval: '5'
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: c-sharp-minor-9th-voicing-2
    start_fret: 1
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
      fret: 1
      pc: 3
      note: D#
      interval: '9'
      role: ninth
      finger: 1
    - string: 3
      fret: 1
      pc: 8
      note: G#
      interval: '5'
      role: fifth
      finger: 1
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
    confidence: 0.85
    tags:
    - computed
    - open
    - barre
  - id: c-sharp-minor-9th-voicing-3
    start_fret: 7
    end_fret: 9
    notes:
    - string: 6
      fret: 9
      pc: 1
      note: C#
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 7
      pc: 4
      note: E
      interval: b3
      role: third
      finger: 1
    - string: 4
      fret: 9
      pc: 11
      note: B
      interval: b7
      role: seventh
      finger: 3
    - string: 3
      fret: 8
      pc: 3
      note: D#
      interval: '9'
      role: ninth
      finger: 2
    - string: 2
      fret: 9
      pc: 8
      note: G#
      interval: '5'
      role: fifth
      finger: 3
    - string: 1
      fret: 7
      pc: 11
      note: B
      interval: b7
      role: seventh
      finger: 1
    confidence: 0.85
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (4.5 2.4 4.3 4.2 4.1).2
  midi_root: 61
musical_forms:
- form: pitch_set
  id: c-sharp-minor-9th-chord-pitch-set
  pitch_classes:
  - 1
  - 4
  - 8
  - 11
  - 3
  intervals:
  - '1'
  - b3
  - '5'
  - b7
  - '9'
  spelling:
  - C#
  - E
  - G#
  - B
  - D#
- form: voicing
  id: c-sharp-minor-9th-voicing-1
  position_ref: c-sharp-minor-9th-voicing-1
  start_fret: 2
  tab: x-4-2-4-4-4
- form: voicing
  id: c-sharp-minor-9th-voicing-2
  position_ref: c-sharp-minor-9th-voicing-2
  start_fret: 1
  tab: x-4-1-1-0-0
- form: voicing
  id: c-sharp-minor-9th-voicing-3
  position_ref: c-sharp-minor-9th-voicing-3
  start_fret: 7
  tab: 9-7-9-8-9-7
- form: tab_excerpt
  id: c-sharp-minor-9th-chord-tab
  alphatex: \tempo 90 (4.5 2.4 4.3 4.2 4.1).2
---

# C# Minor 9th Chord

The C# minor 9th chord is a minor seventh extended with the ninth — smooth and modern. Its tones are **C# – E – G# – B – D#**: C# (root), E (minor third), G# (perfect fifth), B (minor seventh), D# (ninth).

| Interval | Note | Name |
|---|---|---|
| 1 | C# | root |
| b3 | E | minor third |
| 5 | G# | perfect fifth |
| b7 | B | minor seventh |
| 9 | D# | ninth |

**Playing it:** the primary shape is `x-4-2-4-4-4` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [C# Minor 7th Chord](/chords/c-sharp-minor-7th-chord.md)
