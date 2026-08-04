package services

import "cloud-financial-analytics/internal/models"

func GetAlerts() models.Alerts {

	return models.Alerts{
		ServerStatus: "Healthy",
		CPUAlert:     false,
		MemoryAlert:  true,
		DiskAlert:    false,
		NetworkAlert: false,
	}
}
