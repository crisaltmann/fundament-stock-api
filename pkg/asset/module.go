package asset

import (
	"database/sql"
	asset_api2 "github.com/crisaltmann/fundament-stock-api/pkg/asset/api"
	asset_repository2 "github.com/crisaltmann/fundament-stock-api/pkg/asset/repository"
	asset_service2 "github.com/crisaltmann/fundament-stock-api/pkg/asset/service"

	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
	fx.Invoke(asset_api2.MapRouter),
)

var factories = fx.Provide(
	NewRepository,
	NewService,
	NewHandler,
)

func NewRepository(db *sql.DB) *asset_repository2.Repository {
	return &asset_repository2.Repository{DB: db}
}

func NewService(repository *asset_repository2.Repository) *asset_service2.Service {
	return &asset_service2.Service{Repository: repository}
}

func NewHandler(service *asset_service2.Service) *asset_api2.Handler {
	return &asset_api2.Handler{Service: service}
}
