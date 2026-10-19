package car_model

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/Adisey/car-rental-api/internal/api_errors"
	"github.com/Adisey/car-rental-api/internal/api_models"
	shared_validation "github.com/Adisey/car-rental-api/internal/shared/validation"
)

func ModelsMainHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	switch r.Method {
	case http.MethodGet:
		getModelsHandler(w, r)

	case http.MethodPost:
		createModelHandler(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func ModelByIDMainHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	switch r.Method {
	case http.MethodGet:
		getModelByIDHandler(w, r)

	case http.MethodPatch:
		updateModelHandler(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getModelsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	log.Println("GET /models")

	models, err := GetModelsService(
		r.Context(),
		nil,
	)

	if err != nil {
		api_errors.HandleApiError(
			w,
			err,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	_ = json.NewEncoder(w).Encode(
		models,
	)
}

func getModelByIDHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	id := strings.TrimPrefix(
		r.URL.Path,
		"/models/",
	)

	log.Println(
		"GET /models/",
		id,
	)

	model, err := GetModelByIDService(
		r.Context(),
		id,
	)

	if err != nil {
		api_errors.HandleApiError(
			w,
			err,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	_ = json.NewEncoder(w).Encode(
		model,
	)
}

func createModelHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	var request api_models.CreateModelRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {

		w.WriteHeader(
			http.StatusBadRequest,
		)

		_ = json.NewEncoder(w).Encode(
			map[string]string{
				"error": "invalid request",
			},
		)

		return
	}

	model, err := CreateModelService(
		r.Context(),
		request,
	)

	if err != nil {

		var validationResult *shared_validation.ValidationResult

		if errors.As(
			err,
			&validationResult,
		) {

			w.WriteHeader(
				http.StatusBadRequest,
			)

			_ = json.NewEncoder(w).Encode(
				validationResult.Response(),
			)

			return
		}

		api_errors.HandleApiError(
			w,
			err,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(
		http.StatusCreated,
	)

	_ = json.NewEncoder(w).Encode(
		model,
	)
}

func updateModelHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	id := strings.TrimPrefix(
		r.URL.Path,
		"/models/",
	)

	var request api_models.UpdateModelRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {

		w.WriteHeader(
			http.StatusBadRequest,
		)

		_ = json.NewEncoder(w).Encode(
			map[string]string{
				"error": "invalid request",
			},
		)

		return
	}

	model, err := UpdateModelService(
		r.Context(),
		id,
		request,
	)

	if err != nil {

		var validationResult *shared_validation.ValidationResult

		if errors.As(
			err,
			&validationResult,
		) {

			w.WriteHeader(
				http.StatusBadRequest,
			)

			_ = json.NewEncoder(w).Encode(
				validationResult.Response(),
			)

			return
		}

		api_errors.HandleApiError(
			w,
			err,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	_ = json.NewEncoder(w).Encode(
		model,
	)
}
