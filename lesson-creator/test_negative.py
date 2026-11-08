"""Negative mutation tests: prove the validator catches real errors, not just passes everything.

Run: python test_negative.py  (exits 0 if all mutations are caught)
"""
import copy
import json
import sys
from pathlib import Path

from validate import CATALOG_PATH, validate_lesson

LESSON_PATH = Path(__file__).parent / "lessons" / "a-minor-pentatonic-positions.json"
lesson = json.loads(LESSON_PATH.read_text())

# stashed pristine catalog for the missing-destination case
_PRISTINE_CATALOG = CATALOG_PATH.read_text()

CASES = []


def case(name):
    def deco(fn):
        CASES.append((name, fn))
        return fn
    return deco


@case("root flag flipped to False")
def _(l):
    for n in l["musical_context"]["diagrams"][0]["render"]["notes"]:
        if n["string"] == 6 and n["fret"] == 5:
            n["is_root"] = False


@case("extra out-of-scale note added")
def _(l):
    l["musical_context"]["diagrams"][0]["render"]["notes"].append(
        {"string": 2, "fret": 5, "pitch_class": "F", "is_root": False}
    )


@case("note removed from position 1")
def _(l):
    notes = l["musical_context"]["diagrams"][0]["render"]["notes"]
    notes[:] = [n for n in notes if not (n["string"] == 3 and n["fret"] == 5)]


@case("related topic id not in catalog")
def _(l):
    l["related_topics"][0]["id"] = "made-up-topic"


@case("broken diagram ref")
def _(l):
    l["sections"][0]["diagram_refs"] = ["diagram-nope"]


@case("invalid audience enum")
def _(l):
    l["metadata"]["audience_level"] = "expert"


@case("declared fret bounds exclude notes (regression: low=100, high=103)")
def _(l):
    l["musical_context"]["diagrams"][0]["render"]["frets"] = {"low": 100, "high": 103}
    # leave the notes unchanged; the validator must catch that every note is outside


@case("declared fret range reversed (low > high)")
def _(l):
    l["musical_context"]["diagrams"][0]["render"]["frets"] = {"low": 9, "high": 5}


@case("unsupported root (E minor)")
def _(l):
    l["musical_context"]["key"] = "E minor"


def main():
    failures = []

    # first nine cases are pure in-memory mutations
    pure_cases = [c for c in CASES if c[0] != "related topic with catalog entry but missing destination file"]
    for name, mutate in pure_cases:
        mutated = copy.deepcopy(lesson)
        mutate(mutated)
        rep = validate_lesson(mutated)
        if rep.ok:
            failures.append(name)
            print(f"MISSED: {name} (validator passed but should have failed)")
        else:
            print(f"CAUGHT: {name}")
            for e in rep.errors[:2]:
                print(f"    -> {e[:100]}")

    # final case needs a file mutation; run it with try/finally so the catalog is always restored
    entries = json.loads(_PRISTINE_CATALOG)
    for entry in entries:
        if entry["id"] == "blues-scale":
            entry["dest"] = "topics/does-not-exist.md"
            break
    try:
        CATALOG_PATH.write_text(json.dumps(entries, indent=2))
        mutated = copy.deepcopy(lesson)
        rep = validate_lesson(mutated)
        name = "related topic with catalog entry but missing destination file"
        if rep.ok:
            failures.append(name)
            print(f"MISSED: {name} (validator passed but should have failed)")
        else:
            print(f"CAUGHT: {name}")
            for e in rep.errors[:2]:
                print(f"    -> {e[:100]}")
            if not any("destination file missing" in e for e in rep.errors):
                print("    -> WARNING: failure was NOT a missing-destination error")
                failures.append(name + " (wrong reason)")
    finally:
        CATALOG_PATH.write_text(_PRISTINE_CATALOG)

    if failures:
        print(f"\n{len(failures)} mutation(s) NOT caught by the validator")
        sys.exit(1)
    total = len(CASES) + 1  # the missing-destination case runs separately in main()
    print(f"\nAll {total} mutations caught by the validator")


if __name__ == "__main__":
    main()
