package router

import (
	"net/http"

	aiHandler "github.com/talgat-ruby/lessons-go/lesson4/example1/internal/api/handler/ai"
)

func ai(mux *http.ServeMux, apiKey string) {
	h := aiHandler.New(apiKey)

	mux.Handle("POST /prompt", http.HandlerFunc(h.Prompt))
}
