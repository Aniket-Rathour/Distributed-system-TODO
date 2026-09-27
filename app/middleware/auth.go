package middleware

import (
	"errors"
	"net/http"
	//"strings"
	"todo/app/db"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TokenCheck(pool *pgxpool.Pool , next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter , r *http.Request){

		//raw , ok := strings.CutPrefix(r.Header.Get("Token") , "Bearer ")
		raw := r.Header.Get("Token")
		ok := true
		//println(raw)
		if !ok  {
			http.Error(w, "ok error", http.StatusUnauthorized)
			return
		}else if raw == ""{
			http.Error(w, "raw was empty ", http.StatusUnauthorized)
			return
		}
		_ , err := db.UseridFromTocken(r.Context() , pool , raw)
		//println(id)
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		if err != nil {
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w,r)

	})
}