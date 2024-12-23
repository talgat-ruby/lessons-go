package router

import (
	"context"
	"net/http"

	"github.com/talgat-ruby/lessons-go/projects/movie-reservation/internal/api/handler"
	"github.com/talgat-ruby/lessons-go/projects/movie-reservation/internal/api/middleware"
)

type Router struct {
	router  *http.ServeMux
	handler *handler.Handler
	midd    *middleware.Middleware
}

func New(handler *handler.Handler, midd *middleware.Middleware) *Router {
	mux := http.NewServeMux()

	return &Router{
		router:  mux,
		handler: handler,
		midd:    midd,
	}
}

func (r *Router) Start(ctx context.Context) *http.ServeMux {
	r.auth(ctx)
	r.movies(ctx)

	return r.router
}
