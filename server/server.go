package server

import (
	"github.com/crisaltmann/fundament-stock-api/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config *config.Config
	Server *gin.Engine
}

func InitServer(s *Server) {
	s.Server.Run(s.config.ApplicationConfig.Address)
}