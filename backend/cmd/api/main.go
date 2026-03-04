package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/svdx9/conjugate-cc/backend/internal/config"
	httpserver "github.com/svdx9/conjugate-cc/backend/internal/http"
	statusapi "github.com/svdx9/conjugate-cc/backend/internal/status/api"
	statusservice "github.com/svdx9/conjugate-cc/backend/internal/status/service"
)

var (
	serviceGitSHA    = "dev"     //nolint:gochecknoglobals // set via ldflags at build time
	serviceBuildTime = "unknown" //nolint:gochecknoglobals // set via ldflags at build time
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	cfg, err := config.Load()
	if err != nil {
		logger.Error("backend exited", "error", err)
		os.Exit(1)
	}

	logger = newLogger(cfg)
	service := statusservice.New(serviceGitSHA, serviceBuildTime)
	handler := statusapi.NewHandler(logger, service)
	router := httpserver.NewRouter(logger, handler)
	server := httpserver.NewServer(cfg, logger, router)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	err = server.Run(ctx)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error("backend exited", "error", err)
		os.Exit(1)
	}
}

func newLogger(cfg config.Config) *slog.Logger {
	options := &slog.HandlerOptions{Level: cfg.LogLevel} //nolint:exhaustruct
	handler := slog.NewTextHandler(os.Stdout, options)
	return slog.New(handler)
}
