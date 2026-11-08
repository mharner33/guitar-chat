---
type: Concept
slug: /triads/c-minor-triad
title: C Minor Triad
description: 'The minor triad built on C: C, Eb, G.'
tags:
- triad
- minor
- c
parent: /triads
prerequisites:
- /scales/c-natural-minor-scale
used_by:
- /chords/c-minor-7th-chord
- /chords/c-minor-major-7th-chord
- /chords/c-minor-6th-chord
theory:
  category: triad
  quality: minor
  root:
    pc: 0
    name: C
  intervals:
  - '1'
  - b3
  - '5'
  semitones:
  - 0
  - 3
  - 7
  pitch_classes:
  - 0
  - 3
  - 7
  spelling:
  - C
  - Eb
  - G
  enharmonic_policy: prefer_flats
  key_context: C
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
  - 0
  - 3
  - 7
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  preferred_positions:
  - id: c-minor-voicing-1
    start_fret: 1
    end_fret: 3
    notes:
    - string: 5
      fret: 3
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 1
      pc: 3
      note: Eb
      interval: b3
      role: third
      finger: 1
    - string: 3
      fret: 0
      pc: 7
      note: G
      interval: '5'
      role: fifth
      finger: 0
    - string: 2
      fret: 1
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 1
    - string: 1
      fret: 3
      pc: 7
      note: G
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
  - id: c-minor-voicing-2
    start_fret: 3
    end_fret: 5
    notes:
    - string: 5
      fret: 3
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: '5'
      role: fifth
      finger: 3
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 3
    - string: 2
      fret: 4
      pc: 3
      note: Eb
      interval: b3
      role: third
      finger: 2
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: '5'
      role: fifth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: c-minor-voicing-3
    start_fret: 5
    end_fret: 8
    notes:
    - string: 6
      fret: 8
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 6
      pc: 3
      note: Eb
      interval: b3
      role: third
      finger: 2
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: '5'
      role: fifth
      finger: 1
    - string: 3
      fret: 5
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 1
    - string: 2
      fret: 8
      pc: 7
      note: G
      interval: '5'
      role: fifth
      finger: 3
    - string: 1
      fret: 8
      pc: 0
      note: C
      interval: '1'
      role: root
      finger: 3
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (3.5 1.4 0.3 1.2 3.1).2
  midi_root: 60
musical_forms:
- form: pitch_set
  id: c-minor-triad-pitch-set
  pitch_classes:
  - 0
  - 3
  - 7
  intervals:
  - '1'
  - b3
  - '5'
  spelling:
  - C
  - Eb
  - G
- form: voicing
  id: c-minor-voicing-1
  position_ref: c-minor-voicing-1
  start_fret: 1
  tab: x-3-1-0-1-3
- form: voicing
  id: c-minor-voicing-2
  position_ref: c-minor-voicing-2
  start_fret: 3
  tab: x-3-5-5-4-3
- form: voicing
  id: c-minor-voicing-3
  position_ref: c-minor-voicing-3
  start_fret: 5
  tab: 8-6-5-5-8-8
- form: tab_excerpt
  id: c-minor-triad-tab
  alphatex: \tempo 90 (3.5 1.4 0.3 1.2 3.1).2
---

# C Minor Triad

The C minor triad is a darker but still consonant chord. Its tones are **C – Eb – G**.

**Construction:** stacked thirds — a minor third (C–Eb) under a major third (Eb–G).

| Interval | Note | Name |
|---|---|---|
| 1 | C | root |
| b3 | Eb | minor third |
| 5 | G | perfect fifth |

**Playing it:** the primary shape is `x-3-1-0-1-3` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [C Natural Minor Scale](/scales/c-natural-minor-scale.md)
- **Used by:** [C Minor 7th Chord](/chords/c-minor-7th-chord.md), [C Minor Major 7th Chord](/chords/c-minor-major-7th-chord.md), [C Minor 6th Chord](/chords/c-minor-6th-chord.md)
