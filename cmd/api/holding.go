package api

import (
	asset_service "github.com/crisaltmann/fundament-stock-api/pkg/asset/service"
	"github.com/crisaltmann/fundament-stock-api/pkg/holding/api"
	"github.com/crisaltmann/fundament-stock-api/pkg/holding/service"
	portfolio_service "github.com/crisaltmann/fundament-stock-api/pkg/portfolio/service"
	"go.uber.org/fx"
)

var Holding = fx.Options(
	holdingfactories,
	fx.Invoke(holding_api.MapRouter),
)

var holdingfactories = fx.Provide(
	func(assetService asset_service.Service) holding_service.AssetService { return assetService },
	func(portfolioService portfolio_service.Service) holding_service.PortfolioService { return portfolioService },

	holding_service.NewService,
	func(service holding_service.Service) holding_api.Service {return service},

	holding_api.NewHandler,
)