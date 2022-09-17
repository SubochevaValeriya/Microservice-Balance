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

	addTransactionQuery := fmt.Sprintf("INSERT INTO %s (user_id, amount, reason, transaction_date, transfer_id) values ($1, $2, $3, $4, $5)", transactionTable)
	_, err = tx.Exec(addTransactionQuery, id, user.Balance, ReasonOpening, time.Now(), id)
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
	var row microservice.UsersBalances
	var balance int
	getBalanceQuery := fmt.Sprintf("SELECT (balance) FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&balance, getBalanceQuery, userId)
	if err != nil {
		return row, err
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersTable)

	err = r.db.Get(&row, query, userId)

	return row, err
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

func (r *ApiPostgres) DeleteAllUsersBalances() error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	deleteTransactionsQuery := fmt.Sprintf("DELETE FROM %s", transactionTable)
	if _, err := r.db.Exec(deleteTransactionsQuery); err != nil {
		tx.Rollback()
		return err
	}

	deleteBalanceQuery := fmt.Sprintf("DELETE FROM %s", usersTable)
	_, err = tx.Exec(deleteBalanceQuery)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *ApiPostgres) ChangeBalanceById(userId int, transaction microservice.Transactions) (microservice.Transactions, error) {
	var balance, transactionId int
	var reason string
	var row microservice.Transactions

	if transaction.Amount == 0 {
		return row, errors.New("empty transaction amount")
	}
	getBalanceQuery := fmt.Sprintf("SELECT (balance) FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&balance, getBalanceQuery, userId)
	if err != nil {
		return row, err
	}

	if transaction.Amount > 0 {
		reason = ReasonReplenishment
	} else {
		if balance+transaction.Amount < 0 {
			return row, errors.New("balance can't be negative")
		}

		reason = ReasonWithdrawal
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return row, err
	}
	changeBalanceQuery := fmt.Sprintf("UPDATE %s SET balance=$1 WHERE id=$2", usersTable)
	_, err = tx.Exec(changeBalanceQuery, transaction.Amount+balance, userId)
	if err != nil {
		tx.Rollback()
		return row, err
	}

	addTransactionQuery := fmt.Sprintf("INSERT INTO %s (user_id, amount, reason, transaction_date, transfer_id) values ($1, $2, $3, $4, $5) RETURNING id", transactionTable)
	rowInserted := r.db.QueryRow(addTransactionQuery, userId, transaction.Amount, reason, time.Now(), userId)
	if err := rowInserted.Scan(&transactionId); err != nil {
		tx.Rollback()
		return row, err
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", transactionTable)
	err = r.db.Get(&row, query, transactionId)
	if err != nil {
		return row, err
	}

	return row, tx.Commit()
}

func (r *ApiPostgres) ChangeBalances(transaction microservice.Transactions) (microservice.Transactions, error) {
	reason := ReasonTransfer
	var row microservice.Transactions

	if transaction.Amount == 0 {
		return row, errors.New("empty transaction amount")
	}

	if transaction.Amount < 0 {
		return row, errors.New("transaction amount should be positive")
	}

	balanceSender, err := r.GetBalanceAndCheck(transaction.UserId, -transaction.Amount)
	if err != nil {
		return row, err
	}

	balanceReceiver, err := r.GetBalanceAndCheck(transaction.UserId, transaction.Amount)
	if err != nil {
		return row, err
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return row, err
	}

	transactionId, err := r.UpdateBalanceAndInsertTransaction(tx, transaction.UserId, balanceSender-transaction.Amount, -transaction.Amount, transaction.TransferId, reason)
	if err != nil {
		return row, err
		tx.Rollback()
	}
	_, err = r.UpdateBalanceAndInsertTransaction(tx, transaction.TransferId, balanceReceiver+transaction.Amount, transaction.Amount, transaction.UserId, reason)
	if err != nil {
		return row, err
		tx.Rollback()
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", transactionTable)
	err = r.db.Get(&row, query, transactionId)
	if err != nil {
		return row, err
	}

	return row, tx.Commit()
}

func (r *ApiPostgres) GetTransactionsById(userId int) ([]microservice.Transactions, error) {
	var transactions []microservice.Transactions

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", transactionTable)
	err := r.db.Select(&transactions, query, userId)

	return transactions, err
}

func (r *ApiPostgres) GetBalanceAndCheck(userId, amount int) (int, error) {

	balance, err := r.GetBalanceById(userId)
	if err != nil {
		return 0, fmt.Errorf("user not found in the system: %w", err)
	}

	if balance.Balance+amount < 0 {
		return 0, errors.New("balance can't be negative")
	}

	return balance.Balance, nil
}

func (r *ApiPostgres) UpdateBalanceAndInsertTransaction(tx *sqlx.Tx, userId, balance, amount, transferId int, reason string) (int, error) {

	var transactionId int

	changeBalanceQuery := fmt.Sprintf("UPDATE %s SET balance=$1 WHERE id=$2", usersTable)
	_, err := tx.Exec(changeBalanceQuery, balance, userId)
	if err != nil {
		return 0, err
	}

	addTransactionQuery := fmt.Sprintf("INSERT INTO %s (user_id, amount, reason, transaction_date, transfer_id) values ($1, $2, $3, $4, $5) RETURNING id", transactionTable)
	rowInserted := r.db.QueryRow(addTransactionQuery, userId, amount, reason, time.Now(), transferId)
	if err = rowInserted.Scan(&transactionId); err != nil {
		return 0, err
	}

	return transactionId, nil
}
