package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Aniket-Rathour/Distibuted-system-TODO/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

func UserHandler(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter , r *http.Request){
		switch r.Method{
		case http.MethodGet:
			users , err := db.GetUser(r.Context() , pool)
			if err != nil{
				http.Error(w , err.Error() , http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(users)
		case http.MethodPost:
			var body struct{Username string ; Email string}
			json.NewDecoder(r.Body).Decode(&body)
			u, err := db.CreateUser(r.Context() , pool , body.Username , body.Email)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(u)
		default:
			http.Error(w , "method not allowed" , http.StatusMethodNotAllowed)
						
		}
	}
}
