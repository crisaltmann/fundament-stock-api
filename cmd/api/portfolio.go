package api

import (
	"github.com/crisaltmann/fundament-stock-api/pkg/portfolio/api"
	"github.com/crisaltmann/fundament-stock-api/pkg/portfolio/repository"
	"github.com/crisaltmann/fundament-stock-api/pkg/portfolio/service"
	quarter_service "github.com/crisaltmann/fundament-stock-api/pkg/quarter/service"
	"go.uber.org/fx"
)

var Portfolio = fx.Options(
	portfoliofactories,
	fx.Invoke(portfolio_api.MapRouter),
)

var portfoliofactories = fx.Provide(
	portfolio_repository.NewRepository,
	func(repository portfolio_repository.Repository) portfolio_service.Repository { return repository },
	func(quarterService quarter_service.Service) portfolio_service.QuarterService { return quarterService },

	portfolio_service.NewService,
	func(service portfolio_service.Service) portfolio_api.Service { return service },

	portfolio_api.NewHandler,
)




