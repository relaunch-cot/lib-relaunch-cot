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
	status                  string
	CreatedAt               string
	CreatedBy               string
}
