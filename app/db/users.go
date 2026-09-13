package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct{
	ID int
	username	string
	email	string
} 

func GetUser(ctx context.Context , pool *pgxpool.Pool) ([]User , error){
	rows , err := pool.Query(ctx , "SELECT id, username FROM users")
	if err != nil {
		return nil , err
	}
}