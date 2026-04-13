package db

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/svdx9/conjugate-cc/backend/internal/auth"
	"github.com/svdx9/conjugate-cc/backend/internal/db/queries"
)

// PostgreSQL error codes
const (
	pgErrCodeUniqueViolation     = "23505" // unique constraint violation
	pgErrCodeForeignKeyViolation = "23503" // foreign key constraint violation
)

// AuthStore handles all authentication-related database operations
type AuthStore struct {
	queries *queries.Queries
	pool    *pgxpool.Pool
	logger  *slog.Logger
}

// NewAuthStore creates a new AuthStore
func NewAuthStore(pool *pgxpool.Pool, logger *slog.Logger) *AuthStore {
	return &AuthStore{
		queries: queries.New(pool),
		pool:    pool,
		logger:  logger,
	}
}

// CreateUser creates a new user with the given email
func (s *AuthStore) CreateUser(ctx context.Context, email string) (*auth.User, error) {
	// CreateUser is an UPSERT operation
	row, err := s.queries.CreateUser(ctx, email)
	if err != nil {
		s.logger.Error("failed to create user", "email", email, "error", err)
		return nil, auth.ErrInternal
	}
	return toAuthUser(&row), nil
}

// FindUserByEmail finds a user by their email address
func (s *AuthStore) FindUserByEmail(ctx context.Context, email string) (*auth.User, error) {
	row, err := s.queries.FindUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, auth.ErrUserNotFound
		}
		s.logger.Error("failed to find user by email", "email", email, "error", err)
		return nil, auth.ErrInternal
	}
	return toAuthUser(&row), nil
}

// FindUserByID finds a user by their ID
func (s *AuthStore) FindUserByID(ctx context.Context, userID string) (*auth.User, error) {
	uid, err := parseUUID(userID)
	if err != nil {
		return nil, auth.ErrUserNotFound
	}
	row, err := s.queries.GetUserByID(ctx, uid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, auth.ErrUserNotFound
		}
		s.logger.Error("failed to find user by id", "user_id", userID, "error", err)
		return nil, auth.ErrInternal
	}
	return toAuthUser(&row), nil
}

// CreateOrUpdateMagicLinkToken creates or updates a magic link token for a user
// This handles race conditions atomically at the database level using UPSERT:
// - If no unconsumed magic link exists for this user, creates a new one
// - If an unconsumed magic link exists, updates its token and expiration
// This prevents multiple concurrent requests from creating conflicting tokens
func (s *AuthStore) CreateOrUpdateMagicLinkToken(ctx context.Context, userID string, tokenHash []byte, expiresAt time.Time) (*auth.MagicLink, error) {
	uid, err := parseUUID(userID)
	if err != nil {
		return nil, auth.ErrUserNotFound
	}
	row, err := s.queries.CreateOrUpdateMagicLinkToken(ctx, queries.CreateOrUpdateMagicLinkTokenParams{
		UserID:    uid,
		TokenHash: tokenHash,
		ExpiresAt: timestamptzFromTime(expiresAt),
	})
	if err != nil {
		// Check for foreign key constraint violation (user doesn't exist)
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgErrCodeForeignKeyViolation {
			return nil, auth.ErrUserNotFound
		}
		s.logger.Error("failed to create or update magic link token", "user_id", userID, "error", err)
		return nil, auth.ErrInternal
	}
	return toAuthMagicLink(&row), nil
}

// FindMagicLinkByTokenHash finds an unconsumed, non-expired magic link by its token hash
func (s *AuthStore) FindMagicLinkByTokenHash(ctx context.Context, tokenHash []byte) (*auth.MagicLink, error) {
	row, err := s.queries.FindMagicLinkByTokenHash(ctx, tokenHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, auth.ErrMagicLinkNotFound
		}
		s.logger.Error("failed to find magic link by token hash", "error", err)
		return nil, auth.ErrInternal
	}
	return toAuthMagicLinkRow(&row), nil
}

// ConsumeMagicLink marks a magic link as consumed
// Returns ErrMagicLinkConsumed if the magic link was already consumed (no rows updated)
func (s *AuthStore) ConsumeMagicLink(ctx context.Context, magicLinkID string) error {
	mlid, err := parseUUID(magicLinkID)
	if err != nil {
		return auth.ErrMagicLinkNotFound
	}
	_, err = s.queries.ConsumeMagicLink(ctx, mlid)
	if err != nil {
		// Check if this is a "no rows" scenario (magic link already consumed)
		if errors.Is(err, pgx.ErrNoRows) {
			return auth.ErrMagicLinkConsumed
		}
		s.logger.Error("failed to consume magic link", "magic_link_id", magicLinkID, "error", err)
		return auth.ErrInternal
	}
	return nil
}

// CreateSession creates a new session token for a user
func (s *AuthStore) CreateSession(ctx context.Context, userID string, tokenHash []byte, expiresAt time.Time) (*auth.Session, error) {
	uid, err := parseUUID(userID)
	if err != nil {
		return nil, auth.ErrUserNotFound
	}
	row, err := s.queries.CreateSession(ctx, queries.CreateSessionParams{
		UserID:    uid,
		TokenHash: tokenHash,
		ExpiresAt: timestamptzFromTime(expiresAt),
	})
	if err != nil {
		// Check for foreign key constraint violation (user doesn't exist)
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgErrCodeForeignKeyViolation {
			return nil, auth.ErrUserNotFound
		}
		s.logger.Error("failed to create session", "user_id", userID, "error", err)
		return nil, auth.ErrInternal
	}
	return toAuthSession(&row), nil
}

