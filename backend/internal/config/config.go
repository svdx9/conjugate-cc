// Package config loads backend runtime configuration from environment variables.
// Defaults are provided for local development where practical, and validation
// fails fast before the HTTP server starts.
package config

import (
	"errors"
	"fmt"
	"log/slog"
	"net"
	"os"
	"strconv"
)

const (
	defaultEnvironment = "dev"
	defaultLogLevel    = "info"
	defaultPort        = 8080
)

// Config contains typed backend runtime settings.
type Config struct {
	Environment string
	LogLevel    slog.Level
	Port        int
	DatabaseURL string
}

// Load parses configuration exactly once at startup.
func Load() (Config, error) {
	cfg := Config{}

	cfg.Environment = getEnvOrDefault("ENV", defaultEnvironment)

	levelText := getEnvOrDefault("LOG_LEVEL", defaultLogLevel)
	err := cfg.LogLevel.UnmarshalText([]byte(levelText))
	if err != nil {
		return Config{}, fmt.Errorf("parse LOG_LEVEL: %w", err)
	}

	portText := getEnvOrDefault("PORT", strconv.Itoa(defaultPort))
	port, err := strconv.Atoi(portText)
	if err != nil {
		return Config{}, fmt.Errorf("parse PORT: %w", err)
	}
	cfg.Port = port

	cfg.DatabaseURL = os.Getenv("DATABASE_URL")

	err = cfg.Validate()
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

// Validate enforces startup constraints on the parsed config.
func (c Config) Validate() error {
	if c.Environment == "" {
		return errors.New("ENV must not be empty")
	}

	if c.Port < 1 || c.Port > 65535 {
		return fmt.Errorf("PORT must be between 1 and 65535, got %d", c.Port)
	}

	return nil
}

// ListenAddress returns the service bind address.
func (c Config) ListenAddress() string {
	return net.JoinHostPort("", strconv.Itoa(c.Port))
}

func getEnvOrDefault(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
