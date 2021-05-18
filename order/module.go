package order

import (
	"database/sql"
	asset_service "github.com/crisaltmann/fundament-stock-api/asset/service"
	order_api "github.com/crisaltmann/fundament-stock-api/order/api"
	order_repository "github.com/crisaltmann/fundament-stock-api/order/repository"
	order_service "github.com/crisaltmann/fundament-stock-api/order/service"

	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
	fx.Invoke(order_api.MapRouter),
)

var factories = fx.Provide(
	NewRepository,
	NewService,
	NewHandler,
)

func NewRepository(db *sql.DB) *order_repository.Repository {
	return &order_repository.Repository{DB: db}
}

func NewService(repository *order_repository.Repository, assetService *asset_service.Service) *order_service.Service {
	return &order_service.Service{Repository: repository, AssetService: assetService}
}

func NewHandler(service *order_service.Service) *order_api.Handler {
	return &order_api.Handler{Service: service}
}
