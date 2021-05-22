package asset_service

import (
	"github.com/crisaltmann/fundament-stock-api/pkg/asset/domain"
	"log"
	"time"
)

type Service struct {
	Repository Repository
	StockPriceRepository StockPriceRepository
	AssetQuarterlyResultRepository AssetQuarterlyResultRepository
}

type Repository interface {
	UpdateAsset(asset asset_domain.Asset) (asset_domain.Asset, error)
	InsertAsset(asset asset_domain.Asset) (bool, error)
	GetAllAsset() ([]asset_domain.Asset, error)
	GetById(id int64) (asset_domain.Asset, error)
	ExistById(id int64) (bool, error)
	UpdateAssetPrice(id int64, price float32) (bool, error)
}

type StockPriceRepository interface {
	InsertAssetPrice(assetPrice asset_domain.AssetPrice) (bool, error)
	GetByAtivoEData(idAtivo int64, data time.Time) (asset_domain.AssetPrice, error)
}

type AssetQuarterlyResultRepository interface {
	InsertAssetQuarterlyResult(aqResult asset_domain.AssetQuarterlyResult) (bool, error)
	ExistAssetQuarterlyResult(idAtivo int64, idTrimestre int64) (bool, error)
}

func NewService(repository Repository, stockPriceRepository StockPriceRepository, assetQResultRepository AssetQuarterlyResultRepository) Service {
	return Service{
		Repository: repository,
		StockPriceRepository: stockPriceRepository,
		AssetQuarterlyResultRepository: assetQResultRepository,
	}
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

func (s Service) UpdateAssetPrice(id int64, price float32, data time.Time) (bool, error) {
	ok, err := s.Repository.UpdateAssetPrice(id, price)
	if err != nil {
		return false, err
	}
	s.InsertAssetPrice(id, price, data)
	return ok, nil
}

func (s Service) InsertAssetPrice(id int64, price float32, data time.Time) (bool, error) {
	assetPrice, err := s.StockPriceRepository.GetByAtivoEData(id, data)
	if err != nil {
		return false, err
	}
	if assetPrice.Id != 0 {
		return false, nil
	}
	updateAssetPrice := asset_domain.AssetPrice{
		Ativo:   id,
		Cotacao: price,
		Data:    data,
	}
	return s.StockPriceRepository.InsertAssetPrice(updateAssetPrice)
}

func (s Service) InsertAssetQuarterlyResult(aqResult asset_domain.AssetQuarterlyResult) (bool, error) {
	exist, err := s.AssetQuarterlyResultRepository.ExistAssetQuarterlyResult(aqResult.Ativo, aqResult.Trimestre)
	if exist {
		log.Println("Já existe um resultado cadastro para este trimestre.")
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return s.AssetQuarterlyResultRepository.InsertAssetQuarterlyResult(aqResult)
}