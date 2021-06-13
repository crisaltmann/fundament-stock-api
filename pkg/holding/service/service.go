package holding_service

import (
	"errors"
	asset_domain "github.com/crisaltmann/fundament-stock-api/pkg/asset/domain"
	holding_domain "github.com/crisaltmann/fundament-stock-api/pkg/holding/domain"
	portfolio_domain "github.com/crisaltmann/fundament-stock-api/pkg/portfolio/domain"
	quarter_domain "github.com/crisaltmann/fundament-stock-api/pkg/quarter/domain"
	"github.com/rs/zerolog/log"
	"strconv"
)

type Service struct {
	PortfolioService PortfolioService
	AssetService AssetService
	QuarterService QuarterService
	repository Repository
}

type PortfolioService interface {
	GetPortfolio(usuario string) ([]portfolio_domain.Portfolio, error)
}

type AssetService interface {
	GetAssetQuarterlyResults(assetId int64) ([]asset_domain.AssetQuarterlyResult, error)
	GetById(id int64) (asset_domain.Asset, error)
}

type QuarterService interface {
	GetQuarter(id int64) (quarter_domain.Trimestre, error)
}

type Repository interface {
	GetResultadoPortfolio(usuario string) ([]holding_domain.HoldingAtivo, error)
}

func NewService(portfolioService PortfolioService, assetService AssetService, quarterService QuarterService, repository Repository) Service {
	return Service{
		PortfolioService: portfolioService,
		AssetService:     assetService,
		QuarterService: quarterService,
		repository: repository,
	}
}

func (s Service) GetHolding(usuario string) (holding_domain.Holdings, error) {
	resultados, err := s.repository.GetResultadoPortfolio(usuario)
	if err != nil {
		log.Print("Erro ao buscar holding.")
		return holding_domain.Holdings{}, err
	}
	holdingMap := make(map[string]*holding_domain.Holding)

	for _, resultado := range resultados {
		key := strconv.FormatInt(resultado.Trimestre, 10)
		holding, found := holdingMap[key]
		if !found {
			holding = &holding_domain.Holding{}
			holdingMap[key] = holding
		}
		trimestre, err := s.QuarterService.GetQuarter(resultado.Trimestre)
		if err != nil {
			log.Print("Erro ao buscar trimestre na busca de holdings.")
			return holding_domain.Holdings{}, err
		}

		holding.HoldingsAtivo = append(holding.HoldingsAtivo, resultado)
		holding.DividaLiquida += resultado.DividaLiquida
		holding.ReceitaLiquida += resultado.ReceitaLiquida
		holding.Ebitda += resultado.Ebitda
		holding.LucroLiquido += resultado.LucroLiquido
		holding.Trimestre = trimestre
	}

	holdings := holding_domain.Holdings{}

	for _, holdingMap := range holdingMap {
		holdings.Holdings = append(holdings.Holdings, holdingMap.ToStruct())
	}

	return holdings, nil
}

func (s Service) CalculateHolding(usuario string) (holding_domain.Holdings, error) {
	//TODO adicionar filtro por data de trimestre quando for persistido.
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
	resultadosHoldingByAtivo := make(map[string]*holding_domain.HoldingAtivo)

	for _, portfolioItem := range portfolio {
		quarterlyResults, err := s.AssetService.GetAssetQuarterlyResults(portfolioItem.Ativo.Id)
		if err != nil {
			log.Print("Erro ao buscar resultados trimestrais dos ativos no portfolio no calculo de holding.")
			return holding_domain.Holdings{}, errors.New("Erro ao buscar resultados trimestrais dos ativos no portfolio no calculo de holding.")
		}

		if len(quarterlyResults) == 0 {
			log.Print("Não foram encontrados resultados trimestrais dos ativos no portfolio do usuario " + usuario)
			continue
		}

		for _, quarterlyItem := range quarterlyResults {
			err2 := s.buildHoldingQuarterlyResult(quarterlyItem, portfolioItem, resultadosHolding, resultadosHoldingByAtivo)
			if err2 != nil {
				return holding_domain.Holdings{}, err2
			}
		}
	}

	return s.buildHoldingReturn(resultadosHolding, resultadosHoldingByAtivo)
}

