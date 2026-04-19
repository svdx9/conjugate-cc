package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/svdx9/conjugate-cc/backend/internal/api/v1"
	"github.com/svdx9/conjugate-cc/backend/internal/status"
)

// CompositeHandler combines multiple handlers to implement the full ServerInterface
type CompositeHandler struct {
	api.Unimplemented
	statusHandler *status.Handler
}

// NewCompositeHandler creates a new composite handler
func NewCompositeHandler(statusHandler *status.Handler) *CompositeHandler {
	return &CompositeHandler{ //nolint:exhaustruct
		statusHandler: statusHandler,
	}
}

// GetStatus implements the status endpoint
func (h *CompositeHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	h.statusHandler.GetStatus(w, r)
}

// GetMetadata implements the metadata endpoint
func (h *CompositeHandler) GetMetadata(w http.ResponseWriter, r *http.Request) {
	h.statusHandler.GetMetadata(w, r)
}

// NewRouter creates a new chi router and mounts the API handlers.
func NewRouter(statusHandler *status.Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)

	// Create composite handler that implements all required interfaces
	compositeHandler := NewCompositeHandler(statusHandler)

	// Mount the generated handlers at /api prefix
	api.HandlerFromMuxWithBaseURL(compositeHandler, r, "/api")

	return r
}
