package job

import (
	"github.com/crisaltmann/fundament-stock-api/pkg/asset-sync"
	"github.com/crisaltmann/fundament-stock-api/pkg/asset-sync/alphavantage"
	asset_sync_api "github.com/crisaltmann/fundament-stock-api/pkg/asset-sync/api"
	"github.com/crisaltmann/fundament-stock-api/pkg/asset/service"
	"go.uber.org/fx"
)

var AssetSync = fx.Options(
	factories,
	fx.Invoke(asset_sync.ConfigureJob),
	fx.Invoke(asset_sync_api.MapRouter),
)

var factories = fx.Provide(
	asset_sync.NewService,
	func(assetService asset_service.Service) asset_sync.AssetFinder { return assetService },
	func(assetService asset_service.Service) asset_sync.AssetUpdater { return assetService },
	func(client alphavantage.Client) asset_sync.StockPriceFinder { return client },
	asset_sync.NewAssetSync,
	alphavantage.NewAlphaVantageClient,

	asset_sync_api.NewHandler,
)