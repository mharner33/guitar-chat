"""Music theory engine for lesson building.

Deterministic note/fret computation for standard tuning. String numbering is
conventional: 1 = high E (thinnest), 6 = low E (thickest).

Position shapes come from fixtures.py (independently checked reference
diagrams), NOT from a fret-window heuristic, so the generator cannot
invent geometry that disagrees with trusted shapes.
"""
from __future__ import annotations

from fixtures import POSITIONS, fixture_notes

CHROMATIC = ["A", "A#/Bb", "B", "C", "C#/Db", "D", "D#/Eb", "E", "F", "F#/Gb", "G", "G#/Ab"]

STANDARD_TUNING = {1: "E", 2: "B", 3: "G", 4: "D", 5: "A", 6: "E"}

INTERVAL_SEMITONES = {
    "1": 0, "b3": 3, "4": 5, "5": 7, "b7": 10,
    "2": 2, "3": 4, "b6": 8, "6": 9, "7": 11, "b2": 1, "b5": 6,
}

MINOR_PENTATONIC = ["1", "b3", "4", "5", "b7"]


def pc_index(pc: str) -> int:
    return CHROMATIC.index(pc)


def note_at(string: int, fret: int) -> str:
    """Return pitch class of the given string (1=high E ... 6=low E) at fret."""
    return CHROMATIC[(pc_index(STANDARD_TUNING[string]) + fret) % 12]


def scale_notes(root: str, intervals: list[str]) -> list[str]:
    root_idx = pc_index(root)
    return [CHROMATIC[(root_idx + INTERVAL_SEMITONES[i]) % 12] for i in intervals]


def transpose_fretboard_shape(shapes: list[dict], semitones: int) -> list[dict]:
    """Shift every note in a list of {string,fret,is_root,...} by N frets."""
    out = []
    for n in shapes:
        out.append({**n, "fret": n["fret"] + semitones})
    return out


def minor_pentatonic_positions_for_root(root: str) -> list[dict]:
    """Return the five reference positions transposed so the root is `root`.

    The fixtures are defined for A minor pentatonic (root A on string 6 fret 5).
    We compute the shift needed to move the root to the requested pitch class
    on string 6 within frets 0-11, and shift all fixture coordinates by that.
    """
    a_root_fret = 5  # A on low E string
    target_fret = None
    for f in range(0, 12):
        if note_at(6, f) == root:
            target_fret = f
            break
    if target_fret is None:
        raise ValueError(f"root {root} not found on string 6 within frets 0-11")

    shift = target_fret - a_root_fret
    positions = []
    for pos in POSITIONS:
        notes = []
        for s, frets in pos["strings"].items():
            for f in frets:
                pc = note_at(s, f + shift)
                is_root = pc == root
                notes.append({
                    "string": s,
                    "fret": f + shift,
                    "pitch_class": pc,
                    "is_root": is_root,
                })
        low = min(n["fret"] for n in notes)
        high = max(n["fret"] for n in notes)
        positions.append({
            "position": pos["name"],
            "fret_range": [low, high],
            "notes": notes,
        })
    return positions
