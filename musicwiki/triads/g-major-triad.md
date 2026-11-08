---
type: Concept
slug: /triads/g-major-triad
title: G Major Triad
description: 'The major triad built on G: G, B, D.'
tags:
- triad
- major
- g
parent: /triads
prerequisites:
- /scales/g-major-scale
used_by:
- /chords/g-dominant-7th-chord
- /chords/g-major-7th-chord
- /chords/g-6th-chord
- /chords/g-suspended-2-chord
- /chords/g-suspended-4-chord
theory:
  category: triad
  quality: major
  root:
    pc: 7
    name: G
  intervals:
  - '1'
  - '3'
  - '5'
  semitones:
  - 0
  - 4
  - 7
  pitch_classes:
  - 7
  - 11
  - 2
  spelling:
  - G
  - B
  - D
  enharmonic_policy: prefer_sharps
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
  - 11
  - 2
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  preferred_positions:
  - id: g-major-voicing-1
    start_fret: 2
    end_fret: 3
    notes:
    - string: 6
      fret: 3
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 2
    - string: 5
      fret: 2
      pc: 11
      note: B
      interval: '3'
      role: third
      finger: 1
    - string: 4
      fret: 0
      pc: 2
      note: D
      interval: '5'
      role: fifth
      finger: 0
    - string: 3
      fret: 0
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 0
    - string: 2
      fret: 0
      pc: 11
      note: B
      interval: '3'
      role: third
      finger: 0
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
    - open
    - barre
  - id: g-major-voicing-2
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
      fret: 5
      pc: 2
      note: D
      interval: '5'
      role: fifth
      finger: 3
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 3
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '3'
      role: third
      finger: 2
    - string: 2
      fret: 3
      pc: 2
      note: D
      interval: '5'
      role: fifth
      finger: 1
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: g-major-voicing-3
    start_fret: 5
    end_fret: 8
    notes:
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 7
      pc: 2
      note: D
      interval: '5'
      role: fifth
      finger: 2
    - string: 2
      fret: 8
      pc: 7
      note: G
      interval: '1'
      role: root
      finger: 3
    - string: 1
      fret: 7
      pc: 11
      note: B
      interval: '3'
      role: third
      finger: 2
    mute_strings:
    - 5
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (3.6 2.5 0.4 0.3 0.2 3.1).2
  midi_root: 67
musical_forms:
- form: pitch_set
  id: g-major-triad-pitch-set
  pitch_classes:
  - 7
  - 11
  - 2
  intervals:
  - '1'
  - '3'
  - '5'
  spelling:
  - G
  - B
  - D
- form: voicing
  id: g-major-voicing-1
  position_ref: g-major-voicing-1
  start_fret: 2
  tab: 3-2-0-0-0-3
- form: voicing
  id: g-major-voicing-2
  position_ref: g-major-voicing-2
  start_fret: 3
  tab: 3-5-5-4-3-3
- form: voicing
  id: g-major-voicing-3
  position_ref: g-major-voicing-3
  start_fret: 5
  tab: x-x-5-7-8-7
- form: tab_excerpt
  id: g-major-triad-tab
  alphatex: \tempo 90 (3.6 2.5 0.4 0.3 0.2 3.1).2
---

# G Major Triad

The G major triad is a bright, stable, consonant chord. Its tones are **G – B – D**.

**Construction:** stacked thirds — a major third (G–B) under a minor third (B–D).

| Interval | Note | Name |
|---|---|---|
| 1 | G | root |
| 3 | B | major third |
| 5 | D | perfect fifth |

**Playing it:** the primary shape is `3-2-0-0-0-3` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [G Major Scale](/scales/g-major-scale.md)
- **Used by:** [G Dominant 7th Chord](/chords/g-dominant-7th-chord.md), [G Major 7th Chord](/chords/g-major-7th-chord.md), [G 6th Chord](/chords/g-6th-chord.md), [G Suspended 2 Chord](/chords/g-suspended-2-chord.md), [G Suspended 4 Chord](/chords/g-suspended-4-chord.md)
