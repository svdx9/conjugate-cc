package apphttp

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	apiv1 "github.com/svdx9/conjugate-cc/backend/internal/api/v1"
	"github.com/svdx9/conjugate-cc/backend/internal/config"
)

const shutdownTimeout = 5 * time.Second

// Server owns the HTTP server lifecycle.
type Server struct {
	logger *slog.Logger
	server *http.Server
}

// NewRouter wires shared middleware and generated routes.
func NewRouter(logger *slog.Logger, handler apiv1.ServerInterface) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(requestLogger(logger))
	router.Use(middleware.Recoverer)
	apiv1.HandlerFromMux(handler, router)
	return router
}

// NewServer constructs the HTTP server with deterministic defaults.
func NewServer(cfg config.Config, logger *slog.Logger, handler http.Handler) *Server {
	httpServer := &http.Server{
		Addr:              cfg.ListenAddress(),
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
	}

	return &Server{
		logger: logger,
		server: httpServer,
	}
}

// Run starts the server and shuts it down when the context is cancelled.
func (s *Server) Run(ctx context.Context) error {
	errCh := make(chan error, 1)

	go func() {
		s.logger.Info("http server starting", "addr", s.server.Addr)

		err := s.server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- fmt.Errorf("listen and serve: %w", err)
			return
		}

		errCh <- nil
	}()

	select {
	case <-ctx.Done():
		s.logger.Info("http server shutting down")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		err := s.server.Shutdown(shutdownCtx)
		if err != nil {
			return fmt.Errorf("shutdown http server: %w", err)
		}

		err = <-errCh
		if err != nil {
			return err
		}

		return nil
	case err := <-errCh:
		return err
	}
}
