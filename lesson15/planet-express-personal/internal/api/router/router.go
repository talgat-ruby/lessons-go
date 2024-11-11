package router

import (
	"net/http"

	"github.com/talgat-ruby/lessons-go/lesson8/planet-express-personal/internal/db"
)

type ErrJson struct {
	Message string `json:"message"`
}

func SetupRoutes(mux *http.ServeMux, model *db.Model) {
	emplH := &employeeHandler{
		model: model,
	}

	mux.HandleFunc("GET /ping", GetPingHandler)
	mux.HandleFunc("GET /employee/{id}", emplH.GetEmployeeHandler)
	mux.HandleFunc("GET /employee", emplH.GetEmployeesHandler)
	mux.HandleFunc("POST /employee", emplH.CreateEmployeeHandler)
}
