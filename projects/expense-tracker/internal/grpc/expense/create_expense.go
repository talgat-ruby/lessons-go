package expense

import (
	"context"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	expensev1 "github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/grpc/generated/expense-tracker/expense/v1"
)

func (s *Expense) CreateExpense(ctx context.Context, req *expensev1.CreateExpenseRequest) (*expensev1.CreateExpenseResponse, error) {
	log := s.log.With("method", "CreateExpense")

	_, err := s.ctrl.NewExpense(ctx, newCtrlReqCreateExpense(req))
	if err != nil {
		log.ErrorContext(ctx, "fail", slog.Any("error", err))
		return nil, status.Error(codes.Internal, "fail to create expense")
	}

	log.InfoContext(
		ctx,
		"success",
	)
	return &expensev1.CreateExpenseResponse{}, nil
}

type ctrlReqCreateExpense struct {
	*expensev1.CreateExpenseRequest
}

func newCtrlReqCreateExpense(req *expensev1.CreateExpenseRequest) *ctrlReqCreateExpense {
	return &ctrlReqCreateExpense{req}
}

func (s *ctrlReqCreateExpense) GetCategory() string {
	return s.CreateExpenseRequest.GetCategory().String()
}
