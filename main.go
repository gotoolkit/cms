package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gotoolkit/hook/pkg/handlers"
	"github.com/gotoolkit/hook/pkg/version"
)

func main() {
	log.Printf("Starting the service...\ncommit: %s, build time: %s, release: %s", version.Commit, version.BuildTime, version.Release)
	log.Print("Starting the service...")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}
	router := handlers.Router(version.BuildTime, version.Commit, version.Release)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	srv := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	go func() {
		srv.ListenAndServe()
	}()
	log.Print("The service is ready to listen and serve.")

	killSignal := <-interrupt

	switch killSignal {
	case os.Interrupt:
		log.Println("Got SIGINT...")
	case syscall.SIGTERM:
		log.Println("Got SIGTERM...")
	}
	log.Print("The service is shutting down...")
	srv.Shutdown(context.Background())
	log.Println("Done")
}
