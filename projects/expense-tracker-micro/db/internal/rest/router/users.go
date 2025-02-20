package router

import (
	"context"
	"net/http"
)

func (r *Router) users(_ context.Context) {
	r.router.Handle("POST /users", http.HandlerFunc(r.handler.PostUser))
	r.router.Handle("GET /users/{email}", http.HandlerFunc(r.handler.GetUser))
}
