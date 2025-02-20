package sanitary

import (
	sanitaryv1 "github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/grpc/generated/expense-tracker/sanitary/v1"
	"log/slog"
)

type Sanitary struct {
	sanitaryv1.UnimplementedSanitaryServiceServer
	log *slog.Logger
}

func New(log *slog.Logger) *Sanitary {
	return &Sanitary{
		log: log,
	}
}
