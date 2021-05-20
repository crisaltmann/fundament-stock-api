package portfolio

import (
	"database/sql"
	portfolio_repository "github.com/crisaltmann/fundament-stock-api/portfolio/repository"
	portfolio_service "github.com/crisaltmann/fundament-stock-api/portfolio/service"

	portfolio_api "github.com/crisaltmann/fundament-stock-api/portfolio/api"

	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
	fx.Invoke(portfolio_api.MapRouter),
)

var factories = fx.Provide(
	NewRepository,
	NewService,
	NewHandler,
)

func NewRepository(db *sql.DB) *portfolio_repository.Repository {
	return &portfolio_repository.Repository{DB: db}
}

func NewService(repository *portfolio_repository.Repository) *portfolio_service.Service {
	return &portfolio_service.Service{Repository: repository}
}

func NewHandler(service *portfolio_service.Service) *portfolio_api.Handler {
	return &portfolio_api.Handler{Service: service}
}
