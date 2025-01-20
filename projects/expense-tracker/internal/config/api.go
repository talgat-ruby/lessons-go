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
	GraphQL     *APIGraphQLConfig
	Grpc        *APIGrpcConfig
}

func newApiConfig(ctx context.Context) *APIConfig {
	c := &APIConfig{
		TokenSecret: os.Getenv("API_TOKEN_SECRET"),
		Pepper:      os.Getenv("API_PEPPER"),
		Rest:        newApiRestConfig(ctx),
		GraphQL:     newApiGraphQLConfig(ctx),
		Grpc:        newAPIGrpcConfig(ctx),
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

type APIGraphQLConfig struct {
	Host string
	Port int
}

func newApiGraphQLConfig(_ context.Context) *APIGraphQLConfig {
	port, _ := strconv.Atoi(os.Getenv("API_GRAPHQL_PORT"))

	c := &APIGraphQLConfig{
		Host: os.Getenv("API_GRAPHQL_HOST"),
		Port: port,
	}

	flag.StringVar(&c.Host, "api-graphql-host", c.Host, "api graphql host [API_GRAPHQL_HOST]")
	flag.IntVar(&c.Port, "api-graphql-port", c.Port, "api graphql port [API_GRAPHQL_PORT]")

	return c
}

type APIGrpcConfig struct {
	Host string
	Port int
}

func newAPIGrpcConfig(_ context.Context) *APIGrpcConfig {
	port, _ := strconv.Atoi(os.Getenv("API_GRPC_PORT"))

	c := &APIGrpcConfig{
		Host: os.Getenv("API_GRPC_HOST"),
		Port: port,
	}

	flag.StringVar(&c.Host, "api-grpc-host", c.Host, "api grpc host [API_GRPC_HOST]")
	flag.IntVar(&c.Port, "api-grpc-port", c.Port, "api grpc port [API_GRPC_PORT]")

	return c
}
