package services

import "cloud-financial-analytics/internal/models"

func GetBudget() models.Budget {
	return models.Budget{
		TotalBudget:     1000,
		Spent:           400.00,
		RemainingBudget: 600.00,
		Currency:        "USD",
	}
}
