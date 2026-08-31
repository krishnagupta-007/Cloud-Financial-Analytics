package models

type Budget struct {
	TotalBudget     float64 `json:"total_budget"`
	Spent           float64 `json:"spent"`
	RemainingBudget float64 `json:"remaining_budget"`
	Currency        string  `json:"currency"`
}
