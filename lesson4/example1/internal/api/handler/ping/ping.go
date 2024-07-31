package pingHandler

import (
	"log/slog"
	"net/http"

	"github.com/talgat-ruby/lessons-go/lesson4/example1/pkg/response"
)

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	slog.InfoContext(ctx, "pingHandler::Ping::start")

	data := &pingResponseBody{
		Message: "pong",
	}
	if err := response.JSON(w, data); err != nil {
		slog.ErrorContext(
			ctx,
			"pingHandler::Ping::fail",
			"error", err,
		)
		return
	}
}
