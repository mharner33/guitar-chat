---
type: Concept
slug: /chords/d-flat-13th-chord
title: Db 13th Chord
description: 'The 13th chord built on Db: Db, F, Ab, Cb, Eb, Bb.'
tags:
- chord
- 13th
- d
parent: /chords
prerequisites:
- /chords/d-flat-dominant-7th-chord
theory:
  category: chord
  quality: 13th
  root:
    pc: 1
    name: Db
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '13'
  semitones:
  - 0
  - 4
  - 7
  - 10
  - 14
  - 21
  pitch_classes:
  - 1
  - 5
  - 8
  - 11
  - 3
  - 10
  spelling:
  - Db
  - F
  - Ab
  - Cb
  - Eb
  - Bb
  enharmonic_policy: prefer_flats
  key_context: Db
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
  - 1
  - 5
  - 8
  - 11
  - 3
  - 10
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '13'
  preferred_positions:
  - id: d-flat-13th-voicing-1
    start_fret: 7
    end_fret: 9
    notes:
    - string: 6
      fret: 9
      pc: 1
      note: Db
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 8
      pc: 5
      note: F
      interval: '3'
      role: third
      finger: 2
    - string: 4
      fret: 8
      pc: 10
      note: Bb
      interval: '13'
      role: thirteenth
      finger: 2
    - string: 3
      fret: 8
      pc: 3
      note: Eb
      interval: '9'
      role: ninth
      finger: 2
    - string: 2
      fret: 9
      pc: 8
      note: Ab
      interval: '5'
      role: fifth
      finger: 3
    - string: 1
      fret: 7
      pc: 11
      note: Cb
      interval: b7
      role: seventh
      finger: 1
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: d-flat-13th-voicing-2
    start_fret: 6
    end_fret: 9
    notes:
    - string: 6
      fret: 9
      pc: 1
      note: Db
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 8
      pc: 5
      note: F
      interval: '3'
      role: third
      finger: 2
    - string: 4
      fret: 9
      pc: 11
      note: Cb
      interval: b7
      role: seventh
      finger: 3
    - string: 3
      fret: 8
      pc: 3
      note: Eb
      interval: '9'
      role: ninth
      finger: 2
    - string: 2
      fret: 9
      pc: 8
      note: Ab
      interval: '5'
      role: fifth
      finger: 3
    - string: 1
      fret: 6
      pc: 10
      note: Bb
      interval: '13'
      role: thirteenth
      finger: 1
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: d-flat-13th-voicing-3
    start_fret: 1
    end_fret: 4
    notes:
    - string: 5
      fret: 4
      pc: 1
      note: Db
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 1
      pc: 3
      note: Eb
      interval: '9'
      role: ninth
      finger: 1
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '13'
      role: thirteenth
      finger: 2
    - string: 2
      fret: 0
      pc: 11
      note: Cb
      interval: b7
      role: seventh
      finger: 0
    - string: 1
      fret: 1
      pc: 5
      note: F
      interval: '3'
      role: third
      finger: 1
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - open
    - barre
playback:
  alphatex: \tempo 90 (9.6 8.5 8.4 8.3 9.2 7.1).2
  midi_root: 61
musical_forms:
- form: pitch_set
  id: d-flat-13th-chord-pitch-set
  pitch_classes:
  - 1
  - 5
  - 8
  - 11
  - 3
  - 10
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '13'
  spelling:
  - Db
  - F
  - Ab
  - Cb
  - Eb
  - Bb
- form: voicing
  id: d-flat-13th-voicing-1
  position_ref: d-flat-13th-voicing-1
  start_fret: 7
  tab: 9-8-8-8-9-7
- form: voicing
  id: d-flat-13th-voicing-2
  position_ref: d-flat-13th-voicing-2
  start_fret: 6
  tab: 9-8-9-8-9-6
- form: voicing
  id: d-flat-13th-voicing-3
  position_ref: d-flat-13th-voicing-3
  start_fret: 1
  tab: x-4-1-3-0-1
- form: tab_excerpt
  id: d-flat-13th-chord-tab
  alphatex: \tempo 90 (9.6 8.5 8.4 8.3 9.2 7.1).2
---

# Db 13th Chord

The Db 13th chord is a dominant chord extended through the thirteenth — the fullest dominant color. Its tones are **Db – F – Ab – Cb – Eb – Bb**: Db (root), F (major third), Ab (perfect fifth), Cb (minor seventh), Eb (ninth), Bb (thirteenth).

| Interval | Note | Name |
|---|---|---|
| 1 | Db | root |
| 3 | F | major third |
| 5 | Ab | perfect fifth |
| b7 | Cb | minor seventh |
| 9 | Eb | ninth |
| 13 | Bb | thirteenth |

**Playing it:** the primary shape is `9-8-8-8-9-7` (strings low→high, `x` = muted), around fret 7. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Db Dominant 7th Chord](/chords/d-flat-dominant-7th-chord.md)
