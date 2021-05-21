package api

import (
	"github.com/crisaltmann/fundament-stock-api/pkg/quarter/api"
	"github.com/crisaltmann/fundament-stock-api/pkg/quarter/repository"
	"github.com/crisaltmann/fundament-stock-api/pkg/quarter/service"
	"go.uber.org/fx"
)

var Quarter = fx.Options(
	quarterfactories,
	fx.Invoke(quarter_api.MapRouter),
)

var quarterfactories = fx.Provide(
	quarter_repository.NewRepository,
	func(repository quarter_repository.Repository) quarter_service.Repository { return repository},

	quarter_service.NewService,
	func(service quarter_service.Service) quarter_api.QuarterService { return service},

	quarter_api.NewHandler,
)