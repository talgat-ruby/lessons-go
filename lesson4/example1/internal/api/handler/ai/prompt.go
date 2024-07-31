package ai

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/talgat-ruby/lessons-go/lesson4/example1/internal/gemini"
	"github.com/talgat-ruby/lessons-go/lesson4/example1/pkg/response"
)

type promptReqBody struct {
	Prompt string `json:"prompt"`
}

func (h *Handler) Prompt(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	slog.InfoContext(ctx, "aiHandler::Prompt::start")

	var b promptReqBody
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		slog.ErrorContext(
			ctx,
			"aiHandler::Prompt::fail",
			"error", err,
		)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	g := gemini.NewGemini(h.apiKey)
	text, err := g.SendPrompt(b.Prompt)
	if err != nil {
		slog.ErrorContext(
			ctx,
			"aiHandler::Prompt::fail",
			"error", err,
		)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data := &promptResponseBody{
		Message: text,
	}
	if err := response.JSON(w, data); err != nil {
		slog.ErrorContext(
			ctx,
			"aiHandler::Prompt::fail",
			"error", err,
		)
		return
	}
}
