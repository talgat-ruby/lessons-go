package types

import (
	"context"

	"github.com/talgat-ruby/lessons-go/lesson4/example1/configs"
)

type Api interface {
	Start(ctx context.Context, cancel context.CancelFunc)
	Config() *configs.ApiConfig
}
