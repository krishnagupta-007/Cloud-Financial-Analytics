package main

import (
	"fmt"
	"net/http"

	"cloud-financial-analytics/internal/api"
	"cloud-financial-analytics/internal/monitoring"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Register Prometheus metrics
	monitoring.RegisterMetrics()

	// Start background metric collector
	monitoring.StartCollector()

	// Existing API handlers wrapped with metrics middleware to record HTTP metrics
	http.HandleFunc("/", monitoring.MetricsMiddleware("/", api.HomeHandler))
	http.HandleFunc("/health", monitoring.MetricsMiddleware("/health", api.HealthHandler))
	http.HandleFunc("/system-metrics", monitoring.MetricsMiddleware("/system-metrics", api.MetricsHandler))
	http.HandleFunc("/cost", monitoring.MetricsMiddleware("/cost", api.CostHandler))
	http.HandleFunc("/alerts", monitoring.MetricsMiddleware("/alerts", api.AlertsHandler))
	http.HandleFunc("/budget", monitoring.MetricsMiddleware("/budget", api.BudgetHandler))
	http.HandleFunc("/resources", monitoring.MetricsMiddleware("/resources", api.ResourcesHandler))

	// Expose Prometheus metrics endpoint
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Cloud-Financial-Analytics API is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
