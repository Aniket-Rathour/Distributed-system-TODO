package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func insertPost(ctx context.Context , pool *pgxpool.Pool , id int ,title , content string) (int , error){
	insertSchema :=`
	INSERT INTO posts (user_id , title , content)
	VALUES ($!,$2,$3)
	RETURNING id;`
	var postId int
	err := pool.QueryRow(ctx,insertSchema , id , title , content).Scan(&postId)
	if err != nil {
		return 0 , err
	}

	return postId , nil
}

func deletePOst(ctx context.Context , pool *pgxpool.Pool){
	deletion schema := `
	DELETE FROM posts
	WHERE id = &1
	`
}