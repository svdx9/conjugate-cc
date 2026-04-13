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
	row, err := s.queries.CreateUser(ctx, email)
	if err != nil {
		return nil, s.handleDBError(err)
	}
	return toAuthUser(&row), nil
}

// FindUserByEmail finds a user by their email address
func (s *AuthStore) FindUserByEmail(ctx context.Context, email string) (*auth.User, error) {
	row, err := s.queries.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, s.handleDBError(err)
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
		return nil, s.handleDBError(err)
	}
	return toAuthUser(&row), nil
}

// CreateMagicLink creates a new magic link token for a user
func (s *AuthStore) CreateMagicLink(ctx context.Context, userID string, tokenHash []byte, expiresAt time.Time) (*auth.MagicLink, error) {
	uid, err := parseUUID(userID)
	if err != nil {
		return nil, auth.ErrUserNotFound
	}
	row, err := s.queries.CreateMagicLink(ctx, queries.CreateMagicLinkParams{
		UserID:    uid,
		TokenHash: tokenHash,
		ExpiresAt: timestamptzFromTime(expiresAt),
	})
	if err != nil {
		return nil, s.handleDBError(err)
	}
	return toAuthMagicLink(&row), nil
}

// FindMagicLinkByTokenHash finds an unconsumed, non-expired magic link by its token hash
func (s *AuthStore) FindMagicLinkByTokenHash(ctx context.Context, tokenHash []byte) (*auth.MagicLink, error) {
	row, err := s.queries.FindMagicLinkByTokenHash(ctx, tokenHash)
	if err != nil {
		return nil, s.handleDBError(err)
	}
	return toAuthMagicLinkRow(&row), nil
}

// ConsumeMagicLink marks a magic link as consumed
func (s *AuthStore) ConsumeMagicLink(ctx context.Context, magicLinkID string) error {
	mlid, err := parseUUID(magicLinkID)
	if err != nil {
		return auth.ErrMagicLinkNotFound
	}
	err = s.queries.ConsumeMagicLink(ctx, mlid)
	if err != nil {
		return s.handleDBError(err)
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
		return nil, s.handleDBError(err)
	}
	return toAuthSession(&row), nil
}

// FindSessionByTokenHash finds a non-expired session by its token hash
func (s *AuthStore) FindSessionByTokenHash(ctx context.Context, tokenHash []byte) (*auth.Session, error) {
	row, err := s.queries.FindSessionByTokenHash(ctx, tokenHash)
	if err != nil {
		return nil, s.handleDBError(err)
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
		return s.handleDBError(err)
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
		return s.handleDBError(err)
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
	ts.Scan(t)
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
		ID:        ml.ID.String(),
		UserID:    ml.UserID.String(),
		TokenHash: ml.TokenHash,
		ExpiresAt: ml.ExpiresAt.Time,
		CreatedAt: ml.CreatedAt.Time,
	}
}

// toAuthMagicLinkRow converts a sqlc FindMagicLinkByTokenHashRow to an auth.MagicLink
func toAuthMagicLinkRow(ml *queries.FindMagicLinkByTokenHashRow) *auth.MagicLink {
	return &auth.MagicLink{
		ID:        ml.ID.String(),
		UserID:    ml.UserID.String(),
		TokenHash: ml.TokenHash,
		ExpiresAt: ml.ExpiresAt.Time,
		CreatedAt: ml.CreatedAt.Time,
	}
}

// toAuthSession converts a sqlc Session to an auth.Session
func toAuthSession(s *queries.Session) *auth.Session {
	return &auth.Session{
		ID:        s.ID.String(),
		UserID:    s.UserID.String(),
		TokenHash: s.TokenHash,
		ExpiresAt: s.ExpiresAt.Time,
		CreatedAt: s.CreatedAt.Time,
	}
}

// toAuthSessionRow converts a sqlc FindSessionByTokenHashRow to an auth.Session
func toAuthSessionRow(s *queries.FindSessionByTokenHashRow) *auth.Session {
	return &auth.Session{
		ID:        s.ID.String(),
		UserID:    s.UserID.String(),
		TokenHash: s.TokenHash,
		ExpiresAt: s.ExpiresAt.Time,
		CreatedAt: s.CreatedAt.Time,
	}
}

// handleDBError maps database errors to domain errors
func (s *AuthStore) handleDBError(err error) error {
	if err == nil {
		return nil
	}

	// Check for "no rows" error
	if errors.Is(err, pgx.ErrNoRows) {
		// Different endpoints need different error messages, but we can't distinguish here
		// The caller should check the context to determine which error to return
		return auth.ErrUserNotFound // Default, may be overridden by caller
	}

	// Check for unique constraint violations
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		// unique_violation error code is 23505
		if pgErr.Code == "23505" {
			// Check which constraint was violated based on constraint name
			if pgErr.ConstraintName == "users_email_key" {
				return auth.ErrEmailTaken
			}
		}
	}

	// Unknown error
	return err
}
