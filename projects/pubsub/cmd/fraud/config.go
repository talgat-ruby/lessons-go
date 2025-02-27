package main

import (
	"os"
	"strconv"
)

type Config struct {
	RedisHost     string
	RedisPort     int
	RedisPassword string
	RedisChannel  string
}

func LoadConfig() *Config {
	// Set default values
	cfg := &Config{
		RedisHost:     "localhost",
		RedisPort:     6379,
		RedisPassword: "",
		RedisChannel:  "TransactionCompleted",
	}

	// Override with environment variables if they exist
	if host := os.Getenv("REDIS_HOST"); host != "" {
		cfg.RedisHost = host
	}

	if port := os.Getenv("REDIS_PORT"); port != "" {
		portInt, err := strconv.Atoi(port)
		if err == nil {
			cfg.RedisPort = portInt
		}
	}

	if password := os.Getenv("REDIS_PASSWORD"); password != "" {
		cfg.RedisPassword = password
	}

	if channel := os.Getenv("REDIS_CHANNEL"); channel != "" {
		cfg.RedisChannel = channel
	}

	return cfg
}
