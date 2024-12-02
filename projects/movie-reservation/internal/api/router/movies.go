package router

import (
	"context"
)

func (r *Router) movies(ctx context.Context) {
	r.router.HandleFunc("GET /movies", r.handler.FindMovies)
	r.router.HandleFunc("GET /movies/{id}", r.handler.FindMovie)
	r.router.HandleFunc("POST /movies", r.handler.CreateMovie)
	r.router.HandleFunc("PUT /movies/{id}", r.handler.UpdateMovie)
	r.router.HandleFunc("DELETE /movies/{id}", r.handler.DeleteMovie)
}
