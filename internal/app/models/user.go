package models

import "fmt"

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Cpf       string `json:"cpf"`
	Password  string `json:"-"`
	Balance   int    `json:"balance"`
	CreatedAt string `json:"createdAt"`
}

func (user User) String() string {
	return fmt.Sprintf("User { Id: %d, Name: '%s', Email: '%s', Cpf: '%s', Balance: %d, CreatedAt: '%s' }", user.Id, user.Name, user.Email, user.Cpf, user.Balance, user.CreatedAt)
}
