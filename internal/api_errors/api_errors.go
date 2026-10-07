package api_errors

import (
	"encoding/json"
	stdErrors "errors"
	"log"
	"net/http"

	"github.com/Adisey/car-rental-api/internal/validation"
)

func HandleApiError(
	w http.ResponseWriter,
	err error,
) {

	log.Printf("HandleApiError: %v", err)

	var validationResult *validation.ValidationResult

	if stdErrors.As(err, &validationResult) {

		w.Header().Set(
			"Content-Type",
			"application/json",
		)

		w.WriteHeader(http.StatusBadRequest)

		_ = json.NewEncoder(w).Encode(
			validationResult.Response(),
		)

		return
	}

	if stdErrors.Is(err, ErrNotFound) {

		w.Header().Set(
			"Content-Type",
			"application/json",
		)

		w.WriteHeader(http.StatusNotFound)

		_ = json.NewEncoder(w).Encode(
			map[string]string{
				"error": "not found",
			},
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(http.StatusInternalServerError)

	_ = json.NewEncoder(w).Encode(
		map[string]string{
			"error": "internal error",
		},
	)
}
