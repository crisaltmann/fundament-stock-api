package insight_api

import (
	"fmt"
	insight_domain "github.com/crisaltmann/fundament-stock-api/pkg/insghts/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

const Path = "/insights"

type Handler struct {
	insightService InsightService
}

type InsightService interface {
	GetInsights(usuario int64) ([]insight_domain.Insight, error)
	GetSummaryInsights(usuario int64) (insight_domain.InsightsSummary, error)
}

func NewHandler(service InsightService) Handler {
	return Handler{insightService: service}
}

// GetInsights godoc
// @Summary Retorna insights
// @Produce json
// @Success 200 {object} []insight_api.Insight
// @Param usuario path int true "user id"
// @Router /insights [get]
func (h Handler) GetInsights(c *gin.Context) {
	usuarioParam := c.Query("usuario")
	if strings.EqualFold(usuarioParam, "") {
		err := fmt.Errorf("id do usuário é informação obrigatória.")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	userId, err := strconv.ParseInt(usuarioParam, 10, 64)
	if err != nil {
		err := fmt.Errorf("id do usuário informado é inválido.")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	insights, err := h.insightService.GetInsights(userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	insightsResponse := convertInsightsDomainToDto(insights)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, insightsResponse)
}

//
//// GetQuarter godoc
//// @Summary Retorna o trimestre
//// @Produce json
//// @Success 200 {object} quarter_api.TrimestreGetResponse
//// @Param id path int true "quarter id"
//// @Router /quarters/{id} [get]
//func (h Handler) GetQuarter(c *gin.Context) {
//	idParam := c.Param("id")
//	id, err := strconv.ParseInt(idParam, 10, 64)
//	if err != nil {
//		c.AbortWithError(http.StatusBadRequest, err)
//		return
//	}
//	quarter, err := h.QuarterService.GetQuarter(id)
//	if err != nil {
//		c.AbortWithError(http.StatusInternalServerError, err)
//		return
//	}
//	quarterResponse := convertDomainToDto(quarter)
//	if err != nil {
//		c.AbortWithError(http.StatusInternalServerError, err)
//		return
//	}
//	c.JSON(http.StatusOK, quarterResponse)
//}