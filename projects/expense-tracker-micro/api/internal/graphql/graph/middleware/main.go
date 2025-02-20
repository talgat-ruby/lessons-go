package middleware

import (
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/types/controller"
)

type Middleware struct {
	logger *slog.Logger
	ctrl   controller.Controller
}

func New(
	logger *slog.Logger,
	ctrl controller.Controller,
) *Middleware {
	return &Middleware{
		logger: logger,
		ctrl:   ctrl,
	}
}
