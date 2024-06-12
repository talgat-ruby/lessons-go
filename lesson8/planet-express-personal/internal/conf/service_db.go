package conf

import (
	"flag"
)

type DBConfig struct {
	*SharedConfig `env:",prefix=DB_"`
	User          string `env:"DB_USER"`
	Password      string `env:"DB_PASSWORD"`
	DBName        string `env:"DB_NAME"`
}

func newDBConfig() (*DBConfig, error) {
	c := &DBConfig{}

	flag.StringVar(&c.Host, "db-host", c.Host, "database host [DB_HOST]")
	flag.IntVar(&c.Port, "db-port", c.Port, "database port [DB_PORT]")
	flag.StringVar(&c.User, "db-user", c.User, "database user [DB_USER]")
	flag.StringVar(&c.Password, "db-password", c.Password, "database password [DB_PASSWORD]")
	flag.StringVar(&c.DBName, "db-name", c.DBName, "database name [DB_NAME]")

	return c, nil
}
