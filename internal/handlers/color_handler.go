package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/Adisey/car-rental-api/internal/api_errors"
	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/services"
	"github.com/Adisey/car-rental-api/internal/validation"
)

func ColorsMainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getColorsHandler(w, r)

	case http.MethodPost:
		createColorHandler(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func ColorByIDMainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getColorByIDHandler(w, r)

	case http.MethodPatch:
		updateColorHandler(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getColorsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	colors, err := services.GetColorsService(r.Context())
	if err != nil {
		api_errors.HandleApiError(w, err)
		return
	}

	_ = json.NewEncoder(w).Encode(colors)
}

func getColorByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/colors/")
	color, err := services.GetColorByIDService(r.Context(), id)
	if err != nil {
		api_errors.HandleApiError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(color)
}

func createColorHandler(w http.ResponseWriter, r *http.Request) {
	var request api_models.CreateColorRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid request",
		})
		return
	}
	color, err := services.CreateColorService(
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
	_ = json.NewEncoder(w).Encode(color)
}

func updateColorHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/colors/")
	var request api_models.UpdateColorRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid request",
		})
		return
	}
	color, err := services.UpdateColorService(
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
	_ = json.NewEncoder(w).Encode(color)
}
