package rest

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/internal/config"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/internal/rest/handler"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/internal/rest/router"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/internal/types/database"
)

type Rest struct {
	conf   *config.APIRestConfig
	logger *slog.Logger
	router *router.Router
}

func New(conf *config.APIRestConfig, logger *slog.Logger, db database.Database) *Rest {
	h := handler.New(logger.With(slog.String("module", "handler")), db)
	r := router.New(h)

	return &Rest{
		conf:   conf,
		logger: logger,
		router: r,
	}
}

func (a *Rest) Start(ctx context.Context) error {
	mux := a.router.Start(ctx)

	errLogger := slog.NewLogLogger(a.logger.Handler(), slog.LevelError)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.conf.Port),
		Handler: mux,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
		ErrorLog: errLogger,
	}

	a.logger.InfoContext(ctx, "starting server", slog.Int("port", a.conf.Port))

	if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
