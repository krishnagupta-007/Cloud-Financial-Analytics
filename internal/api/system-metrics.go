package api

import (
	"encoding/json"
	"net/http"

	"cloud-financial-analytics/internal/services"
)

func MetricsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	metrics := services.GetMetrics()
	json.NewEncoder(w).Encode(metrics)
}
