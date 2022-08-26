package repository

import (
	"errors"
	"fmt"
	"github.com/SubochevaValeriya/microservice-balance"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"net/http"
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

func (r *ApiPostgres) GetBalanceById(userId int, ccy string) (microservice.UsersBalances, error) {
	// https://apilayer.com/marketplace/exchangerates_data-api?preview=true#documentation-tab
	var row microservice.UsersBalances
	var balance int
	getBalanceQuery := fmt.Sprintf("SELECT (balance) FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&balance, getBalanceQuery, userId)
	if err != nil {
		return row, err
	}

	if ccy != "" {

	} else {
		url := "https://api.apilayer.com/exchangerates_data/convert?to={to}&from={from}&amount={amount}"

		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Set("apikey", "4EIJgz5GX9n5N8QUdRXHwQE01DfqGSqs")

		if err != nil {
			fmt.Println(err)
		}
		res, err := client.Do(req)
		if res.Body != nil {
			defer res.Body.Close()
		}
		body, err := ioutil.ReadAll(res.Body)

		fmt.Println(string(body))

		json.Unmarshal(responseData, &responseObject)
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersTable)

	err := r.db.Get(&row, query, userId)

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

	addTransactionQuery := fmt.Sprintf("INSERT INTO %s (user_id, amount, reason, transaction_date) values ($1, $2, $3, $4) RETURNING id", transactionTable)
	rowInserted := r.db.QueryRow(addTransactionQuery, userId, transaction.Amount, reason, time.Now())
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
