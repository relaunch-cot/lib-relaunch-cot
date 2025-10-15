package models

type User struct {
	UserId   string
	Name     string
	Email    string
	Password string
	Settings UserSettings
	Type     string
}

type UserSettings struct {
	Phone       string
	Cpf         string
	DateOfBirth string
	Biography   string
	Skills      []string
}
