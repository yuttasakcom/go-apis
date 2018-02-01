package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/yuttasakcom/go-apis/events"
	"github.com/yuttasakcom/go-apis/routes"
)

func main() {

	go events.ClearLastRequestsIPs()
	go events.ClearBlockedIPs()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	r := routes.Router()

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	log.Println("go-apis running at port:" + port)
	log.Println("Press CTRL-C to stop")

	killSignal := <-stop
	switch killSignal {
	case os.Interrupt:
		log.Print("Got SIGINT...")
	case syscall.SIGTERM:
		log.Print("Got SIGTERM...")
	}

	log.Println("The service is shutting down...")
	srv.Shutdown(context.Background())
	log.Println("Server gracefully stopped")
}
