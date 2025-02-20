package users

import (
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/internal/types/database"
)

type Users struct {
	logger *slog.Logger
	db     database.Database
}

func New(logger *slog.Logger, db database.Database) *Users {
	return &Users{
		logger: logger,
		db:     db,
	}
}
