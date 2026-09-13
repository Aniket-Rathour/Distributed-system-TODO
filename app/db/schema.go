package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

const schema = `
CREATE TABLE IF NOT EXISTS users(
	id	SERIAL PRIMARY KEY,
	username	VARCHAR(100) NOT NULL,
	email 	VARCHAR(100) UNIQUE NOT NULL,
	created_at TIMESTAMP DEFAULT NOW()
);`

func Migrate(ctx context.Context , pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx ,schema )
	fmt.Println("succefully created the dtable...")
	return err
}

