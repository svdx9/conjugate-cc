package api

import (
	"net/http"
	"log/slog"
	apiv1 "github.com/svdx9/conjugate-cc/backend/internal/api/v1"
	httpserver "github.com/svdx9/conjugate-cc/backend/internal/http"
	statusservice "github.com/svdx9/conjugate-cc/backend/internal/status/service"
)

// Handler adapts the status service to the generated API surface.
type Handler struct {
	logger *slog.Logger
	service *statusservice.Service
}

// NewHandler constructs a status API handler.
func NewHandler(logger *slog.Logger, service *statusservice.Service) *Handler {
	return &Handler{logger: logger, service: service}
}

// GetBuildInfo handles GET /v1/build-info.
func (h *Handler) GetBuildInfo(w http.ResponseWriter, r *http.Request) {
	response := h.service.BuildInfo(r.Context())

	payload := apiv1.BuildInfoResponse{
		BuildTime: response.BuildTime,
		GitSha:    response.GitSHA,
	}

	err := httpserver.WriteJSON(w, http.StatusOK, payload)
	if err != nil {
		h.logger.Error("failed to write response", "error", err)
	}
}

// GetHealth handles GET /v1/health.
func (h *Handler) GetHealth(w http.ResponseWriter, r *http.Request) {
	response := h.service.Health(r.Context())

	payload := apiv1.HealthResponse{
		Status: response.Status,
	}

	err := httpserver.WriteJSON(w, http.StatusOK, payload)
	if err != nil {
		h.logger.Error("failed to write response", "error", err)
	}
}
