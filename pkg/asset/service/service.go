package asset_service

import (
	"github.com/crisaltmann/fundament-stock-api/internal"
	"github.com/crisaltmann/fundament-stock-api/pkg/asset/domain"
	"github.com/crisaltmann/fundament-stock-api/pkg/asset/event"
	"github.com/rs/zerolog/log"
	"time"
)

type Service struct {
	repository                     Repository
	stockPriceRepository           StockPriceRepository
	assetQuarterlyResultRepository AssetQuarterlyResultRepository
	quarterlyProducer              event.QuarterlyResultProducer
}

type Repository interface {
	UpdateAsset(asset asset_domain.Asset) (asset_domain.Asset, error)
	InsertAsset(asset asset_domain.Asset) (bool, error)
	GetAllAsset() ([]asset_domain.Asset, error)
	GetById(id int64) (asset_domain.Asset, error)
	GetByCode(code string) (asset_domain.Asset, error)
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
	GetAssetQuarterlyResults(idAtivo int64, idTrimestre int64) ([]asset_domain.AssetQuarterlyResult, error)
}

func NewService(repository Repository, stockPriceRepository StockPriceRepository, assetQResultRepository AssetQuarterlyResultRepository,
	quarterlyProducer event.QuarterlyResultProducer) Service {
	return Service{
		repository:                     repository,
		stockPriceRepository:           stockPriceRepository,
		assetQuarterlyResultRepository: assetQResultRepository,
		quarterlyProducer:              quarterlyProducer,
	}
}

func (s Service) GetAllAssets() ([]asset_domain.Asset, error) {
	return s.repository.GetAllAsset()
}

func (s Service) ExistById(id int64) (bool, error) {
	return s.repository.ExistById(id)
}

func (s Service) GetById(id int64) (asset_domain.Asset, error) {
	return s.repository.GetById(id)
}

func (s Service) GetByCode(code string) (asset_domain.Asset, error) {
	return s.repository.GetByCode(code)
}

func (s Service) InsertAsset(asset asset_domain.Asset) (bool, error) {
	return s.repository.InsertAsset(asset)
}

func (s Service) UpdateAsset(asset asset_domain.Asset) (asset_domain.Asset, error) {
	return s.repository.UpdateAsset(asset)
}

func (s Service) UpdateAssetPrice(id int64, price float32, data time.Time) (bool, error) {
	ok, err := s.repository.UpdateAssetPrice(id, price)
	if err != nil {
		return false, err
	}
	s.InsertAssetPrice(id, price, data)
	return ok, nil
}

func (s Service) InsertAssetPrice(idAtivo int64, price float32, data time.Time) (bool, error) {
	assetPrice, err := s.stockPriceRepository.GetByAtivoEData(idAtivo, data)
	if err != nil {
		return false, err
	}
	if assetPrice.Id != 0 {
		log.Print("Asset Price já cadastrado %d - %v.", idAtivo, data)
		return false, nil
	}
	updateAssetPrice := asset_domain.AssetPrice{
		Ativo:   idAtivo,
		Cotacao: price,
		Data:    data,
	}
	return s.stockPriceRepository.InsertAssetPrice(updateAssetPrice)
}

func (s Service) ExistAssetQuarterlyResult(idAtivo int64, idTrimestre int64) (bool, error) {
	return s.assetQuarterlyResultRepository.ExistAssetQuarterlyResult(idAtivo, idTrimestre)
}

func (s Service) InsertAssetQuarterlyResult(aqResult asset_domain.AssetQuarterlyResult) (bool, error) {
	exist, err := s.assetQuarterlyResultRepository.ExistAssetQuarterlyResult(aqResult.Ativo, aqResult.Trimestre)
	if exist {
		log.Print("Já existe um resultado cadastro para este trimestre.")
		return false, nil
	}
	if err != nil {
		return false, err
	}
	s.quarterlyProducer.PublishQuarterlyResultEvent(aqResult)
	return s.assetQuarterlyResultRepository.InsertAssetQuarterlyResult(aqResult)
}

func (s Service) GetAssetQuarterlyResults(assetId int64) ([]asset_domain.AssetQuarterlyResult, error) {
	quarterlyResults, err := s.assetQuarterlyResultRepository.GetAssetQuarterlyResults(assetId, 0)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(quarterlyResults); i++ {
		calcularMargens(&quarterlyResults[i])
	}
	return quarterlyResults, nil
}

func (s Service) GetAssetQuarterlyResultsByTrimestre(assetId int64, trimestre int64) ([]asset_domain.AssetQuarterlyResult, error) {
	quarterlyResults, err := s.assetQuarterlyResultRepository.GetAssetQuarterlyResults(assetId, trimestre)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(quarterlyResults); i++ {
		calcularMargens(&quarterlyResults[i])
	}
	return quarterlyResults, nil
}

func calcularMargens(result *asset_domain.AssetQuarterlyResult) {
	result.MargemEbitda = internal.RoundFloat(float32(result.Ebitda) / float32(result.ReceitaLiquida))
	result.MargemLiquida = internal.RoundFloat(float32(result.LucroLiquido) / float32(result.ReceitaLiquida))
	if result.DividaLiquida > 0 {
		result.DivEbitda = internal.RoundFloat(float32(result.DividaLiquida) / float32(result.Ebitda))
	}
}