package microservice

import "time"

type UsersBalances struct {
	Id      int `json:"-"`
	Balance int `json:"balance"`
}

type Transactions struct {
	Id         int
	UserId     int
	Amount     int    `json:"amount"`
	Reason     string `json:"reason"`
	TransferId int
	Date       time.Time `json:"transaction_date"`
}

func main() {
}
