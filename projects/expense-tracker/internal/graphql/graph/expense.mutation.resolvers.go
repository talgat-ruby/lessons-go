package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.63

import (
	"context"
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/graphql/graph/model"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/graphql/graph/request"
)

// CreateExpense is the resolver for the createExpense field.
func (r *mutationResolver) CreateExpense(ctx context.Context, amount int32, category model.Category) (bool, error) {
	log := r.logger.With("method", "PostExpense")

	// request parse
	reqBody := request.NewCtrlCreateExpenseRequest(int64(amount), string(category))
	_, err := r.ctrl.NewExpense(ctx, reqBody)
	if err != nil {
		log.ErrorContext(ctx, "fail", slog.Any("error", err))
		return false, err
	}

	log.InfoContext(
		ctx,
		"success",
	)
	return true, nil
}
