package asset_sync

import (
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/pkg/asset/asset-sync/alphavantage"
	"github.com/crisaltmann/fundament-stock-api/pkg/asset/service"
	"github.com/rs/zerolog/log"
	"strconv"
)

type Service struct {
	AssetService asset_service.Service
	Client 		 alphavantage.Client
}

func (s Service) updateAssetPrice() {
	assets, err := s.AssetService.GetAllAssets()
	if err != nil {
		err = fmt.Errorf("Ocorreu um erro ao executar atualização de preço ativos.", err)
		log.Err(err)
		return
	}
	for _, asset := range assets {
		quote, err := s.Client.GetGlobalQuote(asset.Codigo)
		if err != nil {
			err = fmt.Errorf("Ocorreu um erro ao buscar a cotação do ativo: " + asset.Codigo, err)
			log.Err(err)
		}
		price, err := strconv.ParseFloat(quote.Price, 32)
		ok, err := s.AssetService.UpdateAssetPrice(asset.Id, float32(price))
		if !ok || err != nil {
			err = fmt.Errorf("Ocorreu um erro na atualizacao da cotação do ativo: " + asset.Codigo, err)
			log.Err(err)
		}
	}
}
