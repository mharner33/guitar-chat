// Command eval runs the offline eval harness over evals/golden.jsonl: it runs
// the full RAG pipeline per golden row, scores with an LLM-as-judge, records one
// eval_runs row, and exits non-zero when the quality gate fails.
//
// Phase 0: loads config only. The harness is built in Phase 4 (see docs/plan.md).
package main

import (
	"log/slog"
	"os"

	"github.com/mharner33/guitar-chat/internal/config"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	if _, err := config.Load(); err != nil {
		slog.Error("eval failed", "err", err)
		os.Exit(1)
	}

	slog.Info("eval: config loaded; harness added in Phase 4")
}
