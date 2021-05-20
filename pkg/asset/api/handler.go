package asset_api

import (
	asset_service2 "github.com/crisaltmann/fundament-stock-api/pkg/asset/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const Path = "/assets"

type Handler struct {
	Service *asset_service2.Service
}

// GetAssets godoc
// @Summary Retorna a lista de ativos
// @Produce json
// @Success 200 {object} asset_api.AssetResponse
// @Router /assets [get]
func (h Handler) GetAllAssets(c *gin.Context) {
	assets, err := h.Service.GetAllAssets()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, convertToDtos(assets))
}

// InsertAssets godoc
// @Summary Insere Ativo
// @Produce json
// @Param user body asset_api.AssetPostRequest true "User-Data"
// @Success 201
// @Router /assets [post]
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

// UpdateAssets godoc
// @Summary Atualiza Ativo
// @Produce json
// @Param user body asset_api.AssetPutRequest true "User-Data"
// @Param id path int true "Asset ID"
// @Success 200
// @Router /assets/{id} [put]
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

// GetAssets godoc
// @Summary Retorna a ativo
// @Produce json
// @Param id path int true "Asset ID"
// @Success 200 {object} asset_api.AssetResponse
// @Router /assets/{id} [get]
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