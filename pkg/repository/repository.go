package repository

import (
	"database/sql"
	"github.com/SubochevaValeriya/microservice-balance"
	"github.com/jmoiron/sqlx"
)

type Balance interface {
	CreateUser(user microservice.UsersBalances) (int, error)
	GetAllUsersBalances(user microservice.UsersBalances) (*sql.Row, error)
	GetBalanceById(user microservice.UsersBalances) (*sql.Row, error)
}

type Repository struct {
	Balance
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{NewApiPostgres(db)}
}
