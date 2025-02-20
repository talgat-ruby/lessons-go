package handler

import (
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/internal/rest/handler/users"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/internal/types/database"
)

type Handler struct {
	*users.Users
}

func New(logger *slog.Logger, db database.Database) *Handler {
	return &Handler{
		Users: users.New(logger.With(slog.String("component", "users")), db),
	}
}
