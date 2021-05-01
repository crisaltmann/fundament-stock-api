package main

import (
	"github.com/crisaltmann/fundament-stock-api/asset"
	"github.com/crisaltmann/fundament-stock-api/config"
	"github.com/crisaltmann/fundament-stock-api/infrastructure"
	"github.com/crisaltmann/fundament-stock-api/server"
	"go.uber.org/fx"
	"log"
)

func main() {
	log.Println("Iniciando...")

	app := fx.New(
		infrastructure.Module,
		config.Module,
		server.Module,
		asset.Module,
		fx.Invoke(
			server.InitServer,
		),
	)
	app.Run()
}
