package db

import (
	"context"
	"errors"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// WithTx runs a function within a database transaction.
// The transaction is automatically committed if the function returns nil,
// or rolled back if it returns an error.
// Panics within fn are safely handled by ensuring rollback via defer.
func WithTx(ctx context.Context, pool *pgxpool.Pool, logger *slog.Logger, fn func(pgx.Tx) error) error {
	tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		rbErr := tx.Rollback(ctx)
		if rbErr != nil {
			logger.Error("rollback failed", "error", errors.Join(err, rbErr))
		}
	}()

	err = fn(tx)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
