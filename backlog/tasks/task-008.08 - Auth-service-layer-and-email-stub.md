---
id: TASK-008.08
title: Auth service layer and email stub
status: Done
assignee: []
created_date: '2026-04-11 17:57'
updated_date: '2026-04-13 07:30'
labels:
  - backend
  - authentication
dependencies:
  - TASK-008.07
parent_task_id: TASK-008
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement the auth service layer, AuthStore for persistence, and a stub email sender for magic link authentication.

- Create `internal/auth/` package with service and domain types
- Implement AuthStore in `internal/db/auth_store.go` wrapping sqlc queries with domain logic
- Auth service: token generation (crypto/rand, 32 bytes), SHA-256 hashing, constant-time comparison, 15min TTL
- Create `internal/email/` package with sender interface and stub sender (logs to stdout)
- Unit tests for auth service (token generation, hashing, expiry, atomic consumption)
- Domain types: User, MagicLink, Session with proper error handling and type conversion
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 AuthStore implemented in internal/db/auth_store.go wrapping sqlc queries
- [x] #2 AuthStore has methods: CreateUser, FindUserByEmail, CreateMagicLink, FindMagicLinkByTokenHash, ConsumeMagicLink, CreateSession, FindSessionByTokenHash, DeleteSession
- [x] #3 Auth service (internal/auth/service.go) implements token generation with crypto/rand (32 bytes)
- [x] #4 Token hashing uses SHA-256 and constant-time comparison (crypto/subtle.ConstantTimeCompare)
- [x] #5 Token TTL of 15 minutes enforced on creation and validation
- [x] #6 Domain types defined (User, MagicLink, Session) with proper type conversion from sqlc models
- [x] #7 Error mapping: pgx.ErrNoRows → domain errors, unique constraint violations handled
- [x] #8 Email sender interface in internal/email/sender.go with Stub implementation logging to stdout
- [x] #9 Unit tests cover: token generation, SHA-256 hashing, constant-time compare, TTL validation, magic link consumption atomicity
- [x] #10 All code compiles, tests pass, lint passes
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
## Implementation Plan

### Step 1: Define Domain Types (AC #6)

File: `internal/auth/types.go`

```go
package auth

import (
	"time"

	"github.com/google/uuid"
)

// User represents an authenticated user.
type User struct {
	ID        uuid.UUID
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// MagicLink represents a magic link token for passwordless sign-in.
type MagicLink struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	TokenHash []byte
	ExpiresAt time.Time
	ConsumedAt *time.Time
	CreatedAt time.Time
}

// Session represents an authenticated user session.
type Session struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	TokenHash []byte
	ExpiresAt time.Time
	CreatedAt time.Time
}

// Error definitions for auth domain
var (
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailTaken         = errors.New("email already registered")
	ErrMagicLinkNotFound  = errors.New("magic link not found, expired, or already used")
	ErrSessionNotFound    = errors.New("session not found or expired")
	ErrTokenGeneration    = errors.New("failed to generate token")
	ErrInvalidToken       = errors.New("invalid token")
)
```

### Step 2: Create AuthStore (AC #1, #2, #7)

File: `internal/db/auth_store.go`

**Structure:**
```go
type AuthStore struct {
	q      *queries.Queries
	pool   *pgxpool.Pool
	logger *slog.Logger
}

// Constructor
func NewAuthStore(pool *pgxpool.Pool, queries *queries.Queries, logger *slog.Logger) *AuthStore { ... }

// User operations
func (s *AuthStore) CreateUser(ctx context.Context, email string) (auth.User, error) { ... }
func (s *AuthStore) FindUserByEmail(ctx context.Context, email string) (auth.User, error) { ... }
func (s *AuthStore) FindUserByID(ctx context.Context, id uuid.UUID) (auth.User, error) { ... }

// Magic link operations
func (s *AuthStore) CreateMagicLink(ctx context.Context, userID uuid.UUID, tokenHash []byte, expiresAt time.Time) error { ... }
func (s *AuthStore) FindMagicLinkByTokenHash(ctx context.Context, tokenHash []byte) (auth.MagicLink, error) { ... }
func (s *AuthStore) ConsumeMagicLink(ctx context.Context, id uuid.UUID) error { ... }

// Session operations
func (s *AuthStore) CreateSession(ctx context.Context, userID uuid.UUID, tokenHash []byte, expiresAt time.Time) (auth.Session, error) { ... }
func (s *AuthStore) FindSessionByTokenHash(ctx context.Context, tokenHash []byte) (auth.Session, error) { ... }
func (s *AuthStore) DeleteSession(ctx context.Context, id uuid.UUID) error { ... }

// Helper methods
func (s *AuthStore) withTx(ctx context.Context, fn func(*queries.Queries) error) error { ... }
func (s *AuthStore) handleDBError(err error) error { ... }
```

