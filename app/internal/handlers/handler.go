package handlers

import (
	"github.com/4halavandala/l0/app/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api/v1")
	{
		api.GET("/orders", h.GetAllOrders)
		api.GET("/orders/:id", h.GetOrderById)
		api.POST("/orders", h.CreateOrder)
	}

	return router
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}
