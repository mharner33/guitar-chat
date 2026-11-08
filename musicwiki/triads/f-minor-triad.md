---
type: Concept
slug: /triads/f-minor-triad
title: F Minor Triad
description: 'The minor triad built on F: F, Ab, C.'
tags:
- triad
- minor
- f
parent: /triads
prerequisites:
- /scales/f-natural-minor-scale
used_by:
- /chords/f-minor-7th-chord
- /chords/f-minor-major-7th-chord
- /chords/f-minor-6th-chord
theory:
  category: triad
  quality: minor
  root:
    pc: 5
    name: F
  intervals:
  - '1'
  - b3
  - '5'
  semitones:
  - 0
  - 3
  - 7
  pitch_classes:
  - 5
  - 8
  - 0
  spelling:
  - F
  - Ab
  - C
  enharmonic_policy: prefer_flats
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
  - 8
  - 0
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  preferred_positions:
  - id: f-minor-voicing-1
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
      fret: 3
      pc: 0
      note: C
      interval: '5'
      role: fifth
      finger: 2
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 2
    - string: 3
      fret: 1
      pc: 8
      note: Ab
      interval: b3
      role: third
      finger: 1
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
    - barre
  - id: f-minor-voicing-2
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
      finger: 3
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 4
    - string: 1
      fret: 4
      pc: 8
      note: Ab
      interval: b3
      role: third
      finger: 2
    mute_strings:
    - 5
    - 6
    confidence: 0.75
    tags:
    - computed
  - id: f-minor-voicing-3
    start_fret: 5
    end_fret: 8
    notes:
    - string: 5
      fret: 8
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 6
      pc: 8
      note: Ab
      interval: b3
      role: third
      finger: 2
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
      fret: 8
      pc: 0
      note: C
      interval: '5'
      role: fifth
      finger: 3
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (1.6 3.5 3.4 1.3 1.2 1.1).2
  midi_root: 65
musical_forms:
- form: pitch_set
  id: f-minor-triad-pitch-set
  pitch_classes:
  - 5
  - 8
  - 0
  intervals:
  - '1'
  - b3
  - '5'
  spelling:
  - F
  - Ab
  - C
- form: voicing
  id: f-minor-voicing-1
  position_ref: f-minor-voicing-1
  start_fret: 1
  tab: 1-3-3-1-1-1
- form: voicing
  id: f-minor-voicing-2
  position_ref: f-minor-voicing-2
  start_fret: 3
  tab: x-x-3-5-6-4
- form: voicing
  id: f-minor-voicing-3
  position_ref: f-minor-voicing-3
  start_fret: 5
  tab: x-8-6-5-6-8
- form: tab_excerpt
  id: f-minor-triad-tab
  alphatex: \tempo 90 (1.6 3.5 3.4 1.3 1.2 1.1).2
---

# F Minor Triad

The F minor triad is a darker but still consonant chord. Its tones are **F – Ab – C**.

**Construction:** stacked thirds — a minor third (F–Ab) under a major third (Ab–C).

| Interval | Note | Name |
|---|---|---|
| 1 | F | root |
| b3 | Ab | minor third |
| 5 | C | perfect fifth |

**Playing it:** the primary shape is `1-3-3-1-1-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [F Natural Minor Scale](/scales/f-natural-minor-scale.md)
- **Used by:** [F Minor 7th Chord](/chords/f-minor-7th-chord.md), [F Minor Major 7th Chord](/chords/f-minor-major-7th-chord.md), [F Minor 6th Chord](/chords/f-minor-6th-chord.md)
