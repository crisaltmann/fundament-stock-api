package holding_repository

import (
	"database/sql"
	"errors"
	"fmt"
	asset_domain "github.com/crisaltmann/fundament-stock-api/pkg/asset/domain"
	holding_domain "github.com/crisaltmann/fundament-stock-api/pkg/holding/domain"
)

type Repository struct {
	DB *sql.DB
	assetRepository AssetRepository
}

type AssetRepository interface {
	GetById(id int64) (asset_domain.Asset, error)
}

func NewRepository(db *sql.DB, assetRepository AssetRepository) Repository {
	return Repository{
		DB: db,
		assetRepository: assetRepository,
	}
}

func (r Repository) GetResultadoPortfolio(usuario string) ([]holding_domain.HoldingAtivo, error) {
	rows, err := r.DB.Query("SELECT id, id_trimestre, id_usuario, id_ativo, receita_liquida, ebitda, lucro_liquido, divida_liquida " +
		" FROM portfolio_trimestre WHERE id_usuario = $1 " +
		" ORDER BY id_trimestre ASC", usuario)
	defer rows.Close()

	if err != nil {
		err = fmt.Errorf("Erro ao executar busca do resultado de portfolio", err)
		return nil, err
	}
	defer rows.Close()
	resultPortfolio := []holding_domain.HoldingAtivo{}
	for rows.Next() {
		item := holding_domain.HoldingAtivo{}
		var idAtivo int64
		err := rows.Scan(&item.Id, &item.Trimestre, &idAtivo, &item.ReceitaLiquida, &item.Ebitda, &item.LucroLiquido,
			&item.DividaLiquida)
		if err != nil {
			err = fmt.Errorf("Erro ao executar busca do resultado de portfolio", err)
			return nil, err
		}
		ativo, err := r.assetRepository.GetById(idAtivo)
		if err != nil {
			err = errors.New("Erro ao executar busca de ativo no resultado de portfolio")
			return nil, err
		}
		item.Ativo = ativo
		resultPortfolio = append(resultPortfolio, item)
	}
	return resultPortfolio, nil
}