---
type: Concept
slug: /chords/b-diminished-7th-chord
title: B Diminished 7th Chord
description: 'The diminished 7th chord built on B: B, D, F, Ab.'
tags:
- chord
- diminished-7th
- b
parent: /chords
prerequisites:
- /triads/b-diminished-triad
theory:
  category: chord
  quality: diminished-7th
  root:
    pc: 11
    name: B
  intervals:
  - '1'
  - b3
  - b5
  - bb7
  semitones:
  - 0
  - 3
  - 6
  - 9
  pitch_classes:
  - 11
  - 2
  - 5
  - 8
  spelling:
  - B
  - D
  - F
  - Ab
  enharmonic_policy: prefer_flats
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
  - 8
  emphasize_intervals:
  - '1'
  - b3
  - b5
  - bb7
  preferred_positions:
  - id: b-diminished-7th-voicing-1
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
      note: Ab
      interval: bb7
      role: seventh
      finger: 1
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
  - id: b-diminished-7th-voicing-2
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
      fret: 3
      pc: 5
      note: F
      interval: b5
      role: fifth
      finger: 2
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
      note: Ab
      interval: bb7
      role: seventh
      finger: 3
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: b-diminished-7th-voicing-3
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
      note: Ab
      interval: bb7
      role: seventh
      finger: 2
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
  alphatex: \tempo 90 (2.5 0.4 1.3 0.2 1.1).2
  midi_root: 71
musical_forms:
- form: pitch_set
  id: b-diminished-7th-chord-pitch-set
  pitch_classes:
  - 11
  - 2
  - 5
  - 8
  intervals:
  - '1'
  - b3
  - b5
  - bb7
  spelling:
  - B
  - D
  - F
  - Ab
- form: voicing
  id: b-diminished-7th-voicing-1
  position_ref: b-diminished-7th-voicing-1
  start_fret: 1
  tab: x-2-0-1-0-1
- form: voicing
  id: b-diminished-7th-voicing-2
  position_ref: b-diminished-7th-voicing-2
  start_fret: 2
  tab: x-2-3-4-3-4
- form: voicing
  id: b-diminished-7th-voicing-3
  position_ref: b-diminished-7th-voicing-3
  start_fret: 5
  tab: 7-5-6-7-6-7
- form: tab_excerpt
  id: b-diminished-7th-chord-tab
  alphatex: \tempo 90 (2.5 0.4 1.3 0.2 1.1).2
---

# B Diminished 7th Chord

The B diminished 7th chord is four stacked minor thirds — symmetrical and maximally tense, able to resolve many directions. Its tones are **B – D – F – Ab**: B (root), D (minor third), F (diminished fifth), Ab (diminished seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | B | root |
| b3 | D | minor third |
| b5 | F | diminished fifth |
| bb7 | Ab | diminished seventh |

**Playing it:** the primary shape is `x-2-0-1-0-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [B Diminished Triad](/triads/b-diminished-triad.md)
