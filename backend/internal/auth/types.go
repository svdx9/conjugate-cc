package auth

import (
	"errors"
	"time"
)

// User represents a user account
type User struct {
	ID        string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// MagicLink represents a magic link token for passwordless auth
type MagicLink struct {
	ID        string
	UserID    string
	TokenHash []byte
	ExpiresAt time.Time
	CreatedAt time.Time
}

// Session represents an authenticated session
type Session struct {
	ID        string
	UserID    string
	TokenHash []byte
	ExpiresAt time.Time
	CreatedAt time.Time
}

// Domain errors
var (
	ErrUserNotFound      = errors.New("user not found")
	ErrEmailTaken        = errors.New("email already taken")
	ErrMagicLinkNotFound = errors.New("magic link not found")
	ErrMagicLinkExpired  = errors.New("magic link expired")
	ErrSessionNotFound   = errors.New("session not found")
	ErrSessionExpired    = errors.New("session expired")
	ErrInvalidToken      = errors.New("invalid token")
	ErrTokenHashMismatch = errors.New("token hash does not match")
	ErrMagicLinkConsumed = errors.New("magic link already consumed")
)
