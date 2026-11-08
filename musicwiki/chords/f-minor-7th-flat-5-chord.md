---
type: Concept
slug: /chords/f-minor-7th-flat-5-chord
title: F Minor 7th Flat 5 Chord
description: 'The minor 7th flat 5 chord built on F: F, Ab, Cb, Eb.'
tags:
- chord
- minor-7th-flat-5
- f
parent: /chords
prerequisites:
- /triads/f-diminished-triad
theory:
  category: chord
  quality: minor-7th-flat-5
  root:
    pc: 5
    name: F
  intervals:
  - '1'
  - b3
  - b5
  - b7
  semitones:
  - 0
  - 3
  - 6
  - 10
  pitch_classes:
  - 5
  - 8
  - 11
  - 3
  spelling:
  - F
  - Ab
  - Cb
  - Eb
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
  - 11
  - 3
  emphasize_intervals:
  - '1'
  - b3
  - b5
  - b7
  preferred_positions:
  - id: f-minor-7th-flat-5-voicing-1
    start_fret: 1
    end_fret: 2
    notes:
    - string: 6
      fret: 1
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 2
      pc: 11
      note: Cb
      interval: b5
      role: fifth
      finger: 2
    - string: 4
      fret: 1
      pc: 3
      note: Eb
      interval: b7
      role: seventh
      finger: 1
    - string: 3
      fret: 1
      pc: 8
      note: Ab
      interval: b3
      role: third
      finger: 1
    - string: 2
      fret: 0
      pc: 11
      note: Cb
      interval: b5
      role: fifth
      finger: 0
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
  - id: f-minor-7th-flat-5-voicing-2
    start_fret: 3
    end_fret: 4
    notes:
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 4
      pc: 11
      note: Cb
      interval: b5
      role: fifth
      finger: 2
    - string: 2
      fret: 4
      pc: 3
      note: Eb
      interval: b7
      role: seventh
      finger: 2
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
    - barre
  - id: f-minor-7th-flat-5-voicing-3
    start_fret: 6
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
      finger: 1
    - string: 3
      fret: 8
      pc: 3
      note: Eb
      interval: b7
      role: seventh
      finger: 3
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '1'
      role: root
      finger: 1
    - string: 1
      fret: 7
      pc: 11
      note: Cb
      interval: b5
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (1.6 2.5 1.4 1.3 0.2 1.1).2
  midi_root: 65
musical_forms:
- form: pitch_set
  id: f-minor-7th-flat-5-chord-pitch-set
  pitch_classes:
  - 5
  - 8
  - 11
  - 3
  intervals:
  - '1'
  - b3
  - b5
  - b7
  spelling:
  - F
  - Ab
  - Cb
  - Eb
- form: voicing
  id: f-minor-7th-flat-5-voicing-1
  position_ref: f-minor-7th-flat-5-voicing-1
  start_fret: 1
  tab: 1-2-1-1-0-1
- form: voicing
  id: f-minor-7th-flat-5-voicing-2
  position_ref: f-minor-7th-flat-5-voicing-2
  start_fret: 3
  tab: x-x-3-4-4-4
- form: voicing
  id: f-minor-7th-flat-5-voicing-3
  position_ref: f-minor-7th-flat-5-voicing-3
  start_fret: 6
  tab: x-8-6-8-6-7
- form: tab_excerpt
  id: f-minor-7th-flat-5-chord-tab
  alphatex: \tempo 90 (1.6 2.5 1.4 1.3 0.2 1.1).2
---

# F Minor 7th Flat 5 Chord

The F minor 7th flat 5 chord is the half-diminished chord: a diminished triad plus a minor seventh, the ii chord of a minor key. Its tones are **F – Ab – Cb – Eb**: F (root), Ab (minor third), Cb (diminished fifth), Eb (minor seventh).

| Interval | Note | Name |
|---|---|---|
| 1 | F | root |
| b3 | Ab | minor third |
| b5 | Cb | diminished fifth |
| b7 | Eb | minor seventh |

**Playing it:** the primary shape is `1-2-1-1-0-1` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [F Diminished Triad](/triads/f-diminished-triad.md)
