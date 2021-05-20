package asset_service

import (
	"github.com/crisaltmann/fundament-stock-api/pkg/asset/domain"
)

type Service struct {
	Repository Repository
}

type Repository interface {
	UpdateAsset(asset asset_domain.Asset) (asset_domain.Asset, error)
	InsertAsset(asset asset_domain.Asset) (bool, error)
	GetAllAsset() ([]asset_domain.Asset, error)
	GetById(id int64) (asset_domain.Asset, error)
	ExistById(id int64) (bool, error)
	UpdateAssetPrice(id int64, price float32) (bool, error)
}

func NewService(repository Repository) Service {
	return Service{Repository: repository}
}

func (s Service) GetAllAssets() ([]asset_domain.Asset, error) {
	return s.Repository.GetAllAsset()
}

func (s Service) ExistById(id int64) (bool, error) {
	return s.Repository.ExistById(id)
}

func (s Service) GetById(id int64) (asset_domain.Asset, error) {
	return s.Repository.GetById(id)
}

func (s Service) InsertAsset(asset asset_domain.Asset) (bool, error) {
	return s.Repository.InsertAsset(asset)
}

func (s Service) UpdateAsset(asset asset_domain.Asset) (asset_domain.Asset, error) {
	return s.Repository.UpdateAsset(asset)
}

func (s Service) UpdateAssetPrice(id int64, price float32) (bool, error) {
	return s.Repository.UpdateAssetPrice(id, price)
}
