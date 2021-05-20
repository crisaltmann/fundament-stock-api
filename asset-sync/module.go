package asset_sync

import (
	"github.com/crisaltmann/fundament-stock-api/asset-sync/alphavantage"
	asset_service "github.com/crisaltmann/fundament-stock-api/pkg/asset/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
	fx.Invoke(configureJob),
)

var factories = fx.Provide(
	newService,
	newAssetSync,
	newAlphaVantageClient,
)

func newAssetSync(service *Service) *AssetSync {
	return &AssetSync{
		Service: service,
	}
}

func newService(assetService *asset_service.Service, client *alphavantage.Client) *Service {
	return &Service{
		AssetService: assetService,
		Client: client,
	}
}

func newAlphaVantageClient() *alphavantage.Client {
	return &alphavantage.Client{}
}