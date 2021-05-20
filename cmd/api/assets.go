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

	asset_service.NewService,
	func(service asset_service.Service) asset_api.Service { return service },

	asset_api.NewHandler,
)
