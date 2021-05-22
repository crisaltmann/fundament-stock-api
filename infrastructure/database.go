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

	url := os.Getenv("DATABASE_URL")
	if len(url) == 0 {
		url = config.Url
	}
	db, err := sql.Open("postgres", url)
	defer db.Close()
	if err != nil {
		log.Fatal("Erro ao contectar no banco", err)
		panic(err)
	}
	log.Println("Connected")
	return db
}
