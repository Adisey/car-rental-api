package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db"
)

func Health(w http.ResponseWriter, r *http.Request) {
	if err := db.CheckHealth(r.Context()); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)

		_ = json.NewEncoder(w).Encode(api_models.HealthResponse{
			Status:   "ok",
			Version:  "1.0.6",
			Database: "ok",
		})

		return
	}

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(api_models.HealthResponse{
		Status:   "ok",
		Version:  "1.0.6",
		Database: "ok",
	})
}
