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

//func (h Handler) GetAllAssets(c *gin.Context) {
//	assets, err := h.Service.GetAllAssets()
//	if err != nil {
//		c.AbortWithError(http.StatusInternalServerError, err)
//	}
//	c.JSON(http.StatusOK, convertToDtos(assets))
//}
//
func (h Handler) InsertOrder(c *gin.Context) {
	order := OrderPostRequest{}
	c.BindJSON(&order)
	_, err := h.Service.InsertOrder(convertPostRequestToDomain(order))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusCreated, nil)
}
//
//func (h Handler) UpdateAsset(c *gin.Context) {
//	asset := AssetPutRequest{}
//	c.BindJSON(&asset)
//	domainAsset, err := convertPutRequestToDomain(asset, c.Param("id"))
//	if err != nil {
//		c.AbortWithError(http.StatusBadRequest, err)
//	}
//	rasset, err := h.Service.UpdateAsset(domainAsset)
//	if err != nil {
//		c.AbortWithError(http.StatusInternalServerError, err)
//	}
//	c.JSON(http.StatusOK, rasset)
//}
//
//func (h Handler) GetById(c *gin.Context) {
//	idParam := c.Param("id")
//	id, err := strconv.ParseInt(idParam, 10, 64)
//	if err != nil {
//		c.AbortWithError(http.StatusBadRequest, err)
//	}
//	asset, err := h.Service.GetById(id)
//	if err != nil {
//		c.AbortWithError(http.StatusInternalServerError, err)
//	}
//	c.JSON(http.StatusOK, asset)
//}
