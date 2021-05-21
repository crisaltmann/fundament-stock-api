package quarter_repository

import (
	"database/sql"
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/pkg/quarter/domain"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{DB: db}
}

func (r Repository) GetQuarter(id int64) (quarter_domain.Trimestre, error) {
	rows, err := r.DB.Query("SELECT id, codigo, ano, trimestre, data_inicio, data_fim FROM trimestre WHERE id = $1", id)
	defer rows.Close()

	if err != nil {
		err = fmt.Errorf("Erro ao executar busca de trimestre", err)
		return quarter_domain.Trimestre{}, err
	}
	defer rows.Close()
	trimestre := quarter_domain.Trimestre{}
	for rows.Next() {
		err := rows.Scan(&trimestre.Id, &trimestre.Codigo, &trimestre.Ano, &trimestre.Trimestre, &trimestre.DataInicio, &trimestre.DataFim)
		if err != nil {
			err = fmt.Errorf("Erro ao executar busca do trimestre", err)
			return quarter_domain.Trimestre{}, err
		}
	}
	return trimestre, nil
}

func (r Repository) GetQuarters() ([]quarter_domain.Trimestre, error) {
	rows, err := r.DB.Query("SELECT id, codigo, ano, trimestre, data_inicio, data_fim FROM trimestre")
	defer rows.Close()

	if err != nil {
		err = fmt.Errorf("Erro ao executar busca de trimestres", err)
		return nil, err
	}
	defer rows.Close()
	trimestres := []quarter_domain.Trimestre{}
	for rows.Next() {
		trimestre := quarter_domain.Trimestre{}
		err := rows.Scan(&trimestre.Id, &trimestre.Codigo, &trimestre.Ano, &trimestre.Trimestre, &trimestre.DataInicio, &trimestre.DataFim)
		if err != nil {
			err = fmt.Errorf("Erro ao executar busca do trimestres", err)
			return nil, err
		}
		trimestres = append(trimestres, trimestre)
	}
	return trimestres, nil
}