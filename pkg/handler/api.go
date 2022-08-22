package handler

import "github.com/gin-gonic/gin"

func (h *Handler) createUser(c *gin.Context) {
	//INSERT
}

func (h *Handler) getAllUsersBalances(c *gin.Context) {
	// SELECT * FROM UsersBalances
}

func (h *Handler) changeUsersBalances(c *gin.Context) {
	//UPDATE
}

func (h *Handler) deleteAllUsersBalances(c *gin.Context) {
	// DELETE
}

func (h *Handler) getBalanceByID(c *gin.Context) {
	// SELECT
}

func (h *Handler) changeBalanceByID(c *gin.Context) {
	//UPDATE
}

func (h *Handler) deleteByID(c *gin.Context) {
	//DELETE
}
