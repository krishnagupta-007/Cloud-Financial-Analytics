package models

type Resource struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Status string  `json:"status"`
	Cost   float64 `json:"cost"`
	Region string  `json:"region"`
}
