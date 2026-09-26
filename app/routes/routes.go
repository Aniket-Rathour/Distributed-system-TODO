package routes

import (
	"net/http"
	"todo/app/middleware"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewHandler(pool *pgxpool.Pool, mux *http.ServeMux) *http.ServeMux {

	mux.HandleFunc("POST /users", Users(pool))
	mux.HandleFunc("GET /users", Users(pool))
	mux.Handle("POST /posts", middleware.TokenCheck(pool, Posts(pool)))
	mux.Handle("GET /posts/{Id}", middleware.TokenCheck(pool, Posts(pool)))
	mux.Handle("PUT /posts", middleware.TokenCheck(pool, Posts(pool)))
	return mux

}
