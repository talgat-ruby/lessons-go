package config

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/constant"
)

type Config struct {
	ENV      constant.Environment
	API      *APIConfig
	Postgres *PostgresConfig
}

func New(ctx context.Context) *Config {
	var env constant.Environment
	e := os.Getenv("ENV")
	if e == "" {
		env = constant.EnvironmentLocal
	} else {
		env = constant.Environment(e)
	}

	load(env)

	return &Config{
		ENV:      env,
		API:      newApiConfig(ctx),
		Postgres: newPostgresConfig(ctx),
	}
}

func load(env constant.Environment) error {
	switch env {
	case constant.EnvironmentProd:
		return godotenv.Load(".env", ".env.local", ".env.prod")
	case constant.EnvironmentDev:
		return godotenv.Load(".env", ".env.local", ".env.dev")
	case constant.EnvironmentTest:
		return godotenv.Load(".env", ".env.local", ".env.test")
	case constant.EnvironmentLocal:
		return godotenv.Load(".env", ".env.local")
	default:
		return nil
	}
}
