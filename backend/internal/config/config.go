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
	"time"
)

const (
	defaultEnvironment = "dev"
	defaultLogLevel    = "info"
	defaultPort        = 8080
)

const (
	defaultReadHeaderTimeout = 5 * time.Second
	defaultReadTimeout       = 10 * time.Second
	defaultWriteTimeout      = 10 * time.Second
	defaultIdleTimeout       = 120 * time.Second
)

// Config contains typed backend runtime settings.
type Config struct {
	Environment       string
	LogLevel          slog.Level
	Port              int
	DatabaseURL       string
	ReadHeaderTimeout time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
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

	cfg.ReadHeaderTimeout, err = getDurationEnvOrDefault("HTTP_READ_HEADER_TIMEOUT", defaultReadHeaderTimeout)
	if err != nil {
		return Config{}, fmt.Errorf("parse HTTP_READ_HEADER_TIMEOUT: %w", err)
	}

	cfg.ReadTimeout, err = getDurationEnvOrDefault("HTTP_READ_TIMEOUT", defaultReadTimeout)
	if err != nil {
		return Config{}, fmt.Errorf("parse HTTP_READ_TIMEOUT: %w", err)
	}

	cfg.WriteTimeout, err = getDurationEnvOrDefault("HTTP_WRITE_TIMEOUT", defaultWriteTimeout)
	if err != nil {
		return Config{}, fmt.Errorf("parse HTTP_WRITE_TIMEOUT: %w", err)
	}

	cfg.IdleTimeout, err = getDurationEnvOrDefault("HTTP_IDLE_TIMEOUT", defaultIdleTimeout)
	if err != nil {
		return Config{}, fmt.Errorf("parse HTTP_IDLE_TIMEOUT: %w", err)
	}

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

	if c.ReadHeaderTimeout <= 0 {
		return fmt.Errorf("HTTP_READ_HEADER_TIMEOUT must be positive, got %s", c.ReadHeaderTimeout)
	}

	if c.ReadTimeout <= 0 {
		return fmt.Errorf("HTTP_READ_TIMEOUT must be positive, got %s", c.ReadTimeout)
	}

	if c.WriteTimeout <= 0 {
		return fmt.Errorf("HTTP_WRITE_TIMEOUT must be positive, got %s", c.WriteTimeout)
	}

	if c.IdleTimeout <= 0 {
		return fmt.Errorf("HTTP_IDLE_TIMEOUT must be positive, got %s", c.IdleTimeout)
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

func getDurationEnvOrDefault(key string, fallback time.Duration) (time.Duration, error) {
	value := os.Getenv(key)
	if value == "" {
		return fallback, nil
	}

	duration, err := time.ParseDuration(value)
	if err != nil {
		return 0, err
	}

	return duration, nil
}
