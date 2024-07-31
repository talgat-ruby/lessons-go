package configs

import (
	"github.com/joho/godotenv"
)

func (c *Config) loadDotEnvFiles() error {
	return godotenv.Load(".env", ".env.local")
}
