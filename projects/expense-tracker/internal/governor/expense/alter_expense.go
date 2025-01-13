package expense

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/types/controller"
)

func (r *Expense) AlterExpense(ctx context.Context, req controller.AlterExpenseReq) (controller.AlterExpenseResp, error) {
	log := r.logger.With(slog.String("handler", "AlterExpense"))

	if req == nil {
		log.ErrorContext(ctx, "req is nil")
		return nil, fmt.Errorf("req is nil")
	}

	// get user
	// user, ok := ctx.Value("user").(*auth.UserData)
	// if !ok {
	// 	return nil, fmt.Errorf("user not found in context")
	// }

	dbReq := newUpdateExpenseDBReq(req.GetID(), req.GetAmount(), req.GetCategory())
	_, err := r.db.UpdateExpense(ctx, dbReq)
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

type updateExpenseDBReq struct {
	id       string
	amount   *int64
	category *string
}

func newUpdateExpenseDBReq(id string, amount *int64, category *string) *updateExpenseDBReq {
	return &updateExpenseDBReq{
		id:       id,
		amount:   amount,
		category: category,
	}
}

func (req *updateExpenseDBReq) GetID() string {
	return req.id
}

func (req *updateExpenseDBReq) GetAmount() *int64 {
	return req.amount
}

func (req *updateExpenseDBReq) GetCategory() *string {
	return req.category
}
