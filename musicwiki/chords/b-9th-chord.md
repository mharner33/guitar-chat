---
type: Concept
slug: /chords/b-9th-chord
title: B 9th Chord
description: 'The 9th chord built on B: B, D#, F#, A, C#.'
tags:
- chord
- 9th
- b
parent: /chords
prerequisites:
- /chords/b-dominant-7th-chord
theory:
  category: chord
  quality: 9th
  root:
    pc: 11
    name: B
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
  - 11
  - 3
  - 6
  - 9
  - 1
  spelling:
  - B
  - D#
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
  - 3
  - 6
  - 9
  - 1
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  preferred_positions:
  - id: b-9th-voicing-1
    start_fret: 1
    end_fret: 2
    notes:
    - string: 5
      fret: 2
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 1
      pc: 3
      note: D#
      interval: '3'
      role: third
      finger: 1
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: b7
      role: seventh
      finger: 2
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: '9'
      role: ninth
      finger: 2
    - string: 1
      fret: 2
      pc: 6
      note: F#
      interval: '5'
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: b-9th-voicing-2
    start_fret: 6
    end_fret: 7
    notes:
    - string: 6
      fret: 7
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 6
      pc: 3
      note: D#
      interval: '3'
      role: third
      finger: 1
    - string: 4
      fret: 7
      pc: 9
      note: A
      interval: b7
      role: seventh
      finger: 2
    - string: 3
      fret: 6
      pc: 1
      note: C#
      interval: '9'
      role: ninth
      finger: 1
    - string: 2
      fret: 7
      pc: 6
      note: F#
      interval: '5'
      role: fifth
      finger: 2
    - string: 1
      fret: 7
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 2
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: b-9th-voicing-3
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
      fret: 6
      pc: 3
      note: D#
      interval: '3'
      role: third
      finger: 2
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
playback:
  alphatex: \tempo 90 (2.5 1.4 2.3 2.2 2.1).2
  midi_root: 71
musical_forms:
- form: pitch_set
  id: b-9th-chord-pitch-set
  pitch_classes:
  - 11
  - 3
  - 6
  - 9
  - 1
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  spelling:
  - B
  - D#
  - F#
  - A
  - C#
- form: voicing
  id: b-9th-voicing-1
  position_ref: b-9th-voicing-1
  start_fret: 1
  tab: x-2-1-2-2-2
- form: voicing
  id: b-9th-voicing-2
  position_ref: b-9th-voicing-2
  start_fret: 6
  tab: 7-6-7-6-7-7
- form: voicing
  id: b-9th-voicing-3
  position_ref: b-9th-voicing-3
  start_fret: 5
  tab: 7-6-7-6-7-5
- form: tab_excerpt
  id: b-9th-chord-tab
  alphatex: \tempo 90 (2.5 1.4 2.3 2.2 2.1).2
---

# B 9th Chord

The B 9th chord is a dominant seventh extended with the ninth — a fuller dominant color. Its tones are **B – D# – F# – A – C#**: B (root), D# (major third), F# (perfect fifth), A (minor seventh), C# (ninth).

| Interval | Note | Name |
|---|---|---|
| 1 | B | root |
| 3 | D# | major third |
| 5 | F# | perfect fifth |
| b7 | A | minor seventh |
| 9 | C# | ninth |

**Playing it:** the primary shape is `x-2-1-2-2-2` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [B Dominant 7th Chord](/chords/b-dominant-7th-chord.md)
