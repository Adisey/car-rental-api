package customers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/Adisey/car-rental-api/internal/api_errors"
	"github.com/Adisey/car-rental-api/internal/api_models"
	shared_validation "github.com/Adisey/car-rental-api/internal/shared/validation"
)

func CustomersMainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getCustomersHandler(w, r)
	case http.MethodPost:
		createCustomerHandler(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func CustomerByIDMainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getCustomerByIDHandler(w, r)
	case http.MethodPatch:
		updateCustomerHandler(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getCustomersHandler(w http.ResponseWriter, r *http.Request) {
	customers, err := GetCustomersService(r.Context())
	if err != nil {
		api_errors.HandleApiError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(customers)
}

func getCustomerByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/customers/")
	customer, err := GetCustomerByIDService(r.Context(), id)
	if err != nil {
		api_errors.HandleApiError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(customer)
}

func createCustomerHandler(w http.ResponseWriter, r *http.Request) {
	var request api_models.CreateCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			map[string]string{
				"error": "invalid request",
			},
		)
		return
	}
	customer, err := CreateCustomerService(r.Context(), request)
	if err != nil {
		var validationResult *shared_validation.ValidationResult
		if errors.As(err, &validationResult) {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(validationResult.Response())
			return
		}
		api_errors.HandleApiError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(customer)
}

func updateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/customers/")
	var request api_models.UpdateCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			map[string]string{
				"error": "invalid request",
			},
		)
		return
	}
	customer, err := UpdateCustomerService(r.Context(), id, request)
	if err != nil {
		var validationResult *shared_validation.ValidationResult
		if errors.As(err, &validationResult) {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(validationResult.Response())
			return
		}
		api_errors.HandleApiError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(customer)
}
