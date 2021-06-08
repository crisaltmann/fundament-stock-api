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
		log.Print("Não foram encontrados ativos no portfolio do usuario %d.", usuario)
		holdings := make([]holding_domain.Holding, 0)
		return holding_domain.Holdings{ Holdings: holdings}, nil
	}

	resultadosHolding := make(map[int64]holding_domain.Holding)

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
			log.Print("Não foram encontrados resultados trimestrais dos ativos no portfolio do usuario %d.", usuario)
			holdings := make([]holding_domain.Holding, 0)
			return holding_domain.Holdings{ Holdings: holdings}, nil
		}

		for _, quarterlyItem := range quarterlyResults {
			holdingQuarterly, exist := resultadosHolding[quarterlyItem.Trimestre]
			if !exist {
				holdingQuarterly = holding_domain.Holding{
					Id:             0,
					Trimestre:      quarterlyItem.Trimestre,
				}
				resultadosHolding[quarterlyItem.Trimestre] = holdingQuarterly
			}
			//calcular o percentual
			percentualDetido := float32(item.Quantidade) / float32(item.Ativo.Total)
			//somar os percentuais do trimestre
			holdingQuarterly.ReceitaLiquida += int64(float32(quarterlyItem.ReceitaLiquida) * percentualDetido)
		}
	}

	holdings := make([]holding_domain.Holding, len(resultadosHolding))
	for _, result := range resultadosHolding {
		holdings = append(holdings, result)
	}
	return holding_domain.Holdings{Holdings: holdings}, nil
}

func (s Service) GetHoldingMock(usuario string) (holding_domain.Holdings, error) {
	holdings := make([]holding_domain.Holding, 0)

	h1 := holding_domain.Holding{
		Id:             1,
		Trimestre:      2,
		ReceitaLiquida: 100,
		Ebitda:         200,
		MargemEbitda:   0,
		LucroLiquido:   0,
		MargemLiquida:  0,
		DividaLiquida:  0,
		DivEbitda:      0,
	}

	h2 := holding_domain.Holding{
		Id:             2,
		Trimestre:      2,
		ReceitaLiquida: 100,
		Ebitda:         200,
		MargemEbitda:   2222,
		LucroLiquido:   0,
		MargemLiquida:  0,
		DividaLiquida:  0123111111,
		DivEbitda:      0,
	}

	holdings = append(holdings, h1)
	holdings = append(holdings, h2)

	return holding_domain.Holdings{Holdings: holdings}, nil
}