package api

import (
	"cloud-financial-analytics/internal/services"
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
	Version string `json:"version"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	health := services.GetHealth()

	json.NewEncoder(w).Encode(health)
}
