package repository

import (
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

func (r *ApiPostgres) GetAllUsersBalances() ([]microservice.UsersBalances, error) {
	var usersBalances []microservice.UsersBalances

	query := fmt.Sprintf("SELECT * FROM %s", usersTable)
	err := r.db.Select(&usersBalances, query)

	return usersBalances, err
}

func (r *ApiPostgres) GetBalanceById(userId int) (microservice.UsersBalances, error) {
	var list microservice.UsersBalances
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersTable)
	err := r.db.Get(&list, query, userId)
	//if err := row.Scan(&id); err != nil {
	//	return 0, err
	//}

	return list, err
}

func (r *ApiPostgres) DeleteUserById(userId int) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	deleteTransactionsQuery := fmt.Sprintf("DELETE FROM %s WHERE id = $1", transactionTable)
	if _, err := r.db.Exec(deleteTransactionsQuery, userId); err != nil {
		tx.Rollback()
		return err
	}

	deleteBalanceQuery := fmt.Sprintf("DELETE FROM %s WHERE id = $1", usersTable)
	_, err = tx.Exec(deleteBalanceQuery, userId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
