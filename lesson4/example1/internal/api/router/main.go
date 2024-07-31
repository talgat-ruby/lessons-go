package router

import (
	"net/http"

	apiT "github.com/talgat-ruby/lessons-go/lesson4/example1/internal/api/types"
)

// SetupRoutes setup router api
func SetupRoutes(mux *http.ServeMux, api apiT.Api) {
	ping(mux)
	ai(mux, api.Config().AIApiKey)
}
