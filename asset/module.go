package asset

import (
	"database/sql"
	"github.com/crisaltmann/fundament-stock-api/asset/api"
	"github.com/crisaltmann/fundament-stock-api/asset/repository"
	"github.com/crisaltmann/fundament-stock-api/asset/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
	fx.Invoke(api.MapRouter),
)

var factories = fx.Provide(
	NewRepository,
	NewService,
	NewHandler,
)

func NewRepository(db *sql.DB) *repository.Repository {
	return &repository.Repository{DB: db}
}

func NewService(repository *repository.Repository) *service.Service {
	return &service.Service{Repository: repository}
}

func NewHandler(service *service.Service) *api.Handler {
 	return &api.Handler{Service: service}
}
