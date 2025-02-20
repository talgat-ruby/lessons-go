package expense

import (
	"slices"
	"strings"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/postgres/generated/expenses/public/model"
)

func mapStringToModelCategory(category string) model.ExpenseCategory {
	ind := slices.IndexFunc(model.ExpenseCategoryAllValues, func(v model.ExpenseCategory) bool {
		return strings.EqualFold(string(v), category)
	})

	if ind < 0 {
		return model.ExpenseCategory_Other
	}
	return model.ExpenseCategoryAllValues[ind]
}