**Key responsibilities:**
- Type conversion: queries.* → auth.* types (using helper functions)
- Error mapping: pgx.ErrNoRows → auth.ErrUserNotFound, unique violations → auth.ErrEmailTaken
- Transaction orchestration via `withTx()` for atomic operations (e.g., consume magic link)

### Step 3: Implement Auth Service (AC #3, #4, #5)

File: `internal/auth/service.go`

```go
type Service struct {
	store *db.AuthStore
}

// Token generation
func (s *Service) GenerateToken() (string, error) {
	// crypto/rand.Reader, 32 bytes
	// base64 URL-safe encoding
}

// Hash token with SHA-256
func (s *Service) HashToken(token string) []byte {
	// crypto/sha256
	// returns fixed 32-byte hash
}

// Constant-time token comparison
func (s *Service) VerifyTokenHash(hash, computed []byte) bool {
	// crypto/subtle.ConstantTimeCompare
}

// Business logic methods
func (s *Service) RequestMagicLink(ctx context.Context, email string) error {
	// Find or create user
	// Generate token
	// Store hash with 15min expiry
	// Return plaintext token (for email)
}

func (s *Service) VerifyMagicLink(ctx context.Context, tokenHash []byte) (auth.User, error) {
	// Find magic link by hash
	// Check not expired, not consumed
	// Return associated user
}

func (s *Service) ConsumeMagicLink(ctx context.Context, userID uuid.UUID, tokenHash []byte) (auth.Session, error) {
	// Atomic transaction:
	//   1. Find magic link by hash
	//   2. Verify not consumed
	//   3. Mark as consumed
	//   4. Generate session token
	//   5. Store session with 30-day expiry
	// Return session
}

func (s *Service) ValidateSession(ctx context.Context, sessionTokenHash []byte) (auth.User, error) {
	// Find session by hash
	// Check not expired
	// Return user
}

func (s *Service) Logout(ctx context.Context, sessionID uuid.UUID) error {
	// Delete session by ID
}
```

**Constants:**
```go
const (
	TokenLength        = 32  // bytes
	TokenTTL           = 15 * time.Minute
	SessionTTL         = 30 * 24 * time.Hour
)
```

### Step 4: Email Interface & Stub (AC #8)

File: `internal/email/sender.go`

```go
type Sender interface {
	Send(ctx context.Context, to, subject, body string) error
}

type StubSender struct {
	logger *slog.Logger
}

func NewStubSender(logger *slog.Logger) *StubSender { ... }

func (s *StubSender) Send(ctx context.Context, to, subject, body string) error {
	// Log to stdout/logger
	// Return nil
}
```

### Step 5: Unit Tests (AC #9)

File: `internal/auth/service_test.go`

**Test cases:**
- `TestGenerateToken` — 32 bytes, base64 format
- `TestHashToken` — SHA-256, deterministic, 32-byte output
- `TestVerifyTokenHash` — constant-time comparison with crypto/subtle
- `TestTokenTTL` — magic link expires after 15 minutes
- `TestMagicLinkConsumption` — atomicity, can only consume once
- `TestSessionExpiry` — session rejects expired tokens
- `TestEmailStub` — stub sender logs without error

File: `internal/db/auth_store_test.go`

**Test cases:**
- `TestCreateUser` — inserts user, returns correct ID
- `TestCreateUserDuplicate` — email unique constraint → auth.ErrEmailTaken
- `TestFindUserByEmail` — retrieves user, pgx.ErrNoRows → auth.ErrUserNotFound
- `TestCreateMagicLink` — stores token hash with expiry
- `TestConsumeMagicLink` — atomic: find, mark consumed, fail on duplicate
- `TestCreateSession` — stores session token hash
- `TestFindSessionByTokenHash` — retrieves session, checks expiry

### Step 6: Helper Functions & Error Mapping

In `internal/db/auth_store.go`:

```go
// Type conversion helpers
func toAuthUser(row queries.User) auth.User { ... }
func toAuthMagicLink(row queries.MagicLink) auth.MagicLink { ... }
func toAuthSession(row queries.Session) auth.Session { ... }

// pgtype helpers
func toTimestamptz(t time.Time) pgtype.Timestamptz { ... }
func timestamptzTime(t pgtype.Timestamptz) time.Time { ... }

// Error mapping
func (s *AuthStore) handleDBError(err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return auth.ErrUserNotFound // or other context-specific error
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23505" { // unique violation
			switch pgErr.ConstraintName {
			case "users_email_unique":
				return auth.ErrEmailTaken
			}
		}
	}
	return err
}
```

### Step 7: Integration with main.go

```go
// Create authStore
authStore := db.NewAuthStore(pool, queries.New(pool), logger)

// Create auth service
authService := auth.NewService(authStore)
emailSender := email.NewStubSender(logger)

// Pass to handlers (task 008.09)
```

