package asset_api

import (
	asset_service "github.com/crisaltmann/fundament-stock-api/asset/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const Path = "/assets"

type Handler struct {
	Service *asset_service.Service
}

func (h Handler) GetAllAssets(c *gin.Context) {
	assets, err := h.Service.GetAllAssets()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, convertToDtos(assets))
}

func (h Handler) InsertAsset(c *gin.Context) {
	asset := AssetPostRequest{}
	c.BindJSON(&asset)
	_, err := h.Service.InsertAsset(convertPostRequestToDomain(asset))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, nil)
}

func (h Handler) UpdateAsset(c *gin.Context) {
	asset := AssetPutRequest{}
	c.BindJSON(&asset)
	domainAsset, err := convertPutRequestToDomain(asset, c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	rasset, err := h.Service.UpdateAsset(domainAsset)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, rasset)
}

func (h Handler) GetById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	asset, err := h.Service.GetById(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, asset)
}
