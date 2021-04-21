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
	//asset := AssetPostRequest{}
	//c.BindJSON(&asset)
	//_, err := h.Service.InsertAsset(convertPostRequestToDomain(asset))
	//if err != nil {
	//	c.AbortWithError(http.StatusInternalServerError, err)
	//}
	//c.JSON(http.StatusCreated, nil)
}
