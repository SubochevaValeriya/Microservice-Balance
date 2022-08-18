package handler

import (
	"github.com/SubochevaValeriya/microservice-balance/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/", h.createUser)
		api.GET("/", h.getAllUsersBalances)
		api.PUT("/", h.changeUsersBalances)
		api.DELETE("/", h.deleteAllUsersBalances)
		api.GET("/:id", h.getBalanceByID)
		api.PUT("/:id", h.changeBalanceByID)
		api.DELETE("/:id", h.deleteByID)
	}

	return router
}
