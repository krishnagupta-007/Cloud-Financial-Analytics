package api

import (
	"encoding/json"
	"net/http"

	"cloud-financial-analytics/internal/services"
)

func CostHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	cost := services.GetCost()

	json.NewEncoder(w).Encode(cost)
}
