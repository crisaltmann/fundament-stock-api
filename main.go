package main

import (
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/asset"
	"github.com/crisaltmann/fundament-stock-api/config"
	"github.com/crisaltmann/fundament-stock-api/server"
	_ "github.com/denisenkom/go-mssqldb"
	"go.uber.org/fx"
)

func main() {
	fmt.Println("Iniciando...")

	app := fx.New(
		config.Module,
		server.Module,
		asset.Module,
		fx.Invoke(
			server.InitServer,
		),
	)
	app.Run()
}
