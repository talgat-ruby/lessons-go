package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/talgat-ruby/lessons-go/lesson8/planet-express-personal/internal/conf"
)

type Model struct {
	conf *conf.DBConfig
	db   *sql.DB
}

func NewModel(conf *conf.DBConfig) (*Model, error) {
	db, err := NewDB(
		conf.Host,
		conf.Port,
		conf.User,
		conf.Password,
		conf.DBName,
	)

	if err != nil {
		return nil, err
	}

	m := &Model{
		conf: conf,
		db:   db,
	}

	return m, nil
}

func NewDB(host string, port int, user string, password string, dbname string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
