package api

import (
	user_api "github.com/crisaltmann/fundament-stock-api/pkg/user/api"
	user_repository "github.com/crisaltmann/fundament-stock-api/pkg/user/repository"
	user_service "github.com/crisaltmann/fundament-stock-api/pkg/user/service"
	"go.uber.org/fx"
)

var User = fx.Options(
	userfactories,
	fx.Invoke(user_api.MapRouter),
)

var userfactories = fx.Provide(
	user_repository.NewRepository,
	func(repository user_repository.Repository) user_service.Repository { return repository },

	user_service.NewService,
	func(service user_service.Service) user_api.Service { return service },

	user_api.NewHandler,
)
