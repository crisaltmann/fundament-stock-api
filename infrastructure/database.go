package infrastructure

import (
	"database/sql"
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/config"
	"go.uber.org/fx"
	"os"
)

var Module = fx.Options(
	factories,
)

var factories = fx.Provide(
	//LoadDatabase,
)

func CreateConnection(config *config.Config) *sql.DB {
	fmt.Println("conectando base ...")
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))


	//connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;port=%d",
	//	config.Hostname, config.User, config.Password, config.Scheme, config.Port)
	//fmt.Println("init connection")
	//db, err := sql.Open("sqlserver", connString)

	if err != nil {
		fmt.Errorf("Erro ao contectar no banco #{err}")

	}
	fmt.Println("Connected")
	//defer db.Close()
	//
	//sqlCreate := "CREATE TABLE ATIVO (id bigint PRIMARY KEY, nome varchar(60))"
	//_, err = db.ExecContext(ctx, sqlCreate)
	//if err != nil {
	//	fmt.Println("err")
	//	fmt.Errorf("Erro ao criar tabela", err)
	//}
	return db
}
