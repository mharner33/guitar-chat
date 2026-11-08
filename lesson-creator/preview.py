"""Render a lesson JSON as readable Markdown for human review, with embedded diagram images and resolved topic links."""
import json
import sys
from pathlib import Path


def _load_catalog():
    catalog_path = Path(__file__).parent / "catalog.json"
    if catalog_path.exists():
        return {t["id"]: t for t in json.loads(catalog_path.read_text())}
    return {}


def _topic_link(catalog: dict, topic: dict, base_dir: Path) -> str:
    entry = catalog.get(topic["id"])
    if not entry:
        return f"**{topic['title']}** (`{topic['id']}` — no catalog entry)"
    dest = entry.get("dest")
    if not dest:
        return f"**{topic['title']}** (`{topic['id']}` — no destination)"
    # compute a relative path from the lesson file's directory to the dest file
    target = (Path(__file__).parent / dest).resolve()
    try:
        rel = target.relative_to(base_dir)
        return f"[{topic['title']}]({rel.as_posix()})"
    except ValueError:
        # fall back to walking up
        import os
        return f"[{topic['title']}]({os.path.relpath(target, base_dir).replace(os.sep, '/')})"


def render_markdown(lesson: dict) -> str:
    lines = []
    lines.append(f"# {lesson['title']}\n")
    m = lesson["metadata"]
    lines.append(f"**Level:** {m['audience_level']}  ")
    lines.append(f"**Estimated time:** {m['estimated_minutes']} minutes  ")
    lines.append(f"**Key:** {lesson['musical_context']['key']}  ")
    lines.append(f"**Position convention:** {lesson['musical_context']['position_convention']}\n")
    lines.append("## Objectives\n")
    for o in m["objectives"]:
        lines.append(f"- {o}")
    lines.append("\n## Prerequisites\n")
    for p in m["prerequisites"]:
        lines.append(f"- {p}")
    lines.append("\n## Sections\n")
    for sec in lesson["sections"]:
        lines.append(f"### {sec['title']}\n")
        lines.append(sec["body"])
        if sec.get("diagram_refs"):
            lines.append("\nDiagrams: " + ", ".join(f"`{r}`" for r in sec["diagram_refs"]))
        lines.append("")
    lines.append("## Practice Steps\n")
    for i, step in enumerate(lesson["practice_steps"], start=1):
        lines.append(f"### Step {i}: {step['goal']}\n")
        lines.append(f"**Action:** {step['action']}\n")
        extra = []
        if "bpm" in step:
            extra.append(f"bpm={step['bpm']}")
        if "reps" in step:
            extra.append(f"reps={step['reps']}")
        lines.append(f"**Duration:** {step['duration_minutes']} min" + (f" ({', '.join(extra)})" if extra else ""))
        lines.append(f"**Success criterion:** {step['success_criterion']}\n")
        if step.get("diagram_refs"):
            lines.append(f"Diagrams: {', '.join(step['diagram_refs'])}\n")
    lines.append("## Tips\n")
    for t in lesson["tips"]:
        lines.append(f"- {t}")
    lines.append("\n## Related Topics\n")
    catalog = _load_catalog()
    lesson_dir = Path(sys.argv[1]).resolve().parent if len(sys.argv) > 1 else Path.cwd()
    for topic in lesson["related_topics"]:
        link = _topic_link(catalog, topic, lesson_dir)
        lines.append(f"- {link} — {topic.get('reason', '')}")
    lines.append("\n## Diagrams\n")
    for d in lesson["musical_context"]["diagrams"]:
        png = Path("diagrams") / f"{d['id']}.png"
        lines.append(f"### {d['name']}\n")
        lines.append(f"![{d['name']}]({png.as_posix()})")
        lines.append(f"\n(`{d['id']}`, frets {d['render']['frets']['low']}-{d['render']['frets']['high']})\n")
    return "\n".join(lines)


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("usage: python preview.py <lesson.json>")
        sys.exit(2)
    lesson = json.loads(Path(sys.argv[1]).read_text())
    out = Path(sys.argv[1]).with_suffix(".md")
    out.write_text(render_markdown(lesson))
    print(f"wrote {out}")
