package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/svdx9/conjugate-cc/backend/internal/api/v1"
	"github.com/svdx9/conjugate-cc/backend/internal/status"
)

// NewRouter creates a new chi router and mounts the API handlers.
func NewRouter(statusHandler *status.Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)

	// Mount the generated handlers
	api.HandlerFromMux(statusHandler, r)

	return r
}
