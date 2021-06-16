package asset_repository

import (
	"database/sql"
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/pkg/asset/domain"
)

type AssetQuarterlyResultRepository struct {
	DB *sql.DB
}

func NewAssetQuarterlyResultRepository(db *sql.DB) AssetQuarterlyResultRepository {
	return AssetQuarterlyResultRepository{DB: db}
}

func (r AssetQuarterlyResultRepository) ExistAssetQuarterlyResult(idAtivo int64, idTrimestre int64) (bool, error) {
	rows, err := r.DB.Query("SELECT count(*) FROM resultado_trimestre WHERE id_ativo = $1 AND id_trimestre = $2", idAtivo, idTrimestre)
	defer rows.Close()

	if err != nil {
		err = fmt.Errorf("Erro ao executar busca de exist para resultados de empresas", err)
		return false, err
	}
	defer rows.Close()
	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			err = fmt.Errorf("Erro ao executar busca de exist para resultados de empresas", err)
			return false, err
		}
	}
	return count > 0, nil
}

func (r AssetQuarterlyResultRepository) InsertAssetQuarterlyResult(aqResult asset_domain.AssetQuarterlyResult) (bool, error) {
	prepare, err := r.DB.Prepare("INSERT INTO RESULTADO_TRIMESTRE (id_trimestre, id_ativo, receita_liquida, ebitda, lucro_liquido, divida_liquida) " +
		"VALUES ($1, $2, $3, $4, $5, $6)")

	if err != nil {
		err = fmt.Errorf("Erro ao executar insert de resultado trimestral de ativo", err)
		return false, err
	}

	_, err = prepare.Exec(aqResult.Trimestre, aqResult.Ativo, aqResult.ReceitaLiquida, aqResult.Ebitda, aqResult.LucroLiquido, aqResult.DividaLiquida)
	if err != nil {
		err = fmt.Errorf("Erro ao executar insert de resultado trimestral de ativo", err)
		return false, err
	}
	return true, nil
}

func (r AssetQuarterlyResultRepository) GetAssetQuarterlyResults(idAtivo int64, idTrimestre int64) ([]asset_domain.AssetQuarterlyResult, error) {
	query := "SELECT id, id_trimestre, id_ativo, receita_liquida, ebitda, lucro_liquido, divida_liquida " +
		" FROM resultado_trimestre WHERE id_ativo = $1 "

	if idTrimestre > 0 {
		query = query + " AND id_trimestre = $2 "
	}

	rows, err := r.DB.Query(query, idAtivo, idTrimestre)
	defer rows.Close()

	if err != nil {
		err = fmt.Errorf("Erro ao executar busca de resultados de empresas", err)
		return nil, err
	}
	defer rows.Close()
	quarterlyResults := []asset_domain.AssetQuarterlyResult{}
	for rows.Next() {
		quarter := asset_domain.AssetQuarterlyResult{}
		err := rows.Scan(&quarter.Id, &quarter.Trimestre, &quarter.Ativo, &quarter.ReceitaLiquida, &quarter.Ebitda,
			&quarter.LucroLiquido, &quarter.DividaLiquida)
		if err != nil {
			err = fmt.Errorf("Erro ao executar busca de resultados de empresas", err)
			return nil, err
		}
		quarterlyResults = append(quarterlyResults, quarter)
	}
	return quarterlyResults, nil
}