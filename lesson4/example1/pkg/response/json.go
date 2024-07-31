package response

import (
	"encoding/json"
	"net/http"
)

type DataResponse struct {
	Data interface{} `json:"data"`
}

type Error struct {
	Message string `json:"message"`
}

type ErrorsResponse struct {
	Errors []Error `json:"errors"`
}

func JSON(w http.ResponseWriter, data interface{}, errs ...error) error {
	if data != nil {
		return json.NewEncoder(w).Encode(&DataResponse{Data: data})
	}

	errList := make([]Error, len(errs))
	for i, err := range errs {
		errList[i] = Error{
			Message: err.Error(),
		}
	}

	return json.NewEncoder(w).Encode(&ErrorsResponse{Errors: errList})
}
