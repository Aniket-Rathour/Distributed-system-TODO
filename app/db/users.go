package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct{
	ID int
	Username	string
	Email	string
} 

func GetUser(ctx context.Context , pool *pgxpool.Pool) ([]User , error){
	rows , err := pool.Query(ctx , "SELECT id, username , email FROM users")
	if err != nil {
		return nil , err
	}
	defer rows.Close()
	var users []User
	for rows.Next(){
		var u  User
		err = rows.Scan(&u.ID , &u.Username , &u.Email )
		if err!= nil {
			return nil ,err
		}
		users = append(users, u)
	}
	return users,  nil
}

func CreateUser(ctx context.Context , pool *pgxpool.Pool, username , email string ) (User,error){
	var u User
	err := pool.QueryRow(ctx,
			"INSERT INTO users (username, email) VALUES ($1 , $2) RETURNING id, username , email",
			username ,email).Scan(&u.ID ,&u.Username , &u.Email )	
	return u , err
}

