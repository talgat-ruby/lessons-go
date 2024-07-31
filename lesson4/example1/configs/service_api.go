package configs

import (
	"context"
	"flag"

	"github.com/sethvargo/go-envconfig"
)

type ApiConfig struct {
	Port     int    `env:"PORT"`
	AIApiKey string `env:"AI_API_KEY"`
}

func newApiConfig(ctx context.Context) (*ApiConfig, error) {
	c := &ApiConfig{}

	if err := envconfig.Process(ctx, c); err != nil {
		return nil, err
	}

	flag.IntVar(&c.Port, "port", c.Port, "server port [PORT]")
	flag.StringVar(&c.AIApiKey, "token-key", c.AIApiKey, "Open AI API key [AI_API_KEY]")

	return c, nil
}
