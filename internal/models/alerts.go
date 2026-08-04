package models

type Alerts struct {
	ServerStatus string `json:"server_status"`
	CPUAlert     bool   `json:"cpu_alert"`
	MemoryAlert  bool   `json:"memory_alert"`
	DiskAlert    bool   `json:"disk_alert"`
	NetworkAlert bool   `json:"network_alert"`
}
