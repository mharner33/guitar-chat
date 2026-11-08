---
type: Concept
slug: /triads/f-major-triad
title: F Major Triad
description: 'The major triad built on F: F, A, C.'
tags:
- triad
- major
- f
parent: /triads
prerequisites:
- /scales/f-major-scale
used_by:
- /chords/f-dominant-7th-chord
- /chords/f-major-7th-chord
- /chords/f-6th-chord
- /chords/f-suspended-2-chord
- /chords/f-suspended-4-chord
theory:
  category: triad
  quality: major
  root:
    pc: 5
    name: F
  intervals:
  - '1'
  - '3'
  - '5'
  semitones:
  - 0
  - 4
  - 7
  pitch_classes:
  - 5
  - 9
  - 0
  spelling:
  - F
  - A
  - C
  enharmonic_policy: prefer_sharps
  key_context: F
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
  - 5
  - 9
  - 0
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  preferred_positions:
  - id: f-major-voicing-1
    start_fret: 1
    end_fret: 3
    notes:
    - string: 6
      fret: 1
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 0
      pc: 9
      note: A
      interval: '3'
      role: third
      finger: 0
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 3
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: '3'
      role: third
      finger: 2
    - string: 2
      fret: 1
      pc: 0
      note: C
      interval: '5'
      role: fifth
      finger: 1
    - string: 1
      fret: 1
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: f-major-voicing-2
    start_fret: 3
    end_fret: 6
    notes:
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: '5'
      role: fifth
      finger: 2
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 3
    - string: 1
      fret: 5
      pc: 9
      note: A
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
  - id: f-major-voicing-3
    start_fret: 5
    end_fret: 8
    notes:
    - string: 5
      fret: 8
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 4
    - string: 4
      fret: 7
      pc: 9
      note: A
      interval: '3'
      role: third
      finger: 3
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: '5'
      role: fifth
      finger: 1
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 2
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: '3'
      role: third
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (1.6 0.5 3.4 2.3 1.2 1.1).2
  midi_root: 65
musical_forms:
- form: pitch_set
  id: f-major-triad-pitch-set
  pitch_classes:
  - 5
  - 9
  - 0
  intervals:
  - '1'
  - '3'
  - '5'
  spelling:
  - F
  - A
  - C
- form: voicing
  id: f-major-voicing-1
  position_ref: f-major-voicing-1
  start_fret: 1
  tab: 1-0-3-2-1-1
- form: voicing
  id: f-major-voicing-2
  position_ref: f-major-voicing-2
  start_fret: 3
  tab: x-x-3-5-6-5
- form: voicing
  id: f-major-voicing-3
  position_ref: f-major-voicing-3
  start_fret: 5
  tab: x-8-7-5-6-5
- form: tab_excerpt
  id: f-major-triad-tab
  alphatex: \tempo 90 (1.6 0.5 3.4 2.3 1.2 1.1).2
---

# F Major Triad

The F major triad is a bright, stable, consonant chord. Its tones are **F – A – C**.

**Construction:** stacked thirds — a major third (F–A) under a minor third (A–C).

| Interval | Note | Name |
|---|---|---|
| 1 | F | root |
| 3 | A | major third |
| 5 | C | perfect fifth |

**Playing it:** the primary shape is `1-0-3-2-1-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [F Major Scale](/scales/f-major-scale.md)
- **Used by:** [F Dominant 7th Chord](/chords/f-dominant-7th-chord.md), [F Major 7th Chord](/chords/f-major-7th-chord.md), [F 6th Chord](/chords/f-6th-chord.md), [F Suspended 2 Chord](/chords/f-suspended-2-chord.md), [F Suspended 4 Chord](/chords/f-suspended-4-chord.md)
