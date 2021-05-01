package orderrepository

import (
	"database/sql"
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/asset/domain"
	orderdomain "github.com/crisaltmann/fundament-stock-api/portfolio/domain"
)

type Repository struct {
	DB *sql.DB
}

//func (r Repository) UpdateAsset(asset domain.Asset) (domain.Asset, error) {
//	prepare, err := r.DB.Prepare("UPDATE ATIVO SET CODIGO = @p2, NOME = @p3 WHERE ID = @p1")
//	if err != nil {
//		err = fmt.Errorf("Erro ao executar update de ativos", err)
//		return domain.Asset{}, err
//	}
//	defer prepare.Close()
//
//	_, err = prepare.Exec(asset.Id, asset.Codigo, asset.Nome)
//	if err != nil {
//		err = fmt.Errorf("Erro ao executar update de ativos", err)
//		return domain.Asset{}, err
//	}
//	return asset, nil
//}
//
func (r Repository) InsertOrder(order orderdomain.Order) (bool, error) {
	prepare, err := r.DB.Prepare("INSERT INTO ORDE (CODIGO, NOME) VALUES (@p1, @p2)")
	defer prepare.Close()

	if err != nil {
		err = fmt.Errorf("Erro ao executar insert de ordem", err)
		return false, err
	}

	_, err = prepare.Exec(asset.Codigo, asset.Nome)
	if err != nil {
		err = fmt.Errorf("Erro ao executar insert de ordem", err)
		return false, err
	}
	return true, nil
}
//
//func (r Repository) GetAllAsset() ([]domain.Asset, error) {
//	rows, err := r.DB.Query("select id, codigo, nome FROM ATIVO")
//	defer rows.Close()
//
//	if err != nil {
//		err = fmt.Errorf("Erro ao executar busca de ativos", err)
//		return nil, err
//	}
//	defer rows.Close()
//	assets := []domain.Asset{}
//	for rows.Next() {
//		asset := domain.Asset{}
//		err := rows.Scan(&asset.Id,
//			&asset.Nome, &asset.Codigo)
//		if err != nil {
//			err = fmt.Errorf("Erro ao executar busca de ativos", err)
//			return nil, err
//		}
//		assets = append(assets, asset)
//	}
//	return assets, nil
//}
//
//func (r Repository) GetById(id int64) (domain.Asset, error) {
//	rows, err := r.DB.Query("SELECT id, codigo, nome FROM ATIVO WHERE id = @p1", id)
//	defer rows.Close()
//
//	if err != nil {
//		err = fmt.Errorf("Erro ao executar busca de ativos por id", err)
//		return domain.Asset{}, err
//	}
//	defer rows.Close()
//	asset := domain.Asset{}
//	for rows.Next() {
//		err := rows.Scan(&asset.Id,
//			&asset.Nome, &asset.Codigo)
//		if err != nil {
//			err = fmt.Errorf("Erro ao executar busca de ativos por id", err)
//			return domain.Asset{}, err
//		}
//	}
//	return asset, nil
//}