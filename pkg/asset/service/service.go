package asset_service

import (
	asset_domain2 "github.com/crisaltmann/fundament-stock-api/pkg/asset/domain"
	asset_repository2 "github.com/crisaltmann/fundament-stock-api/pkg/asset/repository"
)

type Service struct {
	Repository *asset_repository2.Repository
}

func (s Service) GetAllAssets() ([]asset_domain2.Asset, error) {
	return s.Repository.GetAllAsset()
}

func (s Service) ExistById(id int64) (bool, error) {
	return s.Repository.ExistById(id)
}

func (s Service) GetById(id int64) (asset_domain2.Asset, error) {
	return s.Repository.GetById(id)
}

func (s Service) InsertAsset(asset asset_domain2.Asset) (bool, error) {
	return s.Repository.InsertAsset(asset)
}

func (s Service) UpdateAsset(asset asset_domain2.Asset) (asset_domain2.Asset, error) {
	return s.Repository.UpdateAsset(asset)
}
