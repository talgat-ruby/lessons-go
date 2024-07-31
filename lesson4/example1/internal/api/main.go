package api

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/talgat-ruby/lessons-go/lesson4/example1/configs"
	"github.com/talgat-ruby/lessons-go/lesson4/example1/internal/api/router"
	apiT "github.com/talgat-ruby/lessons-go/lesson4/example1/internal/api/types"
)

type server struct {
	conf *configs.ApiConfig
}

func New(conf *configs.ApiConfig) apiT.Api {
	s := &server{
		conf: conf,
	}

	return s
}

func (s *server) Config() *configs.ApiConfig {
	return s.conf
}

func (s *server) Start(ctx context.Context, cancel context.CancelFunc) {
	mux := http.NewServeMux()
	router.SetupRoutes(mux, s)

	// start up HTTP
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.conf.Port),
		Handler: mux,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	// Listen from s different goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.ErrorContext(ctx, "server error", "error", err)
		}

		cancel()
	}()

	slog.InfoContext(
		ctx,
		"starting api service",
		"port", s.conf.Port,
		"playground", fmt.Sprintf("http://localhost:%d/", s.conf.Port),
	)

	shutdown := make(chan os.Signal, 1)   // Create channel to signify s signal being sent
	signal.Notify(shutdown, os.Interrupt) // When an interrupt is sent, notify the channel

	go func() {
		sig := <-shutdown

		slog.WarnContext(ctx, "signal received - shutting down...", "signal", sig)
		if err := srv.Shutdown(ctx); err != nil {
			slog.ErrorContext(ctx, "server shutdown error", "error", err)
		}
	}()
}
