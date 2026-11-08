---
type: Concept
slug: /chords/e-flat-11th-chord
title: Eb 11th Chord
description: 'The 11th chord built on Eb: Eb, G, Bb, Db, F, Ab.'
tags:
- chord
- 11th
- e
parent: /chords
prerequisites:
- /chords/e-flat-dominant-7th-chord
theory:
  category: chord
  quality: 11th
  root:
    pc: 3
    name: Eb
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '11'
  semitones:
  - 0
  - 4
  - 7
  - 10
  - 14
  - 17
  pitch_classes:
  - 3
  - 7
  - 10
  - 1
  - 5
  - 8
  spelling:
  - Eb
  - G
  - Bb
  - Db
  - F
  - Ab
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
  - 7
  - 10
  - 1
  - 5
  - 8
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '11'
  preferred_positions:
  - id: e-flat-11th-voicing-1
    start_fret: 4
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
      fret: 5
      pc: 7
      note: G
      interval: '3'
      role: third
      finger: 2
    - string: 3
      fret: 6
      pc: 1
      note: Db
      interval: b7
      role: seventh
      finger: 3
    - string: 2
      fret: 6
      pc: 5
      note: F
      interval: '9'
      role: ninth
      finger: 3
    - string: 1
      fret: 4
      pc: 8
      note: Ab
      interval: '11'
      role: eleventh
      finger: 1
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: e-flat-11th-voicing-2
    start_fret: 8
    end_fret: 11
    notes:
    - string: 6
      fret: 11
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 4
    - string: 5
      fret: 10
      pc: 7
      note: G
      interval: '3'
      role: third
      finger: 3
    - string: 4
      fret: 8
      pc: 10
      note: Bb
      interval: '5'
      role: fifth
      finger: 1
    - string: 3
      fret: 10
      pc: 5
      note: F
      interval: '9'
      role: ninth
      finger: 3
    - string: 2
      fret: 9
      pc: 8
      note: Ab
      interval: '11'
      role: eleventh
      finger: 2
    - string: 1
      fret: 9
      pc: 1
      note: Db
      interval: b7
      role: seventh
      finger: 2
    confidence: 0.85
    tags:
    - computed
    - barre
  - id: e-flat-11th-voicing-3
    start_fret: 3
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
      fret: 6
      pc: 8
      note: Ab
      interval: '11'
      role: eleventh
      finger: 2
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
      fret: 3
      pc: 7
      note: G
      interval: '3'
      role: third
      finger: 1
    mute_strings:
    - 6
    confidence: 0.85
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (6.5 5.4 6.3 6.2 4.1).2
  midi_root: 63
musical_forms:
- form: pitch_set
  id: e-flat-11th-chord-pitch-set
  pitch_classes:
  - 3
  - 7
  - 10
  - 1
  - 5
  - 8
  intervals:
  - '1'
  - '3'
  - '5'
  - b7
  - '9'
  - '11'
  spelling:
  - Eb
  - G
  - Bb
  - Db
  - F
  - Ab
- form: voicing
  id: e-flat-11th-voicing-1
  position_ref: e-flat-11th-voicing-1
  start_fret: 4
  tab: x-6-5-6-6-4
- form: voicing
  id: e-flat-11th-voicing-2
  position_ref: e-flat-11th-voicing-2
  start_fret: 8
  tab: 11-10-8-10-9-9
- form: voicing
  id: e-flat-11th-voicing-3
  position_ref: e-flat-11th-voicing-3
  start_fret: 3
  tab: x-6-6-6-6-3
- form: tab_excerpt
  id: e-flat-11th-chord-tab
  alphatex: \tempo 90 (6.5 5.4 6.3 6.2 4.1).2
---

# Eb 11th Chord

The Eb 11th chord is a dominant chord extended through the eleventh. Its tones are **Eb – G – Bb – Db – F – Ab**: Eb (root), G (major third), Bb (perfect fifth), Db (minor seventh), F (ninth), Ab (eleventh).

| Interval | Note | Name |
|---|---|---|
| 1 | Eb | root |
| 3 | G | major third |
| 5 | Bb | perfect fifth |
| b7 | Db | minor seventh |
| 9 | F | ninth |
| 11 | Ab | eleventh |

**Playing it:** the primary shape is `x-6-5-6-6-4` (strings low→high, `x` = muted), around fret 4. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Eb Dominant 7th Chord](/chords/e-flat-dominant-7th-chord.md)
