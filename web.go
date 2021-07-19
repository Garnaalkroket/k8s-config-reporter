package main

import (
	"encoding/json"
	"net/http"
)

func resourceHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/resources" || r.Method != "GET" {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resources())
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Test</h1>"))
}
