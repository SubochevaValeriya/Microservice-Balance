package handler

import (
	"github.com/SubochevaValeriya/microservice-balance"
	"github.com/gin-gonic/gin"
	"net/http"
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

func (h *Handler) getAllUsersBalances(c *gin.Context) {
	// SELECT * FROM UsersBalances
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

func (h *Handler) changeUsersBalances(c *gin.Context) {
	//UPDATE
	//userId, err := c.Get()
}

func (h *Handler) deleteAllUsersBalances(c *gin.Context) {
	// DELETE
}

func (h *Handler) getBalanceByID(c *gin.Context) {
	// SELECT
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

func (h *Handler) changeBalanceByID(c *gin.Context) {
	//UPDATE
}

func (h *Handler) deleteByID(c *gin.Context) {
	//DELETE
}
