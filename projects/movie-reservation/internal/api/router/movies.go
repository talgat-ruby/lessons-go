package router

import (
	"context"
)

func (r *Router) movies(ctx context.Context) {
	r.router.HandleFunc("GET /movies", r.handler.FindMovies)
}
