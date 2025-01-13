package expense

import (
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/config"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/types/database"
)

type Expense struct {
	conf   *config.Config
	logger *slog.Logger
	db     database.Database
}

func New(conf *config.Config, logger *slog.Logger, db database.Database) *Expense {
	return &Expense{
		conf:   conf,
		logger: logger,
		db:     db,
	}
}
