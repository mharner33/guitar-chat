package corpus

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestLoad(t *testing.T) {
	root := writeTree(t, map[string]string{
		"musicwiki/intervals/u.md": "---\ntype: interval\nslug: /intervals/u\ntitle: U\ntheory: {}\n---\n",
		"guides/g.md":              "---\ntitle: G\nrelated: [/intervals/u]\n---\n\n## A\nB.\n",
		"empty/.keep":              "",
		"notadir":                  "file",
	})
	wiki := filepath.Join(root, "musicwiki")
	guides := filepath.Join(root, "guides")

	t.Run("guides listed first still resolve wiki slugs", func(t *testing.T) {
		got, err := Load([]string{guides, wiki})
		if err != nil {
			t.Fatalf("Load: %v", err)
		}
		if len(got) != 2 || got[0].Prefix != WikiPrefix || got[1].Prefix != GuidePrefix {
			t.Fatalf("sources = %+v, want wiki then guides", got)
		}
	})
	t.Run("trailing slash and ./ still detect musicwiki", func(t *testing.T) {
		got, err := Load([]string{wiki + "/", guides})
		if err != nil || got[0].Prefix != WikiPrefix {
			t.Fatalf("Load = %+v, %v; want wiki source", got, err)
		}
	})

	errs := []struct {
		name    string
		dirs    []string
		wantErr string
	}{
		{"missing dir", []string{wiki, filepath.Join(root, "gudies")}, "gudies"},
		{"not a directory", []string{filepath.Join(root, "notadir")}, "not a directory"},
		{"empty guides dir", []string{wiki, filepath.Join(root, "empty")}, "no documents"},
		{"two guide dirs share a prefix", []string{guides, filepath.Join(root, "empty")}, "at most one"},
		{"two wiki dirs", []string{wiki, wiki}, "at most one"},
		{"guide related needs the wiki", []string{guides}, `related slug "/intervals/u"`},
	}
	for _, tt := range errs {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Load(tt.dirs)
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("err = %v, want containing %q", err, tt.wantErr)
			}
		})
	}
}
