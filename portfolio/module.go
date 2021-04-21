package asset

import (
	"database/sql"
	orderapi "github.com/crisaltmann/fundament-stock-api/portfolio/order/api"
	orderrepository "github.com/crisaltmann/fundament-stock-api/portfolio/order/repository"
	orderservice "github.com/crisaltmann/fundament-stock-api/portfolio/order/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
	fx.Invoke(orderapi.MapRouter),
)

var factories = fx.Provide(
	NewOrderRepository,
	NewOrderService,
	NewOrderHandler,
)

func NewOrderRepository(db *sql.DB) *orderrepository.Repository {
	return &orderrepository.Repository{DB: db}
}

func NewOrderService(repository *orderrepository.Repository) *orderservice.Service {
	return &orderservice.Service{Repository: repository}
}

func NewOrderHandler() *orderapi.Handler {
	return &orderapi.Handler{}
}
