package api

import (
	asset_repository "github.com/crisaltmann/fundament-stock-api/pkg/asset/repository"
	asset_service "github.com/crisaltmann/fundament-stock-api/pkg/asset/service"
	"github.com/crisaltmann/fundament-stock-api/pkg/holding/api"
	holding_repository "github.com/crisaltmann/fundament-stock-api/pkg/holding/repository"
	"github.com/crisaltmann/fundament-stock-api/pkg/holding/service"
	order_service "github.com/crisaltmann/fundament-stock-api/pkg/order/service"
	portfolio_service "github.com/crisaltmann/fundament-stock-api/pkg/portfolio/service"
	quarter_service "github.com/crisaltmann/fundament-stock-api/pkg/quarter/service"
	"go.uber.org/fx"
)

var Holding = fx.Options(
	holdingfactories,
	fx.Invoke(holding_api.MapRouter),
)

var holdingfactories = fx.Provide(
	func(repository asset_repository.Repository) holding_repository.AssetRepository {return repository},
	func(repository holding_repository.Repository) holding_service.Repository { return repository },
	holding_repository.NewRepository,

	func(assetService asset_service.Service) holding_service.AssetService { return assetService },
	func(portfolioService portfolio_service.Service) holding_service.PortfolioService { return portfolioService },
	func(quarterService quarter_service.Service) holding_service.QuarterService {return quarterService},
	func(orderService order_service.Service) holding_service.OrderService {return orderService},

	holding_service.NewService,
	func(service holding_service.Service) holding_api.Service {return service},

	holding_api.NewHandler,
)