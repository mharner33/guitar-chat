---
type: Concept
slug: /triads/b-minor-triad
title: B Minor Triad
description: 'The minor triad built on B: B, D, F#.'
tags:
- triad
- minor
- b
parent: /triads
prerequisites:
- /scales/b-natural-minor-scale
used_by:
- /chords/b-minor-7th-chord
- /chords/b-minor-major-7th-chord
- /chords/b-minor-6th-chord
theory:
  category: triad
  quality: minor
  root:
    pc: 11
    name: B
  intervals:
  - '1'
  - b3
  - '5'
  semitones:
  - 0
  - 3
  - 7
  pitch_classes:
  - 11
  - 2
  - 6
  spelling:
  - B
  - D
  - F#
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
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  preferred_positions:
  - id: b-minor-voicing-1
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
      fret: 0
      pc: 2
      note: D
      interval: b3
      role: third
      finger: 0
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 2
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
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: b-minor-voicing-2
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
      fret: 5
      pc: 2
      note: D
      interval: b3
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
  - id: b-minor-voicing-3
    start_fret: 3
    end_fret: 7
    notes:
    - string: 6
      fret: 7
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 4
    - string: 5
      fret: 5
      pc: 2
      note: D
      interval: b3
      role: third
      finger: 3
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
      fret: 3
      pc: 2
      note: D
      interval: b3
      role: third
      finger: 1
    - string: 1
      fret: 7
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 4
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (2.5 0.4 4.3 0.2 2.1).2
  midi_root: 71
musical_forms:
- form: pitch_set
  id: b-minor-triad-pitch-set
  pitch_classes:
  - 11
  - 2
  - 6
  intervals:
  - '1'
  - b3
  - '5'
  spelling:
  - B
  - D
  - F#
- form: voicing
  id: b-minor-voicing-1
  position_ref: b-minor-voicing-1
  start_fret: 2
  tab: x-2-0-4-0-2
- form: voicing
  id: b-minor-voicing-2
  position_ref: b-minor-voicing-2
  start_fret: 4
  tab: 7-5-4-4-7-7
- form: voicing
  id: b-minor-voicing-3
  position_ref: b-minor-voicing-3
  start_fret: 3
  tab: 7-5-4-4-3-7
- form: tab_excerpt
  id: b-minor-triad-tab
  alphatex: \tempo 90 (2.5 0.4 4.3 0.2 2.1).2
---

# B Minor Triad

The B minor triad is a darker but still consonant chord. Its tones are **B – D – F#**.

**Construction:** stacked thirds — a minor third (B–D) under a major third (D–F#).

| Interval | Note | Name |
|---|---|---|
| 1 | B | root |
| b3 | D | minor third |
| 5 | F# | perfect fifth |

**Playing it:** the primary shape is `x-2-0-4-0-2` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [B Natural Minor Scale](/scales/b-natural-minor-scale.md)
- **Used by:** [B Minor 7th Chord](/chords/b-minor-7th-chord.md), [B Minor Major 7th Chord](/chords/b-minor-major-7th-chord.md), [B Minor 6th Chord](/chords/b-minor-6th-chord.md)
