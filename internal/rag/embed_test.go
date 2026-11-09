package rag

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func embedServer(t *testing.T, handler func(w http.ResponseWriter, input []string)) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/embeddings" || r.Header.Get("Authorization") != "Bearer k" {
			t.Errorf("got %s %s auth=%q", r.Method, r.URL.Path, r.Header.Get("Authorization"))
		}
		var req struct {
			Model string   `json:"model"`
			Input []string `json:"input"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Model != "m" {
			t.Errorf("bad request body: %v model=%q", err, req.Model)
		}
		handler(w, req.Input)
	}))
	t.Cleanup(srv.Close)
	return srv
}

func TestHTTPEmbedder(t *testing.T) {
	tests := []struct {
		name    string
		reply   string
		status  int
		wantErr string
		want    [][]float32
	}{
		{
			name:  "reorders by index",
			reply: `{"data":[{"index":1,"embedding":[3,4]},{"index":0,"embedding":[1,2]}]}`,
			want:  [][]float32{{1, 2}, {3, 4}},
		},
		{name: "wrong dimension", reply: `{"data":[{"index":0,"embedding":[1]},{"index":1,"embedding":[3,4]}]}`, wantErr: "got 1 dimensions, want 2"},
		{name: "wrong count", reply: `{"data":[{"index":0,"embedding":[1,2]}]}`, wantErr: "1 embeddings for 2 inputs"},
		{name: "duplicate index", reply: `{"data":[{"index":0,"embedding":[1,2]},{"index":0,"embedding":[1,2]}]}`, wantErr: "bad index 0"},
		{name: "http error", status: 429, reply: `{"error":"rate limited"}`, wantErr: "429"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := embedServer(t, func(w http.ResponseWriter, _ []string) {
				if tt.status != 0 {
					w.WriteHeader(tt.status)
				}
				_, _ = w.Write([]byte(tt.reply))
			})
			e := &HTTPEmbedder{BaseURL: srv.URL + "/v1/", APIKey: "k", Model: "m", Dimensions: 2}
			got, err := e.Embed(context.Background(), []string{"a", "b"})
			if tt.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("err = %v, want containing %q", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("Embed: %v", err)
			}
			for i := range tt.want {
				if got[i][0] != tt.want[i][0] || got[i][1] != tt.want[i][1] {
					t.Errorf("vec %d = %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestHTTPEmbedderEmptyInputMakesNoRequest(t *testing.T) {
	e := &HTTPEmbedder{BaseURL: "http://127.0.0.1:1", Model: "m", Dimensions: 2}
	got, err := e.Embed(context.Background(), nil)
	if err != nil || got != nil {
		t.Fatalf("Embed(nil) = %v, %v; want nil, nil", got, err)
	}
}
