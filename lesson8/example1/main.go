// main.go
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

func main() {
	w := os.Getenv("NAME")

	ctx := context.Background()
	//slog.InfoContext(ctx, fmt.Sprintf("%d + %d is %d", 3, 5, math.Add(3, 5)))
	//slog.InfoContext(ctx, fmt.Sprintf("%d + %d is %d", 3, 5, math.Multiply(3, 5)))
	//
	//slog.InfoContext(ctx, fmt.Sprintf("Second Law is %d", internal.Force(2, 4)))

	slog.InfoContext(ctx, fmt.Sprintf("Hello %s!", w))
}

func init() {
	slog.Info("Hello from init")
}
