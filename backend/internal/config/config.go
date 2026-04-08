package config

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	Dev string = "dev"
)

const (
	envKey  = "ENV"
	portKey = "PORT"
	hostKey = "HOST"
)

const (
	defaultPort = 8080
	defaultEnv  = Dev
	defaultHost = "0.0.0.0"
)

var (
	ErrPortOutOfRange = errors.New("port out of range")
	ErrInvalidHost    = errors.New("invalid host")
)

func getEnvOrDefault(key, defaultVal string) string {
	val, exists := os.LookupEnv(key)
	if !exists || strings.TrimSpace(val) == "" {
		return defaultVal
	}
	return strings.TrimSpace(val)
}

func getEnvOrDefaultInt(key string, defaultVal int) (int, error) {
	val := getEnvOrDefault(key, "")
	if val == "" {
		return defaultVal, nil
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", key, err)
	}
	return i, nil
}

// Config holds the application configuration.
type Config struct {
	Host string
	Port int
	Env  string
}

// FromEnv parses environment variables and returns a Config struct.
func FromEnv() (Config, error) {
	port, err := getEnvOrDefaultInt(portKey, defaultPort)
	if err != nil {
		return Config{}, fmt.Errorf("%w: %w", ErrPortOutOfRange, err)
	}
	if port < 0 || port > 65535 {
		return Config{}, fmt.Errorf("%w: %d", ErrPortOutOfRange, port)
	}

	host := getEnvOrDefault(hostKey, defaultHost)
	err = validateHost(host)
	if err != nil {
		return Config{}, fmt.Errorf("%w: %w", ErrInvalidHost, err)
	}

	cfg := Config{
		Port: port,
		Env:  getEnvOrDefault(envKey, defaultEnv),
		Host: host,
	}

	return cfg, nil
}

func validateHost(host string) error {
	if ip := net.ParseIP(host); ip != nil {
		return nil
	}
	_, err := net.LookupHost(host)
	if err != nil {
		return ErrInvalidHost
	}
	return nil
}
