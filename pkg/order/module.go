package order

import (
	"database/sql"
	asset_service2 "github.com/crisaltmann/fundament-stock-api/pkg/asset/service"
	order_api2 "github.com/crisaltmann/fundament-stock-api/pkg/order/api"
	order_repository2 "github.com/crisaltmann/fundament-stock-api/pkg/order/repository"
	order_service2 "github.com/crisaltmann/fundament-stock-api/pkg/order/service"

	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
	fx.Invoke(order_api2.MapRouter),
)

var factories = fx.Provide(
	NewRepository,
	NewService,
	NewHandler,
)

func NewRepository(db *sql.DB) *order_repository2.Repository {
	return &order_repository2.Repository{DB: db}
}

func NewService(repository *order_repository2.Repository, assetService *asset_service2.Service) *order_service2.Service {
	return &order_service2.Service{Repository: repository, AssetService: assetService}
}

func NewHandler(service *order_service2.Service) *order_api2.Handler {
	return &order_api2.Handler{Service: service}
}
