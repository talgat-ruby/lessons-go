package router

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/talgat-ruby/lessons-go/lesson8/planet-express-personal/internal/db"
)

type employeeHandler struct {
	model *db.Model
}

type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
	IsRobot  bool   `json:"is_robot"`
}

func (h *employeeHandler) GetEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	strId := r.PathValue("id")

	slog.Info("Get employee with id: ", strId)

	id, err := strconv.Atoi(strId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrJson{Message: err.Error()})
		return
	}

	emplModel, err := h.model.SelectEmployee(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrJson{Message: err.Error()})
		return
	}

	empl := &Employee{
		ID:       emplModel.ID,
		Name:     emplModel.Name,
		Position: emplModel.Position,
		IsRobot:  emplModel.IsRobot,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(empl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h *employeeHandler) GetEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Get employees")

	emplsModel, err := h.model.SelectEmployees()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrJson{Message: err.Error()})
		return
	}

	empls := make([]*Employee, len(emplsModel))
	for i, empl := range emplsModel {
		empls[i] = &Employee{
			ID:       empl.ID,
			Name:     empl.Name,
			Position: empl.Position,
			IsRobot:  empl.IsRobot,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(empls)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h *employeeHandler) CreateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	empl := &Employee{}

	if r.Header.Get("Content-Type") == "application/json" {
		if err := json.NewDecoder(r.Body).Decode(&empl); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	emplModel := &db.EmployeeModel{
		Name:     empl.Name,
		Position: empl.Position,
		IsRobot:  empl.IsRobot,
	}

	if err := h.model.InsertEmployee(emplModel); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrJson{Message: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
