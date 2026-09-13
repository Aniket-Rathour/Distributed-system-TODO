package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

const schema = `
CREATE TABLE IF NOT EXIST users(
	id	SERIAL PRIMARY KEY,
	name	VARCHAR(100) NOT NULL,
	email 	VARCHER(100) UNIQUE NOT NULL,
	created_at TIMESTAMP DEFAULT NOW()
);`

func Mirate(ctx context.Context , pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx ,schema )
	return err
}

