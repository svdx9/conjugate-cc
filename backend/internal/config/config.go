package config

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	Dev string = "dev"
)

const (
	envKey              = "ENV"
	portKey             = "PORT"
	hostKey             = "HOST"
	databaseURLKey      = "DATABASE_URL"
	logLevelKey         = "LOG_LEVEL"
	uiPathKey           = "UI_PATH"
	authDevBypassKey    = "AUTH_DEV_BYPASS"
	authCookieSecureKey = "AUTH_COOKIE_SECURE"
	authSessionTTLKey   = "AUTH_SESSION_TTL"
	authMagicLinkTTLKey = "AUTH_MAGIC_LINK_TTL"
	SiteURLKey          = "SITE_URL"
)

const (
	defaultPort          = 8080
	defaultEnv           = Dev
	defaultHost          = "0.0.0.0"
	defaultUiPath        = "ui"
	defaultSessionTTL    = 30 * 24 * time.Hour
	defaultMagicLinkTTL  = 15 * time.Minute
	defaultAuthDevBypass = false
)

var (
	ErrPortOutOfRange      = errors.New("port out of range")
	ErrMissingEnv          = errors.New("missing required env var")
	ErrInvalidBool         = errors.New("invalid boolean value")
	ErrInvalidDuration     = errors.New("invalid duration value")
	ErrInvalidURL          = errors.New("invalid URL")
	ErrInvalidEnv          = errors.New("invalid environment value")
	ErrMagicLinkTTLRange   = errors.New("magic link TTL must be shorter than session TTL")
	ErrMissingMagicLinkURL = errors.New("missing magic link URL")
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

func requireEnv(key string) (string, error) {
	val, exists := os.LookupEnv(key)
	if !exists || strings.TrimSpace(val) == "" {
		return "", fmt.Errorf("%w: %s", ErrMissingEnv, key)
	}
	return strings.TrimSpace(val), nil
}

func parseBool(val string, defaultVal bool) (bool, error) {
	val = strings.TrimSpace(val)
	if val == "" {
		return defaultVal, nil
	}
	switch strings.ToLower(val) {
	case "1", "true", "yes", "y", "on":
		return true, nil
	case "0", "false", "no", "n", "off":
		return false, nil
	default:
		return false, fmt.Errorf("%w: %q", ErrInvalidBool, val)
	}
}

func parseDuration(val string, defaultVal time.Duration) (time.Duration, error) {
	val = strings.TrimSpace(val)
	if val == "" {
		return defaultVal, nil
	}
	d, err := time.ParseDuration(val)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", ErrInvalidDuration, err)
	}
	return d, nil
}

// Config holds the application configuration.
type Config struct {
	Host             string
	Port             int
	Env              string
	LogLevel         string
	DatabaseURL      string
	UiPath           string
	AuthDevBypass    bool
	AuthCookieSecure bool
	AuthSessionTTL   time.Duration
	AuthMagicLinkTTL time.Duration
	SiteURL          string
}

// ConfigRedacted holds safe fields for logging (excludes secrets like DatabaseURL).
type ConfigRedacted struct {
	Host             string
	Port             int
	Env              string
	LogLevel         string
	UiPath           string
	AuthDevBypass    bool
	AuthCookieSecure bool
	AuthSessionTTL   time.Duration
	AuthMagicLinkTTL time.Duration
	SiteURL          string
}

// Addr returns the "host:port" string for net.Listen().
func (c Config) Addr() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}

// Redacted returns a safe representation for logging (excludes DatabaseURL).
func (c Config) Redacted() ConfigRedacted {
	return ConfigRedacted{
		Host:             c.Host,
		Port:             c.Port,
		Env:              c.Env,
		LogLevel:         c.LogLevel,
		UiPath:           c.UiPath,
		AuthDevBypass:    c.AuthDevBypass,
		AuthCookieSecure: c.AuthCookieSecure,
		AuthSessionTTL:   c.AuthSessionTTL,
		AuthMagicLinkTTL: c.AuthMagicLinkTTL,
		SiteURL:          c.SiteURL,
	}
}

