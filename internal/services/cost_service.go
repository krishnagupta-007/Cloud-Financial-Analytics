package services

import "cloud-financial-analytics/internal/models"

func GetCost() models.Cost {

	return models.Cost{
		TodayCost:       100.0,
		MonthlyCost:     2000.0,
		Budget:          5000.0,
		RemainingBudget: 3000.0,
	}
}
