package repository

import (
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/asset/domain"
	"github.com/crisaltmann/fundament-stock-api/config"
	"github.com/crisaltmann/fundament-stock-api/infrastructure"
	"log"
)

type Repository struct {
	Config *config.Config
}

func (r Repository) GetAllAsset() ([]domain.Asset, error) {
	db := infrastructure.CreateConnection(r.Config)
	defer db.Close()

	rows, err := db.Query("select id, codigo, nome FROM ATIVO")
	if err != nil {
		err = fmt.Errorf("Erro ao executar busca de ativos", err)
		return nil, err
	}
	defer rows.Close()
	assets := []domain.Asset{}
	for rows.Next() {
		asset := domain.Asset{}
		err := rows.Scan(&asset.Id,
			&asset.Nome, &asset.Codigo)
		if err != nil {
			log.Fatal(err)
		}
		assets = append(assets, asset)
	}
	return assets, nil
}
