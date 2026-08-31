package monitoring

import (
	"math/rand"
	"runtime"
	"time"
)

// StartCollector runs a background goroutine to periodically collect
// and update CPU and Memory metrics for Prometheus scraping.
func StartCollector() {
	go func() {
		// Ticker fires every 5 seconds to update metric values
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			var memStats runtime.MemStats
			runtime.ReadMemStats(&memStats)

			// Total memory allocated from OS (bytes)
			totalMem := float64(memStats.Sys)
			// Currently allocated heap memory (bytes)
			usedMem := float64(memStats.Alloc)
			// Memory available (bytes)
			availMem := totalMem - usedMem

			// Calculate memory usage percentage
			var memUsagePct float64
			if totalMem > 0 {
				memUsagePct = (usedMem / totalMem) * 100.0
			}

			// Simulate dynamic CPU usage percentage for demo/learning
			cpuUsagePct := 15.0 + rand.Float64()*10.0

			// Update Gauges with latest system metrics
			CPUUsage.Set(cpuUsagePct)
			MemoryUsage.Set(memUsagePct)
			TotalMemory.Set(totalMem)
			UsedMemory.Set(usedMem)
			AvailableMemory.Set(availMem)
		}
	}()
}
