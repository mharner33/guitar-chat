---
type: Concept
slug: /triads/f-sharp-minor-triad
title: F# Minor Triad
description: 'The minor triad built on F#: F#, A, C#.'
tags:
- triad
- minor
- f
parent: /triads
prerequisites:
- /scales/f-sharp-natural-minor-scale
used_by:
- /chords/f-sharp-minor-7th-chord
- /chords/f-sharp-minor-major-7th-chord
- /chords/f-sharp-minor-6th-chord
theory:
  category: triad
  quality: minor
  root:
    pc: 6
    name: F#
  intervals:
  - '1'
  - b3
  - '5'
  semitones:
  - 0
  - 3
  - 7
  pitch_classes:
  - 6
  - 9
  - 1
  spelling:
  - F#
  - A
  - C#
  enharmonic_policy: prefer_sharps
  key_context: F#
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
  - 6
  - 9
  - 1
  emphasize_intervals:
  - '1'
  - b3
  - '5'
  preferred_positions:
  - id: f-sharp-minor-voicing-1
    start_fret: 2
    end_fret: 4
    notes:
    - string: 6
      fret: 2
      pc: 6
      note: F#
      interval: '1'
      role: root
      finger: 1
    - string: 5
      fret: 0
      pc: 9
      note: A
      interval: b3
      role: third
      finger: 0
    - string: 4
      fret: 4
      pc: 6
      note: F#
      interval: '1'
      role: root
      finger: 2
    - string: 3
      fret: 2
      pc: 9
      note: A
      interval: b3
      role: third
      finger: 1
    - string: 2
      fret: 2
      pc: 1
      note: C#
      interval: '5'
      role: fifth
      finger: 1
    - string: 1
      fret: 2
      pc: 6
      note: F#
      interval: '1'
      role: root
      finger: 1
    confidence: 0.75
    tags:
    - computed
    - open
    - barre
  - id: f-sharp-minor-voicing-2
    start_fret: 4
    end_fret: 7
    notes:
    - string: 4
      fret: 4
      pc: 6
      note: F#
      interval: '1'
      role: root
      finger: 1
    - string: 3
      fret: 6
      pc: 1
      note: C#
      interval: '5'
      role: fifth
      finger: 3
    - string: 2
      fret: 7
      pc: 6
      note: F#
      interval: '1'
      role: root
      finger: 4
    - string: 1
      fret: 5
      pc: 9
      note: A
      interval: b3
      role: third
      finger: 2
    mute_strings:
    - 5
    - 6
    confidence: 0.75
    tags:
    - computed
  - id: f-sharp-minor-voicing-3
    start_fret: 6
    end_fret: 9
    notes:
    - string: 5
      fret: 9
      pc: 6
      note: F#
      interval: '1'
      role: root
      finger: 3
    - string: 4
      fret: 7
      pc: 9
      note: A
      interval: b3
      role: third
      finger: 2
    - string: 3
      fret: 6
      pc: 1
      note: C#
      interval: '5'
      role: fifth
      finger: 1
    - string: 2
      fret: 7
      pc: 6
      note: F#
      interval: '1'
      role: root
      finger: 2
    - string: 1
      fret: 9
      pc: 1
      note: C#
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
  alphatex: \tempo 90 (2.6 0.5 4.4 2.3 2.2 2.1).2
  midi_root: 66
musical_forms:
- form: pitch_set
  id: f-sharp-minor-triad-pitch-set
  pitch_classes:
  - 6
  - 9
  - 1
  intervals:
  - '1'
  - b3
  - '5'
  spelling:
  - F#
  - A
  - C#
- form: voicing
  id: f-sharp-minor-voicing-1
  position_ref: f-sharp-minor-voicing-1
  start_fret: 2
  tab: 2-0-4-2-2-2
- form: voicing
  id: f-sharp-minor-voicing-2
  position_ref: f-sharp-minor-voicing-2
  start_fret: 4
  tab: x-x-4-6-7-5
- form: voicing
  id: f-sharp-minor-voicing-3
  position_ref: f-sharp-minor-voicing-3
  start_fret: 6
  tab: x-9-7-6-7-9
- form: tab_excerpt
  id: f-sharp-minor-triad-tab
  alphatex: \tempo 90 (2.6 0.5 4.4 2.3 2.2 2.1).2
---

# F# Minor Triad

The F# minor triad is a darker but still consonant chord. Its tones are **F# – A – C#**.

**Construction:** stacked thirds — a minor third (F#–A) under a major third (A–C#).

| Interval | Note | Name |
|---|---|---|
| 1 | F# | root |
| b3 | A | minor third |
| 5 | C# | perfect fifth |

**Playing it:** the primary shape is `2-0-4-2-2-2` (strings low→high, `x` = muted), around fret 2. 3 computed fingerings are in the render block, with alphaTex playback.

## Related

- **Build on:** [F# Natural Minor Scale](/scales/f-sharp-natural-minor-scale.md)
- **Used by:** [F# Minor 7th Chord](/chords/f-sharp-minor-7th-chord.md), [F# Minor Major 7th Chord](/chords/f-sharp-minor-major-7th-chord.md), [F# Minor 6th Chord](/chords/f-sharp-minor-6th-chord.md)
