package interceptor

import (
	"context"
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/authentication"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/constant"
)

func (i *Interceptor) Authenticator(ctx context.Context, tokenString string) (context.Context, error) {
	log := i.logger.With(slog.String("interceptor", "Authenticator"))

	userData, err := authentication.ParseToken(tokenString, i.conf.API.TokenSecret)
	if err != nil {
		log.ErrorContext(
			ctx,
			"fail authentication",
			slog.Any("error", err),
		)
		return nil, err
	}

	newCtx := context.WithValue(ctx, constant.ContextUser, userData)
	return newCtx, nil
}
