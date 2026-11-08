// Package rag holds the retrieval-augmented generation pipeline: chunking,
// embedding, retrieval, and prompt assembly.
package rag

import (
	"fmt"
	"regexp"
	"strings"
)

// Chunk is one embeddable piece of a document.
type Chunk struct {
	Section string
	Text    string
}

// Size limits approximate tokens as ~4 characters each, so no tokenizer
// dependency is needed: ~800 tokens per chunk, ~100 tokens of overlap.
const (
	maxChunkChars     = 3200
	overlapChunkChars = 400
)

var paragraphBreak = regexp.MustCompile(`\n[ \t]*\n`)

// ChunkMarkdown splits a Markdown body on "## " headings. Text before the first
// heading becomes the "Introduction" section; "# " title lines are dropped
// because the document title is carried separately. Sections longer than
// maxChunkChars are split at paragraph boundaries with a small overlap.
func ChunkMarkdown(body string) []Chunk {
	var out []Chunk
	section := "Introduction"
	var buf []string
	flush := func() {
		text := strings.TrimSpace(strings.Join(buf, "\n"))
		buf = buf[:0]
		if text != "" {
			out = append(out, splitSection(section, text)...)
		}
	}
	for _, line := range strings.Split(strings.ReplaceAll(body, "\r\n", "\n"), "\n") {
		switch {
		case strings.HasPrefix(line, "## "):
			flush()
			section = strings.TrimSpace(strings.TrimPrefix(line, "## "))
		case strings.HasPrefix(line, "# "):
			// Title line: the document title is stored on the document.
		default:
			buf = append(buf, line)
		}
	}
	flush()
	return out
}

func splitSection(section, text string) []Chunk {
	if len(text) <= maxChunkChars {
		return []Chunk{{Section: section, Text: text}}
	}
	var paras []string
	for _, p := range paragraphBreak.Split(text, -1) {
		if p = strings.TrimSpace(p); p != "" {
			paras = append(paras, p)
		}
	}

	var chunks []Chunk
	var cur []string
	curLen := 0
	emit := func() {
		name := section
		if len(chunks) > 0 {
			name = fmt.Sprintf("%s (part %d)", section, len(chunks)+1)
		}
		chunks = append(chunks, Chunk{Section: name, Text: strings.Join(cur, "\n\n")})
	}
	for _, p := range paras {
		if curLen > 0 && curLen+len(p)+2 > maxChunkChars {
			emit()
			cur, curLen = overlapTail(cur)
		}
		cur = append(cur, p)
		curLen += len(p) + 2
	}
	emit()
	return chunks
}

// overlapTail returns the trailing paragraphs of paras whose combined length
// fits in overlapChunkChars, as a fresh slice, plus that length.
func overlapTail(paras []string) ([]string, int) {
	n, total := 0, 0
	for i := len(paras) - 1; i >= 0; i-- {
		if total+len(paras[i])+2 > overlapChunkChars {
			break
		}
		total += len(paras[i]) + 2
		n++
	}
	return append([]string(nil), paras[len(paras)-n:]...), total
}
