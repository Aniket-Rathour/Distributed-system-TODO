package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo/app/db"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type bodytype struct {
	Name string `json:"name"`
	Pass string
}

func Users(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var body bodytype
			err := json.NewDecoder(r.Body).Decode(&body)
			if err != nil {
				fmt.Fprintln(w, err)
			}
			if body.Name == "" || len(body.Pass) < 8 || len(body.Pass) > 72 {
				http.Error(w, "username req, and pass much be 8 <char < 72", http.StatusBadRequest)
				return
			}
			hash, err := bcrypt.GenerateFromPassword([]byte(body.Pass), bcrypt.DefaultCost)
			if err != nil {
				http.Error(w, "server error", http.StatusInternalServerError)
				return
			}
			id, created, err ,token:= db.Insert(r.Context(), pool, body.Name, string(hash))
			fmt.Fprintln(w, "this is id ", id, "this is time ", created, err ,"\n your token = " , token)
		case http.MethodGet:
			final, err := db.GetUsers(r.Context(), pool)
			if err != nil {
				fmt.Fprintln(w, err)
			}
			for n, name := range final {
				fmt.Fprintf(w, "id %d = %s\n", n, name)
			}
		case http.MethodPut:
			var b bodytype
			err := json.NewDecoder(r.Body).Decode(&b)
			if err != nil {
				http.Error(w, "body decode error", http.StatusBadRequest)
			}
			var id int
			var hash string
			err = pool.QueryRow(r.Context(), "SELECT id, password FROM users WHERE user_name = $1", b.Name).Scan(&id, &hash)
			if err != nil {
				http.Error(w, "mthcing your pass from db", http.StatusBadRequest)
			}
			err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(b.Pass))
			if err != nil {
				http.Error(w, "password doesnt match dude", http.StatusConflict)
			}
			token , err := db.CreateSession(r.Context(),pool, id)
			if err!= nil {
				http.Error(w, "error in token creation ",0)
			}
			fmt.Fprintln(w ,"you have succufull loged in \n your token is =", token)
		}
	}
}



