---
type: Concept
slug: /triads/e-minor-triad
title: E Minor Triad
description: 'The minor triad built on E: E, G, B.'
tags:
- triad
- minor
- e
parent: /triads
prerequisites:
- /scales/e-natural-minor-scale
used_by:
- /chords/e-minor-7th-chord
- /chords/e-minor-major-7th-chord
- /chords/e-minor-6th-chord
theory:
  category: triad
  quality: minor
  root:
    pc: 4
    name: E
  intervals:
  - '1'
  - b3
  - '5'
  semitones:
  - 0
  - 3
  - 7
  pitch_classes:
  - 4
  - 7
  - 11
  spelling:
  - E
  - G
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
  - 7
  - 11
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  preferred_positions:
  - id: e-minor-voicing-1
    start_fret: 2
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
      finger: 1
    - string: 4
      fret: 2
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 0
      pc: 7
      note: G
      interval: b3
      role: third
      finger: 0
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
  - id: e-minor-voicing-2
    start_fret: 4
    end_fret: 7
    notes:
    - string: 5
      fret: 7
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: b3
      role: third
      finger: 2
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
      fret: 7
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 3
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: e-minor-voicing-3
    start_fret: 3
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
      fret: 5
      pc: 7
      note: G
      interval: b3
      role: third
      finger: 3
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 2
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 3
    - string: 1
      fret: 3
      pc: 7
      note: G
      interval: b3
      role: third
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (0.6 2.5 2.4 0.3 0.2 0.1).2
  midi_root: 64
musical_forms:
- form: pitch_set
  id: e-minor-triad-pitch-set
  pitch_classes:
  - 4
  - 7
  - 11
  intervals:
  - '1'
  - b3
  - '5'
  spelling:
  - E
  - G
  - B
- form: voicing
  id: e-minor-voicing-1
  position_ref: e-minor-voicing-1
  start_fret: 2
  tab: 0-2-2-0-0-0
- form: voicing
  id: e-minor-voicing-2
  position_ref: e-minor-voicing-2
  start_fret: 4
  tab: x-7-5-4-5-7
- form: voicing
  id: e-minor-voicing-3
  position_ref: e-minor-voicing-3
  start_fret: 3
  tab: x-7-5-4-5-3
- form: tab_excerpt
  id: e-minor-triad-tab
  alphatex: \tempo 90 (0.6 2.5 2.4 0.3 0.2 0.1).2
---

# E Minor Triad

The E minor triad is a darker but still consonant chord. Its tones are **E – G – B**.

**Construction:** stacked thirds — a minor third (E–G) under a major third (G–B).

| Interval | Note | Name |
|---|---|---|
| 1 | E | root |
| b3 | G | minor third |
| 5 | B | perfect fifth |

**Playing it:** the primary shape is `0-2-2-0-0-0` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [E Natural Minor Scale](/scales/e-natural-minor-scale.md)
- **Used by:** [E Minor 7th Chord](/chords/e-minor-7th-chord.md), [E Minor Major 7th Chord](/chords/e-minor-major-7th-chord.md), [E Minor 6th Chord](/chords/e-minor-6th-chord.md)
