package order_api

import (
	"github.com/crisaltmann/fundament-stock-api/pkg/order/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

const Path = "/orders"

type Handler struct {
	Service Service
}

type Service interface {
	GetAllOrders() ([]order_domain.Order, error)
	InsertOrder(order order_domain.Order) (bool, error)
}

func NewHandler(service Service) Handler {
	return Handler{Service: service}
}

// GetOrders godoc
// @Summary Retorna a lista de ordens
// @Produce json
// @Success 200 {object} order_api.OrderGetResponse
// @Router /orders [get]
func (h Handler) GetAllOrders(c *gin.Context) {
	orders, err := h.Service.GetAllOrders()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ordersResponse, err := convertDomainsToDtos(orders)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, ordersResponse)
}

// InsertOrders godoc
// @Summary Insere Ordem
// @Produce json
// @Param user body order_api.OrderPostRequest true "User-Data"
// @Success 201
// @Router /orders [post]
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
