package routes

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewHandler(pool *pgxpool.Pool ,mux *http.ServeMux) *http.ServeMux{

	mux.HandleFunc("POST /users",Users(pool))
	mux.HandleFunc("GET /users",Users(pool))
	mux.HandleFunc("POST /posts",Posts(pool))
	mux.HandleFunc("GET /posts/{Id}",Posts(pool))
	mux.HandleFunc("PUT /posts",Posts(pool))

	return mux

	
}