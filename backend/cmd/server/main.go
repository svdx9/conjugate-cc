package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/svdx9/conjugate-cc/backend/internal/auth"
	"github.com/svdx9/conjugate-cc/backend/internal/config"
	"github.com/svdx9/conjugate-cc/backend/internal/db"
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
	cfg, err := config.FromEnv()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level:       cfg.LogLevel,
		ReplaceAttr: nil,
		AddSource:   true,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	logger.Info("starting conjugate-cc backend",
		"config", cfg.Redacted(),
		"git_sha", GitSHA,
		"build_time", BuildTime,
	)

	// Create database pool
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	pool, err := db.NewPool(ctx, cfg.DatabaseURL)
	cancel()
	if err != nil {
		logger.Error("failed to create database pool", "error", err)
		os.Exit(1)
	}
	defer pool.Close()

	// Dependency injection - Authentication layer
	authStore := db.NewAuthStore(pool, logger)
	authService := auth.NewService(authStore, cfg.AuthMagicLinkTTL, cfg.AuthSessionTTL)

	// Dependency injection - HTTP layer
	statusHandler := status.NewHandler(logger, GitSHA, BuildTime)
	router := internalhttp.NewRouter(statusHandler)

	// Wire auth service into router (when HTTP handlers are created)
	// TODO: Add route handlers that use authService when implementing Task 008.09
	_ = authService

	//nolint:exhaustruct
	server := &http.Server{
		Addr:              cfg.Addr(),
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
