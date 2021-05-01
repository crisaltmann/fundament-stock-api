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
	port := s.config.ApplicationConfig.Address
	//port = os.Getenv("PORT")
	s.Server.Run(port)
}