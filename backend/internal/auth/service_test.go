package auth_test

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"testing"
	"time"

	"github.com/svdx9/conjugate-cc/backend/internal/auth"
)

const (
	magicLinkTTL = 15 * time.Minute
	sessionTTL   = 30 * 24 * time.Hour
)

// MockClock implements the Clock interface for testing with fixed time
type MockClock struct {
	currentTime time.Time
}

func NewMockClock(t time.Time) *MockClock {
	return &MockClock{currentTime: t}
}

func (m *MockClock) Now() time.Time {
	return m.currentTime
}

// MockStore implements the Store interface for testing
type MockStore struct {
	users         map[string]*auth.User
	magicLinks    map[string]*auth.MagicLink
	consumedLinks map[string]bool // Track which links have been consumed
	sessions      map[string]*auth.Session
}

func NewMockStore() *MockStore {
	return &MockStore{
		users:         make(map[string]*auth.User),
		magicLinks:    make(map[string]*auth.MagicLink),
		consumedLinks: make(map[string]bool),
		sessions:      make(map[string]*auth.Session),
	}
}

func (m *MockStore) CreateUser(ctx context.Context, email string) (*auth.User, error) {
	user := &auth.User{
		ID:        "user-1",
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	m.users[email] = user
	return user, nil
}

func (m *MockStore) FindUserByEmail(ctx context.Context, email string) (*auth.User, error) {
	user, ok := m.users[email]
	if !ok {
		return nil, auth.ErrUserNotFound
	}
	return user, nil
}

func (m *MockStore) FindUserByID(ctx context.Context, userID string) (*auth.User, error) {
	for _, user := range m.users {
		if user.ID == userID {
			return user, nil
		}
	}
	return nil, auth.ErrUserNotFound
}

func (m *MockStore) CreateMagicLink(ctx context.Context, userID string, tokenHash []byte, expiresAt time.Time) (*auth.MagicLink, error) {
	magicLink := &auth.MagicLink{
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

func (m *MockStore) CreateOrUpdateMagicLinkToken(ctx context.Context, userID string, tokenHash []byte, expiresAt time.Time) (*auth.MagicLink, error) {
	// Upsert logic: check if unconsumed magic link exists for this user
	for id, ml := range m.magicLinks {
		// Check if it's for same user and unconsumed (we don't have consumed_at in mock, so check if not marked as consumed)
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
	magicLink := &auth.MagicLink{
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

func (m *MockStore) FindMagicLinkByTokenHash(ctx context.Context, tokenHash []byte) (*auth.MagicLink, error) {
	for id, ml := range m.magicLinks {
		if bytes.Equal(ml.TokenHash, tokenHash) { // Simple comparison for testing
			// Check if already consumed
			if m.consumedLinks[id] {
				return nil, auth.ErrMagicLinkNotFound
			}
			return ml, nil
		}
	}
	return nil, auth.ErrMagicLinkNotFound
}

func (m *MockStore) ConsumeMagicLink(ctx context.Context, magicLinkID string) error {
	_, ok := m.magicLinks[magicLinkID]
	if !ok {
		return auth.ErrMagicLinkNotFound
	}
	// Check if already consumed
	if m.consumedLinks[magicLinkID] {
		return auth.ErrMagicLinkConsumed
	}
	// Mark as consumed
	m.consumedLinks[magicLinkID] = true
	return nil
}

func (m *MockStore) CreateSession(ctx context.Context, userID string, tokenHash []byte, expiresAt time.Time) (*auth.Session, error) {
	s := &auth.Session{
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

func (m *MockStore) FindSessionByTokenHash(ctx context.Context, tokenHash []byte) (*auth.Session, error) {
	for _, s := range m.sessions {
		if bytes.Equal(s.TokenHash, tokenHash) { // Simple comparison for testing
			return s, nil
		}
	}
	return nil, auth.ErrSessionNotFound
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

func TestGenerateToken(t *testing.T) {
	t.Parallel()
	store := NewMockStore()
	svc := auth.NewService(store, magicLinkTTL, sessionTTL)

	pair, err := svc.GenerateToken()
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}

	// Token should be non-empty
	if pair.Token == "" {
		t.Error("Token is empty")
	}

	// TokenHash should be non-empty and match SHA-256 hash
	if len(pair.TokenHash) != 32 { // SHA-256 produces 32 bytes
		t.Errorf("TokenHash length = %d, want 32", len(pair.TokenHash))
	}

	// Verify the hash is correct
	decodedToken, _ := base64.URLEncoding.DecodeString(pair.Token)
	expectedHash := sha256.Sum256(decodedToken)
	for i, b := range expectedHash[:] {
		if pair.TokenHash[i] != b {
			t.Errorf("TokenHash mismatch at position %d", i)
			break
		}
	}

	// Two generated tokens should be different
	pair2, _ := svc.GenerateToken()
	if pair.Token == pair2.Token {
		t.Error("Two generated tokens are identical, expected unique tokens")
	}
}

func TestVerifyToken_Valid(t *testing.T) {
	t.Parallel()
	store := NewMockStore()
	svc := auth.NewService(store, magicLinkTTL, sessionTTL)

	// Generate a token
	pair, _ := svc.GenerateToken()

	// Verify the token
	err := svc.VerifyToken(pair.Token, pair.TokenHash)
	if err != nil {
		t.Errorf("VerifyToken failed: %v", err)
	}
}

func TestVerifyToken_Invalid(t *testing.T) {
	t.Parallel()
	store := NewMockStore()
	svc := auth.NewService(store, magicLinkTTL, sessionTTL)

	// Generate a token
	pair, _ := svc.GenerateToken()

	// Try to verify with wrong token
	err := svc.VerifyToken("invalid-token", pair.TokenHash)
	if !errors.Is(err, auth.ErrInvalidToken) {
		t.Errorf("VerifyToken with invalid token = %v, want %v", err, auth.ErrInvalidToken)
	}
}

func TestVerifyToken_MismatchHash(t *testing.T) {
	t.Parallel()
	store := NewMockStore()
	svc := auth.NewService(store, magicLinkTTL, sessionTTL)

	// Generate two tokens
	pair1, _ := svc.GenerateToken()
	pair2, _ := svc.GenerateToken()

	// Try to verify pair1's token with pair2's hash
	err := svc.VerifyToken(pair1.Token, pair2.TokenHash)
	if !errors.Is(err, auth.ErrTokenHashMismatch) {
		t.Errorf("VerifyToken with mismatched hash = %v, want %v", err, auth.ErrTokenHashMismatch)
	}
}

func TestRequestMagicLink_NewUser(t *testing.T) {
	t.Parallel()
	store := NewMockStore()
	svc := auth.NewService(store, magicLinkTTL, sessionTTL)

	email := "test@example.com"
	user, tokenPair, err := svc.RequestMagicLink(context.Background(), email)
	if err != nil {
		t.Fatalf("RequestMagicLink failed: %v", err)
	}

	// User should be created
	if user.Email != email {
		t.Errorf("User email = %s, want %s", user.Email, email)
	}

	// Token should be generated
	if tokenPair.Token == "" {
		t.Error("Token is empty")
	}

	// TokenHash should be generated
	if len(tokenPair.TokenHash) != 32 {
		t.Errorf("TokenHash length = %d, want 32", len(tokenPair.TokenHash))
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
	user, tokenPair, err := svc.RequestMagicLink(context.Background(), email)
	if err != nil {
		t.Fatalf("RequestMagicLink failed: %v", err)
	}

	// User email should match
	if user.Email != email {
		t.Errorf("User email = %s, want %s", user.Email, email)
	}

	// Token should be generated
	if tokenPair.Token == "" {
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
	if token1.Token == token2.Token {
		t.Error("Tokens are identical, expected different tokens")
	}

	// Both tokens should be non-empty
	if token1.Token == "" {
		t.Error("Token1 is empty")
	}
	if token2.Token == "" {
		t.Error("Token2 is empty")
	}
}

func TestCreateSessionForUser(t *testing.T) {
	t.Parallel()
	store := NewMockStore()
	svc := auth.NewService(store, magicLinkTTL, sessionTTL)

	userID := "user-123"
	tokenPair, err := svc.CreateSessionForUser(context.Background(), userID)
	if err != nil {
		t.Fatalf("CreateSessionForUser failed: %v", err)
	}

	// Token should be generated
	if tokenPair.Token == "" {
		t.Error("Token is empty")
	}

	// TokenHash should be generated
	if len(tokenPair.TokenHash) != 32 {
		t.Errorf("TokenHash length = %d, want 32", len(tokenPair.TokenHash))
	}
}

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
	svc := auth.NewService(store, magicLinkTTL, sessionTTL)

	// Generate a token
	tokenPair, err := svc.GenerateToken()
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}

	// Create a magic link in the store
	expiresAt := time.Now().Add(magicLinkTTL)
	_, err = store.CreateOrUpdateMagicLinkToken(context.Background(), "user-1", tokenPair.TokenHash, expiresAt)
	if err != nil {
		t.Fatalf("CreateOrUpdateMagicLinkToken failed: %v", err)
	}

	// Verify the magic link (should NOT consume it)
	user, err := svc.VerifyMagicLink(context.Background(), tokenPair.Token)
	if err != nil {
		t.Fatalf("VerifyMagicLink failed: %v", err)
	}

	if user == nil {
		t.Fatal("VerifyMagicLink returned nil user")
	}
	if user.Email != "test@example.com" {
		t.Errorf("Expected email test@example.com, got %s", user.Email)
	}

	// Verify token was NOT consumed by checking if we can find it again
	user2, err := svc.VerifyMagicLink(context.Background(), tokenPair.Token)
	if err != nil {
		t.Fatalf("Second VerifyMagicLink call failed (token was consumed): %v", err)
	}
	if user2 == nil {
		t.Fatal("Second VerifyMagicLink returned nil user (token was consumed)")
	}
}

func TestVerifyMagicLink_Expired(t *testing.T) {
	t.Parallel()
	store := NewMockStore()
	svc := auth.NewService(store, magicLinkTTL, sessionTTL)

	// Generate a token with expired TTL
	tokenPair, err := svc.GenerateToken()
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}

	// Create a magic link that expires immediately
	expiresAt := time.Now().Add(-1 * time.Second)
	_, err = store.CreateOrUpdateMagicLinkToken(context.Background(), "user-1", tokenPair.TokenHash, expiresAt)
	if err != nil {
		t.Fatalf("CreateOrUpdateMagicLinkToken failed: %v", err)
	}

	// Try to verify the expired link
	_, err = svc.VerifyMagicLink(context.Background(), tokenPair.Token)
	if !errors.Is(err, auth.ErrMagicLinkExpired) {
		t.Errorf("Expected ErrMagicLinkExpired, got %v", err)
	}
}

func TestConsumeMagicLinkAndCreateSession_Valid(t *testing.T) {
	t.Parallel()
	store := NewMockStore()
	svc := auth.NewService(store, magicLinkTTL, sessionTTL)

	// Generate a token
	tokenPair, err := svc.GenerateToken()
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}

	// Create a magic link in the store
	expiresAt := time.Now().Add(magicLinkTTL)
	_, err = store.CreateOrUpdateMagicLinkToken(context.Background(), "user-1", tokenPair.TokenHash, expiresAt)
	if err != nil {
		t.Fatalf("CreateOrUpdateMagicLinkToken failed: %v", err)
	}

	// Consume the link and create session
	user, sessionToken, err := svc.ConsumeMagicLinkAndCreateSession(context.Background(), tokenPair.Token)
	if err != nil {
		t.Fatalf("ConsumeMagicLinkAndCreateSession failed: %v", err)
	}

	if user == nil {
		t.Fatal("ConsumeMagicLinkAndCreateSession returned nil user")
	}
	if user.Email != "test@example.com" {
		t.Errorf("Expected email test@example.com, got %s", user.Email)
	}

	if sessionToken == nil {
		t.Fatal("ConsumeMagicLinkAndCreateSession returned nil session token")
	}
	if sessionToken.Token == "" {
		t.Fatal("ConsumeMagicLinkAndCreateSession returned empty session token")
	}

	// Verify token WAS consumed by trying to use it again
	// The store returns ErrMagicLinkNotFound for consumed tokens (security: don't reveal token was consumed)
	_, _, err = svc.ConsumeMagicLinkAndCreateSession(context.Background(), tokenPair.Token)
	if !errors.Is(err, auth.ErrMagicLinkNotFound) {
		t.Errorf("Expected ErrMagicLinkNotFound on second use (token was consumed), got %v", err)
	}
}

func TestConsumeMagicLinkAndCreateSession_TokenNotFound(t *testing.T) {
	t.Parallel()
	store := NewMockStore()
	svc := auth.NewService(store, magicLinkTTL, sessionTTL)

	// Generate a token that doesn't exist in the store
	tokenPair, err := svc.GenerateToken()
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}

	// Try to consume a non-existent token
	_, _, err = svc.ConsumeMagicLinkAndCreateSession(context.Background(), tokenPair.Token)
	if !errors.Is(err, auth.ErrMagicLinkNotFound) {
		t.Errorf("Expected ErrMagicLinkNotFound, got %v", err)
	}
}
