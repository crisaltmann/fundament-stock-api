package order_api

import (
	order_service "github.com/crisaltmann/fundament-stock-api/order/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

const Path = "/orders"

type Handler struct {
	Service *order_service.Service
}

func (h Handler) InsertOrder(c *gin.Context) {
	order := OrderPostRequest{}
	c.BindJSON(&order)
	_, err := h.Service.InsertOrder(convertPostRequestToDomain(order))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusCreated, nil)
}
