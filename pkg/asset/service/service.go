package asset_service

import (
	"github.com/crisaltmann/fundament-stock-api/internal"
	"github.com/crisaltmann/fundament-stock-api/pkg/asset/domain"
	"github.com/crisaltmann/fundament-stock-api/pkg/asset/event"
	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog/log"
	"strconv"
	"time"
)

type Service struct {
	Repository                     Repository
	Cache 						   *cache.Cache
	StockPriceRepository           StockPriceRepository
	AssetQuarterlyResultRepository AssetQuarterlyResultRepository
	QuarterlyProducer              event.QuarterlyResultProducer
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
	GetAssetQuarterlyResults(idAtivo int64) ([]asset_domain.AssetQuarterlyResult, error)
}

func NewService(repository Repository, stockPriceRepository StockPriceRepository, assetQResultRepository AssetQuarterlyResultRepository,
	quarterlyProducer event.QuarterlyResultProducer) Service {

	cache := cache.New(1*time.Hour, 10*time.Minute)

	return Service{
		Repository: repository,
		StockPriceRepository: stockPriceRepository,
		AssetQuarterlyResultRepository: assetQResultRepository,
		QuarterlyProducer: quarterlyProducer,
		Cache: cache,
	}
}

func (s Service) GetAllAssets() ([]asset_domain.Asset, error) {
	return s.Repository.GetAllAsset()
}

func (s Service) ExistById(id int64) (bool, error) {
	return s.Repository.ExistById(id)
}

func (s Service) GetById(id int64) (asset_domain.Asset, error) {
	asset, found :=  s.Cache.Get(strconv.FormatInt(id, 10))
	if found {
		return asset.(asset_domain.Asset), nil
	}

	assetDB, err := s.Repository.GetById(id)

	if err != nil {
		s.Cache.Add(strconv.FormatInt(id, 10), assetDB, cache.DefaultExpiration)
	}
	return assetDB, err
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

func (s Service) InsertAssetPrice(idAtivo int64, price float32, data time.Time) (bool, error) {
	assetPrice, err := s.StockPriceRepository.GetByAtivoEData(idAtivo, data)
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
	return s.StockPriceRepository.InsertAssetPrice(updateAssetPrice)
}

func (s Service) InsertAssetQuarterlyResult(aqResult asset_domain.AssetQuarterlyResult) (bool, error) {
	exist, err := s.AssetQuarterlyResultRepository.ExistAssetQuarterlyResult(aqResult.Ativo, aqResult.Trimestre)
	if exist {
		log.Print("Já existe um resultado cadastro para este trimestre.")
		return false, nil
	}
	if err != nil {
		return false, err
	}
	s.QuarterlyProducer.PublishQuarterlyResultEvent(aqResult)
	return s.AssetQuarterlyResultRepository.InsertAssetQuarterlyResult(aqResult)
}

func (s Service) GetAssetQuarterlyResults(assetId int64) ([]asset_domain.AssetQuarterlyResult, error) {
	quarterlyResults, err := s.AssetQuarterlyResultRepository.GetAssetQuarterlyResults(assetId)
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