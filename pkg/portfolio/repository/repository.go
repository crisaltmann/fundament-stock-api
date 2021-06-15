package portfolio_repository

import (
	"database/sql"
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/pkg/portfolio/domain"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{DB: db}
}

func (r Repository) GetPortfolio(usuario int64) ([]portfolio_domain.Portfolio, error) {
	rows, err := r.DB.Query("select a.id, a.codigo, a.logo, a.total, a.cotacao, sum(m.quantidade), m.id_usuario from movimentacao m " +
		"inner join ativo a on m.id_ativo = a.id  " +
		"where m.id_usuario = $1" +
		"group by a.id, a.codigo, a.logo, a.cotacao, m.id_usuario ", usuario)
	defer rows.Close()

	if err != nil {
		err = fmt.Errorf("Erro ao executar busca do portfolio", err)
		return nil, err
	}
	defer rows.Close()
	portfolio := []portfolio_domain.Portfolio{}
	for rows.Next() {
		item := portfolio_domain.Portfolio{}
		cotacao := sql.NullFloat64{}
		err := rows.Scan(&item.Ativo.Id, &item.Ativo.Codigo, &item.Ativo.Logo, &item.Ativo.Total, &cotacao, &item.Quantidade, &item.Usuario)
		if cotacao.Valid {
			item.Ativo.Cotacao = float32(cotacao.Float64)
		}
		if err != nil {
			err = fmt.Errorf("Erro ao executar busca do portfolio", err)
			return nil, err
		}
		portfolio = append(portfolio, item)
	}
	return portfolio, nil
}