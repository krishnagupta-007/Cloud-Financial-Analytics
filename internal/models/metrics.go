package models

type Metric struct {
	CPU     int `json:"cpu"`
	Memory  int `json:"memory"`
	Disk    int `json:"disk"`
	Network int `json:"network"`
}
