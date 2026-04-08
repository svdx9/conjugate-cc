package config

import (
	"testing"
)

func TestFromEnv(t *testing.T) {
	t.Run("default values", func(t *testing.T) {
		// Clear environment variables
		t.Setenv("PORT", "")
		t.Setenv("HOST", "")
		t.Setenv("ENV", "")

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
	})

	t.Run("custom values", func(t *testing.T) {
		t.Setenv("PORT", "3000")
		t.Setenv("HOST", "localhost")
		t.Setenv("ENV", "production")

		cfg, err := FromEnv()
		if err != nil {
			t.Fatalf("FromEnv() error = %v, wantErr false", err)
		}

		if cfg.Port != 3000 {
			t.Errorf("Port = %d, want %d", cfg.Port, 3000)
		}

		if cfg.Host != "localhost" {
			t.Errorf("Host = %s, want %s", cfg.Host, "localhost")
		}

		if cfg.Env != "production" {
			t.Errorf("Env = %s, want %s", cfg.Env, "production")
		}
	})

	t.Run("invalid port", func(t *testing.T) {
		t.Setenv("PORT", "invalid")

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

		cfg, err := FromEnv()
		if err != nil {
			t.Fatalf("FromEnv() error = %v, wantErr false", err)
		}

		if cfg.Host != "example.com" {
			t.Errorf("Host = %s, want %s", cfg.Host, "example.com")
		}
	})
}


