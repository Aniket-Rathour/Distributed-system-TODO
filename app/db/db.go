package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDb(ctx context.Context, str string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, str)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
