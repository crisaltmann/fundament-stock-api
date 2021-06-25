package user_api

import (
	user_domain "github.com/crisaltmann/fundament-stock-api/pkg/user/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

const Path_Login = "/login"

type Handler struct {
	service Service
}

type Service interface {
	Login(user string, password string) (user_domain.Login, error)
}

func NewHandler(service Service) Handler {
	return Handler{service: service}
}

// Login godoc
// @Summary Login
// @Produce json
// @Success 200 {object} user_api.LoginResponse
// @Router /login [post]
func (h Handler) Login(c *gin.Context) {
	loginRequest := LoginRequest{}
	err := c.BindJSON(&loginRequest)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	login, err := h.service.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	c.JSON(http.StatusOK, convertDomainToDto(login))
}