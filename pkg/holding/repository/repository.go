package holding_repository

import (
	"database/sql"
	"errors"
	"fmt"
	asset_domain "github.com/crisaltmann/fundament-stock-api/pkg/asset/domain"
	holding_domain "github.com/crisaltmann/fundament-stock-api/pkg/holding/domain"
	"github.com/rs/zerolog/log"
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

func (r Repository) DeleteByAtivoAndTrimestre(idAtivo int64, idTrimestre int64) error {
	prepare, err := r.DB.Prepare("DELETE FROM portfolio_trimestre WHERE id_ativo = $1 AND id_trimestre = $2")

	if err != nil {
		log.Print("Ocorreu um erro ao preparar query de delete de portfolio trimestre")
		return err
	}

	defer prepare.Close()

	_, err = prepare.Exec(idAtivo, idTrimestre)
	if err != nil {
		err = fmt.Errorf("Erro ao executar delete portfolio trimestre", err)
		return err
	}
	return nil
}

func (r Repository) DeleteByUser(idUser int64) error {
	prepare, err := r.DB.Prepare("DELETE FROM portfolio_trimestre WHERE id_usuario = $1")

	if err != nil {
		log.Print("Ocorreu um erro ao preparar query de delete de portfolio trimestre por user")
		return err
	}

	defer prepare.Close()

	_, err = prepare.Exec(idUser)
	if err != nil {
		err = fmt.Errorf("Erro ao executar delete portfolio trimestre por user", err)
		return err
	}
	return nil
}

func (r Repository) SaveResultadoPortfolio(ativo holding_domain.HoldingAtivo) error {
	prepare, err := r.DB.Prepare("INSERT INTO portfolio_trimestre (id_trimestre, id_usuario, id_ativo, " +
		" receita_liquida, ebitda, lucro_liquido, divida_liquida) VALUES($1, $2, $3, $4, $5, $6, $7);")

	if err != nil {
		log.Print("Ocorreu um erro ao preparar query de insert de portfolio trimestre")
		return err
	}

	defer prepare.Close()

	_, err = prepare.Exec(ativo.Trimestre, ativo.Usuario, ativo.Ativo.Id, ativo.ReceitaLiquida, ativo.Ebitda,
		ativo.LucroLiquido, ativo.DividaLiquida)
	if err != nil {
		err = fmt.Errorf("Erro ao executar insert de portfolio trimestre", err)
		return err
	}
	return nil
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
		err := rows.Scan(&item.Id, &item.Trimestre, &item.Usuario, &idAtivo, &item.ReceitaLiquida, &item.Ebitda, &item.LucroLiquido,
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