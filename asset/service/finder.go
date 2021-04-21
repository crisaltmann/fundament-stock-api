package service

import (
	"github.com/crisaltmann/fundament-stock-api/asset/domain"
	"github.com/crisaltmann/fundament-stock-api/asset/repository"
)

type Service struct {
	Repository *repository.Repository
}

func (s Service) GetAllAssets() ([]domain.Asset, error) {
	//return []Asset{
	//	Asset{1, "WEGE3", "WEG"},
	//	Asset{2, "ITUB3", "ITAÚ"},
	//}
	return s.Repository.GetAllAsset()
}

