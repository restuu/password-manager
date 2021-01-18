package database

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// PostgresOpen ...
func PostgresOpen(ctx context.Context, databaseURL string) (*pgx.Conn, error) {
	return pgx.Connect(ctx, databaseURL)
}
