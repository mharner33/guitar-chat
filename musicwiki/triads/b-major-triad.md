---
type: Concept
slug: /triads/b-major-triad
title: B Major Triad
description: 'The major triad built on B: B, D#, F#.'
tags:
- triad
- major
- b
parent: /triads
prerequisites:
- /scales/b-major-scale
used_by:
- /chords/b-dominant-7th-chord
- /chords/b-major-7th-chord
- /chords/b-6th-chord
- /chords/b-suspended-2-chord
- /chords/b-suspended-4-chord
theory:
  category: triad
  quality: major
  root:
    pc: 11
    name: B
  intervals:
  - '1'
  - '3'
  - '5'
  semitones:
  - 0
  - 4
  - 7
  pitch_classes:
  - 11
  - 3
  - 6
  spelling:
  - B
  - D#
  - F#
  enharmonic_policy: prefer_sharps
  key_context: B
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
  - 11
  - 3
  - 6
  emphasize_intervals:
  - '1'
  - '3'
  - '5'
  preferred_positions:
  - id: b-major-voicing-1
    start_fret: 2
    end_fret: 4
    notes:
    - string: 5
      fret: 2
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 1
    - string: 4
      fret: 4
      pc: 6
      note: F#
      interval: '5'
      role: fifth
      finger: 2
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 2
    - string: 2
      fret: 4
      pc: 3
      note: D#
      interval: '3'
      role: third
      finger: 2
    - string: 1
      fret: 2
      pc: 6
      note: F#
      interval: '5'
      role: fifth
      finger: 1
    mute_strings:
    - 6
    confidence: 0.75
    tags:
    - computed
    - barre
  - id: b-major-voicing-2
    start_fret: 1
    end_fret: 4
    notes:
    - string: 5
      fret: 2
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 2
    - string: 4
      fret: 1
      pc: 3
      note: D#
      interval: '3'
      role: third
      finger: 1
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 3
    - string: 2
      fret: 0
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 0
    - string: 1
      fret: 2
      pc: 6
      note: F#
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
  - id: b-major-voicing-3
    start_fret: 4
    end_fret: 7
    notes:
    - string: 6
      fret: 7
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 3
    - string: 5
      fret: 6
      pc: 3
      note: D#
      interval: '3'
      role: third
      finger: 2
    - string: 4
      fret: 4
      pc: 6
      note: F#
      interval: '5'
      role: fifth
      finger: 1
    - string: 3
      fret: 4
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 1
    - string: 2
      fret: 4
      pc: 3
      note: D#
      interval: '3'
      role: third
      finger: 1
    - string: 1
      fret: 7
      pc: 11
      note: B
      interval: '1'
      role: root
      finger: 3
    confidence: 0.75
    tags:
    - computed
    - barre
playback:
  alphatex: \tempo 90 (2.5 4.4 4.3 4.2 2.1).2
  midi_root: 71
musical_forms:
- form: pitch_set
  id: b-major-triad-pitch-set
  pitch_classes:
  - 11
  - 3
  - 6
  intervals:
  - '1'
  - '3'
  - '5'
  spelling:
  - B
  - D#
  - F#
- form: voicing
  id: b-major-voicing-1
  position_ref: b-major-voicing-1
  start_fret: 2
  tab: x-2-4-4-4-2
- form: voicing
  id: b-major-voicing-2
  position_ref: b-major-voicing-2
  start_fret: 1
  tab: x-2-1-4-0-2
- form: voicing
  id: b-major-voicing-3
  position_ref: b-major-voicing-3
  start_fret: 4
  tab: 7-6-4-4-4-7
- form: tab_excerpt
  id: b-major-triad-tab
  alphatex: \tempo 90 (2.5 4.4 4.3 4.2 2.1).2
---

# B Major Triad

The B major triad is a bright, stable, consonant chord. Its tones are **B – D# – F#**.

**Construction:** stacked thirds — a major third (B–D#) under a minor third (D#–F#).

| Interval | Note | Name |
|---|---|---|
| 1 | B | root |
| 3 | D# | major third |
| 5 | F# | perfect fifth |

**Playing it:** the primary shape is `x-2-4-4-4-2` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [B Major Scale](/scales/b-major-scale.md)
- **Used by:** [B Dominant 7th Chord](/chords/b-dominant-7th-chord.md), [B Major 7th Chord](/chords/b-major-7th-chord.md), [B 6th Chord](/chords/b-6th-chord.md), [B Suspended 2 Chord](/chords/b-suspended-2-chord.md), [B Suspended 4 Chord](/chords/b-suspended-4-chord.md)
