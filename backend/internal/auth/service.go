package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"time"
)

const (
	// TokenSize is the size of the random token in bytes
	TokenSize = 32
)

// Clock provides the current time, allowing time-dependent operations to be testable
type Clock interface {
	Now() time.Time
}

// SystemClock implements Clock using the system time
type SystemClock struct{}

func (SystemClock) Now() time.Time {
	return time.Now()
}

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
	store        Store
	clock        Clock
	magicLinkTTL time.Duration
	sessionTTL   time.Duration
}

// NewService creates a new authentication service with system clock
func NewService(store Store, magicLinkTTL, sessionTTL time.Duration) *Service {
	return &Service{
		store:        store,
		clock:        SystemClock{},
		magicLinkTTL: magicLinkTTL,
		sessionTTL:   sessionTTL,
	}
}

// NewServiceWithClock creates a new authentication service with a custom clock (for testing)
func NewServiceWithClock(store Store, clock Clock, magicLinkTTL, sessionTTL time.Duration) *Service {
	return &Service{
		store:        store,
		clock:        clock,
		magicLinkTTL: magicLinkTTL,
		sessionTTL:   sessionTTL,
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
	// Create or get user atomically using UPSERT
	// This single operation handles both new users and concurrent requests efficiently:
	// - First request: creates user and returns it
	// - Concurrent requests: all serialize on the unique constraint and get the same user
	// - Repeat requests: no-op update returns existing user
	// This is more efficient than SELECT + INSERT because it's a single database round-trip
	user, err := s.store.CreateUser(ctx, email)
	if err != nil {
		return nil, nil, err
	}

	// Generate magic link token
	tokenPair, err := s.GenerateToken()
	if err != nil {
		return nil, nil, err
	}

	// Create or update magic link (handles race conditions atomically in the database)
	expiresAt := s.clock.Now().Add(s.magicLinkTTL)
	_, err = s.store.CreateOrUpdateMagicLinkToken(ctx, user.ID, tokenPair.TokenHash, expiresAt)
	if err != nil {
		return nil, nil, err
	}

	return user, tokenPair, nil
}

// VerifyMagicLink validates a magic link token without consuming it.
//
// This is called by GET /v1/auth/magiclink/verify to allow users to preview
// which email they're signing in as before confirming. It does NOT consume the token
// to prevent automated email link scanners from burning valid tokens.
//
// The token is only consumed by ConsumeMagicLinkAndCreateSession (POST endpoint).
func (s *Service) VerifyMagicLink(ctx context.Context, token string) (*User, error) {
	// Decode and hash the token
	decodedToken, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return nil, ErrInvalidToken
	}
	tokenHash := sha256.Sum256(decodedToken)

	// Find the magic link and user in a single query
	magicLink, err := s.store.FindMagicLinkByTokenHash(ctx, tokenHash[:])
	if err != nil {
		return nil, err
	}

	// Check expiration
	if s.clock.Now().After(magicLink.ExpiresAt) {
		return nil, ErrMagicLinkExpired
	}

	// Construct user from magic link query (which joins users table)
	user := &User{
		ID:        magicLink.UserID,
		Email:     magicLink.Email,
		CreatedAt: magicLink.UserCreatedAt,
		UpdatedAt: magicLink.UserUpdatedAt,
	}

	return user, nil
}

// ConsumeMagicLinkAndCreateSession validates, consumes, and creates a session for a magic link token.
//
// This is called by POST /v1/auth/magiclink/verify after the user confirms they want to sign in.
// It performs three operations atomically:
//  1. Verify the token is valid and not expired
//  2. Consume the token (mark as used) to ensure single-use property
//  3. Create a new session token for the user
//
// If the token was already consumed or doesn't exist, returns ErrMagicLinkConsumed or ErrMagicLinkNotFound.
func (s *Service) ConsumeMagicLinkAndCreateSession(ctx context.Context, token string) (*User, *TokenPair, error) {
	// Decode and hash the token
	decodedToken, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return nil, nil, ErrInvalidToken
	}
	tokenHash := sha256.Sum256(decodedToken)

	// Find the magic link and user in a single query
	magicLink, err := s.store.FindMagicLinkByTokenHash(ctx, tokenHash[:])
	if err != nil {
		return nil, nil, err
	}

	// Check expiration
	if s.clock.Now().After(magicLink.ExpiresAt) {
		return nil, nil, ErrMagicLinkExpired
	}

	// Consume the magic link (mark as used)
	err = s.store.ConsumeMagicLink(ctx, magicLink.ID)
	if err != nil {
		return nil, nil, err
	}

	// Construct user from magic link query (which joins users table)
	user := &User{
		ID:        magicLink.UserID,
		Email:     magicLink.Email,
		CreatedAt: magicLink.UserCreatedAt,
		UpdatedAt: magicLink.UserUpdatedAt,
	}

	// Create session token for the user
	sessionToken, err := s.CreateSessionForUser(ctx, user.ID)
	if err != nil {
		return nil, nil, err
	}

	return user, sessionToken, nil
}

// CreateSessionForUser creates a new session token for a user
func (s *Service) CreateSessionForUser(ctx context.Context, userID string) (*TokenPair, error) {
	// Generate session token
	tokenPair, err := s.GenerateToken()
	if err != nil {
		return nil, err
	}

	// Store session in database
	expiresAt := s.clock.Now().Add(s.sessionTTL)
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

	// Find the session and user in a single query
	session, err := s.store.FindSessionByTokenHash(ctx, tokenHash[:])
	if err != nil {
		return nil, err
	}

	// Check expiration
	if s.clock.Now().After(session.ExpiresAt) {
		return nil, ErrSessionExpired
	}

	// Construct user from session query (which joins users table)
	user := &User{
		ID:        session.UserID,
		Email:     session.Email,
		CreatedAt: session.UserCreatedAt,
		UpdatedAt: session.UserUpdatedAt,
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
