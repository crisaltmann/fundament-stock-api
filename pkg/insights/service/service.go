package insight_service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/internal"
	asset_domain "github.com/crisaltmann/fundament-stock-api/pkg/asset/domain"
	holding_domain "github.com/crisaltmann/fundament-stock-api/pkg/holding/domain"
	insight_domain "github.com/crisaltmann/fundament-stock-api/pkg/insights/domain"
	quarter_domain "github.com/crisaltmann/fundament-stock-api/pkg/quarter/domain"
	"log"
	"strconv"
)

type Service struct {
	repository 	Repository
	quarterService QuarterService
	db *sql.DB
}

type Repository interface {
	DeleteByUser(ctx context.Context, tx *sql.Tx, idUser int64) error
	SaveInsights(ctx context.Context, tx *sql.Tx, insight insight_domain.Insight) error
	GetInsights(usuario int64) ([]insight_domain.Insight, error)
}

type QuarterService interface {
	GetQuarter(id int64) (quarter_domain.Trimestre, error)
	GetQuarters() ([]quarter_domain.Trimestre, error)
}

func NewService(repository Repository, quarterService QuarterService, db *sql.DB) Service {
	return Service{
		repository: repository,
		quarterService: quarterService,
		db: db,
	}
}

func (s Service) GetInsights(usuario int64) ([]insight_domain.Insight, error) {
	return s.repository.GetInsights(usuario)
}

func (s Service) GetSummaryInsights(usuario int64) (insight_domain.InsightsSummary, error) {
	insights, err := s.repository.GetInsights(usuario)
	insightsSummary := insight_domain.InsightsSummary{}
	if err != nil {
		log.Print("Erro ao buscar insights")
		return insightsSummary, err
	}

	//monta map do trimestre, armazenando o ativo com maior delta em cada
	insightMap := make(map[int64]*insight_domain.InsightSummary)
	for i := 0; i < len(insights); i++ {
		insight := insights[i]
		summary, found := insightMap[insight.IdTrimestre]
		if !found {
			summary = &insight_domain.InsightSummary{
				Trimestre: insight.IdTrimestre,
			}
			insightMap[insight.IdTrimestre] = summary
		}

		if insight.ReceitaDelta > summary.ReceitaMaiorDelta {
			summary.ReceitaMaiorDelta = insight.ReceitaDelta
			summary.AtivoMaiorReceita = insight.IdAtivo
		}

		if insight.EbitdaDelta > summary.EbitdaMaiorDelta {
			summary.EbitdaMaiorDelta = insight.EbitdaDelta
			summary.AtivoMaiorEbitda = insight.IdAtivo
		}

		if insight.LucroDelta > summary.LucroMaiorDelta {
			summary.LucroMaiorDelta = insight.LucroDelta
			summary.AtivoMaiorLucro = insight.IdAtivo
		}

		if insight.DividaDelta > summary.DividaDelta {
			summary.DividaDelta = insight.DividaDelta
			summary.AtivoMaiorDivida = insight.IdAtivo
		}
	}

	for _, summary := range insightMap {
		insightsSummary.Insights = append(insightsSummary.Insights, summary.ToStruct())
	}

	return insightsSummary, nil
}

