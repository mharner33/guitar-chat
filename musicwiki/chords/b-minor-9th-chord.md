---
type: Concept
slug: /chords/b-minor-9th-chord
title: B Minor 9th Chord
description: 'The minor 9th chord built on B: B, D, F#, A, C#.'
tags:
- chord
- minor-9th
- b
parent: /chords
prerequisites:
- /chords/b-minor-7th-chord
theory:
  category: chord
  quality: minor-9th
  root:
    pc: 11
    name: B
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
  - 11
  - 2
  - 6
  - 9
  - 1
  spelling:
  - B
  - D
  - F#
  - A
  - C#
  enharmonic_policy: prefer_sharps
  key_context: B
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
  - 11
  - 2
  - 6
  - 9
  - 1
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  - b7
  - '9'
  preferred_positions:
  - id: b-minor-9th-voicing-1
    start_fret: 2
    end_fret: 2
    notes:
    - string: 5
      fret: 2
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 0
      pc: 2
      note: D
      interval: b3
      role: third
      finger: 0
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: b7
      role: seventh
      finger: 1
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: '9'
      role: ninth
      finger: 1
    - string: 1
      fret: 2
      pc: 6
      note: F#
      interval: '5'
      role: fifth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - open
    - barre
  - id: b-minor-9th-voicing-2
    start_fret: 5
    end_fret: 7
    notes:
    - string: 6
      fret: 7
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: b3
      role: third
      finger: 1
    - string: 4
      fret: 7
      pc: 9
      note: A
      interval: b7
      role: seventh
      finger: 3
    - string: 3
      fret: 6
      pc: 1
      note: C#
      interval: '9'
      role: ninth
      finger: 2
    - string: 2
      fret: 7
      pc: 6
      note: F#
      interval: '5'
      role: fifth
      finger: 3
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: b7
      role: seventh
      finger: 1
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: b-minor-9th-voicing-3
    start_fret: 4
    end_fret: 7
    notes:
    - string: 6
      fret: 7
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 4
      pc: 1
      note: C#
      interval: '9'
      role: ninth
      finger: 1
    - string: 4
      fret: 4
      pc: 6
      note: F#
      interval: '5'
      role: fifth
      finger: 1
    - string: 3
      fret: 7
      pc: 2
      note: D
      interval: b3
      role: third
      finger: 3
    - string: 2
      fret: 7
      pc: 6
      note: F#
      interval: '5'
      role: fifth
      finger: 3
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: b7
      role: seventh
      finger: 2
    confidence: 0.85
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (2.5 0.4 2.3 2.2 2.1).2
  midi_root: 71
musical_forms:
- form: pitch_set
  id: b-minor-9th-chord-pitch-set
  pitch_classes:
  - 11
  - 2
  - 6
  - 9
  - 1
  intervals:
  - '1'
  - b3
  - '5'
  - b7
  - '9'
  spelling:
  - B
  - D
  - F#
  - A
  - C#
- form: voicing
  id: b-minor-9th-voicing-1
  position_ref: b-minor-9th-voicing-1
  start_fret: 2
  tab: x-2-0-2-2-2
- form: voicing
  id: b-minor-9th-voicing-2
  position_ref: b-minor-9th-voicing-2
  start_fret: 5
  tab: 7-5-7-6-7-5
- form: voicing
  id: b-minor-9th-voicing-3
  position_ref: b-minor-9th-voicing-3
  start_fret: 4
  tab: 7-4-4-7-7-5
- form: tab_excerpt
  id: b-minor-9th-chord-tab
  alphatex: \tempo 90 (2.5 0.4 2.3 2.2 2.1).2
---

# B Minor 9th Chord

The B minor 9th chord is a minor seventh extended with the ninth — smooth and modern. Its tones are **B – D – F# – A – C#**: B (root), D (minor third), F# (perfect fifth), A (minor seventh), C# (ninth).

| Interval | Note | Name |
|---|---|---|
| 1 | B | root |
| b3 | D | minor third |
| 5 | F# | perfect fifth |
| b7 | A | minor seventh |
| 9 | C# | ninth |

**Playing it:** the primary shape is `x-2-0-2-2-2` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [B Minor 7th Chord](/chords/b-minor-7th-chord.md)
