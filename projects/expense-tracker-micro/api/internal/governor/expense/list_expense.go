package expense

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/authentication"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/constant"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/types/controller"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/types/database"
)

func (r *Expense) ListExpense(ctx context.Context, req controller.ListExpenseReq) (controller.ListExpenseResp, error) {
	log := r.logger.With(slog.String("handler", "ListExpense"))

	if req == nil {
		log.ErrorContext(ctx, "req is nil")
		return nil, fmt.Errorf("req is nil")
	}

	// get user
	user, ok := ctx.Value(constant.ContextUser).(*authentication.UserData)
	if !ok {
		return nil, fmt.Errorf("user not found in context")
	}

	dbReq := newListExpenseDBReq(user.ID, req)
	dbResp, err := r.db.ListExpense(ctx, dbReq)
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
	return nil, nil
}

type listExpenseDBReq struct {
	controller.ListExpenseReq
	userId string
}

func newListExpenseDBReq(userId string, r controller.ListExpenseReq) *listExpenseDBReq {
	return &listExpenseDBReq{
		ListExpenseReq: r,
		userId:         userId,
	}
}

func (req *listExpenseDBReq) GetUserID() string {
	return req.userId
}

func (req *listExpenseDBReq) GetFilter() database.ListExpenseReqFilterItem {
	return newDatabaseFilter(req.ListExpenseReq.GetFilter())
}

type listExpenseDBFilter struct {
	controller.ListExpenseReqFilterItem
}

func newDatabaseFilter(ctrl controller.ListExpenseReqFilterItem) *listExpenseDBFilter {
	if ctrl == nil {
		return nil
	}
	return &listExpenseDBFilter{ctrl}
}

func (req *listExpenseDBFilter) FilterAnd() []database.ListExpenseReqFilterItem {
	if req == nil || req.ListExpenseReqFilterItem == nil || req.ListExpenseReqFilterItem.FilterAnd == nil || req.ListExpenseReqFilterItem.FilterAnd() == nil {
		return nil
	}

	result := make([]database.ListExpenseReqFilterItem, 0, len(req.ListExpenseReqFilterItem.FilterAnd()))
	for _, f := range req.ListExpenseReqFilterItem.FilterAnd() {
		result = append(result, newDatabaseFilter(f))
	}
	return result
}

func (req *listExpenseDBFilter) FilterOr() []database.ListExpenseReqFilterItem {
	if req == nil || req.ListExpenseReqFilterItem == nil || req.ListExpenseReqFilterItem.FilterOr == nil || req.ListExpenseReqFilterItem.FilterOr() == nil {
		return nil
	}

	result := make([]database.ListExpenseReqFilterItem, 0, len(req.ListExpenseReqFilterItem.FilterOr()))
	for _, f := range req.ListExpenseReqFilterItem.FilterOr() {
		result = append(result, newDatabaseFilter(f))
	}
	return result
}
