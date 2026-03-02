package httpserver

import (
	"encoding/json"
	"net/http"

	apiv1 "github.com/svdx9/conjugate-cc/backend/internal/api/v1"
)

func WriteJSON[T any](w http.ResponseWriter, status int, payload T) error {
	js, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(js)
	return err
}

func WriteError(w http.ResponseWriter, status int, message string) {
	_ = WriteJSON(w, status, apiv1.ErrorResponse{
		Code:    http.StatusText(status),
		Message: message,
	})
}
