package rest

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/config"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/rest/handler"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/rest/middleware"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/rest/router"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/types/controller"
)

type Rest struct {
	conf   *config.APIRestConfig
	logger *slog.Logger
	router *router.Router
}

func New(conf *config.APIRestConfig, logger *slog.Logger, ctrl controller.Controller) *Rest {
	midd := middleware.New(logger.With(slog.String("module", "middleware")), ctrl)
	h := handler.New(logger.With(slog.String("module", "handler")), ctrl)
	r := router.New(h, midd)

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
