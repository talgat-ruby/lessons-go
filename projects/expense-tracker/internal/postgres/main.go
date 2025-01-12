package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/lib/pq"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/config"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/postgres/model"
)

type Postgres struct {
	*model.Model
}

func New(conf *config.PostgresConfig, logger *slog.Logger) (*Postgres, error) {
	db, err := NewDB(conf)
	if err != nil {
		return nil, err
	}

	return &Postgres{
		Model: model.New(conf, logger.With(slog.String("module", "model")), db),
	}, nil
}

func NewDB(conf *config.PostgresConfig) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Name,
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}
