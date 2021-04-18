package api

import (
	"github.com/crisaltmann/fundament-stock-api/asset/service"
	"github.com/gin-gonic/gin"
)

const Path = "/asset"

type Handler struct {
	Service *service.Service
}

func (h Handler) GetAllAssets(c *gin.Context) {
	assets := h.Service.GetAllAssets()
	c.JSON(200, assets)
}