package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context) (*pgxpool.Pool , error){
	connstr := os.Getenv("DATABASE_URL")
	pool , err := pgxpool.New(ctx , connstr)
	if err!= nil {
		return nil , err
	}
	fmt.Println("succufully connected to db....")
	return pool , nil
}