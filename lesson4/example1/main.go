package main

import (
	"context"
	"log/slog"

	"github.com/talgat-ruby/lessons-go/lesson4/example1/configs"
	"github.com/talgat-ruby/lessons-go/lesson4/example1/internal/api"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// conf
	conf, err := configs.NewConfig(ctx)
	if err != nil {
		panic(err)
	}

	// configure gateway service
	srv := api.New(conf.Api)
	slog.InfoContext(ctx, "initialize service", "service", "api")
	// start gateway service
	srv.Start(ctx, cancel)

	<-ctx.Done()
	// Your cleanup tasks go here
	slog.InfoContext(ctx, "cleaning up ...")

	slog.InfoContext(ctx, "server was successful shutdown.")
}
