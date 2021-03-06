package quarter_api

import (
	"github.com/crisaltmann/fundament-stock-api/pkg/quarter/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const Path = "/quarters"

type Handler struct {
	QuarterService QuarterService
}

type QuarterService interface {
	GetQuarter(id int64) (quarter_domain.Trimestre, error)
	GetQuarters() ([]quarter_domain.Trimestre, error)
}

func NewHandler(service QuarterService) Handler {
	return Handler{QuarterService: service}
}

// GetQuarters godoc
// @Summary Retorna trimestres
// @Produce json
// @Success 200 {object} []quarter_api.TrimestreGetResponse
// @Router /quarters [get]
func (h Handler) GetQuarters(c *gin.Context) {
	quarters, err := h.QuarterService.GetQuarters()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	portfolioResponse := convertToDtos(quarters)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, portfolioResponse)
}

// GetQuarter godoc
// @Summary Retorna o trimestre
// @Produce json
// @Success 200 {object} quarter_api.TrimestreGetResponse
// @Param id path int true "quarter id"
// @Router /quarters/{id} [get]
func (h Handler) GetQuarter(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	quarter, err := h.QuarterService.GetQuarter(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	quarterResponse := convertDomainToDto(quarter)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, quarterResponse)
}