package database

import (
	"time"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/types/shared"
)

type ExpenseSortByField string

const (
	ExpenseSortByFieldUnspecified ExpenseSortByField = "unspecified"
	ExpenseSortByFieldID          ExpenseSortByField = "id"
	ExpenseSortByFieldAmount      ExpenseSortByField = "amount"
	ExpenseSortByFieldCategory    ExpenseSortByField = "category"
	ExpenseSortByFieldCreatedAt   ExpenseSortByField = "createdAt"
	ExpenseSortByFieldUpdatedAt   ExpenseSortByField = "updatedAt"
)

var ExpenseSortByFields = []ExpenseSortByField{
	ExpenseSortByFieldUnspecified,
	ExpenseSortByFieldID,
	ExpenseSortByFieldAmount,
	ExpenseSortByFieldCategory,
	ExpenseSortByFieldCreatedAt,
	ExpenseSortByFieldUpdatedAt,
}

type ListExpenseReqFilter interface {
	FilterAnd() []ListExpenseReqFilter
	FilterOr() []ListExpenseReqFilter
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
	shared.SortBy[ExpenseSortByField]
	ListExpenseReqFilter
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
