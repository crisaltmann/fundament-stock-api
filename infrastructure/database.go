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
)

func CreateConnection(config *config.Config) *sql.DB {
	fmt.Println("conectando base ...")
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Errorf("Erro ao contectar no banco #{err}")
	}
	return db
}
