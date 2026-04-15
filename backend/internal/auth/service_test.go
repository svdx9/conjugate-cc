package auth_test

import (
	"bytes"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/svdx9/conjugate-cc/backend/internal/auth"
)

const (
	magicLinkTTL = 15 * time.Minute
	sessionTTL   = 30 * 24 * time.Hour
)

// MockClock implements the Clock interface for testing with controllable time
type MockClock struct {
	currentTime time.Time
}

func NewMockClock(t time.Time) *MockClock {
	return &MockClock{currentTime: t}
}

func (m *MockClock) Now() time.Time {
	return m.currentTime
}

func (m *MockClock) SetTime(t time.Time) {
	m.currentTime = t
}

// MockStore implements the Store interface for testing
type MockStore struct {
	users         map[string]auth.User
	magicLinks    map[string]auth.MagicLink
	consumedLinks map[string]bool // Track which links have been consumed
	sessions      map[string]auth.Session
}

func NewMockStore() *MockStore {
	return &MockStore{
		users:         make(map[string]auth.User),
		magicLinks:    make(map[string]auth.MagicLink),
		consumedLinks: make(map[string]bool),
		sessions:      make(map[string]auth.Session),
	}
}

func (m *MockStore) CreateUser(ctx context.Context, email string) (auth.User, error) {
	user := auth.User{
		ID:        "user-1",
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	m.users[email] = user
	return user, nil
}

func (m *MockStore) CreateOrUpdateMagicLinkToken(ctx context.Context, userID string, tokenHash auth.TokenHash, expiresAt time.Time) (auth.MagicLink, error) {
	// Upsert logic: check if unconsumed magic link exists for this user
	for id, ml := range m.magicLinks {
		// Skip consumed links - treat as if they don't exist
		if m.consumedLinks[id] {
			continue
		}
		// Check if it's for same user and unconsumed
		if ml.UserID == userID {
			// Update existing
			ml.TokenHash = tokenHash
			ml.ExpiresAt = expiresAt
			ml.CreatedAt = time.Now()
			m.magicLinks[id] = ml
			return ml, nil
		}
	}
	// Create new if doesn't exist
	magicLink := auth.MagicLink{
		ID:            "ml-1",
		UserID:        userID,
		TokenHash:     tokenHash,
		ExpiresAt:     expiresAt,
		CreatedAt:     time.Now(),
		Email:         "test@example.com",
		UserCreatedAt: time.Now(),
		UserUpdatedAt: time.Now(),
	}
	m.magicLinks["ml-1"] = magicLink
	return magicLink, nil
}

func (m *MockStore) FindMagicLinkByTokenHash(ctx context.Context, tokenHash auth.TokenHash) (auth.MagicLink, error) {
	for id, ml := range m.magicLinks {
		if bytes.Equal(ml.TokenHash, tokenHash) { // Simple comparison for testing
			// Check if already consumed
			if m.consumedLinks[id] {
				return auth.MagicLink{}, auth.ErrMagicLinkNotFound
			}
			return ml, nil
		}
	}
	return auth.MagicLink{}, auth.ErrMagicLinkNotFound
}

func (s *MockStore) ConsumeMagicLinkAndCreateSession(ctx context.Context, tokenHash auth.TokenHash, sessionTokenHash auth.TokenHash, now time.Time, sessionExpiresAt time.Time) (auth.User, auth.Session, error) {
	// Find and consume the magic link
	for id, ml := range s.magicLinks {
		// Skip already consumed links
		if s.consumedLinks[id] {
			continue
		}
		if bytes.Equal(ml.TokenHash, tokenHash) {
			// Mark as consumed
			s.consumedLinks[id] = true
			// Create session
			session := auth.Session{
				ID:            "sess-1",
				UserID:        ml.UserID,
				TokenHash:     sessionTokenHash,
				ExpiresAt:     sessionExpiresAt,
				CreatedAt:     now,
				Email:         ml.Email,
				UserCreatedAt: ml.UserCreatedAt,
				UserUpdatedAt: ml.UserUpdatedAt,
			}
			s.sessions["sess-1"] = session
			// Return user
			return auth.User{
				ID:        ml.UserID,
				Email:     ml.Email,
				CreatedAt: ml.UserCreatedAt,
				UpdatedAt: ml.UserUpdatedAt,
			}, session, nil
		}
	}
	return auth.User{}, auth.Session{}, auth.ErrMagicLinkNotFound
}

func (m *MockStore) CreateSession(ctx context.Context, userID string, tokenHash auth.TokenHash, expiresAt time.Time) (auth.Session, error) {
	s := auth.Session{
		ID:            "sess-1",
		UserID:        userID,
		TokenHash:     tokenHash,
		ExpiresAt:     expiresAt,
		CreatedAt:     time.Now(),
		Email:         "test@example.com",
		UserCreatedAt: time.Now(),
		UserUpdatedAt: time.Now(),
	}
	m.sessions["sess-1"] = s
	return s, nil
}

func (m *MockStore) FindSessionByTokenHash(ctx context.Context, tokenHash auth.TokenHash) (auth.Session, error) {
	for _, s := range m.sessions {
		if bytes.Equal(s.TokenHash, tokenHash) { // Simple comparison for testing
			return s, nil
		}
	}
	return auth.Session{}, auth.ErrSessionNotFound
}

func (m *MockStore) DeleteSession(ctx context.Context, sessionID string) error {
	_, ok := m.sessions[sessionID]
	if !ok {
		return auth.ErrSessionNotFound
	}
	delete(m.sessions, sessionID)
	return nil
}

func (m *MockStore) DeleteSessionsByUserID(ctx context.Context, userID string) error {
	for id, s := range m.sessions {
		if s.UserID == userID {
			delete(m.sessions, id)
		}
	}
	return nil
}

// Tests

func TestRequestMagicLink_NewUser(t *testing.T) {
	t.Parallel()
	store := NewMockStore()
	svc := auth.NewService(store, magicLinkTTL, sessionTTL)

	email := "test@example.com"
	user, token, err := svc.RequestMagicLink(context.Background(), email)
	if err != nil {
		t.Fatalf("RequestMagicLink failed: %v", err)
	}

	// User should be created
	if user.Email != email {
		t.Errorf("User email = %s, want %s", user.Email, email)
	}

	// Token should be generated
	if token == "" {
		t.Error("Token is empty")
	}

}

func TestRequestMagicLink_ExistingUser(t *testing.T) {
	t.Parallel()
	store := NewMockStore()
	svc := auth.NewService(store, magicLinkTTL, sessionTTL)

	email := "test@example.com"
	// Create initial user
	_, err := store.CreateUser(context.Background(), email)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	// Request magic link for same email
	user, token, err := svc.RequestMagicLink(context.Background(), email)
	if err != nil {
		t.Fatalf("RequestMagicLink failed: %v", err)
	}

	// User email should match
	if user.Email != email {
		t.Errorf("User email = %s, want %s", user.Email, email)
	}

	// Token should be generated
	if token == "" {
		t.Error("Token is empty")
	}
}

func TestRequestMagicLink_ConcurrentRequests(t *testing.T) {
	t.Parallel()
	// Test that concurrent requests for the same email both succeed
	// The UPSERT at the DB layer ensures they don't conflict
	store := NewMockStore()
	svc := auth.NewService(store, magicLinkTTL, sessionTTL)

	email := "test@example.com"

	// First request
	user1, token1, err := svc.RequestMagicLink(context.Background(), email)
	if err != nil {
		t.Fatalf("First RequestMagicLink failed: %v", err)
	}

	// Second request for same email (simulates concurrent request)
	user2, token2, err := svc.RequestMagicLink(context.Background(), email)
	if err != nil {
		t.Fatalf("Second RequestMagicLink failed: %v", err)
	}

	// Both should succeed and return same user
	if user1.Email != email {
		t.Errorf("User1 email = %s, want %s", user1.Email, email)
	}
	if user2.Email != email {
		t.Errorf("User2 email = %s, want %s", user2.Email, email)
	}
	if user1.ID != user2.ID {
		t.Errorf("User IDs differ: %s vs %s", user1.ID, user2.ID)
	}

	// Tokens should be different (newly generated)
	if token1 == token2 {
		t.Error("Tokens are identical, expected different tokens")
	}

	// Both tokens should be non-empty
	if token1 == "" {
		t.Error("Token1 is empty")
	}
	if token2 == "" {
		t.Error("Token2 is empty")
	}
}

// TestCreateSessionForUser was removed during refactoring.
// Users now authenticate via magic links, so this test is no longer valid.

func TestLogoutSession(t *testing.T) {
	t.Parallel()
	store := NewMockStore()
	svc := auth.NewService(store, magicLinkTTL, sessionTTL)

	// Create a session first
	userID := "user-1"
	_, err := store.CreateUser(context.Background(), "test@example.com")
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}
	_, err = store.CreateSession(context.Background(), userID, []byte("token-hash"), time.Now().Add(30*24*time.Hour))
	if err != nil {
		t.Fatalf("CreateSession failed: %v", err)
	}

	sessionID := "sess-1"
	err = svc.LogoutSession(context.Background(), sessionID)
	if err != nil {
		t.Fatalf("LogoutSession failed: %v", err)
	}
}

