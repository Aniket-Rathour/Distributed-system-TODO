package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

func Insert(ctx context.Context, pool *pgxpool.Pool, name, pass string) (int, time.Time, error ,string) {
	insert := `INSERT INTO users (user_name, password) VALUES ($1 , $2) RETURNING id , created_at;`
	var id int
	var created time.Time
	err := pool.QueryRow(ctx, insert, name, pass).Scan(&id, &created)
	if err != nil {
		return 0, time.Time{}, err, ""
	}
	token ,err  := CreateSession(ctx, pool, id)

	return id, created, nil, token
}


func GetUsers(ctx context.Context, pool *pgxpool.Pool) ([]string, error) {
	getusers := `SELECT user_name FROM users`
	rows, err := pool.Query(ctx, getusers)
	if err != nil {
		return nil, err
	}
	var final []string
	defer rows.Close()
	for rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			return nil, err
		}
		final = append(final, username)
	}
	return final, nil
}



