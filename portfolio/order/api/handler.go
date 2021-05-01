package orderapi

import (
	orderservice "github.com/crisaltmann/fundament-stock-api/portfolio/order/service"
	"github.com/gin-gonic/gin"
)

const Path = "/portfolios/:portfolioId/order"

type Handler struct {
	Service *orderservice.Service
}

func (h Handler) InsertOrder(c *gin.Context) {
	order := OrderPostRequest{}
	c.BindJSON(&order)

	_, err := h.Service.InsertOrder(convertPostRequestToDomain(order))
	//if err != nil {
	//	c.AbortWithError(http.StatusInternalServerError, err)
	//}
	//c.JSON(http.StatusCreated, nil)
}
