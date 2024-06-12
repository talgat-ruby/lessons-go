package conf

import (
	"flag"
)

type ApiConfig struct {
	*SharedConfig
}

func newApiConfig() (*ApiConfig, error) {
	c := &ApiConfig{}

	flag.IntVar(&c.Port, "port", c.Port, "server port [PORT]")

	return c, nil
}
