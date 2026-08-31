package api

import (
	"encoding/json"
	"net/http"

	"cloud-financial-analytics/internal/services"
)

func ResourcesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services.GetResources())
}
