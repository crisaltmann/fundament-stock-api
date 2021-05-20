package portfolio_api

import (
	"fmt"
	portfolio_service2 "github.com/crisaltmann/fundament-stock-api/pkg/portfolio/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const Path = "/portfolio"

type Handler struct {
	Service *portfolio_service2.Service
}

// GetPortfolio godoc
// @Summary Retorna portfolio do usuario
// @Produce json
// @Success 200 {object} portfolio_api.PortfolioGetResponse
// @Param usuario query string true "user id"
// @Router /portfolio [get]
func (h Handler) GetPortfolio(c *gin.Context) {
	usuario := c.Query("usuario")
	if strings.EqualFold(usuario, "") {
		err := fmt.Errorf("id do usuário é informação obrigatória.")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	portfolio, err := h.Service.GetPortfolio(usuario)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	portfolioResponse, err := convertDomainsToDtos(portfolio)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, portfolioResponse)
}