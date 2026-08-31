package models

type Health struct {
	Service string `json:"service"`
	Status  string `json:"status"`
	Uptime  string `json:"uptime"`
	Version string `json:"version"`
}
