package status

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/svdx9/conjugate-cc/backend/internal/api/v1"
)

// Handler implements the api.ServerInterface.
type Handler struct {
	logger    *slog.Logger
	gitSHA    string
	buildTime string
}

// NewHandler creates a new status handler.
func NewHandler(logger *slog.Logger, gitSHA, buildTime string) *Handler {
	return &Handler{
		logger:    logger,
		gitSHA:    gitSHA,
		buildTime: buildTime,
	}
}

// GetStatus handles GET /v1/status.
func (h *Handler) GetStatus(w http.ResponseWriter, r *http.Request) {
	resp := api.Status{
		Status: "ok",
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

// GetMetadata handles GET /v1/metadata.
func (h *Handler) GetMetadata(w http.ResponseWriter, r *http.Request) {
	resp := api.Metadata{
		GitSha:    h.gitSHA,
		BuildTime: h.buildTime,
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		slog.Error("failed to encode response", "error", err)
	}
}
