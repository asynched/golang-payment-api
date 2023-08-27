package models

type Transaction struct {
	Id        int    `json:"id"`
	PayerId   int    `json:"payerId"`
	PayeeId   int    `json:"payeeId"`
	Value     int    `json:"value"`
	CreatedAt string `json:"createdAt"`
}
