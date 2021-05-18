package order_repository

import (
	"database/sql"
	"fmt"
	order_domain "github.com/crisaltmann/fundament-stock-api/order/domain"
)

type Repository struct {
	DB *sql.DB
}

func (r Repository) InsertOrder(order order_domain.Order) (bool, error) {

	prepare, err := r.DB.Prepare("INSERT INTO MOVIMENTACAO (ID_ATIVO, QUANTIDADE, VALOR) VALUES ($1, $2, $3)")

	if err != nil {
		err = fmt.Errorf("Erro ao executar insert de ativos", err)
		return false, err
	}

	_, err = prepare.Exec(order.Ativo, order.Quantidade, order.Valor)
	if err != nil {
		err = fmt.Errorf("Erro ao executar insert de movimentacao", err)
		return false, err
	}
	return true, nil
}