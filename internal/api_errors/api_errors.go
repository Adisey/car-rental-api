package api_errors

import (
	"encoding/json"
	stdErrors "errors"
	"net/http"
	"strings"

	shared_validation "github.com/Adisey/car-rental-api/internal/shared/validation"
)

func HandleApiError(
	w http.ResponseWriter,
	err error,
) {

	// log.Printf("HandleApiError: %v", err)

	var validationResult *shared_validation.ValidationResult

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

	if strings.Contains(err.Error(), "invalid UUID") {
		w.WriteHeader(http.StatusBadRequest)
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
