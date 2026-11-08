---
type: Concept
slug: /chords/e-minor-9th-chord
title: E Minor 9th Chord
description: 'The minor 9th chord built on E: E, G, B, D, F#.'
tags:
- chord
- minor-9th
- e
parent: /chords
prerequisites:
- /chords/e-minor-7th-chord
theory:
  category: chord
  quality: minor-9th
  root:
    pc: 4
    name: E
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
  - 4
  - 7
  - 11
  - 2
  - 6
  spelling:
  - E
  - G
  - B
  - D
  - F#
  enharmonic_policy: prefer_sharps
  key_context: E
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
  - 4
  - 7
  - 11
  - 2
  - 6
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  - b7
  - '9'
  preferred_positions:
  - id: e-minor-9th-voicing-1
    start_fret: 2
    end_fret: 2
    notes:
    - string: 6
      fret: 0
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 0
    - string: 5
      fret: 2
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 1
    - string: 4
      fret: 0
      pc: 2
      note: D
      interval: b7
      role: seventh
      finger: 0
    - string: 3
      fret: 0
      pc: 7
      note: G
      interval: b3
      role: third
      finger: 0
    - string: 2
      fret: 0
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 0
    - string: 1
      fret: 2
      pc: 6
      note: F#
      interval: '9'
      role: ninth
      finger: 1
    confidence: 0.85
    tags:
    - computed
    - open
    - barre
  - id: e-minor-9th-voicing-2
    start_fret: 5
    end_fret: 7
    notes:
    - string: 5
      fret: 7
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: b3
      role: third
      finger: 1
    - string: 3
      fret: 7
      pc: 2
      note: D
      interval: b7
      role: seventh
      finger: 2
    - string: 2
      fret: 7
      pc: 6
      note: F#
      interval: '9'
      role: ninth
      finger: 2
    - string: 1
      fret: 7
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: e-minor-9th-voicing-3
    start_fret: 3
    end_fret: 7
    notes:
    - string: 5
      fret: 7
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 4
      pc: 6
      note: F#
      interval: '9'
      role: ninth
      finger: 2
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 2
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: b7
      role: seventh
      finger: 1
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: b3
      role: third
      finger: 1
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (0.6 2.5 0.4 0.3 0.2 2.1).2
  midi_root: 64
musical_forms:
- form: pitch_set
  id: e-minor-9th-chord-pitch-set
  pitch_classes:
  - 4
  - 7
  - 11
  - 2
  - 6
  intervals:
  - '1'
  - b3
  - '5'
  - b7
  - '9'
  spelling:
  - E
  - G
  - B
  - D
  - F#
- form: voicing
  id: e-minor-9th-voicing-1
  position_ref: e-minor-9th-voicing-1
  start_fret: 2
  tab: 0-2-0-0-0-2
- form: voicing
  id: e-minor-9th-voicing-2
  position_ref: e-minor-9th-voicing-2
  start_fret: 5
  tab: x-7-5-7-7-7
- form: voicing
  id: e-minor-9th-voicing-3
  position_ref: e-minor-9th-voicing-3
  start_fret: 3
  tab: x-7-4-4-3-3
- form: tab_excerpt
  id: e-minor-9th-chord-tab
  alphatex: \tempo 90 (0.6 2.5 0.4 0.3 0.2 2.1).2
---

# E Minor 9th Chord

The E minor 9th chord is a minor seventh extended with the ninth — smooth and modern. Its tones are **E – G – B – D – F#**: E (root), G (minor third), B (perfect fifth), D (minor seventh), F# (ninth).

| Interval | Note | Name |
|---|---|---|
| 1 | E | root |
| b3 | G | minor third |
| 5 | B | perfect fifth |
| b7 | D | minor seventh |
| 9 | F# | ninth |

**Playing it:** the primary shape is `0-2-0-0-0-2` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [E Minor 7th Chord](/chords/e-minor-7th-chord.md)
