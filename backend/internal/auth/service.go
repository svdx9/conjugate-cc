package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"time"
)

const (
	// MagicLinkTTL is the time-to-live for magic links (15 minutes)
	MagicLinkTTL = 15 * time.Minute
	// SessionTTL is the time-to-live for sessions (30 days)
	SessionTTL = 30 * 24 * time.Hour
	// TokenSize is the size of the random token in bytes
	TokenSize = 32
)

// Store defines the interface for database operations
type Store interface {
	CreateUser(ctx context.Context, email string) (*User, error)
	FindUserByEmail(ctx context.Context, email string) (*User, error)
	FindUserByID(ctx context.Context, userID string) (*User, error)
	CreateOrUpdateMagicLinkToken(ctx context.Context, userID string, tokenHash []byte, expiresAt time.Time) (*MagicLink, error)
	FindMagicLinkByTokenHash(ctx context.Context, tokenHash []byte) (*MagicLink, error)
	ConsumeMagicLink(ctx context.Context, magicLinkID string) error
	CreateSession(ctx context.Context, userID string, tokenHash []byte, expiresAt time.Time) (*Session, error)
	FindSessionByTokenHash(ctx context.Context, tokenHash []byte) (*Session, error)
	DeleteSession(ctx context.Context, sessionID string) error
	DeleteSessionsByUserID(ctx context.Context, userID string) error
}

// Service handles authentication business logic
type Service struct {
	store Store
}

// NewService creates a new authentication service
func NewService(store Store) *Service {
	return &Service{
		store: store,
	}
}

// TokenPair represents a token and its hash
type TokenPair struct {
	Token     string // Plain text token (sent to user)
	TokenHash []byte // SHA-256 hash (stored in database)
}

// GenerateToken creates a new random token and returns both the plain token and its hash
func (s *Service) GenerateToken() (*TokenPair, error) {
	// Generate random bytes
	randomBytes := make([]byte, TokenSize)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	// Encode to base64 URL-safe format (sent to user)
	token := base64.URLEncoding.EncodeToString(randomBytes)

	// Hash the token (stored in database)
	hash := sha256.Sum256(randomBytes)

	return &TokenPair{
		Token:     token,
		TokenHash: hash[:],
	}, nil
}

// VerifyToken checks if a plain token matches a stored hash using constant-time comparison
func (s *Service) VerifyToken(plainToken string, storedHash []byte) error {
	if plainToken == "" {
		return ErrInvalidToken
	}

	// Decode the plain token from base64
	decodedToken, err := base64.URLEncoding.DecodeString(plainToken)
	if err != nil {
		return ErrInvalidToken
	}

	// Hash the decoded token
	hash := sha256.Sum256(decodedToken)

	// Compare using constant-time comparison to prevent timing attacks
	if subtle.ConstantTimeCompare(hash[:], storedHash) != 1 {
		return ErrTokenHashMismatch
	}

	return nil
}

// RequestMagicLink handles the magic link request flow
func (s *Service) RequestMagicLink(ctx context.Context, email string) (*User, *TokenPair, error) {
	// Try to find existing user
	user, err := s.store.FindUserByEmail(ctx, email)
	if err != nil && !errors.Is(err, ErrUserNotFound) {
		return nil, nil, err
	}

	// Create user if doesn't exist
	if user == nil {
		user, err = s.store.CreateUser(ctx, email)
		if err != nil {
			return nil, nil, err
		}
	}

	// Generate magic link token
	tokenPair, err := s.GenerateToken()
	if err != nil {
		return nil, nil, err
	}

	// Create or update magic link (handles race conditions atomically in the database)
	expiresAt := time.Now().Add(MagicLinkTTL)
	_, err = s.store.CreateOrUpdateMagicLinkToken(ctx, user.ID, tokenPair.TokenHash, expiresAt)
	if err != nil {
		return nil, nil, err
	}

	return user, tokenPair, nil
}

// VerifyMagicLink validates a magic link token and consumes it
func (s *Service) VerifyMagicLink(ctx context.Context, token string) (*User, error) {
	// Decode and hash the token
	decodedToken, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return nil, ErrInvalidToken
	}
	tokenHash := sha256.Sum256(decodedToken)

	// Find the magic link
	magicLink, err := s.store.FindMagicLinkByTokenHash(ctx, tokenHash[:])
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return nil, ErrMagicLinkNotFound
		}
		return nil, err
	}

	// Check expiration
	if time.Now().After(magicLink.ExpiresAt) {
		return nil, ErrMagicLinkExpired
	}

	// Get user details
	user, err := s.store.FindUserByID(ctx, magicLink.UserID)
	if err != nil {
		return nil, err
	}

	// Consume the magic link
	err = s.store.ConsumeMagicLink(ctx, magicLink.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// CreateSessionForUser creates a new session token for a user
func (s *Service) CreateSessionForUser(ctx context.Context, userID string) (*TokenPair, error) {
	// Generate session token
	tokenPair, err := s.GenerateToken()
	if err != nil {
		return nil, err
	}

	// Store session in database
	expiresAt := time.Now().Add(SessionTTL)
	_, err = s.store.CreateSession(ctx, userID, tokenPair.TokenHash, expiresAt)
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}

// ValidateSessionToken checks if a session token is valid and returns the associated user
func (s *Service) ValidateSessionToken(ctx context.Context, token string) (*User, error) {
	// Decode and hash the token
	decodedToken, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return nil, ErrInvalidToken
	}
	tokenHash := sha256.Sum256(decodedToken)

	// Find the session
	session, err := s.store.FindSessionByTokenHash(ctx, tokenHash[:])
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return nil, ErrSessionNotFound
		}
		return nil, err
	}

	// Check expiration
	if time.Now().After(session.ExpiresAt) {
		return nil, ErrSessionExpired
	}

	// Get user details
	user, err := s.store.FindUserByID(ctx, session.UserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// LogoutSession removes a session by ID
func (s *Service) LogoutSession(ctx context.Context, sessionID string) error {
	return s.store.DeleteSession(ctx, sessionID)
}

// LogoutAllSessions removes all sessions for a user
func (s *Service) LogoutAllSessions(ctx context.Context, userID string) error {
	return s.store.DeleteSessionsByUserID(ctx, userID)
}
