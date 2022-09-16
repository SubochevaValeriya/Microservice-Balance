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

func (h *Handler) changeUsersBalances(c *gin.Context) {
	//UPDATE
	//userId, err := c.Get()
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
	//READ
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

func (h *Handler) changeBalanceByID(c *gin.Context) {
	//UPDATE
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

func (h *Handler) deleteUsersByID(c *gin.Context) {
	//DELETE
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

//func getUserId(c *gin.Context) (int, error) {
//	id, ok := c.Get("userId")
//	if !ok {
//		return 0, errors.New("user id not found")
//	}
//
//	idInt, ok := id.(int)
//	if !ok {
//		return 0, errors.New("user id is of invalid type")
//	}
//
//	return idInt, nil
//}
