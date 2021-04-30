package orderservice

import (
	"github.com/crisaltmann/fundament-stock-api/asset/domain"
	orderdomain "github.com/crisaltmann/fundament-stock-api/portfolio/domain"
	orderrepository "github.com/crisaltmann/fundament-stock-api/portfolio/order/repository"
)

type Service struct {
	Repository *orderrepository.Repository
}

//func (s Service) GetAllAssets() ([]domain.Asset, error) {
//	return s.Repository.GetAllAsset()
//}
//
//func (s Service) GetById(id int64) (domain.Asset, error) {
//	return s.Repository.GetById(id)
//}
//
func (s Service) InsertOrder(order orderdomain.Order) (bool, error) {
	return s.Repository.InsertOrder(order)
}
//
//func (s Service) UpdateAsset(asset domain.Asset) (domain.Asset, error) {
//	return s.Repository.UpdateAsset(asset)
//}

