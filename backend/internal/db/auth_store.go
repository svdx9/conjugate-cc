package db

import (
	"context"
	"errors"
	"fmt"
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
	pgErrCodeForeignKeyViolation = "23503" // foreign key constraint violation
)

// Clock provides the current time, allowing time-dependent operations to be testable
// store is a wrapper around the database queries with transaction support
type store struct {
	queries *queries.Queries
	pool    *pgxpool.Pool
	logger  *slog.Logger
}

func (s *store) withTx(ctx context.Context, fn func(qtx *queries.Queries) error) error {
	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		rbErr := tx.Rollback(ctx)
		if rbErr != nil && !errors.Is(rbErr, pgx.ErrTxClosed) {
			s.logger.Error("rollback failed", "error", errors.Join(err, rbErr))
		}
	}()

	err = fn(s.queries.WithTx(tx))
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// AuthStore handles all authentication-related database operations
type AuthStore struct {
	*store
}

// NewAuthStore creates a new AuthStore
func NewAuthStore(pool *pgxpool.Pool, logger *slog.Logger) *AuthStore {
	return &AuthStore{
		store: &store{
			queries: queries.New(pool),
			pool:    pool,
			logger:  logger,
		},
	}
}

// CreateUser creates a new user with the given email
func (s *AuthStore) CreateUser(ctx context.Context, email string) (auth.User, error) {
	// CreateUser is an UPSERT operation
	row, err := s.queries.CreateUser(ctx, email)
	if err != nil {
		s.logger.Error("failed to create user", "error", err)
		return auth.User{}, auth.ErrInternal
	}
	return toAuthUser(&row), nil
}

// CreateOrUpdateMagicLinkToken creates or updates a magic link token for a user
// This handles race conditions atomically at the database level using UPSERT:
// - If no unconsumed magic link exists for this user, creates a new one
// - If an unconsumed magic link exists, updates its token and expiration
// This prevents multiple concurrent requests from creating conflicting tokens
func (s *AuthStore) CreateOrUpdateMagicLinkToken(ctx context.Context, userID string, tokenHash auth.TokenHash, expiresAt time.Time) (auth.MagicLink, error) {
	uid, err := parseUUID(userID)
	if err != nil {
		return auth.MagicLink{}, auth.ErrUserNotFound
	}
	expiresAtTS, err := timestamptzFromTime(expiresAt)
	if err != nil {
		s.logger.Error("failed to convert expiresAt to timestamp", "error", err)
		return auth.MagicLink{}, auth.ErrInternal
	}
	row, err := s.queries.CreateOrUpdateMagicLinkToken(ctx, queries.CreateOrUpdateMagicLinkTokenParams{
		UserID:    uid,
		TokenHash: tokenHash,
		ExpiresAt: expiresAtTS,
	})
	if err != nil {
		// Check for foreign key constraint violation (user doesn't exist)
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgErrCodeForeignKeyViolation {
			return auth.MagicLink{}, auth.ErrUserNotFound
		}
		s.logger.Error("failed to create or update magic link token", "user_id", userID, "error", err)
		return auth.MagicLink{}, auth.ErrInternal
	}
	return toAuthMagicLink(&row), nil
}

