package result_importer_api

import (
	result_importer_service "github.com/crisaltmann/fundament-stock-api/pkg/result_importer/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

const Path = "/importer"

type Handler struct {
	importer result_importer_service.Importer
}

func NewHandler(importer result_importer_service.Importer) Handler {
	return Handler{importer: importer}
}

func (h Handler) ImportAll(c *gin.Context) {
	err := h.importer.StartImporterAll()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h Handler) Import(c *gin.Context) {
	code := c.Query("code")
	err := h.importer.StartImporter(code)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}