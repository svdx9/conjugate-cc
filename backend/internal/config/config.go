package config

import (
	"errors"
	"fmt"
	"log/slog"
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
	debugKey            = "DEBUG"
	uiPathKey           = "UI_PATH"
	authDevBypassKey    = "AUTH_DEV_BYPASS"
	authCookieSecureKey = "AUTH_COOKIE_SECURE"
	authSessionTTLKey   = "AUTH_SESSION_TTL"
	authMagicLinkTTLKey = "AUTH_MAGIC_LINK_TTL"
	siteURLKey          = "SITE_URL"
)

const (
	defaultPort          = 8080
	defaultEnv           = Dev
	defaultHost          = "0.0.0.0"
	defaultUiPath        = "ui"
	defaultSessionTTL    = 30 * 24 * time.Hour
	defaultMagicLinkTTL  = 15 * time.Minute
	defaultAuthDevBypass = false
	defaultSiteURL       = ""
	minMagicLinkTTL      = 1 * time.Minute
	maxMagicLinkTTL      = 60 * time.Minute
	minSessionTTL        = 1 * time.Hour
	maxSessionTTL        = 8760 * time.Hour // 1 year
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
	ErrMagicLinkTTLBounds  = errors.New("magic link TTL must be between 1m0s and 60m0s")
	ErrSessionTTLBounds    = errors.New("session TTL must be between 1h0m0s and 8760h0m0s")
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
	LogLevel         slog.Level
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
	LogLevel         slog.Level
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
		defaultLogLevel     slog.Level
	)
	switch environment {
	case "production":
		defaultCookieSecure = true
		defaultLogLevel = slog.LevelInfo
	case "staging":
		defaultCookieSecure = true
		defaultLogLevel = slog.LevelInfo
	case "dev":
		defaultCookieSecure = false
		defaultLogLevel = slog.LevelDebug
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

	// Parse DEBUG flag for debug logging
	debugMode, err := parseBool(getEnvOrDefault(debugKey, "false"), false)
	if err != nil {
		return Config{}, fmt.Errorf("%s: %w", debugKey, err)
	}

	// Determine log level: DEBUG if flag set, otherwise environment default
	var logLevel slog.Level
	if debugMode {
		logLevel = slog.LevelDebug
	} else {
		logLevel = defaultLogLevel
	}

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

	// Parse SITE_URL (in dev mode, default to http://localhost:port)
	siteURL := strings.TrimSpace(getEnvOrDefault(siteURLKey, defaultSiteURL))
	if environment == "dev" && siteURL == defaultSiteURL {
		siteURL = "http://" + net.JoinHostPort("localhost", strconv.Itoa(port))
	}
	if siteURL == "" {
		return Config{}, fmt.Errorf("%s: must be provided for env %s: %w", siteURLKey, environment, ErrMissingMagicLinkURL)
	}

	// Validate SITE_URL is a valid URL
	parsedURL, err := url.Parse(siteURL)
	if err != nil {
		return Config{}, fmt.Errorf("%s: %w: %w", siteURLKey, ErrInvalidURL, err)
	}
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return Config{}, fmt.Errorf("%s: must be a valid HTTP URL: %w", siteURLKey, ErrInvalidURL)
	}
	if parsedURL.Host == "" {
		return Config{}, fmt.Errorf("%s: must be a valid HTTP URL: %w", siteURLKey, ErrInvalidURL)
	}

	// Cross-field validation: magic link TTL must be shorter than session TTL
	if authMagicLinkTTL >= authSessionTTL {
		return Config{}, fmt.Errorf("%w: %v >= %v", ErrMagicLinkTTLRange, authMagicLinkTTL, authSessionTTL)
	}

	// Validate TTL bounds
	if authMagicLinkTTL < minMagicLinkTTL || authMagicLinkTTL > maxMagicLinkTTL {
		return Config{}, ErrMagicLinkTTLBounds
	}
	if authSessionTTL < minSessionTTL || authSessionTTL > maxSessionTTL {
		return Config{}, ErrSessionTTLBounds
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
		SiteURL:          siteURL,
	}

	return cfg, nil
}
