package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todo/app/db"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)
type bodytype struct{
	Name string `json:"name"`
	Pass	string
}

func Users(pool *pgxpool.Pool) http.HandlerFunc{
	return func(w http.ResponseWriter , r *http.Request){
		switch r.Method{
		case http.MethodPost:
			var body bodytype
			err := json.NewDecoder(r.Body).Decode(&body)
			if err != nil {
				fmt.Fprintln(w, err)
			}
			if body.Name== "" || len(body.Pass) < 8 || len(body.Pass)> 72{
				http.Error(w, "username req, and pass much be 8 <char < 72", http.StatusBadRequest)
				return 
			}
			hash , err := bcrypt.GenerateFromPassword([]byte(body.Pass) ,bcrypt.DefaultCost )
			if err != nil {
				http.Error(w, "server error", http.StatusInternalServerError)
				return
			}
			id , created ,err := db.Insert(r.Context(), pool , body.Name, string(hash))
			fmt.Fprintln(w ,"this is id ", id , "this is time ", created , err)
		case http.MethodGet:
			final , err := db.GetUsers(r.Context() , pool )
			if err != nil{
				fmt.Fprintln(w, err )
			}
			for n , name  := range final {
				fmt.Fprintf(w , "id %d = %s\n", n,name)
			}
		case http.MethodHead:
			var b bodytype
			err := json.NewDecoder(r.Body).Decode(&b)
			if err!= nil{
				http.Error(w, "body decode error" , http.StatusBadRequest)
			}
			var id int
			var hash string 
			err = pool.QueryRow(r.Context() , "SELECT id, password FROM users WHERE user_name = $1" , b.Name).Scan(&id , &hash)
			if err!= nil{
				http.Error(w, "mthcing your pass from db" , http.StatusBadRequest)
			}
			err = bcrypt.CompareHashAndPassword([]byte(hash) ,[]byte(b.Pass))
			if err !=nil {
				http.Error(w ,"password doesnt match dude", http.StatusConflict)
			}


		}
	}
}

type postBody struct{
	Title string
	Desc	string
	Userid	int
}
func Posts(pool *pgxpool.Pool) http.HandlerFunc{
	return func(w http.ResponseWriter , r *http.Request){
		switch r.Method{
		case http.MethodPost:
			var body postBody
			err := json.NewDecoder(r.Body).Decode(&body)
			if err != nil {
				fmt.Fprintln(w, err)
			}
			fmt.Println(body.Title , body.Desc)
			id , createdate, err := db.InsertPosts(r.Context() , pool , body.Title , body.Desc ,body.Userid )
			fmt.Fprintln(w, "id: ", id , " created : " , createdate)
		case http.MethodGet:
			// type body struct{Id int}
			// var send body
			// json.NewDecoder(r.Body).Decode(&send)
			IdStr := r.PathValue("Id");
			id ,err  := strconv.Atoi(IdStr)
			q:= r.URL.Query()
			pagestring := q.Get("page")
			///idstring := q.Get("id")
			page , err := strconv.Atoi(pagestring)
			//id , err := strconv.Atoi(idstring)

			result ,err := db.Getposts(r.Context(), pool , id ,3,3*(page -1))
			if err != nil {
				fmt.Fprintf(w, "there was. aerror reading. %s", err)
			}
			for n , name := range result{
				//fmt.Printf( "id = %d  = %s. = %s. = %d\n" , n , name.Title  , name.Description ,name.Id )
				fmt.Fprintf(w , "id = %d  = %s. = %s. = %d \n" , n , name.Title  , name.Description ,name.Id )
			}
		case http.MethodPut:
			tables , err :=db.GetPostWithAuthor(r.Context(), pool)
			if err != nil {
				fmt.Fprintln(w, err)
			}
			for n , names  := range tables{
				fmt.Fprintln(w , "num = ", n , " title = ", names.Title ," description = ",names.Description, " user = ", names.User_name)
			}
		}
		
	}
}