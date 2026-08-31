package monitoring

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// Define Prometheus custom metrics
var (
	// CPU usage gauge percentage
	CPUUsage = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "cloud_cpu_usage_percentage",
			Help: "Current CPU usage percentage of the cloud instance",
		},
	)

	// Memory usage gauge percentage
	MemoryUsage = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "cloud_memory_usage_percentage",
			Help: "Current memory usage percentage of the cloud instance",
		},
	)

	// Disk usage gauge percentage (existing)
	DiskUsage = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "cloud_disk_usage_percentage",
			Help: "Current disk usage percentage of the cloud instance",
		},
	)

	// Total memory in bytes
	TotalMemory = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "cloud_memory_total_bytes",
			Help: "Total memory available in bytes",
		},
	)

	// Used memory in bytes
	UsedMemory = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "cloud_memory_used_bytes",
			Help: "Currently used memory in bytes",
		},
	)

	// Available memory in bytes
	AvailableMemory = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "cloud_memory_available_bytes",
			Help: "Currently available memory in bytes",
		},
	)

	// Total HTTP requests counter metric
	HTTPRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests processed",
		},
		[]string{"path", "method", "status"},
	)

	// HTTP request duration histogram metric
	HTTPRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request latency in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method"},
	)
)

// RegisterMetrics registers all Prometheus metrics with the default registry.
func RegisterMetrics() {
	prometheus.MustRegister(CPUUsage)
	prometheus.MustRegister(MemoryUsage)
	prometheus.MustRegister(DiskUsage)
	prometheus.MustRegister(TotalMemory)
	prometheus.MustRegister(UsedMemory)
	prometheus.MustRegister(AvailableMemory)
	prometheus.MustRegister(HTTPRequestsTotal)
	prometheus.MustRegister(HTTPRequestDuration)
}

// statusResponseWriter captures the HTTP status code for metrics tracking.
type statusResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (srw *statusResponseWriter) WriteHeader(code int) {
	srw.statusCode = code
	srw.ResponseWriter.WriteHeader(code)
}

// MetricsMiddleware tracks total requests and execution duration for HTTP endpoints.
func MetricsMiddleware(path string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		srw := &statusResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		handler(srw, r)

		duration := time.Since(start).Seconds()
		statusStr := fmt.Sprintf("%d", srw.statusCode)

		HTTPRequestsTotal.WithLabelValues(path, r.Method, statusStr).Inc()
		HTTPRequestDuration.WithLabelValues(path, r.Method).Observe(duration)
	}
}
