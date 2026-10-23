package car_drive_types

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/Adisey/car-rental-api/internal/api_errors"
	"github.com/Adisey/car-rental-api/internal/api_models"
	shared_validation "github.com/Adisey/car-rental-api/internal/shared/validation"
)

func DriveTypesMainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getDriveTypesHandler(w, r)

	case http.MethodPost:
		createDriveTypeHandler(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func DriveTypeByIDMainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getDriveTypeByIDHandler(w, r)

	case http.MethodPatch:
		updateDriveTypeHandler(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getDriveTypesHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	bodyTypes, err := GetDriveTypesService(r.Context())
	if err != nil {
		api_errors.HandleApiError(w, err)
		return
	}

	_ = json.NewEncoder(w).Encode(bodyTypes)
}

func getDriveTypeByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/bodyTypes/")
	bodyType, err := GetDriveTypeByIDService(r.Context(), id)
	if err != nil {
		api_errors.HandleApiError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(bodyType)
}

func createDriveTypeHandler(w http.ResponseWriter, r *http.Request) {
	var request api_models.CreateDriveTypeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid request",
		})
		return
	}
	bodyType, err := CreateDriveTypeService(
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

func updateDriveTypeHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/bodyTypes/")
	var request api_models.UpdateDriveTypeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid request",
		})
		return
	}
	bodyType, err := UpdateDriveTypeService(
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
