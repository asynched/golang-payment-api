package models

import "fmt"

type User struct {
	Id        int
	Name      string
	Email     string
	Cpf       string
	Password  string
	CreatedAt string
}

func (user User) String() string {
	return fmt.Sprintf("User { Id: %d, Name: '%s', Email: '%s', Cpf: '%s', CreatedAt: '%s' }", user.Id, user.Name, user.Email, user.Cpf, user.CreatedAt)
}
