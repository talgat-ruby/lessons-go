package logger

import (
	"log/slog"
	"os"
)

func New(isJson bool) *slog.Logger {
	if isJson {
		return slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}

	return slog.New(
		slog.NewTextHandler(
			os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		),
	)
}
