"""Lesson generator: takes a prompt, produces a lesson JSON following lesson-schema.json.

Deterministic parts (scale computation) come from theory.py; the LLM (or a
future call) supplies prose. This MVP uses a template for prose so the whole
pipeline is testable without an LLM.
"""
from __future__ import annotations

import json
import re
from pathlib import Path

from theory import MINOR_PENTATONIC, minor_pentatonic_positions_for_root, scale_notes

OUT = Path(__file__).parent / "lessons"


def detect_request(prompt: str) -> dict:
    """Detect what the user wants; defaults applied for anything unspecified."""
    p = prompt.lower()
    if "minor pentatonic" in p:
        scale = "minor pentatonic"
        key_match = re.search(r"\bin ([a-g])(#|b)?\s*(minor|major)?", p)
        key = f"{key_match.group(1).upper()}{key_match.group(2) or ''}" if key_match else "A"
    else:
        scale = "unknown"
        key = "A"
    level = "intermediate"
    if "beginner" in p:
        level = "beginner"
    elif "advanced" in p:
        level = "advanced"
    position_convention = "5-position" if "position" in p else "5-position"
    return {"scale": scale, "root": key, "level": level, "position_convention": position_convention, "prompt": prompt}


def build_brief(req: dict) -> dict:
    return {
        "root": req["root"],
        "scale": req["scale"],
        "tuning": "standard",
        "level": req["level"],
        "position_convention": req["position_convention"],
        "estimated_minutes": 20,
        "prompt": req["prompt"],
    }