func (s Service) buildHoldingQuarterlyResult(quarterlyItem asset_domain.AssetQuarterlyResult, portfolioItem portfolio_domain.Portfolio,
	resultadosHolding map[int64]*holding_domain.Holding,
	resultadosHoldingByAtivo map[string]*holding_domain.HoldingAtivo) error {

	quarter, err := s.QuarterService.GetQuarter(quarterlyItem.Trimestre)
	if err != nil {
		log.Print("Erro ao buscar quarter.")
		return errors.New("Erro ao buscar quarter.")
	}

	ativo, err := s.AssetService.GetById(quarterlyItem.Ativo)
	if err != nil {
		log.Print("Erro ao buscar ativo.")
		return errors.New("Erro ao buscar ativo.")
	}

	s.buildQuarterly(quarterlyItem, portfolioItem, quarter, resultadosHolding)
	s.buildQuarterlyAtivo(quarterlyItem, portfolioItem, quarter, ativo, resultadosHoldingByAtivo)

	return nil
}

func (s Service) buildQuarterly(quarterlyItem asset_domain.AssetQuarterlyResult, portfolioItem portfolio_domain.Portfolio,
	quarter quarter_domain.Trimestre, resultadosHolding map[int64]*holding_domain.Holding) {
	holdingQuarterly, exist := resultadosHolding[quarterlyItem.Trimestre]
	if !exist {
		holdingQuarterly = &holding_domain.Holding{
			Trimestre: quarter,
		}
		resultadosHolding[quarterlyItem.Trimestre] = holdingQuarterly
	}
	holdingQuarterly.ReceitaLiquida += CalcularFundamentos(portfolioItem, quarterlyItem)
}

func (s Service) buildQuarterlyAtivo(quarterlyItem asset_domain.AssetQuarterlyResult, portfolioItem portfolio_domain.Portfolio,
	quarter quarter_domain.Trimestre, ativo asset_domain.Asset, resultadosHoldingByAtivo map[string]*holding_domain.HoldingAtivo) {
	key := quarter.Codigo + "-" + ativo.Codigo
	holdingQuarterlyAtivo, exist := resultadosHoldingByAtivo[key]
	if !exist {
		holdingQuarterlyAtivo = &holding_domain.HoldingAtivo{
			Trimestre: quarter.Id,
			Ativo: ativo,
		}
		resultadosHoldingByAtivo[key] = holdingQuarterlyAtivo
	}

	holdingQuarterlyAtivo.ReceitaLiquida += CalcularFundamentos(portfolioItem, quarterlyItem)
}

func (s Service) buildHoldingReturn(resultadosHolding map[int64]*holding_domain.Holding, resultadosHoldingByAtivo map[string]*holding_domain.HoldingAtivo) (holding_domain.Holdings, error) {
	holdings := make([]holding_domain.Holding, 0)
	for _, result := range resultadosHolding {

		holdingsAtivo := make([]holding_domain.HoldingAtivo, 0)

		for _, holAtivo := range resultadosHoldingByAtivo {
			if holAtivo.Trimestre != result.Trimestre.Id {
				continue
			}
			holdingsAtivo = append(holdingsAtivo, holAtivo.ToStruct())
		}

		result.HoldingsAtivo = holdingsAtivo
		holdings = append(holdings, result.ToStruct())
	}
	return holding_domain.Holdings{Holdings: holdings}, nil
}

func CalcularFundamentos(item portfolio_domain.Portfolio, quarterlyItem asset_domain.AssetQuarterlyResult) int64{
	//calcular o percentual
	percentualDetido := float32(item.Quantidade) / float32(item.Ativo.Total)
	//somar os percentuais do trimestre
	return int64(float32(quarterlyItem.ReceitaLiquida) * percentualDetido)
}