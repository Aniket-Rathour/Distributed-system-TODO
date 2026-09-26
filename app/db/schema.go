package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateUserTable(ctx context.Context, pool *pgxpool.Pool) error {
	tableQueary := `
	CREATE TABLE IF NOT EXISTS users(
	id SERIAL PRIMARY KEY,
	user_name VARCHAR(50) UNIQUE NOT NULL,
	password VARCHAR(100) NOT NULL,
	created_at TIMESTAMP DEFAULT now()
	);`
	tokenQueary := ` 
	CREATE TABLE IF NOT EXISTS sessions(
	token_hash BYTEA PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	created_at TIMESTAMPTZ DEFAULT now(),
	expire_at	TIMESTAMPTZ NOT NULL
	);`

	postQueary := `
	CREATE TABLE IF NOT EXISTS posts(
		id SERIAL PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		description TEXT ,
		user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		created_at TIMESTAMP DEFAULT now()
	);`

	_, err := pool.Exec(ctx, tableQueary)
	if err != nil {
		return err
	}
	_, err = pool.Exec(ctx, postQueary)
	if err != nil {
		return err
	}
	_, err = pool.Exec(ctx, tokenQueary)
	if err != nil {
		return err
	}
	return nil
}
