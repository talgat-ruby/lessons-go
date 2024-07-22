package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", GetHello)

	// start up HTTP
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8081),
		Handler: mux,
	}

	slog.Info("Listening on", "PORT", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		slog.Error("start server error:", "error", err)
	}
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
