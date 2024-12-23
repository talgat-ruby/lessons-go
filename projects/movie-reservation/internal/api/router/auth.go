package router

import (
	"context"
)

func (r *Router) auth(ctx context.Context) {
	r.router.HandleFunc("POST /register", r.handler.Register)
	r.router.HandleFunc("POST /login", r.handler.Login)

	r.router.HandleFunc("POST /access-token", r.handler.AccessToken)
}