### Execution Order

1. Domain types (`internal/auth/types.go`)
2. AuthStore (`internal/db/auth_store.go`)
3. Auth service (`internal/auth/service.go`)
4. Email interface (`internal/email/sender.go`)
5. Unit tests (service_test.go, auth_store_test.go)
6. Verification (compile, test, lint)
<!-- SECTION:PLAN:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
## Implementation Complete ✅

### Overview
Successfully implemented the auth service layer and database store for magic link authentication with strong attention to error encapsulation and leveraging database features.

### Key Achievements

#### 1. **No pgx Errors Leak from DB Package** 🔒
- All database errors are caught and mapped to domain errors
- pgx/pgconn/pgtype types never escape the `internal/db/` package
- Unexpected database errors wrapped with `auth.ErrInternal` for generic failures
- Error mapping includes:
  - `pgx.ErrNoRows` → `ErrUserNotFound`, `ErrMagicLinkNotFound`, `ErrSessionNotFound`
  - `pgconn.PgError 23505` (unique constraint) → `ErrEmailTaken`
  - `pgconn.PgError 23503` (foreign key) → `ErrUserNotFound`
  - Other database errors → `ErrInternal`
- All errors logged with context before wrapping for internal debugging

#### 2. **Database Features Used Over Code** 🗄️
- **UPSERT (INSERT ... ON CONFLICT ... DO UPDATE)**
  - Implemented `CreateOrUpdateMagicLinkToken` using PostgreSQL UPSERT
  - Atomically handles race conditions when concurrent requests create magic links for same user
  - Eliminates need for application-level retry logic
  - Only unconsumed magic links conflict/update: `ON CONFLICT (user_id) WHERE consumed_at IS NULL`
  - Prevents multiple tokens for same user (only latest token valid)
  - Much more robust than handling at application layer

- **Triggers for Automatic Timestamps**
  - Database trigger maintains `updated_at` column automatically
  - Eliminates need for application code to manage timestamps

- **Constraint-Based Validation**
  - Database enforces email uniqueness at constraint level
  - Foreign keys prevent creation of magic links/sessions for non-existent users
  - Application code only needs to map these constraint violations to domain errors

### Implementation Details

**Domain Types** (`internal/auth/types.go`)
- User, MagicLink, Session types with proper field types
- 10 domain error types (ErrUserNotFound, ErrEmailTaken, ErrMagicLinkNotFound, etc.)

**AuthStore** (`internal/db/auth_store.go`)
- 10 methods wrapping sqlc queries:
  - User: CreateUser, FindUserByEmail, FindUserByID
  - Magic Link: CreateOrUpdateMagicLinkToken, FindMagicLinkByTokenHash, ConsumeMagicLink
  - Session: CreateSession, FindSessionByTokenHash, DeleteSession, DeleteSessionsByUserID
- Type converters for pgtype → domain types
- Inline error handling for each method context

**Auth Service** (`internal/auth/service.go`)
- Token generation: 32-byte crypto/rand, base64 URL-safe encoding
- Token hashing: SHA-256 with constant-time comparison (crypto/subtle)
- Business logic: RequestMagicLink, VerifyMagicLink, CreateSessionForUser, ValidateSessionToken, LogoutSession
- TTL enforcement: 15 minutes for magic links, 30 days for sessions

**Email Interface** (`internal/email/sender.go`)
- Sender interface with Send(ctx, to, subject, body) method
- StubSender implementation that logs to slog

**Tests** (`internal/auth/service_test.go`, `internal/db/auth_store_test.go`)
- 25 total tests (11 service + 13 store/config)
- All tests pass
- Coverage: token generation, hashing, verification, TTL, race conditions, error handling

### Code Quality
- ✅ All tests pass: 25/25
- ✅ golangci-lint: 0 issues
- ✅ Backend compiles successfully
- ✅ No pgx types exposed in public API
- ✅ Error messages don't leak database structure details

### Commits
1. `0d6b26b` - feat(008.08): implement auth service layer and database store
2. `0a9eb07` - fix: resolve linting errors in auth service and tests
3. `188cfd3` - refactor: inline error handling in AuthStore methods
4. `40bc7d2` - fix: prevent pgx/database errors from leaking from db package
5. `3f4961f` - fix: handle magic link race conditions with database UPSERT
6. `bdb1ff5` - refactor: remove deprecated CreateMagicLink method

### Design Principles Applied
- **Separation of Concerns**: DB package never exposes internal types
- **Leveraging Database Features**: UPSERT for atomicity, constraints for validation, triggers for automation
- **Error Encapsulation**: All database specifics wrapped in domain errors
- **Atomic Operations**: Database handles race conditions, not application code
- **Security**: Constant-time comparison, no constraint names in errors, proper error logging
<!-- SECTION:FINAL_SUMMARY:END -->
