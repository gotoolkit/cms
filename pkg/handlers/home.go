package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gotoolkit/hook/pkg/version"
)

func home(w http.ResponseWriter, r *http.Request) {
	info := struct {
		BuildTime string `json:"build_time"`
		Release   string `json:"release"`
		Commit    string `json:"commit"`
	}{
		version.BuildTime,
		version.Release,
		version.Commit,
	}
	body, err := json.Marshal(info)
	if err != nil {
		log.Printf("Could not encode info data: %v", err)
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
