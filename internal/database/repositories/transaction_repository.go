package repositories

import (
	"context"
	"database/sql"

	"github.com/asynched/golang-payment-api/internal/app/models"
)

type TransactionRepository struct {
	db *sql.DB
}

type CreateTransactionInput struct {
	PayerId int
	PayeeId int
	Value   int
}

var createTransactionQuery = `
	INSERT INTO transactions (payer_id, payee_id, value)
	VALUES ($1, $2, $3);
`

var updatePayerBalanceQuery = `
	UPDATE users SET balance = balance - $1 WHERE id = $2;
`

var updatePayeeBalanceQuery = `
	UPDATE users SET balance = balance + $1 WHERE id = $2;
`

func (repository *TransactionRepository) CreateTransaction(input CreateTransactionInput) error {
	tx, err := repository.db.BeginTx(context.Background(), nil)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec(createTransactionQuery, input.PayerId, input.PayeeId, input.Value)

	if err != nil {
		return err
	}

	_, err = tx.Exec(updatePayerBalanceQuery, input.Value, input.PayerId)

	if err != nil {
		return err
	}

	_, err = tx.Exec(updatePayeeBalanceQuery, input.Value, input.PayeeId)

	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

type GetTransactionsInput struct {
	UserId int
}

var getTransactionsQuery = `
	SELECT id, payer_id, payee_id, value, created_at FROM transactions WHERE payer_id = $1 OR payee_id = $1;
`

func (repository *TransactionRepository) GetTransactions(input GetTransactionsInput) ([]models.Transaction, error) {
	rows, err := repository.db.Query(getTransactionsQuery, input.UserId)

	if err != nil {
		return nil, err
	}

	transactions := []models.Transaction{}

	for rows.Next() {
		transaction := models.Transaction{}

		err := rows.Scan(&transaction.Id, &transaction.PayerId, &transaction.PayeeId, &transaction.Value, &transaction.CreatedAt)

		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

type GetTransactionByIdInput struct {
	Id     int
	UserId int
}

var getTransactionByIdQuery = `
	SELECT id, payer_id, payee_id, value, created_at FROM transactions WHERE id = $1 AND (payer_id = $2 OR payee_id = $2);
`

func (repositories *TransactionRepository) GetTransactionById(input GetTransactionByIdInput) (models.Transaction, error) {
	row := repositories.db.QueryRow(getTransactionByIdQuery, input.Id, input.UserId)

	transaction := models.Transaction{}

	err := row.Scan(&transaction.Id, &transaction.PayerId, &transaction.PayeeId, &transaction.Value, &transaction.CreatedAt)

	if err != nil {
		return models.Transaction{}, err
	}

	return transaction, nil
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db}
}
