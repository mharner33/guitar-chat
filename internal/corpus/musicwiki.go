package corpus

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// wikiFolders are the OKF concept folders, in ingest order. Bundle-root files
// (index.md, log.md, README.md, LICENSE) and every index.md are not concepts.
var wikiFolders = []string{"intervals", "scales", "triads", "chords", "keys"}

// ReadMusicWiki reads and flattens every concept node in an OKF bundle. All
// nodes are parsed before any is flattened, so relationship slugs resolve to
// titles across the whole bundle.
func ReadMusicWiki(dir string) ([]Document, error) {
	var nodes []node
	titles := map[string]string{}
	for _, sub := range wikiFolders {
		entries, err := os.ReadDir(filepath.Join(dir, sub))
		if errors.Is(err, fs.ErrNotExist) {
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", filepath.Join(dir, sub), err)
		}
		for _, e := range entries {
			if e.IsDir() || filepath.Ext(e.Name()) != ".md" || e.Name() == "index.md" {
				continue
			}
			path := filepath.Join(dir, sub, e.Name())
			raw, err := os.ReadFile(path)
			if err != nil {
				return nil, fmt.Errorf("read %s: %w", path, err)
			}
			n, err := parseNode(path, raw)
			if err != nil {
				return nil, err
			}
			if _, dup := titles[n.Slug]; dup {
				return nil, fmt.Errorf("%s: duplicate slug %q", path, n.Slug)
			}
			titles[n.Slug] = n.Title
			nodes = append(nodes, n)
		}
	}

	docs := make([]Document, 0, len(nodes))
	for _, n := range nodes {
		chunks, err := flatten(n, titles)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", n.Path, err)
		}
		docs = append(docs, Document{ID: WikiPrefix + n.Slug, Title: n.Title, Chunks: chunks})
	}
	return docs, nil
}
