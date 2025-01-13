package expense

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"

	pg "github.com/go-jet/jet/v2/postgres"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/postgres/generated/expenses/public/table"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/types/database"
)

func (m *Expense) DeleteExpense(ctx context.Context, req database.DeleteExpenseReq) (database.DeleteExpenseResp, error) {
	log := m.logger.With(slog.String("handler", "DeleteExpense"))

	if req == nil {
		log.ErrorContext(ctx, "req is nil")
		return nil, fmt.Errorf("req is nil")
	}

	id, err := strconv.Atoi(req.GetID())
	if err != nil {
		return nil, fmt.Errorf("invalid expense id %s", req.GetID())
	}

	stmt := table.Expense.
		DELETE().
		WHERE(table.Expense.ID.EQ(pg.Int(int64(id))))
	if _, err := stmt.ExecContext(ctx, m.db); err != nil {
		switch {
		default:
			log.ErrorContext(ctx, "failed to delete expense", slog.Any("error", err))
			return nil, fmt.Errorf("failed to delete expense: %w", err)
		}
	}

	log.InfoContext(
		ctx,
		"success",
	)
	return nil, nil
}
