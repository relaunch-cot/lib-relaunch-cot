package models

type User struct {
	UserId   string
	Name     string
	Email    string
	Password string
}

type UserSettings struct {
	Phone       string
	Cpf         string
	DateOfBirth string
	Biography   string
	Skills      []string
}
