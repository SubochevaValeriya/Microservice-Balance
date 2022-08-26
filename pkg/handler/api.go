package handler

import (
	"github.com/SubochevaValeriya/microservice-balance"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createUser(c *gin.Context) {
	//INSERT
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

func (h *Handler) getAllUsersBalances(c *gin.Context) {
	// SELECT * FROM UsersBalances

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

func (h *Handler) getBalanceByID(c *gin.Context) {
	// SELECT
	//userId, err := getUserId(c)
	//if err != nil {
	//	return
	//}

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	list, err := h.services.Balance.GetBalanceById(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) changeBalanceByID(c *gin.Context) {
	//UPDATE
	userId, err := strconv.Atoi(c.Param("id"))
	var input microservice.Transactions

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Balance.ChangeBalanceById(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]int{
		"id": id,
	})
}

func (h *Handler) deleteUsersByID(c *gin.Context) {
	//DELETE
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
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
