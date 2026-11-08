---
type: Concept
slug: /triads/g-sharp-diminished-triad
title: G# Diminished Triad
description: 'The diminished triad built on G#: G#, B, D.'
tags:
- triad
- diminished
- g
parent: /triads
used_by:
- /chords/g-sharp-minor-7th-flat-5-chord
- /chords/g-sharp-diminished-7th-chord
theory:
  category: triad
  quality: diminished
  root:
    pc: 8
    name: G#
  intervals:
  - '1'
  - b3
  - b5
  semitones:
  - 0
  - 3
  - 6
  pitch_classes:
  - 8
  - 11
  - 2
  spelling:
  - G#
  - B
  - D
  enharmonic_policy: prefer_sharps
  key_context: G#
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
  - 8
  - 11
  - 2
  emphasize_intervals:
  - '1'
  - b3
  - b5
  preferred_positions:
  - id: g-sharp-diminished-voicing-1
    start_fret: 2
    end_fret: 4
    notes:
    - string: 6
      fret: 4
      pc: 8
      note: G#
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 2
      pc: 11
      note: B
      interval: b3
      role: third
      finger: 1
    - string: 4
      fret: 0
      pc: 2
      note: D
      interval: b5
      role: fifth
      finger: 0
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: b3
      role: third
      finger: 2
    - string: 2
      fret: 0
      pc: 11
      note: B
      interval: b3
      role: third
      finger: 0
    - string: 1
      fret: 4
      pc: 8
      note: G#
      interval: '1'
      role: root
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: g-sharp-diminished-voicing-2
    start_fret: 1
    end_fret: 4
    notes:
    - string: 6
      fret: 4
      pc: 8
      note: G#
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 2
      pc: 11
      note: B
      interval: b3
      role: third
      finger: 2
    - string: 4
      fret: 0
      pc: 2
      note: D
      interval: b5
      role: fifth
      finger: 0
    - string: 3
      fret: 1
      pc: 8
      note: G#
      interval: '1'
      role: root
      finger: 1
    - string: 2
      fret: 0
      pc: 11
      note: B
      interval: b3
      role: third
      finger: 0
    - string: 1
      fret: 4
      pc: 8
      note: G#
      interval: '1'
      role: root
      finger: 3
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: g-sharp-diminished-voicing-3
    start_fret: 3
    end_fret: 6
    notes:
    - string: 6
      fret: 4
      pc: 8
      note: G#
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: b5
      role: fifth
      finger: 3
    - string: 4
      fret: 6
      pc: 8
      note: G#
      interval: '1'
      role: root
      finger: 4
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: b3
      role: third
      finger: 2
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: b5
      role: fifth
      finger: 1
    - string: 1
      fret: 4
      pc: 8
      note: G#
      interval: '1'
      role: root
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (4.6 2.5 0.4 4.3 0.2 4.1).2
  midi_root: 68
musical_forms:
- form: pitch_set
  id: g-sharp-diminished-triad-pitch-set
  pitch_classes:
  - 8
  - 11
  - 2
  intervals:
  - '1'
  - b3
  - b5
  spelling:
  - G#
  - B
  - D
- form: voicing
  id: g-sharp-diminished-voicing-1
  position_ref: g-sharp-diminished-voicing-1
  start_fret: 2
  tab: 4-2-0-4-0-4
- form: voicing
  id: g-sharp-diminished-voicing-2
  position_ref: g-sharp-diminished-voicing-2
  start_fret: 1
  tab: 4-2-0-1-0-4
- form: voicing
  id: g-sharp-diminished-voicing-3
  position_ref: g-sharp-diminished-voicing-3
  start_fret: 3
  tab: 4-5-6-4-3-4
- form: tab_excerpt
  id: g-sharp-diminished-triad-tab
  alphatex: \tempo 90 (4.6 2.5 0.4 4.3 0.2 4.1).2
---

# G# Diminished Triad

The G# diminished triad is a tense, unstable chord of two stacked minor thirds that wants to resolve. Its tones are **G# – B – D**.

**Construction:** stacked thirds — a minor third (G#–B) under a minor third (B–D).

| Interval | Note | Name |
|---|---|---|
| 1 | G# | root |
| b3 | B | minor third |
| b5 | D | diminished fifth |

**Playing it:** the primary shape is `4-2-0-4-0-4` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Used by:** [G# Minor 7th Flat 5 Chord](/chords/g-sharp-minor-7th-flat-5-chord.md), [G# Diminished 7th Chord](/chords/g-sharp-diminished-7th-chord.md)
