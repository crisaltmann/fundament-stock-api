package asset_repository

import (
	"database/sql"
	"fmt"
	asset_domain "github.com/crisaltmann/fundament-stock-api/asset/domain"
)

type Repository struct {
	DB *sql.DB
}

func (r Repository) UpdateAsset(asset asset_domain.Asset) (asset_domain.Asset, error) {
	prepare, err := r.DB.Prepare("UPDATE ATIVO SET CODIGO = $2, NOME = $3 WHERE ID = $1")

	if err != nil {
		err = fmt.Errorf("Erro ao executar update de ativos", err)
		return asset_domain.Asset{}, err
	}
	defer prepare.Close()

	_, err = prepare.Exec(asset.Id, asset.Codigo, asset.Nome)
	if err != nil {
		err = fmt.Errorf("Erro ao executar update de ativos", err)
		return asset_domain.Asset{}, err
	}
	return asset, nil
}

func (r Repository) InsertAsset(asset asset_domain.Asset) (bool, error) {

	prepare, err := r.DB.Prepare("INSERT INTO ATIVO (CODIGO, NOME) VALUES ($1, $2)")

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

func (r Repository) GetAllAsset() ([]asset_domain.Asset, error) {
	rows, err := r.DB.Query("select id, codigo, nome, logo FROM ATIVO")
	defer rows.Close()

	if err != nil {
		err = fmt.Errorf("Erro ao executar busca de ativos", err)
		return nil, err
	}
	defer rows.Close()
	assets := []asset_domain.Asset{}
	for rows.Next() {
		asset := asset_domain.Asset{}
		err := rows.Scan(&asset.Id,
			&asset.Nome, &asset.Codigo, &asset.Logo)
		if err != nil {
			err = fmt.Errorf("Erro ao executar busca de ativos", err)
			return nil, err
		}
		assets = append(assets, asset)
	}
	return assets, nil
}

func (r Repository) GetById(id int64) (asset_domain.Asset, error) {
	rows, err := r.DB.Query("SELECT id, codigo, nome FROM ATIVO WHERE id = $1", id)
	defer rows.Close()

	if err != nil {
		err = fmt.Errorf("Erro ao executar busca de ativos por id", err)
		return asset_domain.Asset{}, err
	}
	defer rows.Close()
	asset := asset_domain.Asset{}
	for rows.Next() {
		err := rows.Scan(&asset.Id,
			&asset.Nome, &asset.Codigo)
		if err != nil {
			err = fmt.Errorf("Erro ao executar busca de ativos por id", err)
			return asset_domain.Asset{}, err
		}
	}
	return asset, nil
}

func (r Repository) ExistById(id int64) (bool, error) {
	rows, err := r.DB.Query("SELECT count(*) FROM ATIVO WHERE id = $1", id)
	defer rows.Close()

	if err != nil {
		err = fmt.Errorf("Erro ao executar busca de exist ativos por id", err)
		return false, err
	}
	defer rows.Close()
	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			err = fmt.Errorf("Erro ao executar busca de exist ativos por id", err)
			return false, err
		}
	}
	return count > 0, nil
}
