package holding_api

import (
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/pkg/holding/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

const Path = "/holdings"

type Handler struct {
	Service Service
}

type Service interface {
	GetHolding(usuario string) (holding_domain.Holdings, error)
}

func NewHandler(service Service) Handler {
	return Handler{Service: service}
}

// GetHolding godoc
// @Summary Retorna Resultados holding do usuario
// @Produce json
// @Success 200 {object} holding_api.Holdings
// @Param usuario query string true "user id"
// @Param expandir query bool false "expandir ativos"
// @Router /holding [get]
func (h Handler) GetHolding(c *gin.Context) {
	usuario := c.Query("usuario")
	expandirAtivos := false
	expandirParam := c.Query("expandir")
	if strings.EqualFold(usuario, "") {
		err := fmt.Errorf("id do usuário é informação obrigatória.")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if !strings.EqualFold(expandirParam, "") {
		expandirAtivos, _ = strconv.ParseBool(expandirParam)
	}
	holdings, err := h.Service.GetHolding(usuario)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	holdingResponse := convertHoldingsDomainToDto(holdings, expandirAtivos)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, holdingResponse)
}