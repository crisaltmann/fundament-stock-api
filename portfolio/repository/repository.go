package portfolio_repository

import (
	"database/sql"
	"fmt"
	portfolio_domain "github.com/crisaltmann/fundament-stock-api/portfolio/domain"
)

type Repository struct {
	DB *sql.DB
}

func (r Repository) GetPortfolio(usuario string) ([]portfolio_domain.Portfolio, error) {
	rows, err := r.DB.Query("select a.id, a.codigo, a.logo, sum(m.quantidade), m.id_usuario from movimentacao m " +
		"inner join ativo a on m.id_ativo = a.id  " +
		"where m.id_usuario = $1" +
		"group by a.id, a.codigo, a.logo, m.id_usuario ", usuario)
	defer rows.Close()

	if err != nil {
		err = fmt.Errorf("Erro ao executar busca do portfolio", err)
		return nil, err
	}
	defer rows.Close()
	portfolio := []portfolio_domain.Portfolio{}
	for rows.Next() {
		item := portfolio_domain.Portfolio{}
		err := rows.Scan(&item.Ativo.Id, &item.Ativo.Codigo, &item.Ativo.Logo, &item.Quantidade, &item.Usuario)
		if err != nil {
			err = fmt.Errorf("Erro ao executar busca do portfolio", err)
			return nil, err
		}
		portfolio = append(portfolio, item)
	}
	return portfolio, nil
}