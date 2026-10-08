package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadDotEnv(t *testing.T) {
	tests := []struct {
		name     string
		contents string
		preset   map[string]string // already in the environment before loading
		want     map[string]string // expected values after loading
		unset    []string          // keys that must remain unset
		wantErr  bool
	}{
		{
			name:     "plain key=value",
			contents: "GC_A=one\nGC_B=two\n",
			want:     map[string]string{"GC_A": "one", "GC_B": "two"},
		},
		{
			name:     "comments, blank lines, export prefix, whitespace",
			contents: "# a comment\n\n  export GC_A = one  \n\t\nGC_B=two # not stripped\n",
			want:     map[string]string{"GC_A": "one", "GC_B": "two # not stripped"},
		},
		{
			name:     "matching quotes are stripped, inner = kept",
			contents: "GC_A=\"postgres://u:p@h/db?x=1\"\nGC_B='single'\nGC_C=\"unbalanced\n",
			want:     map[string]string{"GC_A": "postgres://u:p@h/db?x=1", "GC_B": "single", "GC_C": "\"unbalanced"},
		},
		{
			name:     "real environment wins over the file",
			contents: "GC_A=from-file\nGC_B=from-file\n",
			preset:   map[string]string{"GC_A": "from-env"},
			want:     map[string]string{"GC_A": "from-env", "GC_B": "from-file"},
		},
		{
			name:     "empty value is set",
			contents: "GC_A=\n",
			want:     map[string]string{"GC_A": ""},
		},
		{
			name:     "line without = is an error",
			contents: "GC_A=one\nNOT_AN_ASSIGNMENT\n",
			wantErr:  true,
		},
		{
			name:     "empty key is an error",
			contents: "=value\n",
			wantErr:  true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for _, k := range []string{"GC_A", "GC_B", "GC_C"} {
				t.Setenv(k, "") // registers restore on cleanup
				os.Unsetenv(k)
			}
			for k, v := range tc.preset {
				t.Setenv(k, v)
			}
			path := filepath.Join(t.TempDir(), ".env")
			if err := os.WriteFile(path, []byte(tc.contents), 0o600); err != nil {
				t.Fatal(err)
			}

			loaded, err := LoadDotEnv(path)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !loaded {
				t.Fatal("loaded = false, want true")
			}
			for k, want := range tc.want {
				got, ok := os.LookupEnv(k)
				if !ok || got != want {
					t.Errorf("%s = %q (set=%v), want %q", k, got, ok, want)
				}
			}
		})
	}
}

func TestLoadDotEnvMissingFileIsNotAnError(t *testing.T) {
	loaded, err := LoadDotEnv(filepath.Join(t.TempDir(), "does-not-exist.env"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if loaded {
		t.Fatal("loaded = true, want false")
	}
}
