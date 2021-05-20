package portfolio

import (
	"database/sql"
	portfolio_api2 "github.com/crisaltmann/fundament-stock-api/pkg/portfolio/api"
	portfolio_repository2 "github.com/crisaltmann/fundament-stock-api/pkg/portfolio/repository"
	portfolio_service2 "github.com/crisaltmann/fundament-stock-api/pkg/portfolio/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
	fx.Invoke(portfolio_api2.MapRouter),
)

var factories = fx.Provide(
	NewRepository,
	NewService,
	NewHandler,
)

func NewRepository(db *sql.DB) *portfolio_repository2.Repository {
	return &portfolio_repository2.Repository{DB: db}
}

func NewService(repository *portfolio_repository2.Repository) *portfolio_service2.Service {
	return &portfolio_service2.Service{Repository: repository}
}

func NewHandler(service *portfolio_service2.Service) *portfolio_api2.Handler {
	return &portfolio_api2.Handler{Service: service}
}
