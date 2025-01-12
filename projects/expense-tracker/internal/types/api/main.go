package api

import "context"

type Api interface {
	Start(context.Context) error
}
