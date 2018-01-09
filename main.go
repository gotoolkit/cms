package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Print("Starting the service...")

	http.HandleFunc("/home", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello! Your request was processed.")
	})
	log.Print("The service is ready to listen and serve.")

	http.ListenAndServe(":8000", nil)
}
