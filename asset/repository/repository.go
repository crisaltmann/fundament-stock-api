package repository

import (
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/asset/domain"
	"github.com/crisaltmann/fundament-stock-api/config"
	"github.com/crisaltmann/fundament-stock-api/infrastructure"
)

type Repository struct {
	Config *config.Config
}

func (r Repository) UpdateAsset(asset domain.Asset) (domain.Asset, error) {
	db := infrastructure.CreateConnection(r.Config)
	defer db.Close()

	prepare, err := db.Prepare("UPDATE ATIVO SET CODIGO = @p2, NOME = @p3 WHERE ID = @p1")
	if err != nil {
		err = fmt.Errorf("Erro ao executar update de ativos", err)
		return domain.Asset{}, err
	}

	_, err = prepare.Exec(asset.Id, asset.Codigo, asset.Nome)
	if err != nil {
		err = fmt.Errorf("Erro ao executar update de ativos", err)
		return domain.Asset{}, err
	}
	return asset, nil
}

func (r Repository) InsertAsset(asset domain.Asset) (bool, error) {
	db := infrastructure.CreateConnection(r.Config)
	defer db.Close()

	prepare, err := db.Prepare("INSERT INTO ATIVO (CODIGO, NOME) VALUES (@p1, @p2)")
	if err != nil {
		err = fmt.Errorf("Erro ao executar insert de ativos", err)
		return false, err
	}

	_, err = prepare.Exec(asset.Codigo, asset.Nome)
	if err != nil {
		err = fmt.Errorf("Erro ao executar insert de ativos", err)
		return false, err
	}
	return true, nil
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
			err = fmt.Errorf("Erro ao executar busca de ativos", err)
			return nil, err
		}
		assets = append(assets, asset)
	}
	return assets, nil
}

func (r Repository) GetById(id int64) (domain.Asset, error) {
	db := infrastructure.CreateConnection(r.Config)
	defer db.Close()

	rows, err := db.Query("SELECT id, codigo, nome FROM ATIVO WHERE id = @p1", id)
	if err != nil {
		err = fmt.Errorf("Erro ao executar busca de ativos por id", err)
		return domain.Asset{}, err
	}
	defer rows.Close()
	asset := domain.Asset{}
	for rows.Next() {
		err := rows.Scan(&asset.Id,
			&asset.Nome, &asset.Codigo)
		if err != nil {
			err = fmt.Errorf("Erro ao executar busca de ativos por id", err)
			return domain.Asset{}, err
		}
	}
	return asset, nil
}
