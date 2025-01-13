package expense

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/postgres/generated/expenses/public/model"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/postgres/generated/expenses/public/table"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/types/database"
)

func (m *Expense) CreateExpense(ctx context.Context, req database.CreateExpenseReq) (database.CreateExpenseResp, error) {
	log := m.logger.With(slog.String("handler", "CreateExpense"))

	if req == nil {
		log.ErrorContext(ctx, "req is nil")
		return nil, fmt.Errorf("req is nil")
	}

	userID, err := strconv.Atoi(req.GetUserID())
	if err != nil {
		return nil, fmt.Errorf("invalid user id %s", req.GetUserID())
	}

	mdl := model.Expense{
		UserID:   int32(userID),
		Amount:   req.GetAmount(),
		Category: mapStringToModelCategory(req.GetCategory()),
	}
	stmt := table.Expense.
		INSERT(table.Expense.UserID, table.Expense.Amount, table.Expense.Category).
		MODEL(mdl).
		RETURNING(table.Expense.AllColumns)
	if _, err := stmt.ExecContext(ctx, m.db); err != nil {
		switch {
		default:
			log.ErrorContext(ctx, "failed to insert expense", slog.Any("error", err))
			return nil, fmt.Errorf("failed to insert expense: %w", err)
		}
	}

	log.InfoContext(
		ctx,
		"success",
	)
	return nil, nil
}
