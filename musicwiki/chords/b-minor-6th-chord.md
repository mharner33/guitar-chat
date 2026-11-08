---
type: Concept
slug: /chords/b-minor-6th-chord
title: B Minor 6th Chord
description: 'The minor 6th chord built on B: B, D, F#, G#.'
tags:
- chord
- minor-6th
- b
parent: /chords
prerequisites:
- /triads/b-minor-triad
theory:
  category: chord
  quality: minor-6th
  root:
    pc: 11
    name: B
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
  - 11
  - 2
  - 6
  - 8
  spelling:
  - B
  - D
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
  - 2
  - 6
  - 8
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  - '6'
  preferred_positions:
  - id: b-minor-6th-voicing-1
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
      fret: 0
      pc: 2
      note: D
      interval: b3
      role: third
      finger: 0
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
  - id: b-minor-6th-voicing-2
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
      fret: 4
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 3
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: b3
      role: third
      finger: 2
    - string: 1
      fret: 4
      pc: 8
      note: G#
      interval: '6'
      role: sixth
      finger: 3
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: b-minor-6th-voicing-3
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
      fret: 6
      pc: 8
      note: G#
      interval: '6'
      role: sixth
      finger: 2
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
      fret: 7
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 3
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (2.5 0.4 1.3 0.2 2.1).2
  midi_root: 71
musical_forms:
- form: pitch_set
  id: b-minor-6th-chord-pitch-set
  pitch_classes:
  - 11
  - 2
  - 6
  - 8
  intervals:
  - '1'
  - b3
  - '5'
  - '6'
  spelling:
  - B
  - D
  - F#
  - G#
- form: voicing
  id: b-minor-6th-voicing-1
  position_ref: b-minor-6th-voicing-1
  start_fret: 1
  tab: x-2-0-1-0-2
- form: voicing
  id: b-minor-6th-voicing-2
  position_ref: b-minor-6th-voicing-2
  start_fret: 2
  tab: x-2-4-4-3-4
- form: voicing
  id: b-minor-6th-voicing-3
  position_ref: b-minor-6th-voicing-3
  start_fret: 5
  tab: 7-5-6-7-7-7
- form: tab_excerpt
  id: b-minor-6th-chord-tab
  alphatex: \tempo 90 (2.5 0.4 1.3 0.2 2.1).2
---

# B Minor 6th Chord

The B minor 6th chord is a minor triad plus a major sixth — a bright-edged minor color. Its tones are **B – D – F# – G#**: B (root), D (minor third), F# (perfect fifth), G# (major sixth).

| Interval | Note | Name |
|---|---|---|
| 1 | B | root |
| b3 | D | minor third |
| 5 | F# | perfect fifth |
| 6 | G# | major sixth |

**Playing it:** the primary shape is `x-2-0-1-0-2` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [B Minor Triad](/triads/b-minor-triad.md)
