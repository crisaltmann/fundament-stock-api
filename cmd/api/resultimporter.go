package api

import (
	asset_service "github.com/crisaltmann/fundament-stock-api/pkg/asset/service"
	quarter_service "github.com/crisaltmann/fundament-stock-api/pkg/quarter/service"
	result_importer_api "github.com/crisaltmann/fundament-stock-api/pkg/result_importer/api"
	result_importer_service "github.com/crisaltmann/fundament-stock-api/pkg/result_importer/service"
	"go.uber.org/fx"
)

var ResultImporter = fx.Options(
	importerfactories,
	fx.Invoke(result_importer_api.MapRouter),
)

var importerfactories = fx.Provide(
	func (quarterService quarter_service.Service) result_importer_service.QuarterService { return quarterService },
	func (assetService asset_service.Service) result_importer_service.AssetService { return assetService },
	func (service asset_service.Service) result_importer_service.QuarterlyResultService { return service },
	result_importer_service.NewImporter,

	result_importer_api.NewHandler,
)
