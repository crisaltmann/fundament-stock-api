package main

import (
	"database/sql"
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/config"
	"github.com/crisaltmann/fundament-stock-api/server"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
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

func testDatabase(ctx *gin.Context) {
	scheme := "youasholding"
	username := "youasholding"
	password := "Test123%"
	hostname := "youasholding.database.windows.net"
	port := 1433

	fmt.Println("conectando...")

	//query := url.Values{}
	//query.Add("app name", "MyAppName")
	//
	//u := &url.URL{
	//	Scheme:   scheme,
	//	User:     url.UserPassword(username, password),
	//	Host:     fmt.Sprintf("%s:%d", hostname, port),
	//	// Path:  instance, // if connecting to an instance instead of a port
	//	RawQuery: query.Encode(),
	//}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;port=%d",
		hostname, username, password, scheme, port)
	fmt.Println("init connection")
	db, err := sql.Open("sqlserver", connString)

	if err != nil {
		fmt.Errorf("Erro ao contectar no banco #{err}")

	}
	fmt.Println("Connected")
	defer db.Close()

	sqlCreate := "CREATE TABLE ATIVO (id bigint PRIMARY KEY, nome varchar(60))"
	_, err = db.ExecContext(ctx, sqlCreate)
	if err != nil {
		fmt.Println("err")
		fmt.Errorf("Erro ao criar tabela", err)
	}
}
