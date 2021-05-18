package asset

import (
	"database/sql"
	asset_service "github.com/crisaltmann/fundament-stock-api/asset/service"

	asset_api "github.com/crisaltmann/fundament-stock-api/asset/api"
	asset_repository "github.com/crisaltmann/fundament-stock-api/asset/repository"
	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
	fx.Invoke(asset_api.MapRouter),
)

var factories = fx.Provide(
	NewRepository,
	NewService,
	NewHandler,
)

func NewRepository(db *sql.DB) *asset_repository.Repository {
	return &asset_repository.Repository{DB: db}
}

func NewService(repository *asset_repository.Repository) *asset_service.Service {
	return &asset_service.Service{Repository: repository}
}

func NewHandler(service *asset_service.Service) *asset_api.Handler {
	return &asset_api.Handler{Service: service}
}
