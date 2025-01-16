package router

import (
	"context"
	"net/http"
)

func (r *Router) expense(_ context.Context) {
	r.router.Handle("GET /expenses", r.midd.Authenticator(http.HandlerFunc(r.handler.GetExpenses)))
	r.router.Handle("POST /expenses", r.midd.Authenticator(http.HandlerFunc(r.handler.PostExpense)))
	r.router.Handle("PATCH /expenses/{id}", r.midd.Authenticator(http.HandlerFunc(r.handler.PatchExpense)))
	r.router.Handle("DELETE /expenses/{id}", r.midd.Authenticator(http.HandlerFunc(r.handler.DeleteExpense)))
}
