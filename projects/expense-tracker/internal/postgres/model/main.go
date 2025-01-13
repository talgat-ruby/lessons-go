package model

import (
	"database/sql"
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/config"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/postgres/model/auth"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/postgres/model/expense"
)

type Model struct {
	*auth.Auth
	*expense.Expense
}

func New(
	conf *config.PostgresConfig,
	logger *slog.Logger,
	db *sql.DB,
) *Model {
	return &Model{
		Auth:    auth.New(conf, logger.With(slog.String("component", "auth")), db),
		Expense: expense.New(conf, logger.With(slog.String("component", "expense")), db),
	}
}
