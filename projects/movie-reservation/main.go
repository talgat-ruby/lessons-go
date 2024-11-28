package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	_ = godotenv.Load()

	fmt.Println(os.Getenv("API_PORT"))

	d, err := newDB(slog.With("service", "db"))
	if err != nil {
		panic(err)
	}
	if err := d.Init(ctx); err != nil {
		panic(err)
	}

	a := newApi(slog.With("service", "api"), d)
	if err := a.Start(ctx); err != nil {
		panic(err)
	}
}
