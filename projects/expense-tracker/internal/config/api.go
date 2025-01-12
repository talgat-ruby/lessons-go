package config

import (
	"context"
	"flag"
	"os"
	"strconv"
)

type APIConfig struct {
	TokenSecret string
	Pepper      string
	Rest        *APIRestConfig
}

func newApiConfig(ctx context.Context) *APIConfig {
	c := &APIConfig{
		TokenSecret: os.Getenv("API_TOKEN_SECRET"),
		Pepper:      os.Getenv("API_PEPPER"),
		Rest:        newApiRestConfig(ctx),
	}

	flag.StringVar(&c.TokenSecret, "api-token-secret", c.TokenSecret, "api token secret [API_TOKEN_SECRET]")
	flag.StringVar(&c.Pepper, "api-pepper", c.Pepper, "api pepper [API_PEPPER]")

	return c
}

type APIRestConfig struct {
	Host string
	Port int
}

func newApiRestConfig(_ context.Context) *APIRestConfig {
	port, _ := strconv.Atoi(os.Getenv("API_REST_PORT"))

	c := &APIRestConfig{
		Host: os.Getenv("API_REST_HOST"),
		Port: port,
	}

	flag.StringVar(&c.Host, "api-rest-host", c.Host, "api rest host [API_REST_HOST]")
	flag.IntVar(&c.Port, "api-rest-port", c.Port, "api rest port [API_REST_PORT]")

	return c
}
