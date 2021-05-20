package quarter

import (
	"database/sql"
	"github.com/crisaltmann/fundament-stock-api/pkg/quarter/api"
	"github.com/crisaltmann/fundament-stock-api/pkg/quarter/repository"
	"github.com/crisaltmann/fundament-stock-api/pkg/quarter/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	factories,
	fx.Invoke(quarter_api.MapRouter),
)

var factories = fx.Provide(
	NewRepository,
	NewService,
	NewHandler,
)

func NewRepository(db *sql.DB) *quarter_repository.Repository {
	return &quarter_repository.Repository{DB: db}
}

func NewService(repository *quarter_repository.Repository) *quarter_service.Service {
	return &quarter_service.Service{Repository: repository}
}

func NewHandler(service *quarter_service.Service) *quarter_api.Handler {
	return &quarter_api.Handler{Service: service}
}
