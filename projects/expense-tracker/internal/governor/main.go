package governor

import (
	"context"
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/config"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/governor/auth"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/types/database"
)

type Governor struct {
	*auth.Auth
}

func New(conf *config.Config, logger *slog.Logger) *Governor {
	return &Governor{
		Auth: new(auth.Auth),
	}
}

func (g *Governor) Config(ctx context.Context, conf *config.Config, logger *slog.Logger, db database.Database) {
	val := validator.New(validator.WithRequiredStructEnabled())

	*g.Auth = *auth.New(conf, logger.With(slog.String("component", "auth")), db, val)
}
