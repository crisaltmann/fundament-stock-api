package api

import (
	insight_api "github.com/crisaltmann/fundament-stock-api/pkg/insghts/api"
	insight_repository "github.com/crisaltmann/fundament-stock-api/pkg/insghts/repository"
	insight_service "github.com/crisaltmann/fundament-stock-api/pkg/insghts/service"
	"go.uber.org/fx"
)

var Insight = fx.Options(
	insightfactories,
	fx.Invoke(insight_api.MapRouter),
	fx.Invoke(insight_repository.InitCache),
)

var insightfactories = fx.Provide(
	insight_repository.NewRepository,
	func(repository insight_repository.Repository) insight_service.Repository { return repository},

	insight_service.NewService,
	func(service insight_service.Service) insight_api.InsightService { return service},

	insight_api.NewHandler,
)