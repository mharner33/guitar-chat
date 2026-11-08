---
type: Concept
slug: /chords/c-sharp-minor-6th-chord
title: C# Minor 6th Chord
description: 'The minor 6th chord built on C#: C#, E, G#, A#.'
tags:
- chord
- minor-6th
- c
parent: /chords
prerequisites:
- /triads/c-sharp-minor-triad
theory:
  category: chord
  quality: minor-6th
  root:
    pc: 1
    name: C#
  intervals:
  - '1'
  - b3
  - '5'
  - '6'
  semitones:
  - 0
  - 3
  - 7
  - 9
  pitch_classes:
  - 1
  - 4
  - 8
  - 10
  spelling:
  - C#
  - E
  - G#
  - A#
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
  - 10
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  - '6'
  preferred_positions:
  - id: c-sharp-minor-6th-voicing-1
    start_fret: 2
    end_fret: 4
    notes:
    - string: 5
      fret: 4
      pc: 1
      note: C#
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: b3
      role: third
      finger: 1
    - string: 3
      fret: 3
      pc: 10
      note: A#
      interval: '6'
      role: sixth
      finger: 2
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: '1'
      role: root
      finger: 1
    - string: 1
      fret: 4
      pc: 8
      note: G#
      interval: '5'
      role: fifth
      finger: 3
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: c-sharp-minor-6th-voicing-2
    start_fret: 4
    end_fret: 6
    notes:
    - string: 5
      fret: 4
      pc: 1
      note: C#
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 6
      pc: 8
      note: G#
      interval: '5'
      role: fifth
      finger: 3
    - string: 3
      fret: 6
      pc: 1
      note: C#
      interval: '1'
      role: root
      finger: 3
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: b3
      role: third
      finger: 2
    - string: 1
      fret: 6
      pc: 10
      note: A#
      interval: '6'
      role: sixth
      finger: 3
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: c-sharp-minor-6th-voicing-3
    start_fret: 3
    end_fret: 6
    notes:
    - string: 5
      fret: 4
      pc: 1
      note: C#
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 6
      pc: 8
      note: G#
      interval: '5'
      role: fifth
      finger: 4
    - string: 3
      fret: 3
      pc: 10
      note: A#
      interval: '6'
      role: sixth
      finger: 1
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: b3
      role: third
      finger: 3
    - string: 1
      fret: 4
      pc: 8
      note: G#
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
  alphatex: \tempo 90 (4.5 2.4 3.3 2.2 4.1).2
  midi_root: 61
musical_forms:
- form: pitch_set
  id: c-sharp-minor-6th-chord-pitch-set
  pitch_classes:
  - 1
  - 4
  - 8
  - 10
  intervals:
  - '1'
  - b3
  - '5'
  - '6'
  spelling:
  - C#
  - E
  - G#
  - A#
- form: voicing
  id: c-sharp-minor-6th-voicing-1
  position_ref: c-sharp-minor-6th-voicing-1
  start_fret: 2
  tab: x-4-2-3-2-4
- form: voicing
  id: c-sharp-minor-6th-voicing-2
  position_ref: c-sharp-minor-6th-voicing-2
  start_fret: 4
  tab: x-4-6-6-5-6
- form: voicing
  id: c-sharp-minor-6th-voicing-3
  position_ref: c-sharp-minor-6th-voicing-3
  start_fret: 3
  tab: x-4-6-3-5-4
- form: tab_excerpt
  id: c-sharp-minor-6th-chord-tab
  alphatex: \tempo 90 (4.5 2.4 3.3 2.2 4.1).2
---

# C# Minor 6th Chord

The C# minor 6th chord is a minor triad plus a major sixth — a bright-edged minor color. Its tones are **C# – E – G# – A#**: C# (root), E (minor third), G# (perfect fifth), A# (major sixth).

| Interval | Note | Name |
|---|---|---|
| 1 | C# | root |
| b3 | E | minor third |
| 5 | G# | perfect fifth |
| 6 | A# | major sixth |

**Playing it:** the primary shape is `x-4-2-3-2-4` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [C# Minor Triad](/triads/c-sharp-minor-triad.md)
