package order_service

import (
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/pkg/order/domain"
	order_event "github.com/crisaltmann/fundament-stock-api/pkg/order/event"
)

type Service struct {
	repository    Repository
	assetFinder   AssetFinder
	producer      order_event.OrderProducer
}

type Repository interface {
	InsertOrder(order order_domain.Order) (bool, error)
	GetAllOrders() ([]order_domain.Order, error)
	GetUsersWithOrders(idAtivo int64) ([]int64, error)
}

type AssetFinder interface {
	ExistById(id int64) (bool, error)
}

func NewService(repository Repository, assetFinder AssetFinder, producer order_event.OrderProducer) Service {
	return Service{
		repository: repository,
		assetFinder: assetFinder,
		producer: producer,
	}
}

func (s Service) GetUsersWithOrders(idAtivo int64) ([]int64, error) {
	return s.repository.GetUsersWithOrders(idAtivo)
}

func (s Service) GetAllOrders() ([]order_domain.Order, error) {
	orders, err := s.repository.GetAllOrders()
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
	ativoExist, err := s.assetFinder.ExistById(order.Ativo)
	if err != nil && !ativoExist {
		err = fmt.Errorf("Ativo informado invalido ou não cadastrado")
	}
	order.Quantidade = adjustQtde(order)
	inserted, err := s.repository.InsertOrder(order)
	if inserted && err == nil {
		err := s.producer.PublishOrderEvent(order)
		if err != nil {
			return false, err
		}
	}
	return inserted, err
}

func adjustQtde(order order_domain.Order) int {
	if order.Tipo == order_domain.SellOrder {
		return -1 * order.Quantidade
	}
	return order.Quantidade
}