package expense

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/authentication"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/constant"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/types/controller"
)

func (r *Expense) RemoveExpense(ctx context.Context, req controller.RemoveExpenseReq) (controller.RemoveExpenseResp, error) {
	log := r.logger.With(slog.String("handler", "RemoveExpense"))

	if req == nil {
		log.ErrorContext(ctx, "req is nil")
		return nil, fmt.Errorf("req is nil")
	}

	// get user
	user, ok := ctx.Value(constant.ContextUser).(*authentication.UserData)
	if !ok {
		log.ErrorContext(ctx, "user not found in context")
		return nil, fmt.Errorf("user not found in context")
	}

	dbReq := newDeleteExpenseDBReq(user.ID, req.GetID())
	dbResp, err := r.db.DeleteExpense(ctx, dbReq)
	if err != nil {
		log.ErrorContext(ctx, "db request failed", slog.Any("error", err))
		return nil, fmt.Errorf("db request failed %w", err)
	}
	if dbResp == nil {
		return nil, nil
	}

	log.InfoContext(
		ctx,
		"success",
	)
	return true, nil
}

type deleteExpenseDBReq struct {
	userId string
	id     string
}

func newDeleteExpenseDBReq(userId, id string) *deleteExpenseDBReq {
	return &deleteExpenseDBReq{
		userId: userId,
		id:     id,
	}
}

func (req *deleteExpenseDBReq) GetUserID() string {
	return req.userId
}

func (req *deleteExpenseDBReq) GetID() string {
	return req.id
}
