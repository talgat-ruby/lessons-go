package interceptor

import (
	"context"
	"log/slog"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (inter *Interceptor) authenticate(ctx context.Context) (context.Context, error) {
	log := inter.log.With(slog.String("interceptor", "authenticate"))

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.ErrorContext(
			ctx,
			"metadata is not provided",
		)
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		log.ErrorContext(
			ctx,
			"authorization header is empty",
		)
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	if !strings.HasPrefix(authHeader[0], "Bearer ") {
		log.ErrorContext(
			ctx,
			"invalid authorization header",
		)
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	tokenString := authHeader[0][len("Bearer "):]
	newCtx, err := inter.ctrl.Authenticator(ctx, tokenString)
	if err != nil {
		log.ErrorContext(
			ctx,
			"fail authentication",
			"error", err,
		)
		return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
	}

	return newCtx, nil
}
