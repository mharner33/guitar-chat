---
type: Concept
slug: /chords/a-13th-chord
title: A 13th Chord
description: 'The 13th chord built on A: A, C#, E, G, B, F#.'
tags:
- chord
- 13th
- a
parent: /chords
prerequisites:
- /chords/a-dominant-7th-chord
theory:
  category: chord
  quality: 13th
  root:
    pc: 9
    name: A
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '13'
  semitones:
  - 0
  - 4
  - 7
  - 10
  - 14
  - 21
  pitch_classes:
  - 9
  - 1
  - 4
  - 7
  - 11
  - 6
  spelling:
  - A
  - C#
  - E
  - G
  - B
  - F#
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
  - 6
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '13'
  preferred_positions:
  - id: a-13th-voicing-1
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
      fret: 4
      pc: 6
      note: F#
      interval: '13'
      role: thirteenth
      finger: 2
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
  - id: a-13th-voicing-2
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
      fret: 2
      pc: 6
      note: F#
      interval: '13'
      role: thirteenth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - open
    - barre
  - id: a-13th-voicing-3
    start_fret: 5
    end_fret: 7
    notes:
    - string: 6
      fret: 5
      pc: 9
      note: A
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 7
      pc: 4
      note: E
      interval: '5'
      role: fifth
      finger: 3
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: b7
      role: seventh
      finger: 1
    - string: 3
      fret: 6
      pc: 1
      note: C#
      interval: '3'
      role: third
      finger: 2
    - string: 2
      fret: 7
      pc: 6
      note: F#
      interval: '13'
      role: thirteenth
      finger: 3
    - string: 1
      fret: 7
      pc: 11
      note: B
      interval: '9'
      role: ninth
      finger: 3
    confidence: 0.85
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (5.6 4.5 4.4 4.3 5.2 3.1).2
  midi_root: 69
musical_forms:
- form: pitch_set
  id: a-13th-chord-pitch-set
  pitch_classes:
  - 9
  - 1
  - 4
  - 7
  - 11
  - 6
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '13'
  spelling:
  - A
  - C#
  - E
  - G
  - B
  - F#
- form: voicing
  id: a-13th-voicing-1
  position_ref: a-13th-voicing-1
  start_fret: 3
  tab: 5-4-4-4-5-3
- form: voicing
  id: a-13th-voicing-2
  position_ref: a-13th-voicing-2
  start_fret: 2
  tab: x-0-2-0-2-2
- form: voicing
  id: a-13th-voicing-3
  position_ref: a-13th-voicing-3
  start_fret: 5
  tab: 5-7-5-6-7-7
- form: tab_excerpt
  id: a-13th-chord-tab
  alphatex: \tempo 90 (5.6 4.5 4.4 4.3 5.2 3.1).2
---

# A 13th Chord

The A 13th chord is a dominant chord extended through the thirteenth — the fullest dominant color. Its tones are **A – C# – E – G – B – F#**: A (root), C# (major third), E (perfect fifth), G (minor seventh), B (ninth), F# (thirteenth).

| Interval | Note | Name |
|---|---|---|
| 1 | A | root |
| 3 | C# | major third |
| 5 | E | perfect fifth |
| b7 | G | minor seventh |
| 9 | B | ninth |
| 13 | F# | thirteenth |

**Playing it:** the primary shape is `5-4-4-4-5-3` (strings low→high, `x` = muted), around fret 3. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [A Dominant 7th Chord](/chords/a-dominant-7th-chord.md)
