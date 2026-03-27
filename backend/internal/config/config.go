package config

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
)

type envKey string

const (
	Dev string = "dev"
)

const (
	env  envKey = "ENV"
	port envKey = "PORT"
	host envKey = "HOST"
)
const (
	defaultPort = 8080
	defaultEnv  = Dev
	defaultHost = "0.0.0.0"
)

var (
	errPortOutOfRange = errors.New("port out of range")
	errInvalidHost    = errors.New("invalid host")
)

// Config holds the application configuration.
type Config struct {
	Host string // must be a valid IP
	Port int
	Env  string
}

// Load parses environment variables and returns a Config struct.
func Load() (*Config, error) {

	port, err := getIntFromEnv(port, defaultPort)
	if err != nil {
		return nil, fmt.Errorf("invalid port: %w", err)
	}

	cfg := &Config{
		Port: port,
		Env:  getFromEnv(env, defaultEnv),
		Host: getFromEnv(host, defaultHost),
	}

	err = cfg.Validate()
	if err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	return cfg, nil
}

// Validate ensures all required configuration is present and valid.
func (c *Config) Validate() error {
	if c.Port < 0 || c.Port > 65535 {
		return fmt.Errorf("%w: %d", errPortOutOfRange, c.Port)
	}
	// check that host is a valid IP
	ip := net.ParseIP(c.Host)
	if ip == nil {
		return fmt.Errorf("%w: %s", errInvalidHost, c.Host)
	}
	return nil
}

func getIntFromEnv(key envKey, defaultValue int) (int, error) {
	value := os.Getenv(string(key))
	if value == "" {
		return defaultValue, nil
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("invalid %s: %w", key, err)
	}
	return i, nil
}

func getFromEnv(key envKey, defaultValue string) string {
	value := os.Getenv(string(key))
	if value == "" {
		return defaultValue
	}
	return value
}
