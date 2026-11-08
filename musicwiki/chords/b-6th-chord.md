---
type: Concept
slug: /chords/b-6th-chord
title: B 6th Chord
description: 'The 6th chord built on B: B, D#, F#, G#.'
tags:
- chord
- 6th
- b
parent: /chords
prerequisites:
- /triads/b-major-triad
theory:
  category: chord
  quality: 6th
  root:
    pc: 11
    name: B
  intervals:
  - '1'
  - '3'
  - '5'
  - '6'
  semitones:
  - 0
  - 4
  - 7
  - 9
  pitch_classes:
  - 11
  - 3
  - 6
  - 8
  spelling:
  - B
  - D#
  - F#
  - G#
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
  - 8
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - '6'
  preferred_positions:
  - id: b-6th-voicing-1
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
      fret: 1
      pc: 8
      note: G#
      interval: '6'
      role: sixth
      finger: 1
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
  - id: b-6th-voicing-2
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
      finger: 2
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 2
    - string: 2
      fret: 4
      pc: 3
      note: D#
      interval: '3'
      role: third
      finger: 2
    - string: 1
      fret: 4
      pc: 8
      note: G#
      interval: '6'
      role: sixth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: b-6th-voicing-3
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
      fret: 4
      pc: 8
      note: G#
      interval: '6'
      role: sixth
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (2.5 1.4 1.3 0.2 2.1).2
  midi_root: 71
musical_forms:
- form: pitch_set
  id: b-6th-chord-pitch-set
  pitch_classes:
  - 11
  - 3
  - 6
  - 8
  intervals:
  - '1'
  - '3'
  - '5'
  - '6'
  spelling:
  - B
  - D#
  - F#
  - G#
- form: voicing
  id: b-6th-voicing-1
  position_ref: b-6th-voicing-1
  start_fret: 1
  tab: x-2-1-1-0-2
- form: voicing
  id: b-6th-voicing-2
  position_ref: b-6th-voicing-2
  start_fret: 2
  tab: x-2-4-4-4-4
- form: voicing
  id: b-6th-voicing-3
  position_ref: b-6th-voicing-3
  start_fret: 4
  tab: 7-6-4-4-4-4
- form: tab_excerpt
  id: b-6th-chord-tab
  alphatex: \tempo 90 (2.5 1.4 1.3 0.2 2.1).2
---

# B 6th Chord

The B 6th chord is a major triad plus a major sixth — a sweet, vintage tonic substitute. Its tones are **B – D# – F# – G#**: B (root), D# (major third), F# (perfect fifth), G# (major sixth).

| Interval | Note | Name |
|---|---|---|
| 1 | B | root |
| 3 | D# | major third |
| 5 | F# | perfect fifth |
| 6 | G# | major sixth |

**Playing it:** the primary shape is `x-2-1-1-0-2` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [B Major Triad](/triads/b-major-triad.md)