def generate_minor_pentatonic_lesson(brief: dict) -> dict:
    root = brief["root"]
    positions = minor_pentatonic_positions_for_root(root)

    diagrams = []
    for pos in positions:
        pos_num = int(pos["position"].split()[-1])
        diagrams.append({
            "id": f"diagram-{root.lower()}-min-pent-pos-{pos_num}",
            "type": "scale-diagram",
            "name": f"{root} Minor Pentatonic – Position {pos_num}",
            "render": {
                "format": "svg",
                "strings": 6,
                "frets": {"low": pos["fret_range"][0], "high": pos["fret_range"][1]},
                "notes": pos["notes"],
            },
        })

    sections = [
        {
            "title": "What the minor pentatonic scale is",
            "body": (
                f"The {root} minor pentatonic scale uses five notes derived from the {root} natural minor scale, "
                "omitting the 2nd and 6th degrees. This produces a stable, blues-friendly sound with no "
                "semitone clashes between adjacent scale tones. In " + f"{root} minor the five notes are "
                f"{' / '.join(scale_notes(root, MINOR_PENTATONIC))}, "
                "corresponding to the intervals root (1), minor third (b3), perfect fourth (4), perfect fifth (5), "
                "and minor seventh (b7). On the fretboard it is usually taught as five overlapping positions; "
                "learning to move between them smoothly is the foundation for improvising solos."
            ),
            "diagram_refs": [diagrams[0]["id"]],
        }
    ]
    for i, d in enumerate(diagrams, start=1):
        if i < 5:
            connect = (
                f"connect it to position {i+1} by shifting to the nearest shared note "
                "on the same string rather than jumping"
            )
        else:
            connect = (
                "return toward position 4 through the shared high-E note at fret 15 "
                "(play frets 17, 15, 12 descending)"
            )
        sections.append({
            "title": f"Position {i}: {root} minor pentatonic, frets {d['render']['frets']['low']}-{d['render']['frets']['high']}",
            "body": (
                f"Position {i} covers frets {d['render']['frets']['low']} through {d['render']['frets']['high']} and "
                "contains every note of the scale in a compact span. Play it ascending and descending, "
                "using one finger per fret where possible, and keep your thumb behind the neck. Once you can play "
                f"it cleanly at a moderate tempo, {connect}."
            ),
            "diagram_refs": [d["id"]],
        })

    practice_steps = [
        {
            "goal": "Memorize position 1 shape and play it cleanly",
            "action": (
                f"Play the {root} minor pentatonic position 1 (frets {diagrams[0]['render']['frets']['low']}-{diagrams[0]['render']['frets']['high']}) "
                "up and down at 60 BPM, one note per click, using alternate picking. Start on the root on the low E "
                f"string (string 6, fret {next(n['fret'] for n in diagrams[0]['render']['notes'] if n['string']==6 and n['is_root'])}), "
                "ascend the full box to fret 8 on the high E string, then descend back to fret 5 on string 6. "
                "Add a slight accent on each marked root note (blue circles with R)."
            ),
            "diagram_refs": [diagrams[0]["id"]],
            "duration_minutes": 3,
            "bpm": 60,
            "reps": 4,
            "success_criterion": "You can ascend to fret 8 and descend back to fret 5 twice in a row without stopping or missing a note.",
        },
        {
            "goal": "Move smoothly between positions 1 and 2",
            "action": (
                "On string 1 (high E), play eight quarter notes in two bars of 4/4: "
                "frets 5, 8, 10, 8 | 8, 10, 8, 5. Fret 8 is shared by positions 1 and 2; "
                "fret 10 enters position 2, and the final fret 5 returns to A."
            ),
            "diagram_refs": [diagrams[0]["id"], diagrams[1]["id"]],
            "duration_minutes": 4,
            "bpm": 60,
            "reps": 6,
            "success_criterion": "You can play the full two-bar phrase without stopping and end cleanly on fret 5 (the root A).",
        },
        {
            "goal": "Connect all five positions along one string",
            "action": (
                f"Start on the root of position 1 on string 6 (fret {next(n['fret'] for n in diagrams[0]['render']['notes'] if n['string']==6 and n['is_root'])}) and play "
                "the pentatonic notes ascending along string 6 through position 5, then reverse and descend: "
                "frets 5, 8, 10, 12, 15, 17, 15, 12, 10, 8, 5 — one note per click, no stopping. "
                "This connects all five boxes along one string; do not attempt to jump back from position 5 "
                "to position 1 — they sit an octave apart."
            ),
            "diagram_refs": [d["id"] for d in diagrams],
            "duration_minutes": 5,
            "bpm": 50,
            "reps": 3,
            "success_criterion": "One unbroken ascending and descending pass through all five positions along string 6.",
        },
        {
            "goal": "Improvise using only position 1",
            "action": (
                f"Over an {root} minor backing track, improvise for two minutes using only position 1. "
                "Emphasize the root and the minor third. Rest for the same amount of time between takes."
            ),
            "diagram_refs": [diagrams[0]["id"]],
            "duration_minutes": 4,
            "success_criterion": "You can improvise for the full two minutes without playing outside the scale.",
        },
        {
            "goal": "Use all five positions melodically",
            "action": (
                "Improvise over the same backing track, but now deliberately move between positions every "
                "four bars. Target the root note in each new position as your landing note."
            ),
            "diagram_refs": [d["id"] for d in diagrams],
            "duration_minutes": 4,
            "success_criterion": "You land on the root in at least three different positions across the take.",
        },
    ]

    tips = [
        "Think of the five positions as one connected shape on the fretboard, not five separate patterns — every position shares notes with its neighbours.",
        "Anchor your hand with your index finger at the low end of each position; this keeps your other three fingers free for the higher frets.",
        "Practice positions in reverse order (5 back to 1) occasionally — most players over-rely on positions 1 and 2.",
        "Sing the root note before you start improvising so your ear has a target, not just your fingers.",
    ]

    related_topics = [
        {"id": "blues-scale", "title": "The Blues Scale (adding the b5)", "reason": "One extra note turns minor pentatonic into the blues scale; natural next step."},
        {"id": "major-pentatonic", "title": "Major Pentatonic Scale", "reason": "A minor pentatonic and C major pentatonic share the same notes and fret locations; only the tonal center changes from A to C."},
        {"id": "string-bending", "title": "String Bending Basics", "reason": "Bends from the 4 to the 5 are the classic pentatonic move."},
        {"id": "relative-major-minor", "title": "Relative Major and Minor", "reason": "Explains why A minor pentatonic and C major pentatonic share a shape."},
        {"id": "improvisation-basics", "title": "Improvisation Basics", "reason": "Practical guidance for building melodic solos from these shapes."},
    ]

    return {
        "schema_version": 1,
        "id": f"{root.lower()}-minor-pentatonic-positions",
        "title": f"Using the {root} Minor Pentatonic Scale Across All Five Positions",
        "metadata": {
            "audience_level": brief["level"],
            "objectives": [
                "Identify the five notes of the minor pentatonic scale in any key",
                "Play all five positions cleanly at a moderate tempo",
                "Transition between adjacent positions without pausing",
                "Improvise using the scale in a musical context",
            ],
            "prerequisites": [
                "Comfortable with alternate picking",
                "Can play a one-octave major scale",
                "Knows the note names on the low E string up to the 12th fret",
            ],
            "estimated_minutes": brief["estimated_minutes"],
            "prompt": brief["prompt"],
        },
        "musical_context": {
            "tuning": "standard",
            "key": f"{root} minor",
            "position_convention": brief["position_convention"],
            "diagrams": diagrams,
        },
        "sections": sections,
        "practice_steps": practice_steps,
        "tips": tips,
        "related_topics": related_topics,
    }


def generate(prompt: str) -> Path:
    req = detect_request(prompt)
    brief = build_brief(req)
    if req["scale"] != "minor pentatonic":
        raise ValueError(f"unsupported scale: {req['scale']}")
    if req["root"] != "A":
        raise ValueError(
            f"unsupported root '{req['root']}': this prototype only generates A minor pentatonic "
            "(hand-checked fixtures exist for A only). Add per-root fixtures before other keys."
        )
    lesson = generate_minor_pentatonic_lesson(brief)
    OUT.mkdir(parents=True, exist_ok=True)
    p = OUT / f"{lesson['id']}.json"
    p.write_text(json.dumps(lesson, indent=2))
    return p


if __name__ == "__main__":
    import sys
    prompt = sys.argv[1] if len(sys.argv) > 1 else "create a lesson on how to use the minor pentatonic scale and the various positions"
    p = generate(prompt)
    print(f"wrote {p}")
