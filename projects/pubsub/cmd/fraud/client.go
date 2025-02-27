package main

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func NewClient(cfg *Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       0,
	})

	return client
}
