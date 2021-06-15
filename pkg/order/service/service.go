package order_service

import (
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/pkg/order/domain"
)

type Service struct {
	Repository Repository
	AssetFinder AssetFinder
}

type Repository interface {
	InsertOrder(order order_domain.Order) (bool, error)
	GetAllOrders() ([]order_domain.Order, error)
	GetUsersWithOrders(idAtivo int64) ([]int64, error)
}

type AssetFinder interface {
	ExistById(id int64) (bool, error)
}

func NewService(repository Repository, assetFinder AssetFinder) Service {
	return Service{Repository: repository, AssetFinder: assetFinder}
}

func (s Service) GetUsersWithOrders(idAtivo int64) ([]int64, error) {
	return s.Repository.GetUsersWithOrders(idAtivo)
}

func (s Service) GetAllOrders() ([]order_domain.Order, error) {
	orders, err := s.Repository.GetAllOrders()
	if err != nil {
		return orders, err
	}
	for idx, order := range orders {
		if order.Quantidade >= 0 {
			orders[idx].Tipo = order_domain.BuyOrder
		} else {
			orders[idx].Tipo = order_domain.SellOrder
		}
	}
	return orders, nil
}

func (s Service) InsertOrder(order order_domain.Order) (bool, error) {
	ativoExist, err := s.AssetFinder.ExistById(order.Ativo)
	if err != nil && !ativoExist {
		err = fmt.Errorf("Ativo informado invalido ou não cadastrado")
	}
	order.Quantidade = adjustQtde(order)
	return s.Repository.InsertOrder(order)
}

func adjustQtde(order order_domain.Order) int {
	if order.Tipo == order_domain.SellOrder {
		return -1 * order.Quantidade
	}
	return order.Quantidade
}