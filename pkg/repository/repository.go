package repository

import "github.com/jmoiron/sqlx"

type Balance interface {
}

type Repository struct {
	Balance
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
