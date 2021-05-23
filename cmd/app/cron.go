package app

import (
	"github.com/crisaltmann/fundament-stock-api/infrastructure"

	"go.uber.org/fx"
)

var CronModule = fx.Options(
	cronfactories,
	fx.Invoke(infrastructure.InitCron),
)

var cronfactories = fx.Provide(
	infrastructure.NewCron,
)