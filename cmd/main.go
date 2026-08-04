package main

import (
	"fmt"
	"net/http"

	"cloud-financial-analytics/internal/api"
)

func main() {

	http.HandleFunc("/", api.HomeHandler)
	http.HandleFunc("/health", api.HealthHandler)
	http.HandleFunc("/metrics", api.MetricsHandler)
	http.HandleFunc("/cost", api.CostHandler)
	http.HandleFunc("/alerts", api.AlertsHandler)
	fmt.Println("Cloud-Financial-Analytics API is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
