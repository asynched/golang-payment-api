package repositories

import (
	"database/sql"

	"github.com/asynched/golang-payment-api/internal/app/models"
)

type UserRepository struct {
	db *sql.DB
}

type CreateUserInput struct {
	Name     string
	Email    string
	Cpf      string
	Password string
}

const createUserQuery = `
	INSERT INTO users(name, email, cpf, password)
	VALUES ($1, $2, $3, $4);
`

func (repository *UserRepository) CreateUser(data CreateUserInput) error {
	_, err := repository.db.Exec(createUserQuery, data.Name, data.Email, data.Cpf, data.Password)

	return err
}

type GetUserByEmailInput struct {
	Email string
}

const getUserByEmailQuery = `
	SELECT id, name, email, cpf, password, balance, created_at FROM users WHERE email = $1;
`

func (repository *UserRepository) GetUserByEmail(data GetUserByEmailInput) (*models.User, error) {
	row := repository.db.QueryRow(getUserByEmailQuery, data.Email)

	user := models.User{}

	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Cpf, &user.Password, &user.Balance, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

type GetUserByIdInput struct {
	Id int
}

const getUserByIdQuery = `
	SELECT id, name, email, cpf, password, balance, created_at FROM users WHERE id = $1;
`

func (repository *UserRepository) GetUserById(data GetUserByIdInput) (*models.User, error) {
	row := repository.db.QueryRow(getUserByIdQuery, data.Id)

	user := models.User{}

	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Cpf, &user.Password, &user.Balance, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}
