package config

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

// LoadDotEnv copies KEY=VALUE pairs from the file at path into the process
// environment, for local development convenience. Variables already set in the
// environment win over the file, so explicit exports and container env are
// never overridden. A missing file is not an error (loaded reports false).
//
// Supported syntax: blank lines, full-line # comments, an optional "export "
// prefix, and values wrapped in matching single or double quotes. Everything
// after the first '=' is the value; no escapes or inline comments.
func LoadDotEnv(path string) (loaded bool, err error) {
	f, err := os.Open(path)
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("open %s: %w", path, err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for n := 1; sc.Scan(); n++ {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		line = strings.TrimPrefix(line, "export ")

		key, val, ok := strings.Cut(line, "=")
		key = strings.TrimSpace(key)
		if !ok || key == "" {
			// Report only the line number: the line may contain a secret.
			return false, fmt.Errorf("%s:%d: expected KEY=VALUE", path, n)
		}
		val = unquote(strings.TrimSpace(val))

		if _, set := os.LookupEnv(key); set {
			continue
		}
		if err := os.Setenv(key, val); err != nil {
			return false, fmt.Errorf("%s:%d: set %s: %w", path, n, key, err)
		}
	}
	if err := sc.Err(); err != nil {
		return false, fmt.Errorf("read %s: %w", path, err)
	}
	return true, nil
}

// unquote strips one pair of matching surrounding quotes, if present.
func unquote(s string) string {
	if len(s) >= 2 && (s[0] == '"' || s[0] == '\'') && s[len(s)-1] == s[0] {
		return s[1 : len(s)-1]
	}
	return s
}
