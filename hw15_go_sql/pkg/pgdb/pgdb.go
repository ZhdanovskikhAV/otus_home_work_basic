package pgdb

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(ctx context.Context, dbDSN string, maxOpenConns int32) (*pgxpool.Pool, error) {
	connConfig, err := pgxpool.ParseConfig(dbDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to create DSN for DB connection: %w", err)
	}
	connConfig.MaxConns = maxOpenConns
	connConfig.MinConns = 0
	dbc, err := pgxpool.NewWithConfig(ctx, connConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB : %w", err)
	}
	if err = dbc.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}

	return dbc, nil
}
