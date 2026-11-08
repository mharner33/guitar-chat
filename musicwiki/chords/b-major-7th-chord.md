---
type: Concept
slug: /chords/b-major-7th-chord
title: B Major 7th Chord
description: 'The major 7th chord built on B: B, D#, F#, A#.'
tags:
- chord
- major-7th
- b
parent: /chords
prerequisites:
- /triads/b-major-triad
used_by:
- /chords/b-major-9th-chord
theory:
  category: chord
  quality: major-7th
  root:
    pc: 11
    name: B
  intervals:
  - '1'
  - '3'
  - '5'
  - '7'
  semitones:
  - 0
  - 4
  - 7
  - 11
  pitch_classes:
  - 11
  - 3
  - 6
  - 10
  spelling:
  - B
  - D#
  - F#
  - A#
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
  - 10
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - '7'
  preferred_positions:
  - id: b-major-7th-voicing-1
    start_fret: 1
    end_fret: 3
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
      fret: 3
      pc: 10
      note: A#
      interval: '7'
      role: seventh
      finger: 3
    - string: 2
      fret: 0
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 0
    - string: 1
      fret: 2
      pc: 6
      note: F#
      interval: '5'
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: b-major-7th-voicing-2
    start_fret: 2
    end_fret: 4
    notes:
    - string: 5
      fret: 2
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 4
      pc: 6
      note: F#
      interval: '5'
      role: fifth
      finger: 3
    - string: 3
      fret: 3
      pc: 10
      note: A#
      interval: '7'
      role: seventh
      finger: 2
    - string: 2
      fret: 4
      pc: 3
      note: D#
      interval: '3'
      role: third
      finger: 3
    - string: 1
      fret: 2
      pc: 6
      note: F#
      interval: '5'
      role: fifth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: b-major-7th-voicing-3
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
      fret: 6
      pc: 3
      note: D#
      interval: '3'
      role: third
      finger: 2
    - string: 4
      fret: 4
      pc: 6
      note: F#
      interval: '5'
      role: fifth
      finger: 1
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 1
    - string: 2
      fret: 4
      pc: 3
      note: D#
      interval: '3'
      role: third
      finger: 1
    - string: 1
      fret: 6
      pc: 10
      note: A#
      interval: '7'
      role: seventh
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (2.5 1.4 3.3 0.2 2.1).2
  midi_root: 71
musical_forms:
- form: pitch_set
  id: b-major-7th-chord-pitch-set
  pitch_classes:
  - 11
  - 3
  - 6
  - 10
  intervals:
  - '1'
  - '3'
  - '5'
  - '7'
  spelling:
  - B
  - D#
  - F#
  - A#
- form: voicing
  id: b-major-7th-voicing-1
  position_ref: b-major-7th-voicing-1
  start_fret: 1
  tab: x-2-1-3-0-2
- form: voicing
  id: b-major-7th-voicing-2
  position_ref: b-major-7th-voicing-2
  start_fret: 2
  tab: x-2-4-3-4-2
- form: voicing
  id: b-major-7th-voicing-3
  position_ref: b-major-7th-voicing-3
  start_fret: 4
  tab: 7-6-4-4-4-6
- form: tab_excerpt
  id: b-major-7th-chord-tab
  alphatex: \tempo 90 (2.5 1.4 3.3 0.2 2.1).2
---

# B Major 7th Chord

The B major 7th chord is a major triad plus a major seventh — lush, warm, and at rest. Its tones are **B – D# – F# – A#**: B (root), D# (major third), F# (perfect fifth), A# (major seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | B | root |
| 3 | D# | major third |
| 5 | F# | perfect fifth |
| 7 | A# | major seventh |

**Playing it:** the primary shape is `x-2-1-3-0-2` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [B Major Triad](/triads/b-major-triad.md)
- **Used by:** [B Major 9th Chord](/chords/b-major-9th-chord.md)
