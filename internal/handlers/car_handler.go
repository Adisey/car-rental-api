package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/services"
)

func Cars(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getCars(w, r)

	case http.MethodPost:
		createCar(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func CarByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getCarByIDHandler(w, r)

	case http.MethodPut:
		updateCar(w, r)

	case http.MethodDelete:
		deleteCar(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getCars(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /cars")

	w.Header().Set("Content-Type", "application/json")

	cars, err := services.GetCars(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(cars)
}

func getCarByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/cars/")

	log.Println("Handler GET /cars/", id)

	car, err := services.GetCarByIDService(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": "car not found",
		})

		return
	}

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(car)
}


func createCar(w http.ResponseWriter, r *http.Request) {
	var request api_models.CreateCarRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid request",
		})

		return
	}

	car := services.CreateCar(request)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(car)
}

func updateCar(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/cars/")

	var request api_models.UpdateCarRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid request",
		})

		return
	}

	car, found := services.UpdateCar(id, request)

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

func deleteCar(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/cars/")

	if !services.DeleteCar(id) {
		w.WriteHeader(http.StatusNotFound)

		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": "car not found",
		})

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
