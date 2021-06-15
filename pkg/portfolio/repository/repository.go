package portfolio_repository

import (
	"database/sql"
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/pkg/portfolio/domain"
	"time"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{DB: db}
}

func (r Repository) GetPortfolio(usuario int64, dateFinal time.Time) ([]portfolio_domain.Portfolio, error) {
	query := "select a.id, a.codigo, a.logo, a.total, a.cotacao, sum(m.quantidade), m.id_usuario from movimentacao m " +
		" inner join ativo a on m.id_ativo = a.id  " +
		" where m.id_usuario = $1"

	if !dateFinal.Equal(time.Time{}) {
		query = query + " and m.data <= $2 "
	}

	query = query + " group by a.id, a.codigo, a.logo, a.cotacao, m.id_usuario "
	rows, err := r.DB.Query(query, usuario, dateFinal)
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