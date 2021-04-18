package main

import (
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/config"
	"github.com/crisaltmann/fundament-stock-api/server"
	_ "github.com/denisenkom/go-mssqldb"
	"go.uber.org/fx"
)

func main() {
	fmt.Println("Iniciando...")

	app := fx.New(
		fx.Provide(
			config.LoadConfig,
			server.ConfigureServer,
		),
		fx.Invoke(
			server.InitServer,
		),
	)
	app.Run()
}


