package api

import (
	insight_api "github.com/crisaltmann/fundament-stock-api/pkg/insghts/api"
	insight_event "github.com/crisaltmann/fundament-stock-api/pkg/insghts/event"
	insight_repository "github.com/crisaltmann/fundament-stock-api/pkg/insghts/repository"
	insight_service "github.com/crisaltmann/fundament-stock-api/pkg/insghts/service"
	quarter_service "github.com/crisaltmann/fundament-stock-api/pkg/quarter/service"
	"go.uber.org/fx"
)

var Insight = fx.Options(
	insightfactories,
	fx.Invoke(insight_api.MapRouter),
	fx.Invoke(insight_repository.InitCache),
	fx.Invoke(insight_event.InitializeInsightConsume),
)

var insightfactories = fx.Provide(
	insight_repository.NewRepository,
	func(repository insight_repository.Repository) insight_service.Repository { return repository},

	insight_service.NewService,
	func(service insight_service.Service) insight_api.InsightService { return service},

	func(service insight_service.Service) insight_event.InsightService { return service },
	func(quarterService quarter_service.Service) insight_service.QuarterService { return quarterService },
	insight_event.NewInsightConsumer,

	insight_api.NewHandler,
)