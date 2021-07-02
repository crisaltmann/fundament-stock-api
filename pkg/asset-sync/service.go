package asset_sync

import (
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/pkg/asset/domain"
	"github.com/rs/zerolog/log"
	"time"
)

type JobService struct {
	AssetFinder  AssetFinder
	AssetUpdater AssetUpdater
	StockPriceFinder  StockPriceFinder
}

type AssetFinder interface {
	GetAllAssets() ([]asset_domain.Asset, error)
}

type AssetUpdater interface {
	UpdateAssetPrice(id int64, price float32, data time.Time) (bool, error)
}

type StockPriceFinder interface {
	GetStockPrice(code string) (float32, time.Time, error)
}

func NewService(finder AssetFinder, assetUpdater AssetUpdater, stockFinder StockPriceFinder) JobService {
	return JobService{
		AssetFinder:  finder,
		AssetUpdater: assetUpdater,
		StockPriceFinder:  stockFinder,
	}
}

func (s JobService) UpdateAssetPrice() {
	assets, err := s.AssetFinder.GetAllAssets()
	if err != nil {
		err = fmt.Errorf("Ocorreu um erro ao executar atualização de preço ativos.", err)
		log.Err(err)
		return
	}
	for _, asset := range assets {
		log.Print("Iniciando atualização de cotação do ativo: " + asset.Codigo)
		price, data, err := s.StockPriceFinder.GetStockPrice(asset.Codigo)
		if err != nil {
			err = fmt.Errorf("Ocorreu um erro ao buscar a cotação do ativo: " + asset.Codigo, err)
			log.Err(err)
			continue
		}

		ok, err := s.AssetUpdater.UpdateAssetPrice(asset.Id, price, data)
		if !ok || err != nil {
			err = fmt.Errorf("Ocorreu um erro na atualizacao da cotação do ativo: " + asset.Codigo, err)
			log.Err(err)
		}
	}
}