func (s Service) CalculateInsights(ctx context.Context, holdings holding_domain.Holdings) error {
	/**
		montar map com chave: idAtivo-trimestre

		Terei todos os ativos por trimestre.

	 */
	var user int64
	insights := make([]insight_domain.Insight, 0)
	trimestreMap := make(map[string]holding_domain.HoldingAtivo)
	for i := 0; i < len(holdings.Holdings); i++ {
		holding := holdings.Holdings[i]
		for j := 0; j < len(holding.HoldingsAtivo); j++ {
			user = holding.HoldingsAtivo[j].Usuario
			holdingAtivo := holding.HoldingsAtivo[j]
			trimestreMap[buildKey(holdingAtivo.Ativo, holdingAtivo.Trimestre)] = holdingAtivo
		}
	}

	/**
		Busca o ativo do trimestre anterior e calcula a diferença.
	    Se não existir no trimestre anterior, salva 0%

	 */

	insights, err := s.buildInsights(trimestreMap, insights)
	if err != nil {
		return err
	}

	err = s.salvarInsights(ctx, err, user, insights)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) buildInsights(trimestreMap map[string]holding_domain.HoldingAtivo, insights []insight_domain.Insight) ([]insight_domain.Insight, error) {
	log.Print("Iniciando o processamento de insights")
	quarters, err := s.quarterService.GetQuarters()
	if err != nil {
		log.Print("Ocorreu um erro ao buscar os trimestres")
		return nil, err
	}

	for _, holdingAtivo := range trimestreMap {
		currentQuarter := holdingAtivo.Trimestre
		lastQuarter, err := s.getLastQuarter(currentQuarter, quarters)
		if err != nil {
			return nil, err
		}
		lastQuarterHolding := trimestreMap[buildKey(holdingAtivo.Ativo, lastQuarter.Id)]

		insight := insight_domain.Insight{
			Id:           0,
			IdTrimestre:  currentQuarter,
			Usuario:      holdingAtivo.Usuario,
			IdAtivo:      holdingAtivo.Ativo.Id,
		}

		if lastQuarter.Id > 0 {
			s.calculateDelta(&insight, holdingAtivo, lastQuarterHolding)
		}

		insights = append(insights, insight)
	}
	return insights, nil
}

func (s Service) calculateDelta(insight *insight_domain.Insight, holdingAtivo holding_domain.HoldingAtivo, lastQuarterHolding holding_domain.HoldingAtivo) {
	if lastQuarterHolding.ReceitaLiquida > 0 {
		insight.ReceitaDelta = internal.RoundFloat((float32(holdingAtivo.ReceitaLiquida) - float32(lastQuarterHolding.ReceitaLiquida)) / float32(lastQuarterHolding.ReceitaLiquida))
	}

	if lastQuarterHolding.Ebitda > 0 {
		insight.EbitdaDelta = internal.RoundFloat((float32(holdingAtivo.Ebitda) - float32(lastQuarterHolding.Ebitda)) / float32(lastQuarterHolding.Ebitda))
	}

	if lastQuarterHolding.LucroLiquido > 0 {
		insight.LucroDelta = internal.RoundFloat((float32(holdingAtivo.LucroLiquido) - float32(lastQuarterHolding.LucroLiquido)) / float32(lastQuarterHolding.LucroLiquido))
	}

	if lastQuarterHolding.DividaLiquida > 0 {
		insight.DividaDelta = internal.RoundFloat((float32(holdingAtivo.DividaLiquida) - float32(lastQuarterHolding.DividaLiquida)) / float32(lastQuarterHolding.DividaLiquida))
	}
}

func (s Service) salvarInsights(ctx context.Context, err error, user int64, insights []insight_domain.Insight) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("Erro ao iniciar transacao")
	}
	//deleta todos insights
	err = s.repository.DeleteByUser(ctx, tx, user)
	if err != nil {
		log.Print("Erro ao deletar insights")
		tx.Rollback()
		return err
	}
	//insere novos insights
	for i := 0; i < len(insights); i++ {
		err := s.repository.SaveInsights(ctx, tx, insights[i])
		if err != nil {
			log.Print("Erro ao persistir insights")
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func (s Service) getLastQuarter(quarter int64, quarters []quarter_domain.Trimestre) (quarter_domain.Trimestre, error) {
	currentQuarter, err := s.quarterService.GetQuarter(quarter)
	if err != nil {
		log.Print("Ocorreu um erro ao buscar o quarter")
		return quarter_domain.Trimestre{}, err
	}
	for i := 0; i < len(quarters); i++ {
		if currentQuarter.TrimestreAnterior == quarters[i].Id {
			return quarters[i], nil
		}
	}
	return quarter_domain.Trimestre{}, nil
}

func buildKey(ativo asset_domain.Asset, trimestre int64) string {
	return ativo.Codigo + "-" +strconv.FormatInt(trimestre, 10)
}