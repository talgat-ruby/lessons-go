package expense

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/types/controller"
)

func (r *Expense) RemoveExpense(ctx context.Context, req controller.RemoveExpenseReq) (controller.RemoveExpenseResp, error) {
	log := r.logger.With(slog.String("handler", "RemoveExpense"))

	if req == nil {
		log.ErrorContext(ctx, "req is nil")
		return nil, fmt.Errorf("req is nil")
	}

	// get user
	// user, ok := ctx.Value("user").(*auth.UserData)
	// if !ok {
	// 	return nil, fmt.Errorf("user not found in context")
	// }

	dbReq := newDeleteExpenseDBReq(req.GetID())
	_, err := r.db.DeleteExpense(ctx, dbReq)
	if err != nil {
		log.ErrorContext(ctx, "db request failed", slog.Any("error", err))
		return nil, fmt.Errorf("db request failed %w", err)
	}

	log.InfoContext(
		ctx,
		"success",
	)
	return nil, nil
}

type deleteExpenseDBReq struct {
	id string
}

func newDeleteExpenseDBReq(id string) *deleteExpenseDBReq {
	return &deleteExpenseDBReq{
		id: id,
	}
}

func (req *deleteExpenseDBReq) GetID() string {
	return req.id
}
