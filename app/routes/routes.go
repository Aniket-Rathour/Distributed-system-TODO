package routes

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRoutes(mux *http.ServeMux ,pool *pgxpool.Pool){
	mux.HandleFunc("/users" ,UserHandler(pool))

}