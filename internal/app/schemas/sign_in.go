package schemas

type SignInSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (schema *SignInSchema) IsValid() (bool, []string) {
	return true, []string{}
}
