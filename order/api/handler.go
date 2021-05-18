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
	orderDomain, err := convertPostRequestToDomain(order)
	if err == nil {
		_, err = h.Service.InsertOrder(orderDomain)
	}
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, nil)
}
