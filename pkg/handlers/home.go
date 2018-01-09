package handlers

import "net/http"
import "fmt"

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! Your request was processed.")
}
