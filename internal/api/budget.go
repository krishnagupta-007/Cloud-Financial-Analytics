package api

import (
	"encoding/json"
	"net/http"

	"cloud-financial-analytics/internal/services"
)

func BudgetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	budget := services.GetBudget()
	json.NewEncoder(w).Encode(budget)
}
