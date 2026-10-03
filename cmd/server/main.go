package main

// @title Car Rental API
// @version 1.0
// @description Car Rental backend service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Adisey/car-rental-api/internal/handlers"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"version": "1.0.6",
	})
}

func main() {
	http.HandleFunc("/health", healthHandler)

	http.HandleFunc("/cars", handlers.Cars)
	http.HandleFunc("/cars/", handlers.CarByID)

	log.Println("Server started on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
