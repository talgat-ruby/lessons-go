package expense

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/authentication"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/rest/constant"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/types/controller"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/types/database"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/types/shared"
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

	_ = newLisExpenseDBReq(user.ID, req)
	dbResp, err := r.db.ListExpense(ctx, nil)
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
	userId  string
	ctrlReq controller.ListExpenseReq
}

func newLisExpenseDBReq(userId string, ctrlReq controller.ListExpenseReq) *listExpenseDBReq {
	return &listExpenseDBReq{
		userId:  userId,
		ctrlReq: ctrlReq,
	}
}

func (req *listExpenseDBReq) GetUserID() string {
	return req.userId
}

func (req *listExpenseDBReq) GetLimit() int {
	return req.ctrlReq.GetLimit()
}

func (req *listExpenseDBReq) GetOffset() int {
	return req.ctrlReq.GetOffset()
}

func (req *listExpenseDBReq) GetSortBy() []shared.SortByItem[database.ExpenseSortByField] {
	result := make([]shared.SortByItem[database.ExpenseSortByField], 0, len(req.ctrlReq.GetSortBy()))

	if req.ctrlReq == nil {
		return nil
	}

	for _, it := range req.ctrlReq.GetSortBy() {
		item := &listExpenseSortByItem{
			direction: it.GetDirection(),
			field:     database.ExpenseSortByField(it.GetField()),
		}
		result = append(result, item)
	}
	return result
}

// func (req *listExpenseDBReq) mapToDatabaseFilter(its []controller.ListExpenseReqFilter) []database.ListExpenseReqFilter {
// 	result := make([]database.ListExpenseReqFilter, 0, len(req.ctrlReq.FilterAnd()))
// 	for _, it := range req.ctrlReq.FilterAnd() {
// 		result = append(result, req.mapToDatabaseFilterItem(it))
// 	}
// 	return result
// }
//
// func (req *listExpenseDBReq) mapToDatabaseFilterItem(item controller.ListExpenseReqFilter) database.ListExpenseReqFilter {
// 	result := database.ListExpenseReqFilter(item)
// 	return result
// }
//
// func (req *listExpenseDBReq) FilterAnd() []database.ListExpenseReqFilter {
// 	result := make([]database.ListExpenseReqFilter, 0, len(req.ctrlReq.FilterAnd()))
// 	for _, it := range req.ctrlReq.FilterAnd() {
// 		result = append(result, database.ListExpenseReqFilter(it))
// 	}
// 	return result
// }

func (req *listExpenseDBReq) FilterCreatedAt() shared.TimeExp {
	return req.ctrlReq.FilterCreatedAt()
}

func (req *listExpenseDBReq) FilterUpdatedAt() shared.TimeExp {
	return req.ctrlReq.FilterUpdatedAt()
}

func (req *listExpenseDBReq) FilterID() shared.IDExp {
	return req.ctrlReq.FilterID()
}

func (req *listExpenseDBReq) FilterUserID() shared.IDExp {
	return req.ctrlReq.FilterUserID()
}

func (req *listExpenseDBReq) FilterAmount() shared.Int64Exp {
	return req.ctrlReq.FilterAmount()
}

func (req *listExpenseDBReq) FilterCategory() shared.StringExp {
	return req.ctrlReq.FilterCategory()
}

type listExpenseSortByItem struct {
	direction shared.SortByDirection
	field     database.ExpenseSortByField
}

func (s *listExpenseSortByItem) GetDirection() shared.SortByDirection {
	return s.direction
}

func (s *listExpenseSortByItem) GetField() database.ExpenseSortByField {
	return s.field
}

type listExpenseFilter struct {
	ctrl controller.ListExpenseReqFilter
}

// func newDatabaseFilter(ctrl controller.ListExpenseReqFilter) database.ListExpenseReqFilter {
// 	if ctrl == nil {
// 		return nil
// 	}
// 	return &listExpenseFilter{ctrl: ctrl}
// }
