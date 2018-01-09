package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gotoolkit/hook/pkg/handlers"
)

func main() {
	log.Print("Starting the service...")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}
	router := handlers.Router()
	log.Print("The service is ready to listen and serve.")

	http.ListenAndServe(":"+port, router)
}
