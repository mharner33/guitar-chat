---
type: Concept
slug: /triads/e-major-triad
title: E Major Triad
description: 'The major triad built on E: E, G#, B.'
tags:
- triad
- major
- e
parent: /triads
prerequisites:
- /scales/e-major-scale
used_by:
- /chords/e-dominant-7th-chord
- /chords/e-major-7th-chord
- /chords/e-6th-chord
- /chords/e-suspended-2-chord
- /chords/e-suspended-4-chord
theory:
  category: triad
  quality: major
  root:
    pc: 4
    name: E
  intervals:
  - '1'
  - '3'
  - '5'
  semitones:
  - 0
  - 4
  - 7
  pitch_classes:
  - 4
  - 8
  - 11
  spelling:
  - E
  - G#
  - B
  enharmonic_policy: prefer_sharps
  key_context: E
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
  - 4
  - 8
  - 11
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  preferred_positions:
  - id: e-major-voicing-1
    start_fret: 1
    end_fret: 2
    notes:
    - string: 6
      fret: 0
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 0
    - string: 5
      fret: 2
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 2
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 2
    - string: 3
      fret: 1
      pc: 8
      note: G#
      interval: '3'
      role: third
      finger: 1
    - string: 2
      fret: 0
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 0
    - string: 1
      fret: 0
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 0
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: e-major-voicing-2
    start_fret: 2
    end_fret: 4
    notes:
    - string: 6
      fret: 0
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 0
    - string: 5
      fret: 2
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 1
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 2
    - string: 2
      fret: 0
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 0
    - string: 1
      fret: 4
      pc: 8
      note: G#
      interval: '3'
      role: third
      finger: 2
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: e-major-voicing-3
    start_fret: 4
    end_fret: 7
    notes:
    - string: 5
      fret: 7
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 4
    - string: 4
      fret: 6
      pc: 8
      note: G#
      interval: '3'
      role: third
      finger: 3
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 1
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 2
    - string: 1
      fret: 4
      pc: 8
      note: G#
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
  alphatex: \tempo 90 (0.6 2.5 2.4 1.3 0.2 0.1).2
  midi_root: 64
musical_forms:
- form: pitch_set
  id: e-major-triad-pitch-set
  pitch_classes:
  - 4
  - 8
  - 11
  intervals:
  - '1'
  - '3'
  - '5'
  spelling:
  - E
  - G#
  - B
- form: voicing
  id: e-major-voicing-1
  position_ref: e-major-voicing-1
  start_fret: 1
  tab: 0-2-2-1-0-0
- form: voicing
  id: e-major-voicing-2
  position_ref: e-major-voicing-2
  start_fret: 2
  tab: 0-2-2-4-0-4
- form: voicing
  id: e-major-voicing-3
  position_ref: e-major-voicing-3
  start_fret: 4
  tab: x-7-6-4-5-4
- form: tab_excerpt
  id: e-major-triad-tab
  alphatex: \tempo 90 (0.6 2.5 2.4 1.3 0.2 0.1).2
---

# E Major Triad

The E major triad is a bright, stable, consonant chord. Its tones are **E – G# – B**.

**Construction:** stacked thirds — a major third (E–G#) under a minor third (G#–B).

| Interval | Note | Name |
|---|---|---|
| 1 | E | root |
| 3 | G# | major third |
| 5 | B | perfect fifth |

**Playing it:** the primary shape is `0-2-2-1-0-0` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [E Major Scale](/scales/e-major-scale.md)
- **Used by:** [E Dominant 7th Chord](/chords/e-dominant-7th-chord.md), [E Major 7th Chord](/chords/e-major-7th-chord.md), [E 6th Chord](/chords/e-6th-chord.md), [E Suspended 2 Chord](/chords/e-suspended-2-chord.md), [E Suspended 4 Chord](/chords/e-suspended-4-chord.md)
