package directives

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/99designs/gqlgen/graphql"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/authentication"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/rest/constant"
)

func (d *Directives) AuthDirective(ctx context.Context, obj any, next graphql.Resolver) (res any, err error) {
	log := d.logger.With(slog.String("middleware", "Authenticator"))

	_, ok := ctx.Value(constant.ContextUser).(*authentication.UserData)
	if !ok {
		log.WarnContext(ctx, "not authenticated")
		return nil, fmt.Errorf("not authenticated")
	}

	return next(ctx)
}
