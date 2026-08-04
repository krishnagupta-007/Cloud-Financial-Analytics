package api

import (
	"encoding/json"
	"net/http"

	"cloud-financial-analytics/internal/services"
)

func AlertsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	alerts := services.GetAlerts()
	json.NewEncoder(w).Encode(alerts)
}
