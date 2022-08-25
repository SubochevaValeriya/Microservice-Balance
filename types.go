package microservice

import (
	_ "github.com/gin-gonic/gin"
	"time"
)

type UsersBalances struct {
	Id      int `json:"id" db:"id"`
	Balance int `json:"balance" db:"balance"`
}

type Transactions struct {
	Id         int
	UserId     int       `json:"userId" binding:"required"`
	Amount     int       `json:"amount" binding:"required"`
	Reason     string    `json:"reason" binding:"required"`
	TransferId int       `json:"transferId"`
	Date       time.Time `json:"transaction_date"`
}

func main() {
}

//type UsersBalances struct {
//	Id      int `json:"-" db:"id"`
//	Balance int `json:"balance" binding:"required"`
//}
//
//type Transactions struct {
//	Id         int
//	UserId     int       `json:"userId" binding:"required"`
//	Amount     int       `json:"amount" binding:"required"`
//	Reason     string    `json:"reason" binding:"required"`
//	TransferId int       `json:"transferId"`
//	Date       time.Time `json:"transaction_date"`
//}
