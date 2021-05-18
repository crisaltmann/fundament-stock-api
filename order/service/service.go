package order_service

import (
	"fmt"
	asset_service "github.com/crisaltmann/fundament-stock-api/asset/service"
	order_domain "github.com/crisaltmann/fundament-stock-api/order/domain"
	order_repository "github.com/crisaltmann/fundament-stock-api/order/repository"
)

type Service struct {
	Repository *order_repository.Repository
	AssetService *asset_service.Service
}

func (s Service) InsertOrder(order order_domain.Order) (bool, error) {
	ativoExist, err := s.AssetService.ExistById(order.Ativo)
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