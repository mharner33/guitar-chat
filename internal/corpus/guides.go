package corpus

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go.yaml.in/yaml/v3"

	"github.com/mharner33/guitar-chat/internal/rag"
)

type guideFront struct {
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
	Tags        []string `yaml:"tags"`
	Related     []string `yaml:"related"`
}

// ReadGuides reads every hand-written guide in dir. Each guide's related slugs
// must name musicwiki nodes, so a typo fails ingest instead of leaving a dead
// link.
func ReadGuides(dir string, wikiSlugs map[string]bool) ([]Document, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", dir, err)
	}
	var docs []Document
	for _, e := range entries {
		if e.IsDir() || filepath.Ext(e.Name()) != ".md" || e.Name() == "README.md" {
			continue
		}
		path := filepath.Join(dir, e.Name())
		doc, err := readGuide(path, wikiSlugs)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", path, err)
		}
		docs = append(docs, doc)
	}
	return docs, nil
}

func readGuide(path string, wikiSlugs map[string]bool) (Document, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return Document{}, err
	}
	front, body, err := splitFrontmatter(raw)
	if err != nil {
		return Document{}, err
	}
	var g guideFront
	if err := yaml.Unmarshal(front, &g); err != nil {
		return Document{}, fmt.Errorf("parse frontmatter: %w", err)
	}
	if strings.TrimSpace(g.Title) == "" {
		return Document{}, fmt.Errorf("missing required field %q", "title")
	}
	for _, s := range g.Related {
		if !wikiSlugs[s] {
			return Document{}, fmt.Errorf("related slug %q is not a musicwiki node", s)
		}
	}
	sections := rag.ChunkMarkdown(body)
	if len(sections) == 0 {
		return Document{}, fmt.Errorf("guide has no content")
	}
	chunks := make([]Chunk, len(sections))
	for i, s := range sections {
		chunks[i] = Chunk{Section: s.Section, Text: g.Title + " — " + s.Section + "\n" + s.Text}
	}
	stem := strings.TrimSuffix(filepath.Base(path), ".md")
	return Document{ID: GuidePrefix + stem, Title: g.Title, Chunks: chunks}, nil
}
