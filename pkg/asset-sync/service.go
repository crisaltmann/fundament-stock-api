package asset_sync

import (
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/pkg/asset/domain"
	"github.com/rs/zerolog/log"
)

type JobService struct {
	AssetFinder  AssetFinder
	AssetUpdater AssetUpdater
	StockFinder  StockPriceFinder
}

type AssetFinder interface {
	GetAllAssets() ([]asset_domain.Asset, error)
}

type AssetUpdater interface {
	UpdateAssetPrice(id int64, price float32) (bool, error)
}

type StockPriceFinder interface {
	GetStockPrice(code string) (float32, error)
}

func NewService(finder AssetFinder, assetUpdater AssetUpdater, stockFinder StockPriceFinder) JobService {
	return JobService{
		AssetFinder:  finder,
		AssetUpdater: assetUpdater,
		StockFinder:  stockFinder,
	}
}

func (s JobService) updateAssetPrice() {
	assets, err := s.AssetFinder.GetAllAssets()
	if err != nil {
		err = fmt.Errorf("Ocorreu um erro ao executar atualização de preço ativos.", err)
		log.Err(err)
		return
	}
	for _, asset := range assets {
		price, err := s.StockFinder.GetStockPrice(asset.Codigo)
		if err != nil {
			err = fmt.Errorf("Ocorreu um erro ao buscar a cotação do ativo: " + asset.Codigo, err)
			log.Err(err)
		}

		ok, err := s.AssetUpdater.UpdateAssetPrice(asset.Id, price)
		if !ok || err != nil {
			err = fmt.Errorf("Ocorreu um erro na atualizacao da cotação do ativo: " + asset.Codigo, err)
			log.Err(err)
		}
	}
}
