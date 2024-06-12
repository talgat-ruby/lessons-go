package conf

import (
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

func NewConfig() (*Config, error) {
	conf := &Config{}

	if err := godotenv.Load(".env", ".env.local", ".env.prod"); err != nil {
		return nil, err
	}

	// Api config
	if c, err := newApiConfig(); err != nil {
		return nil, err
	} else {
		conf.Api = c
	}

	// DB config
	if c, err := newDBConfig(); err != nil {
		return nil, err
	} else {
		conf.DB = c
	}

	flag.Parse()

	return conf, nil
}
