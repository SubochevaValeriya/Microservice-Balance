package microservice

import "time"

type UsersBalance struct {
	UserId  int `json:"-"`
	Balance int `json:"balance"`
}

type Transactions struct {
	UserId     int
	Amount     int    `json:"amount"`
	Reason     string `json:"reason"`
	TransferId int
	Date       time.Time `json:"date"`
}

func main() {
}
