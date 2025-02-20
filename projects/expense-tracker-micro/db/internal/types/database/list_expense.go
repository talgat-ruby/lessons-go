package database

import (
	"time"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/internal/constant"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/internal/types/shared"
)

type ListExpenseReqFilterItem interface {
	shared.FilterItem[ListExpenseReqFilterItem]
	FilterCreatedAt() shared.TimeExp
	FilterUpdatedAt() shared.TimeExp
	FilterID() shared.IDExp
	FilterUserID() shared.IDExp
	FilterAmount() shared.Int64Exp
	FilterCategory() shared.StringExp
}

type ListExpenseReq interface {
	GetUserID() string
	shared.PaginationOffset
	shared.SortBy[constant.ExpenseSortByField]
	shared.Filter[ListExpenseReqFilterItem]
}

type ListExpenseResp interface {
	GetList() []ItemExpenseResp
}

type ItemExpenseResp interface {
	GetID() string
	GetUserID() string
	GetAmount() int64
	GetCategory() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}
