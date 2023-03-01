package handlers

import (
	"github.com/4halavandala/l0/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetAllOrdersResponse struct {
	Data *[]entity.Model `json:"Data"`
}

func (h *Handler) GetAllOrders(c *gin.Context) {
	orders, err := h.services.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, GetAllOrdersResponse{
		Data: orders,
	})
}

func (h *Handler) GetOrderById(c *gin.Context) {
	orderId := c.Param("id")
	if orderId == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid param")
		return
	}

	model, err := h.services.TodoOrder.GetById(orderId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, model)
}

func (h *Handler) CreateOrder(c *gin.Context) {
	var input entity.Model
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	model, err := h.services.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, model)
}

func (h *Handler) GetByIdFromCache(c *gin.Context) {
	orderId := c.Param("id")
	if orderId == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid param")
		return
	}

	order, err := h.services.GetByIdFromCache(orderId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *Handler) GetAllFromCache(c *gin.Context) {
	orders, err := h.services.GetAllFromCache()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, GetAllOrdersResponse{
		Data: orders,
	})
}
