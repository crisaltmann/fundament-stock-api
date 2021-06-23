package insight_api

import (
	"fmt"
	insight_domain "github.com/crisaltmann/fundament-stock-api/pkg/insights/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
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


// GetInsightsSummary godoc
// @Summary Retorna sumario insights
// @Produce json
// @Success 200 {object} []insight_api.InsightsSummary
// @Param usuario path int true "user id"
// @Router /insights-summary [get]
func (h Handler) GetInsightsSummary(c *gin.Context) {
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

	insightsSummary, err := h.insightService.GetSummaryInsights(userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	insightsResponse := convertInsightsSummaryDomainToDto(insightsSummary)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	sort.SliceStable(insightsResponse.Insights, func(i, j int) bool {
		return insightsResponse.Insights[i].Trimestre < insightsResponse.Insights[j].Trimestre
	})

	c.JSON(http.StatusOK, insightsResponse)
}