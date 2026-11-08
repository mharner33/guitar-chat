"""Deterministic SVG renderer for scale/chord diagrams.

Reads diagram data (from the theory engine) and produces SVG. Never draws
anything not specified by the data; orientation and string/fret conventions
are fixed here.

String numbering is conventional: 1 = high E (thinnest, rightmost column),
6 = low E (thickest, leftmost column). Strings are drawn low-to-high,
left-to-right: columns 6,5,4,3,2,1.
"""
from __future__ import annotations

from pathlib import Path

# SVG layout constants (pixels)
STRING_GAP = 30
FRET_GAP = 36
MARGIN_X = 70
MARGIN_Y = 60
NOTE_R = 11

# Columns drawn left-to-right are strings 6..1 (low to high)
_DRAWN_STRINGS = [6, 5, 4, 3, 2, 1]


def _string_x(string_num: int) -> float:
    """X coordinate for a given string number (1=high E rightmost)."""
    col_index = _DRAWN_STRINGS.index(string_num)
    return MARGIN_X + col_index * STRING_GAP


def render_scale_svg(diagram: dict) -> str:
    notes = diagram["render"]["notes"]
    low = diagram["render"]["frets"]["low"]
    high = diagram["render"]["frets"]["high"]
    n_frets = high - low + 1  # rows in the grid

    width = MARGIN_X * 2 + STRING_GAP * 5
    height = MARGIN_Y * 2 + FRET_GAP * (n_frets - 1) + 40  # extra room for legend

    parts = [
        f'<svg xmlns="http://www.w3.org/2000/svg" width="{width}" height="{height}" viewBox="0 0 {width} {height}">',
        f'<rect width="{width}" height="{height}" fill="#ffffff"/>',
        f'<text x="{width/2}" y="24" text-anchor="middle" font-family="sans-serif" font-size="16" font-weight="600">{diagram["name"]}</text>',
    ]

    # frets (horizontal lines): one per grid boundary
    for r in range(n_frets):
        y = MARGIN_Y + r * FRET_GAP
        parts.append(f'<line x1="{MARGIN_X}" y1="{y}" x2="{MARGIN_X + STRING_GAP*5}" y2="{y}" stroke="#999" stroke-width="1"/>')
    parts.append(f'<line x1="{MARGIN_X}" y1="{MARGIN_Y + FRET_GAP*n_frets}" x2="{MARGIN_X + STRING_GAP*5}" y2="{MARGIN_Y + FRET_GAP*n_frets}" stroke="#999" stroke-width="1"/>')
    # strings (vertical lines): thickness scales with string gauge
    for col, s in enumerate(_DRAWN_STRINGS):
        x = MARGIN_X + col * STRING_GAP
        stroke_w = 1.0 + (s - 1) * 0.3  # string 6 (low E) thickest
        parts.append(f'<line x1="{x}" y1="{MARGIN_Y}" x2="{x}" y2="{MARGIN_Y + FRET_GAP*n_frets}" stroke="#333" stroke-width="{stroke_w:.1f}"/>')

    # fret labels on the left
    for i, f in enumerate(range(low, high + 1)):
        y = MARGIN_Y + i * FRET_GAP + FRET_GAP / 2
        parts.append(f'<text x="{MARGIN_X - 12}" y="{y + 5}" text-anchor="end" font-family="sans-serif" font-size="12">{f}</text>')

    # string labels at the top
    for col, s in enumerate(_DRAWN_STRINGS):
        x = MARGIN_X + col * STRING_GAP
        parts.append(f'<text x="{x}" y="{MARGIN_Y - 10}" text-anchor="middle" font-family="sans-serif" font-size="12" fill="#555">{s}</text>')

    # root legend — placed inside the reserved bottom strip
    legend_y = height - 15
    parts.append(f'<circle cx="{width - 90}" cy="{legend_y}" r="8" fill="#1f77b4"/>')
    parts.append(f'<text x="{width - 76}" y="{legend_y + 4}" font-family="sans-serif" font-size="11" fill="#333">root</text>')

    # notes — every note in the declared fret range lands on a distinct grid row
    for n in notes:
        s, f = n["string"], n["fret"]
        if not (low <= f <= high):
            continue
        x = _string_x(s)
        row = f - low
        y = MARGIN_Y + row * FRET_GAP + FRET_GAP / 2
        is_root = bool(n.get("is_root"))
        fill = "#1f77b4" if is_root else "#ffffff"
        stroke = "#1f77b4"
        parts.append(f'<circle cx="{x}" cy="{y}" r="{NOTE_R}" fill="{fill}" stroke="{stroke}" stroke-width="2"/>')
        if is_root:
            parts.append(f'<text x="{x}" y="{y + 5}" text-anchor="middle" font-family="sans-serif" font-size="12" fill="#fff">R</text>')

    parts.append('</svg>')
    return "\n".join(parts)


def render_lesson_diagrams(lesson: dict, out_dir: Path) -> list[Path]:
    out_dir.mkdir(parents=True, exist_ok=True)
    written = []
    for d in lesson["musical_context"]["diagrams"]:
        svg = render_scale_svg(d)
        p = out_dir / f"{d['id']}.svg"
        p.write_text(svg)
        written.append(p)
    return written


if __name__ == "__main__":
    import json
    import sys
    if len(sys.argv) != 2:
        print("usage: python render.py <lesson.json>")
        sys.exit(2)
    lesson = json.loads(Path(sys.argv[1]).read_text())
    out = Path(sys.argv[1]).parent / "diagrams"
    for p in render_lesson_diagrams(lesson, out):
        print(f"wrote {p}")
