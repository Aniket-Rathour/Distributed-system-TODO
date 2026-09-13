package db

import (
	"context"
	"os"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context) (*pgxpool.Pool , error){
	connstr := os.Getenv("DATABASE_URL")
	pool , err := pgxpool.New(ctx , connstr)
	if err!= nil {
		return nil , err
	}
	return pool , nil
}