// FindSessionByTokenHash finds a non-expired session by its token hash
func (s *AuthStore) FindSessionByTokenHash(ctx context.Context, tokenHash []byte) (*auth.Session, error) {
	row, err := s.queries.FindSessionByTokenHash(ctx, tokenHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, auth.ErrSessionNotFound
		}
		s.logger.Error("failed to find session by token hash", "error", err)
		return nil, auth.ErrInternal
	}
	return toAuthSessionRow(&row), nil
}

// DeleteSession deletes a session by its ID
func (s *AuthStore) DeleteSession(ctx context.Context, sessionID string) error {
	sid, err := parseUUID(sessionID)
	if err != nil {
		return auth.ErrSessionNotFound
	}
	err = s.queries.DeleteSession(ctx, sid)
	if err != nil {
		s.logger.Error("failed to delete session", "session_id", sessionID, "error", err)
		return auth.ErrInternal
	}
	return nil
}

// DeleteSessionsByUserID deletes all sessions for a user
func (s *AuthStore) DeleteSessionsByUserID(ctx context.Context, userID string) error {
	uid, err := parseUUID(userID)
	if err != nil {
		return auth.ErrUserNotFound
	}
	err = s.queries.DeleteSessionsByUserID(ctx, uid)
	if err != nil {
		s.logger.Error("failed to delete sessions by user id", "user_id", userID, "error", err)
		return auth.ErrInternal
	}
	return nil
}

// Type conversion helpers

// parseUUID parses a string UUID into pgtype.UUID
func parseUUID(s string) (pgtype.UUID, error) {
	var uuid pgtype.UUID
	err := uuid.Scan(s)
	return uuid, err
}

// timestamptzFromTime converts a time.Time to pgtype.Timestamptz
func timestamptzFromTime(t time.Time) pgtype.Timestamptz {
	var ts pgtype.Timestamptz
	// pgtype.Timestamptz.Scan() should not fail with time.Time, but we should handle errors properly
	err := ts.Scan(t)
	if err != nil {
		// This should never happen with valid time.Time input, but log it if it does
		panic("failed to scan time: " + err.Error())
	}
	return ts
}

// toAuthUser converts a sqlc User to an auth.User
func toAuthUser(u *queries.User) *auth.User {
	return &auth.User{
		ID:        u.ID.String(),
		Email:     u.Email,
		CreatedAt: u.CreatedAt.Time,
		UpdatedAt: u.UpdatedAt.Time,
	}
}

// toAuthMagicLink converts a sqlc MagicLink to an auth.MagicLink
func toAuthMagicLink(ml *queries.MagicLink) *auth.MagicLink {
	return &auth.MagicLink{
		ID:            ml.ID.String(),
		UserID:        ml.UserID.String(),
		TokenHash:     ml.TokenHash,
		ExpiresAt:     ml.ExpiresAt.Time,
		CreatedAt:     ml.CreatedAt.Time,
		Email:         "",
		UserCreatedAt: time.Time{},
		UserUpdatedAt: time.Time{},
	}
}

// toAuthMagicLinkRow converts a sqlc FindMagicLinkByTokenHashRow to an auth.MagicLink
func toAuthMagicLinkRow(ml *queries.FindMagicLinkByTokenHashRow) *auth.MagicLink {
	return &auth.MagicLink{
		ID:            ml.ID.String(),
		UserID:        ml.UserID.String(),
		TokenHash:     ml.TokenHash,
		ExpiresAt:     ml.ExpiresAt.Time,
		CreatedAt:     ml.CreatedAt.Time,
		Email:         ml.Email,
		UserCreatedAt: ml.UserCreatedAt.Time,
		UserUpdatedAt: ml.UserUpdatedAt.Time,
	}
}

// toAuthSession converts a sqlc Session to an auth.Session
func toAuthSession(s *queries.Session) *auth.Session {
	return &auth.Session{
		ID:            s.ID.String(),
		UserID:        s.UserID.String(),
		TokenHash:     s.TokenHash,
		ExpiresAt:     s.ExpiresAt.Time,
		CreatedAt:     s.CreatedAt.Time,
		Email:         "",
		UserCreatedAt: time.Time{},
		UserUpdatedAt: time.Time{},
	}
}

// toAuthSessionRow converts a sqlc FindSessionByTokenHashRow to an auth.Session
func toAuthSessionRow(s *queries.FindSessionByTokenHashRow) *auth.Session {
	return &auth.Session{
		ID:            s.ID.String(),
		UserID:        s.UserID.String(),
		TokenHash:     s.TokenHash,
		ExpiresAt:     s.ExpiresAt.Time,
		CreatedAt:     s.CreatedAt.Time,
		Email:         s.Email,
		UserCreatedAt: s.UserCreatedAt.Time,
		UserUpdatedAt: s.UserUpdatedAt.Time,
	}
}
