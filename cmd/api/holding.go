package api

import (
	"github.com/crisaltmann/fundament-stock-api/pkg/holding/api"
	"github.com/crisaltmann/fundament-stock-api/pkg/holding/service"
	"go.uber.org/fx"
)

var Holding = fx.Options(
	holdingfactories,
	fx.Invoke(holding_api.MapRouter),
)

var holdingfactories = fx.Provide(
	//order_repository.NewRepository,
	//func(repository order_repository.Repository) order_service.Repository { return repository },

	holding_service.NewService,
	//func(service order_service.Service) order_api.Service { return service },
	//func(assetService asset_service.Service) order_service.AssetFinder { return assetService },
	//func(service portfolio_service.Service) portfolio_api.Service { return service },
	func(service holding_service.Service) holding_api.Service {return service},

	holding_api.NewHandler,
)