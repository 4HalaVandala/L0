package handlers

import (
	"github.com/4halavandala/l0/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(CORSMiddleware())

	api := router.Group("/api/v1/")
	{

		orders := api.Group("/orders")
		{
			orders.GET("", h.GetAllOrders)
			orders.GET(":id", h.GetOrderById)
			orders.POST("", h.CreateOrder)
		}

		cache := api.Group("/cache")
		{
			cache.GET("", h.GetAllFromCache)
			cache.GET(":id", h.GetByIdFromCache)
		}
	}

	return router
}
