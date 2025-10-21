package models

type Project struct {
	ProjectId               string
	ClientId                string
	FreelancerId            string
	Name                    string
	Description             string
	Category                string
	ProjectDeliveryDeadline string
	Amount                  float32
	RemainingTime           string
	ClientName              string
	FreelancerName          string
	UrlImageProject         string
	Status                  string
	CreatedAt               string
	CreatedBy               string
	UpdatedAt               string
}
