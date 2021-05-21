package app

import (
	"github.com/crisaltmann/fundament-stock-api/infrastructure"
	"go.uber.org/fx"
)

var Database = fx.Options(
	dbfactories,
)

var dbfactories = fx.Provide(
	infrastructure.CreateConnection,
)
