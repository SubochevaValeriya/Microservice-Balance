package handler

import (
	"github.com/SubochevaValeriya/microservice-balance"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// createUser is made for first input of user's transaction
// It's CREATE in CRUD
func (h *Handler) createUser(c *gin.Context) {
	var input microservice.UsersBalances

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Balance.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]int{
		"id": id,
	})
}

type getAllUsersBalancesResponse struct {
	Data []microservice.UsersBalances `json:"data"`
}

// getAllUsersBalances is used to get information of balances of all users
// It's READ from CRUD
func (h *Handler) getAllUsersBalances(c *gin.Context) {

	usersBalances, err := h.services.Balance.GetAllUsersBalances()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllUsersBalancesResponse{
		Data: usersBalances,
	})
}

// UPDATE
func (h *Handler) changeUsersBalances(c *gin.Context) {
	var input microservice.Transactions

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	row, err := h.services.Balance.ChangeBalances(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, row)
}

func (h *Handler) deleteAllUsersBalances(c *gin.Context) {
	err := h.services.Balance.DeleteAllUsersBalances()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// getBalanceByID allows to get balance of specific user and convert in determined currency (if entered)
// It's READ from CRUD
func (h *Handler) getBalanceByID(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	ccy := c.Query("currency")
	list, err := h.services.Balance.GetBalanceById(userId, ccy)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

// changeBalanceByID used for changing balance of user (adding transaction)
// It's UPDATE from CRUD
func (h *Handler) changeBalanceByID(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input microservice.Transactions

	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	row, err := h.services.Balance.ChangeBalanceById(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, row)
}

// deleteUsersByID allows to delete user and all his transactions
// It's DELETE from CRUD
func (h *Handler) deleteUsersByID(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Balance.DeleteUserById(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

type getTransactionsByIDResponse struct {
	Data []microservice.Transactions `json:"data"`
}

// getTransactionsByID allows to get list of transactions of specific user and convert it in determined currency (if entered)
// It's READ from CRUD
func (h *Handler) getTransactionsByID(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	ccy := c.Query("currency")
	transactions, err := h.services.Balance.GetTransactionsById(userId, ccy)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getTransactionsByIDResponse{
		Data: transactions,
	})
}
