package governor

import (
	"context"
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/config"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/governor/auth"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/governor/expense"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/types/database"
)

type Governor struct {
	*auth.Auth
	*expense.Expense
}

func New(conf *config.Config) *Governor {
	return &Governor{
		Auth:    new(auth.Auth),
		Expense: new(expense.Expense),
	}
}

func (g *Governor) Config(ctx context.Context, conf *config.Config, logger *slog.Logger, db database.Database) {
	*g.Auth = *auth.New(conf, logger.With(slog.String("component", "auth")), db)
	*g.Expense = *expense.New(conf, logger.With(slog.String("component", "expense")), db)
}
