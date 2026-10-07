package config

import (
	"testing"
)

// minimalEnv is the smallest set of variables that makes Load succeed.
func minimalEnv() map[string]string {
	return map[string]string{
		"DATABASE_URL":         "postgres://app:app@localhost:5432/theory",
		"LLM_BASE_URL":         "https://openrouter.ai/api/v1",
		"LLM_API_KEY":          "sk-test",
		"LLM_CHAT_MODEL":       "org/small-instruct",
		"LLM_EMBED_MODEL":      "text-embedding-3-small",
		"EMBEDDING_DIMENSIONS": "1536",
	}
}

func setEnv(t *testing.T, env map[string]string) {
	t.Helper()
	for k, v := range env {
		t.Setenv(k, v)
	}
}

func TestLoad(t *testing.T) {
	tests := []struct {
		name    string
		env     map[string]string
		wantErr bool
		check   func(t *testing.T, c Config)
	}{
		{
			name:    "minimal valid config applies defaults",
			env:     minimalEnv(),
			wantErr: false,
			check: func(t *testing.T, c Config) {
				if c.Addr != ":8080" {
					t.Errorf("Addr = %q, want :8080", c.Addr)
				}
				if c.RetrievalTopK != 5 {
					t.Errorf("RetrievalTopK = %d, want 5", c.RetrievalTopK)
				}
				if c.RetrievalMinScore != 0.5 {
					t.Errorf("RetrievalMinScore = %v, want 0.5", c.RetrievalMinScore)
				}
				if c.HistoryMessageLimit != 10 {
					t.Errorf("HistoryMessageLimit = %d, want 10", c.HistoryMessageLimit)
				}
				if c.RewriteHistoryMessages != 4 {
					t.Errorf("RewriteHistoryMessages = %d, want 4", c.RewriteHistoryMessages)
				}
				if c.RewriteMaxTokens != 64 {
					t.Errorf("RewriteMaxTokens = %d, want 64", c.RewriteMaxTokens)
				}
				if c.OnlineEvalSampleRate != 1.0 {
					t.Errorf("OnlineEvalSampleRate = %v, want 1.0", c.OnlineEvalSampleRate)
				}
				if c.DDSite != "datadoghq.com" {
					t.Errorf("DDSite = %q, want datadoghq.com", c.DDSite)
				}
				if c.DDLLMObsMLApp != "guitar-chat" {
					t.Errorf("DDLLMObsMLApp = %q, want guitar-chat", c.DDLLMObsMLApp)
				}
			},
		},
		{
			name: "judge model defaults to chat model",
			env:  minimalEnv(),
			check: func(t *testing.T, c Config) {
				if c.JudgeModel != c.LLMChatModel {
					t.Errorf("JudgeModel = %q, want fallback to LLMChatModel %q", c.JudgeModel, c.LLMChatModel)
				}
			},
		},
		{
			name: "missing required vars errors",
			env: map[string]string{
				"LLM_BASE_URL": "https://openrouter.ai/api/v1",
			},
			wantErr: true,
		},
		{
			name:    "non-numeric embedding dimensions errors",
			env:     withOverride(minimalEnv(), "EMBEDDING_DIMENSIONS", "notanumber"),
			wantErr: true,
		},
		{
			name:    "out-of-range sample rate errors",
			env:     withOverride(minimalEnv(), "ONLINE_EVAL_SAMPLE_RATE", "2.5"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setEnv(t, tt.env)
			c, err := Load()
			if tt.wantErr {
				if err == nil {
					t.Fatal("Load() succeeded, want error")
				}
				return
			}
			if err != nil {
				t.Fatalf("Load() error = %v", err)
			}
			if tt.check != nil {
				tt.check(t, c)
			}
		})
	}
}

func TestEmbedBaseURL(t *testing.T) {
	c := Config{LLMBaseURL: "https://chat", LLMEmbedBaseURL: ""}
	if got := c.EmbedBaseURL(); got != "https://chat" {
		t.Errorf("EmbedBaseURL() = %q, want fallback to chat URL", got)
	}
	c.LLMEmbedBaseURL = "https://embed"
	if got := c.EmbedBaseURL(); got != "https://embed" {
		t.Errorf("EmbedBaseURL() = %q, want embed URL", got)
	}
}

func TestObservabilityEnabled(t *testing.T) {
	c := Config{}
	if c.ObservabilityEnabled() {
		t.Error("ObservabilityEnabled() = true with no DD config, want false")
	}
	c = Config{DDApiKey: "k", DDSite: "datadoghq.com", DDLLMObsMLApp: "guitar-chat"}
	if !c.ObservabilityEnabled() {
		t.Error("ObservabilityEnabled() = false with full DD config, want true")
	}
}

func withOverride(env map[string]string, key, val string) map[string]string {
	out := make(map[string]string, len(env)+1)
	for k, v := range env {
		out[k] = v
	}
	out[key] = val
	return out
}
