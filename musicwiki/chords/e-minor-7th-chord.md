---
type: Concept
slug: /chords/e-minor-7th-chord
title: E Minor 7th Chord
description: 'The minor 7th chord built on E: E, G, B, D.'
tags:
- chord
- minor-7th
- e
parent: /chords
prerequisites:
- /triads/e-minor-triad
used_by:
- /chords/e-minor-9th-chord
theory:
  category: chord
  quality: minor-7th
  root:
    pc: 4
    name: E
  intervals:
  - '1'
  - b3
  - '5'
  - b7
  semitones:
  - 0
  - 3
  - 7
  - 10
  pitch_classes:
  - 4
  - 7
  - 11
  - 2
  spelling:
  - E
  - G
  - B
  - D
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
  - 2
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  - b7
  preferred_positions:
  - id: e-minor-7th-voicing-1
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
      fret: 0
      pc: 2
      note: D
      interval: b7
      role: seventh
      finger: 0
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
  - id: e-minor-7th-voicing-2
    start_fret: 0
    end_fret: 0
    notes:
    - string: 6
      fret: 0
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 0
    - string: 4
      fret: 0
      pc: 2
      note: D
      interval: b7
      role: seventh
      finger: 0
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
    mute_strings:
    - 5
    confidence: 0.68
    tags:
    - computed
    - open
  - id: e-minor-7th-voicing-3
    start_fret: 5
    end_fret: 7
    notes:
    - string: 5
      fret: 7
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 5
      pc: 7
      note: G
      interval: b3
      role: third
      finger: 1
    - string: 3
      fret: 7
      pc: 2
      note: D
      interval: b7
      role: seventh
      finger: 2
    - string: 2
      fret: 5
      pc: 4
      note: E
      interval: '1'
      role: root
      finger: 1
    - string: 1
      fret: 7
      pc: 11
      note: B
      interval: '5'
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (0.6 2.5 0.4 0.3 0.2 0.1).2
  midi_root: 64
musical_forms:
- form: pitch_set
  id: e-minor-7th-chord-pitch-set
  pitch_classes:
  - 4
  - 7
  - 11
  - 2
  intervals:
  - '1'
  - b3
  - '5'
  - b7
  spelling:
  - E
  - G
  - B
  - D
- form: voicing
  id: e-minor-7th-voicing-1
  position_ref: e-minor-7th-voicing-1
  start_fret: 2
  tab: 0-2-0-0-0-0
- form: voicing
  id: e-minor-7th-voicing-2
  position_ref: e-minor-7th-voicing-2
  start_fret: 0
  tab: 0-x-0-0-0-0
- form: voicing
  id: e-minor-7th-voicing-3
  position_ref: e-minor-7th-voicing-3
  start_fret: 5
  tab: x-7-5-7-5-7
- form: tab_excerpt
  id: e-minor-7th-chord-tab
  alphatex: \tempo 90 (0.6 2.5 0.4 0.3 0.2 0.1).2
---

# E Minor 7th Chord

The E minor 7th chord is a minor triad plus a minor seventh — smooth and mellow, the workhorse of jazz and soul. Its tones are **E – G – B – D**: E (root), G (minor third), B (perfect fifth), D (minor seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | E | root |
| b3 | G | minor third |
| 5 | B | perfect fifth |
| b7 | D | minor seventh |

**Playing it:** the primary shape is `0-2-0-0-0-0` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [E Minor Triad](/triads/e-minor-triad.md)
- **Used by:** [E Minor 9th Chord](/chords/e-minor-9th-chord.md)
