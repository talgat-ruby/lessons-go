package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/internal/config"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/internal/constant"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/internal/postgres"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/internal/rest"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/pkg/logger"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// config
	conf := config.New(ctx)

	// logger
	log := logger.New(conf.ENV != constant.EnvironmentLocal)

	p, err := postgres.New(conf.Postgres, log.With(slog.String("service", "postgre")))
	if err != nil {
		log.ErrorContext(ctx, "failed to start postgre", slog.Any("error", err))
		panic(err)
	}

	r := rest.New(conf.API.Rest, log.With(slog.String("service", "rest")), p)
	go func(ctx context.Context, cancelFunc context.CancelFunc) {
		if err := r.Start(ctx); err != nil {
			log.ErrorContext(ctx, "failed to start rest", slog.Any("error", err))
		}

		cancelFunc()
	}(ctx, cancel)

	go func(cancelFunc context.CancelFunc) {
		shutdown := make(chan os.Signal, 1)   // Create channel to signify s signal being sent
		signal.Notify(shutdown, os.Interrupt) // When an interrupt is sent, notify the channel

		sig := <-shutdown
		log.WarnContext(ctx, "signal received - shutting down...", slog.Any("signal", sig))

		cancelFunc()
	}(cancel)

	<-ctx.Done()

	fmt.Println("shutting down gracefully")
}
