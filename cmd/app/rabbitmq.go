package app

import (
	"github.com/crisaltmann/fundament-stock-api/infrastructure"
	"go.uber.org/fx"
)

var RabbitMQ = fx.Options(
	rabbitfactories,
	fx.Invoke(infrastructure.ConfigureQueue),
)

var rabbitfactories = fx.Provide(
	infrastructure.CreateConnection,
	infrastructure.CreateRabbitMQChannel,
)

