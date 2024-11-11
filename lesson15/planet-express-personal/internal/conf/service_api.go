package conf

import (
	"context"
	"flag"

	"github.com/sethvargo/go-envconfig"
)

type ApiConfig struct {
	*SharedConfig
}

func newApiConfig(ctx context.Context) (*ApiConfig, error) {
	c := &ApiConfig{}

	if err := envconfig.Process(ctx, c); err != nil {
		return nil, err
	}

	flag.IntVar(&c.Port, "port", c.Port, "server port [PORT]")

	return c, nil
}
