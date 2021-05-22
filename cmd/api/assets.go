package api

import (
	"github.com/crisaltmann/fundament-stock-api/pkg/asset/api"
	"github.com/crisaltmann/fundament-stock-api/pkg/asset/repository"
	"github.com/crisaltmann/fundament-stock-api/pkg/asset/service"

	"go.uber.org/fx"
)

var Asset = fx.Options(
	assetfactories,
	fx.Invoke(asset_api.MapRouter),
)

var assetfactories = fx.Provide(
	asset_repository.NewRepository,
	func(repository asset_repository.Repository) asset_service.Repository { return repository },

	asset_repository.NewStockPriceRepository,
	func(stockPriceRepository asset_repository.StockPriceRepository) asset_service.StockPriceRepository { return stockPriceRepository },

	asset_repository.NewAssetQuarterlyResultRepository,
	func(quarterlyResultRepository asset_repository.AssetQuarterlyResultRepository) asset_service.AssetQuarterlyResultRepository { return quarterlyResultRepository },

	asset_repository.NewQuarterlyResultProducer,

	asset_service.NewService,
	func(service asset_service.Service) asset_api.Service { return service },

	asset_api.NewHandler,
)
