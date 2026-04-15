package auth

import (
	"context"
	"time"
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
	CreateUser(ctx context.Context, email string) (User, error)
	CreateOrUpdateMagicLinkToken(ctx context.Context, userID string, tokenHash TokenHash, expiresAt time.Time) (MagicLink, error)
	FindMagicLinkByTokenHash(ctx context.Context, tokenHash TokenHash) (MagicLink, error)
	FindSessionByTokenHash(ctx context.Context, tokenHash TokenHash) (Session, error)
	DeleteSession(ctx context.Context, sessionID string) error
	DeleteSessionsByUserID(ctx context.Context, userID string) error
	ConsumeMagicLinkAndCreateSession(ctx context.Context, tokenHash TokenHash, sessionTokenHash TokenHash, now time.Time, sessionExpiresAt time.Time) (User, Session, error)
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

// RequestMagicLink handles the magic link request flow
func (s *Service) RequestMagicLink(ctx context.Context, email string) (User, Token, error) {
	// Create or get user atomically using UPSERT
	// This single operation handles both new users and concurrent requests efficiently:
	// - First request: creates user and returns it
	// - Concurrent requests: all serialize on the unique constraint and get the same user
	// - Repeat requests: no-op update returns existing user
	// This is more efficient than SELECT + INSERT because it's a single database round-trip
	user, err := s.store.CreateUser(ctx, email)
	if err != nil {
		return User{}, "", err
	}

	// Generate magic link token
	tokenPair, err := newTokenPair()
	if err != nil {
		return User{}, "", err
	}

	// Create or update magic link (handles race conditions atomically in the database)
	expiresAt := s.clock.Now().Add(s.magicLinkTTL)
	_, err = s.store.CreateOrUpdateMagicLinkToken(ctx, user.ID, tokenPair.tokenHash, expiresAt)
	if err != nil {
		return User{}, "", err
	}

	return user, tokenPair.token, nil
}

// VerifyMagicLink validates a magic link token without consuming it.
//
// This is called by GET /v1/auth/magiclink/verify to allow users to preview
// which email they're signing in as before confirming. It does NOT consume the token
// to prevent automated email link scanners from burning valid tokens.
//
// The token is only consumed by ConsumeMagicLinkAndCreateSession (POST endpoint).
func (s *Service) VerifyMagicLink(ctx context.Context, token Token) (User, error) {

	tokenHash, err := decodeToken(token)
	if err != nil {
		return User{}, err
	}
	// Find the magic link and user in a single query
	magicLink, err := s.store.FindMagicLinkByTokenHash(ctx, tokenHash)
	if err != nil {
		return User{}, err
	}

	// Check expiration
	if s.clock.Now().After(magicLink.ExpiresAt) {
		return User{}, ErrMagicLinkExpired
	}

	// Construct user from magic link query (which joins users table)
	user := User{
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
func (s *Service) ConsumeMagicLinkAndCreateSession(ctx context.Context, token Token) (User, Token, error) {
	tokenHash, err := decodeToken(token)
	if err != nil {
		return User{}, "", err
	}
	// generate a session token
	sessionTokenPair, err := newTokenPair()
	if err != nil {
		return User{}, "", err
	}

	// Create or update magic link (handles race conditions atomically in the database)
	sessionsExpiresAt := s.clock.Now().Add(s.sessionTTL)

	user, _, err := s.store.ConsumeMagicLinkAndCreateSession(ctx, tokenHash, sessionTokenPair.tokenHash, s.clock.Now(), sessionsExpiresAt)
	if err != nil {
		return User{}, "", err
	}

	return user, sessionTokenPair.token, nil
}

// ValidateSessionToken checks if a session token is valid and returns the associated user
func (s *Service) ValidateSessionToken(ctx context.Context, token Token) (User, error) {
	// Decode and hash the token
	tokenHash, err := decodeToken(token)
	if err != nil {
		return User{}, err
	}

	// Find the session and user in a single query
	session, err := s.store.FindSessionByTokenHash(ctx, tokenHash)
	if err != nil {
		return User{}, err
	}

	// Check expiration
	if s.clock.Now().After(session.ExpiresAt) {
		return User{}, ErrSessionExpired
	}

	// Construct user from session query (which joins users table)
	user := User{
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
