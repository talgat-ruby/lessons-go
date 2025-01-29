package interceptor

import (
	"context"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/grpc/expense"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/grpc/sanitary"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/types/controller"
)

type Interceptor struct {
	log  *slog.Logger
	ctrl controller.Controller
}

func New(log *slog.Logger, ctrl controller.Controller) *Interceptor {
	return &Interceptor{
		log:  log,
		ctrl: ctrl,
	}
}

// wrappedServerStream wraps grpc.ServerStream to modify the context
type wrappedServerStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *wrappedServerStream) Context() context.Context {
	return w.ctx
}

func newWrappedServerStream(s grpc.ServerStream, ctx context.Context) grpc.ServerStream {
	return &wrappedServerStream{s, ctx}
}

func (inter *Interceptor) Unary(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	inter.log = inter.log.With(slog.String("type", "unary"))

	newCtx, err := inter.both(ctx, info.Server)
	if err != nil {
		return nil, err
	}

	return handler(newCtx, req)
}

func (inter *Interceptor) Stream(
	srv interface{},
	stream grpc.ServerStream,
	_ *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	ctx := stream.Context()

	inter.log = inter.log.With(slog.String("type", "stream"))

	newCtx, err := inter.both(ctx, srv)
	if err != nil {
		return err
	}

	return handler(srv, newWrappedServerStream(stream, newCtx))
}

func (inter *Interceptor) both(ctx context.Context, srv any) (context.Context, error) {
	switch t := srv.(type) {
	case *sanitary.Sanitary:
		return ctx, nil
	case *expense.Expense:
		return inter.authenticate(ctx)
	default:
		return nil, status.Errorf(codes.Internal, "unsuported server type: %T", t)
	}
}
