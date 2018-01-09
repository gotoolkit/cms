package main

import (
	"log"
	"net/http"
	"os"

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
	log.Print("The service is ready to listen and serve.")

	http.ListenAndServe(":"+port, router)
}
