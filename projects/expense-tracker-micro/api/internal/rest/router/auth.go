package router

import (
	"context"
	"net/http"
)

func (r *Router) auth(_ context.Context) {
	r.router.Handle("POST /sign-up", http.HandlerFunc(r.handler.SignUp))
	r.router.Handle("POST /sign-in", http.HandlerFunc(r.handler.SignIn))
}
