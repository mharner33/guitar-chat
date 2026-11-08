---
type: Concept
slug: /triads/e-flat-minor-triad
title: Eb Minor Triad
description: 'The minor triad built on Eb: Eb, Gb, Bb.'
tags:
- triad
- minor
- e
parent: /triads
prerequisites:
- /scales/e-flat-natural-minor-scale
used_by:
- /chords/e-flat-minor-7th-chord
- /chords/e-flat-minor-major-7th-chord
- /chords/e-flat-minor-6th-chord
theory:
  category: triad
  quality: minor
  root:
    pc: 3
    name: Eb
  intervals:
  - '1'
  - b3
  - '5'
  semitones:
  - 0
  - 3
  - 7
  pitch_classes:
  - 3
  - 6
  - 10
  spelling:
  - Eb
  - Gb
  - Bb
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
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  preferred_positions:
  - id: e-flat-minor-voicing-1
    start_fret: 1
    end_fret: 4
    notes:
    - string: 4
      fret: 1
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '5'
      role: fifth
      finger: 3
    - string: 2
      fret: 4
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 4
    - string: 1
      fret: 2
      pc: 6
      note: Gb
      interval: b3
      role: third
      finger: 2
    mute_strings:
    - 5
    - 6
    confidence: 0.75
    tags:
    - computed
  - id: e-flat-minor-voicing-2
    start_fret: 3
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
      fret: 4
      pc: 6
      note: Gb
      interval: b3
      role: third
      finger: 2
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '5'
      role: fifth
      finger: 1
    - string: 2
      fret: 4
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 2
    - string: 1
      fret: 6
      pc: 10
      note: Bb
      interval: '5'
      role: fifth
      finger: 3
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: e-flat-minor-voicing-3
    start_fret: 2
    end_fret: 6
    notes:
    - string: 5
      fret: 6
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 4
    - string: 4
      fret: 4
      pc: 6
      note: Gb
      interval: b3
      role: third
      finger: 3
    - string: 3
      fret: 3
      pc: 10
      note: Bb
      interval: '5'
      role: fifth
      finger: 2
    - string: 2
      fret: 4
      pc: 3
      note: Eb
      interval: '1'
      role: root
      finger: 3
    - string: 1
      fret: 2
      pc: 6
      note: Gb
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
  alphatex: \tempo 90 (1.4 3.3 4.2 2.1).2
  midi_root: 63
musical_forms:
- form: pitch_set
  id: e-flat-minor-triad-pitch-set
  pitch_classes:
  - 3
  - 6
  - 10
  intervals:
  - '1'
  - b3
  - '5'
  spelling:
  - Eb
  - Gb
  - Bb
- form: voicing
  id: e-flat-minor-voicing-1
  position_ref: e-flat-minor-voicing-1
  start_fret: 1
  tab: x-x-1-3-4-2
- form: voicing
  id: e-flat-minor-voicing-2
  position_ref: e-flat-minor-voicing-2
  start_fret: 3
  tab: x-6-4-3-4-6
- form: voicing
  id: e-flat-minor-voicing-3
  position_ref: e-flat-minor-voicing-3
  start_fret: 2
  tab: x-6-4-3-4-2
- form: tab_excerpt
  id: e-flat-minor-triad-tab
  alphatex: \tempo 90 (1.4 3.3 4.2 2.1).2
---

# Eb Minor Triad

The Eb minor triad is a darker but still consonant chord. Its tones are **Eb – Gb – Bb**.

**Construction:** stacked thirds — a minor third (Eb–Gb) under a major third (Gb–Bb).

| Interval | Note | Name |
|---|---|---|
| 1 | Eb | root |
| b3 | Gb | minor third |
| 5 | Bb | perfect fifth |

**Playing it:** the primary shape is `x-x-1-3-4-2` (strings low→high, `x` = muted), around fret 1. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [Eb Natural Minor Scale](/scales/e-flat-natural-minor-scale.md)
- **Used by:** [Eb Minor 7th Chord](/chords/e-flat-minor-7th-chord.md), [Eb Minor Major 7th Chord](/chords/e-flat-minor-major-7th-chord.md), [Eb Minor 6th Chord](/chords/e-flat-minor-6th-chord.md)
