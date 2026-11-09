package corpus

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Source is one ingest source: its directory, the id prefix it owns, and its
// documents. Reconcile deletes only documents with this prefix.
type Source struct {
	Dir    string
	Prefix string
	Docs   []Document
}

// Load reads every configured source before anything is written. A directory
// whose base name is "musicwiki" is read as an OKF bundle; any other as
// guides. Wiki sources are read first so guides can validate related slugs.
// It refuses missing or empty directories (a typo must not reconcile a whole
// source away) and more than one directory per kind (they would share a
// prefix and delete each other's documents).
func Load(dirs []string) ([]Source, error) {
	var wikiDirs, guideDirs []string
	for _, d := range dirs {
		if filepath.Base(filepath.Clean(d)) == "musicwiki" {
			wikiDirs = append(wikiDirs, d)
		} else {
			guideDirs = append(guideDirs, d)
		}
	}
	if len(wikiDirs) > 1 || len(guideDirs) > 1 {
		return nil, fmt.Errorf("CORPUS_DIRS %q: at most one musicwiki and one guides directory (each kind owns one id prefix)", strings.Join(dirs, ","))
	}

	var sources []Source
	wikiSlugs := map[string]bool{}
	for _, d := range wikiDirs {
		if err := requireDir(d); err != nil {
			return nil, err
		}
		docs, err := ReadMusicWiki(d)
		if err != nil {
			return nil, err
		}
		if len(docs) == 0 {
			return nil, fmt.Errorf("corpus dir %q: no documents", d)
		}
		for _, doc := range docs {
			wikiSlugs[strings.TrimPrefix(doc.ID, WikiPrefix)] = true
		}
		sources = append(sources, Source{Dir: d, Prefix: WikiPrefix, Docs: docs})
	}
	for _, d := range guideDirs {
		if err := requireDir(d); err != nil {
			return nil, err
		}
		docs, err := ReadGuides(d, wikiSlugs)
		if err != nil {
			return nil, err
		}
		if len(docs) == 0 {
			return nil, fmt.Errorf("corpus dir %q: no documents", d)
		}
		sources = append(sources, Source{Dir: d, Prefix: GuidePrefix, Docs: docs})
	}
	return sources, nil
}

func requireDir(d string) error {
	fi, err := os.Stat(d)
	if err != nil {
		return fmt.Errorf("corpus dir %q: %w", d, err)
	}
	if !fi.IsDir() {
		return fmt.Errorf("corpus dir %q: not a directory", d)
	}
	return nil
}
