package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"
)

func (m *Middleware) Authenticator(
	next http.Handler,
) http.Handler {
	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := m.logger.With(slog.String("middleware", "Authenticator"))

		authorizationHeader := r.Header.Get("Authorization")
		fmt.Println("authorizationHeader", authorizationHeader)
		if authorizationHeader == "" {
			next.ServeHTTP(w, r)
			return
		}

		if !strings.HasPrefix(authorizationHeader, "Bearer ") {
			log.ErrorContext(
				ctx,
				"invalid authorization header",
			)
			http.Error(
				w,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized,
			)
			return
		}

		tokenString := authorizationHeader[len("Bearer "):]

		newCtx, err := m.ctrl.Authenticator(ctx, tokenString)
		if err != nil {
			log.ErrorContext(
				ctx,
				"fail authentication",
				slog.Any("error", err),
			)
			http.Error(
				w,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized,
			)
			return
		}

		next.ServeHTTP(w, r.WithContext(newCtx))
	}

	return http.HandlerFunc(h)
}
