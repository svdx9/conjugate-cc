package config

import (
	"fmt"
	"os"
)

// Config holds the application configuration.
type Config struct {
	Port string
	Env  string
}

// Load parses environment variables and returns a Config struct.
func Load() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	cfg := &Config{
		Port: port,
		Env:  env,
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	return cfg, nil
}

// Validate ensures all required configuration is present and valid.
func (c *Config) Validate() error {
	if c.Port == "" {
		return fmt.Errorf("PORT is required")
	}
	return nil
}
