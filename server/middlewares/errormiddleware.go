package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": c.Errors.Last().Error()})
	}

}
