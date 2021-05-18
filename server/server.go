package server

import (
	"github.com/crisaltmann/fundament-stock-api/config"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

type Server struct {
	config *config.Config
	Server *gin.Engine
}

func InitServer(s *Server) {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("PORTA NAO ENCONTRADA, USANDO VALOR DEFAULT.")
		port = "8080"
	}
	s.Server.Run(":" + port)
}
