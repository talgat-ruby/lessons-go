package main

import (
	"context"
	"log/slog"

	"github.com/talgat-ruby/lessons-go/lesson8/planet-express-personal/internal/api"
	"github.com/talgat-ruby/lessons-go/lesson8/planet-express-personal/internal/conf"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// conf
	configs, err := conf.NewConfig()
	if err != nil {
		panic(err)
	}

	// configure gateway service
	srv := api.NewApi(configs.Api)
	slog.InfoContext(ctx, "initialize service", "service", "api")
	// start gateway service
	srv.Start(ctx, cancel)

	<-ctx.Done()
	// Your cleanup tasks go here
	slog.InfoContext(ctx, "cleaning up ...")

	slog.InfoContext(ctx, "server was successful shutdown.")
}
