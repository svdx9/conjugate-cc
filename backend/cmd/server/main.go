package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/svdx9/conjugate-cc/backend/internal/config"
	internalhttp "github.com/svdx9/conjugate-cc/backend/internal/http"
	"github.com/svdx9/conjugate-cc/backend/internal/status"
)

// Injected at build time via -ldflags
//
//nolint:gochecknoglobals
var (
	GitSHA    = "unknown"
	BuildTime = "unknown"
)

func main() {
	handler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	cfg, err := config.FromEnv()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	logger.Info("starting conjugate-cc backend",
		"port", cfg.Port,
		"env", cfg.Env,
		"git_sha", GitSHA,
		"build_time", BuildTime,
	)

	// Dependency injection
	statusHandler := status.NewHandler(logger, GitSHA, BuildTime)
	router := internalhttp.NewRouter(statusHandler)

	//nolint:exhaustruct
	server := &http.Server{
		Addr:              net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)),
		Handler:           router,
		ErrorLog:          slog.NewLogLogger(handler, slog.LevelError),
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	// Graceful shutdown setup
	serverErr := make(chan error, 1)
	go func() {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- fmt.Errorf("server failed on %s: %w", server.Addr, err)
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErr:
		logger.Error("server error", "error", err)
		os.Exit(1)
	case sig := <-shutdown:
		logger.Info("shutting down server", "signal", sig.String())

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := server.Shutdown(ctx)
		if err != nil {
			logger.Error("graceful shutdown failed", "error", err)
			_ = server.Close()
			os.Exit(1)
		}
		logger.Info("server exited gracefully")
	}
}
