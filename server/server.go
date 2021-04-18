package server

import (
	"github.com/crisaltmann/fundament-stock-api/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	config *config.Config
}

func ConfigureServer(conf *config.Config) *Server {
	server := &Server{conf}
	return server
}

func InitServer(s *Server) {
	r := gin.New()
	addRouterGin(r)
	r.Run(s.config.ApplicationConfig.Address)
}

func addRouterGin(r *gin.Engine) {
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"greet": "hello, world!",
		})
	})

	r.GET("/echo/:echo", func(c *gin.Context) {
		echo := c.Param("echo")
		c.JSON(http.StatusOK, gin.H{
			"echo": echo,
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ping": "Pong",
		})
	})
}