package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/services"
	"github.com/Adisey/car-rental-api/internal/validation"
)

func CarsMainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getCarsHandler(w, r)

	case http.MethodPost:
		createCarHandler(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func CarByIDMainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getCarByIDHandler(w, r)

	case http.MethodPatch:
		updateCarHandler(w, r)

	case http.MethodDelete:
		deleteCarHandler(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getCarsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /cars")

	w.Header().Set("Content-Type", "application/json")

	cars, err := services.GetCarsService(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(cars)
}

func getCarByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/cars/")
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

func createCarHandler(w http.ResponseWriter, r *http.Request) {
	var request api_models.CreateCarRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid request",
		})
		return
	}
	car, err := services.CreateCarService(
		r.Context(),
		request,
	)
	if err != nil {
		var validationResult *validation.ValidationResult
		if errors.As(err, &validationResult) {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(
				validationResult.Response(),
			)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(car)
}

func updateCarHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/cars/")
	var request api_models.UpdateCarRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid request",
		})
		return
	}
	car, err := services.UpdateCarService(
		r.Context(),
		id,
		request,
	)
	if err != nil {
		var validationResult *validation.ValidationResult
		if errors.As(err, &validationResult) {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(
				validationResult.Response(),
			)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	w.Header().Set(
		"Content-Type",
		"application/json",
	)
	_ = json.NewEncoder(w).Encode(car)
}

func deleteCarHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/cars/")

	if !services.DeleteCarService(id) {
		w.WriteHeader(http.StatusNotFound)

		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": "car not found",
		})

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
