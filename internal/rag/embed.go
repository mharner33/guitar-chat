package rag

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Embedder turns texts into vectors. It is defined here, at its consumers
// (ingest and the answer pipeline), so tests can fake it.
type Embedder interface {
	Embed(ctx context.Context, texts []string) ([][]float32, error)
}

// HTTPEmbedder calls an OpenAI-compatible /embeddings endpoint.
type HTTPEmbedder struct {
	BaseURL    string // e.g. https://api.openai.com/v1
	APIKey     string
	Model      string
	Dimensions int // expected vector width (EMBEDDING_DIMENSIONS)
	Client     *http.Client
}

// Embed returns one vector per text, in input order. The caller's context
// carries the deadline.
func (e *HTTPEmbedder) Embed(ctx context.Context, texts []string) ([][]float32, error) {
	if len(texts) == 0 {
		return nil, nil
	}
	body, err := json.Marshal(map[string]any{"model": e.Model, "input": texts})
	if err != nil {
		return nil, fmt.Errorf("embeddings: encode request: %w", err)
	}
	url := strings.TrimRight(e.BaseURL, "/") + "/embeddings"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("embeddings: build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+e.APIKey)

	client := e.Client
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("embeddings request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		msg, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		return nil, fmt.Errorf("embeddings: %s: %s", resp.Status, strings.TrimSpace(string(msg)))
	}

	var out struct {
		Data []struct {
			Index     int       `json:"index"`
			Embedding []float32 `json:"embedding"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, fmt.Errorf("embeddings: decode response: %w", err)
	}
	if len(out.Data) != len(texts) {
		return nil, fmt.Errorf("embeddings: got %d embeddings for %d inputs", len(out.Data), len(texts))
	}
	vecs := make([][]float32, len(texts))
	for _, d := range out.Data {
		if d.Index < 0 || d.Index >= len(texts) || vecs[d.Index] != nil {
			return nil, fmt.Errorf("embeddings: bad index %d", d.Index)
		}
		if len(d.Embedding) != e.Dimensions {
			return nil, fmt.Errorf("embeddings: got %d dimensions, want %d (EMBEDDING_DIMENSIONS)", len(d.Embedding), e.Dimensions)
		}
		vecs[d.Index] = d.Embedding
	}
	return vecs, nil
}
