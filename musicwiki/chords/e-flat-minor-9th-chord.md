---
type: Concept
slug: /chords/e-flat-minor-9th-chord
title: Eb Minor 9th Chord
description: 'The minor 9th chord built on Eb: Eb, Gb, Bb, Db, F.'
tags:
- chord
- minor-9th
- e
parent: /chords
prerequisites:
- /chords/e-flat-minor-7th-chord
theory:
  category: chord
  quality: minor-9th
  root:
    pc: 3
    name: Eb
  intervals:
  - '1'
  - b3
  - '5'
  - b7
  - '9'
  semitones:
  - 0
  - 3
  - 7
  - 10
  - 14
  pitch_classes:
  - 3
  - 6
  - 10
  - 1
  - 5
  spelling:
  - Eb
  - Gb
  - Bb
  - Db
  - F
  enharmonic_policy: prefer_flats
  key_context: Eb
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
  - 3
  - 6
  - 10
  - 1
  - 5
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  - b7
  - '9'
  preferred_positions:
  - id: e-flat-minor-9th-voicing-1
    start_fret: 4
    end_fret: 6
    notes:
    - string: 5
      fret: 6
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 4
      pc: 6
      note: Gb
      interval: b3
      role: third
      finger: 1
    - string: 3
      fret: 6
      pc: 1
      note: Db
      interval: b7
      role: seventh
      finger: 2
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '9'
      role: ninth
      finger: 2
    - string: 1
      fret: 6
      pc: 10
      note: Bb
      interval: '5'
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: e-flat-minor-9th-voicing-2
    start_fret: 2
    end_fret: 6
    notes:
    - string: 5
      fret: 6
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '9'
      role: ninth
      finger: 2
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '5'
      role: fifth
      finger: 2
    - string: 2
      fret: 2
      pc: 1
      note: Db
      interval: b7
      role: seventh
      finger: 1
    - string: 1
      fret: 2
      pc: 6
      note: Gb
      interval: b3
      role: third
      finger: 1
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: e-flat-minor-9th-voicing-3
    start_fret: 3
    end_fret: 7
    notes:
    - string: 5
      fret: 6
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 3
      pc: 5
      note: F
      interval: '9'
      role: ninth
      finger: 1
    - string: 3
      fret: 6
      pc: 1
      note: Db
      interval: b7
      role: seventh
      finger: 2
    - string: 2
      fret: 7
      pc: 6
      note: Gb
      interval: b3
      role: third
      finger: 3
    - string: 1
      fret: 6
      pc: 10
      note: Bb
      interval: '5'
      role: fifth
      finger: 2
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (6.5 4.4 6.3 6.2 6.1).2
  midi_root: 63
musical_forms:
- form: pitch_set
  id: e-flat-minor-9th-chord-pitch-set
  pitch_classes:
  - 3
  - 6
  - 10
  - 1
  - 5
  intervals:
  - '1'
  - b3
  - '5'
  - b7
  - '9'
  spelling:
  - Eb
  - Gb
  - Bb
  - Db
  - F
- form: voicing
  id: e-flat-minor-9th-voicing-1
  position_ref: e-flat-minor-9th-voicing-1
  start_fret: 4
  tab: x-6-4-6-6-6
- form: voicing
  id: e-flat-minor-9th-voicing-2
  position_ref: e-flat-minor-9th-voicing-2
  start_fret: 2
  tab: x-6-3-3-2-2
- form: voicing
  id: e-flat-minor-9th-voicing-3
  position_ref: e-flat-minor-9th-voicing-3
  start_fret: 3
  tab: x-6-3-6-7-6
- form: tab_excerpt
  id: e-flat-minor-9th-chord-tab
  alphatex: \tempo 90 (6.5 4.4 6.3 6.2 6.1).2
---

# Eb Minor 9th Chord

The Eb minor 9th chord is a minor seventh extended with the ninth — smooth and modern. Its tones are **Eb – Gb – Bb – Db – F**: Eb (root), Gb (minor third), Bb (perfect fifth), Db (minor seventh), F (ninth).

| Interval | Note | Name |
|---|---|---|
| 1 | Eb | root |
| b3 | Gb | minor third |
| 5 | Bb | perfect fifth |
| b7 | Db | minor seventh |
| 9 | F | ninth |

**Playing it:** the primary shape is `x-6-4-6-6-6` (strings low→high, `x` = muted), around fret 4. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Eb Minor 7th Chord](/chords/e-flat-minor-7th-chord.md)
