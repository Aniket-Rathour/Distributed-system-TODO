package db

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewToken() (string, []byte , error){
	b:= make([]byte, 32)
	if _,err := rand.Read(b) ; err !=nil {
		return "", nil , err
	}
	raw := base64.RawURLEncoding.EncodeToString(b)
	sum := sha256.Sum256([]byte(raw))
	return raw , sum[:] , nil
}

func CreateSession(ctx context.Context,pool *pgxpool.Pool ,userid int ) (string, error){
	raw, hashed, err := NewToken()
	if err != nil {
        return "", err
    }
	_ , err = pool.Exec(ctx, 
		`INSERT INTO sessions (token_hash ,user_id , expire_at )
		VALUES ($1, $2 , now()+interval '7 days');`, hashed, userid)
	return raw, err
}

func UseridFromTocken(ctx context.Context , pool *pgxpool.Pool , raw string) (int ,error){
	sum := sha256.Sum256([]byte(raw))
	var userID int
	err := pool.QueryRow(ctx , 
	`SELECT user_id FROM sessions WHERE token_hash = $1 AND expires_at > now()` , sum[:]).Scan(&userID)
	return userID , err
}