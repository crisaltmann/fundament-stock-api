package holding_service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/internal"
	asset_domain "github.com/crisaltmann/fundament-stock-api/pkg/asset/domain"
	holding_domain "github.com/crisaltmann/fundament-stock-api/pkg/holding/domain"
	holding_event "github.com/crisaltmann/fundament-stock-api/pkg/holding/event"
	portfolio_domain "github.com/crisaltmann/fundament-stock-api/pkg/portfolio/domain"
	quarter_domain "github.com/crisaltmann/fundament-stock-api/pkg/quarter/domain"
	"github.com/rs/zerolog/log"
	"strconv"
)

type Service struct {
	portfolioService PortfolioService
	assetService AssetService
	quarterService QuarterService
	orderService OrderService
	repository Repository
	producer holding_event.HoldingResultProducer
	db *sql.DB
}

type OrderService interface {
	GetUsersWithOrders(idAtivo int64) ([]int64, error)
}

type PortfolioService interface {
	GetPortfolioByTrimestre(usuario int64, trimestre int64) ([]portfolio_domain.Portfolio, error)
}

type AssetService interface {
	GetAssetQuarterlyResultsByTrimestre(assetId int64, trimestre int64) ([]asset_domain.AssetQuarterlyResult, error)
	GetById(id int64) (asset_domain.Asset, error)
}

type QuarterService interface {
	GetQuarter(id int64) (quarter_domain.Trimestre, error)
	GetQuarters() ([]quarter_domain.Trimestre, error)
}

type Repository interface {
	GetResultadoPortfolio(usuario string) ([]holding_domain.HoldingAtivo, error)
	DeleteByUser(ctx context.Context, tx *sql.Tx, idUser int64) error
	SaveResultadoPortfolio(ctx context.Context, tx *sql.Tx, ativo holding_domain.HoldingAtivo) error
}

func NewService(portfolioService PortfolioService, assetService AssetService, quarterService QuarterService,
	orderService OrderService, repository Repository, producer holding_event.HoldingResultProducer, db *sql.DB) Service {
	return Service{
		portfolioService: portfolioService,
		assetService:     assetService,
		quarterService: quarterService,
		orderService: orderService,
		repository: repository,
		producer: producer,
		db: db,
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
		trimestre, err := s.quarterService.GetQuarter(resultado.Trimestre)
		if err != nil {
			log.Print("Erro ao buscar trimestre na busca de holdings.")
			return holding_domain.Holdings{}, err
		}

		if resultado.ReceitaLiquida > 0 {
			resultado.MargemEbitda = internal.RoundFloat(float32(resultado.Ebitda) / float32(resultado.ReceitaLiquida))
			resultado.MargemLiquida = internal.RoundFloat(float32(resultado.LucroLiquido) / float32(resultado.ReceitaLiquida))
		}

		if resultado.DividaLiquida > 0 && resultado.Ebitda > 0 {
			resultado.DivEbitda = internal.RoundFloat(float32(resultado.DividaLiquida) / float32(resultado.Ebitda))
		}

		holding.HoldingsAtivo = append(holding.HoldingsAtivo, resultado)
		holding.DividaLiquida += resultado.DividaLiquida
		holding.ReceitaLiquida += resultado.ReceitaLiquida
		holding.Ebitda += resultado.Ebitda
		holding.LucroLiquido += resultado.LucroLiquido
		holding.Trimestre = trimestre
		holding.MargemEbitda = internal.RoundFloat(float32(holding.Ebitda) / float32(holding.ReceitaLiquida))
		holding.MargemLiquida = internal.RoundFloat(float32(holding.LucroLiquido) / float32(holding.ReceitaLiquida))

		if holding.DividaLiquida > 0 && holding.Ebitda > 0 {
			holding.DivEbitda = internal.RoundFloat(float32(holding.DividaLiquida) / float32(holding.Ebitda))
		}
	}

	holdings := holding_domain.Holdings{}

	for _, holdingMap := range holdingMap {
		holdings.Holdings = append(holdings.Holdings, holdingMap.ToStruct())
	}

	return holdings, nil
}

func (s Service) CalculateHolding(ctx context.Context, idAtivo int64) error {
	users, err := s.orderService.GetUsersWithOrders(idAtivo)
	if err != nil {
		return err
	}

	//TODO refactor this, because transaction can't be in service layer
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("Erro ao iniciar transacao")
	}

	for i := 0; i < len(users); i++ {
		err := s.repository.DeleteByUser(ctx, tx, users[i])
		if err != nil {
			log.Print("Erro ao deletar resultados de holging por usuario")
			tx.Rollback()
			return err
		}
		holdings, err := s.calculateHoldingGeneral(users[i])
		if err != nil {
			log.Printf("Erro ao calcular os resultados de holding.")
			tx.Rollback()
			return err
		}

		err = s.saveHoldings(ctx, tx, holdings)
		if err != nil {
			log.Printf("Ocorreu um erro ao salvar resultados de holding")
			tx.Rollback()
			return err
		}
		tx.Commit()
		s.producer.PublishHoldingResultEvent(holdings)
	}

	return nil
}

