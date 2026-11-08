---
type: Concept
slug: /triads/g-sharp-minor-triad
title: G# Minor Triad
description: 'The minor triad built on G#: G#, B, D#.'
tags:
- triad
- minor
- g
parent: /triads
prerequisites:
- /scales/g-sharp-natural-minor-scale
used_by:
- /chords/g-sharp-minor-7th-chord
- /chords/a-flat-minor-major-7th-chord
- /chords/a-flat-minor-6th-chord
theory:
  category: triad
  quality: minor
  root:
    pc: 8
    name: G#
  intervals:
  - '1'
  - b3
  - '5'
  semitones:
  - 0
  - 3
  - 7
  pitch_classes:
  - 8
  - 11
  - 3
  spelling:
  - G#
  - B
  - D#
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
  - 3
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  preferred_positions:
  - id: g-sharp-minor-voicing-1
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
      fret: 1
      pc: 3
      note: D#
      interval: '5'
      role: fifth
      finger: 1
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
  - id: g-sharp-minor-voicing-2
    start_fret: 4
    end_fret: 6
    notes:
    - string: 6
      fret: 4
      pc: 8
      note: G#
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 6
      pc: 3
      note: D#
      interval: '5'
      role: fifth
      finger: 2
    - string: 4
      fret: 6
      pc: 8
      note: G#
      interval: '1'
      role: root
      finger: 2
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: b3
      role: third
      finger: 1
    - string: 2
      fret: 4
      pc: 3
      note: D#
      interval: '5'
      role: fifth
      finger: 1
    - string: 1
      fret: 4
      pc: 8
      note: G#
      interval: '1'
      role: root
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: g-sharp-minor-voicing-3
    start_fret: 2
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
      fret: 2
      pc: 11
      note: B
      interval: b3
      role: third
      finger: 1
    - string: 4
      fret: 6
      pc: 8
      note: G#
      interval: '1'
      role: root
      finger: 3
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: b3
      role: third
      finger: 2
    - string: 2
      fret: 4
      pc: 3
      note: D#
      interval: '5'
      role: fifth
      finger: 2
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
  alphatex: \tempo 90 (4.6 2.5 1.4 1.3 0.2 4.1).2
  midi_root: 68
musical_forms:
- form: pitch_set
  id: g-sharp-minor-triad-pitch-set
  pitch_classes:
  - 8
  - 11
  - 3
  intervals:
  - '1'
  - b3
  - '5'
  spelling:
  - G#
  - B
  - D#
- form: voicing
  id: g-sharp-minor-voicing-1
  position_ref: g-sharp-minor-voicing-1
  start_fret: 1
  tab: 4-2-1-1-0-4
- form: voicing
  id: g-sharp-minor-voicing-2
  position_ref: g-sharp-minor-voicing-2
  start_fret: 4
  tab: 4-6-6-4-4-4
- form: voicing
  id: g-sharp-minor-voicing-3
  position_ref: g-sharp-minor-voicing-3
  start_fret: 2
  tab: 4-2-6-4-4-4
- form: tab_excerpt
  id: g-sharp-minor-triad-tab
  alphatex: \tempo 90 (4.6 2.5 1.4 1.3 0.2 4.1).2
---

# G# Minor Triad

The G# minor triad is a darker but still consonant chord. Its tones are **G# – B – D#**.

**Construction:** stacked thirds — a minor third (G#–B) under a major third (B–D#).

| Interval | Note | Name |
|---|---|---|
| 1 | G# | root |
| b3 | B | minor third |
| 5 | D# | perfect fifth |

**Playing it:** the primary shape is `4-2-1-1-0-4` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [G# Natural Minor Scale](/scales/g-sharp-natural-minor-scale.md)
- **Used by:** [G# Minor 7th Chord](/chords/g-sharp-minor-7th-chord.md), [Ab Minor Major 7th Chord](/chords/a-flat-minor-major-7th-chord.md), [Ab Minor 6th Chord](/chords/a-flat-minor-6th-chord.md)
