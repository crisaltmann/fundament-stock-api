package order_repository

import (
	"database/sql"
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/pkg/order/domain"
	"github.com/rs/zerolog/log"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{DB: db}
}

func (r Repository) GetUsersWithOrders(idAtivo int64) ([]int64, error) {
	rows, err := r.DB.Query("select id_usuario FROM MOVIMENTACAO WHERE id_ativo = $1 GROUP BY id_usuario", idAtivo)

	if err != nil {
		log.Print("Erro ao executar busca de usuarios de orders.")
		return nil, err
	}

	defer rows.Close()
	users := make([]int64, 0)
	for rows.Next() {
		user := int64(0)
		err := rows.Scan(&user)
		if err != nil {
			log.Print("Erro ao executar busca de usuarios de orders.")
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r Repository) InsertOrder(order order_domain.Order) (bool, error) {
	prepare, err := r.DB.Prepare("INSERT INTO MOVIMENTACAO (ID_ATIVO, QUANTIDADE, VALOR, DATA, ID_USUARIO) VALUES ($1, $2, $3, $4, $5)")

	if err != nil {
		err = fmt.Errorf("Erro ao executar insert de movimentacao", err)
		return false, err
	}

	_, err = prepare.Exec(order.Ativo, order.Quantidade, order.Valor, order.Data, order.Usuario)
	if err != nil {
		err = fmt.Errorf("Erro ao executar insert de movimentacao", err)
		return false, err
	}
	return true, nil
}

func (r Repository) GetAllOrders() ([]order_domain.Order, error) {
	rows, err := r.DB.Query("select id, id_ativo, quantidade, valor, data, id_usuario FROM MOVIMENTACAO")
	defer rows.Close()

	if err != nil {
		err = fmt.Errorf("Erro ao executar busca de movimentacoes", err)
		return nil, err
	}
	defer rows.Close()
	orders := []order_domain.Order{}
	for rows.Next() {
		order := order_domain.Order{}
		err := rows.Scan(&order.Id, &order.Ativo, &order.Quantidade, &order.Valor, &order.Data, &order.Usuario)
		if err != nil {
			err = fmt.Errorf("Erro ao executar busca de movimentacoes", err)
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}