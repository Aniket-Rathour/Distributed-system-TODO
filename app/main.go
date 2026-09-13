package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load("../.env")
	//ctx1 := make(chan os.Signal ,1 )
	ctx , stop := signal.NotifyContext(context.Background() , os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err != nil {
		fmt.Println("thre was a error reading env")
	}
	mux := http.NewServeMux()
	server := http.Server{
		Addr: os.Getenv("PORT"),
		Handler: mux,
	}
	
	go func(){
		server.ListenAndServe()
	}()

	<-ctx.Done()
	fmt.Println("the server is closing")
	timectx ,close := context.WithTimeout(context.Background() ,10*time.Second , syscall.SIGTERM)
	defer close()
	server.Shutdown(timectx)
	
}