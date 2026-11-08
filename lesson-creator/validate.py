"""JSON Schema validator + semantic validation with independent fixtures.

The schema (lesson-schema.json) is enforced via jsonschema.Draft202012Validator.
Semantic checks only run on schema-valid input. Position patterns are compared
against fixtures.py (hand-checked reference shapes), not regenerated output.
"""
from __future__ import annotations

import json
import sys
from pathlib import Path

try:
    import jsonschema
except ImportError:
    print("pip install jsonschema", file=sys.stderr)
    sys.exit(2)

from fixtures import POSITIONS, fixture_notes
from theory import MINOR_PENTATONIC, note_at, scale_notes

SCHEMA_PATH = Path(__file__).parent / "lesson-schema.json"
CATALOG_PATH = Path(__file__).parent / "catalog.json"
TOPICS_DIR = Path(__file__).parent / "topics"


class ValidationReport:
    def __init__(self):
        self.errors = []
        self.warnings = []

    def error(self, msg):
        self.errors.append(msg)

    def warn(self, msg):
        self.warnings.append(msg)

    @property
    def ok(self):
        return not self.errors


def load_catalog() -> dict:
    """Return {id: entry} for catalog entries, or {} if catalog is absent."""
    if CATALOG_PATH.exists():
        return {t["id"]: t for t in json.loads(CATALOG_PATH.read_text())}
    return {}


def validate_lesson(lesson: dict) -> ValidationReport:
    rep = ValidationReport()
    schema = json.loads(SCHEMA_PATH.read_text())
    validator = jsonschema.Draft202012Validator(schema)

    # --- 1. JSON Schema hard gate ---
    for err in sorted(validator.iter_errors(lesson), key=lambda e: list(e.absolute_path)):
        path = "/".join(str(p) for p in err.absolute_path) or "<root>"
        rep.error(f"schema: {path}: {err.message}")

    if rep.errors:
        return rep  # no point doing semantic checks on invalid structure

    # --- 2. Duplicate diagram IDs ---
    ids = [d["id"] for d in lesson["musical_context"]["diagrams"]]
    if len(ids) != len(set(ids)):
        rep.error("duplicate diagram IDs found")

    # --- 3. Diagram refs resolve ---
    diagram_ids = set(ids)
    for sec in lesson["sections"]:
        for ref in sec["diagram_refs"]:
            if ref not in diagram_ids:
                rep.error(f"section '{sec['title']}' references unknown diagram: {ref}")
    for step in lesson["practice_steps"]:
        for ref in step["diagram_refs"]:
            if ref not in diagram_ids:
                rep.error(f"practice step '{step['goal']}' references unknown diagram: {ref}")

    # --- 4. Related topics resolve against catalog AND the catalog's actual destination exists ---
    catalog = load_catalog()
    for topic in lesson["related_topics"]:
        if topic["id"] not in catalog:
            rep.error(f"related topic '{topic['id']}' not in catalog (catalog has {sorted(catalog)})")
            continue
        dest_rel = catalog[topic["id"]].get("dest")
        if not dest_rel:
            rep.error(f"related topic '{topic['id']}' has no destination in catalog entry")
            continue
        # resolve the catalog's declared destination (same path preview.py links to)
        target = (Path(__file__).parent / dest_rel).resolve()
        if not target.exists():
            rep.error(
                f"related topic '{topic['id']}' destination file missing at '{dest_rel}' (catalog entry); "
                "create it before linking"
            )

    # --- 5. Independent position-fixture comparison ---
    # Compare the lesson's diagrams against fixtures.py — NOT regenerated output.
    fixture = fixture_notes()
    diagrams = lesson["musical_context"]["diagrams"]
    if len(diagrams) != len(POSITIONS):
        rep.error(f"expected {len(POSITIONS)} position diagrams, got {len(diagrams)}")
    for i, d in enumerate(diagrams):
        if i >= len(POSITIONS):
            break
        # 5a. Declared fret bounds must be sane and contain every drawn note.
        low = d["render"]["frets"]["low"]
        high = d["render"]["frets"]["high"]
        if low > high:
            rep.error(f"diagram {d['id']}: declared fret range low={low} > high={high}")
        if not d["render"]["notes"]:
            rep.error(f"diagram {d['id']}: diagram declares notes but contains none")
        for n in d["render"]["notes"]:
            if not (low <= n["fret"] <= high):
                rep.error(
                    f"diagram {d['id']}: note string {n['string']} fret {n['fret']} "
                    f"is outside the declared range [{low}, {high}]"
                )
        expected = {n["string"] and (n["string"], n["fret"]) for n in fixture if n["position"] == POSITIONS[i]["name"]}
        drawn = {(n["string"], n["fret"]) for n in d["render"]["notes"]}
        missing = expected - drawn
        extra = drawn - expected
        if missing:
            rep.error(f"diagram {d['id']} ({POSITIONS[i]['name']}): missing {sorted(missing)}")
        if extra:
            rep.error(f"diagram {d['id']} ({POSITIONS[i]['name']}): unexpected {sorted(extra)}")
        # root correctness
        for n in d["render"]["notes"]:
            actual = note_at(n["string"], n["fret"])
            if actual != n["pitch_class"]:
                rep.error(f"diagram {d['id']}: s{n['string']} f{n['fret']} labeled {n['pitch_class']} but is {actual}")
            expected_root = any(
                f["string"] == n["string"] and f["fret"] == n["fret"] and f["is_root"] for f in fixture
            )
            if bool(n["is_root"]) != expected_root:
                rep.error(f"diagram {d['id']}: s{n['string']} f{n['fret']} is_root={n['is_root']} but fixture says {expected_root}")

    # --- 6. Scope restriction: only A minor pentatonic is supported ---
    # Fixtures are A minor; validating another root against them would produce
    # nonsense rather than a failure, so reject unsupported roots explicitly.
    if lesson["musical_context"]["key"] != "A minor":
        rep.error(
            f"unsupported key '{lesson['musical_context']['key']}': this prototype only "
            "validates A minor pentatonic; add per-root fixtures before other keys"
        )

    # --- 7. Scale membership sanity ---
    key = lesson["musical_context"]["key"]
    if "minor" in key:
        root = key.split()[0]
        allowed = set(scale_notes(root, MINOR_PENTATONIC))
        for d in diagrams:
            for n in d["render"]["notes"]:
                if n["pitch_class"] not in allowed:
                    rep.error(f"diagram {d['id']}: note {n['pitch_class']} not in {root} minor pentatonic")

    # --- 8. Practice steps reference achievable string/fret paths ---
    # (Advisory only — pedagogical quality, not a hard gate)
    for step in lesson["practice_steps"]:
        body = step["action"]
        if "G and B" in body or "string 3" in body and "position 1" in body.lower():
            rep.warn(f"practice step '{step['goal']}': check that referenced strings actually contain the notes described")

    return rep


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("usage: python validate.py <lesson.json>")
        sys.exit(2)
    lesson = json.loads(Path(sys.argv[1]).read_text())
    rep = validate_lesson(lesson)
    for e in rep.errors:
        print(f"ERROR: {e}")
    for w in rep.warnings:
        print(f"WARNING: {w}")
    if rep.ok:
        print(f"OK: lesson '{lesson['title']}' passed validation")
    sys.exit(0 if rep.ok else 1)
