package main

import (
	"github.com/crisaltmann/fundament-stock-api/cmd/api"
	"github.com/crisaltmann/fundament-stock-api/cmd/app"
	"github.com/crisaltmann/fundament-stock-api/cmd/job"
	"github.com/crisaltmann/fundament-stock-api/server"
	"go.uber.org/fx"
	"log"
)

// @title Fundament Stock Api Swagger API
// @version 1.0
// @description Swagger API for Fundament Stock Api.
// @termsOfService http://swagger.io/terms/

// @contact.name Cristiano Altmann
// @contact.email crisaltmann@gmail.com

// @license.name MIT
// @license.url https://github.com/crisaltmann/fundament-stock-api

// @BasePath /
func main() {
	log.Println("Iniciando...")

	app := fx.New(
		app.Config,
		app.CronModule,
		app.Database,
		app.RabbitMQ,
		app.Server,
		api.Asset,
		api.Order,
		api.Portfolio,
		api.Quarter,
		api.Holding,
		job.AssetSync,
		fx.Invoke(
			server.InitServer,
		),
	)
	app.Run()
}