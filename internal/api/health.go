package api

import (
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
	response := HealthResponse{
		Status:  "Healthy",
		Service: "Cloud-Financial-Analytics API",
		Version: "1.0.0",
	}

	json.NewEncoder(w).Encode(response)
}
