package asset_service

import (
	asset_domain "github.com/crisaltmann/fundament-stock-api/asset/domain"
	asset_repository "github.com/crisaltmann/fundament-stock-api/asset/repository"
)

type Service struct {
	Repository *asset_repository.Repository
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
