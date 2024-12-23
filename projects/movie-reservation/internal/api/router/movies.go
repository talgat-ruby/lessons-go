package router

import (
	"context"
	"net/http"
)

func (r *Router) movies(ctx context.Context) {
	r.router.Handle("GET /movies", http.HandlerFunc(r.handler.FindMovies))
	r.router.Handle("GET /movies/{id}", http.HandlerFunc(r.handler.FindMovie))
	r.router.Handle("POST /movies", r.midd.Authenticator(http.HandlerFunc(r.handler.CreateMovie)))
	r.router.Handle("PUT /movies/{id}", http.HandlerFunc(r.handler.UpdateMovie))
	r.router.Handle("DELETE /movies/{id}", http.HandlerFunc(r.handler.DeleteMovie))
}
