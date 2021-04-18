package asset

import (
	"github.com/crisaltmann/fundament-stock-api/asset/api"
	"github.com/crisaltmann/fundament-stock-api/asset/service"
	"github.com/crisaltmann/fundament-stock-api/server"
	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
	fx.Invoke(MapRouter),
)

var factories = fx.Provide(
	NewService,
	NewHandler,
)

func NewService() *service.Service {
	return &service.Service{}
}

func NewHandler(service *service.Service) *api.Handler {
 	return &api.Handler{service}
}

func MapRouter(server *server.Server, handler *api.Handler) {
	server.Server.GET(api.Path + "s", handler.GetAllAssets)
}
