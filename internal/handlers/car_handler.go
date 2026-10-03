package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/Adisey/car-rental-api/internal/services"
)

// GetCars returns all cars
//
// @Summary Get cars
// @Tags cars
// @Produce json
// @Success 200 {array} models.Car
// @Router /cars [get]
func Cars(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /cars")

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(services.GetCars())
}

// GetCarByID returns a car by id
//
// @Summary Get car
// @Tags cars
// @Produce json
// @Param id path string true "Car ID"
// @Success 200 {object} models.Car
// @Failure 404 {object} map[string]string
// @Router /cars/{id} [get]
func CarByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/cars/")

	log.Println("GET /cars/", id)

	car, found := services.GetCarByID(id)

	if !found {
		w.WriteHeader(http.StatusNotFound)

		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": "car not found",
		})

		return
	}

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(car)
}
