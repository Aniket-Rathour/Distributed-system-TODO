package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Aniket-Rathour/Distibuted-system-TODO/db"
	"github.com/Aniket-Rathour/Distibuted-system-TODO/routes"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("thre was a error reading env")
	}
	//ctx1 := make(chan os.Signal ,1 )
	ctx , stop := signal.NotifyContext(context.Background() , os.Interrupt, syscall.SIGTERM)
	defer stop()
	pool , err :=db.Connect(context.Background())
	if err != nil {
		fmt.Println("thre was a connecting to db" ,err)
		os.Exit(1)
	}
	defer pool.Close()
	if err= db.Migrate(context.Background(), pool); err!= nil{
		fmt.Println("failed to migrate schema:", err)
		os.Exit(1)
	}
	
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux ,pool)
	server := http.Server{
		Addr: os.Getenv("PORT"),
		Handler: mux,
	}
	
	go func(){
		server.ListenAndServe()
	}()

	<-ctx.Done()
	fmt.Println("the server is closing")
	timectx ,close := context.WithTimeout(context.Background() ,10*time.Second)
	defer close()
	server.Shutdown(timectx)
	
}