func (s Service) saveHoldings(ctx context.Context, tx *sql.Tx, holdings holding_domain.Holdings) error {
	for i := 0; i < len(holdings.Holdings); i++ {
		holding := holdings.Holdings[i]
		for j := 0; j < len(holding.HoldingsAtivo); j++ {
			err := s.repository.SaveResultadoPortfolio(ctx, tx, holding.HoldingsAtivo[j])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s Service) calculateHoldingGeneral(usuario int64) (holding_domain.Holdings, error) {
	quarters, err := s.quarterService.GetQuarters()
	if err != nil {
		log.Printf("Erro ao buscar trimestres.")
		return holding_domain.Holdings{}, err
	}

	resultadosHolding := make(map[int64]*holding_domain.Holding)
	resultadosHoldingByAtivo := make(map[string]*holding_domain.HoldingAtivo)

	for i := 0; i < len(quarters); i++ {
		currentQuarter := quarters[i]
		portfolio, err := s.portfolioService.GetPortfolioByTrimestre(usuario, currentQuarter.Id)
		if err != nil {
			log.Printf("Erro ao buscar portfolio no calculo de holding %d.", currentQuarter.Id)
			return holding_domain.Holdings{}, errors.New("Erro ao buscar portfolio no calculo de holding.")
		}

		if len(portfolio) == 0 {
			log.Printf("Não foram encontrados ativos no portfolio do usuario %d no trimestre %d", usuario, currentQuarter.Id)
			continue
		}

		for _, portfolioItem := range portfolio {
			quarterlyResults, err := s.assetService.GetAssetQuarterlyResultsByTrimestre(portfolioItem.Ativo.Id, currentQuarter.Id)
			if err != nil {
				log.Print("Erro ao buscar resultados trimestrais dos ativos no portfolio no calculo de holding.")
				return holding_domain.Holdings{}, errors.New("Erro ao buscar resultados trimestrais dos ativos no portfolio no calculo de holding.")
			}

			if len(quarterlyResults) == 0 {
				log.Printf("Não foram encontrados resultados trimestrais dos ativos no portfolio do usuario %d", usuario)
				continue
			}

			for _, quarterlyItem := range quarterlyResults {
				err2 := s.buildHoldingQuarterlyResult(usuario, quarterlyItem, portfolioItem, resultadosHolding, resultadosHoldingByAtivo)
				if err2 != nil {
					return holding_domain.Holdings{}, err2
				}
			}
		}
	}
	return s.buildHoldingReturn(resultadosHolding, resultadosHoldingByAtivo)
}

func (s Service) buildHoldingQuarterlyResult(idUsuario int64, quarterlyItem asset_domain.AssetQuarterlyResult, portfolioItem portfolio_domain.Portfolio,
	resultadosHolding map[int64]*holding_domain.Holding,
	resultadosHoldingByAtivo map[string]*holding_domain.HoldingAtivo) error {

	quarter, err := s.quarterService.GetQuarter(quarterlyItem.Trimestre)
	if err != nil {
		log.Print("Erro ao buscar quarter.")
		return errors.New("Erro ao buscar quarter.")
	}

	ativo, err := s.assetService.GetById(quarterlyItem.Ativo)
	if err != nil {
		log.Print("Erro ao buscar ativo.")
		return errors.New("Erro ao buscar ativo.")
	}

	s.buildQuarterly(quarterlyItem, portfolioItem, quarter, resultadosHolding)
	s.buildQuarterlyAtivo(idUsuario, quarterlyItem, portfolioItem, quarter, ativo, resultadosHoldingByAtivo)

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

	receitaLiquida, ebitda, lucroLiquido, divida := CalcularFundamentos(portfolioItem, quarterlyItem)
	holdingQuarterly.ReceitaLiquida += receitaLiquida
	holdingQuarterly.Ebitda += ebitda
	holdingQuarterly.LucroLiquido += lucroLiquido
	holdingQuarterly.DividaLiquida += divida
}

func (s Service) buildQuarterlyAtivo(usuario int64, quarterlyItem asset_domain.AssetQuarterlyResult, portfolioItem portfolio_domain.Portfolio,
	quarter quarter_domain.Trimestre, ativo asset_domain.Asset, resultadosHoldingByAtivo map[string]*holding_domain.HoldingAtivo) {
	key := quarter.Codigo + "-" + ativo.Codigo
	holdingQuarterlyAtivo, exist := resultadosHoldingByAtivo[key]
	if !exist {
		holdingQuarterlyAtivo = &holding_domain.HoldingAtivo{
			Trimestre: quarter.Id,
			Ativo: ativo,
			Usuario: usuario,
		}
		resultadosHoldingByAtivo[key] = holdingQuarterlyAtivo
	}

	receitaLiquida, ebitda, lucroLiquido, divida := CalcularFundamentos(portfolioItem, quarterlyItem)
	holdingQuarterlyAtivo.ReceitaLiquida += receitaLiquida
	holdingQuarterlyAtivo.Ebitda += ebitda
	holdingQuarterlyAtivo.LucroLiquido += lucroLiquido
	holdingQuarterlyAtivo.DividaLiquida += divida
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

/**
retorno receitaLiquida, ebitda, lucroLiquido, divida
 */
func CalcularFundamentos(item portfolio_domain.Portfolio, quarterlyItem asset_domain.AssetQuarterlyResult) (int64, int64, int64, int64){
	//calcular o percentual
	percentualDetido := float32(item.Quantidade) / float32(item.Ativo.Total)
	//somar os percentuais do trimestre
	return int64(float32(quarterlyItem.ReceitaLiquida) * percentualDetido),
		int64(float32(quarterlyItem.Ebitda) * percentualDetido),
		int64(float32(quarterlyItem.LucroLiquido) * percentualDetido),
		int64(float32(quarterlyItem.DividaLiquida) * percentualDetido)
}