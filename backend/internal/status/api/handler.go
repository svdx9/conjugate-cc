package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	apiv1 "github.com/svdx9/conjugate-cc/backend/internal/api/v1"
	statusservice "github.com/svdx9/conjugate-cc/backend/internal/status/service"
)

// Handler adapts the status service to the generated API surface.
type Handler struct {
	service *statusservice.Service
}

// NewHandler constructs a status API handler.
func NewHandler(service *statusservice.Service) *Handler {
	return &Handler{service: service}
}

// GetBuildInfo handles GET /v1/build-info.
func (h *Handler) GetBuildInfo(w http.ResponseWriter, r *http.Request) {
	response := h.service.BuildInfo(r.Context())

	payload := apiv1.BuildInfoResponse{
		BuildTime: response.BuildTime,
		GitSha:    response.GitSHA,
	}

	writeJSON(w, http.StatusOK, payload)
}

// GetHealth handles GET /v1/health.
func (h *Handler) GetHealth(w http.ResponseWriter, r *http.Request) {
	response := h.service.Health(r.Context())

	payload := apiv1.HealthResponse{
		Status: response.Status,
	}

	writeJSON(w, http.StatusOK, payload)
}

func writeJSON(w http.ResponseWriter, statusCode int, payload any) {
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	err := encoder.Encode(payload)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(buffer.Bytes())
}