func TestLogoutAllSessions(t *testing.T) {
	t.Parallel()
	store := NewMockStore()
	svc := auth.NewService(store, magicLinkTTL, sessionTTL)

	userID := "user-1"
	err := svc.LogoutAllSessions(context.Background(), userID)
	if err != nil {
		t.Fatalf("LogoutAllSessions failed: %v", err)
	}
}

func TestVerifyMagicLink_Valid(t *testing.T) {
	t.Parallel()
	store := NewMockStore()
	clock := NewMockClock(time.Now())
	svc := auth.NewServiceWithClock(store, clock, magicLinkTTL, sessionTTL)

	email := "test@example.com"

	_, token, err := svc.RequestMagicLink(context.Background(), email)
	if err != nil {
		t.Fatalf("RequestMagicLink failed: %v", err)
	}

	user, err := svc.VerifyMagicLink(context.Background(), token)
	if err != nil {
		t.Fatalf("VerifyMagicLink failed: %v", err)
	}

	if user.Email != email {
		t.Errorf("User email = %s, want %s", user.Email, email)
	}

	user2, err := svc.VerifyMagicLink(context.Background(), token)
	if err != nil {
		t.Fatalf("Second VerifyMagicLink failed: %v", err)
	}
	if user2.Email != email {
		t.Errorf("Second call: User email = %s, want %s", user2.Email, email)
	}
}

