package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/SubochevaValeriya/microservice-balance"
	"github.com/jmoiron/sqlx"
	"time"
)

type ApiPostgres struct {
	db *sqlx.DB
}

func NewApiPostgres(db *sqlx.DB) *ApiPostgres {
	return &ApiPostgres{db: db}
}

const (
	ReasonOpening       = "Opening"
	ReasonReplenishment = "Replenishment"
	ReasonWithdrawal    = "Withdrawal"
	ReasonTransfer      = "Transfer"
)

func (r *ApiPostgres) CreateUser(user microservice.UsersBalances) (int, error) {
	// пользователь мне даёт balance и всё (в будущем валюту), соотв я по умолчанию проставляю
	// ризон и амаунт, мне даже это не нужно от него
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}
	var id int

	if user.Balance < 0 {
		return 0, errors.New("balance can't be negative")
	}
	changeBalanceQuery := fmt.Sprintf("INSERT INTO %s (balance) values ($1) RETURNING id", usersTable)
	row := r.db.QueryRow(changeBalanceQuery, user.Balance)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	addTransactionQuery := fmt.Sprintf("INSERT INTO %s (user_id, amount, reason, transaction_date) values ($1, $2, $3, $4)", transactionTable)
	_, err = tx.Exec(addTransactionQuery, id, user.Balance, ReasonOpening, time.Now())
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *ApiPostgres) GetAllUsersBalances(user microservice.UsersBalances) (*sql.Row, error) {
	query := fmt.Sprintf("SELECT * FROM %s", usersTable)
	row := r.db.QueryRow(query)
	//if err := row.Scan(&id); err != nil {
	//	return 0, err
	//}

	return row, nil
}

func (r *ApiPostgres) GetBalanceById(user microservice.UsersBalances) (*sql.Row, error) {
	query := fmt.Sprintf("SELECT (balance) FROM %s WHERE id = $1", usersTable)
	row := r.db.QueryRow(query, user.Id)
	//if err := row.Scan(&id); err != nil {
	//	return 0, err
	//}

	return row, nil
}
