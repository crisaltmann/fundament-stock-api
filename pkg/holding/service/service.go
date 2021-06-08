package holding_service

import (
	"errors"
	asset_domain "github.com/crisaltmann/fundament-stock-api/pkg/asset/domain"
	holding_domain "github.com/crisaltmann/fundament-stock-api/pkg/holding/domain"
	portfolio_domain "github.com/crisaltmann/fundament-stock-api/pkg/portfolio/domain"
	"github.com/rs/zerolog/log"
)

type Service struct {
	PortfolioService PortfolioService
	AssetService AssetService
}

type PortfolioService interface {
	GetPortfolio(usuario string) ([]portfolio_domain.Portfolio, error)
}

type AssetService interface {
	GetAssetQuarterlyResults(assetId int64) ([]asset_domain.AssetQuarterlyResult, error)
}


func NewService(portfolioService PortfolioService, assetService AssetService) Service {
	return Service{
		PortfolioService: portfolioService,
		AssetService:     assetService,
	}
}

func (s Service) GetHolding(usuario string) (holding_domain.Holdings, error) {
	portfolio, err := s.PortfolioService.GetPortfolio(usuario)
	if err != nil {
		log.Print("Erro ao buscar portfolio no calculo de holding.")
		return holding_domain.Holdings{}, errors.New("Erro ao buscar portfolio no calculo de holding.")
	}

	if len(portfolio) == 0 {
		log.Print("Não foram encontrados ativos no portfolio do usuario " + usuario)
		holdings := make([]holding_domain.Holding, 0)
		return holding_domain.Holdings{ Holdings: holdings}, nil
	}

	resultadosHolding := make(map[int64]*holding_domain.Holding)

	for _, item := range portfolio {
		if item.Ativo.Codigo != "WEGE3" {
			continue
		}
		//tem o ativo
		//buscar todos os resultados deste ativo
		quarterlyResults, err := s.AssetService.GetAssetQuarterlyResults(item.Ativo.Id)
		if err != nil {
			log.Print("Erro ao buscar resultados trimestrais dos ativos no portfolio no calculo de holding.")
			return holding_domain.Holdings{}, errors.New("Erro ao buscar resultados trimestrais dos ativos no portfolio no calculo de holding.")
		}

		if len(quarterlyResults) == 0 {
			log.Print("Não foram encontrados resultados trimestrais dos ativos no portfolio do usuario " + usuario)
			holdings := make([]holding_domain.Holding, 0)
			return holding_domain.Holdings{ Holdings: holdings}, nil
		}

		for _, quarterlyItem := range quarterlyResults {
			holdingQuarterly, exist := resultadosHolding[quarterlyItem.Trimestre]
			if !exist {
				holdingQuarterly = &holding_domain.Holding{
					Id:             0,
					Trimestre:      quarterlyItem.Trimestre,
				}
				resultadosHolding[quarterlyItem.Trimestre] = holdingQuarterly
			}
			holdingQuarterly.ReceitaLiquida = CalcularFundamentos(item, quarterlyItem)
		}
	}

	holdings := make([]holding_domain.Holding, 0)
	for _, result := range resultadosHolding {
		holdings = append(holdings, holding_domain.Holding{
			Id:             result.Id,
			Trimestre:      result.Trimestre,
			ReceitaLiquida: result.ReceitaLiquida,
			Ebitda:         0,
			MargemEbitda:   0,
			LucroLiquido:   0,
			MargemLiquida:  0,
			DividaLiquida:  0,
			DivEbitda:      0,
		})
	}
	return holding_domain.Holdings{Holdings: holdings}, nil
}

func CalcularFundamentos(item portfolio_domain.Portfolio, quarterlyItem asset_domain.AssetQuarterlyResult) int64{
	//calcular o percentual
	percentualDetido := float32(item.Quantidade) / float32(item.Ativo.Total)
	//somar os percentuais do trimestre
	return int64(float32(quarterlyItem.ReceitaLiquida) * percentualDetido)
}