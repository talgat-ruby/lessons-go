package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DataResponse struct {
	Data interface{} `json:"data"`
}

func JSON(w http.ResponseWriter, status int, data interface{}) error {
	if data == nil {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return fmt.Errorf("JSON marshal error: %w", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(js); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return fmt.Errorf("writer error: %w", err)
	}

	return nil
}
