package expense

import (
	"database/sql"
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/config"
)

type Expense struct {
	conf   *config.PostgresConfig
	logger *slog.Logger
	db     *sql.DB
}

func New(
	conf *config.PostgresConfig,
	logger *slog.Logger,
	db *sql.DB,
) *Expense {
	return &Expense{
		conf:   conf,
		logger: logger,
		db:     db,
	}
}
