package constant

type ExpenseSortByField string

const (
	ExpenseSortByFieldUnspecified ExpenseSortByField = "unspecified"
	ExpenseSortByFieldID          ExpenseSortByField = "id"
	ExpenseSortByFieldAmount      ExpenseSortByField = "amount"
	ExpenseSortByFieldCategory    ExpenseSortByField = "category"
	ExpenseSortByFieldCreatedAt   ExpenseSortByField = "createdAt"
	ExpenseSortByFieldUpdatedAt   ExpenseSortByField = "updatedAt"
)
