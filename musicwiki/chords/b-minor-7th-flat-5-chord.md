---
type: Concept
slug: /chords/b-minor-7th-flat-5-chord
title: B Minor 7th Flat 5 Chord
description: 'The minor 7th flat 5 chord built on B: B, D, F, A.'
tags:
- chord
- minor-7th-flat-5
- b
parent: /chords
prerequisites:
- /triads/b-diminished-triad
theory:
  category: chord
  quality: minor-7th-flat-5
  root:
    pc: 11
    name: B
  intervals:
  - '1'
  - b3
  - b5
  - b7
  semitones:
  - 0
  - 3
  - 6
  - 10
  pitch_classes:
  - 11
  - 2
  - 5
  - 9
  spelling:
  - B
  - D
  - F
  - A
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
  - 5
  - 9
  emphasize_intervals:
  - '1'
  - b3
  - b5
  - b7
  preferred_positions:
  - id: b-minor-7th-flat-5-voicing-1
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
      fret: 2
      pc: 9
      note: A
      interval: b7
      role: seventh
      finger: 2
    - string: 2
      fret: 0
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 0
    - string: 1
      fret: 1
      pc: 5
      note: F
      interval: b5
      role: fifth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: b-minor-7th-flat-5-voicing-2
    start_fret: 2
    end_fret: 3
    notes:
    - string: 5
      fret: 2
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: b5
      role: fifth
      finger: 2
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: b7
      role: seventh
      finger: 1
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: b3
      role: third
      finger: 2
    mute_strings:
    - 1
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: b-minor-7th-flat-5-voicing-3
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
      fret: 7
      pc: 2
      note: D
      interval: b3
      role: third
      finger: 3
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: b5
      role: fifth
      finger: 2
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: b7
      role: seventh
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (2.5 0.4 2.3 0.2 1.1).2
  midi_root: 71
musical_forms:
- form: pitch_set
  id: b-minor-7th-flat-5-chord-pitch-set
  pitch_classes:
  - 11
  - 2
  - 5
  - 9
  intervals:
  - '1'
  - b3
  - b5
  - b7
  spelling:
  - B
  - D
  - F
  - A
- form: voicing
  id: b-minor-7th-flat-5-voicing-1
  position_ref: b-minor-7th-flat-5-voicing-1
  start_fret: 1
  tab: x-2-0-2-0-1
- form: voicing
  id: b-minor-7th-flat-5-voicing-2
  position_ref: b-minor-7th-flat-5-voicing-2
  start_fret: 2
  tab: x-2-3-2-3-x
- form: voicing
  id: b-minor-7th-flat-5-voicing-3
  position_ref: b-minor-7th-flat-5-voicing-3
  start_fret: 5
  tab: 7-5-7-7-6-5
- form: tab_excerpt
  id: b-minor-7th-flat-5-chord-tab
  alphatex: \tempo 90 (2.5 0.4 2.3 0.2 1.1).2
---

# B Minor 7th Flat 5 Chord

The B minor 7th flat 5 chord is the half-diminished chord: a diminished triad plus a minor seventh, the ii chord of a minor key. Its tones are **B – D – F – A**: B (root), D (minor third), F (diminished fifth), A (minor seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | B | root |
| b3 | D | minor third |
| b5 | F | diminished fifth |
| b7 | A | minor seventh |

**Playing it:** the primary shape is `x-2-0-2-0-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [B Diminished Triad](/triads/b-diminished-triad.md)
