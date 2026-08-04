package services

import "cloud-financial-analytics/internal/models"

func GetMetrics() models.Metric {

	return models.Metric{
		CPU:     50,
		Memory:  60,
		Disk:    70,
		Network: 80,
	}
}
