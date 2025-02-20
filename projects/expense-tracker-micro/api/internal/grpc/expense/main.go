package expense

import (
	"log/slog"

	expensev1 "github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/grpc/generated/expense-tracker/expense/v1"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/types/controller"
)

type Expense struct {
	expensev1.UnimplementedExpenseServiceServer
	log  *slog.Logger
	ctrl controller.Controller
}

func New(log *slog.Logger, ctrl controller.Controller) *Expense {
	return &Expense{
		log:  log,
		ctrl: ctrl,
	}
}
