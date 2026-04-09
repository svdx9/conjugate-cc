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
	statusHandler *status.Handler
}

// NewCompositeHandler creates a new composite handler
func NewCompositeHandler(statusHandler *status.Handler) *CompositeHandler {
	return &CompositeHandler{
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

// RequestMagicLink is a stub implementation for the magic link request endpoint
func (h *CompositeHandler) RequestMagicLink(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement magic link request handler
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// GetMagicLinkVerify is a stub implementation for the magic link verify GET endpoint
func (h *CompositeHandler) GetMagicLinkVerify(w http.ResponseWriter, r *http.Request, params api.GetMagicLinkVerifyParams) {
	// TODO: Implement magic link verify GET handler
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// PostMagicLinkVerify is a stub implementation for the magic link verify POST endpoint
func (h *CompositeHandler) PostMagicLinkVerify(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement magic link verify POST handler
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// DeleteSession is a stub implementation for the session deletion endpoint
func (h *CompositeHandler) DeleteSession(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement session deletion handler
	http.Error(w, "Not implemented", http.StatusNotImplemented)
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
