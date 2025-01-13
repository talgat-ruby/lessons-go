package expense

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"

	pg "github.com/go-jet/jet/v2/postgres"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/postgres/generated/expenses/public/model"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/postgres/generated/expenses/public/table"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/types/database"
)

func (m *Expense) UpdateExpense(ctx context.Context, req database.UpdateExpenseReq) (database.UpdateExpenseResp, error) {
	log := m.logger.With(slog.String("handler", "UpdateExpense"))

	if req == nil {
		log.ErrorContext(ctx, "req is nil")
		return nil, fmt.Errorf("req is nil")
	}

	id, err := strconv.Atoi(req.GetID())
	if err != nil {
		return nil, fmt.Errorf("invalid expense id %s", req.GetID())
	}

	var cols pg.ColumnList
	if req.GetAmount() != nil {
		cols = append(cols, table.Expense.Amount)
	}
	if req.GetCategory() != nil {
		cols = append(cols, table.Expense.Category)
	}

	mdl := model.Expense{
		Amount:   *req.GetAmount(),
		Category: mapStringToModelCategory(*req.GetCategory()),
	}
	stmt := table.Expense.
		UPDATE(cols).
		MODEL(mdl).
		WHERE(table.Expense.ID.EQ(pg.Int(int64(id))))
	if _, err := stmt.ExecContext(ctx, m.db); err != nil {
		switch {
		default:
			log.ErrorContext(ctx, "failed to update expense", slog.Any("error", err))
			return nil, fmt.Errorf("failed to update expense: %w", err)
		}
	}

	log.InfoContext(
		ctx,
		"success",
	)
	return nil, nil
}
