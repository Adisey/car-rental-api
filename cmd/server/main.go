package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
		"version": "1.0.6",
	})
}

func main() {
	http.HandleFunc("/health", healthHandler)

	log.Println("Server started on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}