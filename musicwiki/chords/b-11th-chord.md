---
type: Concept
slug: /chords/b-11th-chord
title: B 11th Chord
description: 'The 11th chord built on B: B, D#, F#, A, C#, E.'
tags:
- chord
- 11th
- b
parent: /chords
prerequisites:
- /chords/b-dominant-7th-chord
theory:
  category: chord
  quality: 11th
  root:
    pc: 11
    name: B
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '11'
  semitones:
  - 0
  - 4
  - 7
  - 10
  - 14
  - 17
  pitch_classes:
  - 11
  - 3
  - 6
  - 9
  - 1
  - 4
  spelling:
  - B
  - D#
  - F#
  - A
  - C#
  - E
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
  - 4
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '11'
  preferred_positions:
  - id: b-11th-voicing-1
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
      fret: 0
      pc: 4
      note: E
      interval: '11'
      role: eleventh
      finger: 0
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - open
    - barre
  - id: b-11th-voicing-2
    start_fret: 4
    end_fret: 7
    notes:
    - string: 6
      fret: 7
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 4
    - string: 5
      fret: 6
      pc: 3
      note: D#
      interval: '3'
      role: third
      finger: 3
    - string: 4
      fret: 4
      pc: 6
      note: F#
      interval: '5'
      role: fifth
      finger: 1
    - string: 3
      fret: 6
      pc: 1
      note: C#
      interval: '9'
      role: ninth
      finger: 3
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '11'
      role: eleventh
      finger: 2
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
  - id: b-11th-voicing-3
    start_fret: 7
    end_fret: 9
    notes:
    - string: 6
      fret: 7
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 7
      pc: 4
      note: E
      interval: '11'
      role: eleventh
      finger: 1
    - string: 4
      fret: 7
      pc: 9
      note: A
      interval: b7
      role: seventh
      finger: 1
    - string: 3
      fret: 8
      pc: 3
      note: D#
      interval: '3'
      role: third
      finger: 2
    - string: 2
      fret: 7
      pc: 6
      note: F#
      interval: '5'
      role: fifth
      finger: 1
    - string: 1
      fret: 9
      pc: 1
      note: C#
      interval: '9'
      role: ninth
      finger: 3
    confidence: 0.85
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (2.5 1.4 2.3 2.2 0.1).2
  midi_root: 71
musical_forms:
- form: pitch_set
  id: b-11th-chord-pitch-set
  pitch_classes:
  - 11
  - 3
  - 6
  - 9
  - 1
  - 4
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '11'
  spelling:
  - B
  - D#
  - F#
  - A
  - C#
  - E
- form: voicing
  id: b-11th-voicing-1
  position_ref: b-11th-voicing-1
  start_fret: 1
  tab: x-2-1-2-2-0
- form: voicing
  id: b-11th-voicing-2
  position_ref: b-11th-voicing-2
  start_fret: 4
  tab: 7-6-4-6-5-5
- form: voicing
  id: b-11th-voicing-3
  position_ref: b-11th-voicing-3
  start_fret: 7
  tab: 7-7-7-8-7-9
- form: tab_excerpt
  id: b-11th-chord-tab
  alphatex: \tempo 90 (2.5 1.4 2.3 2.2 0.1).2
---

# B 11th Chord

The B 11th chord is a dominant chord extended through the eleventh. Its tones are **B – D# – F# – A – C# – E**: B (root), D# (major third), F# (perfect fifth), A (minor seventh), C# (ninth), E (eleventh).

| Interval | Note | Name |
|---|---|---|
| 1 | B | root |
| 3 | D# | major third |
| 5 | F# | perfect fifth |
| b7 | A | minor seventh |
| 9 | C# | ninth |
| 11 | E | eleventh |

**Playing it:** the primary shape is `x-2-1-2-2-0` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [B Dominant 7th Chord](/chords/b-dominant-7th-chord.md)
