package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Adisey/car-rental-api/internal/db"
)

func Health(w http.ResponseWriter, r *http.Request) {
	if err := db.CheckHealth(r.Context()); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)

		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":   "error",
			"version":  "1.0.6",
			"database": "failed",
		})

		return
	}

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(map[string]string{
		"status":   "ok",
		"version":  "1.0.6",
		"database": "ok",
	})
}
