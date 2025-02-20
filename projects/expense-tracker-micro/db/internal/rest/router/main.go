package router

import (
	"context"
	"net/http"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/internal/rest/handler"
)

type Router struct {
	router  *http.ServeMux
	handler *handler.Handler
}

func New(handler *handler.Handler) *Router {
	mux := http.NewServeMux()

	return &Router{
		router:  mux,
		handler: handler,
	}
}

func (r *Router) Start(ctx context.Context) *http.ServeMux {
	r.users(ctx)

	return r.router
}
