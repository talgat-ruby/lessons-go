package expense

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/authentication"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/constant"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/types/controller"
)

func (r *Expense) NewExpense(ctx context.Context, req controller.NewExpenseReq) (controller.NewExpenseResp, error) {
	log := r.logger.With(slog.String("handler", "NewExpense"))

	if req == nil {
		log.ErrorContext(ctx, "req is nil")
		return nil, fmt.Errorf("req is nil")
	}

	// get user
	user, ok := ctx.Value(constant.ContextUser).(*authentication.UserData)
	if !ok {
		return nil, fmt.Errorf("user not found in context")
	}

	dbReq := newCreateExpenseDBReq(user.ID, req.GetAmount(), req.GetCategory())
	dbResp, err := r.db.CreateExpense(ctx, dbReq)
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

type createExpenseDBReq struct {
	userId   string
	amount   int64
	category string
}

func newCreateExpenseDBReq(userId string, amount int64, category string) *createExpenseDBReq {
	return &createExpenseDBReq{
		userId:   userId,
		amount:   amount,
		category: category,
	}
}

func (req *createExpenseDBReq) GetUserID() string {
	return req.userId
}

func (req *createExpenseDBReq) GetAmount() int64 {
	return req.amount
}

func (req *createExpenseDBReq) GetCategory() string {
	return req.category
}
