package main

import (
	"log"
	"net/http"

	"github.com/gotoolkit/hook/pkg/handlers"
)

func main() {
	log.Print("Starting the service...")
	router := handlers.Router()
	log.Print("The service is ready to listen and serve.")

	http.ListenAndServe(":8000", router)
}