func TestVerifyMagicLink_Expired(t *testing.T) {
	t.Parallel()
	store := NewMockStore()
	clock := NewMockClock(time.Now())
	svc := auth.NewServiceWithClock(store, clock, magicLinkTTL, sessionTTL)

	email := "test@example.com"
	futureTime := clock.Now().Add(magicLinkTTL + time.Hour)

	_, token, err := svc.RequestMagicLink(context.Background(), email)
	if err != nil {
		t.Fatalf("RequestMagicLink failed: %v", err)
	}

	// Set clock to future time to simulate expiration
	clock.SetTime(futureTime)

	_, err = svc.VerifyMagicLink(context.Background(), token)
	if !errors.Is(err, auth.ErrMagicLinkExpired) {
		t.Errorf("Expected ErrMagicLinkExpired, got %v", err)
	}
}

func TestConsumeMagicLinkAndCreateSession_Valid(t *testing.T) {
	t.Parallel()
	store := NewMockStore()
	svc := auth.NewService(store, magicLinkTTL, sessionTTL)

	email := "test@example.com"

	_, token, err := svc.RequestMagicLink(context.Background(), email)
	if err != nil {
		t.Fatalf("RequestMagicLink failed: %v", err)
	}

	user, sessionToken, err := svc.ConsumeMagicLinkAndCreateSession(context.Background(), token)
	if err != nil {
		t.Fatalf("ConsumeMagicLinkAndCreateSession failed: %v", err)
	}

	if user.Email != email {
		t.Errorf("User email = %s, want %s", user.Email, email)
	}

	if sessionToken == "" {
		t.Error("Session token is empty")
	}
}

func TestConsumeMagicLinkAndCreateSession_TokenNotFound(t *testing.T) {
	t.Parallel()
	store := NewMockStore()
	svc := auth.NewService(store, magicLinkTTL, sessionTTL)

	email := "test@example.com"

	_, token, err := svc.RequestMagicLink(context.Background(), email)
	if err != nil {
		t.Fatalf("RequestMagicLink failed: %v", err)
	}

	// First consume should succeed
	_, _, err = svc.ConsumeMagicLinkAndCreateSession(context.Background(), token)
	if err != nil {
		t.Fatalf("First ConsumeMagicLinkAndCreateSession failed: %v", err)
	}

	// Second consume should fail (token already consumed)
	_, _, err = svc.ConsumeMagicLinkAndCreateSession(context.Background(), token)
	if !errors.Is(err, auth.ErrMagicLinkNotFound) {
		t.Errorf("Expected ErrMagicLinkNotFound, got %v", err)
	}
}
