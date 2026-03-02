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
	apphttp "github.com/svdx9/conjugate-cc/backend/internal/http"
	statusapi "github.com/svdx9/conjugate-cc/backend/internal/status/api"
	statusservice "github.com/svdx9/conjugate-cc/backend/internal/status/service"
)

var (
	serviceGitSHA    = "dev"
	serviceBuildTime = "unknown"
)

func main() {
	err := run()
	if err != nil {
		slog.Error("backend exited", "error", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	logger := newLogger(cfg)
	service := statusservice.New(serviceGitSHA, serviceBuildTime)
	handler := statusapi.NewHandler(service)
	router := apphttp.NewRouter(logger, handler)
	server := apphttp.NewServer(cfg, logger, router)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	err = server.Run(ctx)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func newLogger(cfg config.Config) *slog.Logger {
	options := &slog.HandlerOptions{Level: cfg.LogLevel}
	handler := slog.NewTextHandler(os.Stdout, options)
	return slog.New(handler)
}
