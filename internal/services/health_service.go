package services

import "cloud-financial-analytics/internal/models"

func GetHealth() models.Health {

	return models.Health{
		Service: "Cloud-Financial-Analytics",
		Status:  "Running",
		Uptime:  "2 hours",
		Version: "1.0.0",
	}
}
