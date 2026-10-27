package car_transmission

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/Adisey/car-rental-api/internal/api_errors"
	"github.com/Adisey/car-rental-api/internal/api_models"
	shared_validation "github.com/Adisey/car-rental-api/internal/shared/validation"
)

func TransmissionTypesMainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTransmissionTypesHandler(w, r)

	case http.MethodPost:
		createTransmissionTypeHandler(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func TransmissionTypeByIDMainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTransmissionTypeByIDHandler(w, r)

	case http.MethodPatch:
		updateTransmissionTypeHandler(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getTransmissionTypesHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	transmissionTypes, err := GetTransmissionTypesService(r.Context())
	if err != nil {
		api_errors.HandleApiError(w, err)
		return
	}

	_ = json.NewEncoder(w).Encode(transmissionTypes)
}

func getTransmissionTypeByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/transmissionTypes/")
	bodyType, err := GetTransmissionTypeByIDService(r.Context(), id)
	if err != nil {
		api_errors.HandleApiError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(bodyType)
}

func createTransmissionTypeHandler(w http.ResponseWriter, r *http.Request) {
	var request api_models.CreateTransmissionTypeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid request",
		})
		return
	}
	bodyType, err := CreateTransmissionTypeService(
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

func updateTransmissionTypeHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/transmissionTypes/")
	var request api_models.UpdateTransmissionTypeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid request",
		})
		return
	}
	bodyType, err := UpdateTransmissionTypeService(
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
