package sanitary

import (
	"context"
	"fmt"

	sanitaryv1 "github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/grpc/generated/expense-tracker/sanitary/v1"
)

func (s *Sanitary) Ping(_ context.Context, req *sanitaryv1.PingRequest) (*sanitaryv1.PingResponse, error) {
	return &sanitaryv1.PingResponse{
		Message: fmt.Sprint("pong: ", req.GetMessage()),
	}, nil
}
