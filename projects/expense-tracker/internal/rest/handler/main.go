package handler

import (
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/rest/handler/auth"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/types/controller"
)

type Handler struct {
	*auth.Auth
}

func New(logger *slog.Logger, ctrl controller.Controller) *Handler {
	return &Handler{
		Auth: auth.New(logger.With(slog.String("component", "auth")), ctrl),
	}
}
