package api

import (
	"github.com/crisaltmann/fundament-stock-api/asset/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

const Path = "/asset"

type Handler struct {
	Service *service.Service
}

func (h Handler) GetAllAssets(c *gin.Context) {
	assets, err := h.Service.GetAllAssets()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, assets)
}