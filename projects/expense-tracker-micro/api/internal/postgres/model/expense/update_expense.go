package expense

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strconv"

	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/postgres/generated/expenses/public/model"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/postgres/generated/expenses/public/table"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/types/database"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/validation"
)

func (m *Expense) UpdateExpense(ctx context.Context, req database.UpdateExpenseReq) (database.UpdateExpenseResp, error) {
	log := m.logger.With(slog.String("handler", "UpdateExpense"))

	if req == nil {
		log.ErrorContext(ctx, "req is nil")
		return nil, fmt.Errorf("req is nil")
	}

	id, err := strconv.Atoi(req.GetID())
	if err != nil {
		log.ErrorContext(ctx, "could not convert id", slog.Any("error", err))
		return nil, validation.NewError("invalid expense id")
	}

	userId, err := strconv.Atoi(req.GetUserID())
	if err != nil {
		log.ErrorContext(ctx, "could not convert id", slog.Any("error", err))
		return nil, validation.NewError("invalid user id")
	}

	var cols pg.ColumnList
	var mdl model.Expense
	if req.GetAmount() != nil {
		cols = append(cols, table.Expense.Amount)
		mdl.Amount = *req.GetAmount()
	}
	if req.GetCategory() != nil {
		cols = append(cols, table.Expense.Category)
		mdl.Category = mapStringToModelCategory(*req.GetCategory())
	}

	stmt := table.Expense.
		UPDATE(cols).
		MODEL(mdl).
		WHERE(
			pg.AND(
				table.Expense.ID.EQ(pg.Int(int64(id))),
				table.Expense.UserID.EQ(pg.Int(int64(userId))),
			),
		)
	if _, err := stmt.ExecContext(ctx, m.db); err != nil {
		switch {
		case errors.Is(err, qrm.ErrNoRows):
			log.WarnContext(ctx, "no rows found")
			return nil, nil
		default:

			log.ErrorContext(ctx, "failed to update expense", slog.Any("error", err))
			return nil, fmt.Errorf("failed to update expense: %w", err)
		}
	}

	log.InfoContext(
		ctx,
		"success",
	)
	return true, nil
}
