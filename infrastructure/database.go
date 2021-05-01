package infrastructure

import (
	"database/sql"
	"github.com/crisaltmann/fundament-stock-api/config"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func CreateConnection(config *config.Config) *sql.DB {
	log.Println("conectando base ...")

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal("Erro ao contectar no banco", err)
		panic(err)
	}
	log.Println("Connected")
	return db
}
