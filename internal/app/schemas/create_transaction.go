package schemas

type CreateTransactionSchema struct {
	PayeeId int `json:"payeeId"`
	Value   int `json:"value"`
}

func (schema *CreateTransactionSchema) IsValid() (bool, []string) {
	var errors []string

	if schema.PayeeId <= 0 {
		errors = append(errors, "PayerId is required")
	}

	if schema.Value <= 0 {
		errors = append(errors, "Value must be greater than 0")
	}

	return len(errors) == 0, errors
}
