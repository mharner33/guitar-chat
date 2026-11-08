---
type: Concept
slug: /chords/f-sharp-11th-chord
title: F# 11th Chord
description: 'The 11th chord built on F#: F#, A#, C#, E, G#, B.'
tags:
- chord
- 11th
- f
parent: /chords
prerequisites:
- /chords/f-sharp-dominant-7th-chord
theory:
  category: chord
  quality: 11th
  root:
    pc: 6
    name: F#
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
  - 6
  - 10
  - 1
  - 4
  - 8
  - 11
  spelling:
  - F#
  - A#
  - C#
  - E
  - G#
  - B
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
  - 10
  - 1
  - 4
  - 8
  - 11
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '11'
  preferred_positions:
  - id: f-sharp-11th-voicing-1
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
      interval: '11'
      role: eleventh
      finger: 1
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: b7
      role: seventh
      finger: 1
    - string: 3
      fret: 3
      pc: 10
      note: A#
      interval: '3'
      role: third
      finger: 2
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: '5'
      role: fifth
      finger: 1
    - string: 1
      fret: 4
      pc: 8
      note: G#
      interval: '9'
      role: ninth
      finger: 3
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: f-sharp-11th-voicing-2
    start_fret: 1
    end_fret: 2
    notes:
    - string: 6
      fret: 2
      pc: 6
      note: F#
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 1
      pc: 10
      note: A#
      interval: '3'
      role: third
      finger: 1
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: b7
      role: seventh
      finger: 2
    - string: 3
      fret: 1
      pc: 8
      note: G#
      interval: '9'
      role: ninth
      finger: 1
    - string: 2
      fret: 0
      pc: 11
      note: B
      interval: '11'
      role: eleventh
      finger: 0
    - string: 1
      fret: 0
      pc: 4
      note: E
      interval: b7
      role: seventh
      finger: 0
    confidence: 0.85
    tags:
    - computed
    - open
    - barre
  - id: f-sharp-11th-voicing-3
    start_fret: 7
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
      fret: 8
      pc: 10
      note: A#
      interval: '3'
      role: third
      finger: 2
    - string: 3
      fret: 9
      pc: 4
      note: E
      interval: b7
      role: seventh
      finger: 3
    - string: 2
      fret: 9
      pc: 8
      note: G#
      interval: '9'
      role: ninth
      finger: 3
    - string: 1
      fret: 7
      pc: 11
      note: B
      interval: '11'
      role: eleventh
      finger: 1
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (2.6 2.5 2.4 3.3 2.2 4.1).2
  midi_root: 66
musical_forms:
- form: pitch_set
  id: f-sharp-11th-chord-pitch-set
  pitch_classes:
  - 6
  - 10
  - 1
  - 4
  - 8
  - 11
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '11'
  spelling:
  - F#
  - A#
  - C#
  - E
  - G#
  - B
- form: voicing
  id: f-sharp-11th-voicing-1
  position_ref: f-sharp-11th-voicing-1
  start_fret: 2
  tab: 2-2-2-3-2-4
- form: voicing
  id: f-sharp-11th-voicing-2
  position_ref: f-sharp-11th-voicing-2
  start_fret: 1
  tab: 2-1-2-1-0-0
- form: voicing
  id: f-sharp-11th-voicing-3
  position_ref: f-sharp-11th-voicing-3
  start_fret: 7
  tab: x-9-8-9-9-7
- form: tab_excerpt
  id: f-sharp-11th-chord-tab
  alphatex: \tempo 90 (2.6 2.5 2.4 3.3 2.2 4.1).2
---

# F# 11th Chord

The F# 11th chord is a dominant chord extended through the eleventh. Its tones are **F# – A# – C# – E – G# – B**: F# (root), A# (major third), C# (perfect fifth), E (minor seventh), G# (ninth), B (eleventh).

| Interval | Note | Name |
|---|---|---|
| 1 | F# | root |
| 3 | A# | major third |
| 5 | C# | perfect fifth |
| b7 | E | minor seventh |
| 9 | G# | ninth |
| 11 | B | eleventh |

**Playing it:** the primary shape is `2-2-2-3-2-4` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [F# Dominant 7th Chord](/chords/f-sharp-dominant-7th-chord.md)
