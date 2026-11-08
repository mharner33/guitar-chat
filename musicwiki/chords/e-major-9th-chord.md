---
type: Concept
slug: /chords/e-major-9th-chord
title: E Major 9th Chord
description: 'The major 9th chord built on E: E, G#, B, D#, F#.'
tags:
- chord
- major-9th
- e
parent: /chords
prerequisites:
- /chords/e-major-7th-chord
theory:
  category: chord
  quality: major-9th
  root:
    pc: 4
    name: E
  intervals:
  - '1'
  - '3'
  - '5'
  - '7'
  - '9'
  semitones:
  - 0
  - 4
  - 7
  - 11
  - 14
  pitch_classes:
  - 4
  - 8
  - 11
  - 3
  - 6
  spelling:
  - E
  - G#
  - B
  - D#
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
  - 8
  - 11
  - 3
  - 6
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - '7'
  - '9'
  preferred_positions:
  - id: e-major-9th-voicing-1
    start_fret: 1
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
      finger: 2
    - string: 4
      fret: 1
      pc: 3
      note: D#
      interval: '7'
      role: seventh
      finger: 1
    - string: 3
      fret: 1
      pc: 8
      note: G#
      interval: '3'
      role: third
      finger: 1
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
      finger: 2
    confidence: 0.85
    tags:
    - computed
    - open
    - barre
  - id: e-major-9th-voicing-2
    start_fret: 2
    end_fret: 4
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
      fret: 4
      pc: 3
      note: D#
      interval: '7'
      role: seventh
      finger: 2
    - string: 1
      fret: 4
      pc: 8
      note: G#
      interval: '3'
      role: third
      finger: 2
    confidence: 0.85
    tags:
    - computed
    - open
    - barre
  - id: e-major-9th-voicing-3
    start_fret: 4
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
      fret: 4
      pc: 6
      note: F#
      interval: '9'
      role: ninth
      finger: 1
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 1
    - string: 2
      fret: 4
      pc: 3
      note: D#
      interval: '7'
      role: seventh
      finger: 1
    - string: 1
      fret: 4
      pc: 8
      note: G#
      interval: '3'
      role: third
      finger: 1
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (0.6 2.5 1.4 1.3 0.2 2.1).2
  midi_root: 64
musical_forms:
- form: pitch_set
  id: e-major-9th-chord-pitch-set
  pitch_classes:
  - 4
  - 8
  - 11
  - 3
  - 6
  intervals:
  - '1'
  - '3'
  - '5'
  - '7'
  - '9'
  spelling:
  - E
  - G#
  - B
  - D#
  - F#
- form: voicing
  id: e-major-9th-voicing-1
  position_ref: e-major-9th-voicing-1
  start_fret: 1
  tab: 0-2-1-1-0-2
- form: voicing
  id: e-major-9th-voicing-2
  position_ref: e-major-9th-voicing-2
  start_fret: 2
  tab: 0-2-4-4-4-4
- form: voicing
  id: e-major-9th-voicing-3
  position_ref: e-major-9th-voicing-3
  start_fret: 4
  tab: x-7-4-4-4-4
- form: tab_excerpt
  id: e-major-9th-chord-tab
  alphatex: \tempo 90 (0.6 2.5 1.4 1.3 0.2 2.1).2
---

# E Major 9th Chord

The E major 9th chord is a major seventh extended with the ninth — expansive and warm. Its tones are **E – G# – B – D# – F#**: E (root), G# (major third), B (perfect fifth), D# (major seventh), F# (ninth).

| Interval | Note | Name |
|---|---|---|
| 1 | E | root |
| 3 | G# | major third |
| 5 | B | perfect fifth |
| 7 | D# | major seventh |
| 9 | F# | ninth |

**Playing it:** the primary shape is `0-2-1-1-0-2` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [E Major 7th Chord](/chords/e-major-7th-chord.md)
