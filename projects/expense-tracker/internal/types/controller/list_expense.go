package controller

type ExpenseSortByField string

const (
	ExpenseSortByFieldID        ExpenseSortByField = "id"
	ExpenseSortByFieldAmount    ExpenseSortByField = "amount"
	ExpenseSortByFieldCategory  ExpenseSortByField = "category"
	ExpenseSortByFieldCreatedAt ExpenseSortByField = "createdAt"
	ExpenseSortByFieldUpdatedAt ExpenseSortByField = "updatedAt"
)

type ListExpenseReq interface {
	GetPagination() Pagination
	GetSortBy() SortBy[ExpenseSortByField]
}

type ListExpenseResp interface {
	GetList() []ItemExpenseResp
}

type ItemExpenseResp interface {
	GetID() string
	GetUserID() string
	GetAmount() string
	GetCategory() string
	GetCreatedAt() string
	GetUpdatedAt() string
}
