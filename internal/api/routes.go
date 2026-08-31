package api

import (
	"encoding/json"
	"net/http"
)

type HomeResponse struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Status  string `json:"status"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := HomeResponse{
		Name:    "Cloud-Financial_Analytics",
		Version: "1.0.0",
		Status:  "Running",
	}
	json.NewEncoder(w).Encode(response)
}
