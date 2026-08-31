package services

import "cloud-financial-analytics/internal/models"

func GetResources() []models.Resource {
	return []models.Resource{
		{
			Name:   "Web-Server-01",
			Type:   "EC2",
			Status: "Running",
			Cost:   50.00,
			Region: "us-east-1",
		},
		{
			Name:   "Database-01",
			Type:   "RDS",
			Status: "Stopped",
			Cost:   100.00,
			Region: "us-west-2",
		},
		{
			Name:   "Storage-01",
			Type:   "S3",
			Status: "Available",
			Cost:   20.00,
			Region: "eu-central-1",
		},
	}
}
