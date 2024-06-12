package conf

import (
	"context"
	"flag"

	"github.com/joho/godotenv"
)

type SharedConfig struct {
	Port int    `env:"PORT"`
	Host string `env:"HOST,default=localhost"`
}

type Config struct {
	Api *ApiConfig
	DB  *DBConfig
}

func NewConfig(ctx context.Context) (*Config, error) {
	conf := &Config{}

	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	// Api config
	if c, err := newApiConfig(ctx); err != nil {
		return nil, err
	} else {
		conf.Api = c
	}

	// DB config
	if c, err := newDBConfig(ctx); err != nil {
		return nil, err
	} else {
		conf.DB = c
	}

	flag.Parse()

	return conf, nil
}
