package configs

import (
	"context"
	"flag"
)

type Config struct {
	Api *ApiConfig
}

func NewConfig(ctx context.Context) (*Config, error) {
	conf := &Config{}

	_ = conf.loadDotEnvFiles()

	// Api config
	if c, err := newApiConfig(ctx); err != nil {
		return nil, err
	} else {
		conf.Api = c
	}

	flag.Parse()

	return conf, nil
}
