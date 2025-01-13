package handler

import (
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/rest/handler/auth"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/rest/handler/expenses"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/types/controller"
)

type Handler struct {
	*auth.Auth
	*expenses.Expenses
}

func New(logger *slog.Logger, ctrl controller.Controller) *Handler {
	return &Handler{
		Auth:     auth.New(logger.With(slog.String("component", "auth")), ctrl),
		Expenses: expenses.New(logger.With(slog.String("component", "expenses")), ctrl),
	}
}
