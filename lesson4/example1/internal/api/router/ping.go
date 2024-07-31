package router

import (
	"net/http"

	pingHandler "github.com/talgat-ruby/lessons-go/lesson4/example1/internal/api/handler/ping"
)

func ping(mux *http.ServeMux) {
	h := pingHandler.New()

	mux.Handle("GET /ping", http.HandlerFunc(h.Ping))
}
