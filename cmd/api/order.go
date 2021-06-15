package api

import (
	"github.com/crisaltmann/fundament-stock-api/pkg/asset/service"
	"github.com/crisaltmann/fundament-stock-api/pkg/order/api"
	order_event "github.com/crisaltmann/fundament-stock-api/pkg/order/event"
	"github.com/crisaltmann/fundament-stock-api/pkg/order/repository"
	"github.com/crisaltmann/fundament-stock-api/pkg/order/service"

	"go.uber.org/fx"
)

var Order = fx.Options(
	orderfactories,
	fx.Invoke(order_api.MapRouter),
)

var orderfactories = fx.Provide(
	order_repository.NewRepository,
	func(repository order_repository.Repository) order_service.Repository { return repository },

	order_service.NewService,
	func(service order_service.Service) order_api.Service { return service },
	func(assetService asset_service.Service) order_service.AssetFinder { return assetService },

	order_event.NewOrderProducer,

	order_api.NewHandler,
)
