package main

import (
	"context"
	"fmt"

	"net/http"
	"os"
	"os/signal"
	"time"
	"todo/app/db"
	"todo/app/routes"

	"github.com/joho/godotenv"
)

func main() {
	servermux := http.NewServeMux()
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Printf("Error loading .env file: %v", err)
	}
	dbstring := os.Getenv("DB_STRING")
	servermux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "this is home dude ")
	})
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	pool, err := db.ConnectDb(ctx, dbstring)
	if err != nil {
		fmt.Println("failed to connet to db ")
	} else {
		fmt.Println("connected to db....")
	}

	err = db.CreateUserTable(ctx, pool)
	if err != nil {
		fmt.Println("there was a error in creadint ht etable ")
	} else {
		fmt.Println("the datels were created....")
	}

	svr := http.Server{
		Handler:     routes.NewHandler(pool, servermux),
		Addr:        ":8080",
		IdleTimeout: 10 * time.Second,
	}
	go func() {
		svr.ListenAndServe()
	}()
	<-ctx.Done()
	fmt.Println("closing the server/...")
	timectx, close := context.WithTimeout(context.Background(), 10*time.Second)
	defer close()
	if err := svr.Shutdown(timectx); err != nil {
		fmt.Println("timeout , or failed to close the server....")
	}
	fmt.Println("server stopped cleanly")
}
