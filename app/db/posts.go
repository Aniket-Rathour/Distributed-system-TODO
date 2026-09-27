package db

import (
	"context"
	"time"
	"github.com/jackc/pgx/v5/pgxpool"
)
type posts struct {
	Description string
	Id          int
	Title       string
	Created_at  time.Time
}


type table struct {
	Title       string
	User_name   string
	Description string
}
func InsertPosts(ctx context.Context, pool *pgxpool.Pool, title, description string, userId int) (int, time.Time, error) {
	insert := `INSERT INTO posts (title, description , user_id ) VALUES ($1 , $2 , $3 ) RETURNING id , created_at;`
	var id int
	var created time.Time
	err := pool.QueryRow(ctx, insert, title, description, userId).Scan(&id, &created)
	if err != nil {
		return 0, time.Time{}, err
	}
	return id, created, nil
}


func Getposts(ctx context.Context, pool *pgxpool.Pool, id, start, end int) ([]posts, error) {
	getusers := `SELECT id, title, description, created_at 
				FROM posts 
				WHERE user_id = $1
				ORDER BY created_at DESC , id DESC
				LIMIT $2 OFFSET $3
				`
	rows, err := pool.Query(ctx, getusers, id, start, end)
	if err != nil {
		return nil, err
	}
	var final []posts
	defer rows.Close()
	for rows.Next() {
		var block posts
		err := rows.Scan(&block.Id, &block.Title, &block.Description, &block.Created_at)
		if err != nil {
			return nil, err
		}
		final = append(final, block)
	}
	return final, nil
}

func GetPostWithAuthor(ctx context.Context, pool *pgxpool.Pool) ([]table, error) {
	getposts := `
	SELECT posts.title,  posts.description , users.user_name 
	FROM posts 
	JOIN users ON posts.user_id = users.id`
	var u []table
	rows, err := pool.Query(ctx, getposts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var t table
		rows.Scan(&t.Title, &t.Description, &t.User_name)
		u = append(u, t)
	}
	return u, nil
}
