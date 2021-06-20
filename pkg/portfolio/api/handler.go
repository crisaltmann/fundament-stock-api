package portfolio_api

import (
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/pkg/portfolio/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

const Path = "/portfolio"

type Handler struct {
	Service Service
}

type Service interface {
	GetPortfolio(usuario int64) ([]portfolio_domain.Portfolio, error)
}

func NewHandler(service Service) Handler {
	return Handler{Service: service}
}

// GetPortfolio godoc
// @Summary Retorna portfolio do usuario
// @Produce json
// @Success 200 {object} portfolio_api.PortfolioGetResponse
// @Param usuario query string true "user id"
// @Router /portfolio [get]
func (h Handler) GetPortfolio(c *gin.Context) {
	usuarioParam := c.Query("usuario")
	if strings.EqualFold(usuarioParam, "") {
		err := fmt.Errorf("id do usuário é informação obrigatória.")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	usuario, err := strconv.ParseInt(usuarioParam, 10, 64)
	if err != nil {
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