package schemas

type CreateUserSchema struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Cpf      string `json:"cpf"`
	Password string `json:"password"`
}

func (schema *CreateUserSchema) IsValid() (bool, []string) {
	errors := make([]string, 0)

	if len(schema.Name) < 3 || len(schema.Name) > 255 {
		errors = append(errors, "Invalid name, length is less than 3 or greater than 255")
	}

	if len(schema.Email) < 3 || len(schema.Email) > 255 {
		errors = append(errors, "Invalid email, length is less than 3 or greater than 255")
	}

	if len(schema.Cpf) != 11 {
		errors = append(errors, "Invalid CPF, length must be 11")
	}

	if len(schema.Password) < 8 || len(schema.Password) > 255 {
		errors = append(errors, "Invalid password, length is less than 8 or greater than 255")
	}

	return len(errors) == 0, errors
}
