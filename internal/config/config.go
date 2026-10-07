// Package config loads all application configuration from environment
// variables exactly once into a single struct. Per the project conventions
// (see agent.md), no other package reads the environment directly — they take
// the values they need from a Config.
package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Config holds every tunable the application reads. It is populated once by
// Load and then passed down explicitly.
type Config struct {
	// HTTP
	Addr string // listen address for cmd/api, e.g. ":8080"

	// Postgres
	DatabaseURL string

	// LLM / embeddings (OpenAI-compatible REST).
	LLMBaseURL      string // chat base URL, e.g. https://openrouter.ai/api/v1
	LLMEmbedBaseURL string // embeddings base URL; empty means use LLMBaseURL
	LLMAPIKey       string
	LLMChatModel    string // answer + rewrite model
	LLMEmbedModel   string // embedding model; output width must equal EmbeddingDimensions

	// EmbeddingDimensions is the vector width. It must match the vector(N)
	// column in the applied migration; api and ingest refuse to start on a
	// mismatch.
	EmbeddingDimensions int

	// Retrieval.
	RetrievalTopK     int     // chunks retrieved per query, before the score floor
	RetrievalMinScore float64 // minimum cosine similarity; below this a chunk is dropped

	// History / follow-up rewriting.
	HistoryMessageLimit    int // messages from the conversation passed to the answer model
	RewriteHistoryMessages int // messages passed to rewrite.query (only when history is non-empty)
	RewriteMaxTokens       int // max tokens for rewrite.query (temperature is 0)

	// Evaluation.
	JudgeModel           string  // model used for LLM-as-judge
	OnlineEvalSampleRate float64 // fraction of generated answers judged (0..1)

	// Datadog Agent Observability. These are optional: when absent the app
	// degrades to logs and never crashes on telemetry failures.
	DDApiKey                 string
	DDSite                   string // e.g. datadoghq.com (US1)
	DDLLMObsMLApp            string // app name in Agent Observability
	DDLLMObsAgentlessEnabled bool
}

// ObservabilityEnabled reports whether enough Datadog configuration is present
// to attempt span export. Callers use this to degrade gracefully to logs.
func (c Config) ObservabilityEnabled() bool {
	return c.DDApiKey != "" && c.DDSite != "" && c.DDLLMObsMLApp != ""
}

// EmbedBaseURL returns the embeddings base URL, falling back to the chat base
// URL when LLM_EMBED_BASE_URL is unset.
func (c Config) EmbedBaseURL() string {
	if c.LLMEmbedBaseURL != "" {
		return c.LLMEmbedBaseURL
	}
	return c.LLMBaseURL
}

// Load reads configuration from the environment. It returns an error listing
// every problem found rather than failing on the first, so a misconfigured
// deployment surfaces all missing variables at once.
func Load() (Config, error) {
	var errs []string
	req := func(key string) string {
		v := strings.TrimSpace(os.Getenv(key))
		if v == "" {
			errs = append(errs, fmt.Sprintf("%s is required", key))
		}
		return v
	}
	optInt := func(key string, def int) int {
		v := strings.TrimSpace(os.Getenv(key))
		if v == "" {
			return def
		}
		n, err := strconv.Atoi(v)
		if err != nil {
			errs = append(errs, fmt.Sprintf("%s must be an integer: %v", key, err))
			return def
		}
		return n
	}
	reqInt := func(key string) int {
		v := strings.TrimSpace(os.Getenv(key))
		if v == "" {
			errs = append(errs, fmt.Sprintf("%s is required", key))
			return 0
		}
		n, err := strconv.Atoi(v)
		if err != nil {
			errs = append(errs, fmt.Sprintf("%s must be an integer: %v", key, err))
			return 0
		}
		return n
	}
	optFloat := func(key string, def float64) float64 {
		v := strings.TrimSpace(os.Getenv(key))
		if v == "" {
			return def
		}
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			errs = append(errs, fmt.Sprintf("%s must be a number: %v", key, err))
			return def
		}
		return f
	}
	optStr := func(key, def string) string {
		v := strings.TrimSpace(os.Getenv(key))
		if v == "" {
			return def
		}
		return v
	}
	optBool := func(key string, def bool) bool {
		v := strings.TrimSpace(os.Getenv(key))
		if v == "" {
			return def
		}
		b, err := strconv.ParseBool(v)
		if err != nil {
			errs = append(errs, fmt.Sprintf("%s must be a boolean: %v", key, err))
			return def
		}
		return b
	}

	c := Config{
		Addr: optStr("HTTP_ADDR", ":8080"),

		DatabaseURL: req("DATABASE_URL"),

		LLMBaseURL:      req("LLM_BASE_URL"),
		LLMEmbedBaseURL: optStr("LLM_EMBED_BASE_URL", ""),
		LLMAPIKey:       req("LLM_API_KEY"),
		LLMChatModel:    req("LLM_CHAT_MODEL"),
		LLMEmbedModel:   req("LLM_EMBED_MODEL"),

		EmbeddingDimensions: reqInt("EMBEDDING_DIMENSIONS"),

		RetrievalTopK:     optInt("RETRIEVAL_TOP_K", 5),
		RetrievalMinScore: optFloat("RETRIEVAL_MIN_SCORE", 0.5),

		HistoryMessageLimit:    optInt("HISTORY_MESSAGE_LIMIT", 10),
		RewriteHistoryMessages: optInt("REWRITE_HISTORY_MESSAGES", 4),
		RewriteMaxTokens:       optInt("REWRITE_MAX_TOKENS", 64),

		JudgeModel:           optStr("JUDGE_MODEL", ""),
		OnlineEvalSampleRate: optFloat("ONLINE_EVAL_SAMPLE_RATE", 1.0),

		DDApiKey:                 optStr("DD_API_KEY", ""),
		DDSite:                   optStr("DD_SITE", "datadoghq.com"),
		DDLLMObsMLApp:            optStr("DD_LLMOBS_ML_APP", "guitar-chat"),
		DDLLMObsAgentlessEnabled: optBool("DD_LLMOBS_AGENTLESS_ENABLED", true),
	}

	// JUDGE_MODEL defaults to the chat model when unset.
	if c.JudgeModel == "" {
		c.JudgeModel = c.LLMChatModel
	}

	if c.EmbeddingDimensions < 0 {
		errs = append(errs, "EMBEDDING_DIMENSIONS must be positive")
	}
	if c.OnlineEvalSampleRate < 0 || c.OnlineEvalSampleRate > 1 {
		errs = append(errs, "ONLINE_EVAL_SAMPLE_RATE must be between 0 and 1")
	}

	if len(errs) > 0 {
		return Config{}, fmt.Errorf("invalid configuration:\n  - %s", strings.Join(errs, "\n  - "))
	}
	return c, nil
}
