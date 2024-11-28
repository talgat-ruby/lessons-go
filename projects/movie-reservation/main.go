package main

import (
	"context"
	"log/slog"
)

func main() {
	ctx := context.Background()

	d, err := newDB(slog.With("service", "db"))
	if err != nil {
		panic(err)
	}
	if err := d.Init(ctx); err != nil {
		panic(err)
	}

	a := newApi(slog.With("service", "api"))
	if err := a.Start(ctx); err != nil {
		panic(err)
	}
}
