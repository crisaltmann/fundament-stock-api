package server

import (
	"github.com/crisaltmann/fundament-stock-api/config"
	"github.com/gin-gonic/gin"

	"log"
	"os"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/crisaltmann/fundament-stock-api/docs"
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

	//url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	s.Server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	s.Server.Run(":" + port)
}
