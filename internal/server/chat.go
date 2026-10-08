package server

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-chi/chi/v5/middleware"
)

const (
	maxQuestionRunes = 4000
	maxBodyBytes     = 64 << 10 // 64 KiB cap on the raw request body
	// stubAnswer is a fixed placeholder: it calls no LLM and asserts no theory,
	// honoring the "never bypass RAG" guardrail. Real answers arrive in Phase 2.
	stubAnswer = "The answer pipeline isn't wired up yet — grounded answers arrive in Phase 2."
)

var uuidRE = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)

type chatRequest struct {
	ConversationID *string `json:"conversation_id"`
	Question       string  `json:"question"`
}

type source struct {
	N       int     `json:"n"`
	Title   string  `json:"title"`
	Section string  `json:"section"`
	Score   float64 `json:"score"`
}

type chatResponse struct {
	ConversationID string   `json:"conversation_id"`
	Answer         string   `json:"answer"`
	Sources        []source `json:"sources"`
	TraceID        string   `json:"trace_id"`
}

func (a *api) handleChat(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)
	var req chatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		var tooBig *http.MaxBytesError
		if errors.As(err, &tooBig) {
			writeError(w, http.StatusBadRequest, "question_too_long", "request body exceeds the maximum size")
			return
		}
		writeError(w, http.StatusBadRequest, "invalid_json", "request body is not valid JSON")
		return
	}

	question := strings.TrimSpace(req.Question)
	if question == "" {
		writeError(w, http.StatusBadRequest, "invalid_request", "question is required")
		return
	}
	if len([]rune(question)) > maxQuestionRunes {
		writeError(w, http.StatusBadRequest, "question_too_long", "question exceeds the maximum length")
		return
	}

	id := ""
	if req.ConversationID != nil && *req.ConversationID != "" {
		if !uuidRE.MatchString(*req.ConversationID) {
			writeError(w, http.StatusBadRequest, "invalid_conversation_id", "conversation_id must be a UUID")
			return
		}
		id = strings.ToLower(*req.ConversationID)
	}

	ctx := r.Context()
	convID, err := a.chat.EnsureConversation(ctx, id)
	if err != nil {
		slog.Error("ensure conversation", "err", err,
			"request_id", middleware.GetReqID(ctx), "conversation_id", id)
		writeError(w, http.StatusInternalServerError, "internal", "could not process the request")
		return
	}
	if err := a.chat.AddMessage(ctx, convID, "user", question); err != nil {
		slog.Error("add user message", "err", err,
			"request_id", middleware.GetReqID(ctx), "conversation_id", convID, "role", "user")
		writeError(w, http.StatusInternalServerError, "internal", "could not process the request")
		return
	}
	if err := a.chat.AddMessage(ctx, convID, "assistant", stubAnswer); err != nil {
		slog.Error("add assistant message", "err", err,
			"request_id", middleware.GetReqID(ctx), "conversation_id", convID, "role", "assistant")
		writeError(w, http.StatusInternalServerError, "internal", "could not process the request")
		return
	}

	writeJSON(w, http.StatusOK, chatResponse{
		ConversationID: convID,
		Answer:         stubAnswer,
		Sources:        []source{},
		TraceID:        "",
	})
}
