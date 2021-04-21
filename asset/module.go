package asset

import (
	"github.com/crisaltmann/fundament-stock-api/asset/api"
	"github.com/crisaltmann/fundament-stock-api/asset/repository"
	"github.com/crisaltmann/fundament-stock-api/asset/service"
	"github.com/crisaltmann/fundament-stock-api/config"
	"github.com/crisaltmann/fundament-stock-api/server"
	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
	fx.Invoke(MapRouter),
)

var factories = fx.Provide(
	NewRepository,
	NewService,
	NewHandler,
)

func NewRepository(config *config.Config) *repository.Repository {
	return &repository.Repository{Config: config}
}

func NewService(repository *repository.Repository) *service.Service {
	return &service.Service{Repository: repository}
}

func NewHandler(service *service.Service) *api.Handler {
 	return &api.Handler{Service: service}
}

func MapRouter(server *server.Server, handler *api.Handler) {
	server.Server.GET(api.Path + "s", handler.GetAllAssets)
}
