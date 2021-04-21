package server

import (
	"github.com/crisaltmann/fundament-stock-api/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
)

var factories = fx.Provide(
	ConfigureServer,
)

func ConfigureServer(conf *config.Config) *Server {
	r := gin.New()
	server := &Server{conf, r}
	return server
}


