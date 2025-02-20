package governor

import (
	"context"
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/config"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/governor/auth"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/governor/expense"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/governor/interceptor"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/types/database"
)

type Governor struct {
	*interceptor.Interceptor
	*auth.Auth
	*expense.Expense
}

func New(conf *config.Config) *Governor {
	return &Governor{
		Interceptor: new(interceptor.Interceptor),
		Auth:        new(auth.Auth),
		Expense:     new(expense.Expense),
	}
}

func (g *Governor) Config(ctx context.Context, conf *config.Config, logger *slog.Logger, db database.Database) {
	*g.Interceptor = *interceptor.New(conf, logger.With(slog.String("module", "interceptor")))
	*g.Auth = *auth.New(conf, logger.With(slog.String("component", "auth")), db)
	*g.Expense = *expense.New(conf, logger.With(slog.String("component", "expense")), db)
}