// ConsumeMagicLinkAndCreateSession
func (s *AuthStore) ConsumeMagicLinkAndCreateSession(ctx context.Context, tokenHash auth.TokenHash, sessionTokenHash auth.TokenHash, now time.Time, sessionExpiresAt time.Time) (auth.User, auth.Session, error) {
	// use the tokenHash to find the magic link
	// if the magic link is not found, return an error
	// if the magic link is found, check if it is expired
	// if the magic link is not expired, consume the magic link (mark as used)
	// create a session token for the user using the given sessionTokenHash and userID from the magic link
	var user auth.User
	var session auth.Session
	err := s.withTx(ctx, func(qtx *queries.Queries) error {
		// Find the magic link and user in a single query
		ml, err := qtx.FindMagicLinkByTokenHash(ctx, tokenHash)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return auth.ErrMagicLinkNotFound
			}
			s.logger.Error("failed to find magic link by token hash", "error", err)
			return auth.ErrInternal
		}

		if !ml.ExpiresAt.Valid {
			return auth.ErrMagicLinkExpired
		}

		// Check expiration
		if now.After(ml.ExpiresAt.Time) {
			return auth.ErrMagicLinkExpired
		}

		// Consume the magic link (mark as used)
		_, err = qtx.ConsumeMagicLink(ctx, ml.ID)
		if err != nil {
			// Check if this is a "no rows" scenario (magic link already consumed)
			if errors.Is(err, pgx.ErrNoRows) {
				return auth.ErrMagicLinkConsumed
			}
			s.logger.Error("failed to consume magic link", "magic_link_id", ml.ID, "error", err)
			return auth.ErrInternal
		}

		// Construct user from magic link query (which joins users table)
		user = auth.User{
			ID:        ml.UserID.String(),
			Email:     ml.Email,
			CreatedAt: ml.UserCreatedAt.Time,
			UpdatedAt: ml.UserUpdatedAt.Time,
		}
		// Store session in database
		expiresAtTS, err := timestamptzFromTime(sessionExpiresAt)
		if err != nil {
			s.logger.Error("failed to convert expiresAt to timestamp", "error", err)
			return auth.ErrInternal
		}
		row, err := qtx.CreateSession(ctx, queries.CreateSessionParams{
			UserID:    ml.UserID,
			TokenHash: sessionTokenHash,
			ExpiresAt: expiresAtTS,
		})
		if err != nil {
			// Check for foreign key constraint violation (user doesn't exist)
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgErr.Code == pgErrCodeForeignKeyViolation {
				return auth.ErrUserNotFound
			}
			s.logger.Error("failed to create session", "user_id", ml.UserID, "error", err)
			return auth.ErrInternal
		}
		session = toAuthSession(&row)
		// Return user and session
		return nil
	})
	if err != nil {
		return auth.User{}, auth.Session{}, err
	}
	return user, session, err

}

// FindMagicLinkByTokenHash finds an unconsumed, non-expired magic link by its token hash
func (s *AuthStore) FindMagicLinkByTokenHash(ctx context.Context, tokenHash auth.TokenHash) (auth.MagicLink, error) {
	row, err := s.queries.FindMagicLinkByTokenHash(ctx, tokenHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return auth.MagicLink{}, auth.ErrMagicLinkNotFound
		}
		s.logger.Error("failed to find magic link by token hash", "error", err)
		return auth.MagicLink{}, auth.ErrInternal
	}
	return toAuthMagicLinkRow(&row), nil
}

// FindSessionByTokenHash finds a non-expired session by its token hash
func (s *AuthStore) FindSessionByTokenHash(ctx context.Context, tokenHash auth.TokenHash) (auth.Session, error) {
	row, err := s.queries.FindSessionByTokenHash(ctx, tokenHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return auth.Session{}, auth.ErrSessionNotFound
		}
		s.logger.Error("failed to find session by token hash", "error", err)
		return auth.Session{}, auth.ErrInternal
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
func timestamptzFromTime(t time.Time) (pgtype.Timestamptz, error) {
	var ts pgtype.Timestamptz
	err := ts.Scan(t)
	if err != nil {
		return pgtype.Timestamptz{}, fmt.Errorf("failed to scan time: %w", err)
	}
	return ts, nil
}

// toAuthUser converts a sqlc User to an auth.User
func toAuthUser(u *queries.User) auth.User {
	return auth.User{
		ID:        u.ID.String(),
		Email:     u.Email,
		CreatedAt: u.CreatedAt.Time,
		UpdatedAt: u.UpdatedAt.Time,
	}
}

// toAuthMagicLink converts a sqlc MagicLink to an auth.MagicLink
func toAuthMagicLink(ml *queries.MagicLink) auth.MagicLink {
	return auth.MagicLink{
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
func toAuthMagicLinkRow(ml *queries.FindMagicLinkByTokenHashRow) auth.MagicLink {
	return auth.MagicLink{
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
func toAuthSession(s *queries.Session) auth.Session {
	return auth.Session{
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
func toAuthSessionRow(s *queries.FindSessionByTokenHashRow) auth.Session {
	return auth.Session{
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
