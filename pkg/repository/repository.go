package repository

import (
	"github.com/SubochevaValeriya/microservice-balance"
	"github.com/jmoiron/sqlx"
)

type Balance interface {
	CreateUser(user microservice.UsersBalances) (int, error)
	GetAllUsersBalances() ([]microservice.UsersBalances, error)
	GetBalanceById(userId int) (microservice.UsersBalances, error)
	DeleteUserById(userId int) error
	DeleteAllUsersBalances() error
	ChangeBalanceById(userId int, transaction microservice.Transactions) (int, error)
}

type Repository struct {
	Balance
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{NewApiPostgres(db)}
}
