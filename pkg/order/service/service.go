package order_service

import (
	"fmt"
	asset_service2 "github.com/crisaltmann/fundament-stock-api/pkg/asset/service"
	order_domain2 "github.com/crisaltmann/fundament-stock-api/pkg/order/domain"
	order_repository2 "github.com/crisaltmann/fundament-stock-api/pkg/order/repository"
)

type Service struct {
	Repository *order_repository2.Repository
	AssetService *asset_service2.Service
}

func (s Service) GetAllOrders() ([]order_domain2.Order, error) {
	orders, err := s.Repository.GetAllOrders()
	if err != nil {
		return orders, err
	}
	for idx, order := range orders {
		if order.Quantidade >= 0 {
			orders[idx].Tipo = order_domain2.BuyOrder
		} else {
			orders[idx].Tipo = order_domain2.SellOrder
		}
	}
	return orders, nil
}

func (s Service) InsertOrder(order order_domain2.Order) (bool, error) {
	ativoExist, err := s.AssetService.ExistById(order.Ativo)
	if err != nil && !ativoExist {
		err = fmt.Errorf("Ativo informado invalido ou não cadastrado")
	}
	order.Quantidade = adjustQtde(order)
	return s.Repository.InsertOrder(order)
}

func adjustQtde(order order_domain2.Order) int {
	if order.Tipo == order_domain2.SellOrder {
		return -1 * order.Quantidade
	}
	return order.Quantidade
}