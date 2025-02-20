package config

import (
	"context"
	"flag"
	"os"
	"strconv"
)

type RemotePostgresConfig struct {
	Host string
	Port int
}

func newRemotePostgresConfig(_ context.Context) *RemotePostgresConfig {
	port, _ := strconv.Atoi(os.Getenv("REMOTE_PG_PORT"))

	c := &RemotePostgresConfig{
		Host: os.Getenv("REMOTE_PG_HOST"),
		Port: port,
	}

	flag.StringVar(&c.Host, "remote-pg-host", c.Host, "remote pg host [REMOTE_PG_HOST]")
	flag.IntVar(&c.Port, "remote-pg-port", c.Port, "remote pg port [REMOTE_PG_PORT]")

	return c
}
