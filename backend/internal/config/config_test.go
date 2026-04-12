package config

import (
	"errors"
	"testing"
	"time"
)

func TestFromEnv(t *testing.T) {
	t.Run("default values", func(t *testing.T) {
		// Clear environment variables
		t.Setenv("PORT", "")
		t.Setenv("HOST", "")
		t.Setenv("ENV", "")
		t.Setenv("DATABASE_URL", "postgres://localhost/test")
		t.Setenv("LOG_LEVEL", "")
		t.Setenv("UI_PATH", "")
		t.Setenv("AUTH_DEV_BYPASS", "")
		t.Setenv("AUTH_COOKIE_SECURE", "")
		t.Setenv("AUTH_SESSION_TTL", "")
		t.Setenv("AUTH_MAGIC_LINK_TTL", "")
		t.Setenv("AUTH_MAGIC_LINK_URL", "")

		cfg, err := FromEnv()
		if err != nil {
			t.Fatalf("FromEnv() error = %v, wantErr false", err)
		}

		if cfg.Port != defaultPort {
			t.Errorf("Port = %d, want %d", cfg.Port, defaultPort)
		}

		if cfg.Host != defaultHost {
			t.Errorf("Host = %s, want %s", cfg.Host, defaultHost)
		}

		if cfg.Env != defaultEnv {
			t.Errorf("Env = %s, want %s", cfg.Env, defaultEnv)
		}

		if cfg.LogLevel != "DEBUG" {
			t.Errorf("LogLevel = %s, want DEBUG", cfg.LogLevel)
		}

		if cfg.UiPath != defaultUiPath {
			t.Errorf("UiPath = %s, want %s", cfg.UiPath, defaultUiPath)
		}

		if cfg.AuthDevBypass != defaultAuthDevBypass {
			t.Errorf("AuthDevBypass = %v, want %v", cfg.AuthDevBypass, defaultAuthDevBypass)
		}

		if cfg.AuthCookieSecure != false {
			t.Errorf("AuthCookieSecure = %v, want false (dev default)", cfg.AuthCookieSecure)
		}

		if cfg.AuthSessionTTL != defaultSessionTTL {
			t.Errorf("AuthSessionTTL = %v, want %v", cfg.AuthSessionTTL, defaultSessionTTL)
		}

		if cfg.AuthMagicLinkTTL != defaultMagicLinkTTL {
			t.Errorf("AuthMagicLinkTTL = %v, want %v", cfg.AuthMagicLinkTTL, defaultMagicLinkTTL)
		}

		if cfg.AuthMagicLinkURL != "http://0.0.0.0:8080/magic-link" {
			t.Errorf("AuthMagicLinkURL = %s, want computed default", cfg.AuthMagicLinkURL)
		}
	})

	t.Run("custom values", func(t *testing.T) {
		t.Setenv("PORT", "3000")
		t.Setenv("HOST", "localhost")
		t.Setenv("ENV", "production")
		t.Setenv("DATABASE_URL", "postgres://prod-db:5432/app")
		t.Setenv("LOG_LEVEL", "INFO")
		t.Setenv("UI_PATH", "assets")
		t.Setenv("AUTH_DEV_BYPASS", "true")
		t.Setenv("AUTH_COOKIE_SECURE", "false")
		t.Setenv("AUTH_SESSION_TTL", "168h")
		t.Setenv("AUTH_MAGIC_LINK_TTL", "10m")
		t.Setenv("AUTH_MAGIC_LINK_URL", "https://example.com/auth")

		cfg, err := FromEnv()
		if err != nil {
			t.Fatalf("FromEnv() error = %v, wantErr false", err)
		}

		if cfg.Port != 3000 {
			t.Errorf("Port = %d, want 3000", cfg.Port)
		}

		if cfg.Host != "localhost" {
			t.Errorf("Host = %s, want localhost", cfg.Host)
		}

		if cfg.Env != "production" {
			t.Errorf("Env = %s, want production", cfg.Env)
		}

		if cfg.DatabaseURL != "postgres://prod-db:5432/app" {
			t.Errorf("DatabaseURL = %s, want postgres://prod-db:5432/app", cfg.DatabaseURL)
		}

		if cfg.LogLevel != "INFO" {
			t.Errorf("LogLevel = %s, want INFO", cfg.LogLevel)
		}

		if cfg.UiPath != "assets" {
			t.Errorf("UiPath = %s, want assets", cfg.UiPath)
		}

		if cfg.AuthDevBypass != true {
			t.Errorf("AuthDevBypass = %v, want true", cfg.AuthDevBypass)
		}

		if cfg.AuthCookieSecure != false {
			t.Errorf("AuthCookieSecure = %v, want false", cfg.AuthCookieSecure)
		}

		if cfg.AuthSessionTTL != 168*time.Hour {
			t.Errorf("AuthSessionTTL = %v, want 168h", cfg.AuthSessionTTL)
		}

		if cfg.AuthMagicLinkTTL != 10*time.Minute {
			t.Errorf("AuthMagicLinkTTL = %v, want 10m", cfg.AuthMagicLinkTTL)
		}

		if cfg.AuthMagicLinkURL != "https://example.com/auth" {
			t.Errorf("AuthMagicLinkURL = %s, want https://example.com/auth", cfg.AuthMagicLinkURL)
		}
	})

	t.Run("invalid port", func(t *testing.T) {
		t.Setenv("PORT", "invalid")
		t.Setenv("DATABASE_URL", "postgres://localhost/test")

		_, err := FromEnv()
		if err == nil {
			t.Fatal("FromEnv() error = nil, wantErr true")
		}

		// Should not wrap with ErrPortOutOfRange
		if err.Error() != "PORT: strconv.Atoi: parsing \"invalid\": invalid syntax" {
			t.Errorf("FromEnv() error = %v, want specific parsing error", err)
		}
	})

	t.Run("port out of range", func(t *testing.T) {
		t.Setenv("PORT", "70000")
		t.Setenv("DATABASE_URL", "postgres://localhost/test")

		_, err := FromEnv()
		if err == nil {
			t.Fatal("FromEnv() error = nil, wantErr true")
		}

		// Should wrap with ErrPortOutOfRange
		if err.Error() != "port out of range: 70000" {
			t.Errorf("FromEnv() error = %v, want range error", err)
		}
	})

	t.Run("invalid host", func(t *testing.T) {
		t.Setenv("HOST", "invalid host!")
		t.Setenv("DATABASE_URL", "postgres://localhost/test")
		// Must provide explicit magic link URL since the computed URL would be invalid
		t.Setenv("AUTH_MAGIC_LINK_URL", "http://localhost:8080/magic-link")

		cfg, err := FromEnv()
		if err != nil {
			t.Fatalf("FromEnv() error = %v, wantErr false", err)
		}

		// Should accept any host value - validation happens at server startup
		if cfg.Host != "invalid host!" {
			t.Errorf("Host = %s, want %s", cfg.Host, "invalid host!")
		}
	})

	t.Run("valid IP host", func(t *testing.T) {
		t.Setenv("HOST", "192.168.1.1")
		t.Setenv("DATABASE_URL", "postgres://localhost/test")

		cfg, err := FromEnv()
		if err != nil {
			t.Fatalf("FromEnv() error = %v, wantErr false", err)
		}

		if cfg.Host != "192.168.1.1" {
			t.Errorf("Host = %s, want %s", cfg.Host, "192.168.1.1")
		}
	})

	t.Run("valid hostname", func(t *testing.T) {
		t.Setenv("HOST", "example.com")
		t.Setenv("DATABASE_URL", "postgres://localhost/test")

		cfg, err := FromEnv()
		if err != nil {
			t.Fatalf("FromEnv() error = %v, wantErr false", err)
		}

		if cfg.Host != "example.com" {
			t.Errorf("Host = %s, want %s", cfg.Host, "example.com")
		}
	})

	t.Run("missing DATABASE_URL", func(t *testing.T) {
		t.Setenv("DATABASE_URL", "")

		_, err := FromEnv()
		if err == nil {
			t.Fatal("FromEnv() error = nil, wantErr true")
		}

		if !errors.Is(err, ErrMissingEnv) {
			t.Errorf("FromEnv() error = %v, want ErrMissingEnv", err)
		}
	})

	t.Run("invalid AUTH_COOKIE_SECURE", func(t *testing.T) {
		t.Setenv("DATABASE_URL", "postgres://localhost/test")
		t.Setenv("AUTH_COOKIE_SECURE", "maybe")

		_, err := FromEnv()
		if err == nil {
			t.Fatal("FromEnv() error = nil, wantErr true")
		}

		if !errors.Is(err, ErrInvalidBool) {
			t.Errorf("FromEnv() error = %v, want ErrInvalidBool", err)
		}
	})

	t.Run("invalid AUTH_SESSION_TTL", func(t *testing.T) {
		t.Setenv("DATABASE_URL", "postgres://localhost/test")
		t.Setenv("AUTH_SESSION_TTL", "not-a-duration")

		_, err := FromEnv()
		if err == nil {
			t.Fatal("FromEnv() error = nil, wantErr true")
		}

		if !errors.Is(err, ErrInvalidDuration) {
			t.Errorf("FromEnv() error = %v, want ErrInvalidDuration", err)
		}
	})

	t.Run("invalid AUTH_MAGIC_LINK_TTL", func(t *testing.T) {
		t.Setenv("DATABASE_URL", "postgres://localhost/test")
		t.Setenv("AUTH_MAGIC_LINK_TTL", "bad-duration")

		_, err := FromEnv()
		if err == nil {
			t.Fatal("FromEnv() error = nil, wantErr true")
		}

		if !errors.Is(err, ErrInvalidDuration) {
			t.Errorf("FromEnv() error = %v, want ErrInvalidDuration", err)
		}
	})

	t.Run("magic link TTL >= session TTL", func(t *testing.T) {
		t.Setenv("DATABASE_URL", "postgres://localhost/test")
		t.Setenv("AUTH_SESSION_TTL", "30m")
		t.Setenv("AUTH_MAGIC_LINK_TTL", "30m")

		_, err := FromEnv()
		if err == nil {
			t.Fatal("FromEnv() error = nil, wantErr true")
		}

		if !errors.Is(err, ErrMagicLinkTTLRange) {
			t.Errorf("FromEnv() error = %v, want ErrMagicLinkTTLRange", err)
		}
	})

	t.Run("invalid AUTH_MAGIC_LINK_URL", func(t *testing.T) {
		t.Setenv("DATABASE_URL", "postgres://localhost/test")
		t.Setenv("AUTH_MAGIC_LINK_URL", "not a valid url://broken[")

		_, err := FromEnv()
		if err == nil {
			t.Fatal("FromEnv() error = nil, wantErr true")
		}

		if !errors.Is(err, ErrInvalidURL) {
			t.Errorf("FromEnv() error = %v, want ErrInvalidURL", err)
		}
	})

	t.Run("invalid ENV value", func(t *testing.T) {
		t.Setenv("DATABASE_URL", "postgres://localhost/test")
		t.Setenv("ENV", "invalid-env")

		_, err := FromEnv()
		if err == nil {
			t.Fatal("FromEnv() error = nil, wantErr true")
		}

		if !errors.Is(err, ErrInvalidEnv) {
			t.Errorf("FromEnv() error = %v, want ErrInvalidEnv", err)
		}
	})

	t.Run("auth cookie secure true in production", func(t *testing.T) {
		t.Setenv("DATABASE_URL", "postgres://localhost/test")
		t.Setenv("ENV", "production")
		t.Setenv("AUTH_COOKIE_SECURE", "")

		cfg, err := FromEnv()
		if err != nil {
			t.Fatalf("FromEnv() error = %v, wantErr false", err)
		}

		if cfg.AuthCookieSecure != true {
			t.Errorf("AuthCookieSecure = %v, want true (production default)", cfg.AuthCookieSecure)
		}
	})

	t.Run("log level info in production", func(t *testing.T) {
		t.Setenv("DATABASE_URL", "postgres://localhost/test")
		t.Setenv("ENV", "production")
		t.Setenv("LOG_LEVEL", "")

		cfg, err := FromEnv()
		if err != nil {
			t.Fatalf("FromEnv() error = %v, wantErr false", err)
		}

		if cfg.LogLevel != "INFO" {
			t.Errorf("LogLevel = %s, want INFO (production default)", cfg.LogLevel)
		}
	})

	t.Run("Addr() method", func(t *testing.T) {
		t.Setenv("PORT", "9000")
		t.Setenv("HOST", "localhost")
		t.Setenv("DATABASE_URL", "postgres://localhost/test")

		cfg, err := FromEnv()
		if err != nil {
			t.Fatalf("FromEnv() error = %v, wantErr false", err)
		}

		addr := cfg.Addr()
		if addr != "localhost:9000" {
			t.Errorf("Addr() = %s, want localhost:9000", addr)
		}
	})

	t.Run("Redacted() excludes DatabaseURL", func(t *testing.T) {
		t.Setenv("DATABASE_URL", "postgres://user:password@host/db")

		cfg, err := FromEnv()
		if err != nil {
			t.Fatalf("FromEnv() error = %v, wantErr false", err)
		}

		redacted := cfg.Redacted()
		// Redacted() should not include DatabaseURL field (it's not in ConfigRedacted struct)
		// Verify other safe fields are present
		if redacted.Host != cfg.Host {
			t.Errorf("Redacted().Host = %s, want %s", redacted.Host, cfg.Host)
		}

		if redacted.Port != cfg.Port {
			t.Errorf("Redacted().Port = %d, want %d", redacted.Port, cfg.Port)
		}

		if redacted.Env != cfg.Env {
			t.Errorf("Redacted().Env = %s, want %s", redacted.Env, cfg.Env)
		}
	})
}
