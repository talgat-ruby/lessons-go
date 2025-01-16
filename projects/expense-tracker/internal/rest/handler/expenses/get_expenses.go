package expenses

import (
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/rest/pkg/filter"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/rest/pkg/httperror"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/types/controller"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/types/shared"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/pkg/httputils/request"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/pkg/httputils/response"
)

func (h *Expenses) GetExpenses(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "GetExpenses")

	// request parse
	reqBody := new(getExpensesReq)
	if err := request.JSON(w, r, reqBody); err != nil {
		log.ErrorContext(
			ctx,
			"failed to parse request body",
			slog.Any("error", err),
		)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	ctrlResp, err := h.ctrl.ListExpense(ctx, reqBody)
	if err != nil {
		log.ErrorContext(ctx, "fail", slog.Any("error", err))
		httperror.
			NewMessage("", "invalid values", "", "").
			HandleError(w, err)
		return
	}

	respBody := new(getExpensesResp)
	for _, it := range ctrlResp.GetList() {
		item := &getExpensesRespDataItem{
			ID:        it.GetID(),
			UserID:    it.GetUserID(),
			Amount:    it.GetAmount(),
			Category:  it.GetCategory(),
			CreatedAt: it.GetCreatedAt(),
			UpdatedAt: it.GetUpdatedAt(),
		}
		respBody.Data = append(respBody.Data, item)
	}

	if err := response.JSON(
		w,
		http.StatusOK,
		respBody,
	); err != nil {
		log.ErrorContext(
			ctx,
			"fail json",
			slog.Any("error", err),
		)
		return
	}

	log.InfoContext(
		ctx,
		"success",
	)
}

type getExpensesReq struct {
	Data *getExpensesReqData `json:"data"`
}

func (req *getExpensesReq) GetOffset() int {
	return req.Data.Pagination.Offset
}

func (req *getExpensesReq) GetLimit() int {
	return req.Data.Pagination.Limit
}

func (req *getExpensesReq) GetSortBy() []shared.SortByItem[controller.ExpenseSortByField] {
	result := make([]shared.SortByItem[controller.ExpenseSortByField], 0, len(req.Data.SortBy))
	for _, item := range req.Data.SortBy {
		result = append(result, item)
	}
	return result
}

func (req *getExpensesReq) FilterAnd() []controller.ListExpenseReqFilter {
	return req.Data.Filter.FilterAnd()
}

func (req *getExpensesReq) FilterOr() []controller.ListExpenseReqFilter {
	return req.Data.Filter.FilterOr()
}

func (req *getExpensesReq) FilterCreatedAt() shared.TimeExp {
	return req.Data.Filter.FilterCreatedAt()
}

func (req *getExpensesReq) FilterUpdatedAt() shared.TimeExp {
	return req.Data.Filter.FilterUpdatedAt()
}

func (req *getExpensesReq) FilterID() shared.IDExp {
	return req.Data.Filter.FilterID()
}

func (req *getExpensesReq) FilterUserID() shared.IDExp {
	return req.Data.Filter.FilterUserID()
}

func (req *getExpensesReq) FilterAmount() shared.Int64Exp {
	return req.Data.Filter.FilterAmount()
}

func (req *getExpensesReq) FilterCategory() shared.StringExp {
	return req.Data.Filter.FilterCategory()
}

type getExpensesReqData struct {
	Pagination *getExpensesPagination   `json:"pagination,omitempty"`
	SortBy     []*getExpensesSortByItem `json:"sortBy,omitempty"`
	Filter     *getExpensesFilter       `json:"filter,omitempty"`
}

type getExpensesPagination struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type getExpensesSortByItem struct {
	Direction string `json:"direction"`
	Field     string `json:"field"`
}

func (s *getExpensesSortByItem) GetDirection() shared.SortByDirection {
	for _, d := range shared.SortByDirections {
		if strings.EqualFold(s.Direction, string(d)) {
			return d
		}
	}

	return shared.SortByDirectionUnspecified
}

func (s *getExpensesSortByItem) GetField() controller.ExpenseSortByField {
	for _, f := range controller.ExpenseSortByFields {
		if strings.EqualFold(s.Field, string(f)) {
			return f
		}
	}

	return controller.ExpenseSortByFieldUnspecified
}

type getExpensesFilter struct {
	And       []*getExpensesFilter             `json:"and,omitempty"`
	Or        []*getExpensesFilter             `json:"or,omitempty"`
	CreatedAt *filter.ComparableExp[time.Time] `json:"createdAt,omitempty"`
	UpdatedAt *filter.ComparableExp[time.Time] `json:"updatedAt,omitempty"`
	ID        *filter.IDExp                    `json:"id,omitempty"`
	UserID    *filter.IDExp                    `json:"userId,omitempty"`
	Amount    *filter.ComparableExp[int64]     `json:"amount,omitempty"`
	Category  *filter.StringExp                `json:"category,omitempty"`
}

func (req *getExpensesFilter) FilterAnd() []controller.ListExpenseReqFilter {
	result := make([]controller.ListExpenseReqFilter, 0, len(req.And))
	for _, item := range req.And {
		result = append(result, item)
	}
	return result
}

func (req *getExpensesFilter) FilterOr() []controller.ListExpenseReqFilter {
	result := make([]controller.ListExpenseReqFilter, 0, len(req.Or))
	for _, item := range req.Or {
		result = append(result, item)
	}
	return result
}

func (req *getExpensesFilter) FilterCreatedAt() shared.TimeExp {
	return req.CreatedAt
}

func (req *getExpensesFilter) FilterUpdatedAt() shared.TimeExp {
	return req.UpdatedAt
}

func (req *getExpensesFilter) FilterID() shared.IDExp {
	return req.ID
}

func (req *getExpensesFilter) FilterUserID() shared.IDExp {
	return req.UserID
}

func (req *getExpensesFilter) FilterAmount() shared.Int64Exp {
	return req.Amount
}

func (req *getExpensesFilter) FilterCategory() shared.StringExp {
	return req.Category
}

type getExpensesResp struct {
	Data []*getExpensesRespDataItem `json:"data"`
}

type getExpensesRespDataItem struct {
	ID        string    `json:"id,omitempty"`
	UserID    string    `json:"userId,omitempty"`
	Amount    int64     `json:"amount,omitempty"`
	Category  string    `json:"category,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
