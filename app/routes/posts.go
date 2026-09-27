package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todo/app/db"

	"github.com/jackc/pgx/v5/pgxpool"
)
type postBody struct {
	Title  string
	Desc   string
	Userid int
}

func Posts(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var body postBody
			err := json.NewDecoder(r.Body).Decode(&body)
			if err != nil {
				fmt.Fprintln(w, err)
			}
			fmt.Println(body.Title, body.Desc)
			id, createdate, err := db.InsertPosts(r.Context(), pool, body.Title, body.Desc, body.Userid)
			fmt.Fprintln(w, "id: ", id, " created : ", createdate)
		case http.MethodGet:
			// type body struct{Id int}
			// var send body
			// json.NewDecoder(r.Body).Decode(&send)
			IdStr := r.PathValue("Id")
			id, err := strconv.Atoi(IdStr)
			q := r.URL.Query()
			pagestring := q.Get("page")
			///idstring := q.Get("id")
			page, err := strconv.Atoi(pagestring)
			//id , err := strconv.Atoi(idstring)

			result, err := db.Getposts(r.Context(), pool, id, 3, 3*(page-1))
			if err != nil {
				fmt.Fprintf(w, "there was. aerror reading. %s", err)
			}
			for n, name := range result {
				//fmt.Printf( "id = %d  = %s. = %s. = %d\n" , n , name.Title  , name.Description ,name.Id )
				fmt.Fprintf(w, "id = %d  = %s. = %s. = %d \n", n, name.Title, name.Description, name.Id)
			}
		case http.MethodPut:
			tables, err := db.GetPostWithAuthor(r.Context(), pool)
			if err != nil {
				fmt.Fprintln(w, err)
			}
			for n, names := range tables {
				fmt.Fprintln(w, "num = ", n, " title = ", names.Title, " description = ", names.Description, " user = ", names.User_name)
			}
		}

	}
}