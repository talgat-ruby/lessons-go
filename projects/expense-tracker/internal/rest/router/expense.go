package router

import (
	"context"
	"net/http"
)

func (r *Router) expense(_ context.Context) {
	// r.router.Handle("GET /expenses", http.HandlerFunc(r.handler.GetExpenses))
	r.router.Handle("POST /expenses", http.HandlerFunc(r.handler.PostExpenses))
	r.router.Handle("PATCH /expenses/{id}", http.HandlerFunc(r.handler.PatchExpense))
	r.router.Handle("DELETE /expenses/{id}", http.HandlerFunc(r.handler.DeleteExpense))
}
