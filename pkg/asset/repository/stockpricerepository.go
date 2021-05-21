package asset_repository

import (
	"database/sql"
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/pkg/asset/domain"
	"time"
)

type StockPriceRepository struct {
	DB *sql.DB
}

func NewStockPriceRepository(db *sql.DB) StockPriceRepository {
	return StockPriceRepository{DB: db}
}

func (r StockPriceRepository) InsertAssetPrice(assetPrice asset_domain.AssetPrice) (bool, error) {
	prepare, err := r.DB.Prepare("INSERT INTO COTACAO_ATIVO (id_ativo, cotacao, data) VALUES ($1, $2, $3)")

	if err != nil {
		err = fmt.Errorf("Erro ao executar insert de cotação de ativo", err)
		return false, err
	}

	_, err = prepare.Exec(assetPrice.Ativo, assetPrice.Cotacao, assetPrice.Data)
	if err != nil {
		err = fmt.Errorf("Erro ao executar insert de ativos", err)
		return false, err
	}
	return true, nil
}

func (r StockPriceRepository) GetByAtivoEData(idAtivo int64, data time.Time) (asset_domain.AssetPrice, error) {
	rows, err := r.DB.Query("SELECT id, id_Ativo, data, cotacao FROM cotacao_ativo WHERE id_ativo = $1 AND data = $2", idAtivo, data)
	defer rows.Close()

	if err != nil {
		err = fmt.Errorf("Erro ao executar busca de cotação de ativos por id ativo e data", err)
		return asset_domain.AssetPrice{}, err
	}
	defer rows.Close()
	asset := asset_domain.AssetPrice{}
	for rows.Next() {
		err := rows.Scan(&asset.Id, &asset.Ativo, &asset.Data, &asset.Cotacao)
		if err != nil {
			err = fmt.Errorf("Erro ao executar busca de cotação de ativos por id ativo e data", err)
			return asset_domain.AssetPrice{}, err
		}
	}
	return asset, nil
}