package asset_sync_api

import (
	"net/http"

	asset_sync "github.com/crisaltmann/fundament-stock-api/pkg/asset-sync"
	"github.com/gin-gonic/gin"
)

const Path = "/assets-price"

type Handler struct {
	service asset_sync.JobService
}

func NewHandler(service asset_sync.JobService) Handler {
	return Handler{service: service}
}

// UpdateAssetPrice godoc
// @Summary Atualiza cotação ativo
// @Success 201
// @Router /assets-price [post]
func (h Handler) UpdateAssetPrice(c *gin.Context) {

	h.service.UpdateAssetPrice()

	c.JSON(http.StatusOK, nil)
}