// FromEnv parses environment variables and returns a Config struct.
func FromEnv() (Config, error) {
	// Environment: determines deployment context and safe defaults
	environment := getEnvOrDefault(envKey, defaultEnv)
	if environment != "dev" && environment != "staging" && environment != "production" {
		return Config{}, fmt.Errorf("%w: must be 'dev', 'staging', or 'production', got %q", ErrInvalidEnv, environment)
	}

	// Compute context-aware defaults based on environment
	var (
		defaultCookieSecure bool
		defaultLogLevel     string
	)
	switch environment {
	case "production":
		defaultCookieSecure = true
		defaultLogLevel = "INFO"
	case "staging":
		defaultCookieSecure = true
		defaultLogLevel = "INFO"
	case "dev":
		defaultCookieSecure = false
		defaultLogLevel = "DEBUG"
	}

	// Parse PORT
	port, err := getEnvOrDefaultInt(portKey, defaultPort)
	if err != nil {
		return Config{}, err
	}
	if port < 0 || port > 65535 {
		return Config{}, fmt.Errorf("%w: %d", ErrPortOutOfRange, port)
	}

	// Parse HOST
	host := getEnvOrDefault(hostKey, defaultHost)

	// Parse UI_PATH
	uiPath := getEnvOrDefault(uiPathKey, defaultUiPath)

	// Parse LOG_LEVEL
	logLevel := getEnvOrDefault(logLevelKey, defaultLogLevel)

	// DATABASE_URL is required
	databaseURL, err := requireEnv(databaseURLKey)
	if err != nil {
		return Config{}, err
	}

	// Parse AUTH_DEV_BYPASS
	authDevBypass, err := parseBool(getEnvOrDefault(authDevBypassKey, "false"), defaultAuthDevBypass)
	if err != nil {
		return Config{}, fmt.Errorf("%s: %w", authDevBypassKey, err)
	}

	// Parse AUTH_COOKIE_SECURE
	authCookieSecure, err := parseBool(
		getEnvOrDefault(authCookieSecureKey, strconv.FormatBool(defaultCookieSecure)),
		defaultCookieSecure)
	if err != nil {
		return Config{}, fmt.Errorf("%s: %w", authCookieSecureKey, err)
	}

	// Parse AUTH_SESSION_TTL
	authSessionTTL, err := parseDuration(
		getEnvOrDefault(authSessionTTLKey, ""),
		defaultSessionTTL)
	if err != nil {
		return Config{}, fmt.Errorf("%s: %w", authSessionTTLKey, err)
	}

	// Parse AUTH_MAGIC_LINK_TTL
	authMagicLinkTTL, err := parseDuration(
		getEnvOrDefault(authMagicLinkTTLKey, ""),
		defaultMagicLinkTTL)
	if err != nil {
		return Config{}, fmt.Errorf("%s: %w", authMagicLinkTTLKey, err)
	}

	// Parse SITE_URL (computed from host:port if not set)
	SiteURL := strings.TrimSpace(getEnvOrDefault(SiteURLKey, ""))
	if SiteURL == "" && environment == "dev" {
		SiteURL = "http://" + net.JoinHostPort(host, strconv.Itoa(port))
	}
	if SiteURL == "" {
		return Config{}, fmt.Errorf("%s: must be provided for env %s: %w", SiteURLKey, environment, ErrMissingMagicLinkURL)
	}

	// Validate SITE_URL is a valid URL
	_, err = url.Parse(SiteURL)
	if err != nil {
		return Config{}, fmt.Errorf("%s: %w: %w", SiteURLKey, ErrInvalidURL, err)
	}

	// Cross-field validation: magic link TTL must be shorter than session TTL
	if authMagicLinkTTL >= authSessionTTL {
		return Config{}, fmt.Errorf("%w: %v >= %v", ErrMagicLinkTTLRange, authMagicLinkTTL, authSessionTTL)
	}

	cfg := Config{
		Host:             host,
		Port:             port,
		Env:              environment,
		LogLevel:         logLevel,
		DatabaseURL:      databaseURL,
		UiPath:           uiPath,
		AuthDevBypass:    authDevBypass,
		AuthCookieSecure: authCookieSecure,
		AuthSessionTTL:   authSessionTTL,
		AuthMagicLinkTTL: authMagicLinkTTL,
		SiteURL:          SiteURL,
	}

	return cfg, nil
}
