package models

type Project struct {
	ProjectId               string
	ClientId                string
	DeveloperId             string
	Category                string
	ProjectDeliveryDeadline string
	Amount                  float64
	RemainingTime           string
	CreatedAt               string
	CreatedBy               string
}
