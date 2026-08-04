package models

type Cost struct {
	TodayCost       float64 `json:"today_cost"`
	MonthlyCost     float64 `json:"monthly_cost"`
	Budget          float64 `json:"budget"`
	RemainingBudget float64 `json:"remaining_budget"`
}
