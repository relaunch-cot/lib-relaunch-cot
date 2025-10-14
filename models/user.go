package models

type User struct {
	UserId         int
	Name           string
	Email          string
	HashedPassword string
}
