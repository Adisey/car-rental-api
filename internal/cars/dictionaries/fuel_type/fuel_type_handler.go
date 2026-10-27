package car_fuel

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/Adisey/car-rental-api/internal/api_errors"
	"github.com/Adisey/car-rental-api/internal/api_models"
	shared_validation "github.com/Adisey/car-rental-api/internal/shared/validation"
)

func FuelTypesMainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getFuelTypesHandler(w, r)

	case http.MethodPost:
		createFuelTypeHandler(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func FuelTypeByIDMainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getFuelTypeByIDHandler(w, r)

	case http.MethodPatch:
		updateFuelTypeHandler(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getFuelTypesHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	furlTypes, err := GetFuelTypesService(r.Context())
	if err != nil {
		api_errors.HandleApiError(w, err)
		return
	}

	_ = json.NewEncoder(w).Encode(furlTypes)
}

func getFuelTypeByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/furlTypes/")
	bodyType, err := GetFuelTypeByIDService(r.Context(), id)
	if err != nil {
		api_errors.HandleApiError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(bodyType)
}

func createFuelTypeHandler(w http.ResponseWriter, r *http.Request) {
	var request api_models.CreateFuelTypeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid request",
		})
		return
	}
	bodyType, err := CreateFuelTypeService(
		r.Context(),
		request,
	)
	if err != nil {
		var validationResult *shared_validation.ValidationResult
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
	_ = json.NewEncoder(w).Encode(bodyType)
}

func updateFuelTypeHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/fuelTypes/")
	var request api_models.UpdateFuelTypeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid request",
		})
		return
	}
	bodyType, err := UpdateFuelTypeService(
		r.Context(),
		id,
		request,
	)
	if err != nil {
		var validationResult *shared_validation.ValidationResult
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
	_ = json.NewEncoder(w).Encode(bodyType)
}
