---
type: Concept
slug: /triads/g-diminished-triad
title: G Diminished Triad
description: 'The diminished triad built on G: G, Bb, Db.'
tags:
- triad
- diminished
- g
parent: /triads
used_by:
- /chords/g-minor-7th-flat-5-chord
- /chords/g-diminished-7th-chord
theory:
  category: triad
  quality: diminished
  root:
    pc: 7
    name: G
  intervals:
  - '1'
  - b3
  - b5
  semitones:
  - 0
  - 3
  - 6
  pitch_classes:
  - 7
  - 10
  - 1
  spelling:
  - G
  - Bb
  - Db
  enharmonic_policy: prefer_flats
  key_context: G
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
  - 7
  - 10
  - 1
  emphasize_intervals:
  - '1'
  - b3
  - b5
  preferred_positions:
  - id: g-diminished-voicing-1
    start_fret: 2
    end_fret: 5
    notes:
    - string: 6
      fret: 3
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 4
      pc: 1
      note: Db
      interval: b5
      role: fifth
      finger: 3
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 4
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: b3
      role: third
      finger: 2
    - string: 2
      fret: 2
      pc: 1
      note: Db
      interval: b5
      role: fifth
      finger: 1
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: g-diminished-voicing-2
    start_fret: 1
    end_fret: 5
    notes:
    - string: 6
      fret: 3
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 1
      pc: 10
      note: Bb
      interval: b3
      role: third
      finger: 1
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 4
    - string: 3
      fret: 0
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 0
    - string: 2
      fret: 2
      pc: 1
      note: Db
      interval: b5
      role: fifth
      finger: 2
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 3
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: g-diminished-voicing-3
    start_fret: 3
    end_fret: 5
    notes:
    - string: 6
      fret: 3
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 4
      pc: 1
      note: Db
      interval: b5
      role: fifth
      finger: 2
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 3
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: b3
      role: third
      finger: 1
    mute_strings:
    - 1
    - 2
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (3.6 4.5 5.4 3.3 2.2 3.1).2
  midi_root: 67
musical_forms:
- form: pitch_set
  id: g-diminished-triad-pitch-set
  pitch_classes:
  - 7
  - 10
  - 1
  intervals:
  - '1'
  - b3
  - b5
  spelling:
  - G
  - Bb
  - Db
- form: voicing
  id: g-diminished-voicing-1
  position_ref: g-diminished-voicing-1
  start_fret: 2
  tab: 3-4-5-3-2-3
- form: voicing
  id: g-diminished-voicing-2
  position_ref: g-diminished-voicing-2
  start_fret: 1
  tab: 3-1-5-0-2-3
- form: voicing
  id: g-diminished-voicing-3
  position_ref: g-diminished-voicing-3
  start_fret: 3
  tab: 3-4-5-3-x-x
- form: tab_excerpt
  id: g-diminished-triad-tab
  alphatex: \tempo 90 (3.6 4.5 5.4 3.3 2.2 3.1).2
---

# G Diminished Triad

The G diminished triad is a tense, unstable chord of two stacked minor thirds that wants to resolve. Its tones are **G – Bb – Db**.

**Construction:** stacked thirds — a minor third (G–Bb) under a minor third (Bb–Db).

| Interval | Note | Name |
|---|---|---|
| 1 | G | root |
| b3 | Bb | minor third |
| b5 | Db | diminished fifth |

**Playing it:** the primary shape is `3-4-5-3-2-3` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Used by:** [G Minor 7th Flat 5 Chord](/chords/g-minor-7th-flat-5-chord.md), [G Diminished 7th Chord](/chords/g-diminished-7th-chord.md)
